package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type AlwaysOnlineListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result fields.Field[AlwaysOnline] `json:"result"`
}

// MarshalJSON serializes AlwaysOnlineListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysOnlineListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysOnlineListResponse) String() (result string) {
	return fmt.Sprintf("&AlwaysOnlineListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AlwaysOnlineUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// When enabled, Cloudflare serves limited copies of web pages available from the
	// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
	// offline. Refer to
	// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
	// more information.
	Result fields.Field[AlwaysOnline] `json:"result"`
}

// MarshalJSON serializes AlwaysOnlineUpdateResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysOnlineUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysOnlineUpdateResponse) String() (result string) {
	return fmt.Sprintf("&AlwaysOnlineUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type AlwaysOnlineUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[AlwaysOnlineValue] `json:"value,required"`
}

// MarshalJSON serializes AlwaysOnlineUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AlwaysOnlineUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r AlwaysOnlineUpdateParams) String() (result string) {
	return fmt.Sprintf("&AlwaysOnlineUpdateParams{Value:%s}", r.Value)
}
