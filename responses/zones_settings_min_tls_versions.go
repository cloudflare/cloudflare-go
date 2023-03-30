package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type MinTlsVersionListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result MinTlsVersion `json:"result"`
	JSON   MinTlsVersionListResponseJSON
}

type MinTlsVersionListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinTlsVersionListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *MinTlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MinTlsVersionUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Only accepts HTTPS requests that use at least the TLS protocol version
	// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
	// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
	Result MinTlsVersion `json:"result"`
	JSON   MinTlsVersionUpdateResponseJSON
}

type MinTlsVersionUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinTlsVersionUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *MinTlsVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
