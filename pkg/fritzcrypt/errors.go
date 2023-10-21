package fritzcrypt

import "errors"

var (
	// ErrNotCiphertext represents an error where the given ciphertext is not a ciphertext.
	ErrNotCiphertext = errors.New("password must start with $$$$")

	// ErrIVLength represents and error of the initial vector size.
	ErrIVLength = errors.New("crypt iv too short")

	// ErrDataLength represents an error with the encryption data.
	ErrDataLength = errors.New("crypt data too short")

	// ErrInvalidChecksum represents an error with the checksum, most likely a wrong password.
	ErrInvalidChecksum = errors.New("invalid checksum: wrong ciphertext, wrong key, or wrong algorithm")
)

// BaseDecodingError represents a decoding error.
type BaseDecodingError struct {
	Err error
}

// Error returns the error message.
func (e *BaseDecodingError) Error() string {
	if e.Err == nil {
		return "base decoding: error not set"
	}
	return "base decoding: " + e.Err.Error()
}

// Unwrap unwraps wrapped errors.
func (e *BaseDecodingError) Unwrap() error {
	return e.Err
}

// CipherError represents a cipher error.
type CipherError struct {
	Call string
	Err  error
}

func (e *CipherError) Error() string {
	switch {
	case e.Call == "" && (e.Err == nil || e.Err.Error() == ""):
		return "calling cipher [-]: error [-]"
	case e.Call == "":
		return "calling cipher [-]: " + e.Err.Error()
	case e.Err == nil || e.Err.Error() == "":
		return "calling cipher " + e.Call + ": error [-]"
	default:
		return "calling cipher " + e.Call + ": " + e.Error()
	}
}

// Unwrap unwraps wrapped errors.
func (e *CipherError) Unwrap() error {
	return e.Err
}

// DecodeSizeError represents a decoding error.
type DecodeSizeError struct {
	Err error
}

// Error returns the error message.
func (e *DecodeSizeError) Error() string {
	if e.Err == nil {
		return "decoding size: error not set"
	}
	return "decoding size: " + e.Err.Error()
}

// Unwrap unwraps wrapped errors.
func (e *DecodeSizeError) Unwrap() error {
	return e.Err
}
