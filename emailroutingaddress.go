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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// EmailRoutingAddressService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingAddressService]
// method instead.
type EmailRoutingAddressService struct {
	Options []option.RequestOption
}

// NewEmailRoutingAddressService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingAddressService(opts ...option.RequestOption) (r *EmailRoutingAddressService) {
	r = &EmailRoutingAddressService{}
	r.Options = opts
	return
}

// Gets information for a specific destination email already created.
func (r *EmailRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific destination address.
func (r *EmailRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *EmailRoutingAddressService) EmailRoutingDestinationAddressesNewADestinationAddress(ctx context.Context, accountIdentifier string, body EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams, opts ...option.RequestOption) (res *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing destination addresses.
func (r *EmailRoutingAddressService) EmailRoutingDestinationAddressesListDestinationAddresses(ctx context.Context, accountIdentifier string, query EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams, opts ...option.RequestOption) (res *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EmailRoutingAddressGetResponse struct {
	Errors   []EmailRoutingAddressGetResponseError   `json:"errors"`
	Messages []EmailRoutingAddressGetResponseMessage `json:"messages"`
	Result   EmailRoutingAddressGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingAddressGetResponseSuccess `json:"success"`
	JSON    emailRoutingAddressGetResponseJSON    `json:"-"`
}

// emailRoutingAddressGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingAddressGetResponse]
type emailRoutingAddressGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    emailRoutingAddressGetResponseErrorJSON `json:"-"`
}

// emailRoutingAddressGetResponseErrorJSON contains the JSON metadata for the
// struct [EmailRoutingAddressGetResponseError]
type emailRoutingAddressGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    emailRoutingAddressGetResponseMessageJSON `json:"-"`
}

// emailRoutingAddressGetResponseMessageJSON contains the JSON metadata for the
// struct [EmailRoutingAddressGetResponseMessage]
type emailRoutingAddressGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseResult struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressGetResponseResultJSON `json:"-"`
}

// emailRoutingAddressGetResponseResultJSON contains the JSON metadata for the
// struct [EmailRoutingAddressGetResponseResult]
type emailRoutingAddressGetResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressGetResponseSuccess bool

const (
	EmailRoutingAddressGetResponseSuccessTrue EmailRoutingAddressGetResponseSuccess = true
)

type EmailRoutingAddressDeleteResponse struct {
	Errors   []EmailRoutingAddressDeleteResponseError   `json:"errors"`
	Messages []EmailRoutingAddressDeleteResponseMessage `json:"messages"`
	Result   EmailRoutingAddressDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingAddressDeleteResponseSuccess `json:"success"`
	JSON    emailRoutingAddressDeleteResponseJSON    `json:"-"`
}

// emailRoutingAddressDeleteResponseJSON contains the JSON metadata for the struct
// [EmailRoutingAddressDeleteResponse]
type emailRoutingAddressDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    emailRoutingAddressDeleteResponseErrorJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseErrorJSON contains the JSON metadata for the
// struct [EmailRoutingAddressDeleteResponseError]
type emailRoutingAddressDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    emailRoutingAddressDeleteResponseMessageJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseMessageJSON contains the JSON metadata for the
// struct [EmailRoutingAddressDeleteResponseMessage]
type emailRoutingAddressDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseResult struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                   `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressDeleteResponseResultJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseResultJSON contains the JSON metadata for the
// struct [EmailRoutingAddressDeleteResponseResult]
type emailRoutingAddressDeleteResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressDeleteResponseSuccess bool

const (
	EmailRoutingAddressDeleteResponseSuccessTrue EmailRoutingAddressDeleteResponseSuccess = true
)

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse struct {
	Errors   []EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError   `json:"errors"`
	Messages []EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage `json:"messages"`
	Result   EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess `json:"success"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON    `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                                                                   `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess bool

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccessTrue EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseSuccess = true
)

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse struct {
	Errors     []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseError    `json:"errors"`
	Messages   []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessage  `json:"messages"`
	Result     []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResult   `json:"result"`
	ResultInfo EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseSuccess `json:"success"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON    `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseErrorJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseError]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessageJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessage]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResult struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time                                                                                     `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResult]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfo struct {
	Count      interface{}                                                                                       `json:"count"`
	Page       interface{}                                                                                       `json:"page"`
	PerPage    interface{}                                                                                       `json:"per_page"`
	TotalCount interface{}                                                                                       `json:"total_count"`
	JSON       emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfoJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfo]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseSuccess bool

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseSuccessTrue EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseSuccess = true
)

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified] `query:"verified"`
}

// URLQuery serializes
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams]'s
// query parameters as `url.Values`.
func (r EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection string

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionAsc  EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "asc"
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionDesc EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "desc"
)

// Filter by verified destination addresses.
type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified bool

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedTrue  EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = true
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedFalse EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = false
)
