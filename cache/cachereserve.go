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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// CacheReserveService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCacheReserveService] method
// instead.
type CacheReserveService struct {
	Options []option.RequestOption
}

// NewCacheReserveService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCacheReserveService(opts ...option.RequestOption) (r *CacheReserveService) {
	r = &CacheReserveService{}
	r.Options = opts
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *CacheReserveService) Clear(ctx context.Context, params CacheReserveClearParams, opts ...option.RequestOption) (res *CacheReserveClearResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveClearResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
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
func (r *CacheReserveService) Edit(ctx context.Context, params CacheReserveEditParams, opts ...option.RequestOption) (res *CacheReserveEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *CacheReserveService) Get(ctx context.Context, query CacheReserveGetParams, opts ...option.RequestOption) (res *CacheReserveGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *CacheReserveService) Status(ctx context.Context, query CacheReserveStatusParams, opts ...option.RequestOption) (res *CacheReserveStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveStatusResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// ID of the zone setting.
type UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968d string

const (
	UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968dCacheReserveClear UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968d = "cache_reserve_clear"
)

func (r UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968d) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968dCacheReserveClear:
		return true
	}
	return false
}

// ID of the zone setting.
type UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7ad string

const (
	UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7adCacheReserve UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7ad = "cache_reserve"
)

func (r UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7ad) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7adCacheReserve:
		return true
	}
	return false
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheReserveClearResponse struct {
	// ID of the zone setting.
	ID UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968d `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveClearResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                     `json:"end_ts" format:"date-time"`
	JSON  cacheReserveClearResponseJSON `json:"-"`
}

// cacheReserveClearResponseJSON contains the JSON metadata for the struct
// [CacheReserveClearResponse]
type cacheReserveClearResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveClearResponseJSON) RawJSON() string {
	return r.raw
}

// The current state of the Cache Reserve Clear operation.
type CacheReserveClearResponseState string

const (
	CacheReserveClearResponseStateInProgress CacheReserveClearResponseState = "In-progress"
	CacheReserveClearResponseStateCompleted  CacheReserveClearResponseState = "Completed"
)

func (r CacheReserveClearResponseState) IsKnown() bool {
	switch r {
	case CacheReserveClearResponseStateInProgress, CacheReserveClearResponseStateCompleted:
		return true
	}
	return false
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
type CacheReserveEditResponse struct {
	// ID of the zone setting.
	ID UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7ad `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CacheReserveEditResponseValue `json:"value,required"`
	JSON  cacheReserveEditResponseJSON  `json:"-"`
}

// cacheReserveEditResponseJSON contains the JSON metadata for the struct
// [CacheReserveEditResponse]
type cacheReserveEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveEditResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the Cache Reserve zone setting.
type CacheReserveEditResponseValue string

const (
	CacheReserveEditResponseValueOn  CacheReserveEditResponseValue = "on"
	CacheReserveEditResponseValueOff CacheReserveEditResponseValue = "off"
)

func (r CacheReserveEditResponseValue) IsKnown() bool {
	switch r {
	case CacheReserveEditResponseValueOn, CacheReserveEditResponseValueOff:
		return true
	}
	return false
}

// Increase cache lifetimes by automatically storing all cacheable files into
// Cloudflare's persistent object storage buckets. Requires Cache Reserve
// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
// to reduce Reserve operations costs. See the
// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
// for more information.
type CacheReserveGetResponse struct {
	// ID of the zone setting.
	ID UnnamedSchemaRef37c385b4ebac5c7a6475b3f81ef9a7ad `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the Cache Reserve zone setting.
	Value CacheReserveGetResponseValue `json:"value,required"`
	JSON  cacheReserveGetResponseJSON  `json:"-"`
}

// cacheReserveGetResponseJSON contains the JSON metadata for the struct
// [CacheReserveGetResponse]
type cacheReserveGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveGetResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the Cache Reserve zone setting.
type CacheReserveGetResponseValue string

const (
	CacheReserveGetResponseValueOn  CacheReserveGetResponseValue = "on"
	CacheReserveGetResponseValueOff CacheReserveGetResponseValue = "off"
)

func (r CacheReserveGetResponseValue) IsKnown() bool {
	switch r {
	case CacheReserveGetResponseValueOn, CacheReserveGetResponseValueOff:
		return true
	}
	return false
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheReserveStatusResponse struct {
	// ID of the zone setting.
	ID UnnamedSchemaRef2b5e755404a4bfd7892291ce97c4968d `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveStatusResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                      `json:"end_ts" format:"date-time"`
	JSON  cacheReserveStatusResponseJSON `json:"-"`
}

