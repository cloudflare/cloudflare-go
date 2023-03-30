package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type MinifyListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result fields.Field[Minify] `json:"result"`
}

// MarshalJSON serializes MinifyListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MinifyListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinifyListResponse) String() (result string) {
	return fmt.Sprintf("&MinifyListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MinifyUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically minify certain assets for your website. Refer to
	// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
	// for more information.
	Result fields.Field[Minify] `json:"result"`
}

// MarshalJSON serializes MinifyUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MinifyUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinifyUpdateResponse) String() (result string) {
	return fmt.Sprintf("&MinifyUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MinifyUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[MinifyValue] `json:"value,required"`
}

// MarshalJSON serializes MinifyUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MinifyUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MinifyUpdateParams) String() (result string) {
	return fmt.Sprintf("&MinifyUpdateParams{Value:%s}", r.Value)
}
