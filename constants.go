package gofair

type OrderProjection string

// https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Betting+Enums

// OrderProjectionEnum describes the orders you want to receive in the response.
var OrderProjectionEnum = struct {
	All, Executable, ExecutionComplete OrderProjection
}{
	All:               "ALL",
	Executable:        "EXECUTABLE",
	ExecutionComplete: "EXECUTION_COMPLETE",
}

type PriceData string

// PriceDataEnum describes the basic price data you want to receive in the response.
var PriceDataEnum = struct {
	SPAvailable, SPTraded, ExBestOffers, ExAllOffers, ExTraded PriceData
}{
	SPAvailable:  "SP_AVAILABLE",
	SPTraded:     "SP_TRADED",
	ExBestOffers: "EX_BEST_OFFERS",
	ExAllOffers:  "EX_ALL_OFFERS",
	ExTraded:     "EX_TRADED",
}

type MatchProjection string

// MatchProjectionEnum describes the matches you want to receive in the response.
var MatchProjectionEnum = struct {
	NoRollup, RolledUpByPrice, RolledUpByAvgPrice MatchProjection
}{
	NoRollup:           "NO_ROLLUP",
	RolledUpByPrice:    "ROLLED_UP_BY_PRICE",
	RolledUpByAvgPrice: "ROLLED_UP_BY_AVG_PRICE",
}

type MarketStatus string

// MarketStatusEnum describes the status of the market, for example OPEN, SUSPENDED, CLOSED (settled), etc.
var MarketStatusEnum = struct {
	Inactive, Open, Suspended, Closed MarketStatus
}{
	Inactive:  "INACTIVE",
	Open:      "OPEN",
	Suspended: "SUSPENDED",
	Closed:    "CLOSED",
}

type PersistenceType string

// PersistenceTypeEnum describes what to do with the order at turn-in-play.
var PersistenceTypeEnum = struct {
	Lapse, Persist, MarketOnClose PersistenceType
}{
	Lapse:         "LAPSE",
	Persist:       "PERSIST",
	MarketOnClose: "MARKET_ON_CLOSE",
}

type OrderType string

// OrderTypeEnum describes the BSP Order type.
var OrderTypeEnum = struct {
	Limit, LimitOnClose, MarketOnClose OrderType
}{
	Limit:         "LIMIT",
	LimitOnClose:  "LIMIT_ON_CLOSE",
	MarketOnClose: "MARKET_ON_CLOSE",
}

type OrderStatus string

// OrderStatusEnum should generally be either EXECUTABLE (an unmatched amount remains) or EXECUTION_COMPLETE (no unmatched amount remains).
var OrderStatusEnum = struct {
	Pending, ExecutionComplete, Executable, Expired OrderStatus
}{
	Pending:           "PENDING",
	ExecutionComplete: "EXECUTION_COMPLETE",
	Executable:        "EXECUTABLE",
	Expired:           "EXPIRED",
}

// Side is the type associated with strings contained in the Side struct.
type Side string

// SideEnum indicates if the bet is a Back or a Lay.
var SideEnum = struct {
	Back, Lay Side
}{
	Back: "BACK",
	Lay:  "LAY",
}

// WeightConstant should be the type specified for const values relating to request limit calculations
// https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Market+Data+Request+Limits
type WeightConstant int

// WeightEnum is a measure used by the Betfair Exchange API to describe the relative amount of data a particular type of request is expected to return.
var WeightEnum = struct {
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

type ExecutionReportStatus string

// ExecutionReportStatusEnum describes the outcome of placing an order.
var ExecutionReportStatusEnum = struct {
	Success, Failure, ProcessedWithErrors, Timeout ExecutionReportStatus
}{
	Success:             "SUCCESS",
	Failure:             "FAILURE",
	ProcessedWithErrors: "PROCESSED_WITH_ERRORS",
	Timeout:             "TIMEOUT",
}

type InstructionReportStatus string

// InstructionReportStatusEnum describes the outcome of a particular instruction being submitted.
var InstructionReportStatusEnum = struct {
	Success, Failure, Timeout InstructionReportStatus
}{
	Success: "SUCCESS",
	Failure: "FAILURE",
	Timeout: "TIMEOUT",
}

type ExecutionReportErrorCode string

// ExecutionReportErrorCodeEnum describes the potential errors contained within a PlaceExecutionReport.
var ExecutionReportErrorCodeEnum = struct {
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
	InvalidProfitRatio ExecutionReportErrorCode
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

type TimeInForce string

var TimeInForceEnum = struct {
	FillOrKill TimeInForce
}{
	FillOrKill:  "FILL_OR_KILL",
}

type BetTargetType string

var BetTargetTypeEnum = struct {
	BackersProfit,
	Payout BetTargetType
}{
	BackersProfit: "BACKERS_PROFIT",
	Payout: "PAYOUT",
}