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

// AccountLoadBalancerSearchService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerSearchService] method instead.
type AccountLoadBalancerSearchService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerSearchService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerSearchService(opts ...option.RequestOption) (r *AccountLoadBalancerSearchService) {
	r = &AccountLoadBalancerSearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *AccountLoadBalancerSearchService) List(ctx context.Context, accountIdentifier string, query AccountLoadBalancerSearchListParams, opts ...option.RequestOption) (res *AccountLoadBalancerSearchListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/search", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountLoadBalancerSearchListResponse struct {
	Errors     []AccountLoadBalancerSearchListResponseError    `json:"errors"`
	Messages   []AccountLoadBalancerSearchListResponseMessage  `json:"messages"`
	Result     AccountLoadBalancerSearchListResponseResult     `json:"result"`
	ResultInfo AccountLoadBalancerSearchListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerSearchListResponseSuccess `json:"success"`
	JSON    accountLoadBalancerSearchListResponseJSON    `json:"-"`
}

// accountLoadBalancerSearchListResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerSearchListResponse]
type accountLoadBalancerSearchListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerSearchListResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountLoadBalancerSearchListResponseErrorJSON `json:"-"`
}

// accountLoadBalancerSearchListResponseErrorJSON contains the JSON metadata for
// the struct [AccountLoadBalancerSearchListResponseError]
type accountLoadBalancerSearchListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerSearchListResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountLoadBalancerSearchListResponseMessageJSON `json:"-"`
}

// accountLoadBalancerSearchListResponseMessageJSON contains the JSON metadata for
// the struct [AccountLoadBalancerSearchListResponseMessage]
type accountLoadBalancerSearchListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerSearchListResponseResult struct {
	// A list of resources matching the search query.
	Resources []AccountLoadBalancerSearchListResponseResultResource `json:"resources"`
	JSON      accountLoadBalancerSearchListResponseResultJSON       `json:"-"`
}

// accountLoadBalancerSearchListResponseResultJSON contains the JSON metadata for
// the struct [AccountLoadBalancerSearchListResponseResult]
type accountLoadBalancerSearchListResponseResultJSON struct {
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a load balancer resource.
type AccountLoadBalancerSearchListResponseResultResource struct {
	// When listed as a reference, the type (direction) of the reference.
	ReferenceType AccountLoadBalancerSearchListResponseResultResourcesReferenceType `json:"reference_type"`
	// A list of references to (referrer) or from (referral) this resource.
	References []interface{} `json:"references"`
	ResourceID interface{}   `json:"resource_id"`
	// The human-identifiable name of the resource.
	ResourceName string `json:"resource_name"`
	// The type of the resource.
	ResourceType AccountLoadBalancerSearchListResponseResultResourcesResourceType `json:"resource_type"`
	JSON         accountLoadBalancerSearchListResponseResultResourceJSON          `json:"-"`
}

// accountLoadBalancerSearchListResponseResultResourceJSON contains the JSON
// metadata for the struct [AccountLoadBalancerSearchListResponseResultResource]
type accountLoadBalancerSearchListResponseResultResourceJSON struct {
	ReferenceType apijson.Field
	References    apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponseResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// When listed as a reference, the type (direction) of the reference.
type AccountLoadBalancerSearchListResponseResultResourcesReferenceType string

const (
	AccountLoadBalancerSearchListResponseResultResourcesReferenceTypeReferral AccountLoadBalancerSearchListResponseResultResourcesReferenceType = "referral"
	AccountLoadBalancerSearchListResponseResultResourcesReferenceTypeReferrer AccountLoadBalancerSearchListResponseResultResourcesReferenceType = "referrer"
)

// The type of the resource.
type AccountLoadBalancerSearchListResponseResultResourcesResourceType string

const (
	AccountLoadBalancerSearchListResponseResultResourcesResourceTypeLoadBalancer AccountLoadBalancerSearchListResponseResultResourcesResourceType = "load_balancer"
	AccountLoadBalancerSearchListResponseResultResourcesResourceTypeMonitor      AccountLoadBalancerSearchListResponseResultResourcesResourceType = "monitor"
	AccountLoadBalancerSearchListResponseResultResourcesResourceTypePool         AccountLoadBalancerSearchListResponseResultResourcesResourceType = "pool"
)

type AccountLoadBalancerSearchListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accountLoadBalancerSearchListResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerSearchListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountLoadBalancerSearchListResponseResultInfo]
type accountLoadBalancerSearchListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerSearchListResponseSuccess bool

const (
	AccountLoadBalancerSearchListResponseSuccessTrue AccountLoadBalancerSearchListResponseSuccess = true
)

type AccountLoadBalancerSearchListParams struct {
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage      param.Field[float64]                                         `query:"per_page"`
	SearchParams param.Field[AccountLoadBalancerSearchListParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [AccountLoadBalancerSearchListParams]'s query parameters as
// `url.Values`.
func (r AccountLoadBalancerSearchListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLoadBalancerSearchListParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[AccountLoadBalancerSearchListParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [AccountLoadBalancerSearchListParamsSearchParams]'s query
// parameters as `url.Values`.
func (r AccountLoadBalancerSearchListParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of references to include ("\*" for all).
type AccountLoadBalancerSearchListParamsSearchParamsReferences string

const (
	AccountLoadBalancerSearchListParamsSearchParamsReferencesEmpty    AccountLoadBalancerSearchListParamsSearchParamsReferences = ""
	AccountLoadBalancerSearchListParamsSearchParamsReferencesStar     AccountLoadBalancerSearchListParamsSearchParamsReferences = "*"
	AccountLoadBalancerSearchListParamsSearchParamsReferencesReferral AccountLoadBalancerSearchListParamsSearchParamsReferences = "referral"
	AccountLoadBalancerSearchListParamsSearchParamsReferencesReferrer AccountLoadBalancerSearchListParamsSearchParamsReferences = "referrer"
)
