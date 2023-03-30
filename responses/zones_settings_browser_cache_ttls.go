package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type BrowserCacheTtlListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result BrowserCacheTtl `json:"result"`
	JSON   BrowserCacheTtlListResponseJSON
}

type BrowserCacheTtlListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCacheTtlListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BrowserCacheTtlListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrowserCacheTtlUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result BrowserCacheTtl `json:"result"`
	JSON   BrowserCacheTtlUpdateResponseJSON
}

type BrowserCacheTtlUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrowserCacheTtlUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BrowserCacheTtlUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
