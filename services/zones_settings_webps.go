package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsWebpService struct {
	Options []options.RequestOption
}

func NewZonesSettingsWebpService(opts ...options.RequestOption) (r *ZonesSettingsWebpService) {
	r = &ZonesSettingsWebpService{}
	r.Options = opts
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *ZonesSettingsWebpService) Update(ctx context.Context, zone_identifier string, body *requests.WebpUpdateParams, opts ...options.RequestOption) (res *responses.WebpUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/webp", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
func (r *ZonesSettingsWebpService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.WebpListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/webp", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
