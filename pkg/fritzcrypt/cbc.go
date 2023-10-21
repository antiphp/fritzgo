package fritzcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5" //nolint:gosec
	"encoding/base32"
	"encoding/hex"
	"strconv"
)

// CBCAlgo decrypts fritz secrets using CBC, Base32, and MD5.
//
// Seems to be used for:
// - FRITZ!Box 7530 & FRITZ!OS 7.57.
type CBCAlgo struct {
	key      []byte
	encoding *base32.Encoding
}

// NewCBCAlgo returns a new CBC algorithm.
func NewCBCAlgo(key []byte) *CBCAlgo {
	// Last half of the 32-byte key is expected to be null bytes.
	buf := make([]byte, 32)
	copy(buf, key[:16])

	return &CBCAlgo{
		key: buf,

		// The standard encoding goes AB..XY..23, while this goes AB..XY..12.
		encoding: base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ123456").WithPadding(base32.NoPadding),
	}
}

// Decrypt decrypts the given ciphertext.
func (a *CBCAlgo) Decrypt(ciphertext []byte) ([]byte, error) {
	if !bytes.HasPrefix(ciphertext, []byte("$$$$")) {
		return nil, ErrNotCiphertext
	}
	ciphertext = ciphertext[4:]

	// Decode.
	decoded := make([]byte, a.encoding.DecodedLen(len(ciphertext)))
	_, err := a.encoding.Decode(decoded, ciphertext)
	if err != nil {
		return nil, &BaseDecodingError{err}
	}

	// Decrypt.
	if len(decoded) < 16 {
		return nil, ErrIVLength
	}
	iv := decoded[:16]

	src := decoded[16:]
	if len(src) == 0 {
		return nil, ErrDataLength
	}
	for len(src)%16 != 0 {
		src = append(src, byte(0))
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return nil, &CipherError{"aes.NewCipher()", err}
	}

	dst := make([]byte, len(src))

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(dst, src)

	// Check.

	h := md5.New() //nolint:gosec
	h.Write(dst[4 : 4+len(dst)-20])
	checksum := h.Sum(nil)

	want := dst[:4]
	if !bytes.HasPrefix(checksum, want) {
		return nil, ErrInvalidChecksum
	}

	// Extract.

	sizeBytes := dst[4:8]
	size, err := strconv.ParseInt(hex.EncodeToString(sizeBytes), 16, 64)
	if err != nil {
		return nil, &DecodeSizeError{err}
	}

	plaintext := bytes.TrimRight(dst[8:8+size], string([]byte{0}))
	return plaintext, nil
}
