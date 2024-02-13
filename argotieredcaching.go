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

// Get Tiered Caching setting
func (r *ArgoTieredCachingService) TieredCachingGetTieredCachingSetting(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ArgoTieredCachingTieredCachingGetTieredCachingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates enablement of Tiered Caching
func (r *ArgoTieredCachingService) TieredCachingPatchTieredCachingSetting(ctx context.Context, zoneID string, body ArgoTieredCachingTieredCachingPatchTieredCachingSettingParams, opts ...option.RequestOption) (res *ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseUnknown] or
// [shared.UnionString].
type ArgoTieredCachingTieredCachingGetTieredCachingSettingResponse interface {
	ImplementsArgoTieredCachingTieredCachingGetTieredCachingSettingResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoTieredCachingTieredCachingGetTieredCachingSettingResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseUnknown] or
// [shared.UnionString].
type ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse interface {
	ImplementsArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelope struct {
	Errors   []ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoTieredCachingTieredCachingGetTieredCachingSettingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeJSON    `json:"-"`
}

// argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelope]
type argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrors]
type argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessages]
type argoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeSuccessTrue ArgoTieredCachingTieredCachingGetTieredCachingSettingResponseEnvelopeSuccess = true
)

type ArgoTieredCachingTieredCachingPatchTieredCachingSettingParams struct {
	// Enables Tiered Caching.
	Value param.Field[ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue] `json:"value,required"`
}

func (r ArgoTieredCachingTieredCachingPatchTieredCachingSettingParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue string

const (
	ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOn  ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "on"
	ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValueOff ArgoTieredCachingTieredCachingPatchTieredCachingSettingParamsValue = "off"
)

type ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelope struct {
	Errors   []ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeJSON    `json:"-"`
}

// argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelope]
type argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrorsJSON `json:"-"`
}

// argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrors]
type argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessagesJSON `json:"-"`
}

// argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessages]
type argoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeSuccessTrue ArgoTieredCachingTieredCachingPatchTieredCachingSettingResponseEnvelopeSuccess = true
)
