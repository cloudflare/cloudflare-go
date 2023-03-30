package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsOriginMaxHTTPVersionService struct {
	Options []options.RequestOption
}

func NewZonesSettingsOriginMaxHTTPVersionService(opts ...options.RequestOption) (r *ZonesSettingsOriginMaxHTTPVersionService) {
	r = &ZonesSettingsOriginMaxHTTPVersionService{}
	r.Options = opts
	return
}

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
func (r *ZonesSettingsOriginMaxHTTPVersionService) Update(ctx context.Context, zone_identifier string, body *requests.OriginMaxHTTPVersionUpdateParams, opts ...options.RequestOption) (res *responses.OriginMaxHTTPVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
func (r *ZonesSettingsOriginMaxHTTPVersionService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.OriginMaxHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_max_http_version", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
