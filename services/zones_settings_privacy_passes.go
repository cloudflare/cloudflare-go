package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsPrivacyPassService struct {
	Options []options.RequestOption
}

func NewZonesSettingsPrivacyPassService(opts ...options.RequestOption) (r *ZonesSettingsPrivacyPassService) {
	r = &ZonesSettingsPrivacyPassService{}
	r.Options = opts
	return
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
func (r *ZonesSettingsPrivacyPassService) Update(ctx context.Context, zone_identifier string, body *requests.PrivacyPassUpdateParams, opts ...options.RequestOption) (res *responses.PrivacyPassUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/privacy_pass", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
func (r *ZonesSettingsPrivacyPassService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.PrivacyPassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/privacy_pass", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
