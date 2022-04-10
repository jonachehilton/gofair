package streaming

type MarketDataFilterFlag string

// MarketDataFilterEnum describes the various field flags that can be passed in to the "fields" array of a MarketSubscription.
// see https://docs.developer.betfair.com/display/1smk3cen4v3lu3yomq5qye0ni/Exchange+Stream+API#ExchangeStreamAPI-Marketdatafieldfiltering/MarketDataFilter
var MarketDataFilterEnum = struct {
	ExBestOffersDisp,
	ExBestOffers,
	ExAllOffers,
	ExTraded,
	ExTradedVol,
	ExLTP,
	ExMarketDef,
	SPTraded,
	SPProjected MarketDataFilterFlag
}{
	ExBestOffersDisp: "EX_BEST_OFFERS_DISP",
	ExBestOffers:     "EX_BEST_OFFERS",
	ExAllOffers:      "EX_ALL_OFFERS",
	ExTraded:         "EX_TRADED",
	ExTradedVol:      "EX_TRADED_VOL",
	ExLTP:            "EX_LTP",
	ExMarketDef:      "EX_MARKET_DEF",
	SPTraded:         "SP_TRADED",
	SPProjected:      "SP_PROJECTED",
}
