package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsHotlinkProtectionService struct {
	Options []options.RequestOption
}

func NewZonesSettingsHotlinkProtectionService(opts ...options.RequestOption) (r *ZonesSettingsHotlinkProtectionService) {
	r = &ZonesSettingsHotlinkProtectionService{}
	r.Options = opts
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZonesSettingsHotlinkProtectionService) Update(ctx context.Context, zone_identifier string, body *requests.HotlinkProtectionUpdateParams, opts ...options.RequestOption) (res *responses.HotlinkProtectionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
func (r *ZonesSettingsHotlinkProtectionService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.HotlinkProtectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/hotlink_protection", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
