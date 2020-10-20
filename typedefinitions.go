package gofair

import "time"

type eventType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EventTypeResult struct {
	MarketCount int       `json:"marketCount"`
	EventType   eventType `json:"eventType"`
}

type competition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CompetitionResult struct {
	MarketCount       int         `json:"marketCount"`
	CompetitionRegion string      `json:"competitionRegion"`
	Competition       competition `json:"competition"`
}

type timeRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type TimeRangeResult struct {
	MarketCount int       `json:"marketCount"`
	TimeRange   timeRange `json:"timeRange"`
}

type event struct {
	ID          string `json:"id"`
	OpenDate    string `json:"openDate"`
	TimeZone    string `json:"timezone"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Venue       string `json:"venue"`
}

type EventResult struct {
	MarketCount int   `json:"marketCount"`
	Event       event `json:"event"`
}

type MarketTypeResult struct {
	MarketCount int    `json:"marketCount"`
	MarketType  string `json:"marketType"`
}

type CountryResult struct {
	MarketCount int    `json:"marketCount"`
	CountryCode string `json:"countryCode"`
}

type VenueResult struct {
	MarketCount int    `json:"marketCount"`
	Venue       string `json:"venue"`
}

type PriceSize struct {
	Price float32 `json:"price"`
	Size  float32 `json:"size"`
}

type startingPrices struct {
	NearPrice         float32     `json:"nearPrice"`
	FarPrice          float32     `json:"farPrice"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken"`
	ActualSP          float32     `json:"actualSP"`
}

type exchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack"`
	AvailableToLay  []PriceSize `json:"availableToLay"`
	TradedVolume    []PriceSize `json:"tradedVolume"`
}

type Runner struct {
	SelectionID       int                `json:"selectionId"`
	Handicap          float32            `json:"handicap"`
	Status            string             `json:"status"`
	AdjustmentFactor  float32            `json:"adjustmentFactor"`
	LastPriceTraded   float32            `json:"lastPriceTraded"`
	TotalMatched      float32            `json:"totalMatched"`
	RemovalDate       time.Time          `json:"removalDate"`
	StartingPrices    startingPrices     `json:"sp"`
	ExchangePrices    exchangePrices     `json:"ex"`
	Orders            []Order            `json:"orders"`
	Matches           []Match            `json:"matches"`
	MatchesByStrategy map[string][]Match `json:"matchesByStrategy"`
}

type marketCatalogueDescription struct {
	BettingType        string    `json:"bettingType"`
	BSPMarket          bool      `json:"bspMarket"`
	DiscountAllowed    bool      `json:"discountAllowed"`
	MarketBaseRate     float32   `json:"marketBaseRate"`
	MarketTime         time.Time `json:"marketTime"`
	MarketType         string    `json:"marketType"`
	PersistenceEnabled bool      `json:"persistenceEnabled"`
	Regulator          string    `json:"regulator"`
	Rules              string    `json:"rules"`
	RulesHasDate       bool      `json:"rulesHasDate"`
	SuspendDate        time.Time `json:"suspendTime"`
	TurnInPlayEnabled  bool      `json:"turnInPlayEnabled"`
	Wallet             string    `json:"wallet"`
	EachWayDivisor     float32   `json:"eachWayDivisor"`
	Clarifications     string    `json:"clarifications"`
}

type metadata struct {
	RunnerID int `json:"runnerId"`
}

type runnerCatalogue struct {
	SelectionID  int     `json:"selectionId"`
	RunnerName   string  `json:"runnerName"`
	SortPriority int     `json:"sortPriority"`
	Handicap     float32 `json:"handicap"`
}

type ExBestOffersOverrides struct {
	BestPricesDepth          int     `json:"bestPricesDepth,omitempty"`
	RollupModel              string  `json:"rollupModel,omitempty"`
	RollupLimit              int     `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float32 `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int     `json:"rollupLiabilityFactor,omitempty"`
}

type PriceProjection struct {
	PriceData             []priceData           `json:"priceData,omitempty"`
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise"`
	RollOverStakes        bool                  `json:"rolloverStakes"`
}

type LimitOrder struct {
	Size            float32         `json:"size,omitempty"`
	Price           float32         `json:"price,omitempty"`
	PersistenceType persistenceType `json:"persistenceType,omitempty"`
}

type LimitOnCloseOrder struct {
	Liability float32 `json:"liability,omitempty"`
	Price     float32 `json:"price,omitempty"`
}

type MarketOnCloseOrder struct {
	Liability float32 `json:"liability,omitempty"`
}

type CancelInstruction struct {
	BetID         string  `json:"betId"`
	SizeReduction float32 `json:"sizeReduction,omitempty"`
}

type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef"`
	Status             string                    `json:"status"`
	ErrorCode          string                    `json:"errorCode"`
	MarketID           string                    `json:"marketId"`
	InstructionReports []CancelInstructionReport `json:"instructionReports"`
}

