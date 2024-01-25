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

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, accountIdentifier interface{}, identifier string, opts ...option.RequestOption) (res *AccountMemberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify an account member.
func (r *AccountMemberService) Update(ctx context.Context, accountIdentifier interface{}, identifier string, body AccountMemberUpdateParams, opts ...option.RequestOption) (res *AccountMemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove a member from an account.
func (r *AccountMemberService) Delete(ctx context.Context, accountIdentifier interface{}, identifier string, opts ...option.RequestOption) (res *AccountMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a user to the list of members for this account.
func (r *AccountMemberService) AccountMembersAddMember(ctx context.Context, accountIdentifier interface{}, body AccountMemberAccountMembersAddMemberParams, opts ...option.RequestOption) (res *AccountMemberAccountMembersAddMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all members of an account.
func (r *AccountMemberService) AccountMembersListMembers(ctx context.Context, accountIdentifier interface{}, query AccountMemberAccountMembersListMembersParams, opts ...option.RequestOption) (res *shared.Page[AccountMemberAccountMembersListMembersResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%v/members", accountIdentifier)
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

type AccountMemberGetResponse struct {
	Errors   []AccountMemberGetResponseError   `json:"errors"`
	Messages []AccountMemberGetResponseMessage `json:"messages"`
	Result   AccountMemberGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMemberGetResponseSuccess `json:"success"`
	JSON    accountMemberGetResponseJSON    `json:"-"`
}

// accountMemberGetResponseJSON contains the JSON metadata for the struct
// [AccountMemberGetResponse]
type accountMemberGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accountMemberGetResponseErrorJSON `json:"-"`
}

// accountMemberGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseError]
type accountMemberGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountMemberGetResponseMessageJSON `json:"-"`
}

// accountMemberGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseMessage]
type accountMemberGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResult struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberGetResponseResultRole `json:"roles,required"`
	Status interface{}                          `json:"status,required"`
	User   AccountMemberGetResponseResultUser   `json:"user,required"`
	JSON   accountMemberGetResponseResultJSON   `json:"-"`
}

// accountMemberGetResponseResultJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseResult]
type accountMemberGetResponseResultJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                         `json:"name,required"`
	Permissions AccountMemberGetResponseResultRolesPermissions `json:"permissions,required"`
	JSON        accountMemberGetResponseResultRoleJSON         `json:"-"`
}

// accountMemberGetResponseResultRoleJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseResultRole]
type accountMemberGetResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissions struct {
	Analytics    AccountMemberGetResponseResultRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberGetResponseResultRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberGetResponseResultRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberGetResponseResultRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberGetResponseResultRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberGetResponseResultRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberGetResponseResultRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberGetResponseResultRolesPermissionsOrganization `json:"organization"`
	Ssl          AccountMemberGetResponseResultRolesPermissionsSsl          `json:"ssl"`
	Waf          AccountMemberGetResponseResultRolesPermissionsWaf          `json:"waf"`
	ZoneSettings AccountMemberGetResponseResultRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberGetResponseResultRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberGetResponseResultRolesPermissionsJSON         `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseResultRolesPermissions]
type accountMemberGetResponseResultRolesPermissionsJSON struct {
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

func (r *AccountMemberGetResponseResultRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsAnalytics struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsAnalyticsJSON contains the JSON
// metadata for the struct
// [AccountMemberGetResponseResultRolesPermissionsAnalytics]
type accountMemberGetResponseResultRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsBilling struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsBillingJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseResultRolesPermissionsBilling]
type accountMemberGetResponseResultRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsCachePurge struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsCachePurgeJSON contains the JSON
// metadata for the struct
// [AccountMemberGetResponseResultRolesPermissionsCachePurge]
type accountMemberGetResponseResultRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsDNS struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsDNSJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseResultRolesPermissionsDNS]
type accountMemberGetResponseResultRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsDNSRecords struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct
// [AccountMemberGetResponseResultRolesPermissionsDNSRecords]
type accountMemberGetResponseResultRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsLb struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsLbJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsLbJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseResultRolesPermissionsLb]
type accountMemberGetResponseResultRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsLogs struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsLogsJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseResultRolesPermissionsLogs]
type accountMemberGetResponseResultRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsOrganization struct {
	Read  bool                                                           `json:"read"`
	Write bool                                                           `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsOrganizationJSON contains the JSON
// metadata for the struct
// [AccountMemberGetResponseResultRolesPermissionsOrganization]
type accountMemberGetResponseResultRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsSsl struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsSslJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsSslJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseResultRolesPermissionsSsl]
type accountMemberGetResponseResultRolesPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsWaf struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsWafJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsWafJSON contains the JSON metadata
// for the struct [AccountMemberGetResponseResultRolesPermissionsWaf]
type accountMemberGetResponseResultRolesPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsZoneSettings struct {
	Read  bool                                                           `json:"read"`
	Write bool                                                           `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct
// [AccountMemberGetResponseResultRolesPermissionsZoneSettings]
type accountMemberGetResponseResultRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultRolesPermissionsZones struct {
	Read  bool                                                    `json:"read"`
	Write bool                                                    `json:"write"`
	JSON  accountMemberGetResponseResultRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberGetResponseResultRolesPermissionsZonesJSON contains the JSON
// metadata for the struct [AccountMemberGetResponseResultRolesPermissionsZones]
type accountMemberGetResponseResultRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberGetResponseResultUser struct {
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
	TwoFactorAuthenticationEnabled bool                                   `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberGetResponseResultUserJSON `json:"-"`
}

// accountMemberGetResponseResultUserJSON contains the JSON metadata for the struct
// [AccountMemberGetResponseResultUser]
type accountMemberGetResponseResultUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberGetResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberGetResponseSuccess bool

const (
	AccountMemberGetResponseSuccessTrue AccountMemberGetResponseSuccess = true
)

type AccountMemberUpdateResponse struct {
	Errors   []AccountMemberUpdateResponseError   `json:"errors"`
	Messages []AccountMemberUpdateResponseMessage `json:"messages"`
	Result   AccountMemberUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMemberUpdateResponseSuccess `json:"success"`
	JSON    accountMemberUpdateResponseJSON    `json:"-"`
}

// accountMemberUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponse]
type accountMemberUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountMemberUpdateResponseErrorJSON `json:"-"`
}

// accountMemberUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponseError]
type accountMemberUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountMemberUpdateResponseMessageJSON `json:"-"`
}

// accountMemberUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponseMessage]
type accountMemberUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResult struct {
	// Membership identifier tag.
	ID string `json:"id,required"`
	// Roles assigned to this member.
	Roles  []AccountMemberUpdateResponseResultRole `json:"roles,required"`
	Status interface{}                             `json:"status,required"`
	User   AccountMemberUpdateResponseResultUser   `json:"user,required"`
	JSON   accountMemberUpdateResponseResultJSON   `json:"-"`
}

// accountMemberUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccountMemberUpdateResponseResult]
type accountMemberUpdateResponseResultJSON struct {
	ID          apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                            `json:"name,required"`
	Permissions AccountMemberUpdateResponseResultRolesPermissions `json:"permissions,required"`
	JSON        accountMemberUpdateResponseResultRoleJSON         `json:"-"`
}

// accountMemberUpdateResponseResultRoleJSON contains the JSON metadata for the
// struct [AccountMemberUpdateResponseResultRole]
type accountMemberUpdateResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissions struct {
	Analytics    AccountMemberUpdateResponseResultRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberUpdateResponseResultRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberUpdateResponseResultRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberUpdateResponseResultRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberUpdateResponseResultRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberUpdateResponseResultRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberUpdateResponseResultRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberUpdateResponseResultRolesPermissionsOrganization `json:"organization"`
	Ssl          AccountMemberUpdateResponseResultRolesPermissionsSsl          `json:"ssl"`
	Waf          AccountMemberUpdateResponseResultRolesPermissionsWaf          `json:"waf"`
	ZoneSettings AccountMemberUpdateResponseResultRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberUpdateResponseResultRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberUpdateResponseResultRolesPermissionsJSON         `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsJSON contains the JSON metadata
// for the struct [AccountMemberUpdateResponseResultRolesPermissions]
type accountMemberUpdateResponseResultRolesPermissionsJSON struct {
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

func (r *AccountMemberUpdateResponseResultRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsAnalytics struct {
	Read  bool                                                           `json:"read"`
	Write bool                                                           `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsAnalyticsJSON contains the JSON
// metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsAnalytics]
type accountMemberUpdateResponseResultRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsBilling struct {
	Read  bool                                                         `json:"read"`
	Write bool                                                         `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsBillingJSON contains the JSON
// metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsBilling]
type accountMemberUpdateResponseResultRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsCachePurge struct {
	Read  bool                                                            `json:"read"`
	Write bool                                                            `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsCachePurgeJSON contains the
// JSON metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsCachePurge]
type accountMemberUpdateResponseResultRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsDNS struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsDNSJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsDNS]
type accountMemberUpdateResponseResultRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsDNSRecords struct {
	Read  bool                                                            `json:"read"`
	Write bool                                                            `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON contains the
// JSON metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsDNSRecords]
type accountMemberUpdateResponseResultRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsLb struct {
	Read  bool                                                    `json:"read"`
	Write bool                                                    `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsLbJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsLbJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsLb]
type accountMemberUpdateResponseResultRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsLogs struct {
	Read  bool                                                      `json:"read"`
	Write bool                                                      `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsLogsJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsLogs]
type accountMemberUpdateResponseResultRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsOrganization struct {
	Read  bool                                                              `json:"read"`
	Write bool                                                              `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsOrganizationJSON contains the
// JSON metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsOrganization]
type accountMemberUpdateResponseResultRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsSsl struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsSslJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsSslJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsSsl]
type accountMemberUpdateResponseResultRolesPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsWaf struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsWafJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsWafJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsWaf]
type accountMemberUpdateResponseResultRolesPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsZoneSettings struct {
	Read  bool                                                              `json:"read"`
	Write bool                                                              `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON contains the
// JSON metadata for the struct
// [AccountMemberUpdateResponseResultRolesPermissionsZoneSettings]
type accountMemberUpdateResponseResultRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultRolesPermissionsZones struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  accountMemberUpdateResponseResultRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberUpdateResponseResultRolesPermissionsZonesJSON contains the JSON
// metadata for the struct [AccountMemberUpdateResponseResultRolesPermissionsZones]
type accountMemberUpdateResponseResultRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberUpdateResponseResultUser struct {
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
	TwoFactorAuthenticationEnabled bool                                      `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberUpdateResponseResultUserJSON `json:"-"`
}

