package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type H2PrioritizationListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result H2Prioritization `json:"result"`
	JSON   H2PrioritizationListResponseJSON
}

type H2PrioritizationListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into H2PrioritizationListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *H2PrioritizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type H2PrioritizationUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// HTTP/2 Edge Prioritization optimises the delivery of resources served through
	// HTTP/2 to improve page load performance. It also supports fine control of
	// content delivery when used in conjunction with Workers.
	Result H2Prioritization `json:"result"`
	JSON   H2PrioritizationUpdateResponseJSON
}

type H2PrioritizationUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// H2PrioritizationUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *H2PrioritizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
