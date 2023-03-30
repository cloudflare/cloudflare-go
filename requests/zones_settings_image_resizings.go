package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ImageResizingListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result fields.Field[ImageResizing] `json:"result"`
}

// MarshalJSON serializes ImageResizingListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ImageResizingListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ImageResizingListResponse) String() (result string) {
	return fmt.Sprintf("&ImageResizingListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ImageResizingUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result fields.Field[ImageResizing] `json:"result"`
}

// MarshalJSON serializes ImageResizingUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ImageResizingUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ImageResizingUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ImageResizingUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ImageResizingUpdateParams struct {
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Value fields.Field[ImageResizing] `json:"value,required"`
}

// MarshalJSON serializes ImageResizingUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ImageResizingUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ImageResizingUpdateParams) String() (result string) {
	return fmt.Sprintf("&ImageResizingUpdateParams{Value:%s}", r.Value)
}
