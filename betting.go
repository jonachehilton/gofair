package gofair

// ListEventTypes returns a list of Event Types (i.e. Sports) associated with the markets selected by the MarketFilter.
func (b *Betting) ListEventTypes(filter MarketFilter) ([]EventTypeResult, error) {
	// create url
	url := createURL(apiBettingURL, "listEventTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []EventTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListCompetitions returns a list of Competitions (i.e., World Cup 2013) associated with the markets selected by the MarketFilter. Currently only Football markets have an associated competition.
func (b *Betting) ListCompetitions(filter MarketFilter) ([]CompetitionResult, error) {
	// create url
	url := createURL(apiBettingURL, "listCompetitions/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []CompetitionResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListTimeRanges returns a list of time ranges in the granularity specified in the request (i.e. 3PM to 4PM, Aug 14th to Aug 15th) associated with the markets selected by the MarketFilter.
func (b *Betting) ListTimeRanges(filter MarketFilter, granularity string) ([]TimeRangeResult, error) {
	// create url
	url := createURL(apiBettingURL, "listTimeRanges/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter
	params.Granularity = granularity

	var response []TimeRangeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListEvents returns a list of Events (i.e, Reading vs. Man United) associated with the markets selected by the MarketFilter.
func (b *Betting) ListEvents(filter MarketFilter) ([]EventResult, error) {
	// create url
	url := createURL(apiBettingURL, "listEvents/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []EventResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListMarketTypes returns a list of market types (i.e. MATCH_ODDS, NEXT_GOAL) associated with the markets selected by the MarketFilter. The market types are always the same, regardless of locale.
func (b *Betting) ListMarketTypes(filter MarketFilter) ([]MarketTypeResult, error) {
	// create url
	url := createURL(apiBettingURL, "listMarketTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []MarketTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListCountries returns a list of Countries associated with the markets selected by the MarketFilter.
func (b *Betting) ListCountries(filter MarketFilter) ([]CountryResult, error) {
	// create url
	url := createURL(apiBettingURL, "listCountries/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []CountryResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListVenues returns a list of Venues (i.e. Cheltenham, Ascot) associated with the markets selected by the MarketFilter. Currently, only Horse Racing markets are associated with a Venue.
func (b *Betting) ListVenues(filter MarketFilter) ([]VenueResult, error) {
	// create url
	url := createURL(apiBettingURL, "listVenues/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []VenueResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListMarketCatalogue returns a list of information about published (ACTIVE/SUSPENDED) markets that does not change (or changes very rarely). You use listMarketCatalogue to retrieve the name of the market, the names of selections and other information about markets.  Market Data Request Limits apply to requests made to listMarketCatalogue.
func (b *Betting) ListMarketCatalogue(filter MarketFilter, marketProjection []string, sort string, maxResults int) (
	[]MarketCatalogue, error) {
	// create url
	url := createURL(apiBettingURL, "listMarketCatalogue/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter
	params.MarketProjection = marketProjection
	params.Sort = sort
	params.MaxResults = maxResults

	var response []MarketCatalogue

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListMarketBook returns a list of dynamic data about markets. Dynamic data includes prices, the status of the market, the status of selections, the traded volume, and the status of any orders you have placed in the market.
func (b *Betting) ListMarketBook(marketIDs []string, displayOrders bool) ([]MarketBook, error) {
	// TODO: At some point need to expand the number of parameters this can take in order to provide more options to user.
	// create url
	url := createURL(apiBettingURL, "listMarketBook/")

	// build request
	params := new(Params)
	params.MarketIDs = marketIDs
	params.IsMarketDataDelayed = false

	if displayOrders == true {
		params.OrderProjection = OrderProjection.Executable
		params.MatchProjection = MatchProjection.RolledUpByAvgPrice
	} else {
		params.OrderProjection = OrderProjection.All
		priceProjection := new(PriceProjection)
		priceProjection.PriceData = append(priceProjection.PriceData, PriceData.ExBestOffers)
		priceProjection.ExBestOffersOverrides.BestPricesDepth = 3
		params.PriceProjection = priceProjection
	}

	var response []MarketBook

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// ListMarketProfitAndLoss retrieves profit and loss for a given list of OPEN markets. The values are calculated using matched bets and optionally settled bets. Only odds (MarketBettingType = ODDS) markets  are implemented, markets of other types are silently ignored.
func (b *Betting) ListMarketProfitAndLoss(marketIDs []string) ([]MarketProfitAndLoss, error) {
	// create url
	url := createURL(apiBettingURL, "listMarketProfitAndLoss/")

	// build request
	params := new(Params)
	params.MarketIDs = marketIDs

	var response []MarketProfitAndLoss

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

// PlaceOrders allows new orders to be submitted into a market. Please note that additional bet sizing rules apply to bets placed into the Italian Exchange.
func (b *Betting) PlaceOrders(marketID string, placeInstructions []PlaceInstruction) (PlaceExecutionReport, error) {
	// create url
	url := createURL(apiBettingURL, "placeOrders/")
	// build request

	params := new(Params)
	params.MarketID = marketID
	params.Instructions = placeInstructions

	var response PlaceExecutionReport

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return response, err
	}
	return response, err
}

// CancelOrders allows the user to cancel all bets OR cancel all bets on a market OR fully or partially cancel particular orders on a market. Only LIMIT orders can be cancelled or partially cancelled once placed.
func (b *Betting) CancelOrders(marketID string, cancelInstructions []CancelInstruction) (CancelExecutionReport, error) {
	// create url
	url := createURL(apiBettingURL, "cancelOrders/")
	// build request

	params := new(Params)
	params.MarketID = marketID
	params.CancelInstructions = cancelInstructions

	var response CancelExecutionReport

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return response, err
	}
	return response, err
}
