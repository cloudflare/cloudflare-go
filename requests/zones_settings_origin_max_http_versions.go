package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type OriginMaxHTTPVersionListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                 `json:"success"`
	Result  fields.Field[OriginMaxHTTPVersion] `json:"result"`
}

// MarshalJSON serializes OriginMaxHTTPVersionListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OriginMaxHTTPVersionListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginMaxHTTPVersionListResponse) String() (result string) {
	return fmt.Sprintf("&OriginMaxHTTPVersionListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OriginMaxHTTPVersionUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool]                 `json:"success"`
	Result  fields.Field[OriginMaxHTTPVersion] `json:"result"`
}

// MarshalJSON serializes OriginMaxHTTPVersionUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OriginMaxHTTPVersionUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginMaxHTTPVersionUpdateResponse) String() (result string) {
	return fmt.Sprintf("&OriginMaxHTTPVersionUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type OriginMaxHTTPVersionUpdateParams struct {
	Value fields.Field[OriginMaxHTTPVersion] `json:"value,required"`
}

// MarshalJSON serializes OriginMaxHTTPVersionUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *OriginMaxHTTPVersionUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r OriginMaxHTTPVersionUpdateParams) String() (result string) {
	return fmt.Sprintf("&OriginMaxHTTPVersionUpdateParams{Value:%s}", r.Value)
}
