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

// ZoneCacheService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneCacheService] method instead.
type ZoneCacheService struct {
	Options []option.RequestOption
}

// NewZoneCacheService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCacheService(opts ...option.RequestOption) (r *ZoneCacheService) {
	r = &ZoneCacheService{}
	r.Options = opts
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *ZoneCacheService) GetCacheReserveClear(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCacheGetCacheReserveClearResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *ZoneCacheService) GetOriginPostQuantumEncryption(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCacheGetOriginPostQuantumEncryptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *ZoneCacheService) GetRegionalTieredCache(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCacheGetRegionalTieredCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *ZoneCacheService) PatchRegionalTieredCache(ctx context.Context, zoneIdentifier string, body ZoneCachePatchRegionalTieredCacheParams, opts ...option.RequestOption) (res *ZoneCachePatchRegionalTieredCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// You can use Cache Reserve Clear to clear your Cache Reserve, but you must first
// disable Cache Reserve. In most cases, this will be accomplished within 24 hours.
// You cannot re-enable Cache Reserve while this process is ongoing. Keep in mind
// that you cannot undo or cancel this operation.
func (r *ZoneCacheService) PostCacheReserveClear(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneCachePostCacheReserveClearResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/cache_reserve_clear", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *ZoneCacheService) PutOriginPostQuantumEncryption(ctx context.Context, zoneIdentifier string, body ZoneCachePutOriginPostQuantumEncryptionParams, opts ...option.RequestOption) (res *ZoneCachePutOriginPostQuantumEncryptionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCacheGetCacheReserveClearResponse struct {
	Errors   []ZoneCacheGetCacheReserveClearResponseError   `json:"errors"`
	Messages []ZoneCacheGetCacheReserveClearResponseMessage `json:"messages"`
	Result   ZoneCacheGetCacheReserveClearResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCacheGetCacheReserveClearResponseSuccess `json:"success"`
	JSON    zoneCacheGetCacheReserveClearResponseJSON    `json:"-"`
}

// zoneCacheGetCacheReserveClearResponseJSON contains the JSON metadata for the
// struct [ZoneCacheGetCacheReserveClearResponse]
type zoneCacheGetCacheReserveClearResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetCacheReserveClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetCacheReserveClearResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneCacheGetCacheReserveClearResponseErrorJSON `json:"-"`
}

// zoneCacheGetCacheReserveClearResponseErrorJSON contains the JSON metadata for
// the struct [ZoneCacheGetCacheReserveClearResponseError]
type zoneCacheGetCacheReserveClearResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetCacheReserveClearResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetCacheReserveClearResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneCacheGetCacheReserveClearResponseMessageJSON `json:"-"`
}

