package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type OriginErrorPagePassThrusListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result OriginErrorPagePassThru `json:"result"`
	JSON   OriginErrorPagePassThrusListResponseJSON
}

type OriginErrorPagePassThrusListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OriginErrorPagePassThrusListResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *OriginErrorPagePassThrusListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type OriginErrorPagePassThrusUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result OriginErrorPagePassThru `json:"result"`
	JSON   OriginErrorPagePassThrusUpdateResponseJSON
}

type OriginErrorPagePassThrusUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// OriginErrorPagePassThrusUpdateResponse using the internal pjson library.
// Unrecognized fields are stored in the `jsonFields` property.
func (r *OriginErrorPagePassThrusUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
