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

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, accountIdentifier interface{}, identifier string, opts ...option.RequestOption) (res *AccountMemberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *AccountMemberService) Update(ctx context.Context, accountIdentifier interface{}, identifier string, body AccountMemberUpdateParams, opts ...option.RequestOption) (res *AccountMemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a member from an account.
func (r *AccountMemberService) Delete(ctx context.Context, accountIdentifier interface{}, identifier string, opts ...option.RequestOption) (res *AccountMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add a user to the list of members for this account.
func (r *AccountMemberService) AccountMembersAddMember(ctx context.Context, accountIdentifier interface{}, body AccountMemberAccountMembersAddMemberParams, opts ...option.RequestOption) (res *AccountMemberAccountMembersAddMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberAccountMembersAddMemberResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *AccountMemberService) AccountMembersListMembers(ctx context.Context, accountIdentifier interface{}, query AccountMemberAccountMembersListMembersParams, opts ...option.RequestOption) (res *[]AccountMemberAccountMembersListMembersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberAccountMembersListMembersResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type AccountMemberUpdateResponse struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberUpdateResponseRole `json:"roles,required"`
	Status interface{}                       `json:"status,required"`
	User   AccountMemberUpdateResponseUser   `json:"user,required"`
	JSON   accountMemberUpdateResponseJSON   `json:"-"`
}

// accountMemberUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponse]
type accountMemberUpdateResponseJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                      `json:"name,required"`
	Permissions AccountMemberUpdateResponseRolesPermissions `json:"permissions,required"`
	JSON        accountMemberUpdateResponseRoleJSON         `json:"-"`
}

// accountMemberUpdateResponseRoleJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponseRole]
type accountMemberUpdateResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissions struct {
	Analytics    AccountMemberUpdateResponseRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberUpdateResponseRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberUpdateResponseRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberUpdateResponseRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberUpdateResponseRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberUpdateResponseRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberUpdateResponseRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberUpdateResponseRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberUpdateResponseRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberUpdateResponseRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberUpdateResponseRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberUpdateResponseRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberUpdateResponseRolesPermissionsJSON         `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsJSON contains the JSON metadata for
// the struct [AccountMemberUpdateResponseRolesPermissions]
type accountMemberUpdateResponseRolesPermissionsJSON struct {
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

func (r *AccountMemberUpdateResponseRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsAnalytics struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsAnalyticsJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseRolesPermissionsAnalytics]
type accountMemberUpdateResponseRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsBilling struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsBillingJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseRolesPermissionsBilling]
type accountMemberUpdateResponseRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsCachePurge struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsCachePurgeJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseRolesPermissionsCachePurge]
type accountMemberUpdateResponseRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsDNS struct {
	Read  bool                                               `json:"read"`
	Write bool                                               `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsDNSJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseRolesPermissionsDNS]
type accountMemberUpdateResponseRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsDNSRecords struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseRolesPermissionsDNSRecords]
type accountMemberUpdateResponseRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsLb struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsLbJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsLbJSON contains the JSON metadata for
// the struct [AccountMemberUpdateResponseRolesPermissionsLb]
type accountMemberUpdateResponseRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsLogs struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsLogsJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseRolesPermissionsLogs]
type accountMemberUpdateResponseRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsOrganization struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsOrganizationJSON contains the JSON
// metadata for the struct
// [AccountMemberUpdateResponseRolesPermissionsOrganization]
type accountMemberUpdateResponseRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsSSL struct {
	Read  bool                                               `json:"read"`
	Write bool                                               `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsSSLJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseRolesPermissionsSSL]
type accountMemberUpdateResponseRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsWAF struct {
	Read  bool                                               `json:"read"`
	Write bool                                               `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsWAFJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseRolesPermissionsWAF]
type accountMemberUpdateResponseRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsZoneSettings struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct
// [AccountMemberUpdateResponseRolesPermissionsZoneSettings]
type accountMemberUpdateResponseRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseRolesPermissionsZones struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  accountMemberUpdateResponseRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberUpdateResponseRolesPermissionsZonesJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseRolesPermissionsZones]
type accountMemberUpdateResponseRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseUser struct {
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
	TwoFactorAuthenticationEnabled bool                                `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberUpdateResponseUserJSON `json:"-"`
}

// accountMemberUpdateResponseUserJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponseUser]
type accountMemberUpdateResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

type AccountMemberAccountMembersAddMemberResponse struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// The unique activation code for the account membership.
	Code string `json:"code"`
	// Roles assigned to this member.
	Roles  []AccountMemberAccountMembersAddMemberResponseRole `json:"roles"`
	Status interface{}                                        `json:"status"`
	User   AccountMemberAccountMembersAddMemberResponseUser   `json:"user"`
	JSON   accountMemberAccountMembersAddMemberResponseJSON   `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseJSON contains the JSON metadata for
// the struct [AccountMemberAccountMembersAddMemberResponse]
type accountMemberAccountMembersAddMemberResponseJSON struct {
	ID          apijson.Field
	Code        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                                       `json:"name,required"`
	Permissions AccountMemberAccountMembersAddMemberResponseRolesPermissions `json:"permissions,required"`
	JSON        accountMemberAccountMembersAddMemberResponseRoleJSON         `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRoleJSON contains the JSON metadata
// for the struct [AccountMemberAccountMembersAddMemberResponseRole]
type accountMemberAccountMembersAddMemberResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissions struct {
	Analytics    AccountMemberAccountMembersAddMemberResponseRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberAccountMembersAddMemberResponseRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberAccountMembersAddMemberResponseRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberAccountMembersAddMemberResponseRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberAccountMembersAddMemberResponseRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberAccountMembersAddMemberResponseRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberAccountMembersAddMemberResponseRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberAccountMembersAddMemberResponseRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberAccountMembersAddMemberResponseRolesPermissionsJSON         `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissions]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsJSON struct {
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

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsAnalytics struct {
	Read  bool                                                                      `json:"read"`
	Write bool                                                                      `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsAnalyticsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsAnalytics]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsBilling struct {
	Read  bool                                                                    `json:"read"`
	Write bool                                                                    `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsBillingJSON contains
// the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsBilling]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurge struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurgeJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurge]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNS struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNS]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecords struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecordsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecords]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsLb struct {
	Read  bool                                                               `json:"read"`
	Write bool                                                               `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsLbJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsLbJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsLb]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsLogs struct {
	Read  bool                                                                 `json:"read"`
	Write bool                                                                 `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsLogsJSON contains
// the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsLogs]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsOrganization struct {
	Read  bool                                                                         `json:"read"`
	Write bool                                                                         `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsOrganizationJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsOrganization]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsSSL struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsSSLJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsSSL]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsWAF struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsWAFJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsWAF]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettings struct {
	Read  bool                                                                         `json:"read"`
	Write bool                                                                         `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettingsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettings]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseRolesPermissionsZones struct {
	Read  bool                                                                  `json:"read"`
	Write bool                                                                  `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseRolesPermissionsZonesJSON contains
// the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseRolesPermissionsZones]
type accountMemberAccountMembersAddMemberResponseRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseUser struct {
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
	TwoFactorAuthenticationEnabled bool                                                 `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberAccountMembersAddMemberResponseUserJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseUserJSON contains the JSON metadata
// for the struct [AccountMemberAccountMembersAddMemberResponseUser]
type accountMemberAccountMembersAddMemberResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersListMembersResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []AccountMemberAccountMembersListMembersResponseRole `json:"roles,required"`
	// A member's status in the organization.
	Status AccountMemberAccountMembersListMembersResponseStatus `json:"status,required"`
	JSON   accountMemberAccountMembersListMembersResponseJSON   `json:"-"`
}

// accountMemberAccountMembersListMembersResponseJSON contains the JSON metadata
// for the struct [AccountMemberAccountMembersListMembersResponse]
type accountMemberAccountMembersListMembersResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersListMembersResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                               `json:"permissions,required"`
	JSON        accountMemberAccountMembersListMembersResponseRoleJSON `json:"-"`
}

// accountMemberAccountMembersListMembersResponseRoleJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersListMembersResponseRole]
type accountMemberAccountMembersListMembersResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A member's status in the organization.
type AccountMemberAccountMembersListMembersResponseStatus string

const (
	AccountMemberAccountMembersListMembersResponseStatusAccepted AccountMemberAccountMembersListMembersResponseStatus = "accepted"
	AccountMemberAccountMembersListMembersResponseStatusInvited  AccountMemberAccountMembersListMembersResponseStatus = "invited"
)

type AccountMemberGetResponseEnvelope struct {
	Errors   []AccountMemberGetResponseEnvelopeErrors   `json:"errors"`
	Messages []AccountMemberGetResponseEnvelopeMessages `json:"messages"`
	Result   AccountMemberGetResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccountMemberGetResponseEnvelopeSuccess `json:"success"`
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

type AccountMemberUpdateParams struct {
	// Roles assigned to this member.
	Roles param.Field[[]AccountMemberUpdateParamsRole] `json:"roles,required"`
}

func (r AccountMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRole struct {
	// Role identifier tag.
	ID          param.Field[string]                                    `json:"id,required"`
	Permissions param.Field[AccountMemberUpdateParamsRolesPermissions] `json:"permissions,required"`
}

func (r AccountMemberUpdateParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissions struct {
	Analytics    param.Field[AccountMemberUpdateParamsRolesPermissionsAnalytics]    `json:"analytics"`
	Billing      param.Field[AccountMemberUpdateParamsRolesPermissionsBilling]      `json:"billing"`
	CachePurge   param.Field[AccountMemberUpdateParamsRolesPermissionsCachePurge]   `json:"cache_purge"`
	DNS          param.Field[AccountMemberUpdateParamsRolesPermissionsDNS]          `json:"dns"`
	DNSRecords   param.Field[AccountMemberUpdateParamsRolesPermissionsDNSRecords]   `json:"dns_records"`
	Lb           param.Field[AccountMemberUpdateParamsRolesPermissionsLb]           `json:"lb"`
	Logs         param.Field[AccountMemberUpdateParamsRolesPermissionsLogs]         `json:"logs"`
	Organization param.Field[AccountMemberUpdateParamsRolesPermissionsOrganization] `json:"organization"`
	SSL          param.Field[AccountMemberUpdateParamsRolesPermissionsSSL]          `json:"ssl"`
	WAF          param.Field[AccountMemberUpdateParamsRolesPermissionsWAF]          `json:"waf"`
	ZoneSettings param.Field[AccountMemberUpdateParamsRolesPermissionsZoneSettings] `json:"zone_settings"`
	Zones        param.Field[AccountMemberUpdateParamsRolesPermissionsZones]        `json:"zones"`
}

func (r AccountMemberUpdateParamsRolesPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsAnalytics struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsBilling struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsCachePurge struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsCachePurge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsDNS struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsDNSRecords struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsDNSRecords) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsLb struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsLb) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsLogs struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsOrganization struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsSSL struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsWAF struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsZoneSettings struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsZoneSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsZones struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsZones) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateResponseEnvelope struct {
	Errors   []AccountMemberUpdateResponseEnvelopeErrors   `json:"errors"`
	Messages []AccountMemberUpdateResponseEnvelopeMessages `json:"messages"`
	Result   AccountMemberUpdateResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccountMemberUpdateResponseEnvelopeSuccess `json:"success"`
	JSON    accountMemberUpdateResponseEnvelopeJSON    `json:"-"`
}

// accountMemberUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountMemberUpdateResponseEnvelope]
type accountMemberUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountMemberUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountMemberUpdateResponseEnvelopeErrors]
type accountMemberUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountMemberUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccountMemberUpdateResponseEnvelopeMessages]
type accountMemberUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberUpdateResponseEnvelopeSuccess bool

