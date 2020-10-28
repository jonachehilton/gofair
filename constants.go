package gofair

type orderProjection string
type priceData string
type matchProjection string
type marketStatus string
type persistenceType string
type orderType string
type orderStatus string
type executionReportStatus string
type instructionReportStatus string
type executionReportErrorCode string

// BackOrLay is the type associated with strings contained in the Side struct.
type BackOrLay string

// WeightConstant should be the type specified for const values relating to request limit calculations
// https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Market+Data+Request+Limits
type WeightConstant int

// OrderProjection describes the orders you want to receive in the response.
var OrderProjection = struct {
	All, Executable, ExecutionComplete orderProjection
}{
	All:               "ALL",
	Executable:        "EXECUTABLE",
	ExecutionComplete: "EXECUTION_COMPLETE",
}

// PriceData describes the basic price data you want to receive in the response.
var PriceData = struct {
	SPAvailable, SPTraded, ExBestOffers, ExAllOffers, ExTraded priceData
}{
	SPAvailable:  "SP_AVAILABLE",
	SPTraded:     "SP_TRADED",
	ExBestOffers: "EX_BEST_OFFERS",
	ExAllOffers:  "EX_ALL_OFFERS",
	ExTraded:     "EX_TRADED",
}

// MatchProjection describes the matches you want to receive in the response.
var MatchProjection = struct {
	NoRollup, RolledUpByPrice, RolledUpByAvgPrice matchProjection
}{
	NoRollup:           "NO_ROLLUP",
	RolledUpByPrice:    "ROLLED_UP_BY_PRICE",
	RolledUpByAvgPrice: "ROLLED_UP_BY_AVG_PRICE",
}

// MarketStatus describes the status of the market, for example OPEN, SUSPENDED, CLOSED (settled), etc.
var MarketStatus = struct {
	Inactive, Open, Suspended, Closed marketStatus
}{
	Inactive:  "INACTIVE",
	Open:      "OPEN",
	Suspended: "SUSPENDED",
	Closed:    "CLOSED",
}

// PersistenceType describes what to do with the order at turn-in-play.
var PersistenceType = struct {
	Lapse, Persist, MarketOnClose persistenceType
}{
	Lapse:         "LAPSE",
	Persist:       "PERSIST",
	MarketOnClose: "MARKET_ON_CLOSE",
}

// OrderType describes the BSP Order type.
var OrderType = struct {
	Limit, LimitOnClose, MarketOnClose orderType
}{
	Limit:         "LIMIT",
	LimitOnClose:  "LIMIT_ON_CLOSE",
	MarketOnClose: "MARKET_ON_CLOSE",
}

// OrderStatus should generally be either EXECUTABLE (an unmatched amount remains) or EXECUTION_COMPLETE (no unmatched amount remains).
var OrderStatus = struct {
	Pending, ExecutionComplete, Executable, Expired orderStatus
}{
	Pending:           "PENDING",
	ExecutionComplete: "EXECUTION_COMPLETE",
	Executable:        "EXECUTABLE",
	Expired:           "EXPIRED",
}

// Side indicates if the bet is a Back or a Lay.
var Side = struct {
	Back, Lay BackOrLay
}{
	Back: "BACK",
	Lay:  "LAY",
}

// Weight is a measure used by the Betfair Exchange API to describe the relative amount of data a particular type of request is expected to return.
var Weight = struct {
	MaxWeight,
	MarketDescription,
	RunnerDescription,
	Event,
	EventType,
	Competition,
	RunnerMetadata,
	MarketStartTime,
	NotSet,
	SPAvailable,
	SPTraded,
	ExBestOffers,
	ExAllOffers,
	ExTraded,
	ExBestOffersAndExTraded,
	ExAllOffersAndExTraded,
	NotApplicable WeightConstant
}{
	MaxWeight:               200,
	MarketDescription:       1,
	RunnerDescription:       1,
	Event:                   0,
	EventType:               0,
	Competition:             0,
	RunnerMetadata:          1,
	MarketStartTime:         0,
	NotSet:                  2,
	SPAvailable:             3,
	SPTraded:                7,
	ExBestOffers:            5,
	ExAllOffers:             17,
	ExTraded:                17,
	ExBestOffersAndExTraded: 20,
	ExAllOffersAndExTraded:  32,
	NotApplicable:           4,
}

