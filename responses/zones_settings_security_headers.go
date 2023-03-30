package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type SecurityHeaderListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare security header for a zone.
	Result SecurityHeader `json:"result"`
	JSON   SecurityHeaderListResponseJSON
}

type SecurityHeaderListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityHeaderListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SecurityHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SecurityHeaderUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare security header for a zone.
	Result SecurityHeader `json:"result"`
	JSON   SecurityHeaderUpdateResponseJSON
}

type SecurityHeaderUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into SecurityHeaderUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *SecurityHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
