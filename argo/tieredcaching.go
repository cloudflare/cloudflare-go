// File generated from our OpenAPI spec by Stainless.

package argo

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

// TieredCachingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTieredCachingService] method
// instead.
type TieredCachingService struct {
	Options []option.RequestOption
}

// NewTieredCachingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTieredCachingService(opts ...option.RequestOption) (r *TieredCachingService) {
	r = &TieredCachingService{}
	r.Options = opts
	return
}

// Updates enablement of Tiered Caching
func (r *TieredCachingService) Edit(ctx context.Context, params TieredCachingEditParams, opts ...option.RequestOption) (res *TieredCachingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TieredCachingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Tiered Caching setting
func (r *TieredCachingService) Get(ctx context.Context, query TieredCachingGetParams, opts ...option.RequestOption) (res *TieredCachingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env TieredCachingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [argo.TieredCachingEditResponseUnknown] or
// [shared.UnionString].
type TieredCachingEditResponse interface {
	ImplementsArgoTieredCachingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TieredCachingEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [argo.TieredCachingGetResponseUnknown] or
// [shared.UnionString].
type TieredCachingGetResponse interface {
	ImplementsArgoTieredCachingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TieredCachingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type TieredCachingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Caching.
	Value param.Field[TieredCachingEditParamsValue] `json:"value,required"`
}

func (r TieredCachingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type TieredCachingEditParamsValue string

const (
	TieredCachingEditParamsValueOn  TieredCachingEditParamsValue = "on"
	TieredCachingEditParamsValueOff TieredCachingEditParamsValue = "off"
)

type TieredCachingEditResponseEnvelope struct {
	Errors   []TieredCachingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TieredCachingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TieredCachingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TieredCachingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    tieredCachingEditResponseEnvelopeJSON    `json:"-"`
}

// tieredCachingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [TieredCachingEditResponseEnvelope]
type tieredCachingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TieredCachingEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    tieredCachingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// tieredCachingEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TieredCachingEditResponseEnvelopeErrors]
type tieredCachingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TieredCachingEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    tieredCachingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// tieredCachingEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TieredCachingEditResponseEnvelopeMessages]
type tieredCachingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TieredCachingEditResponseEnvelopeSuccess bool

const (
	TieredCachingEditResponseEnvelopeSuccessTrue TieredCachingEditResponseEnvelopeSuccess = true
)

type TieredCachingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TieredCachingGetResponseEnvelope struct {
	Errors   []TieredCachingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TieredCachingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TieredCachingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success TieredCachingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tieredCachingGetResponseEnvelopeJSON    `json:"-"`
}

// tieredCachingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TieredCachingGetResponseEnvelope]
type tieredCachingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TieredCachingGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    tieredCachingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// tieredCachingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [TieredCachingGetResponseEnvelopeErrors]
type tieredCachingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TieredCachingGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    tieredCachingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// tieredCachingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [TieredCachingGetResponseEnvelopeMessages]
type tieredCachingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TieredCachingGetResponseEnvelopeSuccess bool

const (
	TieredCachingGetResponseEnvelopeSuccessTrue TieredCachingGetResponseEnvelopeSuccess = true
)
