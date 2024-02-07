// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

// Get a single region mapping.
func (r *LoadBalancerRegionService) Get(ctx context.Context, accountIdentifier string, regionCode LoadBalancerRegionGetParamsRegionCode, opts ...option.RequestOption) (res *LoadBalancerRegionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerRegionGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", accountIdentifier, regionCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all region mappings.
func (r *LoadBalancerRegionService) LoadBalancerRegionsListRegions(ctx context.Context, accountIdentifier string, query LoadBalancerRegionLoadBalancerRegionsListRegionsParams, opts ...option.RequestOption) (res *LoadBalancerRegionLoadBalancerRegionsListRegionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerRegionGetResponse = interface{}

type LoadBalancerRegionLoadBalancerRegionsListRegionsResponse = interface{}

// A list of Cloudflare regions. WNAM: Western North America, ENAM: Eastern North
// America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America,
// SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa,
// SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East
// Asia).
type LoadBalancerRegionGetParamsRegionCode string

const (
	LoadBalancerRegionGetParamsRegionCodeWnam LoadBalancerRegionGetParamsRegionCode = "WNAM"
	LoadBalancerRegionGetParamsRegionCodeEnam LoadBalancerRegionGetParamsRegionCode = "ENAM"
	LoadBalancerRegionGetParamsRegionCodeWeu  LoadBalancerRegionGetParamsRegionCode = "WEU"
	LoadBalancerRegionGetParamsRegionCodeEeu  LoadBalancerRegionGetParamsRegionCode = "EEU"
	LoadBalancerRegionGetParamsRegionCodeNsam LoadBalancerRegionGetParamsRegionCode = "NSAM"
	LoadBalancerRegionGetParamsRegionCodeSsam LoadBalancerRegionGetParamsRegionCode = "SSAM"
	LoadBalancerRegionGetParamsRegionCodeOc   LoadBalancerRegionGetParamsRegionCode = "OC"
	LoadBalancerRegionGetParamsRegionCodeMe   LoadBalancerRegionGetParamsRegionCode = "ME"
	LoadBalancerRegionGetParamsRegionCodeNaf  LoadBalancerRegionGetParamsRegionCode = "NAF"
	LoadBalancerRegionGetParamsRegionCodeSaf  LoadBalancerRegionGetParamsRegionCode = "SAF"
	LoadBalancerRegionGetParamsRegionCodeSas  LoadBalancerRegionGetParamsRegionCode = "SAS"
	LoadBalancerRegionGetParamsRegionCodeSeas LoadBalancerRegionGetParamsRegionCode = "SEAS"
	LoadBalancerRegionGetParamsRegionCodeNeas LoadBalancerRegionGetParamsRegionCode = "NEAS"
)

type LoadBalancerRegionGetResponseEnvelope struct {
	Errors   []LoadBalancerRegionGetResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerRegionGetResponseEnvelopeMessages `json:"messages"`
	// A list of countries and subdivisions mapped to a region.
	Result LoadBalancerRegionGetResponse `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerRegionGetResponseEnvelopeSuccess `json:"success"`
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

type LoadBalancerRegionLoadBalancerRegionsListRegionsParams struct {
	// Two-letter alpha-2 country code followed in ISO 3166-1.
	CountryCodeA2 param.Field[string] `query:"country_code_a2"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCode param.Field[string] `query:"subdivision_code"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCodeA2 param.Field[string] `query:"subdivision_code_a2"`
}

// URLQuery serializes [LoadBalancerRegionLoadBalancerRegionsListRegionsParams]'s
// query parameters as `url.Values`.
func (r LoadBalancerRegionLoadBalancerRegionsListRegionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelope struct {
	Errors   []LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessages `json:"messages"`
	Result   LoadBalancerRegionLoadBalancerRegionsListRegionsResponse                   `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelope]
type loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrors]
type loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessages]
type loadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeSuccess bool

const (
	LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeSuccessTrue LoadBalancerRegionLoadBalancerRegionsListRegionsResponseEnvelopeSuccess = true
)
