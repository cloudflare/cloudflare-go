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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMemberService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMemberService] method
// instead.
type AccountMemberService struct {
	Options []option.RequestOption
}

// NewAccountMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMemberService(opts ...option.RequestOption) (r *AccountMemberService) {
	r = &AccountMemberService{}
	r.Options = opts
	return
}

// Add a user to the list of members for this account.
func (r *AccountMemberService) New(ctx context.Context, accountID interface{}, body AccountMemberNewParams, opts ...option.RequestOption) (res *AccountMemberNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *AccountMemberService) List(ctx context.Context, accountID interface{}, query AccountMemberListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[AccountMemberListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%v/members", accountID)
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

// List all members of an account.
func (r *AccountMemberService) ListAutoPaging(ctx context.Context, accountID interface{}, query AccountMemberListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[AccountMemberListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountID, query, opts...))
}

// Remove a member from an account.
func (r *AccountMemberService) Delete(ctx context.Context, accountID interface{}, memberID string, opts ...option.RequestOption) (res *AccountMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, accountID interface{}, memberID string, opts ...option.RequestOption) (res *AccountMemberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *AccountMemberService) Replace(ctx context.Context, accountID interface{}, memberID string, body AccountMemberReplaceParams, opts ...option.RequestOption) (res *AccountMemberReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountMemberNewResponse struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberNewResponseRole `json:"roles,required"`
	Status interface{}                    `json:"status,required"`
	User   AccountMemberNewResponseUser   `json:"user,required"`
	// The unique activation code for the account membership.
	Code string                       `json:"code"`
	JSON accountMemberNewResponseJSON `json:"-"`
}

// accountMemberNewResponseJSON contains the JSON metadata for the struct
// [AccountMemberNewResponse]
type accountMemberNewResponseJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                   `json:"name,required"`
	Permissions AccountMemberNewResponseRolesPermissions `json:"permissions,required"`
	JSON        accountMemberNewResponseRoleJSON         `json:"-"`
}

// accountMemberNewResponseRoleJSON contains the JSON metadata for the struct
// [AccountMemberNewResponseRole]
type accountMemberNewResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissions struct {
	Analytics    AccountMemberNewResponseRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberNewResponseRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberNewResponseRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberNewResponseRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberNewResponseRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberNewResponseRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberNewResponseRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberNewResponseRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberNewResponseRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberNewResponseRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberNewResponseRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberNewResponseRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberNewResponseRolesPermissionsJSON         `json:"-"`
}

