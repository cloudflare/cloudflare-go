package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsResponseBufferingService struct {
	Options []options.RequestOption
}

func NewZonesSettingsResponseBufferingService(opts ...options.RequestOption) (r *ZonesSettingsResponseBufferingService) {
	r = &ZonesSettingsResponseBufferingService{}
	r.Options = opts
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZonesSettingsResponseBufferingService) Update(ctx context.Context, zone_identifier string, body *requests.ResponseBufferingUpdateParams, opts ...options.RequestOption) (res *responses.ResponseBufferingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
func (r *ZonesSettingsResponseBufferingService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ResponseBufferingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/response_buffering", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
