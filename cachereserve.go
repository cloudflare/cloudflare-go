// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *CacheReserveService) New(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheReserveNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
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
func (r *CacheReserveService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheReserveGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheReserveGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
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
type CacheReserveNewResponse struct {
	// ID of the zone setting.
	ID CacheReserveNewResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveNewResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                   `json:"end_ts" format:"date-time"`
	JSON  cacheReserveNewResponseJSON `json:"-"`
}

// cacheReserveNewResponseJSON contains the JSON metadata for the struct
// [CacheReserveNewResponse]
type cacheReserveNewResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheReserveNewResponseID string

const (
	CacheReserveNewResponseIDCacheReserveClear CacheReserveNewResponseID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheReserveNewResponseState string

const (
	CacheReserveNewResponseStateInProgress CacheReserveNewResponseState = "In-progress"
	CacheReserveNewResponseStateCompleted  CacheReserveNewResponseState = "Completed"
)

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
type CacheReserveGetResponse struct {
	// ID of the zone setting.
	ID CacheReserveGetResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts,required" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveGetResponseState `json:"state,required"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time                   `json:"end_ts" format:"date-time"`
	JSON  cacheReserveGetResponseJSON `json:"-"`
}

// cacheReserveGetResponseJSON contains the JSON metadata for the struct
// [CacheReserveGetResponse]
type cacheReserveGetResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	EndTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheReserveGetResponseID string

const (
	CacheReserveGetResponseIDCacheReserveClear CacheReserveGetResponseID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheReserveGetResponseState string

const (
	CacheReserveGetResponseStateInProgress CacheReserveGetResponseState = "In-progress"
	CacheReserveGetResponseStateCompleted  CacheReserveGetResponseState = "Completed"
)

type CacheReserveNewResponseEnvelope struct {
	Errors   []CacheReserveNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveNewResponseEnvelopeMessages `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
	Result CacheReserveNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheReserveNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheReserveNewResponseEnvelopeJSON    `json:"-"`
}

// cacheReserveNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheReserveNewResponseEnvelope]
type cacheReserveNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    cacheReserveNewResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheReserveNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheReserveNewResponseEnvelopeErrors]
type cacheReserveNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    cacheReserveNewResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheReserveNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheReserveNewResponseEnvelopeMessages]
type cacheReserveNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheReserveNewResponseEnvelopeSuccess bool

const (
	CacheReserveNewResponseEnvelopeSuccessTrue CacheReserveNewResponseEnvelopeSuccess = true
)

type CacheReserveGetResponseEnvelope struct {
	Errors   []CacheReserveGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheReserveGetResponseEnvelopeMessages `json:"messages,required"`
	// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
	// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
	// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
	// that you cannot undo or cancel this operation.
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

// Whether the API call was successful
type CacheReserveGetResponseEnvelopeSuccess bool

const (
	CacheReserveGetResponseEnvelopeSuccessTrue CacheReserveGetResponseEnvelopeSuccess = true
)
