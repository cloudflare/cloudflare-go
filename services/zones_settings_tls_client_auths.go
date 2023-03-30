package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsTlsClientAuthService struct {
	Options []options.RequestOption
}

func NewZonesSettingsTlsClientAuthService(opts ...options.RequestOption) (r *ZonesSettingsTlsClientAuthService) {
	r = &ZonesSettingsTlsClientAuthService{}
	r.Options = opts
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZonesSettingsTlsClientAuthService) Update(ctx context.Context, zone_identifier string, body *requests.TlsClientAuthUpdateParams, opts ...options.RequestOption) (res *responses.TlsClientAuthUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
func (r *ZonesSettingsTlsClientAuthService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.TlsClientAuthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/tls_client_auth", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
