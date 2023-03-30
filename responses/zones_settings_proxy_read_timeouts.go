package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ProxyReadTimeoutListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Maximum time between two read operations from origin.
	Result ProxyReadTimeout `json:"result"`
	JSON   ProxyReadTimeoutListResponseJSON
}

type ProxyReadTimeoutListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ProxyReadTimeoutListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ProxyReadTimeoutListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ProxyReadTimeoutUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Maximum time between two read operations from origin.
	Result ProxyReadTimeout `json:"result"`
	JSON   ProxyReadTimeoutUpdateResponseJSON
}

type ProxyReadTimeoutUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ProxyReadTimeoutUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ProxyReadTimeoutUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
