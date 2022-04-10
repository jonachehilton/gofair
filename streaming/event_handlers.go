package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

const (
	// Ops
	connection          = "connection"
	status              = "status"
	marketChangeMessage = "mcm"
	orderChangeMessage  = "ocm"

	// Change types
	subscribe   = "SUB_IMAGE"
	resubscribe = "RESUB_DELTA"
	heartbeat   = "HEARTBEAT"
)

type IMarketHandler interface {
	OnSubscribe(ChangeMessage models.MarketChangeMessage)
	OnResubscribe(ChangeMessage models.MarketChangeMessage)
	OnHeartbeat(ChangeMessage models.MarketChangeMessage)
	OnUpdate(ChangeMessage models.MarketChangeMessage)
}

type IOrderHandler interface {
	OnSubscribe(ChangeMessage models.OrderChangeMessage)
	OnResubscribe(ChangeMessage models.OrderChangeMessage)
	OnHeartbeat(ChangeMessage models.OrderChangeMessage)
	OnUpdate(ChangeMessage models.OrderChangeMessage)
}

type eventHandler struct {
	Markets  IMarketHandler
	Orders   IOrderHandler
	channels *StreamChannels
}

func newEventHandler(channels *StreamChannels, marketCache *CachedMarkets, orderCache *CachedOrders) *eventHandler {
	handler := new(eventHandler)
	handler.channels = channels
	handler.Markets = newMarketHandler(channels, marketCache)
	handler.Orders = newOrderHandler(channels, orderCache)
	return handler
}

// onData passes a blob to the appropriate event handler based on the op code
func (eh *eventHandler) onData(op string, data []byte) {

	switch op {
	case connection:
		eh.onConnection(data)
	case status:
		eh.onStatus(data)
	case marketChangeMessage:
		eh.onMarketChangeMessage(data)
	case orderChangeMessage:
		eh.onOrderChangeMessage(data)
	}
}

func (stream *eventHandler) onConnection(data []byte) {
}

func (stream *eventHandler) onStatus(data []byte) {

	statusMessage := new(models.StatusMessage)
	err := statusMessage.UnmarshalJSON(data)
	if err != nil {
		return
	}

	stream.channels.Status <- *statusMessage
}

// onMarketChangeMessage passes a MarketChange blob to the appropriate event handler based on the Change type
func (eh *eventHandler) onMarketChangeMessage(data []byte) {

	marketChangeMessage := new(models.MarketChangeMessage)

	err := marketChangeMessage.UnmarshalJSON(data)
	if err != nil {
		return
	}

	switch marketChangeMessage.Ct {
	case subscribe:
		eh.Markets.OnSubscribe(*marketChangeMessage)
	case resubscribe:
		eh.Markets.OnResubscribe(*marketChangeMessage)
	case heartbeat:
		eh.Markets.OnHeartbeat(*marketChangeMessage)
	default:
		eh.Markets.OnUpdate(*marketChangeMessage)
	}
}

// onOrderChangeMessage passes an OrderChange blob to the appropriate event handler based on the Change type
func (eh *eventHandler) onOrderChangeMessage(data []byte) {

	orderChangeMessage := new(models.OrderChangeMessage)

	err := orderChangeMessage.UnmarshalJSON(data)
	if err != nil {
		return
	}

	switch orderChangeMessage.Ct {
	case subscribe:
		eh.Orders.OnSubscribe(*orderChangeMessage)
	case resubscribe:
		eh.Orders.OnResubscribe(*orderChangeMessage)
	case heartbeat:
		eh.Orders.OnHeartbeat(*orderChangeMessage)
	default:
		eh.Orders.OnUpdate(*orderChangeMessage)
	}
}
