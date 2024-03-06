// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
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
func (r *LoadBalancerMonitorReferenceService) Get(ctx context.Context, monitorID string, query LoadBalancerMonitorReferenceGetParams, opts ...option.RequestOption) (res *[]LoadBalancerMonitorReferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerMonitorReferenceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", query.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerMonitorReferenceGetResponse struct {
	ReferenceType LoadBalancerMonitorReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                               `json:"resource_id"`
	ResourceName  string                                               `json:"resource_name"`
	ResourceType  string                                               `json:"resource_type"`
	JSON          loadBalancerMonitorReferenceGetResponseJSON          `json:"-"`
}

// loadBalancerMonitorReferenceGetResponseJSON contains the JSON metadata for the
// struct [LoadBalancerMonitorReferenceGetResponse]
type loadBalancerMonitorReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorReferenceGetResponseReferenceType string

const (
	LoadBalancerMonitorReferenceGetResponseReferenceTypeStar     LoadBalancerMonitorReferenceGetResponseReferenceType = "*"
	LoadBalancerMonitorReferenceGetResponseReferenceTypeReferral LoadBalancerMonitorReferenceGetResponseReferenceType = "referral"
	LoadBalancerMonitorReferenceGetResponseReferenceTypeReferrer LoadBalancerMonitorReferenceGetResponseReferenceType = "referrer"
)

type LoadBalancerMonitorReferenceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerMonitorReferenceGetResponseEnvelope struct {
	Errors   []LoadBalancerMonitorReferenceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerMonitorReferenceGetResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []LoadBalancerMonitorReferenceGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerMonitorReferenceGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerMonitorReferenceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerMonitorReferenceGetResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerMonitorReferenceGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [LoadBalancerMonitorReferenceGetResponseEnvelope]
type loadBalancerMonitorReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorReferenceGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    loadBalancerMonitorReferenceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerMonitorReferenceGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [LoadBalancerMonitorReferenceGetResponseEnvelopeErrors]
type loadBalancerMonitorReferenceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferenceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerMonitorReferenceGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    loadBalancerMonitorReferenceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerMonitorReferenceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [LoadBalancerMonitorReferenceGetResponseEnvelopeMessages]
type loadBalancerMonitorReferenceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferenceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerMonitorReferenceGetResponseEnvelopeSuccess bool

const (
	LoadBalancerMonitorReferenceGetResponseEnvelopeSuccessTrue LoadBalancerMonitorReferenceGetResponseEnvelopeSuccess = true
)

type LoadBalancerMonitorReferenceGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       loadBalancerMonitorReferenceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerMonitorReferenceGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancerMonitorReferenceGetResponseEnvelopeResultInfo]
type loadBalancerMonitorReferenceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMonitorReferenceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerMonitorReferenceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
