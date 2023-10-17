// Package fritzclient contains an HTTP client to access the fritz box.
package fritzclient

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/antiphp/fritzgo/pkg/fritzclient/middleware"
	"github.com/hamba/logger/v2"
)

// Client is an HTTP client to access the FRITZ!Box.
type Client struct {
	http http.RoundTripper
	url  *urlBuilder

	log *logger.Logger //nolint:unused // Might get used in the future.
}

// New returns a new FRITZ! HTTP client.
func New(addr, user, pass string) (*Client, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("creating url: %w", err)
	}

	rt := http.DefaultTransport
	if user != "" {
		rt = middleware.WithBasicAuth(rt, user, pass)
	}

	return &Client{
		http: rt,
		url:  &urlBuilder{u},
	}, nil
}

type urlBuilder struct {
	*url.URL
}

func (u *urlBuilder) WithPort(port uint16) *urlBuilder {
	clone := *u.URL
	clone.Host = clone.Host + ":" + strconv.Itoa(int(port))
	return &urlBuilder{&clone}
}

func (u *urlBuilder) WithPath(path string) *urlBuilder {
	return &urlBuilder{u.JoinPath(path)}
}

func (u *urlBuilder) WithQuery(key, value string) *urlBuilder {
	q := u.Query()
	q.Set(key, value)

	clone := *u.URL
	clone.RawQuery = q.Encode()
	return &urlBuilder{&clone}
}

// HTTPError represents an undesired HTTP response.
type HTTPError struct {
	Status int
	URL    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("unexpected http response: %d %s, expected 2xx for url: %s", e.Status, http.StatusText(e.Status), e.URL)
}
