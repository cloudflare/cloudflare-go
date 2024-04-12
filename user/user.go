// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	AuditLogs     *AuditLogService
	Billing       *BillingService
	Invites       *InviteService
	Organizations *OrganizationService
	Subscriptions *SubscriptionService
	Tokens        *TokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.AuditLogs = NewAuditLogService(opts...)
	r.Billing = NewBillingService(opts...)
	r.Invites = NewInviteService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Tokens = NewTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Edit(ctx context.Context, body UserEditParams, opts ...option.RequestOption) (res *UserEditResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UserEditResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UserGetResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Permission = string

type Role struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []Permission `json:"permissions,required"`
	JSON        roleJSON     `json:"-"`
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
	Analytics    accounts.PermissionGrant `json:"analytics"`
	Billing      accounts.PermissionGrant `json:"billing"`
	CachePurge   accounts.PermissionGrant `json:"cache_purge"`
	DNS          accounts.PermissionGrant `json:"dns"`
	DNSRecords   accounts.PermissionGrant `json:"dns_records"`
	Lb           accounts.PermissionGrant `json:"lb"`
	Logs         accounts.PermissionGrant `json:"logs"`
	Organization accounts.PermissionGrant `json:"organization"`
	SSL          accounts.PermissionGrant `json:"ssl"`
	WAF          accounts.PermissionGrant `json:"waf"`
	ZoneSettings accounts.PermissionGrant `json:"zone_settings"`
	Zones        accounts.PermissionGrant `json:"zones"`
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
	Analytics    param.Field[accounts.PermissionGrantParam] `json:"analytics"`
	Billing      param.Field[accounts.PermissionGrantParam] `json:"billing"`
	CachePurge   param.Field[accounts.PermissionGrantParam] `json:"cache_purge"`
	DNS          param.Field[accounts.PermissionGrantParam] `json:"dns"`
	DNSRecords   param.Field[accounts.PermissionGrantParam] `json:"dns_records"`
	Lb           param.Field[accounts.PermissionGrantParam] `json:"lb"`
	Logs         param.Field[accounts.PermissionGrantParam] `json:"logs"`
	Organization param.Field[accounts.PermissionGrantParam] `json:"organization"`
	SSL          param.Field[accounts.PermissionGrantParam] `json:"ssl"`
	WAF          param.Field[accounts.PermissionGrantParam] `json:"waf"`
	ZoneSettings param.Field[accounts.PermissionGrantParam] `json:"zone_settings"`
	Zones        param.Field[accounts.PermissionGrantParam] `json:"zones"`
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

// Union satisfied by [user.UserEditResponseUnknown] or [shared.UnionString].
type UserEditResponseUnion interface {
	ImplementsUserUserEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [user.UserGetResponseUnknown] or [shared.UnionString].
type UserGetResponseUnion interface {
	ImplementsUserUserGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserEditParams struct {
	// The country in which the user lives.
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
	// User's telephone number
	Telephone param.Field[string] `json:"telephone"`
	// The zipcode or postal code where the user lives.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r UserEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UserEditResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success UserEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userEditResponseEnvelopeJSON    `json:"-"`
}

// userEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelope]
type userEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserEditResponseEnvelopeSuccess bool

const (
	UserEditResponseEnvelopeSuccessTrue UserEditResponseEnvelopeSuccess = true
)

func (r UserEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   UserGetResponseUnion  `json:"result,required"`
	// Whether the API call was successful
	Success UserGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userGetResponseEnvelopeJSON    `json:"-"`
}

// userGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelope]
type userGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserGetResponseEnvelopeSuccess bool

const (
	UserGetResponseEnvelopeSuccessTrue UserGetResponseEnvelopeSuccess = true
)

func (r UserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
