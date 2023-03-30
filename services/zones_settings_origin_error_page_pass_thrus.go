package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsOriginErrorPagePassThrusService struct {
	Options []options.RequestOption
}

func NewZonesSettingsOriginErrorPagePassThrusService(opts ...options.RequestOption) (r *ZonesSettingsOriginErrorPagePassThrusService) {
	r = &ZonesSettingsOriginErrorPagePassThrusService{}
	r.Options = opts
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZonesSettingsOriginErrorPagePassThrusService) Update(ctx context.Context, zone_identifier string, body *requests.OriginErrorPagePassThrusUpdateParams, opts ...options.RequestOption) (res *responses.OriginErrorPagePassThrusUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
func (r *ZonesSettingsOriginErrorPagePassThrusService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.OriginErrorPagePassThrusListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/origin_error_page_pass_thru", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
