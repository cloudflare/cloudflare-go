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
func (r *LoadBalancerPoolReferenceService) List(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *[]LoadBalancerPoolReferenceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPoolReferenceListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerPoolReferenceListResponse struct {
	ReferenceType LoadBalancerPoolReferenceListResponseReferenceType `json:"reference_type"`
	ResourceID    string                                             `json:"resource_id"`
	ResourceName  string                                             `json:"resource_name"`
	ResourceType  string                                             `json:"resource_type"`
	JSON          loadBalancerPoolReferenceListResponseJSON          `json:"-"`
}

// loadBalancerPoolReferenceListResponseJSON contains the JSON metadata for the
// struct [LoadBalancerPoolReferenceListResponse]
type loadBalancerPoolReferenceListResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceListResponseReferenceType string

const (
	LoadBalancerPoolReferenceListResponseReferenceTypeStar     LoadBalancerPoolReferenceListResponseReferenceType = "*"
	LoadBalancerPoolReferenceListResponseReferenceTypeReferral LoadBalancerPoolReferenceListResponseReferenceType = "referral"
	LoadBalancerPoolReferenceListResponseReferenceTypeReferrer LoadBalancerPoolReferenceListResponseReferenceType = "referrer"
)

type LoadBalancerPoolReferenceListResponseEnvelope struct {
	Errors   []LoadBalancerPoolReferenceListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolReferenceListResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []LoadBalancerPoolReferenceListResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerPoolReferenceListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerPoolReferenceListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerPoolReferenceListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerPoolReferenceListResponseEnvelopeJSON contains the JSON metadata for
// the struct [LoadBalancerPoolReferenceListResponseEnvelope]
type loadBalancerPoolReferenceListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceListResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    loadBalancerPoolReferenceListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolReferenceListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferenceListResponseEnvelopeErrors]
type loadBalancerPoolReferenceListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolReferenceListResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    loadBalancerPoolReferenceListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolReferenceListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LoadBalancerPoolReferenceListResponseEnvelopeMessages]
type loadBalancerPoolReferenceListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolReferenceListResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolReferenceListResponseEnvelopeSuccessTrue LoadBalancerPoolReferenceListResponseEnvelopeSuccess = true
)

type LoadBalancerPoolReferenceListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       loadBalancerPoolReferenceListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerPoolReferenceListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancerPoolReferenceListResponseEnvelopeResultInfo]
type loadBalancerPoolReferenceListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolReferenceListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
