package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type DevelopmentModeListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result DevelopmentMode `json:"result"`
	JSON   DevelopmentModeListResponseJSON
}

type DevelopmentModeListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DevelopmentModeListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DevelopmentModeListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DevelopmentModeUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result DevelopmentMode `json:"result"`
	JSON   DevelopmentModeUpdateResponseJSON
}

type DevelopmentModeUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DevelopmentModeUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DevelopmentModeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
