package cloudflare

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error messages
const (
	errEmptyCredentials     = "invalid credentials: key & email must not be empty"
	errMakeRequestError     = "error from makeRequest"
	errUnmarshalError       = "error unmarshalling the JSON response"
	errRequestNotSuccessful = "error reported by API"
)

// APIError represents error returned from CloudFlare API
type APIError struct {
	msg string
	// HTTP status code
	StatusCode int
	// Optional list of errors
	Errors []ResponseInfo
}

// NewAPIError creates a new APIError instance
func NewAPIError(statusCode int, respBody []byte) error {
	return &APIError{
		msg:        generateAPIErrorMessage(statusCode, respBody),
		StatusCode: statusCode,
		Errors:     parseErrors(respBody),
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("HTTP status %d: %q", e.StatusCode, e.msg)
}

// HasErrorCode checks if the errors list contains a specific error code.
func (e *APIError) HasErrorCode(errorCode int) bool {
	for _, error := range e.Errors {
		if error.Code == errorCode {
			return true
		}
	}
	return false
}

func generateAPIErrorMessage(statusCode int, respBody []byte) string {
	switch {
	case statusCode == http.StatusUnauthorized:
		return "invalid credentials"
	case statusCode == http.StatusForbidden:
		return "insufficient permissions"
	case statusCode == http.StatusServiceUnavailable,
		statusCode == http.StatusBadGateway,
		statusCode == http.StatusGatewayTimeout,
		statusCode == 522,
		statusCode == 523,
		statusCode == 524:
		return "service failure"
	default:
		var s string
		if respBody != nil {
			s = string(respBody)
		}
		return s
	}
}

func parseErrors(respBody []byte) []ResponseInfo {
	var rawResponse RawResponse
	err := json.Unmarshal(respBody, &rawResponse)
	if err != nil {
		return nil
	}
	return rawResponse.Errors
}
