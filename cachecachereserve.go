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
func (r *CacheCacheReserveService) List(ctx context.Context, query CacheCacheReserveListParams, opts ...option.RequestOption) (res *CacheCacheReserveListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveListResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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
func (r *CacheCacheReserveService) Clear(ctx context.Context, body CacheCacheReserveClearParams, opts ...option.RequestOption) (res *CacheCacheReserveClearResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveClearResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
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
func (r *CacheCacheReserveService) Edit(ctx context.Context, params CacheCacheReserveEditParams, opts ...option.RequestOption) (res *CacheCacheReserveEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
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
func (r *CacheCacheReserveService) Status(ctx context.Context, query CacheCacheReserveStatusParams, opts ...option.RequestOption) (res *CacheCacheReserveStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheCacheReserveStatusResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
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

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheCacheReserveClearResponse struct {
	// ID of the zone setting.
	ID CacheCacheReserveClearResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheCacheReserveClearResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                          `json:"end_ts" format:"date-time"`
	JSON  cacheCacheReserveClearResponseJSON `json:"-"`
}

// cacheCacheReserveClearResponseJSON contains the JSON metadata for the struct
// [CacheCacheReserveClearResponse]
type cacheCacheReserveClearResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheCacheReserveClearResponseID string

const (
	CacheCacheReserveClearResponseIDCacheReserveClear CacheCacheReserveClearResponseID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheCacheReserveClearResponseState string

const (
	CacheCacheReserveClearResponseStateInProgress CacheCacheReserveClearResponseState = "In-progress"
	CacheCacheReserveClearResponseStateCompleted  CacheCacheReserveClearResponseState = "Completed"
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

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheCacheReserveStatusResponse struct {
	// ID of the zone setting.
	ID CacheCacheReserveStatusResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheCacheReserveStatusResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                           `json:"end_ts" format:"date-time"`
	JSON  cacheCacheReserveStatusResponseJSON `json:"-"`
}

// cacheCacheReserveStatusResponseJSON contains the JSON metadata for the struct
// [CacheCacheReserveStatusResponse]
type cacheCacheReserveStatusResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheCacheReserveStatusResponseID string

const (
	CacheCacheReserveStatusResponseIDCacheReserveClear CacheCacheReserveStatusResponseID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheCacheReserveStatusResponseState string

const (
	CacheCacheReserveStatusResponseStateInProgress CacheCacheReserveStatusResponseState = "In-progress"
	CacheCacheReserveStatusResponseStateCompleted  CacheCacheReserveStatusResponseState = "Completed"
)

type CacheCacheReserveListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

type CacheCacheReserveClearParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheCacheReserveClearResponseEnvelope struct {
	Errors   []CacheCacheReserveClearResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheCacheReserveClearResponseEnvelopeMessages `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result CacheCacheReserveClearResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheCacheReserveClearResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheCacheReserveClearResponseEnvelopeJSON    `json:"-"`
}

// cacheCacheReserveClearResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheCacheReserveClearResponseEnvelope]
type cacheCacheReserveClearResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveClearResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveClearResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    cacheCacheReserveClearResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheCacheReserveClearResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CacheCacheReserveClearResponseEnvelopeErrors]
type cacheCacheReserveClearResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveClearResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveClearResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    cacheCacheReserveClearResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheCacheReserveClearResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CacheCacheReserveClearResponseEnvelopeMessages]
type cacheCacheReserveClearResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveClearResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheCacheReserveClearResponseEnvelopeSuccess bool

const (
	CacheCacheReserveClearResponseEnvelopeSuccessTrue CacheCacheReserveClearResponseEnvelopeSuccess = true
)

type CacheCacheReserveEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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

type CacheCacheReserveStatusParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheCacheReserveStatusResponseEnvelope struct {
	Errors   []CacheCacheReserveStatusResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheCacheReserveStatusResponseEnvelopeMessages `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result CacheCacheReserveStatusResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheCacheReserveStatusResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheCacheReserveStatusResponseEnvelopeJSON    `json:"-"`
}

// cacheCacheReserveStatusResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheCacheReserveStatusResponseEnvelope]
type cacheCacheReserveStatusResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveStatusResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveStatusResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    cacheCacheReserveStatusResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheCacheReserveStatusResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CacheCacheReserveStatusResponseEnvelopeErrors]
type cacheCacheReserveStatusResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveStatusResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheCacheReserveStatusResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cacheCacheReserveStatusResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheCacheReserveStatusResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CacheCacheReserveStatusResponseEnvelopeMessages]
type cacheCacheReserveStatusResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheCacheReserveStatusResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheCacheReserveStatusResponseEnvelopeSuccess bool

const (
	CacheCacheReserveStatusResponseEnvelopeSuccessTrue CacheCacheReserveStatusResponseEnvelopeSuccess = true
)
