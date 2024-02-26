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

// ArgoSmartRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewArgoSmartRoutingService] method
// instead.
type ArgoSmartRoutingService struct {
	Options []option.RequestOption
}

// NewArgoSmartRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewArgoSmartRoutingService(opts ...option.RequestOption) (r *ArgoSmartRoutingService) {
	r = &ArgoSmartRoutingService{}
	r.Options = opts
	return
}

// Updates enablement of Argo Smart Routing.
func (r *ArgoSmartRoutingService) Edit(ctx context.Context, params ArgoSmartRoutingEditParams, opts ...option.RequestOption) (res *ArgoSmartRoutingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoSmartRoutingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/smart_routing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Argo Smart Routing setting
func (r *ArgoSmartRoutingService) Get(ctx context.Context, query ArgoSmartRoutingGetParams, opts ...option.RequestOption) (res *ArgoSmartRoutingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoSmartRoutingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/smart_routing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ArgoSmartRoutingEditResponseUnknown] or
// [shared.UnionString].
type ArgoSmartRoutingEditResponse interface {
	ImplementsArgoSmartRoutingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoSmartRoutingEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [ArgoSmartRoutingGetResponseUnknown] or [shared.UnionString].
type ArgoSmartRoutingGetResponse interface {
	ImplementsArgoSmartRoutingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoSmartRoutingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ArgoSmartRoutingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Argo Smart Routing.
	Value param.Field[ArgoSmartRoutingEditParamsValue] `json:"value,required"`
}

func (r ArgoSmartRoutingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type ArgoSmartRoutingEditParamsValue string

const (
	ArgoSmartRoutingEditParamsValueOn  ArgoSmartRoutingEditParamsValue = "on"
	ArgoSmartRoutingEditParamsValueOff ArgoSmartRoutingEditParamsValue = "off"
)

type ArgoSmartRoutingEditResponseEnvelope struct {
	Errors   []ArgoSmartRoutingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoSmartRoutingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoSmartRoutingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoSmartRoutingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoSmartRoutingEditResponseEnvelopeJSON    `json:"-"`
}

// argoSmartRoutingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoSmartRoutingEditResponseEnvelope]
type argoSmartRoutingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingEditResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    argoSmartRoutingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// argoSmartRoutingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ArgoSmartRoutingEditResponseEnvelopeErrors]
type argoSmartRoutingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingEditResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    argoSmartRoutingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// argoSmartRoutingEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ArgoSmartRoutingEditResponseEnvelopeMessages]
type argoSmartRoutingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoSmartRoutingEditResponseEnvelopeSuccess bool

const (
	ArgoSmartRoutingEditResponseEnvelopeSuccessTrue ArgoSmartRoutingEditResponseEnvelopeSuccess = true
)

type ArgoSmartRoutingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ArgoSmartRoutingGetResponseEnvelope struct {
	Errors   []ArgoSmartRoutingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoSmartRoutingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoSmartRoutingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoSmartRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoSmartRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// argoSmartRoutingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoSmartRoutingGetResponseEnvelope]
type argoSmartRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    argoSmartRoutingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// argoSmartRoutingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ArgoSmartRoutingGetResponseEnvelopeErrors]
type argoSmartRoutingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    argoSmartRoutingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// argoSmartRoutingGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ArgoSmartRoutingGetResponseEnvelopeMessages]
type argoSmartRoutingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoSmartRoutingGetResponseEnvelopeSuccess bool

const (
	ArgoSmartRoutingGetResponseEnvelopeSuccessTrue ArgoSmartRoutingGetResponseEnvelopeSuccess = true
)
