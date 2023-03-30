package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsNelService struct {
	Options []options.RequestOption
}

func NewZonesSettingsNelService(opts ...options.RequestOption) (r *ZonesSettingsNelService) {
	r = &ZonesSettingsNelService{}
	r.Options = opts
	return
}

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to our [blog post](http://blog.cloudflare.com/nel-solving-mobile-speed)
// for more information.
func (r *ZonesSettingsNelService) Update(ctx context.Context, zone_identifier string, body *requests.NelUpdateParams, opts ...options.RequestOption) (res *responses.NelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enable Network Error Logging reporting on your zone. (Beta)
func (r *ZonesSettingsNelService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.NelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/nel", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
