package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type MarketEventHandler struct {
	channels   *StreamChannels
	cache      map[string]MarketCache
	initialClk string
	clk        string
}

func NewMarketHandler(channels *StreamChannels) *MarketEventHandler {
	marketStream := new(MarketEventHandler)
	marketStream.channels = channels
	marketStream.cache = make(map[string]MarketCache)
	return marketStream
}

func (handler *MarketEventHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		marketCache := CreateMarketCache(&changeMessage, marketChange)
		handler.cache[marketChange.ID] = *marketCache
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	handler.channels.MarketSubscriptionResponse <- *response
}

func (handler *MarketEventHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	handler.channels.MarketSubscriptionResponse <- *response
}

func (handler *MarketEventHandler) OnHeartbeat(changeMessage models.MarketChangeMessage) {
}

func (handler *MarketEventHandler) OnUpdate(changeMessage models.MarketChangeMessage) {

	if handler.initialClk == "" {
		handler.initialClk = changeMessage.Clk
	}

	handler.clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := handler.cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			handler.channels.MarketUpdate <- marketCache.Snap()
		} else {
			marketCache := CreateMarketCache(&changeMessage, marketChange)
			handler.cache[marketChange.ID] = *marketCache
			handler.channels.MarketUpdate <- marketCache.Snap()
		}
	}
}
