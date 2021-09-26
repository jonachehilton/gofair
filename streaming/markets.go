package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type marketEventHandler struct {
	channels   *StreamChannels
	cache      CachedMarkets
	initialClk string
	clk        string
}

func newMarketHandler(channels *StreamChannels, marketCache *CachedMarkets) *marketEventHandler {
	marketStream := new(marketEventHandler)
	marketStream.channels = channels
	marketStream.cache = *marketCache
	return marketStream
}

func (handler *marketEventHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		marketCache := newMarketCache(&changeMessage, marketChange)
		handler.cache[marketChange.ID] = marketCache
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	handler.channels.MarketSubscriptionResponse <- *response
}

func (handler *marketEventHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	handler.channels.MarketSubscriptionResponse <- *response
}

func (handler *marketEventHandler) OnHeartbeat(changeMessage models.MarketChangeMessage) {
}

func (handler *marketEventHandler) OnUpdate(changeMessage models.MarketChangeMessage) {

	if handler.initialClk == "" {
		handler.initialClk = changeMessage.Clk
	}

	handler.clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := handler.cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			handler.channels.MarketUpdate <- marketCache.Snap()
		} else {
			marketCache := newMarketCache(&changeMessage, marketChange)
			handler.cache[marketChange.ID] = marketCache
			handler.channels.MarketUpdate <- marketCache.Snap()
		}
	}
}
