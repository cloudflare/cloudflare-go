package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type EarlyHintListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result EarlyHints `json:"result"`
	JSON   EarlyHintListResponseJSON
}

type EarlyHintListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EarlyHintListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EarlyHintListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type EarlyHintUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When enabled, Cloudflare will attempt to speed up overall page loads by serving
	// `103` responses with `Link` headers from the final response. Refer to
	// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
	// more information.
	Result EarlyHints `json:"result"`
	JSON   EarlyHintUpdateResponseJSON
}

type EarlyHintUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into EarlyHintUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *EarlyHintUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
