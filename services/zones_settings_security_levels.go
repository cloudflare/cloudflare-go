package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsSecurityLevelService struct {
	Options []options.RequestOption
}

func NewZonesSettingsSecurityLevelService(opts ...options.RequestOption) (r *ZonesSettingsSecurityLevelService) {
	r = &ZonesSettingsSecurityLevelService{}
	r.Options = opts
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *ZonesSettingsSecurityLevelService) Update(ctx context.Context, zone_identifier string, body *requests.SecurityLevelUpdateParams, opts ...options.RequestOption) (res *responses.SecurityLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_level", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
func (r *ZonesSettingsSecurityLevelService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.SecurityLevelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/security_level", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
