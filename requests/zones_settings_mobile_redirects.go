package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type MobileRedirectListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result fields.Field[MobileRedirect] `json:"result"`
}

// MarshalJSON serializes MobileRedirectListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MobileRedirectListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MobileRedirectListResponse) String() (result string) {
	return fmt.Sprintf("&MobileRedirectListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MobileRedirectUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Automatically redirect visitors on mobile devices to a mobile-optimized
	// subdomain. Refer to
	// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
	// for more information.
	Result fields.Field[MobileRedirect] `json:"result"`
}

// MarshalJSON serializes MobileRedirectUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MobileRedirectUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MobileRedirectUpdateResponse) String() (result string) {
	return fmt.Sprintf("&MobileRedirectUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type MobileRedirectUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[MobileRedirectValue] `json:"value,required"`
}

// MarshalJSON serializes MobileRedirectUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *MobileRedirectUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r MobileRedirectUpdateParams) String() (result string) {
	return fmt.Sprintf("&MobileRedirectUpdateParams{Value:%s}", r.Value)
}
