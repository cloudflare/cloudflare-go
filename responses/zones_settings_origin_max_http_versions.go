package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type OriginMaxHTTPVersionListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                 `json:"success"`
	Result  OriginMaxHTTPVersion `json:"result"`
	JSON    OriginMaxHTTPVersionListResponseJSON
}

type OriginMaxHTTPVersionListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OriginMaxHTTPVersionListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *OriginMaxHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OriginMaxHTTPVersionUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool                 `json:"success"`
	Result  OriginMaxHTTPVersion `json:"result"`
	JSON    OriginMaxHTTPVersionUpdateResponseJSON
}

type OriginMaxHTTPVersionUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OriginMaxHTTPVersionUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *OriginMaxHTTPVersionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
