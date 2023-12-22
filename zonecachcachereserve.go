// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneCachCacheReserveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCachCacheReserveService]
// method instead.
type ZoneCachCacheReserveService struct {
	Options []option.RequestOption
}

// NewZoneCachCacheReserveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCachCacheReserveService(opts ...option.RequestOption) (r *ZoneCachCacheReserveService) {
	r = &ZoneCachCacheReserveService{}
	r.Options = opts
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *ZoneCachCacheReserveService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachCacheReserveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *ZoneCachCacheReserveService) ZoneCacheSettingsChangeCacheReserveSetting(ctx context.Context, zoneIdentifier string, body ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams, opts ...option.RequestOption) (res *ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneCachCacheReserveListResponse struct {
	Errors   []ZoneCachCacheReserveListResponseError   `json:"errors"`
	Messages []ZoneCachCacheReserveListResponseMessage `json:"messages"`
	Result   ZoneCachCacheReserveListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachCacheReserveListResponseSuccess `json:"success"`
	JSON    zoneCachCacheReserveListResponseJSON    `json:"-"`
}

// zoneCachCacheReserveListResponseJSON contains the JSON metadata for the struct
// [ZoneCachCacheReserveListResponse]
type zoneCachCacheReserveListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneCachCacheReserveListResponseErrorJSON `json:"-"`
}

// zoneCachCacheReserveListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneCachCacheReserveListResponseError]
type zoneCachCacheReserveListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneCachCacheReserveListResponseMessageJSON `json:"-"`
}

// zoneCachCacheReserveListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCachCacheReserveListResponseMessage]
type zoneCachCacheReserveListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveListResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachCacheReserveListResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value ZoneCachCacheReserveListResponseResultValue `json:"value"`
	JSON  zoneCachCacheReserveListResponseResultJSON  `json:"-"`
}

// zoneCachCacheReserveListResponseResultJSON contains the JSON metadata for the
// struct [ZoneCachCacheReserveListResponseResult]
type zoneCachCacheReserveListResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachCacheReserveListResponseResultID string

const (
	ZoneCachCacheReserveListResponseResultIDCacheReserve ZoneCachCacheReserveListResponseResultID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type ZoneCachCacheReserveListResponseResultValue string

const (
	ZoneCachCacheReserveListResponseResultValueOn  ZoneCachCacheReserveListResponseResultValue = "on"
	ZoneCachCacheReserveListResponseResultValueOff ZoneCachCacheReserveListResponseResultValue = "off"
)

// Whether the API call was successful
type ZoneCachCacheReserveListResponseSuccess bool

const (
	ZoneCachCacheReserveListResponseSuccessTrue ZoneCachCacheReserveListResponseSuccess = true
)

type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse struct {
	Errors   []ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseError   `json:"errors"`
	Messages []ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessage `json:"messages"`
	Result   ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseSuccess `json:"success"`
	JSON    zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON    `json:"-"`
}

// zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON
// contains the JSON metadata for the struct
// [ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse]
type zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseErrorJSON `json:"-"`
}

// zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseError]
type zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessageJSON `json:"-"`
}

// zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessage]
type zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValue `json:"value"`
	JSON  zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultJSON  `json:"-"`
}

// zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResult]
type zoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultID string

const (
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultIDCacheReserve ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValue string

const (
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValueOn  ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValue = "on"
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValueOff ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseResultValue = "off"
)

// Whether the API call was successful
type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseSuccess bool

const (
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseSuccessTrue ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseSuccess = true
)

type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue] `json:"value,required"`
}

func (r ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Cache Reserve zone setting.
type ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue string

const (
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValueOn  ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue = "on"
	ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValueOff ZoneCachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue = "off"
)
