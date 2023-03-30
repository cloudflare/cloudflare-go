package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsAlwaysUseHTTPService struct {
	Options []options.RequestOption
}

func NewZonesSettingsAlwaysUseHTTPService(opts ...options.RequestOption) (r *ZonesSettingsAlwaysUseHTTPService) {
	r = &ZonesSettingsAlwaysUseHTTPService{}
	r.Options = opts
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZonesSettingsAlwaysUseHTTPService) Update(ctx context.Context, zone_identifier string, body *requests.AlwaysUseHTTPUpdateParams, opts ...options.RequestOption) (res *responses.AlwaysUseHTTPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
func (r *ZonesSettingsAlwaysUseHTTPService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.AlwaysUseHTTPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/always_use_https", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
