// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

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
	Members *AccountMemberService
	Roles   *AccountRoleService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Members = NewAccountMemberService(opts...)
	r.Roles = NewAccountRoleService(opts...)
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, params AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
func (r *AccountService) Get(ctx context.Context, query AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Account struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Account settings
	Settings AccountSettings `json:"settings"`
	JSON     accountJSON     `json:"-"`
}

// accountJSON contains the JSON metadata for the struct [Account]
type accountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	CreatedOn   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Account) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Account settings
type AccountSettings struct {
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
	DefaultNameservers AccountSettingsDefaultNameservers `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
	UseAccountCustomNsByDefault bool                `json:"use_account_custom_ns_by_default"`
	JSON                        accountSettingsJSON `json:"-"`
}

// accountSettingsJSON contains the JSON metadata for the struct [AccountSettings]
type accountSettingsJSON struct {
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNsByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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
type AccountSettingsDefaultNameservers string

const (
	AccountSettingsDefaultNameserversCloudflareStandard AccountSettingsDefaultNameservers = "cloudflare.standard"
	AccountSettingsDefaultNameserversCustomAccount      AccountSettingsDefaultNameservers = "custom.account"
	AccountSettingsDefaultNameserversCustomTenant       AccountSettingsDefaultNameservers = "custom.tenant"
)

// Union satisfied by [AccountUpdateResponseUnknown] or [shared.UnionString].
type AccountUpdateResponse interface {
	ImplementsAccountUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

type AccountUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
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

type AccountUpdateResponseEnvelope struct {
	Errors   []AccountUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountUpdateResponseEnvelopeJSON    `json:"-"`
}

// accountUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountUpdateResponseEnvelope]
type accountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accountUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountUpdateResponseEnvelopeErrors]
type accountUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accountUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountUpdateResponseEnvelopeMessages]
type accountUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountUpdateResponseEnvelopeSuccess bool

const (
	AccountUpdateResponseEnvelopeSuccessTrue AccountUpdateResponseEnvelopeSuccess = true
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

type AccountGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

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
