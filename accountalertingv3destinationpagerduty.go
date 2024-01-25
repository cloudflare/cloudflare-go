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

// AccountAlertingV3DestinationPagerdutyService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAlertingV3DestinationPagerdutyService] method instead.
type AccountAlertingV3DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3DestinationPagerdutyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationPagerdutyService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationPagerdutyService) {
	r = &AccountAlertingV3DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Get a list of all configured PagerDuty services.
func (r *AccountAlertingV3DestinationPagerdutyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *AccountAlertingV3DestinationPagerdutyService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *AccountAlertingV3DestinationPagerdutyService) Connect(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyConnectResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Links PagerDuty with the account using the integration token.
func (r *AccountAlertingV3DestinationPagerdutyService) Link(ctx context.Context, accountID string, tokenID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAlertingV3DestinationPagerdutyListResponse struct {
	Errors     []AccountAlertingV3DestinationPagerdutyListResponseError    `json:"errors"`
	Messages   []AccountAlertingV3DestinationPagerdutyListResponseMessage  `json:"messages"`
	Result     []AccountAlertingV3DestinationPagerdutyListResponseResult   `json:"result"`
	ResultInfo AccountAlertingV3DestinationPagerdutyListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationPagerdutyListResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationPagerdutyListResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationPagerdutyListResponse]
type accountAlertingV3DestinationPagerdutyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyListResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyListResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationPagerdutyListResponseError]
type accountAlertingV3DestinationPagerdutyListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyListResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyListResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseMessage]
type accountAlertingV3DestinationPagerdutyListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyListResponseResult struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string                                                      `json:"name"`
	JSON accountAlertingV3DestinationPagerdutyListResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseResult]
type accountAlertingV3DestinationPagerdutyListResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseResultInfo]
type accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyListResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyListResponseSuccessTrue AccountAlertingV3DestinationPagerdutyListResponseSuccess = true
)

type AccountAlertingV3DestinationPagerdutyDeleteResponse struct {
	Errors     []AccountAlertingV3DestinationPagerdutyDeleteResponseError    `json:"errors"`
	Messages   []AccountAlertingV3DestinationPagerdutyDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                                 `json:"result,nullable"`
	ResultInfo AccountAlertingV3DestinationPagerdutyDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationPagerdutyDeleteResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationPagerdutyDeleteResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyDeleteResponseJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationPagerdutyDeleteResponse]
type accountAlertingV3DestinationPagerdutyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyDeleteResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyDeleteResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyDeleteResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyDeleteResponseError]
type accountAlertingV3DestinationPagerdutyDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyDeleteResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyDeleteResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyDeleteResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyDeleteResponseMessage]
type accountAlertingV3DestinationPagerdutyDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       accountAlertingV3DestinationPagerdutyDeleteResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyDeleteResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyDeleteResponseResultInfo]
type accountAlertingV3DestinationPagerdutyDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyDeleteResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyDeleteResponseSuccessTrue AccountAlertingV3DestinationPagerdutyDeleteResponseSuccess = true
)

type AccountAlertingV3DestinationPagerdutyConnectResponse struct {
	Errors   []AccountAlertingV3DestinationPagerdutyConnectResponseError   `json:"errors"`
	Messages []AccountAlertingV3DestinationPagerdutyConnectResponseMessage `json:"messages"`
	Result   AccountAlertingV3DestinationPagerdutyConnectResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationPagerdutyConnectResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationPagerdutyConnectResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectResponseJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationPagerdutyConnectResponse]
type accountAlertingV3DestinationPagerdutyConnectResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyConnectResponseError struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyConnectResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectResponseError]
type accountAlertingV3DestinationPagerdutyConnectResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyConnectResponseMessage struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyConnectResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectResponseMessage]
type accountAlertingV3DestinationPagerdutyConnectResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyConnectResponseResult struct {
	// UUID
	ID   string                                                         `json:"id"`
	JSON accountAlertingV3DestinationPagerdutyConnectResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyConnectResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyConnectResponseResult]
type accountAlertingV3DestinationPagerdutyConnectResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyConnectResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyConnectResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyConnectResponseSuccessTrue AccountAlertingV3DestinationPagerdutyConnectResponseSuccess = true
)

type AccountAlertingV3DestinationPagerdutyLinkResponse struct {
	Errors   []AccountAlertingV3DestinationPagerdutyLinkResponseError   `json:"errors"`
	Messages []AccountAlertingV3DestinationPagerdutyLinkResponseMessage `json:"messages"`
	Result   AccountAlertingV3DestinationPagerdutyLinkResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAlertingV3DestinationPagerdutyLinkResponseSuccess `json:"success"`
	JSON    accountAlertingV3DestinationPagerdutyLinkResponseJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyLinkResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationPagerdutyLinkResponse]
type accountAlertingV3DestinationPagerdutyLinkResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyLinkResponseError struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyLinkResponseErrorJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyLinkResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationPagerdutyLinkResponseError]
type accountAlertingV3DestinationPagerdutyLinkResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyLinkResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyLinkResponseMessage struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountAlertingV3DestinationPagerdutyLinkResponseMessageJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyLinkResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyLinkResponseMessage]
type accountAlertingV3DestinationPagerdutyLinkResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyLinkResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAlertingV3DestinationPagerdutyLinkResponseResult struct {
	// UUID
	ID   string                                                      `json:"id"`
	JSON accountAlertingV3DestinationPagerdutyLinkResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyLinkResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyLinkResponseResult]
type accountAlertingV3DestinationPagerdutyLinkResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyLinkResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyLinkResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyLinkResponseSuccessTrue AccountAlertingV3DestinationPagerdutyLinkResponseSuccess = true
)