// accountMemberUpdateResponseResultUserJSON contains the JSON metadata for the
// struct [AccountMemberUpdateResponseResultUser]
type accountMemberUpdateResponseResultUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberUpdateResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberUpdateResponseSuccess bool

const (
	AccountMemberUpdateResponseSuccessTrue AccountMemberUpdateResponseSuccess = true
)

type AccountMemberDeleteResponse struct {
	Errors   []AccountMemberDeleteResponseError   `json:"errors"`
	Messages []AccountMemberDeleteResponseMessage `json:"messages"`
	Result   AccountMemberDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMemberDeleteResponseSuccess `json:"success"`
	JSON    accountMemberDeleteResponseJSON    `json:"-"`
}

// accountMemberDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMemberDeleteResponse]
type accountMemberDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberDeleteResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountMemberDeleteResponseErrorJSON `json:"-"`
}

// accountMemberDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountMemberDeleteResponseError]
type accountMemberDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberDeleteResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountMemberDeleteResponseMessageJSON `json:"-"`
}

// accountMemberDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AccountMemberDeleteResponseMessage]
type accountMemberDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberDeleteResponseResult struct {
	// Identifier
	ID   string                                `json:"id,required"`
	JSON accountMemberDeleteResponseResultJSON `json:"-"`
}

// accountMemberDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccountMemberDeleteResponseResult]
type accountMemberDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberDeleteResponseSuccess bool

const (
	AccountMemberDeleteResponseSuccessTrue AccountMemberDeleteResponseSuccess = true
)

type AccountMemberAccountMembersAddMemberResponse struct {
	Errors   []AccountMemberAccountMembersAddMemberResponseError   `json:"errors"`
	Messages []AccountMemberAccountMembersAddMemberResponseMessage `json:"messages"`
	Result   AccountMemberAccountMembersAddMemberResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountMemberAccountMembersAddMemberResponseSuccess `json:"success"`
	JSON    accountMemberAccountMembersAddMemberResponseJSON    `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseJSON contains the JSON metadata for
// the struct [AccountMemberAccountMembersAddMemberResponse]
type accountMemberAccountMembersAddMemberResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseError struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accountMemberAccountMembersAddMemberResponseErrorJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseErrorJSON contains the JSON metadata
// for the struct [AccountMemberAccountMembersAddMemberResponseError]
type accountMemberAccountMembersAddMemberResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseMessage struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    accountMemberAccountMembersAddMemberResponseMessageJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseMessageJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersAddMemberResponseMessage]
type accountMemberAccountMembersAddMemberResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResult struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// The unique activation code for the account membership.
	Code string `json:"code"`
	// Roles assigned to this member.
	Roles  []AccountMemberAccountMembersAddMemberResponseResultRole `json:"roles"`
	Status interface{}                                              `json:"status"`
	User   AccountMemberAccountMembersAddMemberResponseResultUser   `json:"user"`
	JSON   accountMemberAccountMembersAddMemberResponseResultJSON   `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersAddMemberResponseResult]
