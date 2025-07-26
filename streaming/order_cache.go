package streaming

import (
	"github.com/jonachehilton/gofair/streaming/models"
)

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

func (cache *OrderBookCache) updateMatchedLays(selectionID int64, update [][]float64) {

	if cache.Runners[selectionID] == nil {
		cache.Runners[selectionID] = new(models.OrderRunnerChange)
	}

	for _, entry := range update {
		price := entry[0]
		priceFound := false
		// Check if this price exists in our cached lays, if so, replace it.
		for i, cachedEntry := range cache.Runners[selectionID].Ml {
			if cachedEntry[0] == price {
				priceFound = true
				cache.Runners[selectionID].Ml[i] = entry
				break
			}
		}

		if !priceFound {
			cache.Runners[selectionID].Ml = append(cache.Runners[selectionID].Ml, entry)
		}
	}
}

func (cache *OrderBookCache) updateMatchedBacks(selectionID int64, update [][]float64) {

	if cache.Runners[selectionID] == nil {
		cache.Runners[selectionID] = new(models.OrderRunnerChange)
	}

	for _, entry := range update {
		price := entry[0]
		priceFound := false
		// Check if this price exists in our cached backs, if so, replace it.
		for i, cachedEntry := range cache.Runners[selectionID].Mb {
			if cachedEntry[0] == price {
				priceFound = true
				cache.Runners[selectionID].Mb[i] = entry
				break
			}
		}

		if !priceFound {
			cache.Runners[selectionID].Mb = append(cache.Runners[selectionID].Mb, entry)
		}
	}
}

func (cache *OrderBookCache) updateUnmatchedOrders(selectionID int64, orders []*models.Order) {
	cache.Runners[selectionID].Uo = orders
}

func (cache *OrderBookCache) update(data *models.OrderMarketChange, publishTime int64) {

	cache.LastPublishTime = publishTime

	if data.Closed {
		cache.Closed = true
	}

	// Enumerate over Runners
	for _, orderChange := range data.Orc {

		// Check if this is a Full Image (and therefore need to replace our cached entry) or if it hasn't already been cached
		_, found := cache.Runners[orderChange.ID]

		if orderChange.FullImage || !found {
			cache.Runners[orderChange.ID] = orderChange
		} else {
			// Each entry in runner.Mb/runner.Ml is effectively a tuple of (price, size) e.g. [1.51, 2.0]
			cache.updateMatchedBacks(orderChange.ID, orderChange.Mb)
			cache.updateMatchedLays(orderChange.ID, orderChange.Ml)
			cache.updateUnmatchedOrders(orderChange.ID, orderChange.Uo)
		}
	}
}

func (cache *OrderBookCache) Snap() *OrderBookCache {
	return cache
}

// CachedOrders maps MarketID's to Orders on the Exchange
type CachedOrders map[string]*OrderBookCache
