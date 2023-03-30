package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type HotlinkProtectionListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result HotlinkProtection `json:"result"`
	JSON   HotlinkProtectionListResponseJSON
}

type HotlinkProtectionListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into HotlinkProtectionListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *HotlinkProtectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type HotlinkProtectionUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result HotlinkProtection `json:"result"`
	JSON   HotlinkProtectionUpdateResponseJSON
}

type HotlinkProtectionUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// HotlinkProtectionUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *HotlinkProtectionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
