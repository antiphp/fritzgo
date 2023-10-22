// Package fritzclient contains an HTTP client to access the fritz box.
package fritzclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/antiphp/fritzgo/pkg/fritzclient/middleware"
)

// URLBuilder represents a URL builder with a fluent interface.
type URLBuilder interface {
	WithPort(uint16) URLBuilder
	WithPath(string) URLBuilder
	WithQuery(string, string) URLBuilder
	String() string
}

// OptFunc applies optional client configuration.
type OptFunc func(*Client)

// WithURLBuilder applies a custom URL builder.
//
// The main purpose is testing.
func WithURLBuilder(ub URLBuilder) OptFunc {
	return func(c *Client) {
		c.url = ub
	}
}

// Client is an HTTP client to access the FRITZ!Box.
type Client struct {
	http http.RoundTripper
	url  URLBuilder
}

// New returns a new FRITZ! HTTP client.
func New(addr, user, pass string, opts ...OptFunc) (*Client, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("creating url: %w", err)
	}

	rt := http.DefaultTransport
	if user != "" {
		rt = middleware.WithBasicAuth(rt, user, pass)
	}

	cl := Client{
		http: rt,
		url:  &URLBuilderImpl{u},
	}

	for _, opt := range opts {
		opt(&cl)
	}

	return &cl, nil
}

// HTTPError represents an undesired HTTP response.
type HTTPError struct {
	Status int
	URL    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("unexpected http response: %d %s, expected 2xx for url: %s", e.Status, http.StatusText(e.Status), e.URL)
}
