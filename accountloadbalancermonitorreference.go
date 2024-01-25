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
func (r *AccountLoadBalancerMonitorReferenceService) AccountLoadBalancerMonitorsListMonitorReferences(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse struct {
	Errors   []AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError   `json:"errors"`
	Messages []AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage `json:"messages"`
	// List of resources that reference a given monitor.
	Result     []AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult   `json:"result"`
	ResultInfo AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess `json:"success"`
	JSON    accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON    `json:"-"`
}

// accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse]
type accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON `json:"-"`
}

// accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError]
type accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON `json:"-"`
}

// accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage]
type accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult struct {
	ReferenceType AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                                                                                         `json:"resource_id"`
	ResourceName  string                                                                                                         `json:"resource_name"`
	ResourceType  string                                                                                                         `json:"resource_type"`
	JSON          accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON          `json:"-"`
}

// accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult]
type accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType string

const (
	AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeStar     AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "*"
	AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferral AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referral"
	AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceTypeReferrer AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultReferenceType = "referrer"
)

type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON `json:"-"`
}

// accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo]
type accountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess bool

const (
	AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccessTrue AccountLoadBalancerMonitorReferenceAccountLoadBalancerMonitorsListMonitorReferencesResponseSuccess = true
)
