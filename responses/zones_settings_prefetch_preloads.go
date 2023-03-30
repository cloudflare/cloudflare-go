package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type PrefetchPreloadListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result PrefetchPreload `json:"result"`
	JSON   PrefetchPreloadListResponseJSON
}

type PrefetchPreloadListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrefetchPreloadListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PrefetchPreloadListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PrefetchPreloadUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result PrefetchPreload `json:"result"`
	JSON   PrefetchPreloadUpdateResponseJSON
}

type PrefetchPreloadUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrefetchPreloadUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PrefetchPreloadUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