const (
	AccountMemberUpdateResponseEnvelopeSuccessTrue AccountMemberUpdateResponseEnvelopeSuccess = true
)

type AccountMemberDeleteResponseEnvelope struct {
	Errors   []AccountMemberDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccountMemberDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccountMemberDeleteResponse                   `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMemberDeleteResponseEnvelopeSuccess `json:"success"`
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

type AccountMemberAccountMembersAddMemberParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                                         `json:"roles,required"`
	Status param.Field[AccountMemberAccountMembersAddMemberParamsStatus] `json:"status"`
}

func (r AccountMemberAccountMembersAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberAccountMembersAddMemberParamsStatus string

const (
	AccountMemberAccountMembersAddMemberParamsStatusAccepted AccountMemberAccountMembersAddMemberParamsStatus = "accepted"
	AccountMemberAccountMembersAddMemberParamsStatusPending  AccountMemberAccountMembersAddMemberParamsStatus = "pending"
)

type AccountMemberAccountMembersAddMemberResponseEnvelope struct {
	Errors   []AccountMemberAccountMembersAddMemberResponseEnvelopeErrors   `json:"errors"`
	Messages []AccountMemberAccountMembersAddMemberResponseEnvelopeMessages `json:"messages"`
	Result   AccountMemberAccountMembersAddMemberResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccountMemberAccountMembersAddMemberResponseEnvelopeSuccess `json:"success"`
	JSON    accountMemberAccountMembersAddMemberResponseEnvelopeJSON    `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseEnvelopeJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersAddMemberResponseEnvelope]
type accountMemberAccountMembersAddMemberResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountMemberAccountMembersAddMemberResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseEnvelopeErrors]
type accountMemberAccountMembersAddMemberResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountMemberAccountMembersAddMemberResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseEnvelopeMessages]
type accountMemberAccountMembersAddMemberResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberAccountMembersAddMemberResponseEnvelopeSuccess bool

const (
	AccountMemberAccountMembersAddMemberResponseEnvelopeSuccessTrue AccountMemberAccountMembersAddMemberResponseEnvelopeSuccess = true
)

type AccountMemberAccountMembersListMembersParams struct {
	// Direction to order results.
	Direction param.Field[AccountMemberAccountMembersListMembersParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountMemberAccountMembersListMembersParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[AccountMemberAccountMembersListMembersParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountMemberAccountMembersListMembersParams]'s query
// parameters as `url.Values`.
func (r AccountMemberAccountMembersListMembersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountMemberAccountMembersListMembersParamsDirection string

const (
	AccountMemberAccountMembersListMembersParamsDirectionAsc  AccountMemberAccountMembersListMembersParamsDirection = "asc"
	AccountMemberAccountMembersListMembersParamsDirectionDesc AccountMemberAccountMembersListMembersParamsDirection = "desc"
)

// Field to order results by.
type AccountMemberAccountMembersListMembersParamsOrder string

const (
	AccountMemberAccountMembersListMembersParamsOrderUserFirstName AccountMemberAccountMembersListMembersParamsOrder = "user.first_name"
	AccountMemberAccountMembersListMembersParamsOrderUserLastName  AccountMemberAccountMembersListMembersParamsOrder = "user.last_name"
	AccountMemberAccountMembersListMembersParamsOrderUserEmail     AccountMemberAccountMembersListMembersParamsOrder = "user.email"
	AccountMemberAccountMembersListMembersParamsOrderStatus        AccountMemberAccountMembersListMembersParamsOrder = "status"
)

// A member's status in the account.
type AccountMemberAccountMembersListMembersParamsStatus string

const (
	AccountMemberAccountMembersListMembersParamsStatusAccepted AccountMemberAccountMembersListMembersParamsStatus = "accepted"
	AccountMemberAccountMembersListMembersParamsStatusPending  AccountMemberAccountMembersListMembersParamsStatus = "pending"
	AccountMemberAccountMembersListMembersParamsStatusRejected AccountMemberAccountMembersListMembersParamsStatus = "rejected"
)

type AccountMemberAccountMembersListMembersResponseEnvelope struct {
	Errors     []AccountMemberAccountMembersListMembersResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccountMemberAccountMembersListMembersResponseEnvelopeMessages `json:"messages"`
	Result     []AccountMemberAccountMembersListMembersResponse                 `json:"result"`
	ResultInfo AccountMemberAccountMembersListMembersResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountMemberAccountMembersListMembersResponseEnvelopeSuccess `json:"success"`
	JSON    accountMemberAccountMembersListMembersResponseEnvelopeJSON    `json:"-"`
}

// accountMemberAccountMembersListMembersResponseEnvelopeJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersListMembersResponseEnvelope]
type accountMemberAccountMembersListMembersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersListMembersResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountMemberAccountMembersListMembersResponseEnvelopeErrorsJSON `json:"-"`
}

// accountMemberAccountMembersListMembersResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersListMembersResponseEnvelopeErrors]
type accountMemberAccountMembersListMembersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersListMembersResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountMemberAccountMembersListMembersResponseEnvelopeMessagesJSON `json:"-"`
}

// accountMemberAccountMembersListMembersResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccountMemberAccountMembersListMembersResponseEnvelopeMessages]
type accountMemberAccountMembersListMembersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersListMembersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       accountMemberAccountMembersListMembersResponseEnvelopeResultInfoJSON `json:"-"`
}

// accountMemberAccountMembersListMembersResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [AccountMemberAccountMembersListMembersResponseEnvelopeResultInfo]
type accountMemberAccountMembersListMembersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersListMembersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberAccountMembersListMembersResponseEnvelopeSuccess bool

const (
	AccountMemberAccountMembersListMembersResponseEnvelopeSuccessTrue AccountMemberAccountMembersListMembersResponseEnvelopeSuccess = true
)
