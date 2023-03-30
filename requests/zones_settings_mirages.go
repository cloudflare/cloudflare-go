package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type MirageListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result fields.Field[Mirage] `json:"result"`
}

// MarshalJSON serializes MirageListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MirageListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MirageListResponse) String() (result string) {
	return fmt.Sprintf("&MirageListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MirageUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically optimize image loading for website visitors on mobile devices.
	// Refer to
	// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
	// more information.
	Result fields.Field[Mirage] `json:"result"`
}

// MarshalJSON serializes MirageUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MirageUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MirageUpdateResponse) String() (result string) {
	return fmt.Sprintf("&MirageUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MirageUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[MirageValue] `json:"value,required"`
}

// MarshalJSON serializes MirageUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *MirageUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MirageUpdateParams) String() (result string) {
	return fmt.Sprintf("&MirageUpdateParams{Value:%s}", r.Value)
}
