package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsAutomaticPlatformOptimizationService struct {
	Options []options.RequestOption
}

func NewZonesSettingsAutomaticPlatformOptimizationService(opts ...options.RequestOption) (r *ZonesSettingsAutomaticPlatformOptimizationService) {
	r = &ZonesSettingsAutomaticPlatformOptimizationService{}
	r.Options = opts
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *ZonesSettingsAutomaticPlatformOptimizationService) Update(ctx context.Context, zone_identifier string, body *requests.AutomaticPlatformOptimizationUpdateParams, opts ...options.RequestOption) (res *responses.AutomaticPlatformOptimizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
func (r *ZonesSettingsAutomaticPlatformOptimizationService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.AutomaticPlatformOptimizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_platform_optimization", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
