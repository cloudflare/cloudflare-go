package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type Http2ListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// HTTP2 enabled for this zone.
	Result fields.Field[Http2] `json:"result"`
}

// MarshalJSON serializes Http2ListResponse into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Http2ListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Http2ListResponse) String() (result string) {
	return fmt.Sprintf("&Http2ListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Http2UpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// HTTP2 enabled for this zone.
	Result fields.Field[Http2] `json:"result"`
}

// MarshalJSON serializes Http2UpdateResponse into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *Http2UpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Http2UpdateResponse) String() (result string) {
	return fmt.Sprintf("&Http2UpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type Http2UpdateParams struct {
	// Value of the HTTP2 setting.
	Value fields.Field[Http2Value] `json:"value,required"`
}

// MarshalJSON serializes Http2UpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *Http2UpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Http2UpdateParams) String() (result string) {
	return fmt.Sprintf("&Http2UpdateParams{Value:%s}", r.Value)
}
