package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type PrefetchPreloadListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result fields.Field[PrefetchPreload] `json:"result"`
}

// MarshalJSON serializes PrefetchPreloadListResponse into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PrefetchPreloadListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrefetchPreloadListResponse) String() (result string) {
	return fmt.Sprintf("&PrefetchPreloadListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PrefetchPreloadUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will prefetch any URLs that are included in the response headers.
	// This is limited to Enterprise Zones.
	Result fields.Field[PrefetchPreload] `json:"result"`
}

// MarshalJSON serializes PrefetchPreloadUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *PrefetchPreloadUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrefetchPreloadUpdateResponse) String() (result string) {
	return fmt.Sprintf("&PrefetchPreloadUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type PrefetchPreloadUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[PrefetchPreloadValue] `json:"value,required"`
}

// MarshalJSON serializes PrefetchPreloadUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *PrefetchPreloadUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r PrefetchPreloadUpdateParams) String() (result string) {
	return fmt.Sprintf("&PrefetchPreloadUpdateParams{Value:%s}", r.Value)
}
