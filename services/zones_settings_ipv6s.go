package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsIpv6Service struct {
	Options []options.RequestOption
}

func NewZonesSettingsIpv6Service(opts ...options.RequestOption) (r *ZonesSettingsIpv6Service) {
	r = &ZonesSettingsIpv6Service{}
	r.Options = opts
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZonesSettingsIpv6Service) Update(ctx context.Context, zone_identifier string, body *requests.Ipv6UpdateParams, opts ...options.RequestOption) (res *responses.Ipv6UpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
func (r *ZonesSettingsIpv6Service) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.Ipv6ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ipv6", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
