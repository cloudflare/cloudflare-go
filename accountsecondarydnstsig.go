// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountSecondaryDNSTsigService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountSecondaryDNSTsigService] method instead.
type AccountSecondaryDNSTsigService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSTsigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSTsigService(opts ...option.RequestOption) (r *AccountSecondaryDNSTsigService) {
	r = &AccountSecondaryDNSTsigService{}
	r.Options = opts
	return
}

// Get TSIG.
func (r *AccountSecondaryDNSTsigService) Get(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify TSIG.
func (r *AccountSecondaryDNSTsigService) Update(ctx context.Context, accountIdentifier interface{}, identifier interface{}, body AccountSecondaryDNSTsigUpdateParams, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete TSIG.
func (r *AccountSecondaryDNSTsigService) Delete(ctx context.Context, accountIdentifier interface{}, identifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create TSIG.
func (r *AccountSecondaryDNSTsigService) SecondaryDNSTsigNewTsig(ctx context.Context, accountIdentifier interface{}, body AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List TSIGs.
func (r *AccountSecondaryDNSTsigService) SecondaryDNSTsigListTsiGs(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/secondary_dns/tsigs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountSecondaryDNSTsigGetResponse struct {
	Errors   []AccountSecondaryDNSTsigGetResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSTsigGetResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSTsigGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSTsigGetResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSTsigGetResponseJSON    `json:"-"`
}

// accountSecondaryDNSTsigGetResponseJSON contains the JSON metadata for the struct
// [AccountSecondaryDNSTsigGetResponse]
type accountSecondaryDNSTsigGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountSecondaryDNSTsigGetResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSTsigGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigGetResponseError]
type accountSecondaryDNSTsigGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountSecondaryDNSTsigGetResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSTsigGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigGetResponseMessage]
type accountSecondaryDNSTsigGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigGetResponseResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                       `json:"secret,required"`
	JSON   accountSecondaryDNSTsigGetResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigGetResponseResultJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigGetResponseResult]
type accountSecondaryDNSTsigGetResponseResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSTsigGetResponseSuccess bool

const (
	AccountSecondaryDNSTsigGetResponseSuccessTrue AccountSecondaryDNSTsigGetResponseSuccess = true
)

type AccountSecondaryDNSTsigUpdateResponse struct {
	Errors   []AccountSecondaryDNSTsigUpdateResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSTsigUpdateResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSTsigUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSTsigUpdateResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSTsigUpdateResponseJSON    `json:"-"`
}

// accountSecondaryDNSTsigUpdateResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigUpdateResponse]
type accountSecondaryDNSTsigUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountSecondaryDNSTsigUpdateResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSTsigUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigUpdateResponseError]
type accountSecondaryDNSTsigUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountSecondaryDNSTsigUpdateResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSTsigUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigUpdateResponseMessage]
type accountSecondaryDNSTsigUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigUpdateResponseResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                          `json:"secret,required"`
	JSON   accountSecondaryDNSTsigUpdateResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigUpdateResponseResult]
type accountSecondaryDNSTsigUpdateResponseResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSTsigUpdateResponseSuccess bool

const (
	AccountSecondaryDNSTsigUpdateResponseSuccessTrue AccountSecondaryDNSTsigUpdateResponseSuccess = true
)

type AccountSecondaryDNSTsigDeleteResponse struct {
	Errors   []AccountSecondaryDNSTsigDeleteResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSTsigDeleteResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSTsigDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSTsigDeleteResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSTsigDeleteResponseJSON    `json:"-"`
}

// accountSecondaryDNSTsigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountSecondaryDNSTsigDeleteResponse]
type accountSecondaryDNSTsigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountSecondaryDNSTsigDeleteResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSTsigDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigDeleteResponseError]
type accountSecondaryDNSTsigDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountSecondaryDNSTsigDeleteResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSTsigDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigDeleteResponseMessage]
type accountSecondaryDNSTsigDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigDeleteResponseResult struct {
	ID   interface{}                                     `json:"id"`
	JSON accountSecondaryDNSTsigDeleteResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountSecondaryDNSTsigDeleteResponseResult]
type accountSecondaryDNSTsigDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSTsigDeleteResponseSuccess bool

const (
	AccountSecondaryDNSTsigDeleteResponseSuccessTrue AccountSecondaryDNSTsigDeleteResponseSuccess = true
)

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponse struct {
	Errors   []AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseError   `json:"errors"`
	Messages []AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessage `json:"messages"`
	Result   AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON    `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON contains the JSON
// metadata for the struct [AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponse]
type accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseError struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseError]
type accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessage struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessage]
type accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                                           `json:"secret,required"`
	JSON   accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResultJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResult]
type accountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseSuccess bool

const (
	AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseSuccessTrue AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigResponseSuccess = true
)

type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse struct {
	Errors     []AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseError    `json:"errors"`
	Messages   []AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessage  `json:"messages"`
	Result     []AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResult   `json:"result"`
	ResultInfo AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseSuccess `json:"success"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON    `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON contains the JSON
// metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse]
type accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseErrorJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseError]
type accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessageJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessage]
type accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResult struct {
	ID interface{} `json:"id,required"`
	// TSIG algorithm.
	Algo string `json:"algo,required"`
	// TSIG key name.
	Name string `json:"name,required"`
	// TSIG secret.
	Secret string                                                             `json:"secret,required"`
	JSON   accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultJSON contains the
// JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResult]
type accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultJSON struct {
	ID          apijson.Field
	Algo        apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfoJSON `json:"-"`
}

// accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfoJSON contains
// the JSON metadata for the struct
// [AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfo]
type accountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseSuccess bool

const (
	AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseSuccessTrue AccountSecondaryDNSTsigSecondaryDNSTsigListTsiGsResponseSuccess = true
)

type AccountSecondaryDNSTsigUpdateParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r AccountSecondaryDNSTsigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams struct {
	// TSIG algorithm.
	Algo param.Field[string] `json:"algo,required"`
	// TSIG key name.
	Name param.Field[string] `json:"name,required"`
	// TSIG secret.
	Secret param.Field[string] `json:"secret,required"`
}

func (r AccountSecondaryDNSTsigSecondaryDNSTsigNewTsigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
