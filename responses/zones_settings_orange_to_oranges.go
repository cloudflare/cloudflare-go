package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type OrangeToOrangeListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result OrangeToOrange `json:"result"`
	JSON   OrangeToOrangeListResponseJSON
}

type OrangeToOrangeListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OrangeToOrangeListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *OrangeToOrangeListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OrangeToOrangeUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
	// on Cloudflare.
	Result OrangeToOrange `json:"result"`
	JSON   OrangeToOrangeUpdateResponseJSON
}

type OrangeToOrangeUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into OrangeToOrangeUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *OrangeToOrangeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
