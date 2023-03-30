package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsH2PrioritizationService struct {
	Options []options.RequestOption
}

func NewZonesSettingsH2PrioritizationService(opts ...options.RequestOption) (r *ZonesSettingsH2PrioritizationService) {
	r = &ZonesSettingsH2PrioritizationService{}
	r.Options = opts
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZonesSettingsH2PrioritizationService) Update(ctx context.Context, zone_identifier string, body *requests.H2PrioritizationUpdateParams, opts ...options.RequestOption) (res *responses.H2PrioritizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets HTTP/2 Edge Prioritization setting.
func (r *ZonesSettingsH2PrioritizationService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.H2PrioritizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/h2_prioritization", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
