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
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheService) UpdateRegionalTieredCache(ctx context.Context, zoneID string, body CacheUpdateRegionalTieredCacheParams, opts ...option.RequestOption) (res *CacheUpdateRegionalTieredCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type CacheRegionalTieredCachesResponse struct {
	Errors   []CacheRegionalTieredCachesResponseError   `json:"errors"`
	Messages []CacheRegionalTieredCachesResponseMessage `json:"messages"`
	Result   CacheRegionalTieredCachesResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CacheRegionalTieredCachesResponseSuccess `json:"success"`
	JSON    cacheRegionalTieredCachesResponseJSON    `json:"-"`
}

// cacheRegionalTieredCachesResponseJSON contains the JSON metadata for the struct
// [CacheRegionalTieredCachesResponse]
type cacheRegionalTieredCachesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCachesResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    cacheRegionalTieredCachesResponseErrorJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseErrorJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCachesResponseError]
type cacheRegionalTieredCachesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCachesResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheRegionalTieredCachesResponseMessageJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseMessageJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCachesResponseMessage]
type cacheRegionalTieredCachesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCachesResponseResult struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCachesResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheRegionalTieredCachesResponseResultValue `json:"value"`
	JSON  cacheRegionalTieredCachesResponseResultJSON  `json:"-"`
}

// cacheRegionalTieredCachesResponseResultJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCachesResponseResult]
type cacheRegionalTieredCachesResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCachesResponseResultID string

const (
	CacheRegionalTieredCachesResponseResultIDTcRegional CacheRegionalTieredCachesResponseResultID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCachesResponseResultValue struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCachesResponseResultValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                        `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheRegionalTieredCachesResponseResultValueJSON `json:"-"`
}

// cacheRegionalTieredCachesResponseResultValueJSON contains the JSON metadata for
// the struct [CacheRegionalTieredCachesResponseResultValue]
type cacheRegionalTieredCachesResponseResultValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCachesResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCachesResponseResultValueID string

const (
	CacheRegionalTieredCachesResponseResultValueIDTcRegional CacheRegionalTieredCachesResponseResultValueID = "tc_regional"
)

// Whether the API call was successful
type CacheRegionalTieredCachesResponseSuccess bool

const (
	CacheRegionalTieredCachesResponseSuccessTrue CacheRegionalTieredCachesResponseSuccess = true
)

type CacheUpdateRegionalTieredCacheResponse struct {
	Errors   []CacheUpdateRegionalTieredCacheResponseError   `json:"errors"`
	Messages []CacheUpdateRegionalTieredCacheResponseMessage `json:"messages"`
	Result   CacheUpdateRegionalTieredCacheResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CacheUpdateRegionalTieredCacheResponseSuccess `json:"success"`
	JSON    cacheUpdateRegionalTieredCacheResponseJSON    `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseJSON contains the JSON metadata for the
// struct [CacheUpdateRegionalTieredCacheResponse]
type cacheUpdateRegionalTieredCacheResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheUpdateRegionalTieredCacheResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    cacheUpdateRegionalTieredCacheResponseErrorJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseErrorJSON contains the JSON metadata for
// the struct [CacheUpdateRegionalTieredCacheResponseError]
type cacheUpdateRegionalTieredCacheResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheUpdateRegionalTieredCacheResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cacheUpdateRegionalTieredCacheResponseMessageJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseMessageJSON contains the JSON metadata for
// the struct [CacheUpdateRegionalTieredCacheResponseMessage]
type cacheUpdateRegionalTieredCacheResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheUpdateRegionalTieredCacheResponseResult struct {
	// ID of the zone setting.
	ID CacheUpdateRegionalTieredCacheResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheUpdateRegionalTieredCacheResponseResultValue `json:"value"`
	JSON  cacheUpdateRegionalTieredCacheResponseResultJSON  `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseResultJSON contains the JSON metadata for
// the struct [CacheUpdateRegionalTieredCacheResponseResult]
type cacheUpdateRegionalTieredCacheResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheUpdateRegionalTieredCacheResponseResultID string

const (
	CacheUpdateRegionalTieredCacheResponseResultIDTcRegional CacheUpdateRegionalTieredCacheResponseResultID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheUpdateRegionalTieredCacheResponseResultValue struct {
	// ID of the zone setting.
	ID CacheUpdateRegionalTieredCacheResponseResultValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                             `json:"modified_on,nullable" format:"date-time"`
	JSON       cacheUpdateRegionalTieredCacheResponseResultValueJSON `json:"-"`
}

// cacheUpdateRegionalTieredCacheResponseResultValueJSON contains the JSON metadata
// for the struct [CacheUpdateRegionalTieredCacheResponseResultValue]
type cacheUpdateRegionalTieredCacheResponseResultValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheUpdateRegionalTieredCacheResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheUpdateRegionalTieredCacheResponseResultValueID string

const (
	CacheUpdateRegionalTieredCacheResponseResultValueIDTcRegional CacheUpdateRegionalTieredCacheResponseResultValueID = "tc_regional"
)

// Whether the API call was successful
type CacheUpdateRegionalTieredCacheResponseSuccess bool

const (
	CacheUpdateRegionalTieredCacheResponseSuccessTrue CacheUpdateRegionalTieredCacheResponseSuccess = true
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
