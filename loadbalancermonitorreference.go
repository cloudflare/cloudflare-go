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
func (r *LoadBalancerMonitorReferenceService) AccountLoadBalancerMonitorsListMonitorReferences(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *[]LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse struct {
	ReferenceType LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                                                                            `json:"resource_id"`
	ResourceName  string                                                                                            `json:"resource_name"`
	ResourceType  string                                                                                            `json:"resource_type"`
	JSON          loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON          `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceType string

const (
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeStar     LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "*"
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeReferral LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "referral"
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceTypeReferrer LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseReferenceType = "referrer"
)

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelope struct {
	Errors   []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors   `json:"errors"`
	Messages []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages `json:"messages"`
	// List of resources that reference a given monitor.
	Result     []LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse                 `json:"result"`
	ResultInfo LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess `json:"success"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelope]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                    `json:"total_count"`
	JSON       loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo]
type loadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccessTrue LoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseEnvelopeSuccess = true
)
