package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type WebpListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result Webp `json:"result"`
	JSON   WebpListResponseJSON
}

type WebpListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WebpListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WebpListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type WebpUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result Webp `json:"result"`
	JSON   WebpUpdateResponseJSON
}

type WebpUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into WebpUpdateResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *WebpUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
