package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type Tls_1_3ListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result Tls_1_3 `json:"result"`
	JSON   Tls_1_3ListResponseJSON
}

type Tls_1_3ListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Tls_1_3ListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Tls_1_3ListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Tls_1_3UpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enables Crypto TLS 1.3 feature for a zone.
	Result Tls_1_3 `json:"result"`
	JSON   Tls_1_3UpdateResponseJSON
}

type Tls_1_3UpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Tls_1_3UpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Tls_1_3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
