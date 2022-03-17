package cloudflare

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_Error(t *testing.T) {
	tests := map[string]struct {
		response []ResponseInfo
		want     string
	}{
		"basic complete response": {
			response: []ResponseInfo{{
				Code:    10000,
				Message: "Authentication error",
			}},
			want: "Authentication error (10000)",
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
			want: "Authentication error (10000), Not authentication error (10001)",
		},
		"missing internal error code": {
			response: []ResponseInfo{{
				Message: "something is broke",
			}},
			want: "something is broke",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := &RequestError{cloudflareError: &Error{
				StatusCode: 400,
				Errors:     tc.response,
			}}

			assert.Equal(t, got.Error(), tc.want)
		})
	}
}

func TestError_HTTPStatusCode(t *testing.T) {
	err := &Error{StatusCode: 999}
	assert.Equal(t, err.HTTPStatusCode(), 999)
}

func TestError_InternalErrorCodes(t *testing.T) {
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
			got := &Error{
				StatusCode: 999,
				Errors:     tc.response,
			}

			assert.Equal(t, got.InternalErrorCodes(), tc.want)
		})
	}
}

func TestError_ErrorMessages(t *testing.T) {
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
			got := &Error{
				StatusCode: 999,
				Errors:     tc.response,
			}

			assert.Equal(t, got.ErrorMessages(), tc.want)
		})
	}
}

func TestError_ServiceError(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		500: {want: true},
		599: {want: true},
		600: {want: false},
	}

	for name, tc := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			got := &Error{StatusCode: name}
			assert.Equal(t, got.ServiceError(), tc.want)
		})
	}
}

func TestError_ClientError(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		400: {want: true},
		499: {want: true},
		500: {want: false},
	}

	for name, tc := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			got := &Error{StatusCode: name}
			assert.Equal(t, got.ClientError(), tc.want)
		})
	}
}

func TestError_ClientRateLimited(t *testing.T) {
	tests := map[int]struct {
		want bool
	}{
		429: {want: true},
		499: {want: false},
		500: {want: false},
	}

	for name, tc := range tests {
		t.Run(strconv.Itoa(name), func(t *testing.T) {
			got := &Error{StatusCode: name}
			assert.Equal(t, got.ClientRateLimited(), tc.want)
		})
	}
}

func TestError_InternalErrorCodeIs(t *testing.T) {
	err := &Error{Errors: []ResponseInfo{
		{Code: 1001},
		{Code: 2001},
		{Code: 3001},
	}}
	assert.Equal(t, err.InternalErrorCodeIs(3001), true)
}

func TestError_ErrorMessageContains(t *testing.T) {
	err := &Error{Errors: []ResponseInfo{
		{Message: "dns thing broke"},
		{Message: "application thing broke"},
		{Message: "network thing broke"},
	}}
	assert.Equal(t, err.ErrorMessageContains("application thing broke"), true)
}
