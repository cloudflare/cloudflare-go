// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo

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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r TieredCachingEditParamsValue) IsKnown() bool {
	switch r {
	case TieredCachingEditParamsValueOn, TieredCachingEditParamsValueOff:
		return true
	}
	return false
}

type TieredCachingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   TieredCachingEditResponse `json:"result,required"`
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

// Whether the API call was successful
type TieredCachingEditResponseEnvelopeSuccess bool

const (
	TieredCachingEditResponseEnvelopeSuccessTrue TieredCachingEditResponseEnvelopeSuccess = true
)

func (r TieredCachingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TieredCachingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TieredCachingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TieredCachingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   TieredCachingGetResponse `json:"result,required"`
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

// Whether the API call was successful
type TieredCachingGetResponseEnvelopeSuccess bool

const (
	TieredCachingGetResponseEnvelopeSuccessTrue TieredCachingGetResponseEnvelopeSuccess = true
)

func (r TieredCachingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TieredCachingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
