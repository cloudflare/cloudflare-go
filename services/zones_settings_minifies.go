package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsMinifyService struct {
	Options []options.RequestOption
}

func NewZonesSettingsMinifyService(opts ...options.RequestOption) (r *ZonesSettingsMinifyService) {
	r = &ZonesSettingsMinifyService{}
	r.Options = opts
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *ZonesSettingsMinifyService) Update(ctx context.Context, zone_identifier string, body *requests.MinifyUpdateParams, opts ...options.RequestOption) (res *responses.MinifyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/minify", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
func (r *ZonesSettingsMinifyService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.MinifyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/minify", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
