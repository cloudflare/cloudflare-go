package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsAutomaticHTTPsRewriteService struct {
	Options []options.RequestOption
}

func NewZonesSettingsAutomaticHTTPsRewriteService(opts ...options.RequestOption) (r *ZonesSettingsAutomaticHTTPsRewriteService) {
	r = &ZonesSettingsAutomaticHTTPsRewriteService{}
	r.Options = opts
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZonesSettingsAutomaticHTTPsRewriteService) Update(ctx context.Context, zone_identifier string, body *requests.AutomaticHTTPsRewriteUpdateParams, opts ...options.RequestOption) (res *responses.AutomaticHTTPsRewriteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enable the Automatic HTTPS Rewrites feature for this zone.
func (r *ZonesSettingsAutomaticHTTPsRewriteService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.AutomaticHTTPsRewriteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/automatic_https_rewrites", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
