// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// CacheSmartTieredCachedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCacheSmartTieredCachedService]
// method instead.
type CacheSmartTieredCachedService struct {
	Options []option.RequestOption
}

// NewCacheSmartTieredCachedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCacheSmartTieredCachedService(opts ...option.RequestOption) (r *CacheSmartTieredCachedService) {
	r = &CacheSmartTieredCachedService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *CacheSmartTieredCachedService) Delete(ctx context.Context, body CacheSmartTieredCachedDeleteParams, opts ...option.RequestOption) (res *CacheSmartTieredCachedDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheSmartTieredCachedDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Cache
func (r *CacheSmartTieredCachedService) Edit(ctx context.Context, params CacheSmartTieredCachedEditParams, opts ...option.RequestOption) (res *CacheSmartTieredCachedEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheSmartTieredCachedEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Smart Tiered Cache setting
func (r *CacheSmartTieredCachedService) Get(ctx context.Context, query CacheSmartTieredCachedGetParams, opts ...option.RequestOption) (res *CacheSmartTieredCachedGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheSmartTieredCachedGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CacheSmartTieredCachedDeleteResponseUnknown] or
// [shared.UnionString].
type CacheSmartTieredCachedDeleteResponse interface {
	ImplementsCacheSmartTieredCachedDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheSmartTieredCachedDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CacheSmartTieredCachedEditResponseUnknown] or
// [shared.UnionString].
type CacheSmartTieredCachedEditResponse interface {
	ImplementsCacheSmartTieredCachedEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheSmartTieredCachedEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CacheSmartTieredCachedGetResponseUnknown] or
// [shared.UnionString].
type CacheSmartTieredCachedGetResponse interface {
	ImplementsCacheSmartTieredCachedGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheSmartTieredCachedGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CacheSmartTieredCachedDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheSmartTieredCachedDeleteResponseEnvelope struct {
	Errors   []CacheSmartTieredCachedDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheSmartTieredCachedDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheSmartTieredCachedDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheSmartTieredCachedDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheSmartTieredCachedDeleteResponseEnvelopeJSON    `json:"-"`
}

// cacheSmartTieredCachedDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [CacheSmartTieredCachedDeleteResponseEnvelope]
type cacheSmartTieredCachedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    cacheSmartTieredCachedDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheSmartTieredCachedDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CacheSmartTieredCachedDeleteResponseEnvelopeErrors]
type cacheSmartTieredCachedDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    cacheSmartTieredCachedDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheSmartTieredCachedDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheSmartTieredCachedDeleteResponseEnvelopeMessages]
type cacheSmartTieredCachedDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheSmartTieredCachedDeleteResponseEnvelopeSuccess bool

const (
	CacheSmartTieredCachedDeleteResponseEnvelopeSuccessTrue CacheSmartTieredCachedDeleteResponseEnvelopeSuccess = true
)

type CacheSmartTieredCachedEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Cache.
	Value param.Field[CacheSmartTieredCachedEditParamsValue] `json:"value,required"`
}

func (r CacheSmartTieredCachedEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type CacheSmartTieredCachedEditParamsValue string

const (
	CacheSmartTieredCachedEditParamsValueOn  CacheSmartTieredCachedEditParamsValue = "on"
	CacheSmartTieredCachedEditParamsValueOff CacheSmartTieredCachedEditParamsValue = "off"
)

type CacheSmartTieredCachedEditResponseEnvelope struct {
	Errors   []CacheSmartTieredCachedEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheSmartTieredCachedEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheSmartTieredCachedEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheSmartTieredCachedEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheSmartTieredCachedEditResponseEnvelopeJSON    `json:"-"`
}

// cacheSmartTieredCachedEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [CacheSmartTieredCachedEditResponseEnvelope]
type cacheSmartTieredCachedEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedEditResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    cacheSmartTieredCachedEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheSmartTieredCachedEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CacheSmartTieredCachedEditResponseEnvelopeErrors]
type cacheSmartTieredCachedEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedEditResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    cacheSmartTieredCachedEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheSmartTieredCachedEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CacheSmartTieredCachedEditResponseEnvelopeMessages]
type cacheSmartTieredCachedEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheSmartTieredCachedEditResponseEnvelopeSuccess bool

const (
	CacheSmartTieredCachedEditResponseEnvelopeSuccessTrue CacheSmartTieredCachedEditResponseEnvelopeSuccess = true
)

type CacheSmartTieredCachedGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheSmartTieredCachedGetResponseEnvelope struct {
	Errors   []CacheSmartTieredCachedGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheSmartTieredCachedGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheSmartTieredCachedGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheSmartTieredCachedGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheSmartTieredCachedGetResponseEnvelopeJSON    `json:"-"`
}

// cacheSmartTieredCachedGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CacheSmartTieredCachedGetResponseEnvelope]
type cacheSmartTieredCachedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    cacheSmartTieredCachedGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheSmartTieredCachedGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CacheSmartTieredCachedGetResponseEnvelopeErrors]
type cacheSmartTieredCachedGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CacheSmartTieredCachedGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    cacheSmartTieredCachedGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheSmartTieredCachedGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CacheSmartTieredCachedGetResponseEnvelopeMessages]
type cacheSmartTieredCachedGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheSmartTieredCachedGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cacheSmartTieredCachedGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CacheSmartTieredCachedGetResponseEnvelopeSuccess bool

const (
	CacheSmartTieredCachedGetResponseEnvelopeSuccessTrue CacheSmartTieredCachedGetResponseEnvelopeSuccess = true
)