// cacheReserveStatusResponseJSON contains the JSON metadata for the struct
// [CacheReserveStatusResponse]
type cacheReserveStatusResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveStatusResponseJSON) RawJSON() string {
	return r.raw
}

// The current state of the Cache Reserve Clear operation.
type CacheReserveStatusResponseState string

const (
	CacheReserveStatusResponseStateInProgress CacheReserveStatusResponseState = "In-progress"
	CacheReserveStatusResponseStateCompleted  CacheReserveStatusResponseState = "Completed"
)

func (r CacheReserveStatusResponseState) IsKnown() bool {
	switch r {
	case CacheReserveStatusResponseStateInProgress, CacheReserveStatusResponseStateCompleted:
		return true
	}
	return false
}

type CacheReserveClearParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r CacheReserveClearParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CacheReserveClearResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result CacheReserveClearResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheReserveClearResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheReserveClearResponseEnvelopeJSON    `json:"-"`
}

// cacheReserveClearResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheReserveClearResponseEnvelope]
type cacheReserveClearResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveClearResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheReserveClearResponseEnvelopeSuccess bool

const (
	CacheReserveClearResponseEnvelopeSuccessTrue CacheReserveClearResponseEnvelopeSuccess = true
)

func (r CacheReserveClearResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CacheReserveClearResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CacheReserveEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Cache Reserve zone setting.
	Value param.Field[CacheReserveEditParamsValue] `json:"value,required"`
}

func (r CacheReserveEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Cache Reserve zone setting.
type CacheReserveEditParamsValue string

const (
	CacheReserveEditParamsValueOn  CacheReserveEditParamsValue = "on"
	CacheReserveEditParamsValueOff CacheReserveEditParamsValue = "off"
)

func (r CacheReserveEditParamsValue) IsKnown() bool {
	switch r {
	case CacheReserveEditParamsValueOn, CacheReserveEditParamsValueOff:
		return true
	}
	return false
}

type CacheReserveEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CacheReserveEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheReserveEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheReserveEditResponseEnvelopeJSON    `json:"-"`
}

// cacheReserveEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheReserveEditResponseEnvelope]
type cacheReserveEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheReserveEditResponseEnvelopeSuccess bool

const (
	CacheReserveEditResponseEnvelopeSuccessTrue CacheReserveEditResponseEnvelopeSuccess = true
)

func (r CacheReserveEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CacheReserveEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CacheReserveGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheReserveGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// Increase cache lifetimes by automatically storing all cacheable files into
	// Cloudflare's persistent object storage buckets. Requires Cache Reserve
	// subscription. Note: using Tiered Cache with Cache Reserve is highly recommended
	// to reduce Reserve operations costs. See the
	// [developer docs](https://developers.cloudflare.com/cache/about/cache-reserve)
	// for more information.
	Result CacheReserveGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheReserveGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheReserveGetResponseEnvelopeJSON    `json:"-"`
}

// cacheReserveGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheReserveGetResponseEnvelope]
type cacheReserveGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheReserveGetResponseEnvelopeSuccess bool

const (
	CacheReserveGetResponseEnvelopeSuccessTrue CacheReserveGetResponseEnvelopeSuccess = true
)

func (r CacheReserveGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CacheReserveGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CacheReserveStatusParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheReserveStatusResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result CacheReserveStatusResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheReserveStatusResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheReserveStatusResponseEnvelopeJSON    `json:"-"`
}

// cacheReserveStatusResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheReserveStatusResponseEnvelope]
type cacheReserveStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveStatusResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheReserveStatusResponseEnvelopeSuccess bool

const (
	CacheReserveStatusResponseEnvelopeSuccessTrue CacheReserveStatusResponseEnvelopeSuccess = true
)

func (r CacheReserveStatusResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CacheReserveStatusResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
