package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsSecurityHeaderService struct {
	Options []options.RequestOption
}

func NewZonesSettingsSecurityHeaderService(opts ...options.RequestOption) (r *ZonesSettingsSecurityHeaderService) {
	r = &ZonesSettingsSecurityHeaderService{}
	r.Options = opts
	return
}

// Cloudflare security header for a zone.
func (r *ZonesSettingsSecurityHeaderService) Update(ctx context.Context, zone_identifier string, body *requests.SecurityHeaderUpdateParams, opts ...options.RequestOption) (res *responses.SecurityHeaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_header", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Cloudflare security header for a zone.
func (r *ZonesSettingsSecurityHeaderService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.SecurityHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_header", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
