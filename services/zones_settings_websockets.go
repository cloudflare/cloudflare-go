package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsWebsocketService struct {
	Options []options.RequestOption
}

func NewZonesSettingsWebsocketService(opts ...options.RequestOption) (r *ZonesSettingsWebsocketService) {
	r = &ZonesSettingsWebsocketService{}
	r.Options = opts
	return
}

// Changes Websockets setting. For more information about Websockets, please refer
// to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *ZonesSettingsWebsocketService) Update(ctx context.Context, zone_identifier string, body *requests.WebsocketUpdateParams, opts ...options.RequestOption) (res *responses.WebsocketUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/websockets", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets Websockets setting. For more information about Websockets, please refer to
// [Using Cloudflare with WebSockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Using-Cloudflare-with-WebSockets).
func (r *ZonesSettingsWebsocketService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.WebsocketListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/websockets", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
