package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type IMarketStream interface {
	Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter)
	OnSubscribe(ChangeMessage models.MarketChangeMessage)
	OnResubscribe(ChangeMessage models.MarketChangeMessage)
	OnHeartbeat(ChangeMessage models.MarketChangeMessage)
	OnUpdate(ChangeMessage models.MarketChangeMessage)
}

type IOrderStream interface {
	OnSubscribe(ChangeMessage models.OrderChangeMessage)
	OnResubscribe(ChangeMessage models.OrderChangeMessage)
	OnHeartbeat(ChangeMessage models.OrderChangeMessage)
	OnUpdate(ChangeMessage models.OrderChangeMessage)
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

func (stream *Stream) onData(op string, data []byte) {

	switch op {
	case "connection":
		stream.onConnection(data)
	case "status":
		stream.onStatus(data)
	case "mcm":
		stream.onMarketChangeMessage(stream.MarketStream, data)
	case "ocm":
		stream.onOrderChangeMessage(stream.OrderStream, data)
	}
}

func (stream *Stream) onConnection(data []byte) {
	stream.log.Debug("Connected")
}

func (stream *Stream) onStatus(data []byte) {
	stream.log.Debug("Status Message Received")
}

func (stream *Stream) onMarketChangeMessage(Stream IMarketStream, data []byte) {

	marketChangeMessage := new(models.MarketChangeMessage)

	err := marketChangeMessage.UnmarshalJSON(data)
	if err != nil {
		stream.log.Error("Failed to unmarshal MarketChangeMessage.")
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

func (stream *Stream) onOrderChangeMessage(Stream IOrderStream, data []byte) {

	orderChangeMessage := new(models.OrderChangeMessage)

	err := orderChangeMessage.UnmarshalJSON(data)
	if err != nil {
		stream.log.Error("Failed to unmarshal OrderChangeMessage.")
		return
	}

	switch orderChangeMessage.Ct {
	case "SUB_IMAGE":
		Stream.OnSubscribe(*orderChangeMessage)
	case "RESUB_DELTA":
		Stream.OnResubscribe(*orderChangeMessage)
	case "HEARTBEAT":
		Stream.OnHeartbeat(*orderChangeMessage)
	default:
		Stream.OnUpdate(*orderChangeMessage)
	}
}