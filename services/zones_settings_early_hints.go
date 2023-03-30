package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsEarlyHintService struct {
	Options []options.RequestOption
}

func NewZonesSettingsEarlyHintService(opts ...options.RequestOption) (r *ZonesSettingsEarlyHintService) {
	r = &ZonesSettingsEarlyHintService{}
	r.Options = opts
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZonesSettingsEarlyHintService) Update(ctx context.Context, zone_identifier string, body *requests.EarlyHintUpdateParams, opts ...options.RequestOption) (res *responses.EarlyHintUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/early_hints", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
func (r *ZonesSettingsEarlyHintService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.EarlyHintListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/early_hints", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
