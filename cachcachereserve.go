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

// CachCacheReserveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCachCacheReserveService] method
// instead.
type CachCacheReserveService struct {
	Options []option.RequestOption
}

// NewCachCacheReserveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCachCacheReserveService(opts ...option.RequestOption) (r *CachCacheReserveService) {
	r = &CachCacheReserveService{}
	r.Options = opts
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *CachCacheReserveService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CachCacheReserveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachCacheReserveListResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *CachCacheReserveService) ZoneCacheSettingsChangeCacheReserveSetting(ctx context.Context, zoneID string, body CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams, opts ...option.RequestOption) (res *CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
type CachCacheReserveListResponse struct {
	// ID of the zone setting.
	ID CachCacheReserveListResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CachCacheReserveListResponseValue `json:"value,required"`
	JSON  cachCacheReserveListResponseJSON  `json:"-"`
}

// cachCacheReserveListResponseJSON contains the JSON metadata for the struct
// [CachCacheReserveListResponse]
type cachCacheReserveListResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CachCacheReserveListResponseID string

const (
	CachCacheReserveListResponseIDCacheReserve CachCacheReserveListResponseID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type CachCacheReserveListResponseValue string

const (
	CachCacheReserveListResponseValueOn  CachCacheReserveListResponseValue = "on"
	CachCacheReserveListResponseValueOff CachCacheReserveListResponseValue = "off"
)

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse struct {
	// ID of the zone setting.
	ID CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValue `json:"value,required"`
	JSON  cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON  `json:"-"`
}

// cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON contains
// the JSON metadata for the struct
// [CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse]
type cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseID string

const (
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseIDCacheReserve CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValue string

const (
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValueOn  CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValue = "on"
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValueOff CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseValue = "off"
)

type CachCacheReserveListResponseEnvelope struct {
	Errors   []CachCacheReserveListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachCacheReserveListResponseEnvelopeMessages `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CachCacheReserveListResponse `json:"result,required"`
	// Whether the API call was successful
	Success CachCacheReserveListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachCacheReserveListResponseEnvelopeJSON    `json:"-"`
}

// cachCacheReserveListResponseEnvelopeJSON contains the JSON metadata for the
// struct [CachCacheReserveListResponseEnvelope]
type cachCacheReserveListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachCacheReserveListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    cachCacheReserveListResponseEnvelopeErrorsJSON `json:"-"`
}

// cachCacheReserveListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CachCacheReserveListResponseEnvelopeErrors]
type cachCacheReserveListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachCacheReserveListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    cachCacheReserveListResponseEnvelopeMessagesJSON `json:"-"`
}

// cachCacheReserveListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CachCacheReserveListResponseEnvelopeMessages]
type cachCacheReserveListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachCacheReserveListResponseEnvelopeSuccess bool

const (
	CachCacheReserveListResponseEnvelopeSuccessTrue CachCacheReserveListResponseEnvelopeSuccess = true
)

type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue] `json:"value,required"`
}

func (r CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Cache Reserve zone setting.
type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue string

const (
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValueOn  CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue = "on"
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValueOff CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingParamsValue = "off"
)

type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelope struct {
	Errors   []CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessages `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponse `json:"result,required"`
	// Whether the API call was successful
	Success CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeJSON    `json:"-"`
}

// cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelope]
type cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrors struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrors]
type cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessages struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessages]
type cachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeSuccess bool

const (
	CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeSuccessTrue CachCacheReserveZoneCacheSettingsChangeCacheReserveSettingResponseEnvelopeSuccess = true
)
