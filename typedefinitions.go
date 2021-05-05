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

type Competition struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CompetitionResult struct {
	MarketCount       int         `json:"marketCount"`
	CompetitionRegion string      `json:"competitionRegion"`
	Competition       Competition `json:"competition"`
}

type timeRange struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

type TimeRangeResult struct {
	MarketCount int       `json:"marketCount"`
	TimeRange   timeRange `json:"timeRange"`
}

type Event struct {
	ID          string `json:"id"`
	OpenDate    string `json:"openDate"`
	TimeZone    string `json:"timezone"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Venue       string `json:"venue"`
}

type EventResult struct {
	MarketCount int   `json:"marketCount"`
	Event       Event `json:"event"`
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
	PriceData             []PriceData           `json:"priceData,omitempty"`
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise"`
	RollOverStakes        bool                  `json:"rolloverStakes"`
}

type LimitOrder struct {
	Size            float32         `json:"size,omitempty"`
	Price           float32         `json:"price,omitempty"`
	PersistenceType PersistenceType `json:"persistenceType,omitempty"`
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
	OrderType          OrderType           `json:"orderType,omitempty"`
	SelectionID        int                 `json:"selectionId,omitempty"`
	Handicap           float32             `json:"handicap"`
	Side               Side                `json:"side,omitempty"`
	LimitOrder         LimitOrder          `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string              `json:"customerOrderRef,omitempty"`
}

type PlaceInstructionReport struct {
	Status              InstructionReportStatus `json:"status"`
	ErrorCode           string                  `json:"errorCode"`
	OrderStatus         OrderStatus             `json:"orderStatus"`
	Instruction         PlaceInstruction        `json:"instruction"`
	BetID               string                  `json:"betId"`
	PlacedDate          time.Time               `json:"placedDate"`
	AveragePriceMatched float32                 `json:"averagePriceMatched"`
	SizeMatched         float32                 `json:"sizeMatched"`
}

type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef"`
	Status             ExecutionReportStatus    `json:"status"`
	ErrorCode          ExecutionReportErrorCode `json:"errorCode"`
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

// KeyLineSelection provides a description of a markets key line selection, comprising the selectionId and handicap of the team it is applied to.
type KeyLineSelection struct {
	SelectionID int     `json:"selectionId"`
	Handicap    float32 `json:"handicap"`
}

// KeyLineDescription provides a description of a markets key line for valid market types
type KeyLineDescription struct {
	KeyLine []KeyLineSelection `json:"keyLine"`
}

type Match struct {
	BetID     string    `json:"betId"`
	MatchID   string    `json:"matchId"`
	Side      Side      `json:"side"`
	Price     float32   `json:"price"`
	Size      float32   `json:"size"`
	MatchDate time.Time `json:"matchDate"`
}

// MarketCatalogue holds the static data in a market.
type MarketCatalogue struct {
	MarketID                   string                     `json:"marketId"`
	MarketName                 string                     `json:"marketName"`
	TotalMatched               float32                    `json:"totalMatched"`
	MarketStartTime            time.Time                  `json:"marketStartTime"`
	Competition                Competition                `json:"competition"`
	Event                      Event                      `json:"event"`
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

// RunnerProfitAndLoss contains potential changes in winnings in the event of a particular selection winning, losing or placing.
type RunnerProfitAndLoss struct {
	SelectionID int     `json:"selectionId"`
	IfWin       float32 `json:"ifWin"`
	IfLose      float32 `json:"ifLose"`
	IfPlace     float32 `json:"ifPlace"`
}

// MarketProfitAndLoss contains changes in winnings depending on the performance of selections associated with a given market.
type MarketProfitAndLoss struct {
	MarketID          string                `json:"marketId"`
	CommissionApplied float32               `json:"commissionApplied"`
	ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses"`
}

type TimeRangeFilter struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

// MarketFilter is the filter to select desired markets. All markets that match the criteria in the filter are selected.
type MarketFilter struct {
	TextQuery          string           `json:"textQuery,omitempty"`
	EventTypeIds       []string         `json:"eventTypeIds,omitempty"`
	MarketCountries    []string         `json:"marketCountries,omitempty"`
	MarketIds          []string         `json:"marketIds,omitempty"`
	EventIds           []string         `json:"eventIds,omitempty"`
	CompetitionIds     []string         `json:"competitionIds,omitempty"`
	BSPOnly            bool             `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  bool             `json:"turnInPLayEnabled,omitempty"`
	InPlayOnly         bool             `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []string         `json:"marketBettingTypes,omitempty"`
	MarketTypeCodes    []string         `json:"marketTypeCodes,omitempty"`
	MarketStartTime    *TimeRangeFilter `json:"marketStartTime,omitempty"`
	WithOrders         string           `json:"withOrders,omitempty"`
}

type MarketDataFilter struct {
	Fields       []string `json:"fields"`
	LadderLevels int64    `json:"ladderLevels"`
}

type CurrentOrderSummary struct {
	BetID               string          `json:"betId"`
	MarketID            string          `json:"marketId"`
	SelectionID         string          `json:"selectionId"`
	Handicap            float64         `json:"handicap"`
	PriceSize           PriceSize       `json:"priceSize"`
	BSPLiability        float64         `json:"bspLiability"`
	Side                Side            `json:"side"`
	Status              OrderStatus     `json:"status"`
	PersistenceType     PersistenceType `json:"persistenceType"`
	OrderType           OrderType       `json:"orderType"`
	PlacedDate          time.Time       `json:"placedDate"`
	MatchedDate         time.Time       `json:"matchedDate"`
	AveragePriceMatched float64         `json:"averagePriceMatched,omitempty"`
	SizeMatched         float64         `json:"sizeMatched,omitempty"`
	SizeRemaining       float64         `json:"sizeRemaining,omitempty"`
	SizeLapsed          float64         `json:"sizeLapsed,omitempty"`
	SizeCancelled       float64         `json:"sizeCancelled,omitempty"`
	SizeVoided          float64         `json:"sizeVoided,omitempty"`
	RegulatorAuthCode   string          `json:"regulatorAuthCode,omitempty"`
	RegulatorCode       string          `json:"regulatorCode,omitempty"`
	CustomerOrderRef    string          `json:"customerOrderRef,omitempty"`
	CustomerStrategyRef string          `json:"customerOrderStrategy,omitempty"`
}

type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}
