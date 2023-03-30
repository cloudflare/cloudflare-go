package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsOrangeToOrangeService struct {
	Options []options.RequestOption
}

func NewZonesSettingsOrangeToOrangeService(opts ...options.RequestOption) (r *ZonesSettingsOrangeToOrangeService) {
	r = &ZonesSettingsOrangeToOrangeService{}
	r.Options = opts
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZonesSettingsOrangeToOrangeService) Update(ctx context.Context, zone_identifier string, body *requests.OrangeToOrangeUpdateParams, opts ...options.RequestOption) (res *responses.OrangeToOrangeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
func (r *ZonesSettingsOrangeToOrangeService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.OrangeToOrangeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/orange_to_orange", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
