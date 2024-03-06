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
func (r *LoadBalancerPoolReferenceService) Get(ctx context.Context, poolID string, query LoadBalancerPoolReferenceGetParams, opts ...option.RequestOption) (res *[]LoadBalancerPoolReferenceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolReferenceGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerPoolReferenceGetResponse struct {
	ReferenceType LoadBalancerPoolReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                            `json:"resource_id"`
	ResourceName  string                                            `json:"resource_name"`
	ResourceType  string                                            `json:"resource_type"`
	JSON          loadBalancerPoolReferenceGetResponseJSON          `json:"-"`
}

// loadBalancerPoolReferenceGetResponseJSON contains the JSON metadata for the
// struct [LoadBalancerPoolReferenceGetResponse]
type loadBalancerPoolReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolReferenceGetResponseReferenceType string

const (
	LoadBalancerPoolReferenceGetResponseReferenceTypeStar     LoadBalancerPoolReferenceGetResponseReferenceType = "*"
	LoadBalancerPoolReferenceGetResponseReferenceTypeReferral LoadBalancerPoolReferenceGetResponseReferenceType = "referral"
	LoadBalancerPoolReferenceGetResponseReferenceTypeReferrer LoadBalancerPoolReferenceGetResponseReferenceType = "referrer"
)

type LoadBalancerPoolReferenceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerPoolReferenceGetResponseEnvelope struct {
	Errors   []LoadBalancerPoolReferenceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolReferenceGetResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []LoadBalancerPoolReferenceGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerPoolReferenceGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerPoolReferenceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerPoolReferenceGetResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerPoolReferenceGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancerPoolReferenceGetResponseEnvelope]
type loadBalancerPoolReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolReferenceGetResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    loadBalancerPoolReferenceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolReferenceGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferenceGetResponseEnvelopeErrors]
type loadBalancerPoolReferenceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferenceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerPoolReferenceGetResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    loadBalancerPoolReferenceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolReferenceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferenceGetResponseEnvelopeMessages]
type loadBalancerPoolReferenceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferenceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPoolReferenceGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolReferenceGetResponseEnvelopeSuccessTrue LoadBalancerPoolReferenceGetResponseEnvelopeSuccess = true
)

type LoadBalancerPoolReferenceGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       loadBalancerPoolReferenceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolReferenceGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferenceGetResponseEnvelopeResultInfo]
type loadBalancerPoolReferenceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerPoolReferenceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
