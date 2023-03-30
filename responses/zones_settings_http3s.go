package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type Http3ListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP3 enabled for this zone.
	Result Http3 `json:"result"`
	JSON   Http3ListResponseJSON
}

type Http3ListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http3ListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Http3ListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Http3UpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP3 enabled for this zone.
	Result Http3 `json:"result"`
	JSON   Http3UpdateResponseJSON
}

type Http3UpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http3UpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Http3UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
