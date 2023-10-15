package fritzclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/antiphp/fritzgo/pkg/fritzclient/middleware"
	"github.com/hamba/logger/v2"
)

type Client struct {
	http http.RoundTripper
	url  *urlBuilder

	log *logger.Logger
}

func New(addr, user, pass string) (*Client, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("creating url: %w", err)
	}

	var rt http.RoundTripper
	rt = http.DefaultTransport
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
