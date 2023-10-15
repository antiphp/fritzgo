package middleware

import (
	"net/http"
	"slices"
)

type basicAuth struct {
	user string
	pass string

	skipPaths []string

	rt http.RoundTripper
}

func WithBasicAuth(rt http.RoundTripper, user, pass string, skipPaths ...string) http.RoundTripper {
	return &basicAuth{
		user:      user,
		pass:      pass,
		skipPaths: skipPaths,

		rt: rt,
	}
}

func (a *basicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	if !slices.Contains(a.skipPaths, req.URL.Path) {
		req = req.Clone(req.Context())
		req.SetBasicAuth(a.user, a.pass)
	}
	return a.rt.RoundTrip(req)
}
