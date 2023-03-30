package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type Http2ListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP2 enabled for this zone.
	Result Http2 `json:"result"`
	JSON   Http2ListResponseJSON
}

type Http2ListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http2ListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Http2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Http2UpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP2 enabled for this zone.
	Result Http2 `json:"result"`
	JSON   Http2UpdateResponseJSON
}

type Http2UpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Http2UpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *Http2UpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
