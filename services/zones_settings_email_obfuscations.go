package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsEmailObfuscationService struct {
	Options []options.RequestOption
}

func NewZonesSettingsEmailObfuscationService(opts ...options.RequestOption) (r *ZonesSettingsEmailObfuscationService) {
	r = &ZonesSettingsEmailObfuscationService{}
	r.Options = opts
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZonesSettingsEmailObfuscationService) Update(ctx context.Context, zone_identifier string, body *requests.EmailObfuscationUpdateParams, opts ...options.RequestOption) (res *responses.EmailObfuscationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
func (r *ZonesSettingsEmailObfuscationService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.EmailObfuscationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/email_obfuscation", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
