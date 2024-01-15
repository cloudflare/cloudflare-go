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
func (r *ZoneCachTieredCacheSmartTopologyEnableService) Delete(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachTieredCacheSmartTopologyEnableDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get Smart Tiered Cache setting
func (r *ZoneCachTieredCacheSmartTopologyEnableService) SmartTieredCacheGetSmartTieredCacheSetting(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates enablement of Tiered Cache
func (r *ZoneCachTieredCacheSmartTopologyEnableService) SmartTieredCachePatchSmartTieredCacheSetting(ctx context.Context, zoneIdentifier string, body ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams, opts ...option.RequestOption) (res *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneCachTieredCacheSmartTopologyEnableDeleteResponse struct {
	Errors   []ZoneCachTieredCacheSmartTopologyEnableDeleteResponseError   `json:"errors"`
	Messages []ZoneCachTieredCacheSmartTopologyEnableDeleteResponseMessage `json:"messages"`
	Result   interface{}                                                   `json:"result"`
	// Whether the API call was successful
	Success ZoneCachTieredCacheSmartTopologyEnableDeleteResponseSuccess `json:"success"`
	JSON    zoneCachTieredCacheSmartTopologyEnableDeleteResponseJSON    `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableDeleteResponseJSON contains the JSON
// metadata for the struct [ZoneCachTieredCacheSmartTopologyEnableDeleteResponse]
type zoneCachTieredCacheSmartTopologyEnableDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableDeleteResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableDeleteResponseErrorJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableDeleteResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableDeleteResponseError]
type zoneCachTieredCacheSmartTopologyEnableDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableDeleteResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableDeleteResponseMessageJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableDeleteResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableDeleteResponseMessage]
type zoneCachTieredCacheSmartTopologyEnableDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachTieredCacheSmartTopologyEnableDeleteResponseSuccess bool

const (
	ZoneCachTieredCacheSmartTopologyEnableDeleteResponseSuccessTrue ZoneCachTieredCacheSmartTopologyEnableDeleteResponseSuccess = true
)

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse struct {
	Errors   []ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseError   `json:"errors"`
	Messages []ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                                       `json:"result"`
	// Whether the API call was successful
	Success ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseSuccess `json:"success"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseJSON    `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseErrorJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseError]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessageJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessage]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseSuccess bool

const (
	ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseSuccessTrue ZoneCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseSuccess = true
)

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse struct {
	Errors   []ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseError   `json:"errors"`
	Messages []ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessage `json:"messages"`
	Result   interface{}                                                                                         `json:"result"`
	// Whether the API call was successful
	Success ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseSuccess `json:"success"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseJSON    `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseError struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseErrorJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseError]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessage struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessageJSON `json:"-"`
}

// zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessage]
type zoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseSuccess bool

const (
	ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseSuccessTrue ZoneCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseSuccess = true
)

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