type accountMemberAccountMembersAddMemberResponseResultJSON struct {
	ID          apijson.Field
	Code        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string                                                             `json:"name,required"`
	Permissions AccountMemberAccountMembersAddMemberResponseResultRolesPermissions `json:"permissions,required"`
	JSON        accountMemberAccountMembersAddMemberResponseResultRoleJSON         `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRoleJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersAddMemberResponseResultRole]
type accountMemberAccountMembersAddMemberResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissions struct {
	Analytics    AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalytics    `json:"analytics"`
	Billing      AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsBilling      `json:"billing"`
	CachePurge   AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurge   `json:"cache_purge"`
	DNS          AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNS          `json:"dns"`
	DNSRecords   AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecords   `json:"dns_records"`
	Lb           AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLb           `json:"lb"`
	Logs         AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogs         `json:"logs"`
	Organization AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganization `json:"organization"`
	Ssl          AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsSsl          `json:"ssl"`
	Waf          AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsWaf          `json:"waf"`
	ZoneSettings AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettings `json:"zone_settings"`
	Zones        AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZones        `json:"zones"`
	JSON         accountMemberAccountMembersAddMemberResponseResultRolesPermissionsJSON         `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsJSON contains
// the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissions]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsJSON struct {
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

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalytics struct {
	Read  bool                                                                            `json:"read"`
	Write bool                                                                            `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalyticsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalyticsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalytics]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsBilling struct {
	Read  bool                                                                          `json:"read"`
	Write bool                                                                          `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsBillingJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsBillingJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsBilling]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurge struct {
	Read  bool                                                                             `json:"read"`
	Write bool                                                                             `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurgeJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurgeJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurge]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNS struct {
	Read  bool                                                                      `json:"read"`
	Write bool                                                                      `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNS]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecords struct {
	Read  bool                                                                             `json:"read"`
	Write bool                                                                             `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecordsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecordsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecords]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLb struct {
	Read  bool                                                                     `json:"read"`
	Write bool                                                                     `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLbJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLbJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLb]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogs struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogs]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganization struct {
	Read  bool                                                                               `json:"read"`
	Write bool                                                                               `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganizationJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganizationJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganization]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsSsl struct {
	Read  bool                                                                      `json:"read"`
	Write bool                                                                      `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsSslJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsSslJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsSsl]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsWaf struct {
	Read  bool                                                                      `json:"read"`
	Write bool                                                                      `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsWafJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsWafJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsWaf]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettings struct {
	Read  bool                                                                               `json:"read"`
	Write bool                                                                               `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettingsJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettingsJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettings]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZones struct {
	Read  bool                                                                        `json:"read"`
	Write bool                                                                        `json:"write"`
	JSON  accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZonesJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZonesJSON
// contains the JSON metadata for the struct
// [AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZones]
type accountMemberAccountMembersAddMemberResponseResultRolesPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultRolesPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMemberAccountMembersAddMemberResponseResultUser struct {
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
	TwoFactorAuthenticationEnabled bool                                                       `json:"two_factor_authentication_enabled"`
	JSON                           accountMemberAccountMembersAddMemberResponseResultUserJSON `json:"-"`
}

// accountMemberAccountMembersAddMemberResponseResultUserJSON contains the JSON
// metadata for the struct [AccountMemberAccountMembersAddMemberResponseResultUser]
type accountMemberAccountMembersAddMemberResponseResultUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountMemberAccountMembersAddMemberResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMemberAccountMembersAddMemberResponseSuccess bool

const (
	AccountMemberAccountMembersAddMemberResponseSuccessTrue AccountMemberAccountMembersAddMemberResponseSuccess = true
)

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
	Ssl          param.Field[AccountMemberUpdateParamsRolesPermissionsSsl]          `json:"ssl"`
	Waf          param.Field[AccountMemberUpdateParamsRolesPermissionsWaf]          `json:"waf"`
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

type AccountMemberUpdateParamsRolesPermissionsSsl struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsRolesPermissionsWaf struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r AccountMemberUpdateParamsRolesPermissionsWaf) MarshalJSON() (data []byte, err error) {
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
