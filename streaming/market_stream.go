package streaming

import (
	"github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair/streaming/models"
)

type MarketStreamHandler struct {
	stream *Stream

	Cache      map[string]MarketCache
	InitialClk string
	Clk        string
}

func NewMarketStream(client *Client) *MarketStreamHandler {
	marketStream := new(MarketStreamHandler)
	marketStream.Cache = make(map[string]MarketCache)
	return marketStream
}

func (ms *MarketStreamHandler) Subscribe(marketFilter *models.MarketFilter, marketDataFilter *models.MarketDataFilter) {

	request := new(models.MarketSubscriptionMessage)
	request.SetID(ms.stream.uid)
	ms.stream.uid++
	request.MarketFilter = marketFilter
	request.MarketDataFilter = marketDataFilter

	ms.stream.channels.MarketSubscriptionRequest <- *request
}

func getMarketIDs(mcm models.MarketChangeMessage) []string {
	marketIDs := make([]string, 0)
	for _, marketChange := range mcm.Mc {
		marketIDs = append(marketIDs, marketChange.ID)
	}
	return marketIDs
}

func (ms *MarketStreamHandler) OnSubscribe(changeMessage models.MarketChangeMessage) {

	marketIDs := getMarketIDs(changeMessage)

	ms.stream.log.WithFields(logrus.Fields{
		"marketIDs": marketIDs,
	}).Debug("Subscribed to Betfair Market Changes")

	for _, marketChange := range changeMessage.Mc {

		marketCache := CreateMarketCache(&changeMessage, marketChange)
		ms.Cache[marketChange.ID] = *marketCache
	}
}

func (ms *MarketStreamHandler) OnResubscribe(changeMessage models.MarketChangeMessage) {
	marketIDs := getMarketIDs(changeMessage)
	ms.stream.log.WithFields(logrus.Fields{
		"marketIDs": marketIDs,
	}).Debug("Resubscribed to Betfair Market Changes")
}

func (ms *MarketStreamHandler) OnHeartbeat(changeMessage models.MarketChangeMessage) {
	ms.stream.log.Debug("Heartbeat")
}

func (ms *MarketStreamHandler) OnUpdate(changeMessage models.MarketChangeMessage) {

	if ms.InitialClk == "" {
		ms.InitialClk = changeMessage.Clk
	}

	ms.Clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := ms.Cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			ms.stream.channels.MarketUpdate <- marketCache.Snap()
		} else {
			marketCache := CreateMarketCache(&changeMessage, marketChange)
			ms.Cache[marketChange.ID] = *marketCache
			ms.stream.channels.MarketUpdate <- marketCache.Snap()
		}
	}
}
