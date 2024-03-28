// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// MemberService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r *MemberService) {
	r = &MemberService{}
	r.Options = opts
	return
}

// Add a user to the list of members for this account.
func (r *MemberService) New(ctx context.Context, params MemberNewParams, opts ...option.RequestOption) (res *AccountMemberWithID, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *MemberService) Update(ctx context.Context, memberID string, params MemberUpdateParams, opts ...option.RequestOption) (res *AccountMember, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *MemberService) List(ctx context.Context, params MemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[MemberListResponse], err error) {
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
func (r *MemberService) ListAutoPaging(ctx context.Context, params MemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[MemberListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a member from an account.
func (r *MemberService) Delete(ctx context.Context, memberID string, body MemberDeleteParams, opts ...option.RequestOption) (res *MemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", body.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *MemberService) Get(ctx context.Context, memberID string, query MemberGetParams, opts ...option.RequestOption) (res *AccountMember, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

type MemberListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []Role `json:"roles,required"`
	// A member's status in the organization.
	Status MemberListResponseStatus `json:"status,required"`
	JSON   memberListResponseJSON   `json:"-"`
}

// memberListResponseJSON contains the JSON metadata for the struct
// [MemberListResponse]
type memberListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponseJSON) RawJSON() string {
	return r.raw
}

// A member's status in the organization.
type MemberListResponseStatus string

const (
	MemberListResponseStatusAccepted MemberListResponseStatus = "accepted"
	MemberListResponseStatusInvited  MemberListResponseStatus = "invited"
)

func (r MemberListResponseStatus) IsKnown() bool {
	switch r {
	case MemberListResponseStatusAccepted, MemberListResponseStatusInvited:
		return true
	}
	return false
}

type MemberDeleteResponse struct {
	// Identifier
	ID   string                   `json:"id,required"`
	JSON memberDeleteResponseJSON `json:"-"`
}

// memberDeleteResponseJSON contains the JSON metadata for the struct
// [MemberDeleteResponse]
type memberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MemberNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]              `json:"roles,required"`
	Status param.Field[MemberNewParamsStatus] `json:"status"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsStatus string

const (
	MemberNewParamsStatusAccepted MemberNewParamsStatus = "accepted"
	MemberNewParamsStatusPending  MemberNewParamsStatus = "pending"
)

func (r MemberNewParamsStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsStatusAccepted, MemberNewParamsStatusPending:
		return true
	}
	return false
}

type MemberNewResponseEnvelope struct {
	Errors   []MemberNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMemberWithID                 `json:"result,required"`
	// Whether the API call was successful
	Success MemberNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberNewResponseEnvelopeJSON    `json:"-"`
}

// memberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelope]
type memberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    memberNewResponseEnvelopeErrorsJSON `json:"-"`
}

// memberNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeErrors]
type memberNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    memberNewResponseEnvelopeMessagesJSON `json:"-"`
}

// memberNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeMessages]
type memberNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberNewResponseEnvelopeSuccess bool

const (
	MemberNewResponseEnvelopeSuccessTrue MemberNewResponseEnvelopeSuccess = true
)

func (r MemberNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Roles assigned to this member.
	Roles param.Field[[]MemberUpdateParamsRole] `json:"roles,required"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRole struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissions struct {
	Analytics    param.Field[MemberUpdateParamsRolesPermissionsAnalytics]    `json:"analytics"`
	Billing      param.Field[MemberUpdateParamsRolesPermissionsBilling]      `json:"billing"`
	CachePurge   param.Field[MemberUpdateParamsRolesPermissionsCachePurge]   `json:"cache_purge"`
	DNS          param.Field[MemberUpdateParamsRolesPermissionsDNS]          `json:"dns"`
	DNSRecords   param.Field[MemberUpdateParamsRolesPermissionsDNSRecords]   `json:"dns_records"`
	Lb           param.Field[MemberUpdateParamsRolesPermissionsLb]           `json:"lb"`
	Logs         param.Field[MemberUpdateParamsRolesPermissionsLogs]         `json:"logs"`
	Organization param.Field[MemberUpdateParamsRolesPermissionsOrganization] `json:"organization"`
	SSL          param.Field[MemberUpdateParamsRolesPermissionsSSL]          `json:"ssl"`
	WAF          param.Field[MemberUpdateParamsRolesPermissionsWAF]          `json:"waf"`
	ZoneSettings param.Field[MemberUpdateParamsRolesPermissionsZoneSettings] `json:"zone_settings"`
	Zones        param.Field[MemberUpdateParamsRolesPermissionsZones]        `json:"zones"`
}

func (r MemberUpdateParamsRolesPermissions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsAnalytics struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsAnalytics) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsBilling struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsBilling) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsCachePurge struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsCachePurge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsDNS struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsDNSRecords struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsDNSRecords) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsLb struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsLb) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsLogs struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsLogs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsOrganization struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsOrganization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsSSL struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsWAF struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsWAF) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsZoneSettings struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsZoneSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsRolesPermissionsZones struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r MemberUpdateParamsRolesPermissionsZones) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateResponseEnvelope struct {
	Errors   []MemberUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMember                          `json:"result,required"`
	// Whether the API call was successful
	Success MemberUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberUpdateResponseEnvelopeJSON    `json:"-"`
}

// memberUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelope]
type memberUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    memberUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// memberUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelopeErrors]
type memberUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    memberUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// memberUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberUpdateResponseEnvelopeMessages]
type memberUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberUpdateResponseEnvelopeSuccess bool

const (
	MemberUpdateResponseEnvelopeSuccessTrue MemberUpdateResponseEnvelopeSuccess = true
)

func (r MemberUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Direction to order results.
	Direction param.Field[MemberListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[MemberListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[MemberListParamsStatus] `query:"status"`
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type MemberListParamsDirection string

const (
	MemberListParamsDirectionAsc  MemberListParamsDirection = "asc"
	MemberListParamsDirectionDesc MemberListParamsDirection = "desc"
)

func (r MemberListParamsDirection) IsKnown() bool {
	switch r {
	case MemberListParamsDirectionAsc, MemberListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type MemberListParamsOrder string

const (
	MemberListParamsOrderUserFirstName MemberListParamsOrder = "user.first_name"
	MemberListParamsOrderUserLastName  MemberListParamsOrder = "user.last_name"
	MemberListParamsOrderUserEmail     MemberListParamsOrder = "user.email"
	MemberListParamsOrderStatus        MemberListParamsOrder = "status"
)

func (r MemberListParamsOrder) IsKnown() bool {
	switch r {
	case MemberListParamsOrderUserFirstName, MemberListParamsOrderUserLastName, MemberListParamsOrderUserEmail, MemberListParamsOrderStatus:
		return true
	}
	return false
}

// A member's status in the account.
type MemberListParamsStatus string

const (
	MemberListParamsStatusAccepted MemberListParamsStatus = "accepted"
	MemberListParamsStatusPending  MemberListParamsStatus = "pending"
	MemberListParamsStatusRejected MemberListParamsStatus = "rejected"
)

func (r MemberListParamsStatus) IsKnown() bool {
	switch r {
	case MemberListParamsStatusAccepted, MemberListParamsStatusPending, MemberListParamsStatusRejected:
		return true
	}
	return false
}

type MemberDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type MemberDeleteResponseEnvelope struct {
	Errors   []MemberDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MemberDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberDeleteResponseEnvelopeJSON    `json:"-"`
}

// memberDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelope]
type memberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    memberDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// memberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelopeErrors]
type memberDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    memberDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// memberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberDeleteResponseEnvelopeMessages]
type memberDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberDeleteResponseEnvelopeSuccess bool

const (
	MemberDeleteResponseEnvelopeSuccessTrue MemberDeleteResponseEnvelopeSuccess = true
)

func (r MemberDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type MemberGetResponseEnvelope struct {
	Errors   []MemberGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountMember                       `json:"result,required"`
	// Whether the API call was successful
	Success MemberGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    memberGetResponseEnvelopeJSON    `json:"-"`
}

// memberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelope]
type memberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    memberGetResponseEnvelopeErrorsJSON `json:"-"`
}

// memberGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeErrors]
type memberGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    memberGetResponseEnvelopeMessagesJSON `json:"-"`
}

// memberGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeMessages]
type memberGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MemberGetResponseEnvelopeSuccess bool

const (
	MemberGetResponseEnvelopeSuccessTrue MemberGetResponseEnvelopeSuccess = true
)

func (r MemberGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
