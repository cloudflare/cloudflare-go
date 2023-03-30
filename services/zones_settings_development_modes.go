package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsDevelopmentModeService struct {
	Options []options.RequestOption
}

func NewZonesSettingsDevelopmentModeService(opts ...options.RequestOption) (r *ZonesSettingsDevelopmentModeService) {
	r = &ZonesSettingsDevelopmentModeService{}
	r.Options = opts
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZonesSettingsDevelopmentModeService) Update(ctx context.Context, zone_identifier string, body *requests.DevelopmentModeUpdateParams, opts ...options.RequestOption) (res *responses.DevelopmentModeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/development_mode", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
func (r *ZonesSettingsDevelopmentModeService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.DevelopmentModeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/development_mode", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
