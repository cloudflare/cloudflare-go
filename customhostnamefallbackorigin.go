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

// CustomHostnameFallbackOriginService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewCustomHostnameFallbackOriginService] method instead.
type CustomHostnameFallbackOriginService struct {
	Options []option.RequestOption
}

// NewCustomHostnameFallbackOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCustomHostnameFallbackOriginService(opts ...option.RequestOption) (r *CustomHostnameFallbackOriginService) {
	r = &CustomHostnameFallbackOriginService{}
	r.Options = opts
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Update(ctx context.Context, params CustomHostnameFallbackOriginUpdateParams, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Delete(ctx context.Context, body CustomHostnameFallbackOriginDeleteParams, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Get(ctx context.Context, query CustomHostnameFallbackOriginGetParams, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [CustomHostnameFallbackOriginUpdateResponseUnknown] or
// [shared.UnionString].
type CustomHostnameFallbackOriginUpdateResponse interface {
	ImplementsCustomHostnameFallbackOriginUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomHostnameFallbackOriginUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CustomHostnameFallbackOriginDeleteResponseUnknown] or
// [shared.UnionString].
type CustomHostnameFallbackOriginDeleteResponse interface {
	ImplementsCustomHostnameFallbackOriginDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomHostnameFallbackOriginDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [CustomHostnameFallbackOriginGetResponseUnknown] or
// [shared.UnionString].
type CustomHostnameFallbackOriginGetResponse interface {
	ImplementsCustomHostnameFallbackOriginGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomHostnameFallbackOriginGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type CustomHostnameFallbackOriginUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r CustomHostnameFallbackOriginUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameFallbackOriginUpdateResponseEnvelope struct {
	Errors   []CustomHostnameFallbackOriginUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameFallbackOriginUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameFallbackOriginUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameFallbackOriginUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameFallbackOriginUpdateResponseEnvelopeJSON    `json:"-"`
}

// customHostnameFallbackOriginUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [CustomHostnameFallbackOriginUpdateResponseEnvelope]
type customHostnameFallbackOriginUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    customHostnameFallbackOriginUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameFallbackOriginUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginUpdateResponseEnvelopeErrors]
type customHostnameFallbackOriginUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    customHostnameFallbackOriginUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameFallbackOriginUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginUpdateResponseEnvelopeMessages]
type customHostnameFallbackOriginUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameFallbackOriginUpdateResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginUpdateResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginUpdateResponseEnvelopeSuccess = true
)

type CustomHostnameFallbackOriginDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomHostnameFallbackOriginDeleteResponseEnvelope struct {
	Errors   []CustomHostnameFallbackOriginDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameFallbackOriginDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameFallbackOriginDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameFallbackOriginDeleteResponseEnvelopeJSON    `json:"-"`
}

// customHostnameFallbackOriginDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [CustomHostnameFallbackOriginDeleteResponseEnvelope]
type customHostnameFallbackOriginDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    customHostnameFallbackOriginDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameFallbackOriginDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginDeleteResponseEnvelopeErrors]
type customHostnameFallbackOriginDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    customHostnameFallbackOriginDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameFallbackOriginDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginDeleteResponseEnvelopeMessages]
type customHostnameFallbackOriginDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccess = true
)

type CustomHostnameFallbackOriginGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomHostnameFallbackOriginGetResponseEnvelope struct {
	Errors   []CustomHostnameFallbackOriginGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameFallbackOriginGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameFallbackOriginGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameFallbackOriginGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameFallbackOriginGetResponseEnvelopeJSON    `json:"-"`
}

// customHostnameFallbackOriginGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [CustomHostnameFallbackOriginGetResponseEnvelope]
type customHostnameFallbackOriginGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    customHostnameFallbackOriginGetResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameFallbackOriginGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CustomHostnameFallbackOriginGetResponseEnvelopeErrors]
type customHostnameFallbackOriginGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameFallbackOriginGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    customHostnameFallbackOriginGetResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameFallbackOriginGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginGetResponseEnvelopeMessages]
type customHostnameFallbackOriginGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameFallbackOriginGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameFallbackOriginGetResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginGetResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginGetResponseEnvelopeSuccess = true
)
