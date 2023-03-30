package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ImageResizingListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ImageResizing `json:"result"`
	JSON   ImageResizingListResponseJSON
}

type ImageResizingListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ImageResizingListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ImageResizingListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ImageResizingUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
	// Image Resizing provides on-demand resizing, conversion and optimisation for
	// images served through Cloudflare's network. Refer to the
	// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
	// more information.
	Result ImageResizing `json:"result"`
	JSON   ImageResizingUpdateResponseJSON
}

type ImageResizingUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ImageResizingUpdateResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ImageResizingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
