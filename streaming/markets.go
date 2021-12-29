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

func (handler *marketEventHandler) onChangeMessage(changeMessage models.MarketChangeMessage) {
	
	if handler.initialClk == "" {
		handler.initialClk = changeMessage.Clk
	}

	handler.clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		var marketCache *MarketCache
		var found bool

		if marketCache, found = handler.cache[marketChange.ID]; found {
			marketCache.UpdateCache(&changeMessage, marketChange)
		} else {
			marketCache = newMarketCache(&changeMessage, marketChange)
			handler.cache[marketChange.ID] = marketCache
		}

		handler.channels.MarketUpdate <- marketCache.Snap()
	}
}

func (handler *marketEventHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {
	handler.onChangeMessage(changeMessage)
}

func (handler *marketEventHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {
	handler.onChangeMessage(changeMessage)
}

func (handler *marketEventHandler) OnHeartbeat(changeMessage models.MarketChangeMessage) {
}

func (handler *marketEventHandler) OnUpdate(changeMessage models.MarketChangeMessage) {
	handler.onChangeMessage(changeMessage)
}
