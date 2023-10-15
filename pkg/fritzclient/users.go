package fritzclient

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/antiphp/fritzgo/pkg/fritztypes"
)

type loginSIDResponse struct {
	SID       string `xml:"SID"`
	Challenge string `xml:"Challenge"`
	Users     struct {
		User []struct {
			Name string `xml:",chardata"`
			Last uint8  `xml:"last,attr"`
		} `xml:"User"`
	} `xml:"Users"`
}

func (c *Client) ListUsers(ctx context.Context) ([]fritztypes.User, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url.WithPath("login_sid.lua").WithQuery("version", "2").String(), nil)
	if err != nil {
		return nil, fmt.Errorf("preparing request: %w", err)
	}

	resp, err := c.http.RoundTrip(req)
	if err != nil {
		return nil, fmt.Errorf("requesting: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("unexpected http response: %d %s, expected 2xx for url: %s", resp.StatusCode, http.StatusText(resp.StatusCode), req.URL.String())
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	res := loginSIDResponse{}
	if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	users := make([]fritztypes.User, len(res.Users.User))
	for i, user := range res.Users.User {
		users[i] = fritztypes.User{
			Name:    user.Name,
			Default: user.Last == 1,
		}
	}
	return users, nil
}
