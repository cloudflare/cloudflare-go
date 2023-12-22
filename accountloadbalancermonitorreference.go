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

// AccountLoadBalancerMonitorReferenceService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLoadBalancerMonitorReferenceService] method instead.
type AccountLoadBalancerMonitorReferenceService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerMonitorReferenceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountLoadBalancerMonitorReferenceService(opts ...option.RequestOption) (r *AccountLoadBalancerMonitorReferenceService) {
	r = &AccountLoadBalancerMonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *AccountLoadBalancerMonitorReferenceService) AccountLoadBalancerMonitorsListMonitorReferences(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *ReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%v/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ReferencesResponse struct {
	Errors   []ReferencesResponseError   `json:"errors"`
	Messages []ReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given monitor.
	Result     []ReferencesResponseResult   `json:"result"`
	ResultInfo ReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ReferencesResponseSuccess `json:"success"`
	JSON    referencesResponseJSON    `json:"-"`
}

// referencesResponseJSON contains the JSON metadata for the struct
// [ReferencesResponse]
type referencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReferencesResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    referencesResponseErrorJSON `json:"-"`
}

// referencesResponseErrorJSON contains the JSON metadata for the struct
// [ReferencesResponseError]
type referencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReferencesResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    referencesResponseMessageJSON `json:"-"`
}

// referencesResponseMessageJSON contains the JSON metadata for the struct
// [ReferencesResponseMessage]
type referencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReferencesResponseResult struct {
	ReferenceType ReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                `json:"resource_id"`
	ResourceName  string                                `json:"resource_name"`
	ResourceType  string                                `json:"resource_type"`
	JSON          referencesResponseResultJSON          `json:"-"`
}

// referencesResponseResultJSON contains the JSON metadata for the struct
// [ReferencesResponseResult]
type referencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ReferencesResponseResultReferenceType string

const (
	ReferencesResponseResultReferenceTypeStar     ReferencesResponseResultReferenceType = "*"
	ReferencesResponseResultReferenceTypeReferral ReferencesResponseResultReferenceType = "referral"
	ReferencesResponseResultReferenceTypeReferrer ReferencesResponseResultReferenceType = "referrer"
)

type ReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                          `json:"total_count"`
	JSON       referencesResponseResultInfoJSON `json:"-"`
}

// referencesResponseResultInfoJSON contains the JSON metadata for the struct
// [ReferencesResponseResultInfo]
type referencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ReferencesResponseSuccess bool

const (
	ReferencesResponseSuccessTrue ReferencesResponseSuccess = true
)
