package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type orderHandler struct {
	cache      CachedOrders
	channels   *StreamChannels
	initialClk string
	clk        string
}

func newOrderHandler(channels *StreamChannels, orderCache *CachedOrders) *orderHandler {
	orderStream := new(orderHandler)
	orderStream.cache = *orderCache
	orderStream.channels = channels
	return orderStream
}

func (handler *orderHandler) OnSubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (handler *orderHandler) OnResubscribe(orderChangeMessage models.OrderChangeMessage) {

}

func (orderHandler *orderHandler) OnHeartbeat(orderChangeMessage models.OrderChangeMessage) {

}

func (handler *orderHandler) OnUpdate(orderChangeMessage models.OrderChangeMessage) {

	if handler.initialClk == "" {
		handler.initialClk = orderChangeMessage.Clk
	}

	handler.clk = orderChangeMessage.Clk

	for _, orderMarketChange := range orderChangeMessage.Oc {

		// Check if a cache for the given Market ID exists or if it's a Full Image (Notification to replace item in cache as opposed to just updating it)
		// https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Exchange+Stream+API#ExchangeStreamAPI-OrderSubscriptionMessage

		orderBookCache, found := handler.cache[orderMarketChange.ID]
		if !found || orderMarketChange.FullImage {
			orderBookCache = newOrderBookCache()
			orderBookCache.MarketID = orderMarketChange.ID
			orderBookCache.LastPublishTime = orderChangeMessage.Pt
			handler.cache[orderMarketChange.ID] = orderBookCache
		}

		orderBookCache.update(orderMarketChange, orderChangeMessage.Pt)
		handler.channels.OrderUpdate <- *orderBookCache.Snap()
	}
}
