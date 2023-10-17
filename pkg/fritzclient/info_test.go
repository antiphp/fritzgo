package fritzclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/antiphp/fritzgo/pkg/fritzclient"
	"github.com/antiphp/fritzgo/pkg/fritztypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClient_Info(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)

	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		require.Equal(t, "/logiln_sid.lua?version=2", req.URL.RequestURI())

		rw.WriteHeader(http.StatusOK)
		_, err := rw.Write([]byte(`
			<?xml version="1.0"?>
			<root xmlns="urn:dslforum-org:device-1-0">
				<systemVersion>
					<Display>164.07.57</Display>
				</systemVersion>
				<device>
					<friendlyName>FRITZ!Box 7530</friendlyName>
					<serialNumber>A3:7D:9B:C1:4E:2A</serialNumber>
					<presentationURL>http://fritz.box</presentationURL>
				</device>
			</root>`))
		require.NoError(t, err)
	}))
	t.Cleanup(srv.Close)

	fritz, err := fritzclient.New(srv.URL, "", "")
	require.NoError(t, err)

	info, err := fritz.Info(ctx)
	require.NoError(t, err)

	assert.Equal(t, fritztypes.Info{
		Name:    "FRITZ!Box 7530",
		Version: "164.07.57",
		URL:     "http://fritz.box",
		Mac:     "A3:7D:9B:C1:4E:2A",
	}, info)
}
