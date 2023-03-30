package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsHttp3Service struct {
	Options []options.RequestOption
}

func NewZonesSettingsHttp3Service(opts ...options.RequestOption) (r *ZonesSettingsHttp3Service) {
	r = &ZonesSettingsHttp3Service{}
	r.Options = opts
	return
}

// Value of the HTTP3 setting.
func (r *ZonesSettingsHttp3Service) Update(ctx context.Context, zone_identifier string, body *requests.Http3UpdateParams, opts ...options.RequestOption) (res *responses.Http3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Value of the HTTP3 setting.
func (r *ZonesSettingsHttp3Service) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.Http3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http3", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
