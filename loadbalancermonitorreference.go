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
func (r *LoadBalancerMonitorReferenceService) List(ctx context.Context, accountID string, monitorID string, opts ...option.RequestOption) (res *[]LoadBalancerMonitorReferenceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorReferenceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerMonitorReferenceListResponse struct {
	ReferenceType LoadBalancerMonitorReferenceListResponseReferenceType `json:"reference_type"`
	ResourceID    string                                                `json:"resource_id"`
	ResourceName  string                                                `json:"resource_name"`
	ResourceType  string                                                `json:"resource_type"`
	JSON          loadBalancerMonitorReferenceListResponseJSON          `json:"-"`
}

// loadBalancerMonitorReferenceListResponseJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorReferenceListResponse]
type loadBalancerMonitorReferenceListResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceListResponseReferenceType string

const (
	LoadBalancerMonitorReferenceListResponseReferenceTypeStar     LoadBalancerMonitorReferenceListResponseReferenceType = "*"
	LoadBalancerMonitorReferenceListResponseReferenceTypeReferral LoadBalancerMonitorReferenceListResponseReferenceType = "referral"
	LoadBalancerMonitorReferenceListResponseReferenceTypeReferrer LoadBalancerMonitorReferenceListResponseReferenceType = "referrer"
)

type LoadBalancerMonitorReferenceListResponseEnvelope struct {
	Errors   []LoadBalancerMonitorReferenceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorReferenceListResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []LoadBalancerMonitorReferenceListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerMonitorReferenceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerMonitorReferenceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerMonitorReferenceListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerMonitorReferenceListResponseEnvelopeJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorReferenceListResponseEnvelope]
type loadBalancerMonitorReferenceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    loadBalancerMonitorReferenceListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorReferenceListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [LoadBalancerMonitorReferenceListResponseEnvelopeErrors]
type loadBalancerMonitorReferenceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMonitorReferenceListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    loadBalancerMonitorReferenceListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorReferenceListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [LoadBalancerMonitorReferenceListResponseEnvelopeMessages]
type loadBalancerMonitorReferenceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerMonitorReferenceListResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorReferenceListResponseEnvelopeSuccessTrue LoadBalancerMonitorReferenceListResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorReferenceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       loadBalancerMonitorReferenceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorReferenceListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancerMonitorReferenceListResponseEnvelopeResultInfo]
type loadBalancerMonitorReferenceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
