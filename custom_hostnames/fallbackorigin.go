// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

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

// FallbackOriginService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFallbackOriginService] method
// instead.
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
func (r *FallbackOriginService) Update(ctx context.Context, params FallbackOriginUpdateParams, opts ...option.RequestOption) (res *FallbackOriginUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FallbackOriginUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Delete(ctx context.Context, params FallbackOriginDeleteParams, opts ...option.RequestOption) (res *FallbackOriginDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FallbackOriginDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *FallbackOriginService) Get(ctx context.Context, query FallbackOriginGetParams, opts ...option.RequestOption) (res *FallbackOriginGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FallbackOriginGetResponseEnvelope
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
type FallbackOriginUpdateResponse interface {
	ImplementsCustomHostnamesFallbackOriginUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [custom_hostnames.FallbackOriginDeleteResponseUnknown] or
// [shared.UnionString].
type FallbackOriginDeleteResponse interface {
	ImplementsCustomHostnamesFallbackOriginDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [custom_hostnames.FallbackOriginGetResponseUnknown] or
// [shared.UnionString].
type FallbackOriginGetResponse interface {
	ImplementsCustomHostnamesFallbackOriginGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FallbackOriginGetResponse)(nil)).Elem(),
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
	Errors   []FallbackOriginUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FallbackOriginUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FallbackOriginUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    fallbackOriginUpdateResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginUpdateResponseEnvelope]
type fallbackOriginUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    fallbackOriginUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FallbackOriginUpdateResponseEnvelopeErrors]
type fallbackOriginUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    fallbackOriginUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// fallbackOriginUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FallbackOriginUpdateResponseEnvelopeMessages]
type fallbackOriginUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r FallbackOriginDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FallbackOriginDeleteResponseEnvelope struct {
	Errors   []FallbackOriginDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FallbackOriginDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FallbackOriginDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    fallbackOriginDeleteResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FallbackOriginDeleteResponseEnvelope]
type fallbackOriginDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    fallbackOriginDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FallbackOriginDeleteResponseEnvelopeErrors]
type fallbackOriginDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    fallbackOriginDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// fallbackOriginDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FallbackOriginDeleteResponseEnvelopeMessages]
type fallbackOriginDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []FallbackOriginGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FallbackOriginGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FallbackOriginGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FallbackOriginGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    fallbackOriginGetResponseEnvelopeJSON    `json:"-"`
}

// fallbackOriginGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FallbackOriginGetResponseEnvelope]
type fallbackOriginGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    fallbackOriginGetResponseEnvelopeErrorsJSON `json:"-"`
}

// fallbackOriginGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FallbackOriginGetResponseEnvelopeErrors]
type fallbackOriginGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FallbackOriginGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    fallbackOriginGetResponseEnvelopeMessagesJSON `json:"-"`
}

// fallbackOriginGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FallbackOriginGetResponseEnvelopeMessages]
type fallbackOriginGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackOriginGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackOriginGetResponseEnvelopeMessagesJSON) RawJSON() string {
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
