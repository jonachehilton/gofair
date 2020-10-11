package gofair

func (b *Betting) ListEventTypes(filter MarketFilter) ([]eventTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEventTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []eventTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCompetitions(filter MarketFilter) ([]competitionResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCompetitions/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []competitionResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListTimeRanges(filter MarketFilter, granularity string) ([]timeRangeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listTimeRanges/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter
	params.Granularity = granularity

	var response []timeRangeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListEvents(filter MarketFilter) ([]eventResult, error) {
	// create url
	url := createUrl(api_betting_url, "listEvents/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []eventResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketTypes(filter MarketFilter) ([]marketTypeResult, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketTypes/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []marketTypeResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListCountries(filter MarketFilter) ([]countryResult, error) {
	// create url
	url := createUrl(api_betting_url, "listCountries/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []countryResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListVenues(filter MarketFilter) ([]venueResult, error) {
	// create url
	url := createUrl(api_betting_url, "listVenues/")

	// build request
	params := new(Params)
	params.MarketFilter = &filter

	var response []venueResult

	// make request
	err := b.Request(url, params, &response)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (b *Betting) ListMarketCatalogue(filter MarketFilter, marketProjection []string, sort string, maxResults int) (
	[]MarketCatalogue, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketCatalogue/")

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

// TODO: At some point need to expand the number of parameters this can take in order to provide more options to user.
func (b *Betting) ListMarketBook(marketIDs []string, displayOrders bool) ([]MarketBook, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketBook/")

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

func (b *Betting) ListMarketProfitAndLoss(marketIDs []string) ([]MarketProfitAndLoss, error) {
	// create url
	url := createUrl(api_betting_url, "listMarketProfitAndLoss/")

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

func (b *Betting) PlaceOrders(marketID string, placeInstructions []PlaceInstruction) (PlaceExecutionReport, error) {
	// create url
	url := createUrl(api_betting_url, "placeOrders/")
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
