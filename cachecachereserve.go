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
func (r *CacheCacheReserveService) Update(ctx context.Context, zoneID string, body CacheCacheReserveUpdateParams, opts ...option.RequestOption) (res *CacheCacheReserveUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveUpdateResponseEnvelope
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
type CacheCacheReserveUpdateResponse struct {
	// ID of the zone setting.
	ID CacheCacheReserveUpdateResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CacheCacheReserveUpdateResponseValue `json:"value,required"`
	JSON  cacheCacheReserveUpdateResponseJSON  `json:"-"`
}

// cacheCacheReserveUpdateResponseJSON contains the JSON metadata for the struct
// [CacheCacheReserveUpdateResponse]
type cacheCacheReserveUpdateResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheCacheReserveUpdateResponseID string

const (
	CacheCacheReserveUpdateResponseIDCacheReserve CacheCacheReserveUpdateResponseID = "cache_reserve"
)

// Value of the Cache Reserve zone setting.
type CacheCacheReserveUpdateResponseValue string

const (
	CacheCacheReserveUpdateResponseValueOn  CacheCacheReserveUpdateResponseValue = "on"
	CacheCacheReserveUpdateResponseValueOff CacheCacheReserveUpdateResponseValue = "off"
)

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

type CacheCacheReserveUpdateParams struct {
	// Value of the Cache Reserve zone setting.
	Value param.Field[CacheCacheReserveUpdateParamsValue] `json:"value,required"`
}

func (r CacheCacheReserveUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Cache Reserve zone setting.
type CacheCacheReserveUpdateParamsValue string

const (
	CacheCacheReserveUpdateParamsValueOn  CacheCacheReserveUpdateParamsValue = "on"
	CacheCacheReserveUpdateParamsValueOff CacheCacheReserveUpdateParamsValue = "off"
)

type CacheCacheReserveUpdateResponseEnvelope struct {
	Errors   []CacheCacheReserveUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheCacheReserveUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CacheCacheReserveUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheCacheReserveUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheCacheReserveUpdateResponseEnvelopeJSON    `json:"-"`
}

// cacheCacheReserveUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheCacheReserveUpdateResponseEnvelope]
type cacheCacheReserveUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cacheCacheReserveUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheCacheReserveUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CacheCacheReserveUpdateResponseEnvelopeErrors]
type cacheCacheReserveUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cacheCacheReserveUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheCacheReserveUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CacheCacheReserveUpdateResponseEnvelopeMessages]
type cacheCacheReserveUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheCacheReserveUpdateResponseEnvelopeSuccess bool

const (
	CacheCacheReserveUpdateResponseEnvelopeSuccessTrue CacheCacheReserveUpdateResponseEnvelopeSuccess = true
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
