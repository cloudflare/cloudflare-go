package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsBrotliService struct {
	Options []options.RequestOption
}

func NewZonesSettingsBrotliService(opts ...options.RequestOption) (r *ZonesSettingsBrotliService) {
	r = &ZonesSettingsBrotliService{}
	r.Options = opts
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZonesSettingsBrotliService) Update(ctx context.Context, zone_identifier string, body *requests.BrotliUpdateParams, opts ...options.RequestOption) (res *responses.BrotliUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/brotli", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
func (r *ZonesSettingsBrotliService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.BrotliListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/brotli", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
