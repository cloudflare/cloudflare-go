package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsTls_1_3Service struct {
	Options []options.RequestOption
}

func NewZonesSettingsTls_1_3Service(opts ...options.RequestOption) (r *ZonesSettingsTls_1_3Service) {
	r = &ZonesSettingsTls_1_3Service{}
	r.Options = opts
	return
}

// Changes TLS 1.3 setting.
func (r *ZonesSettingsTls_1_3Service) Update(ctx context.Context, zone_identifier string, body *requests.Tls_1_3UpdateParams, opts ...options.RequestOption) (res *responses.Tls_1_3UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets TLS 1.3 setting enabled for a zone.
func (r *ZonesSettingsTls_1_3Service) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.Tls_1_3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_1_3", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
