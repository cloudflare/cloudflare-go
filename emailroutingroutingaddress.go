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

// EmailRoutingRoutingAddressService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingAddressService] method instead.
type EmailRoutingRoutingAddressService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingAddressService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingAddressService(opts ...option.RequestOption) (r *EmailRoutingRoutingAddressService) {
	r = &EmailRoutingRoutingAddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *EmailRoutingRoutingAddressService) New(ctx context.Context, accountIdentifier string, body EmailRoutingRoutingAddressNewParams, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a specific destination address.
func (r *EmailRoutingRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *EmailRoutingRoutingAddressService) EmailRoutingDestinationAddressesListDestinationAddresses(ctx context.Context, accountIdentifier string, query EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams, opts ...option.RequestOption) (res *[]EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *EmailRoutingRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingAddressNewResponse struct {
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
	Verified time.Time                                 `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressNewResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressNewResponse]
type emailRoutingRoutingAddressNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressDeleteResponse struct {
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
	Verified time.Time                                    `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressDeleteResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressDeleteResponse]
type emailRoutingRoutingAddressDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse struct {
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
	Verified time.Time                                                                                      `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON
// contains the JSON metadata for the struct
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse]
type emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressGetResponse struct {
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
	Verified time.Time                                 `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressGetResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressGetResponse]
type emailRoutingRoutingAddressGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressNewParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r EmailRoutingRoutingAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRoutingRoutingAddressNewResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingAddressNewResponseEnvelope]
type emailRoutingRoutingAddressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressNewResponseEnvelopeErrors]
type emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressNewResponseEnvelopeMessages]
type emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingAddressNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressNewResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressNewResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingAddressDeleteResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingAddressDeleteResponseEnvelope]
type emailRoutingRoutingAddressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors]
type emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages]
type emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified] `query:"verified"`
}

// URLQuery serializes
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams]'s
// query parameters as `url.Values`.
func (r EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection string

const (
	EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionAsc  EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "asc"
	EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirectionDesc EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsDirection = "desc"
)

// Filter by verified destination addresses.
type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified bool

const (
	EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedTrue  EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = true
	EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerifiedFalse EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesParamsVerified = false
)

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope]
type emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors]
type emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages struct {
	Code    int64                                                                                                          `json:"code,required"`
	Message string                                                                                                         `json:"message,required"`
	JSON    emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages]
type emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                          `json:"total_count"`
	JSON       emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo]
type emailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressEmailRoutingDestinationAddressesListDestinationAddressesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingAddressGetResponseEnvelope]
type emailRoutingRoutingAddressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressGetResponseEnvelopeErrors]
type emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRoutingAddressGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressGetResponseEnvelopeMessages]
type emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRoutingAddressGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressGetResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressGetResponseEnvelopeSuccess = true
)
