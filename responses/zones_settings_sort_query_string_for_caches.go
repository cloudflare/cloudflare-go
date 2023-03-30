package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type SortQueryStringForCachListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result SortQueryStringForCache `json:"result"`
	JSON   SortQueryStringForCachListResponseJSON
}

type SortQueryStringForCachListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SortQueryStringForCachListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SortQueryStringForCachListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type SortQueryStringForCachUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result SortQueryStringForCache `json:"result"`
	JSON   SortQueryStringForCachUpdateResponseJSON
}

type SortQueryStringForCachUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// SortQueryStringForCachUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *SortQueryStringForCachUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
