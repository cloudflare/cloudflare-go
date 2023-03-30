package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type NelListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result Nel `json:"result"`
	JSON   NelListResponseJSON
}

type NelListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into NelListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *NelListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type NelUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Enable Network Error Logging reporting on your zone. (Beta)
	Result Nel `json:"result"`
	JSON   NelUpdateResponseJSON
}

type NelUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into NelUpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *NelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
