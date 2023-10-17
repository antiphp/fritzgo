package fritzclient

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/antiphp/fritzgo/pkg/fritztypes"
)

type tr64desc struct {
	SystemVersion struct {
		Display string `xml:"Display"` // Capital D.
	} `xml:"systemVersion"`
	Device struct {
		FriendlyName    string `xml:"friendlyName"`
		SerialNumber    string `xml:"serialNumber"`
		PresentationURL string `xml:"presentationURL"`
	} `xml:"device"`
}

// Info returns fritz basic information.
func (c *Client) Info(ctx context.Context) (fritztypes.Info, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url.WithPort(49000).WithPath("/tr64desc.xml").String(), nil)
	if err != nil {
		return fritztypes.Info{}, fmt.Errorf("preparing request: %w", err)
	}

	resp, err := c.http.RoundTrip(req)
	if err != nil {
		return fritztypes.Info{}, fmt.Errorf("requesting: %w", err)
	}
	if resp.StatusCode/100 != 2 {
		return fritztypes.Info{}, &HTTPError{resp.StatusCode, req.URL.String()}
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	res := tr64desc{}
	if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
		return fritztypes.Info{}, fmt.Errorf("decoding response: %w", err)
	}

	return fritztypes.Info{
		Name:    res.Device.FriendlyName,
		Version: res.SystemVersion.Display,
		URL:     res.Device.PresentationURL,
		Mac:     res.Device.SerialNumber,
	}, nil
}