// accountMemberNewResponseRolesPermissionsJSON contains the JSON metadata for the
// struct [AccountMemberNewResponseRolesPermissions]
type accountMemberNewResponseRolesPermissionsJSON struct {
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

func (r *AccountMemberNewResponseRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsAnalytics struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsAnalyticsJSON contains the JSON metadata
// for the struct [AccountMemberNewResponseRolesPermissionsAnalytics]
type accountMemberNewResponseRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsBilling struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsBillingJSON contains the JSON metadata
// for the struct [AccountMemberNewResponseRolesPermissionsBilling]
type accountMemberNewResponseRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsCachePurge struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsCachePurgeJSON contains the JSON
// metadata for the struct [AccountMemberNewResponseRolesPermissionsCachePurge]
type accountMemberNewResponseRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsDNS struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsDNSJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsDNS]
type accountMemberNewResponseRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsDNSRecords struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct [AccountMemberNewResponseRolesPermissionsDNSRecords]
type accountMemberNewResponseRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsLb struct {
	Read  bool                                           `json:"read"`
	Write bool                                           `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsLbJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsLbJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsLb]
type accountMemberNewResponseRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsLogs struct {
	Read  bool                                             `json:"read"`
	Write bool                                             `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsLogsJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsLogs]
type accountMemberNewResponseRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsOrganization struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsOrganizationJSON contains the JSON
// metadata for the struct [AccountMemberNewResponseRolesPermissionsOrganization]
type accountMemberNewResponseRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsSSL struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsSSLJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsSSL]
type accountMemberNewResponseRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsWAF struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsWAFJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsWAF]
type accountMemberNewResponseRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsZoneSettings struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct [AccountMemberNewResponseRolesPermissionsZoneSettings]
type accountMemberNewResponseRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseRolesPermissionsZones struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  accountMemberNewResponseRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberNewResponseRolesPermissionsZonesJSON contains the JSON metadata for
// the struct [AccountMemberNewResponseRolesPermissionsZones]
type accountMemberNewResponseRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseUser struct {
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
	TwoFactorAuthenticationEnabled bool                             `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberNewResponseUserJSON `json:"-"`
}

// accountMemberNewResponseUserJSON contains the JSON metadata for the struct
// [AccountMemberNewResponseUser]
type accountMemberNewResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberNewResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []AccountMemberListResponseRole `json:"roles,required"`
	// A member's status in the organization.
	Status AccountMemberListResponseStatus `json:"status,required"`
	JSON   accountMemberListResponseJSON   `json:"-"`
}

// accountMemberListResponseJSON contains the JSON metadata for the struct
// [AccountMemberListResponse]
type accountMemberListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberListResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                          `json:"permissions,required"`
	JSON        accountMemberListResponseRoleJSON `json:"-"`
}

// accountMemberListResponseRoleJSON contains the JSON metadata for the struct
// [AccountMemberListResponseRole]
type accountMemberListResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A member's status in the organization.
type AccountMemberListResponseStatus string

const (
	AccountMemberListResponseStatusAccepted AccountMemberListResponseStatus = "accepted"
	AccountMemberListResponseStatusInvited  AccountMemberListResponseStatus = "invited"
)

type AccountMemberDeleteResponse struct {
	// Identifier
	ID   string                          `json:"id,required"`
	JSON accountMemberDeleteResponseJSON `json:"-"`
}

// accountMemberDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMemberDeleteResponse]
type accountMemberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponse struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberGetResponseRole `json:"roles,required"`
	Status interface{}                    `json:"status,required"`
	User   AccountMemberGetResponseUser   `json:"user,required"`
	JSON   accountMemberGetResponseJSON   `json:"-"`
}

// accountMemberGetResponseJSON contains the JSON metadata for the struct
// [AccountMemberGetResponse]
type accountMemberGetResponseJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                   `json:"name,required"`
	Permissions AccountMemberGetResponseRolesPermissions `json:"permissions,required"`
	JSON        accountMemberGetResponseRoleJSON         `json:"-"`
}

// accountMemberGetResponseRoleJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseRole]
type accountMemberGetResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissions struct {
	Analytics    AccountMemberGetResponseRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberGetResponseRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberGetResponseRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberGetResponseRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberGetResponseRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberGetResponseRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberGetResponseRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberGetResponseRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberGetResponseRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberGetResponseRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberGetResponseRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberGetResponseRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberGetResponseRolesPermissionsJSON         `json:"-"`
}

// accountMemberGetResponseRolesPermissionsJSON contains the JSON metadata for the
// struct [AccountMemberGetResponseRolesPermissions]
type accountMemberGetResponseRolesPermissionsJSON struct {
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

func (r *AccountMemberGetResponseRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsAnalytics struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsAnalyticsJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseRolesPermissionsAnalytics]
type accountMemberGetResponseRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsBilling struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsBillingJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseRolesPermissionsBilling]
type accountMemberGetResponseRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsCachePurge struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsCachePurgeJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseRolesPermissionsCachePurge]
type accountMemberGetResponseRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsDNS struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsDNSJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsDNS]
type accountMemberGetResponseRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsDNSRecords struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseRolesPermissionsDNSRecords]
type accountMemberGetResponseRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsLb struct {
	Read  bool                                           `json:"read"`
	Write bool                                           `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsLbJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsLbJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsLb]
type accountMemberGetResponseRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsLogs struct {
	Read  bool                                             `json:"read"`
	Write bool                                             `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsLogsJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsLogs]
type accountMemberGetResponseRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsOrganization struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsOrganizationJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseRolesPermissionsOrganization]
type accountMemberGetResponseRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsSSL struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsSSLJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsSSL]
type accountMemberGetResponseRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsWAF struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsWAFJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsWAF]
type accountMemberGetResponseRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsZoneSettings struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseRolesPermissionsZoneSettings]
type accountMemberGetResponseRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseRolesPermissionsZones struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  accountMemberGetResponseRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberGetResponseRolesPermissionsZonesJSON contains the JSON metadata for
// the struct [AccountMemberGetResponseRolesPermissionsZones]
type accountMemberGetResponseRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseUser struct {
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
	TwoFactorAuthenticationEnabled bool                             `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberGetResponseUserJSON `json:"-"`
}

// accountMemberGetResponseUserJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseUser]
type accountMemberGetResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponse struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberReplaceResponseRole `json:"roles,required"`
	Status interface{}                        `json:"status,required"`
	User   AccountMemberReplaceResponseUser   `json:"user,required"`
	JSON   accountMemberReplaceResponseJSON   `json:"-"`
}

// accountMemberReplaceResponseJSON contains the JSON metadata for the struct
// [AccountMemberReplaceResponse]
type accountMemberReplaceResponseJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                       `json:"name,required"`
	Permissions AccountMemberReplaceResponseRolesPermissions `json:"permissions,required"`
	JSON        accountMemberReplaceResponseRoleJSON         `json:"-"`
}

