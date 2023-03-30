package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ResponseBufferingListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result fields.Field[ResponseBuffering] `json:"result"`
}

// MarshalJSON serializes ResponseBufferingListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ResponseBufferingListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ResponseBufferingListResponse) String() (result string) {
	return fmt.Sprintf("&ResponseBufferingListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ResponseBufferingUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// Enables or disables buffering of responses from the proxied server. Cloudflare
	// may buffer the whole payload to deliver it at once to the client versus allowing
	// it to be delivered in chunks. By default, the proxied server streams directly
	// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
	Result fields.Field[ResponseBuffering] `json:"result"`
}

// MarshalJSON serializes ResponseBufferingUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ResponseBufferingUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ResponseBufferingUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ResponseBufferingUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ResponseBufferingUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[ResponseBufferingValue] `json:"value,required"`
}

// MarshalJSON serializes ResponseBufferingUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ResponseBufferingUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ResponseBufferingUpdateParams) String() (result string) {
	return fmt.Sprintf("&ResponseBufferingUpdateParams{Value:%s}", r.Value)
}
