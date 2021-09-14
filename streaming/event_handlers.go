package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
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

type EventHandler struct {
	Markets IMarketHandler
	Orders IOrderHandler
}

func NewEventHandler() *EventHandler {
	eh := new(EventHandler)
	return eh
}

func (eh *EventHandler) onData(op string, data []byte) {

	switch op {
	case "connection":
		eh.onConnection(data)
	case "status":
		eh.onStatus(data)
	case "mcm":
		eh.onMarketChangeMessage(data)
	case "ocm":
		eh.onOrderChangeMessage(data)
	}
}

func (stream *EventHandler) onConnection(data []byte) {
}

func (stream *EventHandler) onStatus(data []byte) {
}

func (eh *EventHandler) onMarketChangeMessage(data []byte) {

	marketChangeMessage := new(models.MarketChangeMessage)

	err := marketChangeMessage.UnmarshalJSON(data)
	if err != nil {
		return
	}

	switch marketChangeMessage.Ct {
	case "SUB_IMAGE":
		eh.Markets.OnSubscribe(*marketChangeMessage)
	case "RESUB_DELTA":
		eh.Markets.OnResubscribe(*marketChangeMessage)
	case "HEARTBEAT":
		eh.Markets.OnHeartbeat(*marketChangeMessage)
	default:
		eh.Markets.OnUpdate(*marketChangeMessage)
	}
}

func (eh *EventHandler) onOrderChangeMessage(data []byte) {

	orderChangeMessage := new(models.OrderChangeMessage)

	err := orderChangeMessage.UnmarshalJSON(data)
	if err != nil {
		return
	}

	switch orderChangeMessage.Ct {
	case "SUB_IMAGE":
		eh.Orders.OnSubscribe(*orderChangeMessage)
	case "RESUB_DELTA":
		eh.Orders.OnResubscribe(*orderChangeMessage)
	case "HEARTBEAT":
		eh.Orders.OnHeartbeat(*orderChangeMessage)
	default:
		eh.Orders.OnUpdate(*orderChangeMessage)
	}
}