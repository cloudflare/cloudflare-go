package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsBrowserCheckService struct {
	Options []options.RequestOption
}

func NewZonesSettingsBrowserCheckService(opts ...options.RequestOption) (r *ZonesSettingsBrowserCheckService) {
	r = &ZonesSettingsBrowserCheckService{}
	r.Options = opts
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *ZonesSettingsBrowserCheckService) Update(ctx context.Context, zone_identifier string, body *requests.BrowserCheckUpdateParams, opts ...options.RequestOption) (res *responses.BrowserCheckUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_check", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
func (r *ZonesSettingsBrowserCheckService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.BrowserCheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/browser_check", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
