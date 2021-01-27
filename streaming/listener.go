package streaming

import (
	"bufio"
	"crypto/tls"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair"
	"github.com/belmegatron/gofair/streaming/models"
)

// ListenerFactory creates a Listener struct
func ListenerFactory(client *gofair.Client) *Listener {
	l := new(Listener)
	l.client = client
	l.subscribeChannel = make(chan models.MarketSubscriptionMessage, 64)
	l.killChannel = make(chan int, 64)
	l.ResultsChannel = make(chan MarketBook, 64)
	l.ErrorChannel = make(chan error, 64)
	l.addMarketStream()
	l.addOrderStream()

	return l
}

func (l *Listener) Start(errChan *chan error) error {

	success, err := l.connect()

	if err != nil {
		return err
	}

	if success {
		err = l.authenticate()

		if err != nil {
			return err
		}

		go l.readPump(errChan)
		go l.writePump(errChan)
	}

	return nil
}

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
	conn         *tls.Conn
	client       *gofair.Client

	// Public
	MarketStream *MarketStream
	OrderStream  Stream

	// Channels for IPC
	subscribeChannel chan models.MarketSubscriptionMessage
	killChannel      chan int
	ErrorChannel     chan error
	// TODO: change to interface so that OrderBook can be accepted
	ResultsChannel chan MarketBook
}

func (l *Listener) addMarketStream() {
	l.MarketStream = new(MarketStream)
	l.MarketStream.OutputChannel = l.ResultsChannel
	l.MarketStream.Cache = make(map[string]MarketCache)
}

func (l *Listener) addOrderStream() {}

func (l *Listener) connect() (bool, error) {

	var success bool = false

	cfg := &tls.Config{Certificates: []tls.Certificate{*l.client.Certificates}}
	conn, err := tls.Dial("tcp", gofair.Endpoints.StreamIntegration, cfg)
	c := bufio.NewReader(conn)

	if err == nil {

		buf, _, err := c.ReadLine()
		if err != nil {
			return success, err
		}

		connectionMessage := new(models.ConnectionMessage)
		err = connectionMessage.UnmarshalJSON(buf)
		if err != nil {
			return success, err
		}

		if connectionMessage.ConnectionID != "" {
			success = true
			l.connectionID = connectionMessage.ConnectionID
			l.conn = conn
			log.Debug("BetfairStreamAPI - Connected")
		} else {
			err := new(ConnectionError)
			return success, err
		}

	}

	return success, nil
}

func (l *Listener) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	request := new(models.MarketSubscriptionMessage)
	request.SetOp("marketSubscription")
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

func (l *Listener) authenticate() error {

	if l.conn == nil {
		return new(NoConnectionError)
	}

	c := bufio.NewReader(l.conn)

	authenticationMessage := new(models.AuthenticationMessage)
	authenticationMessage.SetOp("authentication")
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

	buf, _, err := c.ReadLine()
	if err != nil {
		return err
	}

	statusMessage := new(models.StatusMessage)
	err = statusMessage.UnmarshalJSON(buf)
	if err != nil {
		return err
	}

	if statusMessage.StatusCode == "FAILURE" {
		authenticationError := new(AuthenticationError)
		err = authenticationError
		log.WithFields(log.Fields{
			"errorCode":    statusMessage.ErrorCode,
			"errorMessage": statusMessage.ErrorMessage,
		}).Error("Betfair Stream API - Failed to Authenticate")
		return err
	}

	log.Debug("Betfair Stream API - Authenticated")

	return nil
}

func (l *Listener) readPump(errChan *chan error) {

	if l.conn == nil {
		err := new(NoConnectionError)
		*errChan <- err
		return
	}

	c := bufio.NewReader(l.conn)

	for {
		select {
		case <-l.killChannel:
			return
		default:
			marketChangeMessage := new(models.MarketChangeMessage)
			// TODO: Handle a disconnect and resubscribe
			buf, _, err := c.ReadLine()
			err = marketChangeMessage.UnmarshalJSON(buf)
			if err == nil {
				l.onData(*marketChangeMessage)
			}
			if err != nil {
				*errChan <- err
				return
			}
		}
	}
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

func (l *Listener) writePump(errChan *chan error) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		l.conn.Close()
	}()
	for {
		select {
		case <-l.killChannel:
			return
		case marketSubscriptionMessage := <-l.subscribeChannel:
			b, err := marketSubscriptionMessage.MarshalJSON()
			if err == nil {
				_, err = l.write(b)
			}
			if err != nil {
				*errChan <- err
				return
			}
		case <-ticker.C:
			_, err := l.write([]byte{})
			if err != nil {
				*errChan <- err
				return
			}
		}
	}
}

func (l *Listener) onData(ChangeMessage models.MarketChangeMessage) {

	switch ChangeMessage.Op() {
	case "connection":
		l.onConnection(ChangeMessage)
	case "status":
		l.onStatus(ChangeMessage)
	case "mcm":
		l.onChangeMessage(l.MarketStream, ChangeMessage)
	case "ocm":
		l.onChangeMessage(l.OrderStream, ChangeMessage)
	}
}

func (l *Listener) onConnection(ChangeMessage models.MarketChangeMessage) {
	log.Debug("BetfairStreamAPI - Connected")
}

func (l *Listener) onStatus(ChangeMessage models.MarketChangeMessage) {
	log.Debug("BetfairStreamAPI - Status Message Received")
}

func (l *Listener) onChangeMessage(Stream Stream, ChangeMessage models.MarketChangeMessage) {
	switch ChangeMessage.Ct {
	case "SUB_IMAGE":
		Stream.OnSubscribe(ChangeMessage)
	case "RESUB_DELTA":
		Stream.OnResubscribe(ChangeMessage)
	case "HEARTBEAT":
		Stream.OnHeartbeat(ChangeMessage)
	default:
		Stream.OnUpdate(ChangeMessage)
	}
}
