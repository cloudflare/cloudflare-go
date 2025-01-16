// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options       []option.RequestOption
	Members       *MemberService
	Roles         *RoleService
	Subscriptions *SubscriptionService
	Tokens        *TokenService
	Logs          *LogService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Members = NewMemberService(opts...)
	r.Roles = NewRoleService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Tokens = NewTokenService(opts...)
	r.Logs = NewLogService(opts...)
	return
}

// Create an account (only available for tenant admins at this time)
func (r *AccountService) New(ctx context.Context, body AccountNewParams, opts ...option.RequestOption) (res *Account, err error) {
	var env AccountNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, params AccountUpdateParams, opts ...option.RequestOption) (res *Account, err error) {
	var env AccountUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all accounts you have ownership or verified access to.
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Account], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Account] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Delete a specific account (only available for tenant admins at this time). This
// is a permanent operation that will delete any zones or other resources under the
// account
func (r *AccountService) Delete(ctx context.Context, body AccountDeleteParams, opts ...option.RequestOption) (res *AccountDeleteResponse, err error) {
	var env AccountDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, query AccountGetParams, opts ...option.RequestOption) (res *Account, err error) {
	var env AccountGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r accountJSON) RawJSON() string {
	return r.raw
}

// Account settings
type AccountSettings struct {
	// Sets an abuse contact email to notify for abuse reports.
	AbuseContactEmail string `json:"abuse_contact_email"`
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
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	DefaultNameservers AccountSettingsDefaultNameservers `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	UseAccountCustomNSByDefault bool                `json:"use_account_custom_ns_by_default"`
	JSON                        accountSettingsJSON `json:"-"`
}

// accountSettingsJSON contains the JSON metadata for the struct [AccountSettings]
type accountSettingsJSON struct {
	AbuseContactEmail           apijson.Field
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNSByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsJSON) RawJSON() string {
	return r.raw
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
//
// Deprecated in favor of
// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
type AccountSettingsDefaultNameservers string

const (
	AccountSettingsDefaultNameserversCloudflareStandard AccountSettingsDefaultNameservers = "cloudflare.standard"
	AccountSettingsDefaultNameserversCustomAccount      AccountSettingsDefaultNameservers = "custom.account"
	AccountSettingsDefaultNameserversCustomTenant       AccountSettingsDefaultNameservers = "custom.tenant"
)

func (r AccountSettingsDefaultNameservers) IsKnown() bool {
	switch r {
	case AccountSettingsDefaultNameserversCloudflareStandard, AccountSettingsDefaultNameserversCustomAccount, AccountSettingsDefaultNameserversCustomTenant:
		return true
	}
	return false
}

type AccountParam struct {
	// Account name
	Name param.Field[string] `json:"name,required"`
	// Account settings
	Settings param.Field[AccountSettingsParam] `json:"settings"`
}

func (r AccountParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type AccountSettingsParam struct {
	// Sets an abuse contact email to notify for abuse reports.
	AbuseContactEmail param.Field[string] `json:"abuse_contact_email"`
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
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	DefaultNameservers param.Field[AccountSettingsDefaultNameservers] `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor param.Field[bool] `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of
	// [DNS Settings](https://developers.cloudflare.com/api/operations/dns-settings-for-an-account-update-dns-settings).
	UseAccountCustomNSByDefault param.Field[bool] `json:"use_account_custom_ns_by_default"`
}

func (r AccountSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDeleteResponse struct {
	// Identifier
	ID   string                    `json:"id,required"`
	JSON accountDeleteResponseJSON `json:"-"`
}

// accountDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDeleteResponse]
type accountDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountNewParams struct {
	// Account name
	Name param.Field[string] `json:"name,required"`
	// the type of account being created. For self-serve customers, use standard. for
	// enterprise customers, use enterprise.
	Type param.Field[AccountNewParamsType] `json:"type,required"`
	// information related to the tenant unit, and optionally, an id of the unit to
	// create the account on. see
	// https://developers.cloudflare.com/tenant/how-to/manage-accounts/
	Unit param.Field[AccountNewParamsUnit] `json:"unit"`
}

func (r AccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// the type of account being created. For self-serve customers, use standard. for
// enterprise customers, use enterprise.
type AccountNewParamsType string

const (
	AccountNewParamsTypeStandard   AccountNewParamsType = "standard"
	AccountNewParamsTypeEnterprise AccountNewParamsType = "enterprise"
)

func (r AccountNewParamsType) IsKnown() bool {
	switch r {
	case AccountNewParamsTypeStandard, AccountNewParamsTypeEnterprise:
		return true
	}
	return false
}

// information related to the tenant unit, and optionally, an id of the unit to
// create the account on. see
// https://developers.cloudflare.com/tenant/how-to/manage-accounts/
type AccountNewParamsUnit struct {
	// Tenant unit ID
	ID param.Field[string] `json:"id"`
}

func (r AccountNewParamsUnit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Account                           `json:"result"`
	JSON    accountNewResponseEnvelopeJSON    `json:"-"`
}

// accountNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountNewResponseEnvelope]
type accountNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountNewResponseEnvelopeSuccess bool

const (
	AccountNewResponseEnvelopeSuccessTrue AccountNewResponseEnvelopeSuccess = true
)

func (r AccountNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccountUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	Account   AccountParam        `json:"account,required"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Account)
}

type AccountUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Account                              `json:"result"`
	JSON    accountUpdateResponseEnvelopeJSON    `json:"-"`
}

// accountUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountUpdateResponseEnvelope]
type accountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountUpdateResponseEnvelopeSuccess bool

const (
	AccountUpdateResponseEnvelopeSuccessTrue AccountUpdateResponseEnvelopeSuccess = true
)

func (r AccountUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccountListParams struct {
	// Direction to order results.
	Direction param.Field[AccountListParamsDirection] `query:"direction"`
	// Name of the account.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountListParams]'s query parameters as `url.Values`.
func (r AccountListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type AccountListParamsDirection string

const (
	AccountListParamsDirectionAsc  AccountListParamsDirection = "asc"
	AccountListParamsDirectionDesc AccountListParamsDirection = "desc"
)

func (r AccountListParamsDirection) IsKnown() bool {
	switch r {
	case AccountListParamsDirectionAsc, AccountListParamsDirectionDesc:
		return true
	}
	return false
}

type AccountDeleteParams struct {
	// The account ID of the account to be deleted
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccountDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccountDeleteResponse                `json:"result,nullable"`
	JSON    accountDeleteResponseEnvelopeJSON    `json:"-"`
}

// accountDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountDeleteResponseEnvelope]
type accountDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountDeleteResponseEnvelopeSuccess bool

const (
	AccountDeleteResponseEnvelopeSuccessTrue AccountDeleteResponseEnvelopeSuccess = true
)

func (r AccountDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccountGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccountGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AccountGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Account                           `json:"result"`
	JSON    accountGetResponseEnvelopeJSON    `json:"-"`
}

// accountGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountGetResponseEnvelope]
type accountGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGetResponseEnvelopeSuccess bool

const (
	AccountGetResponseEnvelopeSuccessTrue AccountGetResponseEnvelopeSuccess = true
)

func (r AccountGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccountGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
