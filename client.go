package gofair

import (
	"crypto/tls"
	"strings"
	"time"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/belmegatron/gofair/streaming"
	"github.com/belmegatron/gofair/config"

)

type Session struct {
	SessionToken string
	LoginTime    time.Time
}

// Client object
type Client struct {
	Config       *config.Config
	Session      *Session
	Certificates *tls.Certificate
	Betting      *Betting
	Account      *Account
	Streaming    *streaming.Stream
}

func createURL(endpoint string, method string) string {
	return endpoint + method
}

// Request issues a HTTP POST to the Betfair Exchange API Endpoint specified.
func (c *Client) request(url string, params interface{}, v interface{}) error {

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
	req.Header.Set("X-Application", c.Config.AppKey)
	req.Header.Set("X-Authentication", c.Session.SessionToken)
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
		return errors.New(resp.Status)
	}
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

// NewClient creates a new Betfair client.
func NewClient(cfg *config.Config) (*Client, error) {

	c := new(Client)
	c.Session = new(Session)

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	c.Certificates = &cert

	// set config
	c.Config = cfg

	// create betting
	c.Betting = &Betting{Client: c}

	// create account
	c.Account = &Account{Client: c}

	// create streaming
	c.Streaming = &streaming.Stream{}

	return c, nil
}

// SessionExpired returns True if client not logged in or expired, betfair requires keep alive every 4hrs (20mins ITA)
func (c *Client) SessionExpired() bool {
	if c.Session.SessionToken == "" {
		return true
	}
	duration := time.Since(c.Session.LoginTime)

	return duration.Minutes() > 200
}
