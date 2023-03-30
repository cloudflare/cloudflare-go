package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsOpportunisticOnionService struct {
	Options []options.RequestOption
}

func NewZonesSettingsOpportunisticOnionService(opts ...options.RequestOption) (r *ZonesSettingsOpportunisticOnionService) {
	r = &ZonesSettingsOpportunisticOnionService{}
	r.Options = opts
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *ZonesSettingsOpportunisticOnionService) Update(ctx context.Context, zone_identifier string, body *requests.OpportunisticOnionUpdateParams, opts ...options.RequestOption) (res *responses.OpportunisticOnionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
func (r *ZonesSettingsOpportunisticOnionService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.OpportunisticOnionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/opportunistic_onion", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
