package streaming

// MarketID: Order
type OrderCache map[string]MarketOrderCache

type MarketOrderCache struct {
	MarketID     string
	RunnerOrders map[int]RunnerOrderCache
}

type RunnerOrderCache struct {
	SelectionID int
	BackMatched float64
	LayMatched  float64
}
