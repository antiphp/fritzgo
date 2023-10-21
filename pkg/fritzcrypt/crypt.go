// Package fritzcrypt decrypts fritz secrets.
package fritzcrypt

import (
	"crypto/md5" //nolint:gosec
	"fmt"
)

// New returns a new decrypting algorithm for the given user passphrase and master cipher.
func New(userPassphrase, masterCipher []byte) (*CBCAlgo, error) {
	hasher := md5.New() //nolint:gosec
	_, err := hasher.Write(userPassphrase)
	if err != nil {
		return nil, fmt.Errorf("md5 hashing: %w", err)
	}

	userKey := hasher.Sum(nil)
	algo := NewCBCAlgo(userKey)

	key, err := algo.Decrypt(masterCipher)
	if err != nil {
		return nil, fmt.Errorf("decrypting master cipher: %w", err)
	}

	return NewCBCAlgo(key), nil
}
