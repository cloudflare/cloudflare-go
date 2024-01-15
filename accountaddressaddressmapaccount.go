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

// AccountAddressAddressMapAccountService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapAccountService] method instead.
type AccountAddressAddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapAccountService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapAccountService(opts ...option.RequestOption) (r *AccountAddressAddressMapAccountService) {
	r = &AccountAddressAddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AccountAddressAddressMapAccountService) Update(ctx context.Context, accountIdentifier1 string, addressMapIdentifier string, accountIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountIdentifier1, addressMapIdentifier, accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove an account as a member of a particular address map.
func (r *AccountAddressAddressMapAccountService) Delete(ctx context.Context, accountIdentifier1 string, addressMapIdentifier string, accountIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountIdentifier1, addressMapIdentifier, accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAddressAddressMapAccountUpdateResponse struct {
	Errors     []AccountAddressAddressMapAccountUpdateResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapAccountUpdateResponseMessage  `json:"messages"`
	Result     []interface{}                                           `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapAccountUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapAccountUpdateResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapAccountUpdateResponseJSON    `json:"-"`
}

// accountAddressAddressMapAccountUpdateResponseJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapAccountUpdateResponse]
type accountAddressAddressMapAccountUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountUpdateResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAddressAddressMapAccountUpdateResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapAccountUpdateResponseErrorJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapAccountUpdateResponseError]
type accountAddressAddressMapAccountUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountUpdateResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAddressAddressMapAccountUpdateResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapAccountUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapAccountUpdateResponseMessage]
type accountAddressAddressMapAccountUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       accountAddressAddressMapAccountUpdateResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapAccountUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAddressAddressMapAccountUpdateResponseResultInfo]
type accountAddressAddressMapAccountUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapAccountUpdateResponseSuccess bool

const (
	AccountAddressAddressMapAccountUpdateResponseSuccessTrue AccountAddressAddressMapAccountUpdateResponseSuccess = true
)

type AccountAddressAddressMapAccountDeleteResponse struct {
	Errors     []AccountAddressAddressMapAccountDeleteResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapAccountDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                           `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapAccountDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapAccountDeleteResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapAccountDeleteResponseJSON    `json:"-"`
}

// accountAddressAddressMapAccountDeleteResponseJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapAccountDeleteResponse]
type accountAddressAddressMapAccountDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountDeleteResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAddressAddressMapAccountDeleteResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapAccountDeleteResponseErrorJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapAccountDeleteResponseError]
type accountAddressAddressMapAccountDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountDeleteResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAddressAddressMapAccountDeleteResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapAccountDeleteResponseMessageJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapAccountDeleteResponseMessage]
type accountAddressAddressMapAccountDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapAccountDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       accountAddressAddressMapAccountDeleteResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapAccountDeleteResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAddressAddressMapAccountDeleteResponseResultInfo]
type accountAddressAddressMapAccountDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapAccountDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapAccountDeleteResponseSuccess bool

const (
	AccountAddressAddressMapAccountDeleteResponseSuccessTrue AccountAddressAddressMapAccountDeleteResponseSuccess = true
)
