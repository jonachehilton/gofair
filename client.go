package gofair

import (
	"crypto/tls"
	"strings"
	"time"
)

// Config holds login data
type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AppKey   string `json:"api_key"`
	CertFile string `json:"ssl_cert"`
	KeyFile  string `json:"ssl_key"`
	Locale   string
}

type session struct {
	SessionToken string
	LoginTime    time.Time
}

// Client object
type Client struct {
	config       *Config
	session      *session
	certificates *tls.Certificate
	Betting      *Betting
	Account      *Account
	Streaming    *Streaming
	Historical   *Historical
}

// Betting object
type Betting struct {
	Client *Client
}

// Account object
type Account struct {
	Client *Client
}

// Streaming object
type Streaming struct {
	Client *Client
}

// Historical object
type Historical struct {
	Client *Client
}

// NewClient creates a new Betfair client.
func NewClient(config *Config) (*Client, error) {

	c := new(Client)
	c.session = new(session)

	var cert tls.Certificate
	var err error
	// create certificates
	// ----- is obviously not a path, therefore load direct from the variables
	if strings.HasPrefix(config.CertFile, "------") {
		cert, err = tls.X509KeyPair([]byte(config.CertFile), []byte(config.KeyFile))
		if err != nil {
			return nil, err
		}
	} else {
		cert, err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			return nil, err
		}
	}
	c.certificates = &cert

	// set config
	c.config = config

	// create betting
	c.Betting = new(Betting)
	c.Betting.Client = c

	// create account
	c.Account = new(Account)
	c.Account.Client = c

	// create streaming
	c.Streaming = new(Streaming)
	c.Streaming.Client = c

	// create historical
	c.Historical = new(Historical)
	c.Historical.Client = c

	return c, nil
}

// SessionExpired returns True if client not logged in or expired, betfair requires keep alive every 4hrs (20mins ITA)
func (c *Client) SessionExpired() bool {
	if c.session.SessionToken == "" {
		return true
	}
	duration := time.Since(c.session.LoginTime)
	if duration.Minutes() > 200 {
		return true
	}
	return false
}
