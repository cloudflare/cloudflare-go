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

// AccountAddressAddressMapZoneService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapZoneService] method instead.
type AccountAddressAddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapZoneService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapZoneService(opts ...option.RequestOption) (r *AccountAddressAddressMapZoneService) {
	r = &AccountAddressAddressMapZoneService{}
	r.Options = opts
	return
}

// Add a zone as a member of a particular address map.
func (r *AccountAddressAddressMapZoneService) Update(ctx context.Context, accountIdentifier string, addressMapIdentifier string, zoneIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountIdentifier, addressMapIdentifier, zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove a zone as a member of a particular address map.
func (r *AccountAddressAddressMapZoneService) Delete(ctx context.Context, accountIdentifier string, addressMapIdentifier string, zoneIdentifier string, opts ...option.RequestOption) (res *AccountAddressAddressMapZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountIdentifier, addressMapIdentifier, zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountAddressAddressMapZoneUpdateResponse struct {
	Errors     []AccountAddressAddressMapZoneUpdateResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapZoneUpdateResponseMessage  `json:"messages"`
	Result     []interface{}                                        `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapZoneUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapZoneUpdateResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapZoneUpdateResponseJSON    `json:"-"`
}

// accountAddressAddressMapZoneUpdateResponseJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapZoneUpdateResponse]
type accountAddressAddressMapZoneUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneUpdateResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAddressAddressMapZoneUpdateResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapZoneUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapZoneUpdateResponseError]
type accountAddressAddressMapZoneUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneUpdateResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAddressAddressMapZoneUpdateResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapZoneUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapZoneUpdateResponseMessage]
type accountAddressAddressMapZoneUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accountAddressAddressMapZoneUpdateResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapZoneUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapZoneUpdateResponseResultInfo]
type accountAddressAddressMapZoneUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapZoneUpdateResponseSuccess bool

const (
	AccountAddressAddressMapZoneUpdateResponseSuccessTrue AccountAddressAddressMapZoneUpdateResponseSuccess = true
)

type AccountAddressAddressMapZoneDeleteResponse struct {
	Errors     []AccountAddressAddressMapZoneDeleteResponseError    `json:"errors"`
	Messages   []AccountAddressAddressMapZoneDeleteResponseMessage  `json:"messages"`
	Result     []interface{}                                        `json:"result,nullable"`
	ResultInfo AccountAddressAddressMapZoneDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAddressAddressMapZoneDeleteResponseSuccess `json:"success"`
	JSON    accountAddressAddressMapZoneDeleteResponseJSON    `json:"-"`
}

// accountAddressAddressMapZoneDeleteResponseJSON contains the JSON metadata for
// the struct [AccountAddressAddressMapZoneDeleteResponse]
type accountAddressAddressMapZoneDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneDeleteResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountAddressAddressMapZoneDeleteResponseErrorJSON `json:"-"`
}

// accountAddressAddressMapZoneDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapZoneDeleteResponseError]
type accountAddressAddressMapZoneDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneDeleteResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountAddressAddressMapZoneDeleteResponseMessageJSON `json:"-"`
}

// accountAddressAddressMapZoneDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountAddressAddressMapZoneDeleteResponseMessage]
type accountAddressAddressMapZoneDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressAddressMapZoneDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       accountAddressAddressMapZoneDeleteResponseResultInfoJSON `json:"-"`
}

// accountAddressAddressMapZoneDeleteResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAddressAddressMapZoneDeleteResponseResultInfo]
type accountAddressAddressMapZoneDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressAddressMapZoneDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressAddressMapZoneDeleteResponseSuccess bool

const (
	AccountAddressAddressMapZoneDeleteResponseSuccessTrue AccountAddressAddressMapZoneDeleteResponseSuccess = true
)
