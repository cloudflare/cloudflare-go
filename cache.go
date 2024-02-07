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

// CacheService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCacheService] method instead.
type CacheService struct {
	Options []option.RequestOption
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheService) RegionalTieredCaches(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheRegionalTieredCachesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheRegionalTieredCachesResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheService) UpdateRegionalTieredCache(ctx context.Context, zoneID string, body CacheUpdateRegionalTieredCacheParams, opts ...option.RequestOption) (res *CacheUpdateRegionalTieredCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheUpdateRegionalTieredCacheResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CacheRegionalTieredCachesResponse struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCachesResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheRegionalTieredCachesResponseValue `json:"value"`
	JSON  cacheRegionalTieredCachesResponseJSON  `json:"-"`
}

// cacheRegionalTieredCachesResponseJSON contains the JSON metadata for the struct
// [CacheRegionalTieredCachesResponse]
type cacheRegionalTieredCachesResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCachesResponseID string

const (
	CacheRegionalTieredCachesResponseIDTcRegional CacheRegionalTieredCachesResponseID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCachesResponseValue struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCachesResponseValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheRegionalTieredCachesResponseValueJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseValueJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCachesResponseValue]
type cacheRegionalTieredCachesResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCachesResponseValueID string

const (
	CacheRegionalTieredCachesResponseValueIDTcRegional CacheRegionalTieredCachesResponseValueID = "tc_regional"
)

type CacheUpdateRegionalTieredCacheResponse struct {
	// ID of the zone setting.
	ID CacheUpdateRegionalTieredCacheResponseID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheUpdateRegionalTieredCacheResponseValue `json:"value"`
	JSON  cacheUpdateRegionalTieredCacheResponseJSON  `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseJSON contains the JSON metadata for the
// struct [CacheUpdateRegionalTieredCacheResponse]
type cacheUpdateRegionalTieredCacheResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheUpdateRegionalTieredCacheResponseID string

const (
	CacheUpdateRegionalTieredCacheResponseIDTcRegional CacheUpdateRegionalTieredCacheResponseID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheUpdateRegionalTieredCacheResponseValue struct {
	// ID of the zone setting.
	ID CacheUpdateRegionalTieredCacheResponseValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheUpdateRegionalTieredCacheResponseValueJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseValueJSON contains the JSON metadata for
// the struct [CacheUpdateRegionalTieredCacheResponseValue]
type cacheUpdateRegionalTieredCacheResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheUpdateRegionalTieredCacheResponseValueID string

const (
	CacheUpdateRegionalTieredCacheResponseValueIDTcRegional CacheUpdateRegionalTieredCacheResponseValueID = "tc_regional"
)

type CacheRegionalTieredCachesResponseEnvelope struct {
	Errors   []CacheRegionalTieredCachesResponseEnvelopeErrors   `json:"errors"`
	Messages []CacheRegionalTieredCachesResponseEnvelopeMessages `json:"messages"`
	Result   CacheRegionalTieredCachesResponse                   `json:"result"`
	// Whether the API call was successful
	Success CacheRegionalTieredCachesResponseEnvelopeSuccess `json:"success"`
	JSON    cacheRegionalTieredCachesResponseEnvelopeJSON    `json:"-"`
}

// cacheRegionalTieredCachesResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCachesResponseEnvelope]
type cacheRegionalTieredCachesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCachesResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cacheRegionalTieredCachesResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CacheRegionalTieredCachesResponseEnvelopeErrors]
type cacheRegionalTieredCachesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCachesResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cacheRegionalTieredCachesResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CacheRegionalTieredCachesResponseEnvelopeMessages]
type cacheRegionalTieredCachesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheRegionalTieredCachesResponseEnvelopeSuccess bool

const (
	CacheRegionalTieredCachesResponseEnvelopeSuccessTrue CacheRegionalTieredCachesResponseEnvelopeSuccess = true
)

type CacheUpdateRegionalTieredCacheParams struct {
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[CacheUpdateRegionalTieredCacheParamsValue] `json:"value,required"`
}

func (r CacheUpdateRegionalTieredCacheParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type CacheUpdateRegionalTieredCacheParamsValue string

const (
	CacheUpdateRegionalTieredCacheParamsValueOn  CacheUpdateRegionalTieredCacheParamsValue = "on"
	CacheUpdateRegionalTieredCacheParamsValueOff CacheUpdateRegionalTieredCacheParamsValue = "off"
)

type CacheUpdateRegionalTieredCacheResponseEnvelope struct {
	Errors   []CacheUpdateRegionalTieredCacheResponseEnvelopeErrors   `json:"errors"`
	Messages []CacheUpdateRegionalTieredCacheResponseEnvelopeMessages `json:"messages"`
	Result   CacheUpdateRegionalTieredCacheResponse                   `json:"result"`
	// Whether the API call was successful
	Success CacheUpdateRegionalTieredCacheResponseEnvelopeSuccess `json:"success"`
	JSON    cacheUpdateRegionalTieredCacheResponseEnvelopeJSON    `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseEnvelopeJSON contains the JSON metadata
// for the struct [CacheUpdateRegionalTieredCacheResponseEnvelope]
type cacheUpdateRegionalTieredCacheResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheUpdateRegionalTieredCacheResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cacheUpdateRegionalTieredCacheResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CacheUpdateRegionalTieredCacheResponseEnvelopeErrors]
type cacheUpdateRegionalTieredCacheResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheUpdateRegionalTieredCacheResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    cacheUpdateRegionalTieredCacheResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheUpdateRegionalTieredCacheResponseEnvelopeMessages]
type cacheUpdateRegionalTieredCacheResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheUpdateRegionalTieredCacheResponseEnvelopeSuccess bool

const (
	CacheUpdateRegionalTieredCacheResponseEnvelopeSuccessTrue CacheUpdateRegionalTieredCacheResponseEnvelopeSuccess = true
)
