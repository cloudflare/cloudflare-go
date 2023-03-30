package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type HotlinkProtectionListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result fields.Field[HotlinkProtection] `json:"result"`
}

// MarshalJSON serializes HotlinkProtectionListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *HotlinkProtectionListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r HotlinkProtectionListResponse) String() (result string) {
	return fmt.Sprintf("&HotlinkProtectionListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type HotlinkProtectionUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, the Hotlink Protection option ensures that other sites cannot suck
	// up your bandwidth by building pages that use images hosted on your site. Anytime
	// a request for an image on your site hits Cloudflare, we check to ensure that
	// it's not another site requesting them. People will still be able to download and
	// view images from your page, but other sites won't be able to steal them for use
	// on their own pages.
	// (https://support.cloudflare.com/hc/en-us/articles/200170026).
	Result fields.Field[HotlinkProtection] `json:"result"`
}

// MarshalJSON serializes HotlinkProtectionUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *HotlinkProtectionUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r HotlinkProtectionUpdateResponse) String() (result string) {
	return fmt.Sprintf("&HotlinkProtectionUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type HotlinkProtectionUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[HotlinkProtectionValue] `json:"value,required"`
}

// MarshalJSON serializes HotlinkProtectionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *HotlinkProtectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r HotlinkProtectionUpdateParams) String() (result string) {
	return fmt.Sprintf("&HotlinkProtectionUpdateParams{Value:%s}", r.Value)
}
