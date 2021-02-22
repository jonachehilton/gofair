package streaming

import (
	"github.com/go-openapi/strfmt"
)

type MarketBook struct {
	PublishTime           int64
	MarketID              string
	Status                string
	BetDelay              int32
	BspReconciled         bool
	Complete              bool
	InPlay                bool
	NumberOfWinners       int32
	NumberOfRunners       int
	NumberOfActiveRunners int32
	TotalMatched          float64
	CrossMatching         bool
	RunnersVoidable       bool
	Version               int64
	Runners               []Runner
}

type Runner struct {
	SelectionID      int64
	Handicap         float64
	Status           string
	AdjustmentFactor float64
	LastPriceTraded  float64
	TotalMatched     float64
	RemovalDate      strfmt.DateTime
	EX               ExchangePrices
}

type ExchangePrices struct {
	BestAvailableToBack PositionPriceSize
	BestAvailableToLay  PositionPriceSize
	AvailableToBack     PriceSize
	AvailableToLay      PriceSize
	TradedVolume        PriceSize
}
