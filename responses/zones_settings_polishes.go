package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type PolishListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result Polish `json:"result"`
	JSON   PolishListResponseJSON
}

type PolishListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PolishListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PolishListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type PolishUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result Polish `json:"result"`
	JSON   PolishUpdateResponseJSON
}

type PolishUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into PolishUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *PolishUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
