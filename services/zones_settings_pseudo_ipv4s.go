package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsPseudoIpv4Service struct {
	Options []options.RequestOption
}

func NewZonesSettingsPseudoIpv4Service(opts ...options.RequestOption) (r *ZonesSettingsPseudoIpv4Service) {
	r = &ZonesSettingsPseudoIpv4Service{}
	r.Options = opts
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZonesSettingsPseudoIpv4Service) Update(ctx context.Context, zone_identifier string, body *requests.PseudoIpv4UpdateParams, opts ...options.RequestOption) (res *responses.PseudoIpv4UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Value of the Pseudo IPv4 setting.
func (r *ZonesSettingsPseudoIpv4Service) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.PseudoIpv4ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/pseudo_ipv4", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
