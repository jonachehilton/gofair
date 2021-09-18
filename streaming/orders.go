package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type OrderHandler struct {
	cache      OrderCache
	stream     *Stream
	InitialClk string
	Clk        string
}

func NewOrderHandler(stream *Stream) *OrderHandler {
	orderStream := new(OrderHandler)
	orderStream.cache = make(OrderCache)
	orderStream.stream = stream
	return orderStream
}

func (handler *OrderHandler) OnSubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (handler *OrderHandler) OnResubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (orderHandler *OrderHandler) OnHeartbeat(orderChangeMessage models.OrderChangeMessage) {

}

func (handler *OrderHandler) OnUpdate(orderChangeMessage models.OrderChangeMessage) {

	if handler.InitialClk == "" {
		handler.InitialClk = orderChangeMessage.Clk
	}

	handler.Clk = orderChangeMessage.Clk

	for _, orderMarketChange := range orderChangeMessage.Oc {

		// Check if a cache for the given Market ID exists
		_, found := handler.cache[orderMarketChange.ID]
		if !found {
			handler.cache[orderMarketChange.ID] = MarketOrderCache{MarketID: orderMarketChange.ID}
		}

		// Check if a cache for the given MarketID/RunnerID combination exists
		for _, orderRunnerChange := range orderMarketChange.Orc {

			_, found = handler.cache[orderMarketChange.ID].RunnerOrders[int(orderRunnerChange.ID)]
			if !found {
				handler.cache[orderMarketChange.ID].RunnerOrders[int(orderRunnerChange.ID)] = RunnerOrderCache{}
			}

			// TODO: Update Cache

		}
	}
}
