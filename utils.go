package gofair

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func createURL(endpoint string, method string) string {
	return endpoint + method
}

type Detail struct{}

type ErrorResponse struct {
	//b'{"faultcode":"Client","faultstring":"DSC-0018","detail":{}}'
	FaultCode   string  `json:"faultcode"`
	FaultString string  `json:"faultstring"`
	Detail      *Detail `json:"detail"`
}

func logError(data []byte) error {
	var errorResp ErrorResponse
	if err := json.Unmarshal(data, &errorResp); err != nil {
		return err
	}
	log.Println("Error:", errorResp, errorResp.Detail)
	return nil
}

func (params *Params) MarshalJSON() ([]byte, error) {

	placeInstructionCount := len(params.PlaceInstructions)
	cancelInstructionCount := len(params.CancelInstructions)

	if (placeInstructionCount > 0) && (cancelInstructionCount == 0) {
		type alias struct {
			MarketID            string              `json:"marketId,omitempty"`
			MarketIDs           []string            `json:"marketIds,omitempty"`
			MarketFilter        *MarketFilter       `json:"filter,omitempty"`
			MaxResults          int                 `json:"maxResults,omitempty"`
			Granularity         string              `json:"granularity,omitempty"`
			MarketProjection    []string            `json:"marketProjection,omitempty"`
			OrderProjection     orderProjection     `json:"orderProjection,omitempty"`
			MatchProjection     matchProjection     `json:"matchProjection,omitempty"`
			PriceProjection     *PriceProjection    `json:"priceProjection,omitempty"`
			Sort                string              `json:"sort,omitempty"`
			Locale              string              `json:"locale,omitempty"`
			IsMarketDataDelayed bool                `json:"isMarketDataDelayed,omitempty"`
			PlaceInstructions   []PlaceInstruction  `json:"instructions,omitempty"`
			CancelInstructions  []CancelInstruction `json:"cancelInstructions,omitempty"`
		}

		var a alias = alias(*params)
		return json.Marshal(&a)

	} else if (cancelInstructionCount > 0) && (placeInstructionCount == 0) {
		type alias struct {
			MarketID            string              `json:"marketId,omitempty"`
			MarketIDs           []string            `json:"marketIds,omitempty"`
			MarketFilter        *MarketFilter       `json:"filter,omitempty"`
			MaxResults          int                 `json:"maxResults,omitempty"`
			Granularity         string              `json:"granularity,omitempty"`
			MarketProjection    []string            `json:"marketProjection,omitempty"`
			OrderProjection     orderProjection     `json:"orderProjection,omitempty"`
			MatchProjection     matchProjection     `json:"matchProjection,omitempty"`
			PriceProjection     *PriceProjection    `json:"priceProjection,omitempty"`
			Sort                string              `json:"sort,omitempty"`
			Locale              string              `json:"locale,omitempty"`
			IsMarketDataDelayed bool                `json:"isMarketDataDelayed,omitempty"`
			PlaceInstructions   []PlaceInstruction  `json:"placeInstructions,omitempty"`
			CancelInstructions  []CancelInstruction `json:"instructions,omitempty"`
		}

		var a alias = alias(*params)
		return json.Marshal(&a)

	} else if (cancelInstructionCount == 0) && (placeInstructionCount == 0) {
		return json.Marshal(params)
	} else {
		empty := make([]byte, 0)
		return empty, errors.New("Multiple Instructions found")
	}
}

func (b *Betting) Request(url string, params *Params, v interface{}) error {
	//params.Locale = b.Client.config.Locale

	bytes, err := params.MarshalJSON()
	if err != nil {
		return err
	}

	body := strings.NewReader(string(bytes))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}

	// set headers
	req.Header.Set("X-Application", b.Client.config.AppKey)
	req.Header.Set("X-Authentication", b.Client.session.SessionToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		logError(data)
		return errors.New(resp.Status)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
