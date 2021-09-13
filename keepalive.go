package gofair

import (
	"encoding/json"
	"time"
)

type KeepAliveResult struct {
	SessionToken string `json:"sessionToken"`
	Token        string `json:"token"`
	Status       string `json:"status"`
	Error        string `json:"error"`
}

func (c *Client) KeepAlive() (KeepAliveResult, error) {
	// build url
	url := createURL(Endpoints.Identity, "keepAlive")

	// make request
	resp, err := logoutRequest(c, url)
	if err != nil {
		return *new(KeepAliveResult), err
	}

	var result KeepAliveResult

	// parse json
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return result, err
	}

	c.Session.SessionToken = result.Token
	c.Session.LoginTime = time.Now().UTC()
	return result, nil
}