// zoneCacheGetCacheReserveClearResponseMessageJSON contains the JSON metadata for
// the struct [ZoneCacheGetCacheReserveClearResponseMessage]
type zoneCacheGetCacheReserveClearResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetCacheReserveClearResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetCacheReserveClearResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheGetCacheReserveClearResponseResultID `json:"id"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State ZoneCacheGetCacheReserveClearResponseResultState `json:"state"`
	JSON  zoneCacheGetCacheReserveClearResponseResultJSON  `json:"-"`
}

// zoneCacheGetCacheReserveClearResponseResultJSON contains the JSON metadata for
// the struct [ZoneCacheGetCacheReserveClearResponseResult]
type zoneCacheGetCacheReserveClearResponseResultJSON struct {
	ID          apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetCacheReserveClearResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCacheGetCacheReserveClearResponseResultID string

const (
	ZoneCacheGetCacheReserveClearResponseResultIDCacheReserveClear ZoneCacheGetCacheReserveClearResponseResultID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type ZoneCacheGetCacheReserveClearResponseResultState string

const (
	ZoneCacheGetCacheReserveClearResponseResultStateInProgress ZoneCacheGetCacheReserveClearResponseResultState = "In-progress"
	ZoneCacheGetCacheReserveClearResponseResultStateCompleted  ZoneCacheGetCacheReserveClearResponseResultState = "Completed"
)

// Whether the API call was successful
type ZoneCacheGetCacheReserveClearResponseSuccess bool

const (
	ZoneCacheGetCacheReserveClearResponseSuccessTrue ZoneCacheGetCacheReserveClearResponseSuccess = true
)

type ZoneCacheGetOriginPostQuantumEncryptionResponse struct {
	Errors   []ZoneCacheGetOriginPostQuantumEncryptionResponseError   `json:"errors"`
	Messages []ZoneCacheGetOriginPostQuantumEncryptionResponseMessage `json:"messages"`
	Result   interface{}                                              `json:"result"`
	// Whether the API call was successful
	Success ZoneCacheGetOriginPostQuantumEncryptionResponseSuccess `json:"success"`
	JSON    zoneCacheGetOriginPostQuantumEncryptionResponseJSON    `json:"-"`
}

// zoneCacheGetOriginPostQuantumEncryptionResponseJSON contains the JSON metadata
// for the struct [ZoneCacheGetOriginPostQuantumEncryptionResponse]
type zoneCacheGetOriginPostQuantumEncryptionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetOriginPostQuantumEncryptionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetOriginPostQuantumEncryptionResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneCacheGetOriginPostQuantumEncryptionResponseErrorJSON `json:"-"`
}

// zoneCacheGetOriginPostQuantumEncryptionResponseErrorJSON contains the JSON
// metadata for the struct [ZoneCacheGetOriginPostQuantumEncryptionResponseError]
type zoneCacheGetOriginPostQuantumEncryptionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetOriginPostQuantumEncryptionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetOriginPostQuantumEncryptionResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneCacheGetOriginPostQuantumEncryptionResponseMessageJSON `json:"-"`
}

// zoneCacheGetOriginPostQuantumEncryptionResponseMessageJSON contains the JSON
// metadata for the struct [ZoneCacheGetOriginPostQuantumEncryptionResponseMessage]
type zoneCacheGetOriginPostQuantumEncryptionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetOriginPostQuantumEncryptionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCacheGetOriginPostQuantumEncryptionResponseSuccess bool

const (
	ZoneCacheGetOriginPostQuantumEncryptionResponseSuccessTrue ZoneCacheGetOriginPostQuantumEncryptionResponseSuccess = true
)

type ZoneCacheGetRegionalTieredCacheResponse struct {
	Errors   []ZoneCacheGetRegionalTieredCacheResponseError   `json:"errors"`
	Messages []ZoneCacheGetRegionalTieredCacheResponseMessage `json:"messages"`
	Result   ZoneCacheGetRegionalTieredCacheResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCacheGetRegionalTieredCacheResponseSuccess `json:"success"`
	JSON    zoneCacheGetRegionalTieredCacheResponseJSON    `json:"-"`
}

// zoneCacheGetRegionalTieredCacheResponseJSON contains the JSON metadata for the
// struct [ZoneCacheGetRegionalTieredCacheResponse]
type zoneCacheGetRegionalTieredCacheResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetRegionalTieredCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetRegionalTieredCacheResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneCacheGetRegionalTieredCacheResponseErrorJSON `json:"-"`
}

// zoneCacheGetRegionalTieredCacheResponseErrorJSON contains the JSON metadata for
// the struct [ZoneCacheGetRegionalTieredCacheResponseError]
type zoneCacheGetRegionalTieredCacheResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetRegionalTieredCacheResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetRegionalTieredCacheResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneCacheGetRegionalTieredCacheResponseMessageJSON `json:"-"`
}

// zoneCacheGetRegionalTieredCacheResponseMessageJSON contains the JSON metadata
// for the struct [ZoneCacheGetRegionalTieredCacheResponseMessage]
type zoneCacheGetRegionalTieredCacheResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetRegionalTieredCacheResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCacheGetRegionalTieredCacheResponseResult struct {
	// ID of the zone setting.
	ID ZoneCacheGetRegionalTieredCacheResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value ZoneCacheGetRegionalTieredCacheResponseResultValue `json:"value"`
	JSON  zoneCacheGetRegionalTieredCacheResponseResultJSON  `json:"-"`
}

// zoneCacheGetRegionalTieredCacheResponseResultJSON contains the JSON metadata for
// the struct [ZoneCacheGetRegionalTieredCacheResponseResult]
type zoneCacheGetRegionalTieredCacheResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetRegionalTieredCacheResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCacheGetRegionalTieredCacheResponseResultID string

const (
	ZoneCacheGetRegionalTieredCacheResponseResultIDTcRegional ZoneCacheGetRegionalTieredCacheResponseResultID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type ZoneCacheGetRegionalTieredCacheResponseResultValue struct {
	// ID of the zone setting.
	ID ZoneCacheGetRegionalTieredCacheResponseResultValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                              `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCacheGetRegionalTieredCacheResponseResultValueJSON `json:"-"`
}

