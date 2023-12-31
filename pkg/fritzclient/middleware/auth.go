// Package middleware contains routing middleware.
package middleware

import (
	"net/http"
)

type basicAuth struct {
	user string
	pass string

	rt http.RoundTripper
}

// WithBasicAuth returns a round tripper with applied basic auth.
func WithBasicAuth(rt http.RoundTripper, user, pass string) http.RoundTripper {
	return &basicAuth{
		user: user,
		pass: pass,

		rt: rt,
	}
}

// RoundTrip implements the http.RoundTripper interface.
func (a *basicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req = req.Clone(req.Context())
	req.SetBasicAuth(a.user, a.pass)

	return a.rt.RoundTrip(req)
}
