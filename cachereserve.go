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
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *CacheReserveService) Clear(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheReserveClearResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CacheReserveNewResponse struct {
	Errors   []CacheReserveNewResponseError   `json:"errors"`
	Messages []CacheReserveNewResponseMessage `json:"messages"`
	Result   CacheReserveNewResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CacheReserveNewResponseSuccess `json:"success"`
	JSON    cacheReserveNewResponseJSON    `json:"-"`
}

// cacheReserveNewResponseJSON contains the JSON metadata for the struct
// [CacheReserveNewResponse]
type cacheReserveNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveNewResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    cacheReserveNewResponseErrorJSON `json:"-"`
}

// cacheReserveNewResponseErrorJSON contains the JSON metadata for the struct
// [CacheReserveNewResponseError]
type cacheReserveNewResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveNewResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    cacheReserveNewResponseMessageJSON `json:"-"`
}

// cacheReserveNewResponseMessageJSON contains the JSON metadata for the struct
// [CacheReserveNewResponseMessage]
type cacheReserveNewResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveNewResponseResult struct {
	// ID of the zone setting.
	ID CacheReserveNewResponseResultID `json:"id"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveNewResponseResultState `json:"state"`
	JSON  cacheReserveNewResponseResultJSON  `json:"-"`
}

// cacheReserveNewResponseResultJSON contains the JSON metadata for the struct
// [CacheReserveNewResponseResult]
type cacheReserveNewResponseResultJSON struct {
	ID          apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheReserveNewResponseResultID string

const (
	CacheReserveNewResponseResultIDCacheReserveClear CacheReserveNewResponseResultID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheReserveNewResponseResultState string

const (
	CacheReserveNewResponseResultStateInProgress CacheReserveNewResponseResultState = "In-progress"
	CacheReserveNewResponseResultStateCompleted  CacheReserveNewResponseResultState = "Completed"
)

// Whether the API call was successful
type CacheReserveNewResponseSuccess bool

const (
	CacheReserveNewResponseSuccessTrue CacheReserveNewResponseSuccess = true
)

type CacheReserveClearResponse struct {
	Errors   []CacheReserveClearResponseError   `json:"errors"`
	Messages []CacheReserveClearResponseMessage `json:"messages"`
	Result   CacheReserveClearResponseResult    `json:"result"`
	// Whether the API call was successful
	Success CacheReserveClearResponseSuccess `json:"success"`
	JSON    cacheReserveClearResponseJSON    `json:"-"`
}

// cacheReserveClearResponseJSON contains the JSON metadata for the struct
// [CacheReserveClearResponse]
type cacheReserveClearResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveClearResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    cacheReserveClearResponseErrorJSON `json:"-"`
}

// cacheReserveClearResponseErrorJSON contains the JSON metadata for the struct
// [CacheReserveClearResponseError]
type cacheReserveClearResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveClearResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    cacheReserveClearResponseMessageJSON `json:"-"`
}

// cacheReserveClearResponseMessageJSON contains the JSON metadata for the struct
// [CacheReserveClearResponseMessage]
type cacheReserveClearResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheReserveClearResponseResult struct {
	// ID of the zone setting.
	ID CacheReserveClearResponseResultID `json:"id"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State CacheReserveClearResponseResultState `json:"state"`
	JSON  cacheReserveClearResponseResultJSON  `json:"-"`
}

// cacheReserveClearResponseResultJSON contains the JSON metadata for the struct
// [CacheReserveClearResponseResult]
type cacheReserveClearResponseResultJSON struct {
	ID          apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheReserveClearResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheReserveClearResponseResultID string

const (
	CacheReserveClearResponseResultIDCacheReserveClear CacheReserveClearResponseResultID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type CacheReserveClearResponseResultState string

const (
	CacheReserveClearResponseResultStateInProgress CacheReserveClearResponseResultState = "In-progress"
	CacheReserveClearResponseResultStateCompleted  CacheReserveClearResponseResultState = "Completed"
)

// Whether the API call was successful
type CacheReserveClearResponseSuccess bool

const (
	CacheReserveClearResponseSuccessTrue CacheReserveClearResponseSuccess = true
)
