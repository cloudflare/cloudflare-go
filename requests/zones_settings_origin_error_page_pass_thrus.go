package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type OriginErrorPagePassThrusListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result fields.Field[OriginErrorPagePassThru] `json:"result"`
}

// MarshalJSON serializes OriginErrorPagePassThrusListResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OriginErrorPagePassThrusListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginErrorPagePassThrusListResponse) String() (result string) {
	return fmt.Sprintf("&OriginErrorPagePassThrusListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OriginErrorPagePassThrusUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will proxy customer error pages on any 502,504 errors on origin
	// server instead of showing a default Cloudflare error page. This does not apply
	// to 522 errors and is limited to Enterprise Zones.
	Result fields.Field[OriginErrorPagePassThru] `json:"result"`
}

// MarshalJSON serializes OriginErrorPagePassThrusUpdateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OriginErrorPagePassThrusUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginErrorPagePassThrusUpdateResponse) String() (result string) {
	return fmt.Sprintf("&OriginErrorPagePassThrusUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OriginErrorPagePassThrusUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[OriginErrorPagePassThruValue] `json:"value,required"`
}

// MarshalJSON serializes OriginErrorPagePassThrusUpdateParams into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *OriginErrorPagePassThrusUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginErrorPagePassThrusUpdateParams) String() (result string) {
	return fmt.Sprintf("&OriginErrorPagePassThrusUpdateParams{Value:%s}", r.Value)
}