// accountMemberReplaceResponseRoleJSON contains the JSON metadata for the struct
// [AccountMemberReplaceResponseRole]
type accountMemberReplaceResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissions struct {
	Analytics    AccountMemberReplaceResponseRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberReplaceResponseRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberReplaceResponseRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberReplaceResponseRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberReplaceResponseRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberReplaceResponseRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberReplaceResponseRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberReplaceResponseRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberReplaceResponseRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberReplaceResponseRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberReplaceResponseRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberReplaceResponseRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberReplaceResponseRolesPermissionsJSON         `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsJSON contains the JSON metadata for
// the struct [AccountMemberReplaceResponseRolesPermissions]
type accountMemberReplaceResponseRolesPermissionsJSON struct {
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

func (r *AccountMemberReplaceResponseRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsAnalytics struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsAnalyticsJSON contains the JSON
// metadata for the struct [AccountMemberReplaceResponseRolesPermissionsAnalytics]
type accountMemberReplaceResponseRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsBilling struct {
	Read  bool                                                    `json:"read"`
	Write bool                                                    `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsBillingJSON contains the JSON
// metadata for the struct [AccountMemberReplaceResponseRolesPermissionsBilling]
type accountMemberReplaceResponseRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsCachePurge struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsCachePurgeJSON contains the JSON
// metadata for the struct [AccountMemberReplaceResponseRolesPermissionsCachePurge]
type accountMemberReplaceResponseRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsDNS struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsDNSJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsDNS]
type accountMemberReplaceResponseRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsDNSRecords struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct [AccountMemberReplaceResponseRolesPermissionsDNSRecords]
type accountMemberReplaceResponseRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsLb struct {
	Read  bool                                               `json:"read"`
	Write bool                                               `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsLbJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsLbJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsLb]
type accountMemberReplaceResponseRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsLogs struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsLogsJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsLogs]
type accountMemberReplaceResponseRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsOrganization struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsOrganizationJSON contains the JSON
// metadata for the struct
// [AccountMemberReplaceResponseRolesPermissionsOrganization]
type accountMemberReplaceResponseRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsSSL struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsSSLJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsSSL]
type accountMemberReplaceResponseRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsWAF struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsWAFJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsWAF]
type accountMemberReplaceResponseRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsZoneSettings struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct
// [AccountMemberReplaceResponseRolesPermissionsZoneSettings]
type accountMemberReplaceResponseRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseRolesPermissionsZones struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberReplaceResponseRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberReplaceResponseRolesPermissionsZonesJSON contains the JSON metadata
// for the struct [AccountMemberReplaceResponseRolesPermissionsZones]
type accountMemberReplaceResponseRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseUser struct {
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
	TwoFactorAuthenticationEnabled bool                                 `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberReplaceResponseUserJSON `json:"-"`
}

// accountMemberReplaceResponseUserJSON contains the JSON metadata for the struct
// [AccountMemberReplaceResponseUser]
type accountMemberReplaceResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                     `json:"roles,required"`
	Status param.Field[AccountMemberNewParamsStatus] `json:"status"`
}

func (r AccountMemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberNewParamsStatus string

const (
	AccountMemberNewParamsStatusAccepted AccountMemberNewParamsStatus = "accepted"
	AccountMemberNewParamsStatusPending  AccountMemberNewParamsStatus = "pending"
)

type AccountMemberNewResponseEnvelope struct {
	Errors   []AccountMemberNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMemberNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountMemberNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountMemberNewResponseEnvelopeJSON    `json:"-"`
}

// accountMemberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountMemberNewResponseEnvelope]
type accountMemberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountMemberNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountMemberNewResponseEnvelopeErrors]
type accountMemberNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberNewResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMemberNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountMemberNewResponseEnvelopeMessages]
type accountMemberNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberNewResponseEnvelopeSuccess bool

const (
	AccountMemberNewResponseEnvelopeSuccessTrue AccountMemberNewResponseEnvelopeSuccess = true
)

