package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsBrowserCacheTtlService struct {
	Options []options.RequestOption
}

func NewZonesSettingsBrowserCacheTtlService(opts ...options.RequestOption) (r *ZonesSettingsBrowserCacheTtlService) {
	r = &ZonesSettingsBrowserCacheTtlService{}
	r.Options = opts
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZonesSettingsBrowserCacheTtlService) Update(ctx context.Context, zone_identifier string, body *requests.BrowserCacheTtlUpdateParams, opts ...options.RequestOption) (res *responses.BrowserCacheTtlUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
func (r *ZonesSettingsBrowserCacheTtlService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.BrowserCacheTtlListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_cache_ttl", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
