// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// MembershipService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMembershipService] method instead.
type MembershipService struct {
	Options []option.RequestOption
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r *MembershipService) {
	r = &MembershipService{}
	r.Options = opts
	return
}

// Get a specific membership.
func (r *MembershipService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *MembershipGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("memberships/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept or reject this account invitation.
func (r *MembershipService) Update(ctx context.Context, identifier string, body MembershipUpdateParams, opts ...option.RequestOption) (res *MembershipUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("memberships/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Remove the associated member from an account.
func (r *MembershipService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *MembershipDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("memberships/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List memberships of accounts the user can access.
func (r *MembershipService) UserSAccountMembershipsListMemberships(ctx context.Context, query MembershipUserSAccountMembershipsListMembershipsParams, opts ...option.RequestOption) (res *shared.Page[MembershipUserSAccountMembershipsListMembershipsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "memberships"
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

type MembershipGetResponse struct {
	Errors   []MembershipGetResponseError   `json:"errors"`
	Messages []MembershipGetResponseMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success MembershipGetResponseSuccess `json:"success"`
	JSON    membershipGetResponseJSON    `json:"-"`
}

// membershipGetResponseJSON contains the JSON metadata for the struct
// [MembershipGetResponse]
type membershipGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipGetResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    membershipGetResponseErrorJSON `json:"-"`
}

// membershipGetResponseErrorJSON contains the JSON metadata for the struct
// [MembershipGetResponseError]
type membershipGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipGetResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    membershipGetResponseMessageJSON `json:"-"`
}

// membershipGetResponseMessageJSON contains the JSON metadata for the struct
// [MembershipGetResponseMessage]
type membershipGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipGetResponseSuccess bool

const (
	MembershipGetResponseSuccessTrue MembershipGetResponseSuccess = true
)

type MembershipUpdateResponse struct {
	Errors   []MembershipUpdateResponseError   `json:"errors"`
	Messages []MembershipUpdateResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success MembershipUpdateResponseSuccess `json:"success"`
	JSON    membershipUpdateResponseJSON    `json:"-"`
}

// membershipUpdateResponseJSON contains the JSON metadata for the struct
// [MembershipUpdateResponse]
type membershipUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUpdateResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    membershipUpdateResponseErrorJSON `json:"-"`
}

// membershipUpdateResponseErrorJSON contains the JSON metadata for the struct
// [MembershipUpdateResponseError]
type membershipUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUpdateResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    membershipUpdateResponseMessageJSON `json:"-"`
}

// membershipUpdateResponseMessageJSON contains the JSON metadata for the struct
// [MembershipUpdateResponseMessage]
type membershipUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipUpdateResponseSuccess bool

const (
	MembershipUpdateResponseSuccessTrue MembershipUpdateResponseSuccess = true
)

type MembershipDeleteResponse struct {
	Errors   []MembershipDeleteResponseError   `json:"errors"`
	Messages []MembershipDeleteResponseMessage `json:"messages"`
	Result   MembershipDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success MembershipDeleteResponseSuccess `json:"success"`
	JSON    membershipDeleteResponseJSON    `json:"-"`
}

// membershipDeleteResponseJSON contains the JSON metadata for the struct
// [MembershipDeleteResponse]
type membershipDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipDeleteResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    membershipDeleteResponseErrorJSON `json:"-"`
}

// membershipDeleteResponseErrorJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseError]
type membershipDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipDeleteResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    membershipDeleteResponseMessageJSON `json:"-"`
}

// membershipDeleteResponseMessageJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseMessage]
type membershipDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipDeleteResponseResult struct {
	// Membership identifier tag.
	ID   string                             `json:"id"`
	JSON membershipDeleteResponseResultJSON `json:"-"`
}

// membershipDeleteResponseResultJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseResult]
type membershipDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipDeleteResponseSuccess bool

const (
	MembershipDeleteResponseSuccessTrue MembershipDeleteResponseSuccess = true
)

type MembershipUserSAccountMembershipsListMembershipsResponse struct {
	// Membership identifier tag.
	ID      string                                                          `json:"id"`
	Account MembershipUserSAccountMembershipsListMembershipsResponseAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// The unique activation code for the account membership.
	Code        string                                                              `json:"code"`
	Permissions MembershipUserSAccountMembershipsListMembershipsResponsePermissions `json:"permissions"`
	// List of role names for the user at the account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipUserSAccountMembershipsListMembershipsResponseStatus `json:"status"`
	JSON   membershipUserSAccountMembershipsListMembershipsResponseJSON   `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseJSON contains the JSON
// metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponse]
type membershipUserSAccountMembershipsListMembershipsResponseJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Code             apijson.Field
	Permissions      apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponseAccount struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Account settings
	Settings MembershipUserSAccountMembershipsListMembershipsResponseAccountSettings `json:"settings"`
	JSON     membershipUserSAccountMembershipsListMembershipsResponseAccountJSON     `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseAccountJSON contains the
// JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseAccount]
type membershipUserSAccountMembershipsListMembershipsResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	CreatedOn   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Account settings
type MembershipUserSAccountMembershipsListMembershipsResponseAccountSettings struct {
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
	DefaultNameservers MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameservers `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default
	UseAccountCustomNsByDefault bool                                                                        `json:"use_account_custom_ns_by_default"`
	JSON                        membershipUserSAccountMembershipsListMembershipsResponseAccountSettingsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseAccountSettingsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseAccountSettings]
type membershipUserSAccountMembershipsListMembershipsResponseAccountSettingsJSON struct {
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNsByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseAccountSettings) UnmarshalJSON(data []byte) (err error) {
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
type MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameservers string

const (
	MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameserversCloudflareStandard MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameservers = "cloudflare.standard"
	MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameserversCustomAccount      MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameservers = "custom.account"
	MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameserversCustomTenant       MembershipUserSAccountMembershipsListMembershipsResponseAccountSettingsDefaultNameservers = "custom.tenant"
)

type MembershipUserSAccountMembershipsListMembershipsResponsePermissions struct {
	Analytics    MembershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalytics    `json:"analytics"`
	Billing      MembershipUserSAccountMembershipsListMembershipsResponsePermissionsBilling      `json:"billing"`
	CachePurge   MembershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurge   `json:"cache_purge"`
	DNS          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNS          `json:"dns"`
	DNSRecords   MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecords   `json:"dns_records"`
	Lb           MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLb           `json:"lb"`
	Logs         MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLogs         `json:"logs"`
	Organization MembershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganization `json:"organization"`
	Ssl          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSsl          `json:"ssl"`
	Waf          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWaf          `json:"waf"`
	ZoneSettings MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettings `json:"zone_settings"`
	Zones        MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZones        `json:"zones"`
	JSON         membershipUserSAccountMembershipsListMembershipsResponsePermissionsJSON         `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsJSON contains
// the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissions]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsJSON struct {
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

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalytics struct {
	Read  bool                                                                             `json:"read"`
	Write bool                                                                             `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalyticsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalyticsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalytics]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsBilling struct {
	Read  bool                                                                           `json:"read"`
	Write bool                                                                           `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsBillingJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsBillingJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsBilling]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurge struct {
	Read  bool                                                                              `json:"read"`
	Write bool                                                                              `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurgeJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurgeJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurge]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNS struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNS]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecords struct {
	Read  bool                                                                              `json:"read"`
	Write bool                                                                              `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecordsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecordsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecords]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLb struct {
	Read  bool                                                                      `json:"read"`
	Write bool                                                                      `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsLbJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsLbJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLb]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLogs struct {
	Read  bool                                                                        `json:"read"`
	Write bool                                                                        `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsLogsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsLogsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLogs]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganization struct {
	Read  bool                                                                                `json:"read"`
	Write bool                                                                                `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganizationJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganizationJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganization]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSsl struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsSslJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsSslJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSsl]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWaf struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsWafJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsWafJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWaf]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettings struct {
	Read  bool                                                                                `json:"read"`
	Write bool                                                                                `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettingsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettingsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettings]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZones struct {
	Read  bool                                                                         `json:"read"`
	Write bool                                                                         `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsZonesJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsZonesJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZones]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of this membership.
type MembershipUserSAccountMembershipsListMembershipsResponseStatus string

const (
	MembershipUserSAccountMembershipsListMembershipsResponseStatusAccepted MembershipUserSAccountMembershipsListMembershipsResponseStatus = "accepted"
	MembershipUserSAccountMembershipsListMembershipsResponseStatusPending  MembershipUserSAccountMembershipsListMembershipsResponseStatus = "pending"
	MembershipUserSAccountMembershipsListMembershipsResponseStatusRejected MembershipUserSAccountMembershipsListMembershipsResponseStatus = "rejected"
)

type MembershipUpdateParams struct {
	// Whether to accept or reject this account invitation.
	Status param.Field[MembershipUpdateParamsStatus] `json:"status,required"`
}

func (r MembershipUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to accept or reject this account invitation.
type MembershipUpdateParamsStatus string

const (
	MembershipUpdateParamsStatusAccepted MembershipUpdateParamsStatus = "accepted"
	MembershipUpdateParamsStatusRejected MembershipUpdateParamsStatus = "rejected"
)

type MembershipUserSAccountMembershipsListMembershipsParams struct {
	Account param.Field[MembershipUserSAccountMembershipsListMembershipsParamsAccount] `query:"account"`
	// Direction to order memberships.
	Direction param.Field[MembershipUserSAccountMembershipsListMembershipsParamsDirection] `query:"direction"`
	// Account name
	Name param.Field[string] `query:"name"`
	// Field to order memberships by.
	Order param.Field[MembershipUserSAccountMembershipsListMembershipsParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of memberships per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of this membership.
	Status param.Field[MembershipUserSAccountMembershipsListMembershipsParamsStatus] `query:"status"`
}

// URLQuery serializes [MembershipUserSAccountMembershipsListMembershipsParams]'s
// query parameters as `url.Values`.
func (r MembershipUserSAccountMembershipsListMembershipsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MembershipUserSAccountMembershipsListMembershipsParamsAccount struct {
	// Account name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes
// [MembershipUserSAccountMembershipsListMembershipsParamsAccount]'s query
// parameters as `url.Values`.
func (r MembershipUserSAccountMembershipsListMembershipsParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order memberships.
type MembershipUserSAccountMembershipsListMembershipsParamsDirection string

const (
	MembershipUserSAccountMembershipsListMembershipsParamsDirectionAsc  MembershipUserSAccountMembershipsListMembershipsParamsDirection = "asc"
	MembershipUserSAccountMembershipsListMembershipsParamsDirectionDesc MembershipUserSAccountMembershipsListMembershipsParamsDirection = "desc"
)

// Field to order memberships by.
type MembershipUserSAccountMembershipsListMembershipsParamsOrder string

const (
	MembershipUserSAccountMembershipsListMembershipsParamsOrderID          MembershipUserSAccountMembershipsListMembershipsParamsOrder = "id"
	MembershipUserSAccountMembershipsListMembershipsParamsOrderAccountName MembershipUserSAccountMembershipsListMembershipsParamsOrder = "account.name"
	MembershipUserSAccountMembershipsListMembershipsParamsOrderStatus      MembershipUserSAccountMembershipsListMembershipsParamsOrder = "status"
)

// Status of this membership.
type MembershipUserSAccountMembershipsListMembershipsParamsStatus string

const (
	MembershipUserSAccountMembershipsListMembershipsParamsStatusAccepted MembershipUserSAccountMembershipsListMembershipsParamsStatus = "accepted"
	MembershipUserSAccountMembershipsListMembershipsParamsStatusPending  MembershipUserSAccountMembershipsListMembershipsParamsStatus = "pending"
	MembershipUserSAccountMembershipsListMembershipsParamsStatusRejected MembershipUserSAccountMembershipsListMembershipsParamsStatus = "rejected"
)
