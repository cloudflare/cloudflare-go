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

// UserLoadBalancerPoolReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerPoolReferenceService] method instead.
type UserLoadBalancerPoolReferenceService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolReferenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolReferenceService(opts ...option.RequestOption) (r *UserLoadBalancerPoolReferenceService) {
	r = &UserLoadBalancerPoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *UserLoadBalancerPoolReferenceService) LoadBalancerPoolsListPoolReferences(ctx context.Context, poolID string, opts ...option.RequestOption) (res *[]UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/references", poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse struct {
	ReferenceType UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceType `json:"reference_type"`
	ResourceID    string                                                                                `json:"resource_id"`
	ResourceName  string                                                                                `json:"resource_name"`
	ResourceType  string                                                                                `json:"resource_type"`
	JSON          userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON          `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceType string

const (
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceTypeStar     UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceType = "*"
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceTypeReferral UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceType = "referral"
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceTypeReferrer UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseReferenceType = "referrer"
)

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessages `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelope]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrors struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrors]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessages struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessages]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeSuccessTrue UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeSuccess = true
)

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                        `json:"total_count"`
	JSON       userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfo]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
