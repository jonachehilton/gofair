package streaming

import (
	"crypto/tls"

	"github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair/streaming/models"
)

type StreamChannels struct {
	// Outgoing Requests
	MarketSubscriptionRequest chan models.MarketSubscriptionMessage
	OrderSubscriptionRequest  chan models.OrderSubscriptionMessage

	// Incoming Responses
	MarketUpdate chan MarketBook

	// TODO: Fix this
	OrderUpdate                chan interface{}
	MarketSubscriptionResponse chan MarketSubscriptionResponse
	Error                      chan error
}

type streamRequests struct {
	
}

type StreamResponses struct {

}

func newStreamChannels() *StreamChannels {

	channels := new(StreamChannels)

	// Set up Outgoing Request Channels
	channels.MarketSubscriptionRequest = make(chan models.MarketSubscriptionMessage, 64)
	channels.OrderSubscriptionRequest = make(chan models.OrderSubscriptionMessage, 1)

	// Set up Incoming Response Channels
	channels.MarketUpdate = make(chan MarketBook, 64)
	channels.MarketSubscriptionResponse = make(chan MarketSubscriptionResponse, 64)
	channels.Error = make(chan error)

	return channels
}

type Stream struct {

	channels *StreamChannels
	session *Session
	eventHandler *EventHandler

	// Logger instance for Stream object
	log *logrus.Logger
}

// NewStreamClient blah blah
func NewStream(endpoint string, log *logrus.Logger, certs *tls.Certificate) (*Stream, error) {

	if endpoint != StreamEndpoint && endpoint != StreamIntegrationEndpoint {
		return nil, &EndpointError{}
	}

	stream := new(Stream)
	stream.log = log
	stream.channels = NewStreamChannels()
	stream.eventHandler = NewEventHandler()

	return stream, nil
}

// Start performs the Connection and Authentication steps and initializes the read/write goroutines
func (stream *Stream) Start() error {

	err := stream.connect()
	if err != nil {
		return err
	}

	err = stream.authenticate()
	if err != nil {
		return err
	}

	go stream.readPump(errChan)
	go stream.writePump(errChan)

	return nil
}

// Stop closes the connection and kills the associated read/write goroutines
func (l *Stream) Stop() error {

	l.killChannel <- 1

	err := l.conn.Close()

	if err != nil {
		return err
	}

	return nil
}

func (stream *Stream) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	marketSubscriptionRequest := &models.MarketSubscriptionMessage{MarketFilter: marketFilter, MarketDataFilter: marketDataFilter}
	marketSubscriptionRequest.SetID(stream.uid)

	stream.channels.MarketSubscriptionRequest <- *marketSubscriptionRequest

	orderSubscriptionRequest := &models.OrderSubscriptionMessage{SegmentationEnabled: true}
	orderSubscriptionRequest.SetID(stream.uid)

	stream.channels.OrderSubscriptionRequest <- *orderSubscriptionRequest

	stream.uid++
}