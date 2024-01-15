// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationMemberService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationMemberService] method
// instead.
type OrganizationMemberService struct {
	Options []option.RequestOption
}

// NewOrganizationMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationMemberService(opts ...option.RequestOption) (r *OrganizationMemberService) {
	r = &OrganizationMemberService{}
	r.Options = opts
	return
}

// Get information about a specific member of an organization.
func (r *OrganizationMemberService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationMemberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change the Roles of an Organization's Member.
func (r *OrganizationMemberService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationMemberUpdateParams, opts ...option.RequestOption) (res *OrganizationMemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Remove a member from an organization.
func (r *OrganizationMemberService) Delete(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all members of a organization.
func (r *OrganizationMemberService) OrganizationMembersListMembers(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *OrganizationMemberOrganizationMembersListMembersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationMemberGetResponse struct {
	Errors   []OrganizationMemberGetResponseError   `json:"errors"`
	Messages []OrganizationMemberGetResponseMessage `json:"messages"`
	Result   OrganizationMemberGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success OrganizationMemberGetResponseSuccess `json:"success"`
	JSON    organizationMemberGetResponseJSON    `json:"-"`
}

// organizationMemberGetResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberGetResponse]
type organizationMemberGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    organizationMemberGetResponseErrorJSON `json:"-"`
}

// organizationMemberGetResponseErrorJSON contains the JSON metadata for the struct
// [OrganizationMemberGetResponseError]
type organizationMemberGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    organizationMemberGetResponseMessageJSON `json:"-"`
}

// organizationMemberGetResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationMemberGetResponseMessage]
type organizationMemberGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResult struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []OrganizationMemberGetResponseResultRole `json:"roles,required"`
	Status interface{}                               `json:"status,required"`
	User   OrganizationMemberGetResponseResultUser   `json:"user,required"`
	JSON   organizationMemberGetResponseResultJSON   `json:"-"`
}

// organizationMemberGetResponseResultJSON contains the JSON metadata for the
// struct [OrganizationMemberGetResponseResult]
type organizationMemberGetResponseResultJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                              `json:"name,required"`
	Permissions OrganizationMemberGetResponseResultRolesPermissions `json:"permissions,required"`
	JSON        organizationMemberGetResponseResultRoleJSON         `json:"-"`
}

