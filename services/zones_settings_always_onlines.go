package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsAlwaysOnlineService struct {
	Options []options.RequestOption
}

func NewZonesSettingsAlwaysOnlineService(opts ...options.RequestOption) (r *ZonesSettingsAlwaysOnlineService) {
	r = &ZonesSettingsAlwaysOnlineService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZonesSettingsAlwaysOnlineService) Update(ctx context.Context, zone_identifier string, body *requests.AlwaysOnlineUpdateParams, opts ...options.RequestOption) (res *responses.AlwaysOnlineUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_online", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
func (r *ZonesSettingsAlwaysOnlineService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.AlwaysOnlineListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_online", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
