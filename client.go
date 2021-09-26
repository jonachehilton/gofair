package gofair

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/belmegatron/gofair/config"
	"github.com/belmegatron/gofair/streaming"
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
func NewClient(cfg *config.Config, streamingEndpoint string) (*Client, error) {

	client := new(Client)
	client.Session = new(Session)

	cert, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return nil, err
	}

	client.Certificates = &cert

	// set config
	client.Config = cfg

	// create betting
	client.Betting = &Betting{Client: client}

	// create account
	client.Account = &Account{Client: client}

	// create streaming
	stream, err := streaming.NewStream(streamingEndpoint, client.Certificates, cfg.AppKey)
	if err != nil {
		return nil, err
	}

	client.Streaming = stream

	return client, nil
}

// SessionExpired returns True if client not logged in or expired, betfair requires keep alive every 4hrs (20mins ITA)
func (c *Client) SessionExpired() bool {
	if c.Session.SessionToken == "" {
		return true
	}
	duration := time.Since(c.Session.LoginTime)

	return duration.Minutes() > 200
}
