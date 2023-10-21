package fritzcrypt_test

import (
	"testing"

	"github.com/antiphp/fritzgo/pkg/fritzcrypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	type pair struct {
		cipher    []byte
		wantPlain string
	}

	tests := []struct {
		name           string
		userPassphrase []byte
		masterCipher   []byte
		decrypt        []pair
	}{
		{
			name:           "decrypting with master key",
			userPassphrase: []byte("foobar"),
			masterCipher:   []byte("$$$$YMBBRA2XMB4GZB6I1WPIAWINEWBXKD2PIYQHHAL3Y4Y1CZ42PDRDFDQHM65FVLI51NS1OTX2F3U4QRKFREZZZ3FRIBL3I42TJJPKZVQA"),
			decrypt: []pair{
				{
					cipher:    []byte("$$$$5MSQMRYHZ62MJLMMKUCQQOJ5QPNPSQTN3SHTFPXGEWDVORINLWOIA3AQXMZ5QWQBJNFLCIZQP45VIAAA"),
					wantPlain: "anonymous@t-online.de",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			algo, err := fritzcrypt.New(test.userPassphrase, test.masterCipher)
			require.NoError(t, err, "Could not select algorithm and decrypt master cipher")

			for i, testcase := range test.decrypt {
				plaintext, err := algo.Decrypt(testcase.cipher)
				require.NoErrorf(t, err, "Could not decrypt testcase '%d'", i)
				assert.Equal(t, testcase.wantPlain, string(plaintext))
			}
		})
	}

}