type AccountMemberListParams struct {
	// Direction to order results.
	Direction param.Field[AccountMemberListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountMemberListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[AccountMemberListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountMemberListParams]'s query parameters as
// `url.Values`.
func (r AccountMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountMemberListParamsDirection string

const (
	AccountMemberListParamsDirectionAsc  AccountMemberListParamsDirection = "asc"
	AccountMemberListParamsDirectionDesc AccountMemberListParamsDirection = "desc"
)

// Field to order results by.
type AccountMemberListParamsOrder string

const (
	AccountMemberListParamsOrderUserFirstName AccountMemberListParamsOrder = "user.first_name"
	AccountMemberListParamsOrderUserLastName  AccountMemberListParamsOrder = "user.last_name"
	AccountMemberListParamsOrderUserEmail     AccountMemberListParamsOrder = "user.email"
	AccountMemberListParamsOrderStatus        AccountMemberListParamsOrder = "status"
)

// A member's status in the account.
type AccountMemberListParamsStatus string

const (
	AccountMemberListParamsStatusAccepted AccountMemberListParamsStatus = "accepted"
	AccountMemberListParamsStatusPending  AccountMemberListParamsStatus = "pending"
	AccountMemberListParamsStatusRejected AccountMemberListParamsStatus = "rejected"
)

type AccountMemberDeleteResponseEnvelope struct {
	Errors   []AccountMemberDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMemberDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountMemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountMemberDeleteResponseEnvelopeJSON    `json:"-"`
}

// accountMemberDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountMemberDeleteResponseEnvelope]
type accountMemberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountMemberDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountMemberDeleteResponseEnvelopeErrors]
type accountMemberDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountMemberDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccountMemberDeleteResponseEnvelopeMessages]
type accountMemberDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberDeleteResponseEnvelopeSuccess bool

const (
	AccountMemberDeleteResponseEnvelopeSuccessTrue AccountMemberDeleteResponseEnvelopeSuccess = true
)

type AccountMemberGetResponseEnvelope struct {
	Errors   []AccountMemberGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMemberGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountMemberGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountMemberGetResponseEnvelopeJSON    `json:"-"`
}

// accountMemberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseEnvelope]
type accountMemberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountMemberGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountMemberGetResponseEnvelopeErrors]
type accountMemberGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountMemberGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountMemberGetResponseEnvelopeMessages]
type accountMemberGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberGetResponseEnvelopeSuccess bool

const (
	AccountMemberGetResponseEnvelopeSuccessTrue AccountMemberGetResponseEnvelopeSuccess = true
)

type AccountMemberReplaceParams struct {
	// Roles assigned to this member.
	Roles param.Field[[]AccountMemberReplaceParamsRole] `json:"roles,required"`
}

func (r AccountMemberReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRole struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r AccountMemberReplaceParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissions struct {
	Analytics    param.Field[AccountMemberReplaceParamsRolesPermissionsAnalytics]    `json:"analytics"`
	Billing      param.Field[AccountMemberReplaceParamsRolesPermissionsBilling]      `json:"billing"`
	CachePurge   param.Field[AccountMemberReplaceParamsRolesPermissionsCachePurge]   `json:"cache_purge"`
	DNS          param.Field[AccountMemberReplaceParamsRolesPermissionsDNS]          `json:"dns"`
	DNSRecords   param.Field[AccountMemberReplaceParamsRolesPermissionsDNSRecords]   `json:"dns_records"`
	Lb           param.Field[AccountMemberReplaceParamsRolesPermissionsLb]           `json:"lb"`
	Logs         param.Field[AccountMemberReplaceParamsRolesPermissionsLogs]         `json:"logs"`
	Organization param.Field[AccountMemberReplaceParamsRolesPermissionsOrganization] `json:"organization"`
	SSL          param.Field[AccountMemberReplaceParamsRolesPermissionsSSL]          `json:"ssl"`
	WAF          param.Field[AccountMemberReplaceParamsRolesPermissionsWAF]          `json:"waf"`
	ZoneSettings param.Field[AccountMemberReplaceParamsRolesPermissionsZoneSettings] `json:"zone_settings"`
	Zones        param.Field[AccountMemberReplaceParamsRolesPermissionsZones]        `json:"zones"`
}

func (r AccountMemberReplaceParamsRolesPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsAnalytics struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsBilling struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsCachePurge struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsCachePurge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsDNS struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsDNSRecords struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsDNSRecords) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsLb struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsLb) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsLogs struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsOrganization struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsSSL struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsWAF struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsZoneSettings struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsZoneSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceParamsRolesPermissionsZones struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberReplaceParamsRolesPermissionsZones) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberReplaceResponseEnvelope struct {
	Errors   []AccountMemberReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMemberReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountMemberReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountMemberReplaceResponseEnvelopeJSON    `json:"-"`
}

// accountMemberReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountMemberReplaceResponseEnvelope]
type accountMemberReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountMemberReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccountMemberReplaceResponseEnvelopeErrors]
type accountMemberReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountMemberReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccountMemberReplaceResponseEnvelopeMessages]
type accountMemberReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberReplaceResponseEnvelopeSuccess bool

const (
	AccountMemberReplaceResponseEnvelopeSuccessTrue AccountMemberReplaceResponseEnvelopeSuccess = true
)
