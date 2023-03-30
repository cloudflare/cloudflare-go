package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type OpportunisticOnionListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result OpportunisticOnion `json:"result"`
	JSON   OpportunisticOnionListResponseJSON
}

type OpportunisticOnionListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OpportunisticOnionListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *OpportunisticOnionListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OpportunisticOnionUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
	// connection to use our onion services instead of exit nodes.
	Result OpportunisticOnion `json:"result"`
	JSON   OpportunisticOnionUpdateResponseJSON
}

type OpportunisticOnionUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OpportunisticOnionUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *OpportunisticOnionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
