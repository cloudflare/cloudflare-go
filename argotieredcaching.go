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

// ArgoTieredCachingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewArgoTieredCachingService] method
// instead.
type ArgoTieredCachingService struct {
	Options []option.RequestOption
}

// NewArgoTieredCachingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewArgoTieredCachingService(opts ...option.RequestOption) (r *ArgoTieredCachingService) {
	r = &ArgoTieredCachingService{}
	r.Options = opts
	return
}

// Updates enablement of Tiered Caching
func (r *ArgoTieredCachingService) Update(ctx context.Context, zoneID string, body ArgoTieredCachingUpdateParams, opts ...option.RequestOption) (res *ArgoTieredCachingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Tiered Caching setting
func (r *ArgoTieredCachingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ArgoTieredCachingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ArgoTieredCachingUpdateResponseUnknown] or
// [shared.UnionString].
type ArgoTieredCachingUpdateResponse interface {
	ImplementsArgoTieredCachingUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoTieredCachingUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [ArgoTieredCachingGetResponseUnknown] or
// [shared.UnionString].
type ArgoTieredCachingGetResponse interface {
	ImplementsArgoTieredCachingGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoTieredCachingGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ArgoTieredCachingUpdateParams struct {
	// Enables Tiered Caching.
	Value param.Field[ArgoTieredCachingUpdateParamsValue] `json:"value,required"`
}

func (r ArgoTieredCachingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ArgoTieredCachingUpdateParamsValue string

const (
	ArgoTieredCachingUpdateParamsValueOn  ArgoTieredCachingUpdateParamsValue = "on"
	ArgoTieredCachingUpdateParamsValueOff ArgoTieredCachingUpdateParamsValue = "off"
)

type ArgoTieredCachingUpdateResponseEnvelope struct {
	Errors   []ArgoTieredCachingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoTieredCachingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoTieredCachingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoTieredCachingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoTieredCachingUpdateResponseEnvelopeJSON    `json:"-"`
}

// argoTieredCachingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoTieredCachingUpdateResponseEnvelope]
type argoTieredCachingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    argoTieredCachingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// argoTieredCachingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ArgoTieredCachingUpdateResponseEnvelopeErrors]
type argoTieredCachingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    argoTieredCachingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// argoTieredCachingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ArgoTieredCachingUpdateResponseEnvelopeMessages]
type argoTieredCachingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoTieredCachingUpdateResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingUpdateResponseEnvelopeSuccessTrue ArgoTieredCachingUpdateResponseEnvelopeSuccess = true
)

type ArgoTieredCachingGetResponseEnvelope struct {
	Errors   []ArgoTieredCachingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoTieredCachingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoTieredCachingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoTieredCachingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoTieredCachingGetResponseEnvelopeJSON    `json:"-"`
}

// argoTieredCachingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoTieredCachingGetResponseEnvelope]
type argoTieredCachingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    argoTieredCachingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// argoTieredCachingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ArgoTieredCachingGetResponseEnvelopeErrors]
type argoTieredCachingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    argoTieredCachingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// argoTieredCachingGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ArgoTieredCachingGetResponseEnvelopeMessages]
type argoTieredCachingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoTieredCachingGetResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingGetResponseEnvelopeSuccessTrue ArgoTieredCachingGetResponseEnvelopeSuccess = true
)
