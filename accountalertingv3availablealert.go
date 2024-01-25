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

// AccountAlertingV3AvailableAlertService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3AvailableAlertService] method instead.
type AccountAlertingV3AvailableAlertService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3AvailableAlertService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAlertingV3AvailableAlertService(opts ...option.RequestOption) (r *AccountAlertingV3AvailableAlertService) {
	r = &AccountAlertingV3AvailableAlertService{}
	r.Options = opts
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AccountAlertingV3AvailableAlertService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3AvailableAlertListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAlertingV3AvailableAlertListResponse struct {
	Errors     []AccountAlertingV3AvailableAlertListResponseError    `json:"errors"`
	Messages   []AccountAlertingV3AvailableAlertListResponseMessage  `json:"messages"`
	Result     interface{}                                           `json:"result"`
	ResultInfo AccountAlertingV3AvailableAlertListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3AvailableAlertListResponseSuccess `json:"success"`
	JSON    accountAlertingV3AvailableAlertListResponseJSON    `json:"-"`
}

// accountAlertingV3AvailableAlertListResponseJSON contains the JSON metadata for
// the struct [AccountAlertingV3AvailableAlertListResponse]
type accountAlertingV3AvailableAlertListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3AvailableAlertListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3AvailableAlertListResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAlertingV3AvailableAlertListResponseErrorJSON `json:"-"`
}

// accountAlertingV3AvailableAlertListResponseErrorJSON contains the JSON metadata
// for the struct [AccountAlertingV3AvailableAlertListResponseError]
type accountAlertingV3AvailableAlertListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3AvailableAlertListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3AvailableAlertListResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAlertingV3AvailableAlertListResponseMessageJSON `json:"-"`
}

// accountAlertingV3AvailableAlertListResponseMessageJSON contains the JSON
// metadata for the struct [AccountAlertingV3AvailableAlertListResponseMessage]
type accountAlertingV3AvailableAlertListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3AvailableAlertListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3AvailableAlertListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       accountAlertingV3AvailableAlertListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3AvailableAlertListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAlertingV3AvailableAlertListResponseResultInfo]
type accountAlertingV3AvailableAlertListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3AvailableAlertListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3AvailableAlertListResponseSuccess bool

const (
	AccountAlertingV3AvailableAlertListResponseSuccessTrue AccountAlertingV3AvailableAlertListResponseSuccess = true
)
