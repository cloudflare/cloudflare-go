package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsCacheLevelService struct {
	Options []options.RequestOption
}

func NewZonesSettingsCacheLevelService(opts ...options.RequestOption) (r *ZonesSettingsCacheLevelService) {
	r = &ZonesSettingsCacheLevelService{}
	r.Options = opts
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *ZonesSettingsCacheLevelService) Update(ctx context.Context, zone_identifier string, body *requests.CacheLevelUpdateParams, opts ...options.RequestOption) (res *responses.CacheLevelUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/cache_level", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
func (r *ZonesSettingsCacheLevelService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.CacheLevelListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/cache_level", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
