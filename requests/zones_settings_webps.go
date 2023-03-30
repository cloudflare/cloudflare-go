package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type WebpListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result fields.Field[Webp] `json:"result"`
}

// MarshalJSON serializes WebpListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WebpListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebpListResponse) String() (result string) {
	return fmt.Sprintf("&WebpListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WebpUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When the client requesting the image supports the WebP image codec, and WebP
	// offers a performance advantage over the original image format, Cloudflare will
	// serve a WebP version of the original image.
	Result fields.Field[Webp] `json:"result"`
}

// MarshalJSON serializes WebpUpdateResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WebpUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebpUpdateResponse) String() (result string) {
	return fmt.Sprintf("&WebpUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type WebpUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[WebpValue] `json:"value,required"`
}

// MarshalJSON serializes WebpUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *WebpUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r WebpUpdateParams) String() (result string) {
	return fmt.Sprintf("&WebpUpdateParams{Value:%s}", r.Value)
}
