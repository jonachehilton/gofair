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

// ErrorResponse contains error information from an unsuccessful HTTP POST request.
type ErrorResponse struct {
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

// Request issues a HTTP POST to the Betfair Exchange API Endpoint specified.
func (b *Betting) Request(url string, params interface{}, v interface{}) error {

	bytes, err := json.Marshal(params)

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

	if resp.StatusCode != http.StatusOK {
		logError(data)
		return errors.New(resp.Status)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}
