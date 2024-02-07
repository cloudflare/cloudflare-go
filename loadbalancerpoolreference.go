// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerPoolReferenceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLoadBalancerPoolReferenceService] method instead.
type LoadBalancerPoolReferenceService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolReferenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolReferenceService(opts ...option.RequestOption) (r *LoadBalancerPoolReferenceService) {
	r = &LoadBalancerPoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *LoadBalancerPoolReferenceService) AccountLoadBalancerPoolsListPoolReferences(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse struct {
	Errors   []LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError   `json:"errors"`
	Messages []LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given pool.
	Result     []LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult   `json:"result"`
	ResultInfo LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess `json:"success"`
	JSON    loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON    `json:"-"`
}

// loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse]
type loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON `json:"-"`
}

// loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError]
type loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON `json:"-"`
}

// loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage]
type loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult struct {
	ReferenceType LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                         `json:"resource_id"`
	ResourceName  string                                                                                         `json:"resource_name"`
	ResourceType  string                                                                                         `json:"resource_type"`
	JSON          loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON          `json:"-"`
}

// loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult]
type loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType string

const (
	LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeStar     LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "*"
	LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferral LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referral"
	LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferrer LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referrer"
)

type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                   `json:"total_count"`
	JSON       loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON `json:"-"`
}

// loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo]
type loadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess bool

const (
	LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccessTrue LoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess = true
)
