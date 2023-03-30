package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsTrueClientIpHeaderService struct {
	Options []options.RequestOption
}

func NewZonesSettingsTrueClientIpHeaderService(opts ...options.RequestOption) (r *ZonesSettingsTrueClientIpHeaderService) {
	r = &ZonesSettingsTrueClientIpHeaderService{}
	r.Options = opts
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZonesSettingsTrueClientIpHeaderService) Update(ctx context.Context, zone_identifier string, body *requests.TrueClientIpHeaderUpdateParams, opts ...options.RequestOption) (res *responses.TrueClientIpHeaderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
func (r *ZonesSettingsTrueClientIpHeaderService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.TrueClientIpHeaderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/true_client_ip_header", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
