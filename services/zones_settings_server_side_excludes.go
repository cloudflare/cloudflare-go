package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsServerSideExcludeService struct {
	Options []options.RequestOption
}

func NewZonesSettingsServerSideExcludeService(opts ...options.RequestOption) (r *ZonesSettingsServerSideExcludeService) {
	r = &ZonesSettingsServerSideExcludeService{}
	r.Options = opts
	return
}

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
func (r *ZonesSettingsServerSideExcludeService) Update(ctx context.Context, zone_identifier string, body *requests.ServerSideExcludeUpdateParams, opts ...options.RequestOption) (res *responses.ServerSideExcludeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

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
func (r *ZonesSettingsServerSideExcludeService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ServerSideExcludeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/server_side_exclude", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
