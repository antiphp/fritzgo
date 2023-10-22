package fritzclient_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
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
		require.Equal(t, "/tr64desc.xml", req.URL.RequestURI())

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

	fritz, err := fritzclient.New(srv.URL, "", "", fritzclient.WithURLBuilder(newTestBuilder(t, srv.URL, uint16(49000))))
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

// testBuilder is a URL builder with a static port assignment, so we can use whatever httptest.NewServer() provides as a URL.
type testBuilder struct {
	url        fritzclient.URLBuilder
	assertPort func(uint16)
}

func newTestBuilder(t *testing.T, u string, wantPort uint16) *testBuilder {
	parsed, err := url.Parse(u)
	require.NoError(t, err)

	return &testBuilder{
		url: &fritzclient.URLBuilderImpl{URL: parsed},
		assertPort: func(got uint16) {
			assert.Equal(t, wantPort, got)
		},
	}
}

func (b *testBuilder) WithPort(port uint16) fritzclient.URLBuilder {
	b.assertPort(port)
	return b
}

func (b *testBuilder) WithPath(s string) fritzclient.URLBuilder {
	return b.url.WithPath(s)
}

func (b *testBuilder) WithQuery(key string, value string) fritzclient.URLBuilder {
	return b.url.WithQuery(key, value)
}

func (b *testBuilder) String() string {
	return b.url.String()
}
