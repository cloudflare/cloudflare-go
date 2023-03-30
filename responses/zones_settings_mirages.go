package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type MirageListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result Mirage `json:"result"`
	JSON   MirageListResponseJSON
}

type MirageListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MirageListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MirageListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MirageUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result Mirage `json:"result"`
	JSON   MirageUpdateResponseJSON
}

type MirageUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MirageUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MirageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
