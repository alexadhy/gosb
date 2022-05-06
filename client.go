package gosb

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// Config provides SendBird config
// contains API Token, Application ID, and API version of Sendbird
// by default it will use version 3 of the Sendbird API
type Config struct {
	APIToken           string
	ApplicationID      string
	SendbirdApiVersion int
}

// BaseURL returns base URL for any endpoints
func (c Config) baseURL() string {
	return fmt.Sprintf("https://api-%s.sendbird.com/v%d", c.ApplicationID, c.SendbirdApiVersion)
}

// Client contains default http client to use
// as well as the base url to use
type Client struct {
	appID   string
	hc      *resty.Client
	baseURL string
}

// NewClient will create a new *Client instance
func NewClient(cfg Config) *Client {
	return &Client{
		appID:   cfg.ApplicationID,
		hc:      defaultHC(cfg.APIToken),
		baseURL: cfg.baseURL(),
	}
}
