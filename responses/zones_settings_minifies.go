package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type MinifyListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result Minify `json:"result"`
	JSON   MinifyListResponseJSON
}

type MinifyListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinifyListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MinifyListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type MinifyUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result Minify `json:"result"`
	JSON   MinifyUpdateResponseJSON
}

type MinifyUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into MinifyUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *MinifyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
