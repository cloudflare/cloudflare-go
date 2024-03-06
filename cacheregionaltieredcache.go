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

// CacheRegionalTieredCacheService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewCacheRegionalTieredCacheService] method instead.
type CacheRegionalTieredCacheService struct {
	Options []option.RequestOption
}

// NewCacheRegionalTieredCacheService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCacheRegionalTieredCacheService(opts ...option.RequestOption) (r *CacheRegionalTieredCacheService) {
	r = &CacheRegionalTieredCacheService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheRegionalTieredCacheService) Edit(ctx context.Context, params CacheRegionalTieredCacheEditParams, opts ...option.RequestOption) (res *CacheRegionalTieredCacheEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheRegionalTieredCacheEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheRegionalTieredCacheService) Get(ctx context.Context, query CacheRegionalTieredCacheGetParams, opts ...option.RequestOption) (res *CacheRegionalTieredCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheRegionalTieredCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCacheEditResponse struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheEditResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheRegionalTieredCacheEditResponseValue `json:"value,required"`
	JSON  cacheRegionalTieredCacheEditResponseJSON  `json:"-"`
}

// cacheRegionalTieredCacheEditResponseJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCacheEditResponse]
type cacheRegionalTieredCacheEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheRegionalTieredCacheEditResponseID string

const (
	CacheRegionalTieredCacheEditResponseIDTcRegional CacheRegionalTieredCacheEditResponseID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCacheEditResponseValue struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheEditResponseValueID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cacheRegionalTieredCacheEditResponseValueJSON `json:"-"`
}

// cacheRegionalTieredCacheEditResponseValueJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCacheEditResponseValue]
type cacheRegionalTieredCacheEditResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheRegionalTieredCacheEditResponseValueID string

const (
	CacheRegionalTieredCacheEditResponseValueIDTcRegional CacheRegionalTieredCacheEditResponseValueID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCacheGetResponse struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheGetResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheRegionalTieredCacheGetResponseValue `json:"value,required"`
	JSON  cacheRegionalTieredCacheGetResponseJSON  `json:"-"`
}

// cacheRegionalTieredCacheGetResponseJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCacheGetResponse]
type cacheRegionalTieredCacheGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheRegionalTieredCacheGetResponseID string

const (
	CacheRegionalTieredCacheGetResponseIDTcRegional CacheRegionalTieredCacheGetResponseID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCacheGetResponseValue struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheGetResponseValueID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                    `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cacheRegionalTieredCacheGetResponseValueJSON `json:"-"`
}

// cacheRegionalTieredCacheGetResponseValueJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCacheGetResponseValue]
type cacheRegionalTieredCacheGetResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheGetResponseValueJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type CacheRegionalTieredCacheGetResponseValueID string

const (
	CacheRegionalTieredCacheGetResponseValueIDTcRegional CacheRegionalTieredCacheGetResponseValueID = "tc_regional"
)

type CacheRegionalTieredCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[CacheRegionalTieredCacheEditParamsValue] `json:"value,required"`
}

func (r CacheRegionalTieredCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type CacheRegionalTieredCacheEditParamsValue string

const (
	CacheRegionalTieredCacheEditParamsValueOn  CacheRegionalTieredCacheEditParamsValue = "on"
	CacheRegionalTieredCacheEditParamsValueOff CacheRegionalTieredCacheEditParamsValue = "off"
)

type CacheRegionalTieredCacheEditResponseEnvelope struct {
	Errors   []CacheRegionalTieredCacheEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheRegionalTieredCacheEditResponseEnvelopeMessages `json:"messages,required"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Result CacheRegionalTieredCacheEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheRegionalTieredCacheEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheRegionalTieredCacheEditResponseEnvelopeJSON    `json:"-"`
}

// cacheRegionalTieredCacheEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [CacheRegionalTieredCacheEditResponseEnvelope]
type cacheRegionalTieredCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheRegionalTieredCacheEditResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    cacheRegionalTieredCacheEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheRegionalTieredCacheEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CacheRegionalTieredCacheEditResponseEnvelopeErrors]
type cacheRegionalTieredCacheEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheRegionalTieredCacheEditResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cacheRegionalTieredCacheEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheRegionalTieredCacheEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheRegionalTieredCacheEditResponseEnvelopeMessages]
type cacheRegionalTieredCacheEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheRegionalTieredCacheEditResponseEnvelopeSuccess bool

const (
	CacheRegionalTieredCacheEditResponseEnvelopeSuccessTrue CacheRegionalTieredCacheEditResponseEnvelopeSuccess = true
)

type CacheRegionalTieredCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheRegionalTieredCacheGetResponseEnvelope struct {
	Errors   []CacheRegionalTieredCacheGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheRegionalTieredCacheGetResponseEnvelopeMessages `json:"messages,required"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Result CacheRegionalTieredCacheGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheRegionalTieredCacheGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheRegionalTieredCacheGetResponseEnvelopeJSON    `json:"-"`
}

// cacheRegionalTieredCacheGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [CacheRegionalTieredCacheGetResponseEnvelope]
type cacheRegionalTieredCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheRegionalTieredCacheGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cacheRegionalTieredCacheGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheRegionalTieredCacheGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CacheRegionalTieredCacheGetResponseEnvelopeErrors]
type cacheRegionalTieredCacheGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheRegionalTieredCacheGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    cacheRegionalTieredCacheGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheRegionalTieredCacheGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheRegionalTieredCacheGetResponseEnvelopeMessages]
type cacheRegionalTieredCacheGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheRegionalTieredCacheGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheRegionalTieredCacheGetResponseEnvelopeSuccess bool

const (
	CacheRegionalTieredCacheGetResponseEnvelopeSuccessTrue CacheRegionalTieredCacheGetResponseEnvelopeSuccess = true
)
