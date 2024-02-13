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
func (r *ArgoSmartRoutingService) Update(ctx context.Context, zoneID string, body ArgoSmartRoutingUpdateParams, opts ...option.RequestOption) (res *ArgoSmartRoutingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoSmartRoutingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Argo Smart Routing setting
func (r *ArgoSmartRoutingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ArgoSmartRoutingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoSmartRoutingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/smart_routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ArgoSmartRoutingUpdateResponseUnknown] or
// [shared.UnionString].
type ArgoSmartRoutingUpdateResponse interface {
	ImplementsArgoSmartRoutingUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoSmartRoutingUpdateResponse)(nil)).Elem(),
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

type ArgoSmartRoutingUpdateParams struct {
	// Enables Argo Smart Routing.
	Value param.Field[ArgoSmartRoutingUpdateParamsValue] `json:"value,required"`
}

func (r ArgoSmartRoutingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Argo Smart Routing.
type ArgoSmartRoutingUpdateParamsValue string

const (
	ArgoSmartRoutingUpdateParamsValueOn  ArgoSmartRoutingUpdateParamsValue = "on"
	ArgoSmartRoutingUpdateParamsValueOff ArgoSmartRoutingUpdateParamsValue = "off"
)

type ArgoSmartRoutingUpdateResponseEnvelope struct {
	Errors   []ArgoSmartRoutingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoSmartRoutingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoSmartRoutingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoSmartRoutingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoSmartRoutingUpdateResponseEnvelopeJSON    `json:"-"`
}

// argoSmartRoutingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoSmartRoutingUpdateResponseEnvelope]
type argoSmartRoutingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    argoSmartRoutingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// argoSmartRoutingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ArgoSmartRoutingUpdateResponseEnvelopeErrors]
type argoSmartRoutingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoSmartRoutingUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    argoSmartRoutingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// argoSmartRoutingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ArgoSmartRoutingUpdateResponseEnvelopeMessages]
type argoSmartRoutingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoSmartRoutingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoSmartRoutingUpdateResponseEnvelopeSuccess bool

const (
	ArgoSmartRoutingUpdateResponseEnvelopeSuccessTrue ArgoSmartRoutingUpdateResponseEnvelopeSuccess = true
)

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
