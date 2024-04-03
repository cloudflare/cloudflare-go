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

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheReserveClearResponse struct {
	// ID of the zone setting.
	ID CacheReserveClearResponseID `json:"id,required"`
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

// ID of the zone setting.
type CacheReserveClearResponseID string

const (
	CacheReserveClearResponseIDCacheReserveClear CacheReserveClearResponseID = "cache_reserve_clear"
)

func (r CacheReserveClearResponseID) IsKnown() bool {
	switch r {
	case CacheReserveClearResponseIDCacheReserveClear:
		return true
	}
	return false
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
	ID CacheReserveEditResponseID `json:"id,required"`
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

// ID of the zone setting.
type CacheReserveEditResponseID string

const (
	CacheReserveEditResponseIDCacheReserve CacheReserveEditResponseID = "cache_reserve"
)

func (r CacheReserveEditResponseID) IsKnown() bool {
	switch r {
	case CacheReserveEditResponseIDCacheReserve:
		return true
	}
	return false
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
	ID CacheReserveGetResponseID `json:"id,required"`
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

// ID of the zone setting.
type CacheReserveGetResponseID string

const (
	CacheReserveGetResponseIDCacheReserve CacheReserveGetResponseID = "cache_reserve"
)

func (r CacheReserveGetResponseID) IsKnown() bool {
	switch r {
	case CacheReserveGetResponseIDCacheReserve:
		return true
	}
	return false
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
	ID CacheReserveStatusResponseID `json:"id,required"`
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

// ID of the zone setting.
type CacheReserveStatusResponseID string

const (
	CacheReserveStatusResponseIDCacheReserveClear CacheReserveStatusResponseID = "cache_reserve_clear"
)

func (r CacheReserveStatusResponseID) IsKnown() bool {
	switch r {
	case CacheReserveStatusResponseIDCacheReserveClear:
		return true
	}
	return false
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
	Errors   []CacheReserveClearResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveClearResponseEnvelopeMessages `json:"messages,required"`
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

type CacheReserveClearResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cacheReserveClearResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheReserveClearResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheReserveClearResponseEnvelopeErrors]
type cacheReserveClearResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveClearResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheReserveClearResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    cacheReserveClearResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheReserveClearResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheReserveClearResponseEnvelopeMessages]
type cacheReserveClearResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveClearResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []CacheReserveEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveEditResponseEnvelopeMessages `json:"messages,required"`
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

type CacheReserveEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    cacheReserveEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheReserveEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheReserveEditResponseEnvelopeErrors]
type cacheReserveEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheReserveEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheReserveEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheReserveEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheReserveEditResponseEnvelopeMessages]
type cacheReserveEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveEditResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []CacheReserveGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveGetResponseEnvelopeMessages `json:"messages,required"`
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

type CacheReserveGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    cacheReserveGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheReserveGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheReserveGetResponseEnvelopeErrors]
type cacheReserveGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheReserveGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cacheReserveGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheReserveGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheReserveGetResponseEnvelopeMessages]
type cacheReserveGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []CacheReserveStatusResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveStatusResponseEnvelopeMessages `json:"messages,required"`
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

type CacheReserveStatusResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheReserveStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheReserveStatusResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheReserveStatusResponseEnvelopeErrors]
type cacheReserveStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveStatusResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheReserveStatusResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    cacheReserveStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheReserveStatusResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CacheReserveStatusResponseEnvelopeMessages]
type cacheReserveStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheReserveStatusResponseEnvelopeMessagesJSON) RawJSON() string {
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