// zoneCacheGetRegionalTieredCacheResponseResultValueJSON contains the JSON
// metadata for the struct [ZoneCacheGetRegionalTieredCacheResponseResultValue]
type zoneCacheGetRegionalTieredCacheResponseResultValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCacheGetRegionalTieredCacheResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCacheGetRegionalTieredCacheResponseResultValueID string

const (
	ZoneCacheGetRegionalTieredCacheResponseResultValueIDTcRegional ZoneCacheGetRegionalTieredCacheResponseResultValueID = "tc_regional"
)

// Whether the API call was successful
type ZoneCacheGetRegionalTieredCacheResponseSuccess bool

const (
	ZoneCacheGetRegionalTieredCacheResponseSuccessTrue ZoneCacheGetRegionalTieredCacheResponseSuccess = true
)

type ZoneCachePatchRegionalTieredCacheResponse struct {
	Errors   []ZoneCachePatchRegionalTieredCacheResponseError   `json:"errors"`
	Messages []ZoneCachePatchRegionalTieredCacheResponseMessage `json:"messages"`
	Result   ZoneCachePatchRegionalTieredCacheResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachePatchRegionalTieredCacheResponseSuccess `json:"success"`
	JSON    zoneCachePatchRegionalTieredCacheResponseJSON    `json:"-"`
}

// zoneCachePatchRegionalTieredCacheResponseJSON contains the JSON metadata for the
// struct [ZoneCachePatchRegionalTieredCacheResponse]
type zoneCachePatchRegionalTieredCacheResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePatchRegionalTieredCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePatchRegionalTieredCacheResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneCachePatchRegionalTieredCacheResponseErrorJSON `json:"-"`
}

// zoneCachePatchRegionalTieredCacheResponseErrorJSON contains the JSON metadata
// for the struct [ZoneCachePatchRegionalTieredCacheResponseError]
type zoneCachePatchRegionalTieredCacheResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePatchRegionalTieredCacheResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePatchRegionalTieredCacheResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zoneCachePatchRegionalTieredCacheResponseMessageJSON `json:"-"`
}

// zoneCachePatchRegionalTieredCacheResponseMessageJSON contains the JSON metadata
// for the struct [ZoneCachePatchRegionalTieredCacheResponseMessage]
type zoneCachePatchRegionalTieredCacheResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePatchRegionalTieredCacheResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePatchRegionalTieredCacheResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachePatchRegionalTieredCacheResponseResultID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Instructs Cloudflare to check a regional hub data center on the way to your
	// upper tier. This can help improve performance for smart and custom tiered cache
	// topologies.
	Value ZoneCachePatchRegionalTieredCacheResponseResultValue `json:"value"`
	JSON  zoneCachePatchRegionalTieredCacheResponseResultJSON  `json:"-"`
}

// zoneCachePatchRegionalTieredCacheResponseResultJSON contains the JSON metadata
// for the struct [ZoneCachePatchRegionalTieredCacheResponseResult]
type zoneCachePatchRegionalTieredCacheResponseResultJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePatchRegionalTieredCacheResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachePatchRegionalTieredCacheResponseResultID string

