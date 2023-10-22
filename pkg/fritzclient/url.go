package fritzclient

import (
	"net/url"
	"strconv"
)

// URLBuilderImpl builds URLs in a fluent interface style.
type URLBuilderImpl struct {
	URL *url.URL
}

// WithPort applies the port.
func (u *URLBuilderImpl) WithPort(port uint16) URLBuilder {
	clone := *u.URL
	clone.Host = clone.Hostname() + ":" + strconv.Itoa(int(port))
	return &URLBuilderImpl{&clone}
}

// WithPath applies the path.
func (u *URLBuilderImpl) WithPath(path string) URLBuilder {
	return &URLBuilderImpl{u.URL.JoinPath(path)}
}

// WithQuery applies a query key-value pair.
func (u *URLBuilderImpl) WithQuery(key, value string) URLBuilder {
	q := u.URL.Query()
	q.Set(key, value)

	clone := *u.URL
	clone.RawQuery = q.Encode()
	return &URLBuilderImpl{&clone}
}

// String returns the URL as a string.
func (u *URLBuilderImpl) String() string {
	return u.URL.String()
}
