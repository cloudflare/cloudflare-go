package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type AdvancedDdoListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
	// website. This is an uneditable value that is 'on' in the case of Business and
	// Enterprise zones.
	Result AdvancedDdos `json:"result"`
	JSON   AdvancedDdoListResponseJSON
}

type AdvancedDdoListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into AdvancedDdoListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AdvancedDdoListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
