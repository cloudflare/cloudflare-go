package cloudflare

import (
	"fmt"
	"net/http"
)

// Error messages
const (
	errEmptyCredentials          = "invalid credentials: key & email must not be empty"
	errEmptyAPIToken             = "invalid credentials: API Token must not be empty"
	errMakeRequestError          = "error from makeRequest"
	errUnmarshalError            = "error unmarshalling the JSON response"
	errUnmarshalErrorBody        = "error unmarshalling the JSON response error body"
	errRequestNotSuccessful      = "error reported by API"
	errMissingAccountID          = "account ID is empty and must be provided"
	errOperationStillRunning     = "bulk operation did not finish before timeout"
	errOperationUnexpectedStatus = "bulk operation returned an unexpected status"
)

// APIRequestError is a type of error raised by API calls made by this library.
type APIRequestError struct {
	StatusCode int
	Message    string
	ErrorCode  int
}

func (e APIRequestError) Error() string {
	errString := ""
	errString += fmt.Sprintf("HTTP status %d", e.StatusCode)

	if e.Message != "" {
		errString += fmt.Sprintf(": %s", e.Message)
	}

	if e.ErrorCode != 0 {
		errString += fmt.Sprintf(" (%d)", e.ErrorCode)
	}

	return errString
}

// HTTPStatusCode exposes the HTTP status from the error response encountered.
func (e APIRequestError) HTTPStatusCode() int {
	return e.StatusCode
}

// ErrorMessage exposes the first error message from the error response
// encountered.
func (e *APIRequestError) ErrorMessage() string {
	return e.Message
}

// InternalErrorCode exposes the first internal error code from the error
// response encountered.
func (e *APIRequestError) InternalErrorCode() int {
	return e.ErrorCode
}

// ServiceError returns a boolean whether or not the raised error was caused by
// an internal service.
func (e *APIRequestError) ServiceError() bool {
	return e.StatusCode >= http.StatusInternalServerError
}

// ClientError returns a boolean whether or not the raised error was caused by
// something client side.
func (e *APIRequestError) ClientError() bool {
	return e.StatusCode >= http.StatusBadRequest &&
		e.StatusCode <= http.StatusInternalServerError &&
		e.StatusCode != http.StatusTooManyRequests
}

// ClientRateLimited returns a boolean whether or not the raised error was
// caused by too many requests from the client.
func (e *APIRequestError) ClientRateLimited() bool {
	return e.StatusCode == http.StatusTooManyRequests
}
