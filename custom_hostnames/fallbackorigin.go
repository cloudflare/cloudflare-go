// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// FallbackOriginService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFallbackOriginService] method instead.
type FallbackOriginService struct {
	Options []option.RequestOption
}

// NewFallbackOriginService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFallbackOriginService(opts ...option.RequestOption) (r *FallbackOriginService) {
	r = &FallbackOriginService{}
	r.Options = opts
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Update(ctx context.Context, params FallbackOriginUpdateParams, opts ...option.RequestOption) (res *FallbackOriginUpdateResponseUnion, err error) {
	var env FallbackOriginUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Delete(ctx context.Context, body FallbackOriginDeleteParams, opts ...option.RequestOption) (res *FallbackOriginDeleteResponseUnion, err error) {
	var env FallbackOriginDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Get(ctx context.Context, query FallbackOriginGetParams, opts ...option.RequestOption) (res *FallbackOriginGetResponseUnion, err error) {
	var env FallbackOriginGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [custom_hostnames.FallbackOriginUpdateResponseUnknown] or
// [shared.UnionString].
type FallbackOriginUpdateResponseUnion interface {
	ImplementsCustomHostnamesFallbackOriginUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [custom_hostnames.FallbackOriginDeleteResponseUnknown] or
// [shared.UnionString].
type FallbackOriginDeleteResponseUnion interface {
	ImplementsCustomHostnamesFallbackOriginDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [custom_hostnames.FallbackOriginGetResponseUnknown] or
// [shared.UnionString].
type FallbackOriginGetResponseUnion interface {
	ImplementsCustomHostnamesFallbackOriginGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FallbackOriginUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r FallbackOriginUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FallbackOriginUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FallbackOriginUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginUpdateResponseUnion           `json:"result"`
	JSON    fallbackOriginUpdateResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginUpdateResponseEnvelope]
type fallbackOriginUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FallbackOriginUpdateResponseEnvelopeSuccess bool

const (
	FallbackOriginUpdateResponseEnvelopeSuccessTrue FallbackOriginUpdateResponseEnvelopeSuccess = true
)

func (r FallbackOriginUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FallbackOriginDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FallbackOriginDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FallbackOriginDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginDeleteResponseUnion           `json:"result"`
	JSON    fallbackOriginDeleteResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginDeleteResponseEnvelope]
type fallbackOriginDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FallbackOriginDeleteResponseEnvelopeSuccess bool

const (
	FallbackOriginDeleteResponseEnvelopeSuccessTrue FallbackOriginDeleteResponseEnvelopeSuccess = true
)

func (r FallbackOriginDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FallbackOriginGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FallbackOriginGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success FallbackOriginGetResponseEnvelopeSuccess `json:"success,required"`
	Result  FallbackOriginGetResponseUnion           `json:"result"`
	JSON    fallbackOriginGetResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FallbackOriginGetResponseEnvelope]
type fallbackOriginGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FallbackOriginGetResponseEnvelopeSuccess bool

const (
	FallbackOriginGetResponseEnvelopeSuccessTrue FallbackOriginGetResponseEnvelopeSuccess = true
)

func (r FallbackOriginGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FallbackOriginGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
