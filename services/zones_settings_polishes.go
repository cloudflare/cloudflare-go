package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsPolishService struct {
	Options []options.RequestOption
}

func NewZonesSettingsPolishService(opts ...options.RequestOption) (r *ZonesSettingsPolishService) {
	r = &ZonesSettingsPolishService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *ZonesSettingsPolishService) Update(ctx context.Context, zone_identifier string, body *requests.PolishUpdateParams, opts ...options.RequestOption) (res *responses.PolishUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/polish", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/polish-solving-mobile-speed)
// for more information.
func (r *ZonesSettingsPolishService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.PolishListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/polish", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
