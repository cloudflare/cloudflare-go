package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type BrotliListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result Brotli `json:"result"`
	JSON   BrotliListResponseJSON
}

type BrotliListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrotliListResponse using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BrotliListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BrotliUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// When the client requesting an asset supports the Brotli compression algorithm,
	// Cloudflare will serve a Brotli compressed version of the asset.
	Result Brotli `json:"result"`
	JSON   BrotliUpdateResponseJSON
}

type BrotliUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BrotliUpdateResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BrotliUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
