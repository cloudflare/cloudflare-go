package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type BrowserCacheTtlListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result fields.Field[BrowserCacheTtl] `json:"result"`
}

// MarshalJSON serializes BrowserCacheTtlListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrowserCacheTtlListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCacheTtlListResponse) String() (result string) {
	return fmt.Sprintf("&BrowserCacheTtlListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrowserCacheTtlUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
	// will remain on your visitors' computers. Cloudflare will honor any larger times
	// specified by your server.
	// (https://support.cloudflare.com/hc/en-us/articles/200168276).
	Result fields.Field[BrowserCacheTtl] `json:"result"`
}

// MarshalJSON serializes BrowserCacheTtlUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *BrowserCacheTtlUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCacheTtlUpdateResponse) String() (result string) {
	return fmt.Sprintf("&BrowserCacheTtlUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type BrowserCacheTtlUpdateParams struct {
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value fields.Field[BrowserCacheTtlValue] `json:"value,required"`
}

// MarshalJSON serializes BrowserCacheTtlUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *BrowserCacheTtlUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r BrowserCacheTtlUpdateParams) String() (result string) {
	return fmt.Sprintf("&BrowserCacheTtlUpdateParams{Value:%s}", r.Value)
}
