// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// LoadBalancerRegionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerRegionService] method
// instead.
type LoadBalancerRegionService struct {
	Options []option.RequestOption
}

// NewLoadBalancerRegionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerRegionService(opts ...option.RequestOption) (r *LoadBalancerRegionService) {
	r = &LoadBalancerRegionService{}
	r.Options = opts
	return
}

// List all region mappings.
func (r *LoadBalancerRegionService) List(ctx context.Context, params LoadBalancerRegionListParams, opts ...option.RequestOption) (res *LoadBalancerRegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerRegionListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single region mapping.
func (r *LoadBalancerRegionService) Get(ctx context.Context, regionID LoadBalancerRegionGetParamsRegionID, query LoadBalancerRegionGetParams, opts ...option.RequestOption) (res *LoadBalancerRegionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerRegionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", query.AccountID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [LoadBalancerRegionListResponseUnknown] or
// [shared.UnionString].
type LoadBalancerRegionListResponse interface {
	ImplementsLoadBalancerRegionListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerRegionListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// A list of countries and subdivisions mapped to a region.
//
// Union satisfied by [LoadBalancerRegionGetResponseUnknown] or
// [shared.UnionString].
type LoadBalancerRegionGetResponse interface {
	ImplementsLoadBalancerRegionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerRegionGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LoadBalancerRegionListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Two-letter alpha-2 country code followed in ISO 3166-1.
	CountryCodeA2 param.Field[string] `query:"country_code_a2"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCode param.Field[string] `query:"subdivision_code"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCodeA2 param.Field[string] `query:"subdivision_code_a2"`
}

// URLQuery serializes [LoadBalancerRegionListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerRegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerRegionListResponseEnvelope struct {
	Errors   []LoadBalancerRegionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerRegionListResponseEnvelopeMessages `json:"messages,required"`
	Result   LoadBalancerRegionListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerRegionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerRegionListResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerRegionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerRegionListResponseEnvelope]
type loadBalancerRegionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerRegionListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerRegionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerRegionListResponseEnvelopeErrors]
type loadBalancerRegionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerRegionListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerRegionListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerRegionListResponseEnvelopeMessages]
type loadBalancerRegionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerRegionListResponseEnvelopeSuccess bool

const (
	LoadBalancerRegionListResponseEnvelopeSuccessTrue LoadBalancerRegionListResponseEnvelopeSuccess = true
)

type LoadBalancerRegionGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

// A list of Cloudflare regions. WNAM: Western North America, ENAM: Eastern North
// America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America,
// SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa,
// SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East
// Asia).
type LoadBalancerRegionGetParamsRegionID string

const (
	LoadBalancerRegionGetParamsRegionIDWnam LoadBalancerRegionGetParamsRegionID = "WNAM"
	LoadBalancerRegionGetParamsRegionIDEnam LoadBalancerRegionGetParamsRegionID = "ENAM"
	LoadBalancerRegionGetParamsRegionIDWeu  LoadBalancerRegionGetParamsRegionID = "WEU"
	LoadBalancerRegionGetParamsRegionIDEeu  LoadBalancerRegionGetParamsRegionID = "EEU"
	LoadBalancerRegionGetParamsRegionIDNsam LoadBalancerRegionGetParamsRegionID = "NSAM"
	LoadBalancerRegionGetParamsRegionIDSsam LoadBalancerRegionGetParamsRegionID = "SSAM"
	LoadBalancerRegionGetParamsRegionIDOc   LoadBalancerRegionGetParamsRegionID = "OC"
	LoadBalancerRegionGetParamsRegionIDMe   LoadBalancerRegionGetParamsRegionID = "ME"
	LoadBalancerRegionGetParamsRegionIDNaf  LoadBalancerRegionGetParamsRegionID = "NAF"
	LoadBalancerRegionGetParamsRegionIDSaf  LoadBalancerRegionGetParamsRegionID = "SAF"
	LoadBalancerRegionGetParamsRegionIDSas  LoadBalancerRegionGetParamsRegionID = "SAS"
	LoadBalancerRegionGetParamsRegionIDSeas LoadBalancerRegionGetParamsRegionID = "SEAS"
	LoadBalancerRegionGetParamsRegionIDNeas LoadBalancerRegionGetParamsRegionID = "NEAS"
)

type LoadBalancerRegionGetResponseEnvelope struct {
	Errors   []LoadBalancerRegionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerRegionGetResponseEnvelopeMessages `json:"messages,required"`
	// A list of countries and subdivisions mapped to a region.
	Result LoadBalancerRegionGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerRegionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerRegionGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerRegionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerRegionGetResponseEnvelope]
type loadBalancerRegionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    loadBalancerRegionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerRegionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerRegionGetResponseEnvelopeErrors]
type loadBalancerRegionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    loadBalancerRegionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerRegionGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerRegionGetResponseEnvelopeMessages]
type loadBalancerRegionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerRegionGetResponseEnvelopeSuccess bool

const (
	LoadBalancerRegionGetResponseEnvelopeSuccessTrue LoadBalancerRegionGetResponseEnvelopeSuccess = true
)
