// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// CachTieredCacheSmartTopologyEnableService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewCachTieredCacheSmartTopologyEnableService] method instead.
type CachTieredCacheSmartTopologyEnableService struct {
	Options []option.RequestOption
}

// NewCachTieredCacheSmartTopologyEnableService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCachTieredCacheSmartTopologyEnableService(opts ...option.RequestOption) (r *CachTieredCacheSmartTopologyEnableService) {
	r = &CachTieredCacheSmartTopologyEnableService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *CachTieredCacheSmartTopologyEnableService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CachTieredCacheSmartTopologyEnableDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachTieredCacheSmartTopologyEnableDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Smart Tiered Cache setting
func (r *CachTieredCacheSmartTopologyEnableService) SmartTieredCacheGetSmartTieredCacheSetting(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Cache
func (r *CachTieredCacheSmartTopologyEnableService) SmartTieredCachePatchSmartTieredCacheSetting(ctx context.Context, zoneID string, body CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams, opts ...option.RequestOption) (res *CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CachTieredCacheSmartTopologyEnableDeleteResponseUnknown] or
// [shared.UnionString].
type CachTieredCacheSmartTopologyEnableDeleteResponse interface {
	ImplementsCachTieredCacheSmartTopologyEnableDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CachTieredCacheSmartTopologyEnableDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseUnknown]
// or [shared.UnionString].
type CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse interface {
	ImplementsCachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseUnknown]
// or [shared.UnionString].
type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse interface {
	ImplementsCachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CachTieredCacheSmartTopologyEnableDeleteResponseEnvelope struct {
	Errors   []CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CachTieredCacheSmartTopologyEnableDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeJSON    `json:"-"`
}

// cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [CachTieredCacheSmartTopologyEnableDeleteResponseEnvelope]
type cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrors struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrors]
type cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessages struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessages]
type cachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeSuccess bool

const (
	CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeSuccessTrue CachTieredCacheSmartTopologyEnableDeleteResponseEnvelopeSuccess = true
)

type CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelope struct {
	Errors   []CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessages `json:"messages,required"`
	Result   CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeJSON    `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelope]
type cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrors]
type cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessages]
type cachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeSuccess bool

const (
	CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeSuccessTrue CachTieredCacheSmartTopologyEnableSmartTieredCacheGetSmartTieredCacheSettingResponseEnvelopeSuccess = true
)

type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams struct {
	// Enables Tiered Cache.
	Value param.Field[CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue] `json:"value,required"`
}

func (r CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue string

const (
	CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValueOn  CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue = "on"
	CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValueOff CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingParamsValue = "off"
)

type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelope struct {
	Errors   []CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessages `json:"messages,required"`
	Result   CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeJSON    `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelope]
type cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrors]
type cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessages]
type cachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeSuccess bool

const (
	CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeSuccessTrue CachTieredCacheSmartTopologyEnableSmartTieredCachePatchSmartTieredCacheSettingResponseEnvelopeSuccess = true
)
