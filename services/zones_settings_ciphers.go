package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsCipherService struct {
	Options []options.RequestOption
}

func NewZonesSettingsCipherService(opts ...options.RequestOption) (r *ZonesSettingsCipherService) {
	r = &ZonesSettingsCipherService{}
	r.Options = opts
	return
}

// Changes ciphers setting.
func (r *ZonesSettingsCipherService) Update(ctx context.Context, zone_identifier string, body *requests.CipherUpdateParams, opts ...options.RequestOption) (res *responses.CipherUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ciphers", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Gets ciphers setting.
func (r *ZonesSettingsCipherService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.CipherListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/ciphers", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
