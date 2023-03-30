package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type TrueClientIpHeaderListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result TrueClientIpHeader `json:"result"`
	JSON   TrueClientIpHeaderListResponseJSON
}

type TrueClientIpHeaderListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TrueClientIpHeaderListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TrueClientIpHeaderListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type TrueClientIpHeaderUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Allows customer to continue to use True Client IP (Akamai feature) in the
	// headers we send to the origin. This is limited to Enterprise Zones.
	Result TrueClientIpHeader `json:"result"`
	JSON   TrueClientIpHeaderUpdateResponseJSON
}

type TrueClientIpHeaderUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// TrueClientIpHeaderUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *TrueClientIpHeaderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
