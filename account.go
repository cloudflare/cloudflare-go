// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	return
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, identifier interface{}, body AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all accounts you have ownership or verified access to.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *AccountListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountGetResponse struct {
	Errors   []AccountGetResponseError   `json:"errors"`
	Messages []AccountGetResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success AccountGetResponseSuccess `json:"success"`
	JSON    accountGetResponseJSON    `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    accountGetResponseErrorJSON `json:"-"`
}

// accountGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountGetResponseError]
type accountGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    accountGetResponseMessageJSON `json:"-"`
}

// accountGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountGetResponseMessage]
type accountGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGetResponseSuccess bool

const (
	AccountGetResponseSuccessTrue AccountGetResponseSuccess = true
)

type AccountUpdateResponse struct {
	Errors   []AccountUpdateResponseError   `json:"errors"`
	Messages []AccountUpdateResponseMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success AccountUpdateResponseSuccess `json:"success"`
	JSON    accountUpdateResponseJSON    `json:"-"`
}

// accountUpdateResponseJSON contains the JSON metadata for the struct
// [AccountUpdateResponse]
type accountUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    accountUpdateResponseErrorJSON `json:"-"`
}

// accountUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountUpdateResponseError]
type accountUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    accountUpdateResponseMessageJSON `json:"-"`
}

// accountUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccountUpdateResponseMessage]
type accountUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountUpdateResponseSuccess bool

const (
	AccountUpdateResponseSuccessTrue AccountUpdateResponseSuccess = true
)

type AccountListResponse struct {
	Errors     []AccountListResponseError    `json:"errors"`
	Messages   []AccountListResponseMessage  `json:"messages"`
	Result     []interface{}                 `json:"result"`
	ResultInfo AccountListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountListResponseSuccess `json:"success"`
	JSON    accountListResponseJSON    `json:"-"`
}

// accountListResponseJSON contains the JSON metadata for the struct
// [AccountListResponse]
type accountListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    accountListResponseErrorJSON `json:"-"`
}

// accountListResponseErrorJSON contains the JSON metadata for the struct
// [AccountListResponseError]
type accountListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    accountListResponseMessageJSON `json:"-"`
}

// accountListResponseMessageJSON contains the JSON metadata for the struct
// [AccountListResponseMessage]
type accountListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                           `json:"total_count"`
	JSON       accountListResponseResultInfoJSON `json:"-"`
}

// accountListResponseResultInfoJSON contains the JSON metadata for the struct
// [AccountListResponseResultInfo]
type accountListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountListResponseSuccess bool

const (
	AccountListResponseSuccessTrue AccountListResponseSuccess = true
)

type AccountUpdateParams struct {
	// Account name
	Name param.Field[string] `json:"name,required"`
	// Account settings
	Settings param.Field[AccountUpdateParamsSettings] `json:"settings"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type AccountUpdateParamsSettings struct {
	// Specifies the default nameservers to be used for new zones added to this
	// account.
	//
	// - `cloudflare.standard` for Cloudflare-branded nameservers
	// - `custom.account` for account custom nameservers
	// - `custom.tenant` for tenant custom nameservers
	//
	// See
	// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
	// for more information.
	DefaultNameservers param.Field[AccountUpdateParamsSettingsDefaultNameservers] `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor param.Field[bool] `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default
	UseAccountCustomNsByDefault param.Field[bool] `json:"use_account_custom_ns_by_default"`
}

func (r AccountUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the default nameservers to be used for new zones added to this
// account.
//
// - `cloudflare.standard` for Cloudflare-branded nameservers
// - `custom.account` for account custom nameservers
// - `custom.tenant` for tenant custom nameservers
//
// See
// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
// for more information.
type AccountUpdateParamsSettingsDefaultNameservers string

const (
	AccountUpdateParamsSettingsDefaultNameserversCloudflareStandard AccountUpdateParamsSettingsDefaultNameservers = "cloudflare.standard"
	AccountUpdateParamsSettingsDefaultNameserversCustomAccount      AccountUpdateParamsSettingsDefaultNameservers = "custom.account"
	AccountUpdateParamsSettingsDefaultNameserversCustomTenant       AccountUpdateParamsSettingsDefaultNameservers = "custom.tenant"
)

type AccountListParams struct {
	// Direction to order results.
	Direction param.Field[AccountListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountListParamsDirection string

const (
	AccountListParamsDirectionAsc  AccountListParamsDirection = "asc"
	AccountListParamsDirectionDesc AccountListParamsDirection = "desc"
)
