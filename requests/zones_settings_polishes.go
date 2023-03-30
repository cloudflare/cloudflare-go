package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type PolishListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result fields.Field[Polish] `json:"result"`
}

// MarshalJSON serializes PolishListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PolishListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PolishListResponse) String() (result string) {
	return fmt.Sprintf("&PolishListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PolishUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Result fields.Field[Polish] `json:"result"`
}

// MarshalJSON serializes PolishUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PolishUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PolishUpdateResponse) String() (result string) {
	return fmt.Sprintf("&PolishUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PolishUpdateParams struct {
	// Removes metadata and compresses your images for faster page load times. Basic
	// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
	// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
	// image loading. Larger JPEGs are converted to progressive images, loading a
	// lower-resolution image first and ending in a higher-resolution version. Not
	// recommended for hi-res photography sites.
	Value fields.Field[Polish] `json:"value,required"`
}

// MarshalJSON serializes PolishUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *PolishUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PolishUpdateParams) String() (result string) {
	return fmt.Sprintf("&PolishUpdateParams{Value:%s}", r.Value)
}
