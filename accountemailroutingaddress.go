// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountEmailRoutingAddressService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountEmailRoutingAddressService] method instead.
type AccountEmailRoutingAddressService struct {
	Options []option.RequestOption
}

// NewAccountEmailRoutingAddressService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEmailRoutingAddressService(opts ...option.RequestOption) (r *AccountEmailRoutingAddressService) {
	r = &AccountEmailRoutingAddressService{}
	r.Options = opts
	return
}

// Gets information for a specific destination email already created.
func (r *AccountEmailRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *AccountEmailRoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific destination address.
func (r *AccountEmailRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *AccountEmailRoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *AccountEmailRoutingAddressService) EmailRoutingDestinationAddressesNewADestinationAddress(ctx context.Context, accountIdentifier string, body AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams, opts ...option.RequestOption) (res *AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing destination addresses.
func (r *AccountEmailRoutingAddressService) EmailRoutingDestinationAddressesListDestinationAddresses(ctx context.Context, accountIdentifier string, query AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams, opts ...option.RequestOption) (res *shared.Page[AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

type AccountEmailRoutingAddressGetResponse struct {
	Errors   []AccountEmailRoutingAddressGetResponseError   `json:"errors"`
	Messages []AccountEmailRoutingAddressGetResponseMessage `json:"messages"`
	Result   AccountEmailRoutingAddressGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountEmailRoutingAddressGetResponseSuccess `json:"success"`
	JSON    accountEmailRoutingAddressGetResponseJSON    `json:"-"`
}

// accountEmailRoutingAddressGetResponseJSON contains the JSON metadata for the
// struct [AccountEmailRoutingAddressGetResponse]
type accountEmailRoutingAddressGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountEmailRoutingAddressGetResponseErrorJSON `json:"-"`
}

// accountEmailRoutingAddressGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountEmailRoutingAddressGetResponseError]
type accountEmailRoutingAddressGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountEmailRoutingAddressGetResponseMessageJSON `json:"-"`
}

// accountEmailRoutingAddressGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountEmailRoutingAddressGetResponseMessage]
type accountEmailRoutingAddressGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressGetResponseResult struct {
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address identifier.
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                       `json:"verified" format:"date-time"`
	JSON     accountEmailRoutingAddressGetResponseResultJSON `json:"-"`
}

// accountEmailRoutingAddressGetResponseResultJSON contains the JSON metadata for
// the struct [AccountEmailRoutingAddressGetResponseResult]
type accountEmailRoutingAddressGetResponseResultJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountEmailRoutingAddressGetResponseSuccess bool

const (
	AccountEmailRoutingAddressGetResponseSuccessTrue AccountEmailRoutingAddressGetResponseSuccess = true
)

type AccountEmailRoutingAddressDeleteResponse struct {
	Errors   []AccountEmailRoutingAddressDeleteResponseError   `json:"errors"`
	Messages []AccountEmailRoutingAddressDeleteResponseMessage `json:"messages"`
	Result   AccountEmailRoutingAddressDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountEmailRoutingAddressDeleteResponseSuccess `json:"success"`
	JSON    accountEmailRoutingAddressDeleteResponseJSON    `json:"-"`
}

// accountEmailRoutingAddressDeleteResponseJSON contains the JSON metadata for the
// struct [AccountEmailRoutingAddressDeleteResponse]
type accountEmailRoutingAddressDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountEmailRoutingAddressDeleteResponseErrorJSON `json:"-"`
}

// accountEmailRoutingAddressDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountEmailRoutingAddressDeleteResponseError]
type accountEmailRoutingAddressDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountEmailRoutingAddressDeleteResponseMessageJSON `json:"-"`
}

// accountEmailRoutingAddressDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountEmailRoutingAddressDeleteResponseMessage]
type accountEmailRoutingAddressDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressDeleteResponseResult struct {
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address identifier.
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                          `json:"verified" format:"date-time"`
	JSON     accountEmailRoutingAddressDeleteResponseResultJSON `json:"-"`
}

// accountEmailRoutingAddressDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountEmailRoutingAddressDeleteResponseResult]
type accountEmailRoutingAddressDeleteResponseResultJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountEmailRoutingAddressDeleteResponseSuccess bool

const (
	AccountEmailRoutingAddressDeleteResponseSuccessTrue AccountEmailRoutingAddressDeleteResponseSuccess = true
)

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse struct {
	Errors   []AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError   `json:"errors"`
	Messages []AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage `json:"messages"`
	Result   AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess `json:"success"`
	JSON    accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON    `json:"-"`
}

// accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON
// contains the JSON metadata for the struct
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse]
type accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON `json:"-"`
}

// accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError]
type accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON `json:"-"`
}

// accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage]
type accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult struct {
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address identifier.
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                                                                          `json:"verified" format:"date-time"`
	JSON     accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON `json:"-"`
}

// accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON
// contains the JSON metadata for the struct
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult]
type accountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess bool

const (
	AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccessTrue AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess = true
)

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse struct {
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address identifier.
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                                                                      `json:"verified" format:"date-time"`
	JSON     accountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON `json:"-"`
}

// accountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON
// contains the JSON metadata for the struct
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse]
type accountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified] `query:"verified"`
}

// URLQuery serializes
// [AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams]'s
// query parameters as `url.Values`.
func (r AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection string

const (
	AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionAsc  AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "asc"
	AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionDesc AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "desc"
)

// Filter by verified destination addresses.
type AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified bool

const (
	AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedTrue  AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = true
	AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedFalse AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = false
)
