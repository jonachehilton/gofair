package gofair

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type LogoutResult struct {
	Token   string `json:"token"`
	Product string `json:"product"`
	Status  string `json:"status"`
	Error   string `json:"error"`
}

// Logout from the current session.
func (c *Client) Logout() (LogoutResult, error) {
	// build url
	url := createURL(Endpoints.Identity, "logout")

	// make request
	resp, err := logoutRequest(c, url)
	if err != nil {
		return *new(LogoutResult), err
	}

	var result LogoutResult

	// parse json
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	c.Session.SessionToken = ""
	c.Session.LoginTime = time.Time{}
	return result, nil
}

func logoutRequest(c *Client, url string) ([]byte, error) {

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Application", c.Config.AppKey)
	req.Header.Set("X-Authentication", c.Session.SessionToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
