// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RegionalTieredCacheService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRegionalTieredCacheService]
// method instead.
type RegionalTieredCacheService struct {
	Options []option.RequestOption
}

// NewRegionalTieredCacheService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalTieredCacheService(opts ...option.RequestOption) (r *RegionalTieredCacheService) {
	r = &RegionalTieredCacheService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *RegionalTieredCacheService) Edit(ctx context.Context, params RegionalTieredCacheEditParams, opts ...option.RequestOption) (res *RegionalTieredCacheEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RegionalTieredCacheEditResponseEnvelope
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
func (r *RegionalTieredCacheService) Get(ctx context.Context, query RegionalTieredCacheGetParams, opts ...option.RequestOption) (res *RegionalTieredCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RegionalTieredCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", query.ZoneID)
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
type RegionalTieredCacheEditResponse struct {
	// ID of the zone setting.
	ID RegionalTieredCacheEditResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value RegionalTieredCacheEditResponseValue `json:"value,required"`
	JSON  regionalTieredCacheEditResponseJSON  `json:"-"`
}

// regionalTieredCacheEditResponseJSON contains the JSON metadata for the struct
// [RegionalTieredCacheEditResponse]
type regionalTieredCacheEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type RegionalTieredCacheEditResponseID string

const (
	RegionalTieredCacheEditResponseIDTcRegional RegionalTieredCacheEditResponseID = "tc_regional"
)

func (r RegionalTieredCacheEditResponseID) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditResponseIDTcRegional:
		return true
	}
	return false
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type RegionalTieredCacheEditResponseValue struct {
	// ID of the zone setting.
	ID RegionalTieredCacheEditResponseValueID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                `json:"modified_on,required,nullable" format:"date-time"`
	JSON       regionalTieredCacheEditResponseValueJSON `json:"-"`
}

// regionalTieredCacheEditResponseValueJSON contains the JSON metadata for the
// struct [RegionalTieredCacheEditResponseValue]
type regionalTieredCacheEditResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type RegionalTieredCacheEditResponseValueID string

const (
	RegionalTieredCacheEditResponseValueIDTcRegional RegionalTieredCacheEditResponseValueID = "tc_regional"
)

func (r RegionalTieredCacheEditResponseValueID) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditResponseValueIDTcRegional:
		return true
	}
	return false
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type RegionalTieredCacheGetResponse struct {
	// ID of the zone setting.
	ID RegionalTieredCacheGetResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value RegionalTieredCacheGetResponseValue `json:"value,required"`
	JSON  regionalTieredCacheGetResponseJSON  `json:"-"`
}

// regionalTieredCacheGetResponseJSON contains the JSON metadata for the struct
// [RegionalTieredCacheGetResponse]
type regionalTieredCacheGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type RegionalTieredCacheGetResponseID string

const (
	RegionalTieredCacheGetResponseIDTcRegional RegionalTieredCacheGetResponseID = "tc_regional"
)

func (r RegionalTieredCacheGetResponseID) IsKnown() bool {
	switch r {
	case RegionalTieredCacheGetResponseIDTcRegional:
		return true
	}
	return false
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type RegionalTieredCacheGetResponseValue struct {
	// ID of the zone setting.
	ID RegionalTieredCacheGetResponseValueID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                               `json:"modified_on,required,nullable" format:"date-time"`
	JSON       regionalTieredCacheGetResponseValueJSON `json:"-"`
}

// regionalTieredCacheGetResponseValueJSON contains the JSON metadata for the
// struct [RegionalTieredCacheGetResponseValue]
type regionalTieredCacheGetResponseValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseValueJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type RegionalTieredCacheGetResponseValueID string

const (
	RegionalTieredCacheGetResponseValueIDTcRegional RegionalTieredCacheGetResponseValueID = "tc_regional"
)

func (r RegionalTieredCacheGetResponseValueID) IsKnown() bool {
	switch r {
	case RegionalTieredCacheGetResponseValueIDTcRegional:
		return true
	}
	return false
}

type RegionalTieredCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[RegionalTieredCacheEditParamsValue] `json:"value,required"`
}

func (r RegionalTieredCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type RegionalTieredCacheEditParamsValue string

const (
	RegionalTieredCacheEditParamsValueOn  RegionalTieredCacheEditParamsValue = "on"
	RegionalTieredCacheEditParamsValueOff RegionalTieredCacheEditParamsValue = "off"
)

func (r RegionalTieredCacheEditParamsValue) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditParamsValueOn, RegionalTieredCacheEditParamsValueOff:
		return true
	}
	return false
}

type RegionalTieredCacheEditResponseEnvelope struct {
	Errors   []RegionalTieredCacheEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegionalTieredCacheEditResponseEnvelopeMessages `json:"messages,required"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Result RegionalTieredCacheEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success RegionalTieredCacheEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    regionalTieredCacheEditResponseEnvelopeJSON    `json:"-"`
}

// regionalTieredCacheEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalTieredCacheEditResponseEnvelope]
type regionalTieredCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalTieredCacheEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    regionalTieredCacheEditResponseEnvelopeErrorsJSON `json:"-"`
}

// regionalTieredCacheEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegionalTieredCacheEditResponseEnvelopeErrors]
type regionalTieredCacheEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalTieredCacheEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    regionalTieredCacheEditResponseEnvelopeMessagesJSON `json:"-"`
}

// regionalTieredCacheEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RegionalTieredCacheEditResponseEnvelopeMessages]
type regionalTieredCacheEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalTieredCacheEditResponseEnvelopeSuccess bool

const (
	RegionalTieredCacheEditResponseEnvelopeSuccessTrue RegionalTieredCacheEditResponseEnvelopeSuccess = true
)

func (r RegionalTieredCacheEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalTieredCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalTieredCacheGetResponseEnvelope struct {
	Errors   []RegionalTieredCacheGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegionalTieredCacheGetResponseEnvelopeMessages `json:"messages,required"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Result RegionalTieredCacheGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success RegionalTieredCacheGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    regionalTieredCacheGetResponseEnvelopeJSON    `json:"-"`
}

// regionalTieredCacheGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalTieredCacheGetResponseEnvelope]
type regionalTieredCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalTieredCacheGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    regionalTieredCacheGetResponseEnvelopeErrorsJSON `json:"-"`
}

// regionalTieredCacheGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegionalTieredCacheGetResponseEnvelopeErrors]
type regionalTieredCacheGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalTieredCacheGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    regionalTieredCacheGetResponseEnvelopeMessagesJSON `json:"-"`
}

// regionalTieredCacheGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RegionalTieredCacheGetResponseEnvelopeMessages]
type regionalTieredCacheGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalTieredCacheGetResponseEnvelopeSuccess bool

const (
	RegionalTieredCacheGetResponseEnvelopeSuccessTrue RegionalTieredCacheGetResponseEnvelopeSuccess = true
)

func (r RegionalTieredCacheGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalTieredCacheGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