const (
	ZoneCachePatchRegionalTieredCacheResponseResultIDTcRegional ZoneCachePatchRegionalTieredCacheResponseResultID = "tc_regional"
)

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
type ZoneCachePatchRegionalTieredCacheResponseResultValue struct {
	// ID of the zone setting.
	ID ZoneCachePatchRegionalTieredCacheResponseResultValueID `json:"id"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneCachePatchRegionalTieredCacheResponseResultValueJSON `json:"-"`
}

// zoneCachePatchRegionalTieredCacheResponseResultValueJSON contains the JSON
// metadata for the struct [ZoneCachePatchRegionalTieredCacheResponseResultValue]
type zoneCachePatchRegionalTieredCacheResponseResultValueJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePatchRegionalTieredCacheResponseResultValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachePatchRegionalTieredCacheResponseResultValueID string

const (
	ZoneCachePatchRegionalTieredCacheResponseResultValueIDTcRegional ZoneCachePatchRegionalTieredCacheResponseResultValueID = "tc_regional"
)

// Whether the API call was successful
type ZoneCachePatchRegionalTieredCacheResponseSuccess bool

const (
	ZoneCachePatchRegionalTieredCacheResponseSuccessTrue ZoneCachePatchRegionalTieredCacheResponseSuccess = true
)

type ZoneCachePostCacheReserveClearResponse struct {
	Errors   []ZoneCachePostCacheReserveClearResponseError   `json:"errors"`
	Messages []ZoneCachePostCacheReserveClearResponseMessage `json:"messages"`
	Result   ZoneCachePostCacheReserveClearResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCachePostCacheReserveClearResponseSuccess `json:"success"`
	JSON    zoneCachePostCacheReserveClearResponseJSON    `json:"-"`
}

// zoneCachePostCacheReserveClearResponseJSON contains the JSON metadata for the
// struct [ZoneCachePostCacheReserveClearResponse]
type zoneCachePostCacheReserveClearResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePostCacheReserveClearResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePostCacheReserveClearResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneCachePostCacheReserveClearResponseErrorJSON `json:"-"`
}

// zoneCachePostCacheReserveClearResponseErrorJSON contains the JSON metadata for
// the struct [ZoneCachePostCacheReserveClearResponseError]
type zoneCachePostCacheReserveClearResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePostCacheReserveClearResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePostCacheReserveClearResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneCachePostCacheReserveClearResponseMessageJSON `json:"-"`
}

// zoneCachePostCacheReserveClearResponseMessageJSON contains the JSON metadata for
// the struct [ZoneCachePostCacheReserveClearResponseMessage]
type zoneCachePostCacheReserveClearResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePostCacheReserveClearResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePostCacheReserveClearResponseResult struct {
	// ID of the zone setting.
	ID ZoneCachePostCacheReserveClearResponseResultID `json:"id"`
	// The time that the latest Cache Reserve Clear operation completed.
	EndTs time.Time `json:"end_ts" format:"date-time"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// The time that the latest Cache Reserve Clear operation started.
	StartTs time.Time `json:"start_ts" format:"date-time"`
	// The current state of the Cache Reserve Clear operation.
	State ZoneCachePostCacheReserveClearResponseResultState `json:"state"`
	JSON  zoneCachePostCacheReserveClearResponseResultJSON  `json:"-"`
}

