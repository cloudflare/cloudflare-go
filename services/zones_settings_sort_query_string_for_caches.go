package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingsSortQueryStringForCachService struct {
	Options []options.RequestOption
}

func NewZonesSettingsSortQueryStringForCachService(opts ...options.RequestOption) (r *ZonesSettingsSortQueryStringForCachService) {
	r = &ZonesSettingsSortQueryStringForCachService{}
	r.Options = opts
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *ZonesSettingsSortQueryStringForCachService) Update(ctx context.Context, zone_identifier string, body *requests.SortQueryStringForCachUpdateParams, opts ...options.RequestOption) (res *responses.SortQueryStringForCachUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
func (r *ZonesSettingsSortQueryStringForCachService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.SortQueryStringForCachListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings/sort_query_string_for_cache", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}
