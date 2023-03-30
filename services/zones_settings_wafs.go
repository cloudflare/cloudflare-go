package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsWafService struct {
	Options []options.RequestOption
}

func NewZonesSettingsWafService(opts ...options.RequestOption) (r *ZonesSettingsWafService) {
	r = &ZonesSettingsWafService{}
	r.Options = opts
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *ZonesSettingsWafService) Update(ctx context.Context, zone_identifier string, body *requests.WafUpdateParams, opts ...options.RequestOption) (res *responses.WafUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/waf", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// The WAF examines HTTP requests to your website. It inspects both GET and POST
// requests and applies rules to help filter out illegitimate traffic from
// legitimate website visitors. The Cloudflare WAF inspects website addresses or
// URLs to detect anything out of the ordinary. If the Cloudflare WAF determines
// suspicious user behavior, then the WAF will 'challenge' the web visitor with a
// page that asks them to submit a CAPTCHA successfully to continue their action.
// If the challenge is failed, the action will be stopped. What this means is that
// Cloudflare's WAF will block any traffic identified as illegitimate before it
// reaches your origin web server.
// (https://support.cloudflare.com/hc/en-us/articles/200172016).
func (r *ZonesSettingsWafService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.WafListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/waf", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
