package responses

import (
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type ServerSideExcludeListResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
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
	Result ServerSideExclude `json:"result"`
	JSON   ServerSideExcludeListResponseJSON
}

type ServerSideExcludeListResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ServerSideExcludeListResponse
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ServerSideExcludeListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type ServerSideExcludeUpdateResponse struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	// Whether the API call was successful
	Success bool `json:"success"`
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
	Result ServerSideExclude `json:"result"`
	JSON   ServerSideExcludeUpdateResponseJSON
}

type ServerSideExcludeUpdateResponseJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Success  pjson.Metadata
	Result   pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// ServerSideExcludeUpdateResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ServerSideExcludeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
