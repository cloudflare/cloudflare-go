package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type SortQueryStringForCachListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result fields.Field[SortQueryStringForCache] `json:"result"`
}

// MarshalJSON serializes SortQueryStringForCachListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *SortQueryStringForCachListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SortQueryStringForCachListResponse) String() (result string) {
	return fmt.Sprintf("&SortQueryStringForCachListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SortQueryStringForCachUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cloudflare will treat files with the same query strings as the same file in
	// cache, regardless of the order of the query strings. This is limited to
	// Enterprise Zones.
	Result fields.Field[SortQueryStringForCache] `json:"result"`
}

// MarshalJSON serializes SortQueryStringForCachUpdateResponse into an array of
// bytes using the gjson library. Members of the `jsonFields` field are serialized
// into the top-level, and will overwrite known members of the same name.
func (r *SortQueryStringForCachUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SortQueryStringForCachUpdateResponse) String() (result string) {
	return fmt.Sprintf("&SortQueryStringForCachUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type SortQueryStringForCachUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[SortQueryStringForCacheValue] `json:"value,required"`
}

// MarshalJSON serializes SortQueryStringForCachUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *SortQueryStringForCachUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r SortQueryStringForCachUpdateParams) String() (result string) {
	return fmt.Sprintf("&SortQueryStringForCachUpdateParams{Value:%s}", r.Value)
}
