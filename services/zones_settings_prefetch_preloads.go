package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsPrefetchPreloadService struct {
	Options []options.RequestOption
}

func NewZonesSettingsPrefetchPreloadService(opts ...options.RequestOption) (r *ZonesSettingsPrefetchPreloadService) {
	r = &ZonesSettingsPrefetchPreloadService{}
	r.Options = opts
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZonesSettingsPrefetchPreloadService) Update(ctx context.Context, zone_identifier string, body *requests.PrefetchPreloadUpdateParams, opts ...options.RequestOption) (res *responses.PrefetchPreloadUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
func (r *ZonesSettingsPrefetchPreloadService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.PrefetchPreloadListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/prefetch_preload", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
