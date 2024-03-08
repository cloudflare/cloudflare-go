// File generated from our OpenAPI spec by Stainless.

package cache

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// SmartTieredCachedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSmartTieredCachedService] method
// instead.
type SmartTieredCachedService struct {
	Options []option.RequestOption
}

// NewSmartTieredCachedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmartTieredCachedService(opts ...option.RequestOption) (r *SmartTieredCachedService) {
	r = &SmartTieredCachedService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *SmartTieredCachedService) Delete(ctx context.Context, body SmartTieredCachedDeleteParams, opts ...option.RequestOption) (res *SmartTieredCachedDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCachedDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Cache
func (r *SmartTieredCachedService) Edit(ctx context.Context, params SmartTieredCachedEditParams, opts ...option.RequestOption) (res *SmartTieredCachedEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCachedEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Smart Tiered Cache setting
func (r *SmartTieredCachedService) Get(ctx context.Context, query SmartTieredCachedGetParams, opts ...option.RequestOption) (res *SmartTieredCachedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SmartTieredCachedGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [cache.SmartTieredCachedDeleteResponseUnknown] or
// [shared.UnionString].
type SmartTieredCachedDeleteResponse interface {
	ImplementsCacheSmartTieredCachedDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCachedDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [cache.SmartTieredCachedEditResponseUnknown] or
// [shared.UnionString].
type SmartTieredCachedEditResponse interface {
	ImplementsCacheSmartTieredCachedEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCachedEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [cache.SmartTieredCachedGetResponseUnknown] or
// [shared.UnionString].
type SmartTieredCachedGetResponse interface {
	ImplementsCacheSmartTieredCachedGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SmartTieredCachedGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SmartTieredCachedDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartTieredCachedDeleteResponseEnvelope struct {
	Errors   []SmartTieredCachedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCachedDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCachedDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCachedDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCachedDeleteResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCachedDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCachedDeleteResponseEnvelope]
type smartTieredCachedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    smartTieredCachedDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCachedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SmartTieredCachedDeleteResponseEnvelopeErrors]
type smartTieredCachedDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    smartTieredCachedDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCachedDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SmartTieredCachedDeleteResponseEnvelopeMessages]
type smartTieredCachedDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCachedDeleteResponseEnvelopeSuccess bool

const (
	SmartTieredCachedDeleteResponseEnvelopeSuccessTrue SmartTieredCachedDeleteResponseEnvelopeSuccess = true
)

type SmartTieredCachedEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Cache.
	Value param.Field[SmartTieredCachedEditParamsValue] `json:"value,required"`
}

func (r SmartTieredCachedEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type SmartTieredCachedEditParamsValue string

const (
	SmartTieredCachedEditParamsValueOn  SmartTieredCachedEditParamsValue = "on"
	SmartTieredCachedEditParamsValueOff SmartTieredCachedEditParamsValue = "off"
)

type SmartTieredCachedEditResponseEnvelope struct {
	Errors   []SmartTieredCachedEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCachedEditResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCachedEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCachedEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCachedEditResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCachedEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCachedEditResponseEnvelope]
type smartTieredCachedEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    smartTieredCachedEditResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCachedEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SmartTieredCachedEditResponseEnvelopeErrors]
type smartTieredCachedEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    smartTieredCachedEditResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCachedEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SmartTieredCachedEditResponseEnvelopeMessages]
type smartTieredCachedEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCachedEditResponseEnvelopeSuccess bool

const (
	SmartTieredCachedEditResponseEnvelopeSuccessTrue SmartTieredCachedEditResponseEnvelopeSuccess = true
)

type SmartTieredCachedGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartTieredCachedGetResponseEnvelope struct {
	Errors   []SmartTieredCachedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SmartTieredCachedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   SmartTieredCachedGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SmartTieredCachedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartTieredCachedGetResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCachedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCachedGetResponseEnvelope]
type smartTieredCachedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    smartTieredCachedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// smartTieredCachedGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SmartTieredCachedGetResponseEnvelopeErrors]
type smartTieredCachedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SmartTieredCachedGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    smartTieredCachedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// smartTieredCachedGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SmartTieredCachedGetResponseEnvelopeMessages]
type smartTieredCachedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCachedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCachedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCachedGetResponseEnvelopeSuccess bool

const (
	SmartTieredCachedGetResponseEnvelopeSuccessTrue SmartTieredCachedGetResponseEnvelopeSuccess = true
)
