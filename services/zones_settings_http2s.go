package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsHttp2Service struct {
	Options []options.RequestOption
}

func NewZonesSettingsHttp2Service(opts ...options.RequestOption) (r *ZonesSettingsHttp2Service) {
	r = &ZonesSettingsHttp2Service{}
	r.Options = opts
	return
}

// Value of the HTTP2 setting.
func (r *ZonesSettingsHttp2Service) Update(ctx context.Context, zone_identifier string, body *requests.Http2UpdateParams, opts ...options.RequestOption) (res *responses.Http2UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Value of the HTTP2 setting.
func (r *ZonesSettingsHttp2Service) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.Http2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/http2", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
