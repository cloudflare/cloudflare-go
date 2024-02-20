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

// Delete Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update Fallback Origin for Custom Hostnames
func (r *CustomHostnameFallbackOriginService) Replace(ctx context.Context, zoneID string, body CustomHostnameFallbackOriginReplaceParams, opts ...option.RequestOption) (res *CustomHostnameFallbackOriginReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameFallbackOriginReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/fallback_origin", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// Union satisfied by [CustomHostnameFallbackOriginReplaceResponseUnknown] or
// [shared.UnionString].
type CustomHostnameFallbackOriginReplaceResponse interface {
	ImplementsCustomHostnameFallbackOriginReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CustomHostnameFallbackOriginReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

// Whether the API call was successful
type CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginDeleteResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type CustomHostnameFallbackOriginGetResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginGetResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginGetResponseEnvelopeSuccess = true
)

type CustomHostnameFallbackOriginReplaceParams struct {
	// Your origin hostname that requests to your custom hostnames will be sent to.
	Origin param.Field[string] `json:"origin,required"`
}

func (r CustomHostnameFallbackOriginReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameFallbackOriginReplaceResponseEnvelope struct {
	Errors   []CustomHostnameFallbackOriginReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameFallbackOriginReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameFallbackOriginReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameFallbackOriginReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameFallbackOriginReplaceResponseEnvelopeJSON    `json:"-"`
}

// customHostnameFallbackOriginReplaceResponseEnvelopeJSON contains the JSON
// metadata for the struct [CustomHostnameFallbackOriginReplaceResponseEnvelope]
type customHostnameFallbackOriginReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameFallbackOriginReplaceResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    customHostnameFallbackOriginReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameFallbackOriginReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CustomHostnameFallbackOriginReplaceResponseEnvelopeErrors]
type customHostnameFallbackOriginReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameFallbackOriginReplaceResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    customHostnameFallbackOriginReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameFallbackOriginReplaceResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [CustomHostnameFallbackOriginReplaceResponseEnvelopeMessages]
type customHostnameFallbackOriginReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameFallbackOriginReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameFallbackOriginReplaceResponseEnvelopeSuccess bool

const (
	CustomHostnameFallbackOriginReplaceResponseEnvelopeSuccessTrue CustomHostnameFallbackOriginReplaceResponseEnvelopeSuccess = true
)
