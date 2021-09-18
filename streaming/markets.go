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

func (ms *MarketEventHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		marketCache := CreateMarketCache(&changeMessage, marketChange)
		ms.cache[marketChange.ID] = *marketCache
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	ms.channels.MarketSubscriptionResponse <- *response
}

func (ms *MarketEventHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {

	response := new(MarketSubscriptionResponse)

	for _, marketChange := range changeMessage.Mc {
		response.SubscribedMarketIDs = append(response.SubscribedMarketIDs, marketChange.ID)
	}

	ms.channels.MarketSubscriptionResponse <- *response
}

func (ms *MarketEventHandler) OnHeartbeat(changeMessage models.MarketChangeMessage) {
}

func (ms *MarketEventHandler) OnUpdate(changeMessage models.MarketChangeMessage) {

	if ms.initialClk == "" {
		ms.initialClk = changeMessage.Clk
	}

	ms.clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := ms.cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			ms.channels.MarketUpdate <- marketCache.Snap()
		} else {
			marketCache := CreateMarketCache(&changeMessage, marketChange)
			ms.cache[marketChange.ID] = *marketCache
			ms.channels.MarketUpdate <- marketCache.Snap()
		}
	}
}
