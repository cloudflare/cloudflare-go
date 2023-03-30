package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type MobileRedirectListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result MobileRedirect `json:"result"`
	JSON   MobileRedirectListResponseJSON
}

type MobileRedirectListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MobileRedirectListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *MobileRedirectListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MobileRedirectUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result MobileRedirect `json:"result"`
	JSON   MobileRedirectUpdateResponseJSON
}

type MobileRedirectUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MobileRedirectUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *MobileRedirectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
