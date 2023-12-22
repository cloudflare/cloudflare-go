// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCachTieredCacheSmartTopologyEnableService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCachTieredCacheSmartTopologyEnableService] method instead.
type ZoneCachTieredCacheSmartTopologyEnableService struct {
	Options []option.RequestOption
}

// NewZoneCachTieredCacheSmartTopologyEnableService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCachTieredCacheSmartTopologyEnableService(opts ...option.RequestOption) (r *ZoneCachTieredCacheSmartTopologyEnableService) {
	r = &ZoneCachTieredCacheSmartTopologyEnableService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *ZoneCachTieredCacheSmartTopologyEnableService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Smart Tiered Cache setting
func (r *ZoneCachTieredCacheSmartTopologyEnableService) SmartTieredCacheGetSmartTieredCacheSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Tiered Cache
func (r *ZoneCachTieredCacheSmartTopologyEnableService) SmartTieredCachePatchSmartTieredCacheSetting(ctx context.Context, zoneIdentifier string, body ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams, opts ...option.RequestOption) (res *SchemasResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams struct {
	// Enables Tiered Cache.
	Value param.Field[ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue] `json:"value,required"`
}

func (r ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue string

const (
	ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValueOn  ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue = "on"
	ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValueOff ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue = "off"
)
