package disgo

import (
	"net/http"
	"net/url"
	"time"
)

// Client for interaction with Discourse HTTP API.
type Client struct {
	// URI of a Discourse to use
	Endpoint string
	// Username admin who make the api key.
	Username string
	// APIKey generated API key.
	APIKey  string
	host      string
	transport http.RoundTripper
	timeout   time.Duration
}

// NewClient instantiates a client.
func NewClient(uri string, username string, apiKey string) (me *Client, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	me = &Client{
		Endpoint: uri,
		host:     u.Host,
		Username: username,
		APIKey: apiKey,
	}

	return me, nil
}

// SetTimeout changes the HTTP timeout that the Client will use.
// By default there is no timeout.
func (c *Client) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}