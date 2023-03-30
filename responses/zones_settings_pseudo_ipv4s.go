package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type PseudoIpv4ListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// The value set for the Pseudo IPv4 setting.
	Result PseudoIpv4 `json:"result"`
	JSON   PseudoIpv4ListResponseJSON
}

type PseudoIpv4ListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PseudoIpv4ListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PseudoIpv4ListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PseudoIpv4UpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// The value set for the Pseudo IPv4 setting.
	Result PseudoIpv4 `json:"result"`
	JSON   PseudoIpv4UpdateResponseJSON
}

type PseudoIpv4UpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PseudoIpv4UpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PseudoIpv4UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