// zoneCachePostCacheReserveClearResponseResultJSON contains the JSON metadata for
// the struct [ZoneCachePostCacheReserveClearResponseResult]
type zoneCachePostCacheReserveClearResponseResultJSON struct {
	ID          apijson.Field
	EndTs       apijson.Field
	ModifiedOn  apijson.Field
	StartTs     apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePostCacheReserveClearResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type ZoneCachePostCacheReserveClearResponseResultID string

const (
	ZoneCachePostCacheReserveClearResponseResultIDCacheReserveClear ZoneCachePostCacheReserveClearResponseResultID = "cache_reserve_clear"
)

// The current state of the Cache Reserve Clear operation.
type ZoneCachePostCacheReserveClearResponseResultState string

const (
	ZoneCachePostCacheReserveClearResponseResultStateInProgress ZoneCachePostCacheReserveClearResponseResultState = "In-progress"
	ZoneCachePostCacheReserveClearResponseResultStateCompleted  ZoneCachePostCacheReserveClearResponseResultState = "Completed"
)

// Whether the API call was successful
type ZoneCachePostCacheReserveClearResponseSuccess bool

const (
	ZoneCachePostCacheReserveClearResponseSuccessTrue ZoneCachePostCacheReserveClearResponseSuccess = true
)

type ZoneCachePutOriginPostQuantumEncryptionResponse struct {
	Errors   []ZoneCachePutOriginPostQuantumEncryptionResponseError   `json:"errors"`
	Messages []ZoneCachePutOriginPostQuantumEncryptionResponseMessage `json:"messages"`
	Result   interface{}                                              `json:"result"`
	// Whether the API call was successful
	Success ZoneCachePutOriginPostQuantumEncryptionResponseSuccess `json:"success"`
	JSON    zoneCachePutOriginPostQuantumEncryptionResponseJSON    `json:"-"`
}

// zoneCachePutOriginPostQuantumEncryptionResponseJSON contains the JSON metadata
// for the struct [ZoneCachePutOriginPostQuantumEncryptionResponse]
type zoneCachePutOriginPostQuantumEncryptionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePutOriginPostQuantumEncryptionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePutOriginPostQuantumEncryptionResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zoneCachePutOriginPostQuantumEncryptionResponseErrorJSON `json:"-"`
}

// zoneCachePutOriginPostQuantumEncryptionResponseErrorJSON contains the JSON
// metadata for the struct [ZoneCachePutOriginPostQuantumEncryptionResponseError]
type zoneCachePutOriginPostQuantumEncryptionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePutOriginPostQuantumEncryptionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCachePutOriginPostQuantumEncryptionResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zoneCachePutOriginPostQuantumEncryptionResponseMessageJSON `json:"-"`
}

// zoneCachePutOriginPostQuantumEncryptionResponseMessageJSON contains the JSON
// metadata for the struct [ZoneCachePutOriginPostQuantumEncryptionResponseMessage]
type zoneCachePutOriginPostQuantumEncryptionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCachePutOriginPostQuantumEncryptionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCachePutOriginPostQuantumEncryptionResponseSuccess bool

const (
	ZoneCachePutOriginPostQuantumEncryptionResponseSuccessTrue ZoneCachePutOriginPostQuantumEncryptionResponseSuccess = true
)

type ZoneCachePatchRegionalTieredCacheParams struct {
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[ZoneCachePatchRegionalTieredCacheParamsValue] `json:"value,required"`
}

func (r ZoneCachePatchRegionalTieredCacheParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type ZoneCachePatchRegionalTieredCacheParamsValue string

const (
	ZoneCachePatchRegionalTieredCacheParamsValueOn  ZoneCachePatchRegionalTieredCacheParamsValue = "on"
	ZoneCachePatchRegionalTieredCacheParamsValueOff ZoneCachePatchRegionalTieredCacheParamsValue = "off"
)

type ZoneCachePutOriginPostQuantumEncryptionParams struct {
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[ZoneCachePutOriginPostQuantumEncryptionParamsValue] `json:"value,required"`
}

func (r ZoneCachePutOriginPostQuantumEncryptionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type ZoneCachePutOriginPostQuantumEncryptionParamsValue string

const (
	ZoneCachePutOriginPostQuantumEncryptionParamsValuePreferred ZoneCachePutOriginPostQuantumEncryptionParamsValue = "preferred"
	ZoneCachePutOriginPostQuantumEncryptionParamsValueSupported ZoneCachePutOriginPostQuantumEncryptionParamsValue = "supported"
	ZoneCachePutOriginPostQuantumEncryptionParamsValueOff       ZoneCachePutOriginPostQuantumEncryptionParamsValue = "off"
)
