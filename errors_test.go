package cloudflare

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIRequestError_Error(t *testing.T) {
	tests := map[string]struct {
		response []ResponseInfo
		want     string
	}{
		"basic complete response": {
			response: []ResponseInfo{{
				Code:    10000,
				Message: "Authentication error",
			}},
			want: "HTTP status 400: Authentication error (10000)",
		},
		"multiple complete response": {
			response: []ResponseInfo{
				{
					Code:    10000,
					Message: "Authentication error",
				},
				{
					Code:    10001,
					Message: "Not authentication error",
				},
			},
			want: "HTTP status 400: Authentication error (10000), Not authentication error (10001)",
		},
		"empty errors payload": {
			response: []ResponseInfo{},
			want:     "HTTP status 400",
		},
		"missing internal error code": {
			response: []ResponseInfo{{
				Message: "something is broke",
			}},
			want: "HTTP status 400: something is broke",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &APIRequestError{
				StatusCode: 400,
				Errors:     tc.response,
			}

			assert.Equal(t, got.Error(), tc.want)
		})
	}
}

func TestAPIRequestError_HTTPStatusCode(t *testing.T) {
	err := &APIRequestError{StatusCode: 999}
	assert.Equal(t, err.HTTPStatusCode(), 999)
}

func TestAPIRequestError_InternalErrorCodes(t *testing.T) {
	tests := map[string]struct {
		response []ResponseInfo
		want     []int
	}{
		"single": {
			response: []ResponseInfo{
				{Code: 1234},
			},
			want: []int{1234},
		},
		"multiple": {
			response: []ResponseInfo{
				{Code: 1234},
				{Code: 4321},
			},
			want: []int{1234, 4321},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &APIRequestError{
				StatusCode: 999,
				Errors:     tc.response,
			}

			assert.Equal(t, got.InternalErrorCodes(), tc.want)
		})
	}
}

func TestAPIRequestError_ErrorMessages(t *testing.T) {
	tests := map[string]struct {
		response []ResponseInfo
		want     []string
	}{
		"single": {
			response: []ResponseInfo{
				{Message: "broke once"},
			},
			want: []string{"broke once"},
		},
		"multiple": {
			response: []ResponseInfo{
				{Message: "broke once"},
				{Message: "broke twice"},
			},
			want: []string{"broke once", "broke twice"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &APIRequestError{
				StatusCode: 999,
				Errors:     tc.response,
			}

			assert.Equal(t, got.ErrorMessages(), tc.want)
		})
	}
}

func TestAPIRequestError_ServiceError(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		500: {want: true},
		599: {want: true},
		600: {want: false},
	}

	for name, tc := range tests {
		t.Run(fmt.Sprintf("%s", strconv.Itoa(name)), func(t *testing.T) {
			got := &APIRequestError{StatusCode: name}
			assert.Equal(t, got.ServiceError(), tc.want)
		})
	}
}

func TestAPIRequestError_ClientError(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		400: {want: true},
		499: {want: true},
		500: {want: false},
	}

	for name, tc := range tests {
		t.Run(fmt.Sprintf("%s", strconv.Itoa(name)), func(t *testing.T) {
			got := &APIRequestError{StatusCode: name}
			assert.Equal(t, got.ClientError(), tc.want)
		})
	}
}

func TestAPIRequestError_ClientRateLimited(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		429: {want: true},
		499: {want: false},
		500: {want: false},
	}

	for name, tc := range tests {
		t.Run(fmt.Sprintf("%s", strconv.Itoa(name)), func(t *testing.T) {
			got := &APIRequestError{StatusCode: name}
			assert.Equal(t, got.ClientRateLimited(), tc.want)
		})
	}
}

func TestAPIRequestError_InternalErrorCodeIs(t *testing.T) {
	err := &APIRequestError{Errors: []ResponseInfo{
		{Code: 1001},
		{Code: 2001},
		{Code: 3001},
	}}
	assert.Equal(t, err.InternalErrorCodeIs(3001), true)
}

func TestAPIRequestError_ErrorMessageContains(t *testing.T) {
	err := &APIRequestError{Errors: []ResponseInfo{
		{Message: "dns thing broke"},
		{Message: "application thing broke"},
		{Message: "network thing broke"},
	}}
	assert.Equal(t, err.ErrorMessageContains("application thing broke"), true)
}
