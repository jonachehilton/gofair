package streaming

import "github.com/belmegatron/gofair/streaming/models"

type OrderBookCache struct {
	MarketID        string
	LastPublishTime int64
	Runners         map[int64]*models.OrderRunnerChange
	Closed          bool
}

func newOrderBookCache() *OrderBookCache {
	cache := new(OrderBookCache)
	cache.Runners = make(map[int64]*models.OrderRunnerChange)
	return cache
}

func (cache *OrderBookCache) update(data *models.OrderMarketChange, publishTime int64) {

	cache.LastPublishTime = publishTime

	if data.Closed {
		cache.Closed = true
	}

	// Enumerate over Runners
	for _, orderChange := range data.Orc {

		// Check if this is a Full Image (and therefore need to replace our cached entry) or if it hasn't already been cached
		runner, found := cache.Runners[orderChange.ID]

		if orderChange.FullImage || !found {
			cache.Runners[orderChange.ID] = orderChange
		} else {
			// Update Matched Backs
			cache.Runners[orderChange.ID].Mb = append(cache.Runners[orderChange.ID].Mb, runner.Mb...)

			// Update Matched Lays
			cache.Runners[orderChange.ID].Ml = append(cache.Runners[orderChange.ID].Ml, runner.Ml...)

			// Updated Unmatched Orders
			cache.Runners[orderChange.ID].Uo = append(cache.Runners[orderChange.ID].Uo, runner.Uo...)
		}
	}
}

func (cache *OrderBookCache) Snap() *OrderBookCache {
	return cache
}

// CachedOrders maps MarketID's to Orders on the Exchange
type CachedOrders map[string]*OrderBookCache