// organizationMemberGetResponseResultRoleJSON contains the JSON metadata for the
// struct [OrganizationMemberGetResponseResultRole]
type organizationMemberGetResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissions struct {
	Analytics    OrganizationMemberGetResponseResultRolesPermissionsAnalytics    `json:"analytics"`
	Billing      OrganizationMemberGetResponseResultRolesPermissionsBilling      `json:"billing"`
	CachePurge   OrganizationMemberGetResponseResultRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          OrganizationMemberGetResponseResultRolesPermissionsDNS          `json:"dns"`
	DNSRecords   OrganizationMemberGetResponseResultRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           OrganizationMemberGetResponseResultRolesPermissionsLb           `json:"lb"`
	Logs         OrganizationMemberGetResponseResultRolesPermissionsLogs         `json:"logs"`
	Organization OrganizationMemberGetResponseResultRolesPermissionsOrganization `json:"organization"`
	Ssl          OrganizationMemberGetResponseResultRolesPermissionsSsl          `json:"ssl"`
	Waf          OrganizationMemberGetResponseResultRolesPermissionsWaf          `json:"waf"`
	ZoneSettings OrganizationMemberGetResponseResultRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        OrganizationMemberGetResponseResultRolesPermissionsZones        `json:"zones"`
	JSON         organizationMemberGetResponseResultRolesPermissionsJSON         `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsJSON contains the JSON
// metadata for the struct [OrganizationMemberGetResponseResultRolesPermissions]
type organizationMemberGetResponseResultRolesPermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	Ssl          apijson.Field
	Waf          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsAnalytics struct {
	Read  bool                                                             `json:"read"`
	Write bool                                                             `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsAnalyticsJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsAnalyticsJSON contains the
// JSON metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsAnalytics]
type organizationMemberGetResponseResultRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsBilling struct {
	Read  bool                                                           `json:"read"`
	Write bool                                                           `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsBillingJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsBillingJSON contains the JSON
// metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsBilling]
type organizationMemberGetResponseResultRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsCachePurge struct {
	Read  bool                                                              `json:"read"`
	Write bool                                                              `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsCachePurgeJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsCachePurgeJSON contains the
// JSON metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsCachePurge]
type organizationMemberGetResponseResultRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsDNS struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsDNSJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsDNSJSON contains the JSON
// metadata for the struct [OrganizationMemberGetResponseResultRolesPermissionsDNS]
type organizationMemberGetResponseResultRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsDNSRecords struct {
	Read  bool                                                              `json:"read"`
	Write bool                                                              `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsDNSRecordsJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsDNSRecordsJSON contains the
// JSON metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsDNSRecords]
type organizationMemberGetResponseResultRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsLb struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsLbJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsLbJSON contains the JSON
// metadata for the struct [OrganizationMemberGetResponseResultRolesPermissionsLb]
type organizationMemberGetResponseResultRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsLogs struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsLogsJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsLogsJSON contains the JSON
// metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsLogs]
type organizationMemberGetResponseResultRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsOrganization struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsOrganizationJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsOrganizationJSON contains the
// JSON metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsOrganization]
type organizationMemberGetResponseResultRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsSsl struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsSslJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsSslJSON contains the JSON
// metadata for the struct [OrganizationMemberGetResponseResultRolesPermissionsSsl]
type organizationMemberGetResponseResultRolesPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsWaf struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsWafJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsWafJSON contains the JSON
// metadata for the struct [OrganizationMemberGetResponseResultRolesPermissionsWaf]
type organizationMemberGetResponseResultRolesPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsZoneSettings struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsZoneSettingsJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsZoneSettingsJSON contains the
// JSON metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsZoneSettings]
type organizationMemberGetResponseResultRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultRolesPermissionsZones struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  organizationMemberGetResponseResultRolesPermissionsZonesJSON `json:"-"`
}

// organizationMemberGetResponseResultRolesPermissionsZonesJSON contains the JSON
// metadata for the struct
// [OrganizationMemberGetResponseResultRolesPermissionsZones]
type organizationMemberGetResponseResultRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberGetResponseResultUser struct {
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
	TwoFactorAuthenticationEnabled bool                                        `json:"two_factor_authentication_enabled"`
	JSON                           organizationMemberGetResponseResultUserJSON `json:"-"`
}

// organizationMemberGetResponseResultUserJSON contains the JSON metadata for the
// struct [OrganizationMemberGetResponseResultUser]
type organizationMemberGetResponseResultUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *OrganizationMemberGetResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationMemberGetResponseSuccess bool

const (
	OrganizationMemberGetResponseSuccessTrue OrganizationMemberGetResponseSuccess = true
)

type OrganizationMemberUpdateResponse struct {
	Errors   []OrganizationMemberUpdateResponseError   `json:"errors"`
	Messages []OrganizationMemberUpdateResponseMessage `json:"messages"`
	Result   OrganizationMemberUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success OrganizationMemberUpdateResponseSuccess `json:"success"`
	JSON    organizationMemberUpdateResponseJSON    `json:"-"`
}

// organizationMemberUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberUpdateResponse]
type organizationMemberUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    organizationMemberUpdateResponseErrorJSON `json:"-"`
}

// organizationMemberUpdateResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationMemberUpdateResponseError]
type organizationMemberUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    organizationMemberUpdateResponseMessageJSON `json:"-"`
}

// organizationMemberUpdateResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationMemberUpdateResponseMessage]
type organizationMemberUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResult struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []OrganizationMemberUpdateResponseResultRole `json:"roles,required"`
	Status interface{}                                  `json:"status,required"`
	User   OrganizationMemberUpdateResponseResultUser   `json:"user,required"`
	JSON   organizationMemberUpdateResponseResultJSON   `json:"-"`
}

// organizationMemberUpdateResponseResultJSON contains the JSON metadata for the
// struct [OrganizationMemberUpdateResponseResult]
type organizationMemberUpdateResponseResultJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                                 `json:"name,required"`
	Permissions OrganizationMemberUpdateResponseResultRolesPermissions `json:"permissions,required"`
	JSON        organizationMemberUpdateResponseResultRoleJSON         `json:"-"`
}

// organizationMemberUpdateResponseResultRoleJSON contains the JSON metadata for
// the struct [OrganizationMemberUpdateResponseResultRole]
type organizationMemberUpdateResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissions struct {
	Analytics    OrganizationMemberUpdateResponseResultRolesPermissionsAnalytics    `json:"analytics"`
	Billing      OrganizationMemberUpdateResponseResultRolesPermissionsBilling      `json:"billing"`
	CachePurge   OrganizationMemberUpdateResponseResultRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          OrganizationMemberUpdateResponseResultRolesPermissionsDNS          `json:"dns"`
	DNSRecords   OrganizationMemberUpdateResponseResultRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           OrganizationMemberUpdateResponseResultRolesPermissionsLb           `json:"lb"`
	Logs         OrganizationMemberUpdateResponseResultRolesPermissionsLogs         `json:"logs"`
	Organization OrganizationMemberUpdateResponseResultRolesPermissionsOrganization `json:"organization"`
	Ssl          OrganizationMemberUpdateResponseResultRolesPermissionsSsl          `json:"ssl"`
	Waf          OrganizationMemberUpdateResponseResultRolesPermissionsWaf          `json:"waf"`
	ZoneSettings OrganizationMemberUpdateResponseResultRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        OrganizationMemberUpdateResponseResultRolesPermissionsZones        `json:"zones"`
	JSON         organizationMemberUpdateResponseResultRolesPermissionsJSON         `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsJSON contains the JSON
// metadata for the struct [OrganizationMemberUpdateResponseResultRolesPermissions]
type organizationMemberUpdateResponseResultRolesPermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	Lb           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	Ssl          apijson.Field
	Waf          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsAnalytics struct {
	Read  bool                                                                `json:"read"`
	Write bool                                                                `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsAnalyticsJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsAnalyticsJSON contains the
// JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsAnalytics]
type organizationMemberUpdateResponseResultRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsBilling struct {
	Read  bool                                                              `json:"read"`
	Write bool                                                              `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsBillingJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsBillingJSON contains the
// JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsBilling]
type organizationMemberUpdateResponseResultRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsCachePurge struct {
	Read  bool                                                                 `json:"read"`
	Write bool                                                                 `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsCachePurgeJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsCachePurgeJSON contains
// the JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsCachePurge]
type organizationMemberUpdateResponseResultRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsDNS struct {
	Read  bool                                                          `json:"read"`
	Write bool                                                          `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsDNSJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsDNSJSON contains the JSON
// metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsDNS]
type organizationMemberUpdateResponseResultRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsDNSRecords struct {
	Read  bool                                                                 `json:"read"`
	Write bool                                                                 `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON contains
// the JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsDNSRecords]
type organizationMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsLb struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsLbJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsLbJSON contains the JSON
// metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsLb]
type organizationMemberUpdateResponseResultRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsLogs struct {
	Read  bool                                                           `json:"read"`
	Write bool                                                           `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsLogsJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsLogsJSON contains the JSON
// metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsLogs]
type organizationMemberUpdateResponseResultRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsOrganization struct {
	Read  bool                                                                   `json:"read"`
	Write bool                                                                   `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsOrganizationJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsOrganizationJSON contains
// the JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsOrganization]
type organizationMemberUpdateResponseResultRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsSsl struct {
	Read  bool                                                          `json:"read"`
	Write bool                                                          `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsSslJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsSslJSON contains the JSON
// metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsSsl]
type organizationMemberUpdateResponseResultRolesPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsWaf struct {
	Read  bool                                                          `json:"read"`
	Write bool                                                          `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsWafJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsWafJSON contains the JSON
// metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsWaf]
type organizationMemberUpdateResponseResultRolesPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsZoneSettings struct {
	Read  bool                                                                   `json:"read"`
	Write bool                                                                   `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON contains
// the JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsZoneSettings]
type organizationMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultRolesPermissionsZones struct {
	Read  bool                                                            `json:"read"`
	Write bool                                                            `json:"write"`
	JSON  organizationMemberUpdateResponseResultRolesPermissionsZonesJSON `json:"-"`
}

// organizationMemberUpdateResponseResultRolesPermissionsZonesJSON contains the
// JSON metadata for the struct
// [OrganizationMemberUpdateResponseResultRolesPermissionsZones]
type organizationMemberUpdateResponseResultRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateResponseResultUser struct {
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
	TwoFactorAuthenticationEnabled bool                                           `json:"two_factor_authentication_enabled"`
	JSON                           organizationMemberUpdateResponseResultUserJSON `json:"-"`
}

// organizationMemberUpdateResponseResultUserJSON contains the JSON metadata for
// the struct [OrganizationMemberUpdateResponseResultUser]
type organizationMemberUpdateResponseResultUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *OrganizationMemberUpdateResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationMemberUpdateResponseSuccess bool

const (
	OrganizationMemberUpdateResponseSuccessTrue OrganizationMemberUpdateResponseSuccess = true
)

type OrganizationMemberDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON organizationMemberDeleteResponseJSON `json:"-"`
}

// organizationMemberDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberDeleteResponse]
type organizationMemberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberOrganizationMembersListMembersResponse struct {
	Errors     []OrganizationMemberOrganizationMembersListMembersResponseError    `json:"errors"`
	Messages   []OrganizationMemberOrganizationMembersListMembersResponseMessage  `json:"messages"`
	Result     []OrganizationMemberOrganizationMembersListMembersResponseResult   `json:"result"`
	ResultInfo OrganizationMemberOrganizationMembersListMembersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OrganizationMemberOrganizationMembersListMembersResponseSuccess `json:"success"`
	JSON    organizationMemberOrganizationMembersListMembersResponseJSON    `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseJSON contains the JSON
// metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponse]
type organizationMemberOrganizationMembersListMembersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberOrganizationMembersListMembersResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    organizationMemberOrganizationMembersListMembersResponseErrorJSON `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseErrorJSON contains the
// JSON metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponseError]
type organizationMemberOrganizationMembersListMembersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberOrganizationMembersListMembersResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    organizationMemberOrganizationMembersListMembersResponseMessageJSON `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseMessageJSON contains the
// JSON metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponseMessage]
type organizationMemberOrganizationMembersListMembersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberOrganizationMembersListMembersResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []OrganizationMemberOrganizationMembersListMembersResponseResultRole `json:"roles,required"`
	// A member's status in the organization.
	Status OrganizationMemberOrganizationMembersListMembersResponseResultStatus `json:"status,required"`
	JSON   organizationMemberOrganizationMembersListMembersResponseResultJSON   `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseResultJSON contains the
// JSON metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponseResult]
type organizationMemberOrganizationMembersListMembersResponseResultJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberOrganizationMembersListMembersResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                                               `json:"permissions,required"`
	JSON        organizationMemberOrganizationMembersListMembersResponseResultRoleJSON `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseResultRoleJSON contains
// the JSON metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponseResultRole]
type organizationMemberOrganizationMembersListMembersResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A member's status in the organization.
type OrganizationMemberOrganizationMembersListMembersResponseResultStatus string

const (
	OrganizationMemberOrganizationMembersListMembersResponseResultStatusAccepted OrganizationMemberOrganizationMembersListMembersResponseResultStatus = "accepted"
	OrganizationMemberOrganizationMembersListMembersResponseResultStatusInvited  OrganizationMemberOrganizationMembersListMembersResponseResultStatus = "invited"
)

type OrganizationMemberOrganizationMembersListMembersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       organizationMemberOrganizationMembersListMembersResponseResultInfoJSON `json:"-"`
}

// organizationMemberOrganizationMembersListMembersResponseResultInfoJSON contains
// the JSON metadata for the struct
// [OrganizationMemberOrganizationMembersListMembersResponseResultInfo]
type organizationMemberOrganizationMembersListMembersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberOrganizationMembersListMembersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationMemberOrganizationMembersListMembersResponseSuccess bool

const (
	OrganizationMemberOrganizationMembersListMembersResponseSuccessTrue OrganizationMemberOrganizationMembersListMembersResponseSuccess = true
)

type OrganizationMemberUpdateParams struct {
	// Array of Roles associated with this Member.
	Roles param.Field[[]string] `json:"roles"`
}

func (r OrganizationMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
