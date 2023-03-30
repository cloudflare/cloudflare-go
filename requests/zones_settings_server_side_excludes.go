package requests

import (
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type ServerSideExcludeListResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result fields.Field[ServerSideExclude] `json:"result"`
}

// MarshalJSON serializes ServerSideExcludeListResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ServerSideExcludeListResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ServerSideExcludeListResponse) String() (result string) {
	return fmt.Sprintf("&ServerSideExcludeListResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ServerSideExcludeUpdateResponse struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	// Whether the API call was successful
	Success fields.Field[bool] `json:"success"`
	// If there is sensitive content on your website that you want visible to real
	// visitors, but that you want to hide from suspicious visitors, all you have to do
	// is wrap the content with Cloudflare SSE tags. Wrap any content that you want to
	// be excluded from suspicious visitors in the following SSE tags:
	// <!--sse--><!--/sse-->. For example: <!--sse--> Bad visitors won't see my phone
	// number, 555-555-5555 <!--/sse-->. Note: SSE only will work with HTML. If you
	// have HTML minification enabled, you won't see the SSE tags in your HTML source
	// when it's served through Cloudflare. SSE will still function in this case, as
	// Cloudflare's HTML minification and SSE functionality occur on-the-fly as the
	// resource moves through our network to the visitor's computer.
	// (https://support.cloudflare.com/hc/en-us/articles/200170036).
	Result fields.Field[ServerSideExclude] `json:"result"`
}

// MarshalJSON serializes ServerSideExcludeUpdateResponse into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ServerSideExcludeUpdateResponse) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ServerSideExcludeUpdateResponse) String() (result string) {
	return fmt.Sprintf("&ServerSideExcludeUpdateResponse{Errors:%s Messages:%s Success:%s Result:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Success, r.Result)
}

type ServerSideExcludeUpdateParams struct {
	// Value of the zone setting.
	Value fields.Field[ServerSideExcludeValue] `json:"value,required"`
}

// MarshalJSON serializes ServerSideExcludeUpdateParams into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ServerSideExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r ServerSideExcludeUpdateParams) String() (result string) {
	return fmt.Sprintf("&ServerSideExcludeUpdateParams{Value:%s}", r.Value)
}
