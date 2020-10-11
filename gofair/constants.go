package gofair

type orderProjection string

var OrderProjection = struct {
	All, Executable, ExecutionComplete orderProjection
}{"ALL", "EXECUTABLE", "EXECUTION_COMPLETE"}

type priceData string

var PriceData = struct {
	SPAvailable, SPTraded, ExBestOffers, ExAllOffers, ExTraded priceData
}{"SP_AVAILABLE", "SP_TRADED", "EX_BEST_OFFERS", "EX_ALL_OFFERS", "EX_TRADED"}

type matchProjection string

var MatchProjection = struct {
	NoRollup, RolledUpByPrice, RolledUpByAvgPrice matchProjection
}{"NO_ROLLUP", "ROLLED_UP_BY_PRICE", "ROLLED_UP_BY_AVG_PRICE"}

type marketStatus string

var MarketStatus = struct {
	Inactive, Open, Suspended, Closed marketStatus
}{"INACTIVE", "OPEN", "SUSPENDED", "CLOSED"}

type persistenceType string

var PersistenceType = struct {
	Lapse, Persist, MarketOnClose persistenceType
}{"LAPSE", "PERSIST", "MARKET_ON_CLOSE"}

type orderType string

var OrderType = struct {
	Limit, LimitOnClose, MarketOnClose orderType
}{"LIMIT", "LIMIT_ON_CLOSE", "MARKET_ON_CLOSE"}

type orderStatus string

var OrderStatus = struct {
	Pending, ExecutionComplete, Executable, Expired orderStatus
}{"PENDING", "EXECUTION_COMPLETE", "EXECUTABLE", "EXPIRED"}

type side string

var Side = struct {
	Back, Lay side
}{"BACK", "LAY"}

type WeightConstant int

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
}{200, 1, 1, 0, 0, 0, 1, 0, 2, 3, 7, 5, 17, 17, 20, 32, 4}

type executionReportStatus string

var ExecutionReportStatus = struct {
	Success, Failure, ProcessedWithErrors, Timeout executionReportStatus
}{"SUCCESS", "FAILURE", "PROCESSED_WITH_ERRORS", "TIMEOUT"}

type instructionReportStatus string

var InstructionReportStatus = struct {
	Success, Failure, Timeout instructionReportStatus
}{"SUCCESS", "FAILURE", "TIMEOUT"}

type executionReportErrorCode string

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
}{"ERROR_IN_MATCHER",
	"PROCESSED_WITH_ERRORS",
	"BET_ACTION_ERROR",
	"INVALID_ACCOUNT_STATE",
	"INVALID_WALLET_STATUS",
	"INSUFFICIENT_FUNDS",
	"LOSS_LIMIT_EXCEEDED",
	"MARKET_SUSPENDED",
	"MARKET_NOT_OPEN_FOR_BETTING",
	"DUPLICATE_TRANSACTION",
	"INVALID_ORDER",
	"INVALID_MARKET_ID",
	"PERMISSION_DENIED",
	"DUPLICATE_BETIDS",
	"NO_ACTION_REQUIRED",
	"SERVICE_UNAVAILABLE",
	"REJECTED_BY_REGULATOR",
	"NO_CHASING",
	"REGULATOR_IS_NOT_AVAILABLE",
	"TOO_MANY_INSTRUCTIONS",
	"INVALID_MARKET_VERSION",
	"INVALID_PROFIT_RATIO"}

const (
	MinStakeSizeGBP = 2.00
)
