// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SearchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *SearchService) Get(ctx context.Context, params SearchGetParams, opts ...option.RequestOption) (res *SearchGetResponse, err error) {
	var env SearchGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SearchGetResponse struct {
	// A list of resources matching the search query.
	Resources []SearchGetResponseResource `json:"resources"`
	JSON      searchGetResponseJSON       `json:"-"`
}

// searchGetResponseJSON contains the JSON metadata for the struct
// [SearchGetResponse]
type searchGetResponseJSON struct {
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a load balancer resource.
type SearchGetResponseResource struct {
	// When listed as a reference, the type (direction) of the reference.
	ReferenceType SearchGetResponseResourcesReferenceType `json:"reference_type"`
	// A list of references to (referrer) or from (referral) this resource.
	References []interface{} `json:"references"`
	ResourceID string        `json:"resource_id"`
	// The human-identifiable name of the resource.
	ResourceName string `json:"resource_name"`
	// The type of the resource.
	ResourceType SearchGetResponseResourcesResourceType `json:"resource_type"`
	JSON         searchGetResponseResourceJSON          `json:"-"`
}

// searchGetResponseResourceJSON contains the JSON metadata for the struct
// [SearchGetResponseResource]
type searchGetResponseResourceJSON struct {
	ReferenceType apijson.Field
	References    apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchGetResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseResourceJSON) RawJSON() string {
	return r.raw
}

// When listed as a reference, the type (direction) of the reference.
type SearchGetResponseResourcesReferenceType string

const (
	SearchGetResponseResourcesReferenceTypeReferral SearchGetResponseResourcesReferenceType = "referral"
	SearchGetResponseResourcesReferenceTypeReferrer SearchGetResponseResourcesReferenceType = "referrer"
)

func (r SearchGetResponseResourcesReferenceType) IsKnown() bool {
	switch r {
	case SearchGetResponseResourcesReferenceTypeReferral, SearchGetResponseResourcesReferenceTypeReferrer:
		return true
	}
	return false
}

// The type of the resource.
type SearchGetResponseResourcesResourceType string

const (
	SearchGetResponseResourcesResourceTypeLoadBalancer SearchGetResponseResourcesResourceType = "load_balancer"
	SearchGetResponseResourcesResourceTypeMonitor      SearchGetResponseResourcesResourceType = "monitor"
	SearchGetResponseResourcesResourceTypePool         SearchGetResponseResourcesResourceType = "pool"
)

func (r SearchGetResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case SearchGetResponseResourcesResourceTypeLoadBalancer, SearchGetResponseResourcesResourceTypeMonitor, SearchGetResponseResourcesResourceTypePool:
		return true
	}
	return false
}

type SearchGetParams struct {
	// Identifier
	AccountID    param.Field[string]                      `path:"account_id,required"`
	Page         param.Field[float64]                     `query:"page"`
	PerPage      param.Field[float64]                     `query:"per_page"`
	SearchParams param.Field[SearchGetParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [SearchGetParams]'s query parameters as `url.Values`.
func (r SearchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SearchGetParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[SearchGetParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [SearchGetParamsSearchParams]'s query parameters as
// `url.Values`.
func (r SearchGetParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of references to include ("\*" for all).
type SearchGetParamsSearchParamsReferences string

const (
	SearchGetParamsSearchParamsReferencesEmpty    SearchGetParamsSearchParamsReferences = ""
	SearchGetParamsSearchParamsReferencesStar     SearchGetParamsSearchParamsReferences = "*"
	SearchGetParamsSearchParamsReferencesReferral SearchGetParamsSearchParamsReferences = "referral"
	SearchGetParamsSearchParamsReferencesReferrer SearchGetParamsSearchParamsReferences = "referrer"
)

func (r SearchGetParamsSearchParamsReferences) IsKnown() bool {
	switch r {
	case SearchGetParamsSearchParamsReferencesEmpty, SearchGetParamsSearchParamsReferencesStar, SearchGetParamsSearchParamsReferencesReferral, SearchGetParamsSearchParamsReferencesReferrer:
		return true
	}
	return false
}

type SearchGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   SearchGetResponse     `json:"result,required"`
	// Whether the API call was successful
	Success    SearchGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SearchGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       searchGetResponseEnvelopeJSON       `json:"-"`
}

// searchGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SearchGetResponseEnvelope]
type searchGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SearchGetResponseEnvelopeSuccess bool

const (
	SearchGetResponseEnvelopeSuccessTrue SearchGetResponseEnvelopeSuccess = true
)

func (r SearchGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SearchGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SearchGetResponseEnvelopeResultInfo struct {
	// Total number of results on the current page
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total number of pages available
	TotalPages float64                                 `json:"total_pages"`
	JSON       searchGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// searchGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [SearchGetResponseEnvelopeResultInfo]
type searchGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
