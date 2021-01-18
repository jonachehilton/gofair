package streaming

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/belmegatron/gofair"
)

// ListenerFactory creates a Listener struct
func ListenerFactory(client *gofair.Client) *Listener {
	l := new(Listener)
	l.Client = client
	l.SubscribeChannel = make(chan MarketSubscriptionRequest, 64)
	l.addMarketStream()
	l.addOrderStream()
	return l
}

type Listener struct {
	UniqueID         int64
	Conn             *tls.Conn
	Client           *gofair.Client
	MarketStream     *MarketStream
	OrderStream      Stream
	SubscribeChannel chan MarketSubscriptionRequest
	OutputChannel    chan MarketBook // TODO: change to interface so that OrderBook can be accepted
}

func (l *Listener) addMarketStream() {
	l.MarketStream = new(MarketStream)
	l.MarketStream.OutputChannel = l.OutputChannel
	l.MarketStream.Cache = make(map[string]MarketCache)
}

func (l *Listener) addOrderStream() {}

func (l *Listener) connect() error {

	cfg := &tls.Config{Certificates: []tls.Certificate{*l.Client.Certificates}}
	conn, err := tls.Dial("tcp", gofair.Endpoints.StreamIntegration, cfg)
	if err != nil {
		return err
	}

	l.Conn = conn

	return nil
}

func (l *Listener) Subscribe(marketFilter *gofair.MarketFilter) {

	request := new(MarketSubscriptionRequest)
	request.OP = "marketSubscription"
	request.ID = l.UniqueID
	l.UniqueID++
	request.MarketFilter = *marketFilter
	request.MarketDataFilter.Fields = []string{"EX_BEST_OFFERS"}
	l.SubscribeChannel <- *request
}

func (l *Listener) authenticate() error {
	msg := new(AuthRequest)
	msg.OP = "authentication"
	msg.ID = l.UniqueID
	msg.AppKey = l.Client.Config.AppKey
	msg.SessionToken = l.Client.Session.SessionToken

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

func (l *Listener) readPump(errChan *chan error) {

	if l.Conn == nil {
		err := new(NoConnectionError)
		*errChan <- err
		return
	}

	for {

		var res []byte
		cb, err := l.Conn.Read(res)
		if err != nil {
			*errChan <- err
			return
		}

		if cb > 0 {
			msg := new(MarketChangeMessage)
			err = json.Unmarshal(res, msg)

			if err != nil {
				*errChan <- err
				return
			}

			l.onData(*msg)
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
		l.Conn.Close()
	}()
	for {
		select {
		case sub := <-l.SubscribeChannel:
			request := new(bytes.Buffer)
			json.NewEncoder(request).Encode(sub)
			b := request.Bytes()
			_, err := l.Conn.Write(b)
			if err != nil {
				*errChan <- err
				return
			}
		case <-ticker.C:
			_, err := l.Conn.Write([]byte{})
			if err != nil {
				*errChan <- err
				return
			}
		}
	}
}

func (l *Listener) onData(ChangeMessage MarketChangeMessage) {
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
