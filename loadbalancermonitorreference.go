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

// LoadBalancerMonitorReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewLoadBalancerMonitorReferenceService] method instead.
type LoadBalancerMonitorReferenceService struct {
	Options []option.RequestOption
}

// NewLoadBalancerMonitorReferenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLoadBalancerMonitorReferenceService(opts ...option.RequestOption) (r *LoadBalancerMonitorReferenceService) {
	r = &LoadBalancerMonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *LoadBalancerMonitorReferenceService) AccountLoadBalancerMonitorsListMonitorReferences(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse struct {
	Errors   []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError   `json:"errors"`
	Messages []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given monitor.
	Result     []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult   `json:"result"`
	ResultInfo LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess `json:"success"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON    `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult struct {
	ReferenceType LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                                  `json:"resource_id"`
	ResourceName  string                                                                                                  `json:"resource_name"`
	ResourceType  string                                                                                                  `json:"resource_type"`
	JSON          loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON          `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType string

const (
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeStar     LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "*"
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferral LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referral"
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferrer LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referrer"
)

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                            `json:"total_count"`
	JSON       loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess bool

const (
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccessTrue LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess = true
)
