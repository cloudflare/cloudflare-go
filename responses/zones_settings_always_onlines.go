package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type AlwaysOnlineListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result AlwaysOnline `json:"result"`
	JSON   AlwaysOnlineListResponseJSON
}

type AlwaysOnlineListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysOnlineListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AlwaysOnlineListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AlwaysOnlineUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result AlwaysOnline `json:"result"`
	JSON   AlwaysOnlineUpdateResponseJSON
}

type AlwaysOnlineUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AlwaysOnlineUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *AlwaysOnlineUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
