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
func (r *CacheRegionalTieredCacheService) Update(ctx context.Context, zoneID string, body CacheRegionalTieredCacheUpdateParams, opts ...option.RequestOption) (res *CacheRegionalTieredCacheUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheRegionalTieredCacheUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *CacheRegionalTieredCacheService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheRegionalTieredCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheRegionalTieredCacheGetResponseEnvelope
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
type CacheRegionalTieredCacheUpdateResponse struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheUpdateResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value CacheRegionalTieredCacheUpdateResponseValue `json:"value,required"`
	JSON  cacheRegionalTieredCacheUpdateResponseJSON  `json:"-"`
}

// cacheRegionalTieredCacheUpdateResponseJSON contains the JSON metadata for the
// struct [CacheRegionalTieredCacheUpdateResponse]
type cacheRegionalTieredCacheUpdateResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCacheUpdateResponseID string

const (
	CacheRegionalTieredCacheUpdateResponseIDTcRegional CacheRegionalTieredCacheUpdateResponseID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type CacheRegionalTieredCacheUpdateResponseValue struct {
	// ID of the zone setting.
	ID CacheRegionalTieredCacheUpdateResponseValueID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cacheRegionalTieredCacheUpdateResponseValueJSON `json:"-"`
}

// cacheRegionalTieredCacheUpdateResponseValueJSON contains the JSON metadata for
// the struct [CacheRegionalTieredCacheUpdateResponseValue]
type cacheRegionalTieredCacheUpdateResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheUpdateResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheRegionalTieredCacheUpdateResponseValueID string

const (
	CacheRegionalTieredCacheUpdateResponseValueIDTcRegional CacheRegionalTieredCacheUpdateResponseValueID = "tc_regional"
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

// ID of the zone setting.
type CacheRegionalTieredCacheGetResponseValueID string

const (
	CacheRegionalTieredCacheGetResponseValueIDTcRegional CacheRegionalTieredCacheGetResponseValueID = "tc_regional"
)

type CacheRegionalTieredCacheUpdateParams struct {
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[CacheRegionalTieredCacheUpdateParamsValue] `json:"value,required"`
}

func (r CacheRegionalTieredCacheUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type CacheRegionalTieredCacheUpdateParamsValue string

const (
	CacheRegionalTieredCacheUpdateParamsValueOn  CacheRegionalTieredCacheUpdateParamsValue = "on"
	CacheRegionalTieredCacheUpdateParamsValueOff CacheRegionalTieredCacheUpdateParamsValue = "off"
)

type CacheRegionalTieredCacheUpdateResponseEnvelope struct {
	Errors   []CacheRegionalTieredCacheUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheRegionalTieredCacheUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Result CacheRegionalTieredCacheUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheRegionalTieredCacheUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheRegionalTieredCacheUpdateResponseEnvelopeJSON    `json:"-"`
}

// cacheRegionalTieredCacheUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [CacheRegionalTieredCacheUpdateResponseEnvelope]
type cacheRegionalTieredCacheUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCacheUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cacheRegionalTieredCacheUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheRegionalTieredCacheUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CacheRegionalTieredCacheUpdateResponseEnvelopeErrors]
type cacheRegionalTieredCacheUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheRegionalTieredCacheUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    cacheRegionalTieredCacheUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheRegionalTieredCacheUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheRegionalTieredCacheUpdateResponseEnvelopeMessages]
type cacheRegionalTieredCacheUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheRegionalTieredCacheUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheRegionalTieredCacheUpdateResponseEnvelopeSuccess bool

const (
	CacheRegionalTieredCacheUpdateResponseEnvelopeSuccessTrue CacheRegionalTieredCacheUpdateResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type CacheRegionalTieredCacheGetResponseEnvelopeSuccess bool

const (
	CacheRegionalTieredCacheGetResponseEnvelopeSuccessTrue CacheRegionalTieredCacheGetResponseEnvelopeSuccess = true
)
