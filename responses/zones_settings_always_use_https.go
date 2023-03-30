package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type AlwaysUseHTTPListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result AlwaysUseHTTPs `json:"result"`
	JSON   AlwaysUseHTTPListResponseJSON
}

type AlwaysUseHTTPListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysUseHTTPListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AlwaysUseHTTPListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AlwaysUseHTTPUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Reply to all requests for URLs that use "http" with a 301 redirect to the
	// equivalent "https" URL. If you only want to redirect for a subset of requests,
	// consider creating an "Always use HTTPS" page rule.
	Result AlwaysUseHTTPs `json:"result"`
	JSON   AlwaysUseHTTPUpdateResponseJSON
}

type AlwaysUseHTTPUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysUseHTTPUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AlwaysUseHTTPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
