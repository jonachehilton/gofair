package streaming

import (
	"crypto/tls"

	"github.com/belmegatron/gofair/streaming/models"
)

type StreamChannels struct {
	// Outgoing Requests
	marketSubscriptionRequest chan models.MarketSubscriptionMessage
	orderSubscriptionRequest  chan models.OrderSubscriptionMessage

	// Incoming Responses
	MarketUpdate chan MarketBook

	// TODO: Fix this
	OrderUpdate                chan interface{}
	MarketSubscriptionResponse chan MarketSubscriptionResponse
	Err                        chan error
}

type streamRequests struct {
}

type StreamResponses struct {
}

func newStreamChannels() *StreamChannels {

	channels := new(StreamChannels)

	// Set up Outgoing Request Channels
	channels.marketSubscriptionRequest = make(chan models.MarketSubscriptionMessage, 64)
	channels.orderSubscriptionRequest = make(chan models.OrderSubscriptionMessage, 1)

	// Set up Incoming Response Channels
	channels.MarketUpdate = make(chan MarketBook, 64)
	channels.MarketSubscriptionResponse = make(chan MarketSubscriptionResponse, 64)
	channels.Err = make(chan error)

	return channels
}

type Stream struct {
	endpoint string
	certs    *tls.Certificate
	Channels *StreamChannels
	session  *Session
}

// NewStreamClient blah blah
func NewStream(endpoint string, certs *tls.Certificate) (*Stream, error) {

	if endpoint != StreamEndpoint && endpoint != StreamIntegrationEndpoint {
		return nil, &EndpointError{}
	}

	stream := new(Stream)
	stream.endpoint = endpoint
	stream.certs = certs
	stream.Channels = newStreamChannels()

	return stream, nil
}

// Start performs the Connection and Authentication steps and initializes the read/write goroutines
func (stream *Stream) Start() error {

	session, err := NewSession(stream.endpoint, stream.certs)

	return nil
}

// Stop tears down the underlying TLS session to the Streaming endpoint
func (stream *Stream) Stop() {
	stream.session.Stop()
}

func (stream *Stream) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	marketSubscriptionRequest := &models.MarketSubscriptionMessage{MarketFilter: marketFilter, MarketDataFilter: marketDataFilter}
	marketSubscriptionRequest.SetID(stream.uid)

	stream.Channels.marketSubscriptionRequest <- *marketSubscriptionRequest

	orderSubscriptionRequest := &models.OrderSubscriptionMessage{SegmentationEnabled: true}
	stream.session.eventHandler.Orders.Subscribe()

	stream.Channels.orderSubscriptionRequest <- *orderSubscriptionRequest

	stream.uid++
}
