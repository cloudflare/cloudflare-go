package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsMobileRedirectService struct {
	Options []options.RequestOption
}

func NewZonesSettingsMobileRedirectService(opts ...options.RequestOption) (r *ZonesSettingsMobileRedirectService) {
	r = &ZonesSettingsMobileRedirectService{}
	r.Options = opts
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZonesSettingsMobileRedirectService) Update(ctx context.Context, zone_identifier string, body *requests.MobileRedirectUpdateParams, opts ...options.RequestOption) (res *responses.MobileRedirectUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
func (r *ZonesSettingsMobileRedirectService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.MobileRedirectListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mobile_redirect", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
