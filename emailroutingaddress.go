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

// Deletes a specific destination address.
func (r *EmailRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingAddressDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *EmailRoutingAddressService) EmailRoutingDestinationAddressesNewADestinationAddress(ctx context.Context, accountIdentifier string, body EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams, opts ...option.RequestOption) (res *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *EmailRoutingAddressService) EmailRoutingDestinationAddressesListDestinationAddresses(ctx context.Context, accountIdentifier string, query EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams, opts ...option.RequestOption) (res *[]EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *EmailRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingAddressGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingAddressDeleteResponse struct {
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
	Verified time.Time                             `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressDeleteResponseJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseJSON contains the JSON metadata for the struct
// [EmailRoutingAddressDeleteResponse]
type emailRoutingAddressDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse struct {
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
	Verified time.Time                                                                             `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse struct {
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
	Verified time.Time                                                                               `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponse struct {
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
	Verified time.Time                          `json:"verified" format:"date-time"`
	JSON     emailRoutingAddressGetResponseJSON `json:"-"`
}

// emailRoutingAddressGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingAddressGetResponse]
type emailRoutingAddressGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseEnvelope struct {
	Errors   []EmailRoutingAddressDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingAddressDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingAddressDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingAddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingAddressDeleteResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingAddressDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingAddressDeleteResponseEnvelope]
type emailRoutingAddressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    emailRoutingAddressDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EmailRoutingAddressDeleteResponseEnvelopeErrors]
type emailRoutingAddressDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    emailRoutingAddressDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingAddressDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingAddressDeleteResponseEnvelopeMessages]
type emailRoutingAddressDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressDeleteResponseEnvelopeSuccess bool

const (
	EmailRoutingAddressDeleteResponseEnvelopeSuccessTrue EmailRoutingAddressDeleteResponseEnvelopeSuccess = true
)

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelope struct {
	Errors   []EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelope]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrors struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrors]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessages struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessages]
type emailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeSuccess bool

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeSuccessTrue EmailRoutingAddressEmailRoutingDestinationAddressesNewADestinationAddressResponseEnvelopeSuccess = true
)

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

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope struct {
	Errors   []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess bool

const (
	EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccessTrue EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess = true
)

type EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                   `json:"total_count"`
	JSON       emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo]
type emailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseEnvelope struct {
	Errors   []EmailRoutingAddressGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingAddressGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingAddressGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingAddressGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingAddressGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingAddressGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingAddressGetResponseEnvelope]
type emailRoutingAddressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    emailRoutingAddressGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingAddressGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingAddressGetResponseEnvelopeErrors]
type emailRoutingAddressGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingAddressGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    emailRoutingAddressGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingAddressGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingAddressGetResponseEnvelopeMessages]
type emailRoutingAddressGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingAddressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingAddressGetResponseEnvelopeSuccess bool

const (
	EmailRoutingAddressGetResponseEnvelopeSuccessTrue EmailRoutingAddressGetResponseEnvelopeSuccess = true
)
