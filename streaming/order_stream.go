package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
	"github.com/sirupsen/logrus"
)

type OrderStreamHandler struct {
	cache      OrderCache
	stream     *Stream
	log        *logrus.Logger
	InitialClk string
	Clk        string
}

func NewOrderStream(stream *Stream) *OrderStreamHandler {
	orderStream := new(OrderStreamHandler)
	orderStream.cache = make(OrderCache)
	orderStream.stream = stream
	return orderStream
}

func (orderStream *OrderStreamHandler) Subscribe(orderSubscription models.OrderSubscriptionMessage) {
	orderStream.listener.orderSubscriptionRequest <- orderSubscription
}

func (orderStream *OrderStreamHandler) OnSubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (orderStream *OrderStreamHandler) OnResubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (orderStream *OrderStreamHandler) OnHeartbeat(orderChangeMessage models.OrderChangeMessage) {

}

func (orderStream *OrderStreamHandler) OnUpdate(orderChangeMessage models.OrderChangeMessage) {

	if orderStream.InitialClk == "" {
		orderStream.InitialClk = orderChangeMessage.Clk
	}

	orderStream.Clk = orderChangeMessage.Clk

	for _, orderMarketChange := range orderChangeMessage.Oc {

		// Check if a cache for the given Market ID exists
		_, found := orderStream.cache[orderMarketChange.ID]
		if !found {
			orderStream.cache[orderMarketChange.ID] = MarketOrderCache{MarketID: orderMarketChange.ID}
		}

		// Check if a cache for the given MarketID/RunnerID combination exists
		for _, orderRunnerChange := range orderMarketChange.Orc {

			_, found = orderStream.cache[orderMarketChange.ID].RunnerOrders[int(orderRunnerChange.ID)]
			if !found {
				orderStream.cache[orderMarketChange.ID].RunnerOrders[int(orderRunnerChange.ID)] = RunnerOrderCache{}
			}

			// TODO: Update Cache

		}
	}
}
