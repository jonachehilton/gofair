package streaming

import (
	"github.com/belmegatron/gofair/streaming/models"
)

type MarketEventHandler struct {
	counter    int32
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

func (eh *MarketEventHandler) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	request := models.MarketSubscriptionMessage{MarketFilter: marketFilter, MarketDataFilter: marketDataFilter}
	request.SetID(eh.counter)
	eh.counter++

	eh.channels.marketSubscriptionRequest <- request
}

func (ms *MarketEventHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {

	for _, marketChange := range changeMessage.Mc {
		marketCache := CreateMarketCache(&changeMessage, marketChange)
		ms.cache[marketChange.ID] = *marketCache
	}
}

func (ms *MarketEventHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {
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
