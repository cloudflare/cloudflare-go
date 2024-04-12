// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/user"
	"github.com/tidwall/gjson"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options []option.RequestOption
	Members *MemberService
	Roles   *RoleService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Members = NewMemberService(opts...)
	r.Roles = NewRoleService(opts...)
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, params AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponseUnion, err error) {
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
func (r *AccountService) List(ctx context.Context, query AccountListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccountListResponse], err error) {
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
func (r *AccountService) ListAutoPaging(ctx context.Context, query AccountListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccountListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, query AccountGetParams, opts ...option.RequestOption) (res *AccountGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v", query.AccountID)
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
	UseAccountCustomNSByDefault bool                `json:"use_account_custom_ns_by_default"`
	JSON                        accountSettingsJSON `json:"-"`
}

// accountSettingsJSON contains the JSON metadata for the struct [AccountSettings]
type accountSettingsJSON struct {
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
	DefaultNameservers param.Field[AccountSettingsDefaultNameservers] `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor param.Field[bool] `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
	UseAccountCustomNSByDefault param.Field[bool] `json:"use_account_custom_ns_by_default"`
}

func (r AccountSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Role struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []user.Permission `json:"permissions,required"`
	JSON        roleJSON          `json:"-"`
}

// roleJSON contains the JSON metadata for the struct [Role]
type roleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Role) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleJSON) RawJSON() string {
	return r.raw
}

type User struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []UserRole  `json:"roles,required"`
	Status interface{} `json:"status,required"`
	User   UserUser    `json:"user,required"`
	JSON   userJSON    `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string               `json:"name,required"`
	Permissions UserRolesPermissions `json:"permissions,required"`
	JSON        userRoleJSON         `json:"-"`
}

// userRoleJSON contains the JSON metadata for the struct [UserRole]
type userRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRoleJSON) RawJSON() string {
	return r.raw
}

type UserRolesPermissions struct {
	Analytics    PermissionGrant          `json:"analytics"`
	Billing      PermissionGrant          `json:"billing"`
	CachePurge   PermissionGrant          `json:"cache_purge"`
	DNS          PermissionGrant          `json:"dns"`
	DNSRecords   PermissionGrant          `json:"dns_records"`
	Lb           PermissionGrant          `json:"lb"`
	Logs         PermissionGrant          `json:"logs"`
	Organization PermissionGrant          `json:"organization"`
	SSL          PermissionGrant          `json:"ssl"`
	WAF          PermissionGrant          `json:"waf"`
	ZoneSettings PermissionGrant          `json:"zone_settings"`
	Zones        PermissionGrant          `json:"zones"`
	JSON         userRolesPermissionsJSON `json:"-"`
}

// userRolesPermissionsJSON contains the JSON metadata for the struct
// [UserRolesPermissions]
type userRolesPermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRolesPermissionsJSON) RawJSON() string {
	return r.raw
}

type UserUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool         `json:"two_factor_authentication_enabled"`
	JSON                           userUserJSON `json:"-"`
}

// userUserJSON contains the JSON metadata for the struct [UserUser]
type userUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *UserUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserJSON) RawJSON() string {
	return r.raw
}

type UserParam struct {
	// Roles assigned to this member.
	Roles param.Field[[]UserRoleParam] `json:"roles,required"`
}

func (r UserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRoleParam struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r UserRoleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRolesPermissionsParam struct {
	Analytics    param.Field[PermissionGrantParam] `json:"analytics"`
	Billing      param.Field[PermissionGrantParam] `json:"billing"`
	CachePurge   param.Field[PermissionGrantParam] `json:"cache_purge"`
	DNS          param.Field[PermissionGrantParam] `json:"dns"`
	DNSRecords   param.Field[PermissionGrantParam] `json:"dns_records"`
	Lb           param.Field[PermissionGrantParam] `json:"lb"`
	Logs         param.Field[PermissionGrantParam] `json:"logs"`
	Organization param.Field[PermissionGrantParam] `json:"organization"`
	SSL          param.Field[PermissionGrantParam] `json:"ssl"`
	WAF          param.Field[PermissionGrantParam] `json:"waf"`
	ZoneSettings param.Field[PermissionGrantParam] `json:"zone_settings"`
	Zones        param.Field[PermissionGrantParam] `json:"zones"`
}

func (r UserRolesPermissionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUserParam struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
}

func (r UserUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [accounts.AccountUpdateResponseUnknown] or
// [shared.UnionString].
type AccountUpdateResponseUnion interface {
	ImplementsAccountsAccountUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountListResponse = interface{}

// Union satisfied by [accounts.AccountGetResponseUnknown] or [shared.UnionString].
type AccountGetResponseUnion interface {
	ImplementsAccountsAccountGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccountUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	Account   AccountParam             `json:"account,required"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Account)
}

type AccountUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   AccountUpdateResponseUnion `json:"result,required"`
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
		NestedFormat: apiquery.NestedQueryFormatBrackets,
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

type AccountGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type AccountGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   AccountGetResponseUnion `json:"result,required"`
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
