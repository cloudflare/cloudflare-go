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

// CacheTieredCacheSmartTopologyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewCacheTieredCacheSmartTopologyService] method instead.
type CacheTieredCacheSmartTopologyService struct {
	Options []option.RequestOption
}

// NewCacheTieredCacheSmartTopologyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCacheTieredCacheSmartTopologyService(opts ...option.RequestOption) (r *CacheTieredCacheSmartTopologyService) {
	r = &CacheTieredCacheSmartTopologyService{}
	r.Options = opts
	return
}

// Remvoves enablement of Smart Tiered Cache
func (r *CacheTieredCacheSmartTopologyService) Delete(ctx context.Context, body CacheTieredCacheSmartTopologyDeleteParams, opts ...option.RequestOption) (res *CacheTieredCacheSmartTopologyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheTieredCacheSmartTopologyDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Cache
func (r *CacheTieredCacheSmartTopologyService) Edit(ctx context.Context, params CacheTieredCacheSmartTopologyEditParams, opts ...option.RequestOption) (res *CacheTieredCacheSmartTopologyEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheTieredCacheSmartTopologyEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Smart Tiered Cache setting
func (r *CacheTieredCacheSmartTopologyService) Get(ctx context.Context, query CacheTieredCacheSmartTopologyGetParams, opts ...option.RequestOption) (res *CacheTieredCacheSmartTopologyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheTieredCacheSmartTopologyGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CacheTieredCacheSmartTopologyDeleteResponseUnknown] or
// [shared.UnionString].
type CacheTieredCacheSmartTopologyDeleteResponse interface {
	ImplementsCacheTieredCacheSmartTopologyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheTieredCacheSmartTopologyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CacheTieredCacheSmartTopologyEditResponseUnknown] or
// [shared.UnionString].
type CacheTieredCacheSmartTopologyEditResponse interface {
	ImplementsCacheTieredCacheSmartTopologyEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheTieredCacheSmartTopologyEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CacheTieredCacheSmartTopologyGetResponseUnknown] or
// [shared.UnionString].
type CacheTieredCacheSmartTopologyGetResponse interface {
	ImplementsCacheTieredCacheSmartTopologyGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CacheTieredCacheSmartTopologyGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CacheTieredCacheSmartTopologyDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheTieredCacheSmartTopologyDeleteResponseEnvelope struct {
	Errors   []CacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheTieredCacheSmartTopologyDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheTieredCacheSmartTopologyDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheTieredCacheSmartTopologyDeleteResponseEnvelopeJSON    `json:"-"`
}

// cacheTieredCacheSmartTopologyDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [CacheTieredCacheSmartTopologyDeleteResponseEnvelope]
type cacheTieredCacheSmartTopologyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrors]
type cacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [CacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessages]
type cacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheTieredCacheSmartTopologyDeleteResponseEnvelopeSuccess bool

const (
	CacheTieredCacheSmartTopologyDeleteResponseEnvelopeSuccessTrue CacheTieredCacheSmartTopologyDeleteResponseEnvelopeSuccess = true
)

type CacheTieredCacheSmartTopologyEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Cache.
	Value param.Field[CacheTieredCacheSmartTopologyEditParamsValue] `json:"value,required"`
}

func (r CacheTieredCacheSmartTopologyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Cache.
type CacheTieredCacheSmartTopologyEditParamsValue string

const (
	CacheTieredCacheSmartTopologyEditParamsValueOn  CacheTieredCacheSmartTopologyEditParamsValue = "on"
	CacheTieredCacheSmartTopologyEditParamsValueOff CacheTieredCacheSmartTopologyEditParamsValue = "off"
)

type CacheTieredCacheSmartTopologyEditResponseEnvelope struct {
	Errors   []CacheTieredCacheSmartTopologyEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheTieredCacheSmartTopologyEditResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheTieredCacheSmartTopologyEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheTieredCacheSmartTopologyEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheTieredCacheSmartTopologyEditResponseEnvelopeJSON    `json:"-"`
}

// cacheTieredCacheSmartTopologyEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [CacheTieredCacheSmartTopologyEditResponseEnvelope]
type cacheTieredCacheSmartTopologyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyEditResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CacheTieredCacheSmartTopologyEditResponseEnvelopeErrors]
type cacheTieredCacheSmartTopologyEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyEditResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CacheTieredCacheSmartTopologyEditResponseEnvelopeMessages]
type cacheTieredCacheSmartTopologyEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheTieredCacheSmartTopologyEditResponseEnvelopeSuccess bool

const (
	CacheTieredCacheSmartTopologyEditResponseEnvelopeSuccessTrue CacheTieredCacheSmartTopologyEditResponseEnvelopeSuccess = true
)

type CacheTieredCacheSmartTopologyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CacheTieredCacheSmartTopologyGetResponseEnvelope struct {
	Errors   []CacheTieredCacheSmartTopologyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheTieredCacheSmartTopologyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CacheTieredCacheSmartTopologyGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CacheTieredCacheSmartTopologyGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheTieredCacheSmartTopologyGetResponseEnvelopeJSON    `json:"-"`
}

// cacheTieredCacheSmartTopologyGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CacheTieredCacheSmartTopologyGetResponseEnvelope]
type cacheTieredCacheSmartTopologyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CacheTieredCacheSmartTopologyGetResponseEnvelopeErrors]
type cacheTieredCacheSmartTopologyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheTieredCacheSmartTopologyGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    cacheTieredCacheSmartTopologyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheTieredCacheSmartTopologyGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CacheTieredCacheSmartTopologyGetResponseEnvelopeMessages]
type cacheTieredCacheSmartTopologyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheTieredCacheSmartTopologyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheTieredCacheSmartTopologyGetResponseEnvelopeSuccess bool

const (
	CacheTieredCacheSmartTopologyGetResponseEnvelopeSuccessTrue CacheTieredCacheSmartTopologyGetResponseEnvelopeSuccess = true
)
