package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsProxyReadTimeoutService struct {
	Options []options.RequestOption
}

func NewZonesSettingsProxyReadTimeoutService(opts ...options.RequestOption) (r *ZonesSettingsProxyReadTimeoutService) {
	r = &ZonesSettingsProxyReadTimeoutService{}
	r.Options = opts
	return
}

// Maximum time between two read operations from origin.
func (r *ZonesSettingsProxyReadTimeoutService) Update(ctx context.Context, zone_identifier string, body *requests.ProxyReadTimeoutUpdateParams, opts ...options.RequestOption) (res *responses.ProxyReadTimeoutUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Maximum time between two read operations from origin.
func (r *ZonesSettingsProxyReadTimeoutService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ProxyReadTimeoutListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/proxy_read_timeout", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
