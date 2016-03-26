package cloudflare

import "errors"

var errEmptyCredentials = errors.New("invalid credentials: key & email must not be empty")

// Error represents an error returned from this library.
type Error interface {
	// Raised when user credentials or configuration is invalid.
	User() bool
	// Raised when a network error occurs.
	Network() bool
	// Contains the original (wrapped) error.
	Cause() error
}

// UserError represents a user-generated error.
type UserError struct {
	error
}

// User is a user-caused error.
func (e UserError) User() bool {
	return true
}

// Network error.
func (e UserError) Network() bool {
	return false
}

// Cause wraps the underlying error.
func (e UserError) Cause() error {
	return e.error
}
