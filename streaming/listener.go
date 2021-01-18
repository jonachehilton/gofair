package streaming

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
)

const (
	STREAM             = "stream-api.betfair.com:443"
	STREAM_INTEGRATION = "stream-api-integration.betfair.com:443"
)

// ListenerFactory creates a Listener struct
func ListenerFactory(AppKey string, SessionToken string, CertPath string, KeyPath string) *Listener {
	l := new(Listener)
	l.AppKey = AppKey
	l.SessionToken = SessionToken
	l.CertPath = CertPath
	l.KeyPath = KeyPath
	return l
}

type AuthMessage struct {
	OP           string `json:"op"`
	ID           int64  `json:"id"`
	AppKey       string `json:"appKey"`
	SessionToken string `json:"session"`
}

type Listener struct {
	UniqueID      int64
	Conn          *tls.Conn
	AppKey        string
	SessionToken  string
	CertPath      string
	KeyPath       string
	MarketStream  *MarketStream
	OrderStream   Stream
	OutputChannel chan MarketBook //todo change to interface so that OrderBook can be accepted
}

func (l *Listener) AddMarketStream() {
	l.MarketStream = new(MarketStream)
	l.MarketStream.OutputChannel = l.OutputChannel
	l.MarketStream.Cache = make(map[string]MarketCache)
}

func (l *Listener) AddOrderStream() {}

func (l *Listener) Connect() error {

	cert, err := tls.LoadX509KeyPair(l.CertPath, l.KeyPath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	conn, err := tls.Dial("tcp", STREAM_INTEGRATION, cfg)
	if err != nil {
		return err
	}

	l.Conn = conn

	return nil
}

type NoConnectionError struct{}

func (err *NoConnectionError) Error() string {
	return fmt.Sprintf("No stream connection exists.")
}

func (l *Listener) auth() error {
	msg := new(AuthMessage)
	msg.OP = "authentication"
	msg.ID = l.UniqueID
	msg.AppKey = l.AppKey
	msg.SessionToken = l.SessionToken

	if l.Conn == nil {
		err := new(NoConnectionError)
		return err
	}

	request := new(bytes.Buffer)
	json.NewEncoder(request).Encode(msg)

	b := request.Bytes()
	l.Conn.Write(b)

	return nil
}

func (l *Listener) ReadLoop() error {

	if l.Conn == nil {
		err := new(NoConnectionError)
		return err
	}

	for {
		var res []byte
		cb, err := l.Conn.Read(res)
		if err != nil {
			log.Fatal(err)
		}

		if cb > 0 {
			fmt.Println(string(res))
			msg := new(MarketChangeMessage)
			err = json.Unmarshal(res, msg)
			if err != nil {
				log.Fatal(err)
			}

			l.OnData(*msg)
		}
	}

}

func (l *Listener) OnData(ChangeMessage MarketChangeMessage) {
	//todo check unique id
	//todo error handler

	switch ChangeMessage.Operation {
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

func (l *Listener) onConnection(ChangeMessage MarketChangeMessage) {
	// todo
}

func (l *Listener) onStatus(ChangeMessage MarketChangeMessage) {
	// todo
}

func (l *Listener) onChangeMessage(Stream Stream, ChangeMessage MarketChangeMessage) {
	switch ChangeMessage.ChangeType {
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
