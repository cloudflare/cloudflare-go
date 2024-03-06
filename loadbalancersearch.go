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

// LoadBalancerSearchService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerSearchService] method
// instead.
type LoadBalancerSearchService struct {
	Options []option.RequestOption
}

// NewLoadBalancerSearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerSearchService(opts ...option.RequestOption) (r *LoadBalancerSearchService) {
	r = &LoadBalancerSearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *LoadBalancerSearchService) Get(ctx context.Context, params LoadBalancerSearchGetParams, opts ...option.RequestOption) (res *[]LoadBalancerSearchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerSearchGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerSearchGetResponse = interface{}

type LoadBalancerSearchGetParams struct {
	// Identifier
	AccountID    param.Field[string]                                  `path:"account_id,required"`
	Page         param.Field[interface{}]                             `query:"page"`
	PerPage      param.Field[interface{}]                             `query:"per_page"`
	SearchParams param.Field[LoadBalancerSearchGetParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [LoadBalancerSearchGetParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerSearchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerSearchGetParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[LoadBalancerSearchGetParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [LoadBalancerSearchGetParamsSearchParams]'s query parameters
// as `url.Values`.
func (r LoadBalancerSearchGetParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of references to include ("\*" for all).
type LoadBalancerSearchGetParamsSearchParamsReferences string

const (
	LoadBalancerSearchGetParamsSearchParamsReferencesEmpty    LoadBalancerSearchGetParamsSearchParamsReferences = ""
	LoadBalancerSearchGetParamsSearchParamsReferencesStar     LoadBalancerSearchGetParamsSearchParamsReferences = "*"
	LoadBalancerSearchGetParamsSearchParamsReferencesReferral LoadBalancerSearchGetParamsSearchParamsReferences = "referral"
	LoadBalancerSearchGetParamsSearchParamsReferencesReferrer LoadBalancerSearchGetParamsSearchParamsReferences = "referrer"
)

type LoadBalancerSearchGetResponseEnvelope struct {
	Errors   []LoadBalancerSearchGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerSearchGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancerSearchGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerSearchGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerSearchGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerSearchGetResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerSearchGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerSearchGetResponseEnvelope]
type loadBalancerSearchGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSearchGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerSearchGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    loadBalancerSearchGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerSearchGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerSearchGetResponseEnvelopeErrors]
type loadBalancerSearchGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSearchGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerSearchGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    loadBalancerSearchGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerSearchGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [LoadBalancerSearchGetResponseEnvelopeMessages]
type loadBalancerSearchGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSearchGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerSearchGetResponseEnvelopeSuccess bool

const (
	LoadBalancerSearchGetResponseEnvelopeSuccessTrue LoadBalancerSearchGetResponseEnvelopeSuccess = true
)

type LoadBalancerSearchGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       loadBalancerSearchGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerSearchGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [LoadBalancerSearchGetResponseEnvelopeResultInfo]
type loadBalancerSearchGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSearchGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSearchGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
