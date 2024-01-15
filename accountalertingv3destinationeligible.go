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

// AccountAlertingV3DestinationEligibleService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationEligibleService] method instead.
type AccountAlertingV3DestinationEligibleService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationEligibleService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationEligibleService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationEligibleService) {
	r = &AccountAlertingV3DestinationEligibleService{}
	r.Options = opts
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *AccountAlertingV3DestinationEligibleService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationEligibleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAlertingV3DestinationEligibleListResponse struct {
	Errors     []AccountAlertingV3DestinationEligibleListResponseError    `json:"errors"`
	Messages   []AccountAlertingV3DestinationEligibleListResponseMessage  `json:"messages"`
	Result     interface{}                                                `json:"result"`
	ResultInfo AccountAlertingV3DestinationEligibleListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationEligibleListResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationEligibleListResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationEligibleListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationEligibleListResponse]
type accountAlertingV3DestinationEligibleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationEligibleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationEligibleListResponseError struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    accountAlertingV3DestinationEligibleListResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationEligibleListResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationEligibleListResponseError]
type accountAlertingV3DestinationEligibleListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationEligibleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationEligibleListResponseMessage struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountAlertingV3DestinationEligibleListResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationEligibleListResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationEligibleListResponseMessage]
type accountAlertingV3DestinationEligibleListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationEligibleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationEligibleListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       accountAlertingV3DestinationEligibleListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationEligibleListResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationEligibleListResponseResultInfo]
type accountAlertingV3DestinationEligibleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationEligibleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationEligibleListResponseSuccess bool

const (
	AccountAlertingV3DestinationEligibleListResponseSuccessTrue AccountAlertingV3DestinationEligibleListResponseSuccess = true
)
