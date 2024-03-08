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
func (r *ArgoTieredCachingService) Edit(ctx context.Context, params ArgoTieredCachingEditParams, opts ...option.RequestOption) (res *ArgoTieredCachingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Tiered Caching setting
func (r *ArgoTieredCachingService) Get(ctx context.Context, query ArgoTieredCachingGetParams, opts ...option.RequestOption) (res *ArgoTieredCachingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ArgoTieredCachingGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [ArgoTieredCachingEditResponseUnknown] or
// [shared.UnionString].
type ArgoTieredCachingEditResponse interface {
	ImplementsArgoTieredCachingEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ArgoTieredCachingEditResponse)(nil)).Elem(),
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

type ArgoTieredCachingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Caching.
	Value param.Field[ArgoTieredCachingEditParamsValue] `json:"value,required"`
}

func (r ArgoTieredCachingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type ArgoTieredCachingEditParamsValue string

const (
	ArgoTieredCachingEditParamsValueOn  ArgoTieredCachingEditParamsValue = "on"
	ArgoTieredCachingEditParamsValueOff ArgoTieredCachingEditParamsValue = "off"
)

type ArgoTieredCachingEditResponseEnvelope struct {
	Errors   []ArgoTieredCachingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ArgoTieredCachingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ArgoTieredCachingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ArgoTieredCachingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    argoTieredCachingEditResponseEnvelopeJSON    `json:"-"`
}

// argoTieredCachingEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ArgoTieredCachingEditResponseEnvelope]
type argoTieredCachingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoTieredCachingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ArgoTieredCachingEditResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    argoTieredCachingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// argoTieredCachingEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ArgoTieredCachingEditResponseEnvelopeErrors]
type argoTieredCachingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoTieredCachingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ArgoTieredCachingEditResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    argoTieredCachingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// argoTieredCachingEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ArgoTieredCachingEditResponseEnvelopeMessages]
type argoTieredCachingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ArgoTieredCachingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r argoTieredCachingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ArgoTieredCachingEditResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingEditResponseEnvelopeSuccessTrue ArgoTieredCachingEditResponseEnvelopeSuccess = true
)

type ArgoTieredCachingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

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

func (r argoTieredCachingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r argoTieredCachingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r argoTieredCachingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ArgoTieredCachingGetResponseEnvelopeSuccess bool

const (
	ArgoTieredCachingGetResponseEnvelopeSuccessTrue ArgoTieredCachingGetResponseEnvelopeSuccess = true
)
