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

func TestClient_ListUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	t.Cleanup(cancel)

	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		require.Equal(t, "/login_sid.lua?version=2", req.URL.RequestURI())

		rw.WriteHeader(http.StatusOK)
		_, err := rw.Write([]byte(`
			<?xml version="1.0" encoding="utf-8"?>
			<SessionInfo>
				<SID>0000000000000000</SID>
				<Challenge>3$55000$5ab3f01c09a46e6f95fb4d7b7d00a9fd$9000$d61e25a6c9a9a9c8acdda12f979a2c8</Challenge>
				<BlockTime>0</BlockTime>
				<Rights></Rights>
				<Users>
					<User last="1">fritz9375</User>
				</Users>
			</SessionInfo>`))
		require.NoError(t, err)
	}))
	t.Cleanup(srv.Close)

	fritz, err := fritzclient.New(srv.URL, "", "")
	require.NoError(t, err)

	users, err := fritz.ListUsers(ctx)
	require.NoError(t, err)

	assert.Equal(t, []fritztypes.User{{Name: "fritz9375", Default: true}}, users)

}
