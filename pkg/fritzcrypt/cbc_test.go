package fritzcrypt_test

import (
	"testing"

	"github.com/antiphp/fritzgo/pkg/fritzcrypt"
	"github.com/stretchr/testify/require"
)

func TestCbcAlgo_Decrypt(t *testing.T) {
	_, err := fritzcrypt.New([]byte("foobar"), []byte("$$$$YMBBRA2XMB4GZB6I1WPIAWINEWBXKD2PIYQHHAL3Y4Y1CZ42PDRDFDQHM65FVLI51NS1OTX2F3U4QRKFREZZZ3FRIBL3I42TJJPKZVQA"))
	require.NoError(t, err)
}
