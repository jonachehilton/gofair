package gofair

// Betting API Operations
const (
	listEventTypes = "listEventTypes/"
	listCompetitions = "listCompetitions/"
	listTimeRanges = "listTimeRanges/"
	listEvents = "listEvents/"
	listMarketTypes = "listMarketTypes/"
	listCountries = "listCountries/"
	listVenues = "listVenues/"
	listMarketCatalogue = "listMarketCatalogue/"
	listMarketBook = "listMarketBook/"
	listMarketProfitAndLoss = "listMarketProfitAndLoss/"
	placeOrders = "placeOrders/"
	cancelOrders = "cancelOrders/"
	listCurrentOrders = "listCurrentOrders/"
)

// Betting object
type Betting struct {
	Client *Client
}

func (b *Betting) bettingRequest(endpoint string, params interface{}, response interface{}) error{
	
	url := createURL(Endpoints.Betting, endpoint)

	// make request
	err := b.Client.request(url, params, &response)
	if err != nil {
		return err
	}

	return nil
}

// ListEventTypes returns a list of Event Types (i.e. Sports) associated with the markets selected by the MarketFilter.
func (b *Betting) ListEventTypes(filter MarketFilter) ([]EventTypeResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []EventTypeResult

	err := b.bettingRequest(listEventTypes, params, response)

	return response, err
}

// ListCompetitions returns a list of Competitions (i.e., World Cup 2013) associated with the markets selected by the MarketFilter. Currently only Football markets have an associated competition.
func (b *Betting) ListCompetitions(filter MarketFilter) ([]CompetitionResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []CompetitionResult

	err := b.bettingRequest(listCompetitions, params, response)

	return response, err
}

// ListTimeRanges returns a list of time ranges in the granularity specified in the request (i.e. 3PM to 4PM, Aug 14th to Aug 15th) associated with the markets selected by the MarketFilter.
func (b *Betting) ListTimeRanges(filter MarketFilter, granularity string) ([]TimeRangeResult, error) {
	// build request
	params := struct {
		Filter      MarketFilter `json:"filter,omitempty"`
		Granularity string       `json:"granularity,omitempty"`
	}{
		Filter:      filter,
		Granularity: granularity,
	}

	var response []TimeRangeResult

	err := b.bettingRequest(listTimeRanges, params, response)

	return response, err
}

// ListEvents returns a list of Events (i.e, Reading vs. Man United) associated with the markets selected by the MarketFilter.
func (b *Betting) ListEvents(filter MarketFilter) ([]EventResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []EventResult

	err := b.bettingRequest(listEvents, params, response)

	return response, err
}

// ListMarketTypes returns a list of market types (i.e. MATCH_ODDS, NEXT_GOAL) associated with the markets selected by the MarketFilter. The market types are always the same, regardless of locale.
func (b *Betting) ListMarketTypes(filter MarketFilter) ([]MarketTypeResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []MarketTypeResult

	err := b.bettingRequest(listMarketTypes, params, response)

	return response, err
}

// ListCountries returns a list of Countries associated with the markets selected by the MarketFilter.
func (b *Betting) ListCountries(filter MarketFilter) ([]CountryResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []CountryResult

	err := b.bettingRequest(listCountries, params, response)

	return response, err
}

// ListVenues returns a list of Venues (i.e. Cheltenham, Ascot) associated with the markets selected by the MarketFilter. Currently, only Horse Racing markets are associated with a Venue.
func (b *Betting) ListVenues(filter MarketFilter) ([]VenueResult, error) {
	// build request
	params := struct {
		Filter MarketFilter `json:"filter,omitempty"`
	}{
		Filter: filter,
	}

	var response []VenueResult

	err := b.bettingRequest(listVenues, params, response)

	return response, err
}

// ListMarketCatalogue returns a list of information about published (ACTIVE/SUSPENDED) markets that does not change (or changes very rarely). You use listMarketCatalogue to retrieve the name of the market, the names of selections and other information about markets.  Market Data request Limits apply to requests made to listMarketCatalogue.
func (b *Betting) ListMarketCatalogue(filter MarketFilter, marketProjection []string, sort string, maxResults int) ([]MarketCatalogue, error) {
	// build request
	params := struct {
		Filter           MarketFilter `json:"filter,omitempty"`
		MarketProjection []string     `json:"marketProjection,omitempty"`
		Sort             string       `json:"sort,omitempty"`
		MaxResults       int          `json:"maxResults,omitempty"`
	}{
		Filter:           filter,
		MarketProjection: marketProjection,
		Sort:             sort,
		MaxResults:       maxResults,
	}

	var response []MarketCatalogue

	err := b.bettingRequest(listMarketCatalogue, params, response)

	return response, err
}

