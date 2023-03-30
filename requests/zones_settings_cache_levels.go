package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type CacheLevelListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result fields.Field[CacheLevel] `json:"result"`
}

// MarshalJSON serializes CacheLevelListResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CacheLevelListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CacheLevelListResponse) String() (result string) {
	return fmt.Sprintf("&CacheLevelListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type CacheLevelUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Cache Level functions based off the setting level. The basic setting will cache
	// most static resources (i.e., css, images, and JavaScript). The simplified
	// setting will ignore the query string when delivering a cached resource. The
	// aggressive setting will cache all static resources, including ones with a query
	// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
	Result fields.Field[CacheLevel] `json:"result"`
}

// MarshalJSON serializes CacheLevelUpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CacheLevelUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CacheLevelUpdateResponse) String() (result string) {
	return fmt.Sprintf("&CacheLevelUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type CacheLevelUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[CacheLevelValue] `json:"value,required"`
}

// MarshalJSON serializes CacheLevelUpdateParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CacheLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CacheLevelUpdateParams) String() (result string) {
	return fmt.Sprintf("&CacheLevelUpdateParams{Value:%s}", r.Value)
}
