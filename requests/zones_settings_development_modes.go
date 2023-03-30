package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type DevelopmentModeListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result fields.Field[DevelopmentMode] `json:"result"`
}

// MarshalJSON serializes DevelopmentModeListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DevelopmentModeListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DevelopmentModeListResponse) String() (result string) {
	return fmt.Sprintf("&DevelopmentModeListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type DevelopmentModeUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Development Mode temporarily allows you to enter development mode for your
	// websites if you need to make changes to your site. This will bypass Cloudflare's
	// accelerated cache and slow down your site, but is useful if you are making
	// changes to cacheable content (like images, css, or JavaScript) and would like to
	// see those changes right away. Once entered, development mode will last for 3
	// hours and then automatically toggle off.
	Result fields.Field[DevelopmentMode] `json:"result"`
}

// MarshalJSON serializes DevelopmentModeUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *DevelopmentModeUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DevelopmentModeUpdateResponse) String() (result string) {
	return fmt.Sprintf("&DevelopmentModeUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type DevelopmentModeUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[DevelopmentModeValue] `json:"value,required"`
}

// MarshalJSON serializes DevelopmentModeUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DevelopmentModeUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DevelopmentModeUpdateParams) String() (result string) {
	return fmt.Sprintf("&DevelopmentModeUpdateParams{Value:%s}", r.Value)
}
