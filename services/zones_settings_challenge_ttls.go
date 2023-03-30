package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsChallengeTtlService struct {
	Options []options.RequestOption
}

func NewZonesSettingsChallengeTtlService(opts ...options.RequestOption) (r *ZonesSettingsChallengeTtlService) {
	r = &ZonesSettingsChallengeTtlService{}
	r.Options = opts
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZonesSettingsChallengeTtlService) Update(ctx context.Context, zone_identifier string, body *requests.ChallengeTtlUpdateParams, opts ...options.RequestOption) (res *responses.ChallengeTtlUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
func (r *ZonesSettingsChallengeTtlService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ChallengeTtlListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/challenge_ttl", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
