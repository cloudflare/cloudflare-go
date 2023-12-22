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
func (r *AccountEmailRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a specific destination address.
func (r *AccountEmailRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *AccountEmailRoutingAddressService) EmailRoutingDestinationAddressesNewADestinationAddress(ctx context.Context, accountIdentifier string, body AccountEmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing destination addresses.
func (r *AccountEmailRoutingAddressService) EmailRoutingDestinationAddressesListDestinationAddresses(ctx context.Context, accountIdentifier string, query AccountEmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams, opts ...option.RequestOption) (res *DestinationAddressesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DestinationAddressResponseSingle struct {
	Errors   []DestinationAddressResponseSingleError   `json:"errors"`
	Messages []DestinationAddressResponseSingleMessage `json:"messages"`
	Result   DestinationAddressResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success DestinationAddressResponseSingleSuccess `json:"success"`
	JSON    destinationAddressResponseSingleJSON    `json:"-"`
}

// destinationAddressResponseSingleJSON contains the JSON metadata for the struct
// [DestinationAddressResponseSingle]
type destinationAddressResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressResponseSingleError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    destinationAddressResponseSingleErrorJSON `json:"-"`
}

// destinationAddressResponseSingleErrorJSON contains the JSON metadata for the
// struct [DestinationAddressResponseSingleError]
type destinationAddressResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressResponseSingleMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    destinationAddressResponseSingleMessageJSON `json:"-"`
}

// destinationAddressResponseSingleMessageJSON contains the JSON metadata for the
// struct [DestinationAddressResponseSingleMessage]
type destinationAddressResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressResponseSingleResult struct {
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
	Verified time.Time                                  `json:"verified" format:"date-time"`
	JSON     destinationAddressResponseSingleResultJSON `json:"-"`
}

// destinationAddressResponseSingleResultJSON contains the JSON metadata for the
// struct [DestinationAddressResponseSingleResult]
type destinationAddressResponseSingleResultJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DestinationAddressResponseSingleSuccess bool

const (
	DestinationAddressResponseSingleSuccessTrue DestinationAddressResponseSingleSuccess = true
)

type DestinationAddressesResponseCollection struct {
	Errors     []DestinationAddressesResponseCollectionError    `json:"errors"`
	Messages   []DestinationAddressesResponseCollectionMessage  `json:"messages"`
	Result     []DestinationAddressesResponseCollectionResult   `json:"result"`
	ResultInfo DestinationAddressesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success DestinationAddressesResponseCollectionSuccess `json:"success"`
	JSON    destinationAddressesResponseCollectionJSON    `json:"-"`
}

// destinationAddressesResponseCollectionJSON contains the JSON metadata for the
// struct [DestinationAddressesResponseCollection]
type destinationAddressesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressesResponseCollectionError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    destinationAddressesResponseCollectionErrorJSON `json:"-"`
}

// destinationAddressesResponseCollectionErrorJSON contains the JSON metadata for
// the struct [DestinationAddressesResponseCollectionError]
type destinationAddressesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressesResponseCollectionMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    destinationAddressesResponseCollectionMessageJSON `json:"-"`
}

// destinationAddressesResponseCollectionMessageJSON contains the JSON metadata for
// the struct [DestinationAddressesResponseCollectionMessage]
type destinationAddressesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressesResponseCollectionResult struct {
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
	Verified time.Time                                        `json:"verified" format:"date-time"`
	JSON     destinationAddressesResponseCollectionResultJSON `json:"-"`
}

// destinationAddressesResponseCollectionResultJSON contains the JSON metadata for
// the struct [DestinationAddressesResponseCollectionResult]
type destinationAddressesResponseCollectionResultJSON struct {
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressesResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DestinationAddressesResponseCollectionResultInfo struct {
	Count      interface{}                                          `json:"count"`
	Page       interface{}                                          `json:"page"`
	PerPage    interface{}                                          `json:"per_page"`
	TotalCount interface{}                                          `json:"total_count"`
	JSON       destinationAddressesResponseCollectionResultInfoJSON `json:"-"`
}

// destinationAddressesResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [DestinationAddressesResponseCollectionResultInfo]
type destinationAddressesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DestinationAddressesResponseCollectionSuccess bool

const (
	DestinationAddressesResponseCollectionSuccessTrue DestinationAddressesResponseCollectionSuccess = true
)

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
