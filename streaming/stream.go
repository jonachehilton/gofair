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

func NewStreamChannels() *StreamChannels {

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

	// Unique ID that must be assigned to a listener
	uid          int32
	connectionID string

	channels *StreamChannels

	// Live/Testing Streaming API endpoint
	endpoint string

	// Logger instance for listener object
	log *logrus.Logger

	MarketStream IMarketStream
	OrderStream  IOrderStream
}

// NewStreamClient blah blah
func NewStream(endpoint string, log *logrus.Logger, certs *tls.Certificate) (*Stream, error) {

	if endpoint != StreamEndpoint && endpoint != StreamIntegrationEndpoint {
		return nil, &EndpointError{}
	}

	stream := new(Stream)
	stream.endpoint = endpoint
	stream.log = log
	stream.channels = NewStreamChannels()

	stream.MarketStream = NewMarketStream(stream, stream.log, &stream.IncomingMarketData)
	stream.OrderStream = NewOrderStream(stream, stream.log)

	return stream, nil
}

// Start performs the Connection and Authentication steps and initializes the read/write goroutines
func (stream *Stream) Start(errChan *chan error) error {

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
