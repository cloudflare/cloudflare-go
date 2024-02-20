// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// List all accounts you have ownership or verified access to.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[AccountListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
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

// List all accounts you have ownership or verified access to.
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[AccountListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing account.
func (r *AccountService) Replace(ctx context.Context, accountID interface{}, body AccountReplaceParams, opts ...option.RequestOption) (res *AccountReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountListResponse = interface{}

// Union satisfied by [AccountGetResponseUnknown] or [shared.UnionString].
type AccountGetResponse interface {
	ImplementsAccountGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [AccountReplaceResponseUnknown] or [shared.UnionString].
type AccountReplaceResponse interface {
	ImplementsAccountReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

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

type AccountGetResponseEnvelope struct {
	Errors   []AccountGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountGetResponseEnvelopeJSON    `json:"-"`
}

// accountGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountGetResponseEnvelope]
type accountGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accountGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AccountGetResponseEnvelopeErrors]
type accountGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accountGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AccountGetResponseEnvelopeMessages]
type accountGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGetResponseEnvelopeSuccess bool

const (
	AccountGetResponseEnvelopeSuccessTrue AccountGetResponseEnvelopeSuccess = true
)

type AccountReplaceParams struct {
	// Account name
	Name param.Field[string] `json:"name,required"`
	// Account settings
	Settings param.Field[AccountReplaceParamsSettings] `json:"settings"`
}

func (r AccountReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type AccountReplaceParamsSettings struct {
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
	DefaultNameservers param.Field[AccountReplaceParamsSettingsDefaultNameservers] `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor param.Field[bool] `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
	UseAccountCustomNsByDefault param.Field[bool] `json:"use_account_custom_ns_by_default"`
}

func (r AccountReplaceParamsSettings) MarshalJSON() (data []byte, err error) {
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
type AccountReplaceParamsSettingsDefaultNameservers string

const (
	AccountReplaceParamsSettingsDefaultNameserversCloudflareStandard AccountReplaceParamsSettingsDefaultNameservers = "cloudflare.standard"
	AccountReplaceParamsSettingsDefaultNameserversCustomAccount      AccountReplaceParamsSettingsDefaultNameservers = "custom.account"
	AccountReplaceParamsSettingsDefaultNameserversCustomTenant       AccountReplaceParamsSettingsDefaultNameservers = "custom.tenant"
)

type AccountReplaceResponseEnvelope struct {
	Errors   []AccountReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountReplaceResponseEnvelopeJSON    `json:"-"`
}

// accountReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountReplaceResponseEnvelope]
type accountReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountReplaceResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accountReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountReplaceResponseEnvelopeErrors]
type accountReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountReplaceResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accountReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountReplaceResponseEnvelopeMessages]
type accountReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountReplaceResponseEnvelopeSuccess bool

const (
	AccountReplaceResponseEnvelopeSuccessTrue AccountReplaceResponseEnvelopeSuccess = true
)
