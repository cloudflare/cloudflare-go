package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type AutomaticHTTPsRewriteListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result AutomaticHTTPsRewrites `json:"result"`
	JSON   AutomaticHTTPsRewriteListResponseJSON
}

type AutomaticHTTPsRewriteListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AutomaticHTTPsRewriteListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *AutomaticHTTPsRewriteListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type AutomaticHTTPsRewriteUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable the Automatic HTTPS Rewrites feature for this zone.
	Result AutomaticHTTPsRewrites `json:"result"`
	JSON   AutomaticHTTPsRewriteUpdateResponseJSON
}

type AutomaticHTTPsRewriteUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// AutomaticHTTPsRewriteUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *AutomaticHTTPsRewriteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