// ListMarketBook returns a list of dynamic data about markets. Dynamic data includes prices, the status of the market, the status of selections, the traded volume, and the status of any orders you have placed in the market.
func (b *Betting) ListMarketBook(marketIDs []string, displayOrders bool) ([]MarketBook, error) {
	// build request
	priceProjection := new(PriceProjection)

	params := struct {
		MarketIDs           []string         `json:"marketIds,omitempty"`
		IsMarketDataDelayed bool             `json:"isMarketDataDelayed,omitempty"`
		OrderProjection     OrderProjection  `json:"orderProjection,omitempty"`
		MatchProjection     MatchProjection  `json:"matchProjection,omitempty"`
		PriceProjection     *PriceProjection `json:"priceProjection,omitempty"`
	}{
		MarketIDs:           marketIDs,
		IsMarketDataDelayed: false,
		OrderProjection:     OrderProjectionEnum.Executable,
		MatchProjection:     MatchProjectionEnum.RolledUpByAvgPrice,
		PriceProjection:     priceProjection,
	}

	if !displayOrders {
		params.OrderProjection = OrderProjectionEnum.All
		params.MatchProjection = ""
		priceProjection.PriceData = append(priceProjection.PriceData, PriceDataEnum.ExBestOffers)
		priceProjection.ExBestOffersOverrides.BestPricesDepth = 3
	}

	var response []MarketBook

	err := b.bettingRequest(listMarketBook, params, response)

	return response, err
}

// ListMarketProfitAndLoss retrieves profit and loss for a given list of OPEN markets. The values are calculated using matched bets and optionally settled bets. Only odds (MarketBettingType = ODDS) markets  are implemented, markets of other types are silently ignored.
func (b *Betting) ListMarketProfitAndLoss(marketIDs []string) ([]MarketProfitAndLoss, error) {
	// build request
	params := struct {
		MarketIDs []string `json:"marketIds,omitempty"`
	}{
		MarketIDs: marketIDs,
	}

	var response []MarketProfitAndLoss

	err := b.bettingRequest(listMarketProfitAndLoss, params, response)

	return response, err
}

// PlaceOrders allows new orders to be submitted into a market. Please note that additional bet sizing rules apply to bets placed into the Italian Exchange.
func (b *Betting) PlaceOrders(marketID string, placeInstructions []PlaceInstruction) (PlaceExecutionReport, error) {
	// build request
	params := struct {
		MarketID     string             `json:"marketId,omitempty"`
		Instructions []PlaceInstruction `json:"instructions,omitempty"`
	}{
		MarketID:     marketID,
		Instructions: placeInstructions,
	}

	var response PlaceExecutionReport

	err := b.bettingRequest(placeOrders, params, response)

	return response, err
}

// CancelOrders allows the user to cancel all bets OR cancel all bets on a market OR fully or partially cancel particular orders on a market. Only LIMIT orders can be cancelled or partially cancelled once placed.
func (b *Betting) CancelOrders(marketID string, cancelInstructions []CancelInstruction) (CancelExecutionReport, error) {
	// build request
	params := struct {
		MarketID     string              `json:"marketId,omitempty"`
		Instructions []CancelInstruction `json:"instructions,omitempty"`
	}{
		MarketID:     marketID,
		Instructions: cancelInstructions,
	}

	var response CancelExecutionReport

	err := b.bettingRequest(cancelOrders, params, response)

	return response, err
}

func (b *Betting) ListCurrentOrders(betIDs []string, marketIDs []string, orderProjection OrderProjection) (CurrentOrderSummaryReport, error) {
	// build request
	params := struct {
		BetIDs          []string        `json:"betIds,omitempty"`
		MarketIDs       []string        `json:"marketIds,omitempty"`
		OrderProjection OrderProjection `json:"orderProjection,omitempty"`
	}{
		BetIDs:          betIDs,
		MarketIDs:       marketIDs,
		OrderProjection: orderProjection,
	}

	var response CurrentOrderSummaryReport

	err := b.bettingRequest(listCurrentOrders, params, response)

	return response, err
}
