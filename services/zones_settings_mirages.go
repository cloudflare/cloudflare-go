package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsMirageService struct {
	Options []options.RequestOption
}

func NewZonesSettingsMirageService(opts ...options.RequestOption) (r *ZonesSettingsMirageService) {
	r = &ZonesSettingsMirageService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *ZonesSettingsMirageService) Update(ctx context.Context, zone_identifier string, body *requests.MirageUpdateParams, opts ...options.RequestOption) (res *responses.MirageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mirage", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our
// [blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for more
// information.
func (r *ZonesSettingsMirageService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.MirageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/mirage", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
