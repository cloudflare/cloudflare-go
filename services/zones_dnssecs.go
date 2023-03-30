package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesDnssecService struct {
	Options []options.RequestOption
}

func NewZonesDnssecService(opts ...options.RequestOption) (r *ZonesDnssecService) {
	r = &ZonesDnssecService{}
	r.Options = opts
	return
}

// Details about DNSSEC status and configuration.
func (r *ZonesDnssecService) Get(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.DnssecSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Enable or disable DNSSEC.
func (r *ZonesDnssecService) Update(ctx context.Context, zone_identifier string, body *requests.DnssecUpdateParams, opts ...options.RequestOption) (res *responses.DnssecSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Delete DNSSEC.
func (r *ZonesDnssecService) Delete(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.DeleteDnssecResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dnssec", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "DELETE", path, nil, &res, opts...)
	return
}