type CancelInstructionReport struct {
	Status        string            `json:"status"`
	ErrorCode     string            `json:"errorCode"`
	Instruction   CancelInstruction `json:"instruction"`
	SizeCancelled float32           `json:"sizeCancelled"`
	CancelledDate time.Time         `json:"cancelledDate"`
}

type PlaceInstruction struct {
	OrderType          orderType           `json:"orderType,omitempty"`
	SelectionID        int                 `json:"selectionId,omitempty"`
	Handicap           float32             `json:"handicap"`
	Side               side                `json:"side,omitempty"`
	LimitOrder         LimitOrder          `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string              `json:"customerOrderRef,omitempty"`
}

type PlaceInstructionReport struct {
	Status              instructionReportStatus `json:"status"`
	ErrorCode           string                  `json:"errorCode"`
	OrderStatus         orderStatus             `json:"orderStatus"`
	Instruction         PlaceInstruction        `json:"instruction"`
	BetID               string                  `json:"betId"`
	PlacedDate          time.Time               `json:"placedDate"`
	AveragePriceMatched float32                 `json:"averagePriceMatched"`
	SizeMatched         float32                 `json:"sizeMatched"`
}

type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef"`
	Status             executionReportStatus    `json:"status"`
	ErrorCode          executionReportErrorCode `json:"errorCode"`
	MarketID           string                   `json:"marketId"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports"`
}

type Order struct {
	BetID               string    `json:"betId"`
	OrderType           string    `json:"orderType"`
	Status              string    `json:"status"`
	PersistenceType     string    `json:"persistenceType"`
	Side                string    `json:"side"`
	Price               float32   `json:"price"`
	Size                float32   `json:"size"`
	BSPLiability        float32   `json:"bspLiability"`
	PlacedDate          time.Time `json:"placedDate"`
	AvgPriceMatched     float32   `json:"avgPriceMatched"`
	SizeMatched         float32   `json:"sizeMatched"`
	SizeRemaining       float32   `json:"sizeRemaining"`
	SizeLapsed          float32   `json:"sizeLapsed"`
	SizeCancelled       float32   `json:"sizeCancelled"`
	SizeVoided          float32   `json:"sizeVoided"`
	CustomerOrderRef    string    `json:"customerOrderRef"`
	CustomerStrategyRef string    `json:"customerStrategyRef"`
}

type KeyLineSelection struct {
	SelectionID int     `json:"selectionId"`
	Handicap    float32 `json:"handicap"`
}

type KeyLineDescription struct {
	KeyLine []KeyLineSelection `json:"keyLine"`
}

type Match struct {
	BetID     string    `json:"betId"`
	MatchID   string    `json:"matchId"`
	Side      side      `json:"side"`
	Price     float32   `json:"price"`
	Size      float32   `json:"size"`
	MatchDate time.Time `json:"matchDate"`
}

// MarketCatalogue holds information about a market.
type MarketCatalogue struct {
	MarketID                   string                     `json:"marketId"`
	MarketName                 string                     `json:"marketName"`
	TotalMatched               float32                    `json:"totalMatched"`
	MarketStartTime            time.Time                  `json:"marketStartTime"`
	Competition                competition                `json:"competition"`
	Event                      event                      `json:"event"`
	EventType                  eventType                  `json:"eventType"`
	MarketCatalogueDescription marketCatalogueDescription `json:"description"`
	Runners                    []runnerCatalogue          `json:"runners"`
}

// MarketBook holds the dynamic data in a market.
type MarketBook struct {
	MarketID              string             `json:"marketId"`
	IsMarketDataDelayed   bool               `json:"isMarketDataDelayed"`
	Status                string             `json:"status"`
	BetDelay              int                `json:"betDelay"`
	BSPReconciled         bool               `json:"bspReconciled"`
	Complete              bool               `json:"complete"`
	InPlay                bool               `json:"inplay"`
	NumberOfWinners       int                `json:"numberOfWinners"`
	NumberOfRunners       int                `json:"numberOfRunners"`
	NumberOfActiveRunners int                `json:"numberOfActiveRunners"`
	LastMatchTime         time.Time          `json:"lastMatchTime"`
	TotalMatched          float32            `json:"totalMatched"`
	TotalAvailable        float32            `json:"totalAvailable"`
	CrossMatching         bool               `json:"crossMatching"`
	RunnersVoidable       bool               `json:"runnersVoidable"`
	Version               int64              `json:"version"`
	Runners               []Runner           `json:"runners"`
	KeyLineDescription    KeyLineDescription `json:"keyLineDescription"`
}

type RunnerProfitAndLoss struct {
	SelectionID int     `json:"selectionId"`
	IfWin       float32 `json:"ifWin"`
	IfLose      float32 `json:"ifLose"`
	IfPlace     float32 `json:"ifPlace"`
}

type MarketProfitAndLoss struct {
	MarketID          string                `json:"marketId"`
	CommissionApplied float32               `json:"commissionApplied"`
	ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses"`
}
