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

// AccountCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCustomPageService] method
// instead.
type AccountCustomPageService struct {
	Options []option.RequestOption
}

// NewAccountCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCustomPageService(opts ...option.RequestOption) (r *AccountCustomPageService) {
	r = &AccountCustomPageService{}
	r.Options = opts
	return
}

// Fetches the details of a custom page.
func (r *AccountCustomPageService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountCustomPageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration of an existing custom page.
func (r *AccountCustomPageService) Update(ctx context.Context, accountIdentifier string, identifier string, body AccountCustomPageUpdateParams, opts ...option.RequestOption) (res *AccountCustomPageUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all the custom pages at the account level.
func (r *AccountCustomPageService) CustomPagesForAnAccountListCustomPages(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountCustomPageCustomPagesForAnAccountListCustomPagesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCustomPageGetResponse struct {
	Errors   []AccountCustomPageGetResponseError   `json:"errors"`
	Messages []AccountCustomPageGetResponseMessage `json:"messages"`
	Result   interface{}                           `json:"result"`
	// Whether the API call was successful
	Success AccountCustomPageGetResponseSuccess `json:"success"`
	JSON    accountCustomPageGetResponseJSON    `json:"-"`
}

// accountCustomPageGetResponseJSON contains the JSON metadata for the struct
// [AccountCustomPageGetResponse]
type accountCustomPageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountCustomPageGetResponseErrorJSON `json:"-"`
}

// accountCustomPageGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountCustomPageGetResponseError]
type accountCustomPageGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountCustomPageGetResponseMessageJSON `json:"-"`
}

// accountCustomPageGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountCustomPageGetResponseMessage]
type accountCustomPageGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomPageGetResponseSuccess bool

const (
	AccountCustomPageGetResponseSuccessTrue AccountCustomPageGetResponseSuccess = true
)

type AccountCustomPageUpdateResponse struct {
	Errors   []AccountCustomPageUpdateResponseError   `json:"errors"`
	Messages []AccountCustomPageUpdateResponseMessage `json:"messages"`
	Result   interface{}                              `json:"result"`
	// Whether the API call was successful
	Success AccountCustomPageUpdateResponseSuccess `json:"success"`
	JSON    accountCustomPageUpdateResponseJSON    `json:"-"`
}

// accountCustomPageUpdateResponseJSON contains the JSON metadata for the struct
// [AccountCustomPageUpdateResponse]
type accountCustomPageUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountCustomPageUpdateResponseErrorJSON `json:"-"`
}

// accountCustomPageUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountCustomPageUpdateResponseError]
type accountCustomPageUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountCustomPageUpdateResponseMessageJSON `json:"-"`
}

// accountCustomPageUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountCustomPageUpdateResponseMessage]
type accountCustomPageUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomPageUpdateResponseSuccess bool

const (
	AccountCustomPageUpdateResponseSuccessTrue AccountCustomPageUpdateResponseSuccess = true
)

type AccountCustomPageCustomPagesForAnAccountListCustomPagesResponse struct {
	Errors     []AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseError    `json:"errors"`
	Messages   []AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessage  `json:"messages"`
	Result     []interface{}                                                             `json:"result"`
	ResultInfo AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseSuccess `json:"success"`
	JSON    accountCustomPageCustomPagesForAnAccountListCustomPagesResponseJSON    `json:"-"`
}

// accountCustomPageCustomPagesForAnAccountListCustomPagesResponseJSON contains the
// JSON metadata for the struct
// [AccountCustomPageCustomPagesForAnAccountListCustomPagesResponse]
type accountCustomPageCustomPagesForAnAccountListCustomPagesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageCustomPagesForAnAccountListCustomPagesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountCustomPageCustomPagesForAnAccountListCustomPagesResponseErrorJSON `json:"-"`
}

// accountCustomPageCustomPagesForAnAccountListCustomPagesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseError]
type accountCustomPageCustomPagesForAnAccountListCustomPagesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessageJSON `json:"-"`
}

// accountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessage]
type accountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       accountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfoJSON `json:"-"`
}

// accountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfo]
type accountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseSuccess bool

const (
	AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseSuccessTrue AccountCustomPageCustomPagesForAnAccountListCustomPagesResponseSuccess = true
)

type AccountCustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[AccountCustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r AccountCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The custom page state.
type AccountCustomPageUpdateParamsState string

const (
	AccountCustomPageUpdateParamsStateDefault    AccountCustomPageUpdateParamsState = "default"
	AccountCustomPageUpdateParamsStateCustomized AccountCustomPageUpdateParamsState = "customized"
)
