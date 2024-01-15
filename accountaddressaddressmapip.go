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

// AccountAddressAddressMapIPService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapIPService] method instead.
type AccountAddressAddressMapIPService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapIPService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapIPService(opts ...option.RequestOption) (r *AccountAddressAddressMapIPService) {
	r = &AccountAddressAddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AccountAddressAddressMapIPService) Update(ctx context.Context, accountIdentifier string, addressMapIdentifier string, ipAddress string, opts ...option.RequestOption) (res *AccountAddressAddressMapIPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountIdentifier, addressMapIdentifier, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove an IP from a particular address map.
func (r *AccountAddressAddressMapIPService) Delete(ctx context.Context, accountIdentifier string, addressMapIdentifier string, ipAddress string, opts ...option.RequestOption) (res *AccountAddressAddressMapIPDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountIdentifier, addressMapIdentifier, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAddressAddressMapIPUpdateResponse struct {
	Errors     []AccountAddressAddressMapIPUpdateResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapIPUpdateResponseMessage  `json:"messages"`
	Result     []interface{}                                      `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapIPUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapIPUpdateResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapIPUpdateResponseJSON    `json:"-"`
}

// accountAddressAddressMapIPUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapIPUpdateResponse]
type accountAddressAddressMapIPUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAddressAddressMapIPUpdateResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapIPUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapIPUpdateResponseError]
type accountAddressAddressMapIPUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAddressAddressMapIPUpdateResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapIPUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapIPUpdateResponseMessage]
type accountAddressAddressMapIPUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       accountAddressAddressMapIPUpdateResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapIPUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapIPUpdateResponseResultInfo]
type accountAddressAddressMapIPUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapIPUpdateResponseSuccess bool

const (
	AccountAddressAddressMapIPUpdateResponseSuccessTrue AccountAddressAddressMapIPUpdateResponseSuccess = true
)

type AccountAddressAddressMapIPDeleteResponse struct {
	Errors     []AccountAddressAddressMapIPDeleteResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapIPDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                      `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapIPDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapIPDeleteResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapIPDeleteResponseJSON    `json:"-"`
}

// accountAddressAddressMapIPDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAddressAddressMapIPDeleteResponse]
type accountAddressAddressMapIPDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAddressAddressMapIPDeleteResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapIPDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapIPDeleteResponseError]
type accountAddressAddressMapIPDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAddressAddressMapIPDeleteResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapIPDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapIPDeleteResponseMessage]
type accountAddressAddressMapIPDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapIPDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       accountAddressAddressMapIPDeleteResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapIPDeleteResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapIPDeleteResponseResultInfo]
type accountAddressAddressMapIPDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapIPDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapIPDeleteResponseSuccess bool

const (
	AccountAddressAddressMapIPDeleteResponseSuccessTrue AccountAddressAddressMapIPDeleteResponseSuccess = true
)
