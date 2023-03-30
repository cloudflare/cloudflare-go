package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type PrivacyPassListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result PrivacyPass `json:"result"`
	JSON   PrivacyPassListResponseJSON
}

type PrivacyPassListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrivacyPassListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PrivacyPassListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PrivacyPassUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Privacy Pass is a browser extension developed by the Privacy Pass Team to
	// improve the browsing experience for your visitors. Enabling Privacy Pass will
	// reduce the number of CAPTCHAs shown to your visitors.
	// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
	Result PrivacyPass `json:"result"`
	JSON   PrivacyPassUpdateResponseJSON
}

type PrivacyPassUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PrivacyPassUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *PrivacyPassUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
