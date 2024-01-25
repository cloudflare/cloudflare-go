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

// AccountLoadBalancerPoolReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPoolReferenceService] method instead.
type AccountLoadBalancerPoolReferenceService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPoolReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPoolReferenceService(opts ...option.RequestOption) (r *AccountLoadBalancerPoolReferenceService) {
	r = &AccountLoadBalancerPoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *AccountLoadBalancerPoolReferenceService) AccountLoadBalancerPoolsListPoolReferences(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse struct {
	Errors   []AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError   `json:"errors"`
	Messages []AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given pool.
	Result     []AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult   `json:"result"`
	ResultInfo AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess `json:"success"`
	JSON    accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON    `json:"-"`
}

// accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse]
type accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON `json:"-"`
}

// accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError]
type accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON `json:"-"`
}

// accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage]
type accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult struct {
	ReferenceType AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                                `json:"resource_id"`
	ResourceName  string                                                                                                `json:"resource_name"`
	ResourceType  string                                                                                                `json:"resource_type"`
	JSON          accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON          `json:"-"`
}

// accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult]
type accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType string

const (
	AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeStar     AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "*"
	AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferral AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referral"
	AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceTypeReferrer AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultReferenceType = "referrer"
)

type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                          `json:"total_count"`
	JSON       accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo]
type accountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess bool

const (
	AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccessTrue AccountLoadBalancerPoolReferenceAccountLoadBalancerPoolsListPoolReferencesResponseSuccess = true
)
