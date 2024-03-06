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
func (r *AccountMemberService) New(ctx context.Context, params AccountMemberNewParams, opts ...option.RequestOption) (res *AccountMemberWithID, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *AccountMemberService) Update(ctx context.Context, memberID string, params AccountMemberUpdateParams, opts ...option.RequestOption) (res *AccountMember, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *AccountMemberService) List(ctx context.Context, params AccountMemberListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[AccountMemberListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%v/members", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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
func (r *AccountMemberService) ListAutoPaging(ctx context.Context, params AccountMemberListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[AccountMemberListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a member from an account.
func (r *AccountMemberService) Delete(ctx context.Context, memberID string, body AccountMemberDeleteParams, opts ...option.RequestOption) (res *AccountMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", body.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, memberID string, query AccountMemberGetParams, opts ...option.RequestOption) (res *AccountMember, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountMemberGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountMember struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberRole `json:"roles,required"`
	Status interface{}         `json:"status,required"`
	User   AccountMemberUser   `json:"user,required"`
	JSON   accountMemberJSON   `json:"-"`
}

// accountMemberJSON contains the JSON metadata for the struct [AccountMember]
type accountMemberJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                        `json:"name,required"`
	Permissions AccountMemberRolesPermissions `json:"permissions,required"`
	JSON        accountMemberRoleJSON         `json:"-"`
}

// accountMemberRoleJSON contains the JSON metadata for the struct
// [AccountMemberRole]
type accountMemberRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRoleJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissions struct {
	Analytics    AccountMemberRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberRolesPermissionsJSON         `json:"-"`
}

// accountMemberRolesPermissionsJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissions]
type accountMemberRolesPermissionsJSON struct {
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

func (r *AccountMemberRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsAnalytics struct {
	Read  bool                                       `json:"read"`
	Write bool                                       `json:"write"`
	JSON  accountMemberRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberRolesPermissionsAnalyticsJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsAnalytics]
type accountMemberRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsAnalyticsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsBilling struct {
	Read  bool                                     `json:"read"`
	Write bool                                     `json:"write"`
	JSON  accountMemberRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberRolesPermissionsBillingJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsBilling]
type accountMemberRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsBillingJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsCachePurge struct {
	Read  bool                                        `json:"read"`
	Write bool                                        `json:"write"`
	JSON  accountMemberRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberRolesPermissionsCachePurgeJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsCachePurge]
type accountMemberRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsCachePurgeJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsDNS struct {
	Read  bool                                 `json:"read"`
	Write bool                                 `json:"write"`
	JSON  accountMemberRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberRolesPermissionsDNSJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsDNS]
type accountMemberRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsDNSJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsDNSRecords struct {
	Read  bool                                        `json:"read"`
	Write bool                                        `json:"write"`
	JSON  accountMemberRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberRolesPermissionsDNSRecordsJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsDNSRecords]
type accountMemberRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsDNSRecordsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsLb struct {
	Read  bool                                `json:"read"`
	Write bool                                `json:"write"`
	JSON  accountMemberRolesPermissionsLbJSON `json:"-"`
}

// accountMemberRolesPermissionsLbJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsLb]
type accountMemberRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsLbJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsLogs struct {
	Read  bool                                  `json:"read"`
	Write bool                                  `json:"write"`
	JSON  accountMemberRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberRolesPermissionsLogsJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsLogs]
type accountMemberRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsLogsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsOrganization struct {
	Read  bool                                          `json:"read"`
	Write bool                                          `json:"write"`
	JSON  accountMemberRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberRolesPermissionsOrganizationJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsOrganization]
type accountMemberRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsOrganizationJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsSSL struct {
	Read  bool                                 `json:"read"`
	Write bool                                 `json:"write"`
	JSON  accountMemberRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberRolesPermissionsSSLJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsSSL]
type accountMemberRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsSSLJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsWAF struct {
	Read  bool                                 `json:"read"`
	Write bool                                 `json:"write"`
	JSON  accountMemberRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberRolesPermissionsWAFJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsWAF]
type accountMemberRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsWAFJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsZoneSettings struct {
	Read  bool                                          `json:"read"`
	Write bool                                          `json:"write"`
	JSON  accountMemberRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberRolesPermissionsZoneSettingsJSON contains the JSON metadata for the
// struct [AccountMemberRolesPermissionsZoneSettings]
type accountMemberRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsZoneSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberRolesPermissionsZones struct {
	Read  bool                                   `json:"read"`
	Write bool                                   `json:"write"`
	JSON  accountMemberRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberRolesPermissionsZonesJSON contains the JSON metadata for the struct
// [AccountMemberRolesPermissionsZones]
type accountMemberRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberRolesPermissionsZonesJSON) RawJSON() string {
	return r.raw
}

type AccountMemberUser struct {
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
	TwoFactorAuthenticationEnabled bool                  `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberUserJSON `json:"-"`
}

// accountMemberUserJSON contains the JSON metadata for the struct
// [AccountMemberUser]
type accountMemberUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberUserJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithID struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberWithIDRole `json:"roles,required"`
	Status interface{}               `json:"status,required"`
	User   AccountMemberWithIDUser   `json:"user,required"`
	// The unique activation code for the account membership.
	Code string                  `json:"code"`
	JSON accountMemberWithIDJSON `json:"-"`
}

// accountMemberWithIDJSON contains the JSON metadata for the struct
// [AccountMemberWithID]
type accountMemberWithIDJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                              `json:"name,required"`
	Permissions AccountMemberWithIDRolesPermissions `json:"permissions,required"`
	JSON        accountMemberWithIDRoleJSON         `json:"-"`
}

// accountMemberWithIDRoleJSON contains the JSON metadata for the struct
// [AccountMemberWithIDRole]
type accountMemberWithIDRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRoleJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissions struct {
	Analytics    AccountMemberWithIDRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberWithIDRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberWithIDRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberWithIDRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberWithIDRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberWithIDRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberWithIDRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberWithIDRolesPermissionsOrganization `json:"organization"`
	SSL          AccountMemberWithIDRolesPermissionsSSL          `json:"ssl"`
	WAF          AccountMemberWithIDRolesPermissionsWAF          `json:"waf"`
	ZoneSettings AccountMemberWithIDRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberWithIDRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberWithIDRolesPermissionsJSON         `json:"-"`
}

// accountMemberWithIDRolesPermissionsJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissions]
type accountMemberWithIDRolesPermissionsJSON struct {
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

func (r *AccountMemberWithIDRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsAnalytics struct {
	Read  bool                                             `json:"read"`
	Write bool                                             `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsAnalyticsJSON contains the JSON metadata for
// the struct [AccountMemberWithIDRolesPermissionsAnalytics]
type accountMemberWithIDRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsAnalyticsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsBilling struct {
	Read  bool                                           `json:"read"`
	Write bool                                           `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsBillingJSON contains the JSON metadata for
// the struct [AccountMemberWithIDRolesPermissionsBilling]
type accountMemberWithIDRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsBillingJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsCachePurge struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsCachePurgeJSON contains the JSON metadata for
// the struct [AccountMemberWithIDRolesPermissionsCachePurge]
type accountMemberWithIDRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsCachePurgeJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsDNS struct {
	Read  bool                                       `json:"read"`
	Write bool                                       `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsDNSJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsDNS]
type accountMemberWithIDRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsDNSJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsDNSRecords struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsDNSRecordsJSON contains the JSON metadata for
// the struct [AccountMemberWithIDRolesPermissionsDNSRecords]
type accountMemberWithIDRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsDNSRecordsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsLb struct {
	Read  bool                                      `json:"read"`
	Write bool                                      `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsLbJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsLbJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsLb]
type accountMemberWithIDRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsLbJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsLogs struct {
	Read  bool                                        `json:"read"`
	Write bool                                        `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsLogsJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsLogs]
type accountMemberWithIDRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsLogsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsOrganization struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsOrganizationJSON contains the JSON metadata
// for the struct [AccountMemberWithIDRolesPermissionsOrganization]
type accountMemberWithIDRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsOrganizationJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsSSL struct {
	Read  bool                                       `json:"read"`
	Write bool                                       `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsSSLJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsSSLJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsSSL]
type accountMemberWithIDRolesPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsSSLJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsWAF struct {
	Read  bool                                       `json:"read"`
	Write bool                                       `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsWAFJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsWAFJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsWAF]
type accountMemberWithIDRolesPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsWAFJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsZoneSettings struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsZoneSettingsJSON contains the JSON metadata
// for the struct [AccountMemberWithIDRolesPermissionsZoneSettings]
type accountMemberWithIDRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsZoneSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDRolesPermissionsZones struct {
	Read  bool                                         `json:"read"`
	Write bool                                         `json:"write"`
	JSON  accountMemberWithIDRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberWithIDRolesPermissionsZonesJSON contains the JSON metadata for the
// struct [AccountMemberWithIDRolesPermissionsZones]
type accountMemberWithIDRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberWithIDRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDRolesPermissionsZonesJSON) RawJSON() string {
	return r.raw
}

type AccountMemberWithIDUser struct {
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
	TwoFactorAuthenticationEnabled bool                        `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberWithIDUserJSON `json:"-"`
}

// accountMemberWithIDUserJSON contains the JSON metadata for the struct
// [AccountMemberWithIDUser]
type accountMemberWithIDUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberWithIDUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberWithIDUserJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberListResponseJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberListResponseRoleJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMemberNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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
	Result   AccountMemberWithID                        `json:"result,required"`
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

func (r accountMemberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMemberNewResponseEnvelopeSuccess bool

const (
	AccountMemberNewResponseEnvelopeSuccessTrue AccountMemberNewResponseEnvelopeSuccess = true
)

type AccountMemberUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Roles assigned to this member.
	Roles param.Field[[]AccountMemberUpdateParamsRole] `json:"roles,required"`
}

func (r AccountMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRole struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
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
	Errors   []AccountMemberUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMember                                 `json:"result,required"`
	// Whether the API call was successful
	Success AccountMemberUpdateResponseEnvelopeSuccess `json:"success,required"`
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

func (r accountMemberUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMemberUpdateResponseEnvelopeSuccess bool

const (
	AccountMemberUpdateResponseEnvelopeSuccessTrue AccountMemberUpdateResponseEnvelopeSuccess = true
)

type AccountMemberListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
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

type AccountMemberDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

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

func (r accountMemberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMemberDeleteResponseEnvelopeSuccess bool

const (
	AccountMemberDeleteResponseEnvelopeSuccessTrue AccountMemberDeleteResponseEnvelopeSuccess = true
)

type AccountMemberGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type AccountMemberGetResponseEnvelope struct {
	Errors   []AccountMemberGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountMemberGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMember                              `json:"result,required"`
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

func (r accountMemberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accountMemberGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMemberGetResponseEnvelopeSuccess bool

const (
	AccountMemberGetResponseEnvelopeSuccessTrue AccountMemberGetResponseEnvelopeSuccess = true
)
