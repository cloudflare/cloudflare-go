package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsMinTlsVersionService struct {
	Options []options.RequestOption
}

func NewZonesSettingsMinTlsVersionService(opts ...options.RequestOption) (r *ZonesSettingsMinTlsVersionService) {
	r = &ZonesSettingsMinTlsVersionService{}
	r.Options = opts
	return
}

// Changes Minimum TLS Version setting.
func (r *ZonesSettingsMinTlsVersionService) Update(ctx context.Context, zone_identifier string, body *requests.MinTlsVersionUpdateParams, opts ...options.RequestOption) (res *responses.MinTlsVersionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets Minimum TLS Version setting.
func (r *ZonesSettingsMinTlsVersionService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.MinTlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/min_tls_version", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