// ExecutionReportStatus describes the outcome of placing an order.
var ExecutionReportStatus = struct {
	Success, Failure, ProcessedWithErrors, Timeout executionReportStatus
}{
	Success:             "SUCCESS",
	Failure:             "FAILURE",
	ProcessedWithErrors: "PROCESSED_WITH_ERRORS",
	Timeout:             "TIMEOUT",
}

// InstructionReportStatus describes the outcome of a particular instruction being submitted.
var InstructionReportStatus = struct {
	Success, Failure, Timeout instructionReportStatus
}{
	Success: "SUCCESS",
	Failure: "FAILURE",
	Timeout: "TIMEOUT",
}

// ExecutionReportErrorCode describes the potential errors contained within a PlaceExecutionReport.
var ExecutionReportErrorCode = struct {
	ErrorInMatcher,
	ProcessedWithErrors,
	BetActionError,
	InvalidAccountState,
	InvalidWalletStatus,
	InsufficientFunds,
	LossLimitExceeded,
	MarketSuspended,
	MarketNotOpenForBetting,
	DuplicateTransaction,
	InvalidOrder,
	InvalidMarketID,
	PermissionDenied,
	DuplicateBetIDs,
	NoActionRequired,
	ServiceUnavailable,
	RejectedByRegulator,
	NoChasing,
	RegulatorIsNotAvailable,
	TooManyInstructions,
	InvalidMarketVersion,
	InvalidProfitRatio executionReportErrorCode
}{
	ErrorInMatcher:          "ERROR_IN_MATCHER",
	ProcessedWithErrors:     "PROCESSED_WITH_ERRORS",
	BetActionError:          "BET_ACTION_ERROR",
	InvalidAccountState:     "INVALID_ACCOUNT_STATE",
	InvalidWalletStatus:     "INVALID_WALLET_STATUS",
	InsufficientFunds:       "INSUFFICIENT_FUNDS",
	LossLimitExceeded:       "LOSS_LIMIT_EXCEEDED",
	MarketSuspended:         "MARKET_SUSPENDED",
	MarketNotOpenForBetting: "MARKET_NOT_OPEN_FOR_BETTING",
	DuplicateTransaction:    "DUPLICATE_TRANSACTION",
	InvalidOrder:            "INVALID_ORDER",
	InvalidMarketID:         "INVALID_MARKET_ID",
	PermissionDenied:        "PERMISSION_DENIED",
	DuplicateBetIDs:         "DUPLICATE_BETIDS",
	NoActionRequired:        "NO_ACTION_REQUIRED",
	ServiceUnavailable:      "SERVICE_UNAVAILABLE",
	RejectedByRegulator:     "REJECTED_BY_REGULATOR",
	NoChasing:               "NO_CHASING",
	RegulatorIsNotAvailable: "REGULATOR_IS_NOT_AVAILABLE",
	TooManyInstructions:     "TOO_MANY_INSTRUCTIONS",
	InvalidMarketVersion:    "INVALID_MARKET_VERSION",
	InvalidProfitRatio:      "INVALID_PROFIT_RATIO",
}

const (
	MinimumStakeSizeGBP = 2.00
)

// Endpoints contains all the Betfair Exchange API endpoints.
var Endpoints = struct {
	Login,
	Identity,
	Betting,
	Account,
	Navigation string
}{
	Login:      "https://identitysso-api.betfair.com/api/",
	Identity:   "https://identitysso.betfair.com/api/",
	Betting:    "https://api.betfair.com/exchange/betting/rest/v1.0/",
	Account:    "https://api.betfair.com/exchange/account/rest/v1.0/",
	Navigation: "https://api.betfair.com/exchange/betting/rest/v1/en/navigation/menu.json",
}
