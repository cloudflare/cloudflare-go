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

// CacheCacheReserveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCacheCacheReserveService] method
// instead.
type CacheCacheReserveService struct {
	Options []option.RequestOption
}

// NewCacheCacheReserveService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCacheCacheReserveService(opts ...option.RequestOption) (r *CacheCacheReserveService) {
	r = &CacheCacheReserveService{}
	r.Options = opts
	return
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
func (r *CacheCacheReserveService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheCacheReserveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveListResponseEnvelope
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
func (r *CacheCacheReserveService) Edit(ctx context.Context, zoneID string, body CacheCacheReserveEditParams, opts ...option.RequestOption) (res *CacheCacheReserveEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveEditResponseEnvelope
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
type CacheCacheReserveListResponse struct {
	// ID of the zone setting.
	ID CacheCacheReserveListResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CacheCacheReserveListResponseValue `json:"value,required"`
	JSON  cacheCacheReserveListResponseJSON  `json:"-"`
}

// cacheCacheReserveListResponseJSON contains the JSON metadata for the struct
// [CacheCacheReserveListResponse]
type cacheCacheReserveListResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheCacheReserveListResponseID string

const (
	CacheCacheReserveListResponseIDCacheReserve CacheCacheReserveListResponseID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type CacheCacheReserveListResponseValue string

const (
	CacheCacheReserveListResponseValueOn  CacheCacheReserveListResponseValue = "on"
	CacheCacheReserveListResponseValueOff CacheCacheReserveListResponseValue = "off"
)

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
type CacheCacheReserveEditResponse struct {
	// ID of the zone setting.
	ID CacheCacheReserveEditResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CacheCacheReserveEditResponseValue `json:"value,required"`
	JSON  cacheCacheReserveEditResponseJSON  `json:"-"`
}

// cacheCacheReserveEditResponseJSON contains the JSON metadata for the struct
// [CacheCacheReserveEditResponse]
type cacheCacheReserveEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheCacheReserveEditResponseID string

const (
	CacheCacheReserveEditResponseIDCacheReserve CacheCacheReserveEditResponseID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type CacheCacheReserveEditResponseValue string

const (
	CacheCacheReserveEditResponseValueOn  CacheCacheReserveEditResponseValue = "on"
	CacheCacheReserveEditResponseValueOff CacheCacheReserveEditResponseValue = "off"
)

type CacheCacheReserveListResponseEnvelope struct {
	Errors   []CacheCacheReserveListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheCacheReserveListResponseEnvelopeMessages `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CacheCacheReserveListResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheCacheReserveListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheCacheReserveListResponseEnvelopeJSON    `json:"-"`
}

// cacheCacheReserveListResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheCacheReserveListResponseEnvelope]
type cacheCacheReserveListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    cacheCacheReserveListResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheCacheReserveListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CacheCacheReserveListResponseEnvelopeErrors]
type cacheCacheReserveListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cacheCacheReserveListResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheCacheReserveListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CacheCacheReserveListResponseEnvelopeMessages]
type cacheCacheReserveListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheCacheReserveListResponseEnvelopeSuccess bool

const (
	CacheCacheReserveListResponseEnvelopeSuccessTrue CacheCacheReserveListResponseEnvelopeSuccess = true
)

type CacheCacheReserveEditParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[CacheCacheReserveEditParamsValue] `json:"value,required"`
}

func (r CacheCacheReserveEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Cache Reserve zone setting.
type CacheCacheReserveEditParamsValue string

const (
	CacheCacheReserveEditParamsValueOn  CacheCacheReserveEditParamsValue = "on"
	CacheCacheReserveEditParamsValueOff CacheCacheReserveEditParamsValue = "off"
)

type CacheCacheReserveEditResponseEnvelope struct {
	Errors   []CacheCacheReserveEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheCacheReserveEditResponseEnvelopeMessages `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CacheCacheReserveEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheCacheReserveEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheCacheReserveEditResponseEnvelopeJSON    `json:"-"`
}

// cacheCacheReserveEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheCacheReserveEditResponseEnvelope]
type cacheCacheReserveEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    cacheCacheReserveEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheCacheReserveEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CacheCacheReserveEditResponseEnvelopeErrors]
type cacheCacheReserveEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cacheCacheReserveEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheCacheReserveEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CacheCacheReserveEditResponseEnvelopeMessages]
type cacheCacheReserveEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheCacheReserveEditResponseEnvelopeSuccess bool

const (
	CacheCacheReserveEditResponseEnvelopeSuccessTrue CacheCacheReserveEditResponseEnvelopeSuccess = true
)
