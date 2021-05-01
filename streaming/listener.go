package streaming

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair"
	"github.com/belmegatron/gofair/streaming/models"
)

// ListenerFactory creates a Listener struct
func ListenerFactory(client *gofair.Client, endpoint string, log *logrus.Logger) (*Listener, error) {
	l := new(Listener)
	l.client = client

	l.log = log

	if endpoint != gofair.Endpoints.Stream && endpoint != gofair.Endpoints.StreamIntegration {
		return nil, &EndpointError{}
	}

	l.endpoint = endpoint

	l.subscribeChannel = make(chan models.MarketSubscriptionMessage, 64)
	l.killChannel = make(chan int)
	l.ResultsChannel = make(chan MarketBook, 64)
	l.ErrorChannel = make(chan error)

	l.addMarketStream(l.log)
	l.addOrderStream()

	return l, nil
}

// Start performs the Connection and Authentication steps and initializes the read/write goroutines
func (l *Listener) Start(errChan *chan error) error {

	err := l.connect()

	if err != nil {
		return err
	}

	err = l.authenticate()

	if err != nil {
		return err
	}

	go l.readPump(errChan)
	go l.writePump(errChan)

	return nil
}

// Stop closes the connection and kills the associated read/write goroutines
func (l *Listener) Stop() error {

	l.killChannel <- 1

	err := l.conn.Close()

	if err != nil {
		return err
	}

	return nil
}

type Listener struct {

	// Private
	uniqueID     int32
	connectionID string
	endpoint     string
	log          *logrus.Logger
	conn         *tls.Conn
	client       *gofair.Client
	scanner      bufio.Scanner

	// Public
	MarketStream *MarketStream
	OrderStream  IStream

	// Channels for IPC
	subscribeChannel chan models.MarketSubscriptionMessage
	killChannel      chan int
	ErrorChannel     chan error
	// TODO: change to interface so that OrderBook can be accepted
	ResultsChannel chan MarketBook
}

func (l *Listener) addMarketStream(log *logrus.Logger) {
	l.MarketStream = new(MarketStream)
	l.MarketStream.log = log
	l.MarketStream.OutputChannel = l.ResultsChannel
	l.MarketStream.Cache = make(map[string]MarketCache)
}

func (l *Listener) addOrderStream() {}

func (l *Listener) connect() error {

	cfg := &tls.Config{Certificates: []tls.Certificate{*l.client.Certificates}}
	conn, err := tls.Dial("tcp", l.endpoint, cfg)

	if err != nil {
		return err
	}

	c := bufio.NewReader(conn)
	buf, _, err := c.ReadLine()
	if err != nil {
		return err
	}

	connectionMessage := new(models.ConnectionMessage)
	err = connectionMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if connectionMessage.ConnectionID == "" {
		return &ConnectionError{}
	}

	l.connectionID = connectionMessage.ConnectionID
	l.conn = conn
	// This scanner allows us to keep reading bytes from the connection until we encounter "\r\n"
	l.scanner = *bufio.NewScanner(l.conn)
	l.scanner.Split(ScanCRLF)

	return nil
}

func (l *Listener) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	request := new(models.MarketSubscriptionMessage)
	request.SetID(l.uniqueID)
	l.uniqueID++
	request.MarketFilter = marketFilter
	request.MarketDataFilter = marketDataFilter

	l.subscribeChannel <- *request
}

func (l *Listener) write(b []byte) (int, error) {
	// Every message is in json & terminated with a line feed (CRLF)
	b = append(b, []byte{'\r', '\n'}...)
	return l.conn.Write(b)
}

func (l *Listener) read() ([]byte, error) {

	l.scanner.Scan()

	if err := l.scanner.Err(); err != nil {
		return []byte{}, err
	}

	return l.scanner.Bytes(), nil
}

func (l *Listener) authenticate() error {

	if l.conn == nil {
		return &NoConnectionError{}
	}

	authenticationMessage := new(models.AuthenticationMessage)
	authenticationMessage.SetID(l.uniqueID)
	authenticationMessage.AppKey = l.client.Config.AppKey
	authenticationMessage.Session = l.client.Session.SessionToken

	b, err := authenticationMessage.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = l.write(b)
	if err != nil {
		return err
	}

	buf, err := l.read()
	if err != nil {
		return err
	}

	statusMessage := new(models.StatusMessage)
	err = statusMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if statusMessage.StatusCode == "FAILURE" {
		err := new(AuthenticationError)
		l.log.WithFields(logrus.Fields{
			"errorCode":    statusMessage.ErrorCode,
			"errorMessage": statusMessage.ErrorMessage,
		}).Error("Failed to Authenticate")
		return err
	}

	l.log.Debug("Authenticated")

	return nil
}

// dropCR drops a terminal \r from the data.
func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanCRLF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte{'\r', '\n'}); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func (l *Listener) readPump(errChan *chan error) {

	if l.conn == nil {
		err := new(NoConnectionError)
		*errChan <- err
		return
	}

	for {
		select {
		case <-l.killChannel:
			return
		default:
			buf, err := l.read()
			if err != nil {
				*errChan <- err
				return
			}

			tmp := make(map[string]json.RawMessage)
			var op string
			err = json.Unmarshal(buf, &tmp)
			if err != nil {
				*errChan <- err
				return
			}

			err = json.Unmarshal(tmp["op"], &op)
			if err != nil {
				*errChan <- err
				return
			}

			l.onData(op, buf)
		}
	}
}

func (l *Listener) writePump(errChan *chan error) {
	for {
		select {
		case <-l.killChannel:
			return
		case marketSubscriptionMessage := <-l.subscribeChannel:
			b, err := marketSubscriptionMessage.MarshalJSON()
			if err != nil {
				*errChan <- err
				return
			}
			l.write(b)
		}
	}
}

func (l *Listener) onData(op string, data []byte) {

	switch op {
	case "connection":
		l.onConnection(data)
	case "status":
		l.onStatus(data)
	case "mcm":
		l.onChangeMessage(l.MarketStream, data)
	case "ocm":
		l.onChangeMessage(l.OrderStream, data)
	}
}

func (l *Listener) onConnection(data []byte) {
	l.log.Debug("Connected")
}

func (l *Listener) onStatus(data []byte) {
	l.log.Debug("Status Message Received")
}

func (l *Listener) onChangeMessage(Stream IStream, data []byte) {

	marketChangeMessage := new(models.MarketChangeMessage)
	err := marketChangeMessage.UnmarshalJSON(data)
	if err != nil {
		l.log.Error("Failed to unmarshal MarketChangeMessage.")
		return
	}

	switch marketChangeMessage.Ct {
	case "SUB_IMAGE":
		Stream.OnSubscribe(*marketChangeMessage)
	case "RESUB_DELTA":
		Stream.OnResubscribe(*marketChangeMessage)
	case "HEARTBEAT":
		Stream.OnHeartbeat(*marketChangeMessage)
	default:
		Stream.OnUpdate(*marketChangeMessage)
	}
}
