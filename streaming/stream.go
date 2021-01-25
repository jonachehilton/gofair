package streaming

import (
	log "github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair/streaming/models"
)

type Stream interface {
	OnSubscribe(ChangeMessage models.MarketChangeMessage)
	OnResubscribe(ChangeMessage models.MarketChangeMessage)
	OnHeartbeat(ChangeMessage models.MarketChangeMessage)
	OnUpdate(ChangeMessage models.MarketChangeMessage)
}

type MarketStream struct {
	OutputChannel chan MarketBook
	Cache         map[string]MarketCache
}

func getMarketIDs(mcm models.MarketChangeMessage) []string {
	marketIDs := make([]string, 0)
	for _, marketChange := range mcm.Mc {
		marketIDs = append(marketIDs, marketChange.ID)
	}
	return marketIDs
}

func (ms *MarketStream) OnSubscribe(changeMessage models.MarketChangeMessage) {
	marketIDs := getMarketIDs(changeMessage)
	log.WithFields(log.Fields{
		"marketIDs": marketIDs,
	}).Debug("BetfairStreamAPI - Subscribed to Betfair Market Changes")
}

func (ms *MarketStream) OnResubscribe(changeMessage models.MarketChangeMessage) {
	marketIDs := getMarketIDs(changeMessage)
	log.WithFields(log.Fields{
		"marketIDs": marketIDs,
	}).Debug("BetfairStreamAPI - Resubscribed to Betfair Market Changes")
}

func (ms *MarketStream) OnHeartbeat(changeMessage models.MarketChangeMessage) {
	log.Debug("BetfairStreamAPI - Heartbeat")
}

func (ms *MarketStream) OnUpdate(changeMessage models.MarketChangeMessage) {
	// todo update clk/initialClk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := ms.Cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			ms.OutputChannel <- marketCache.Snap()
		} else {
			marketCache := CreateMarketCache(&changeMessage, marketChange)
			ms.Cache[marketChange.ID] = *marketCache
			ms.OutputChannel <- marketCache.Snap()

			log.WithFields(log.Fields{
				"marketID": marketChange.ID,
			}).Debug("BetfairStreamAPI - Created new market cache")
		}
	}
}
