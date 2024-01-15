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
func (r *UserLoadBalancerPoolReferenceService) LoadBalancerPoolsListPoolReferences(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s/references", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse struct {
	Errors   []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseError   `json:"errors"`
	Messages []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given pool.
	Result     []UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResult   `json:"result"`
	ResultInfo UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseSuccess `json:"success"`
	JSON    userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON    `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseErrorJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseError]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessageJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessage]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResult struct {
	ReferenceType UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                      `json:"resource_id"`
	ResourceName  string                                                                                      `json:"resource_name"`
	ResourceType  string                                                                                      `json:"resource_type"`
	JSON          userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultJSON          `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResult]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceType string

const (
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeStar     UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "*"
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferral UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referral"
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferrer UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referrer"
)

type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                `json:"total_count"`
	JSON       userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON `json:"-"`
}

// userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfo]
type userLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseSuccess bool

const (
	UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseSuccessTrue UserLoadBalancerPoolReferenceLoadBalancerPoolsListPoolReferencesResponseSuccess = true
)
