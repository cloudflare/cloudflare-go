// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// SmartTieredCacheService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSmartTieredCacheService] method
// instead.
type SmartTieredCacheService struct {
	Options []option.RequestOption
}

// NewSmartTieredCacheService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmartTieredCacheService(opts ...option.RequestOption) (r *SmartTieredCacheService) {
	r = &SmartTieredCacheService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *SmartTieredCacheService) Delete(ctx context.Context, params SmartTieredCacheDeleteParams, opts ...option.RequestOption) (res *SmartTieredCacheDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCacheDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Cache
func (r *SmartTieredCacheService) Edit(ctx context.Context, params SmartTieredCacheEditParams, opts ...option.RequestOption) (res *SmartTieredCacheEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCacheEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Smart Tiered Cache setting
func (r *SmartTieredCacheService) Get(ctx context.Context, query SmartTieredCacheGetParams, opts ...option.RequestOption) (res *SmartTieredCacheGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCacheGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [cache.SmartTieredCacheDeleteResponseUnknown] or
// [shared.UnionString].
type SmartTieredCacheDeleteResponse interface {
	ImplementsCacheSmartTieredCacheDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCacheDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [cache.SmartTieredCacheEditResponseUnknown] or
// [shared.UnionString].
type SmartTieredCacheEditResponse interface {
	ImplementsCacheSmartTieredCacheEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCacheEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [cache.SmartTieredCacheGetResponseUnknown] or
// [shared.UnionString].
type SmartTieredCacheGetResponse interface {
	ImplementsCacheSmartTieredCacheGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCacheGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SmartTieredCacheDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r SmartTieredCacheDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SmartTieredCacheDeleteResponseEnvelope struct {
	Errors   []SmartTieredCacheDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCacheDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCacheDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCacheDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCacheDeleteResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheDeleteResponseEnvelope]
type smartTieredCacheDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    smartTieredCacheDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCacheDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SmartTieredCacheDeleteResponseEnvelopeErrors]
type smartTieredCacheDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    smartTieredCacheDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCacheDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SmartTieredCacheDeleteResponseEnvelopeMessages]
type smartTieredCacheDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheDeleteResponseEnvelopeSuccess bool

const (
	SmartTieredCacheDeleteResponseEnvelopeSuccessTrue SmartTieredCacheDeleteResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartTieredCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Cache.
	Value param.Field[SmartTieredCacheEditParamsValue] `json:"value,required"`
}

func (r SmartTieredCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type SmartTieredCacheEditParamsValue string

const (
	SmartTieredCacheEditParamsValueOn  SmartTieredCacheEditParamsValue = "on"
	SmartTieredCacheEditParamsValueOff SmartTieredCacheEditParamsValue = "off"
)

func (r SmartTieredCacheEditParamsValue) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditParamsValueOn, SmartTieredCacheEditParamsValueOff:
		return true
	}
	return false
}

type SmartTieredCacheEditResponseEnvelope struct {
	Errors   []SmartTieredCacheEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCacheEditResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCacheEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCacheEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCacheEditResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheEditResponseEnvelope]
type smartTieredCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    smartTieredCacheEditResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCacheEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SmartTieredCacheEditResponseEnvelopeErrors]
type smartTieredCacheEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    smartTieredCacheEditResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCacheEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SmartTieredCacheEditResponseEnvelopeMessages]
type smartTieredCacheEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheEditResponseEnvelopeSuccess bool

const (
	SmartTieredCacheEditResponseEnvelopeSuccessTrue SmartTieredCacheEditResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartTieredCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartTieredCacheGetResponseEnvelope struct {
	Errors   []SmartTieredCacheGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCacheGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCacheGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCacheGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCacheGetResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheGetResponseEnvelope]
type smartTieredCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    smartTieredCacheGetResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCacheGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SmartTieredCacheGetResponseEnvelopeErrors]
type smartTieredCacheGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCacheGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    smartTieredCacheGetResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCacheGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SmartTieredCacheGetResponseEnvelopeMessages]
type smartTieredCacheGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheGetResponseEnvelopeSuccess bool

const (
	SmartTieredCacheGetResponseEnvelopeSuccessTrue SmartTieredCacheGetResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
