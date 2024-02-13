// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Accept or reject this account invitation.
func (r *MembershipService) Update(ctx context.Context, membershipID string, body MembershipUpdateParams, opts ...option.RequestOption) (res *MembershipUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MembershipUpdateResponseEnvelope
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove the associated member from an account.
func (r *MembershipService) Delete(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *MembershipDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MembershipDeleteResponseEnvelope
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific membership.
func (r *MembershipService) Get(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *MembershipGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MembershipGetResponseEnvelope
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List memberships of accounts the user can access.
func (r *MembershipService) UserSAccountMembershipsListMemberships(ctx context.Context, query MembershipUserSAccountMembershipsListMembershipsParams, opts ...option.RequestOption) (res *[]MembershipUserSAccountMembershipsListMembershipsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MembershipUserSAccountMembershipsListMembershipsResponseEnvelope
	path := "memberships"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [MembershipUpdateResponseUnknown] or [shared.UnionString].
type MembershipUpdateResponse interface {
	ImplementsMembershipUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MembershipUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type MembershipDeleteResponse struct {
	// Membership identifier tag.
	ID   string                       `json:"id"`
	JSON membershipDeleteResponseJSON `json:"-"`
}

// membershipDeleteResponseJSON contains the JSON metadata for the struct
// [MembershipDeleteResponse]
type membershipDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [MembershipGetResponseUnknown] or [shared.UnionString].
type MembershipGetResponse interface {
	ImplementsMembershipGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MembershipGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type MembershipUserSAccountMembershipsListMembershipsResponse struct {
	// Membership identifier tag.
	ID      string                                                          `json:"id"`
	Account MembershipUserSAccountMembershipsListMembershipsResponseAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// The unique activation code for the account membership.
	Code string `json:"code"`
	// All access permissions for the user at the account.
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
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
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

// All access permissions for the user at the account.
type MembershipUserSAccountMembershipsListMembershipsResponsePermissions struct {
	Analytics    MembershipUserSAccountMembershipsListMembershipsResponsePermissionsAnalytics    `json:"analytics"`
	Billing      MembershipUserSAccountMembershipsListMembershipsResponsePermissionsBilling      `json:"billing"`
	CachePurge   MembershipUserSAccountMembershipsListMembershipsResponsePermissionsCachePurge   `json:"cache_purge"`
	DNS          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNS          `json:"dns"`
	DNSRecords   MembershipUserSAccountMembershipsListMembershipsResponsePermissionsDNSRecords   `json:"dns_records"`
	Lb           MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLb           `json:"lb"`
	Logs         MembershipUserSAccountMembershipsListMembershipsResponsePermissionsLogs         `json:"logs"`
	Organization MembershipUserSAccountMembershipsListMembershipsResponsePermissionsOrganization `json:"organization"`
	SSL          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSSL          `json:"ssl"`
	WAF          MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWAF          `json:"waf"`
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
	SSL          apijson.Field
	WAF          apijson.Field
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

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSSL struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsSSLJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsSSLJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSSL]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWAF struct {
	Read  bool                                                                       `json:"read"`
	Write bool                                                                       `json:"write"`
	JSON  membershipUserSAccountMembershipsListMembershipsResponsePermissionsWAFJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponsePermissionsWAFJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWAF]
type membershipUserSAccountMembershipsListMembershipsResponsePermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponsePermissionsWAF) UnmarshalJSON(data []byte) (err error) {
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

type MembershipUpdateResponseEnvelope struct {
	Errors   []MembershipUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MembershipUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MembershipUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    membershipUpdateResponseEnvelopeJSON    `json:"-"`
}

// membershipUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipUpdateResponseEnvelope]
type membershipUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    membershipUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// membershipUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipUpdateResponseEnvelopeErrors]
type membershipUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    membershipUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// membershipUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipUpdateResponseEnvelopeMessages]
type membershipUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipUpdateResponseEnvelopeSuccess bool

const (
	MembershipUpdateResponseEnvelopeSuccessTrue MembershipUpdateResponseEnvelopeSuccess = true
)

type MembershipDeleteResponseEnvelope struct {
	Errors   []MembershipDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MembershipDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MembershipDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    membershipDeleteResponseEnvelopeJSON    `json:"-"`
}

// membershipDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseEnvelope]
type membershipDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipDeleteResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    membershipDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// membershipDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipDeleteResponseEnvelopeErrors]
type membershipDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipDeleteResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    membershipDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// membershipDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipDeleteResponseEnvelopeMessages]
type membershipDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipDeleteResponseEnvelopeSuccess bool

const (
	MembershipDeleteResponseEnvelopeSuccessTrue MembershipDeleteResponseEnvelopeSuccess = true
)

type MembershipGetResponseEnvelope struct {
	Errors   []MembershipGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MembershipGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success MembershipGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    membershipGetResponseEnvelopeJSON    `json:"-"`
}

// membershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipGetResponseEnvelope]
type membershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    membershipGetResponseEnvelopeErrorsJSON `json:"-"`
}

// membershipGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipGetResponseEnvelopeErrors]
type membershipGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    membershipGetResponseEnvelopeMessagesJSON `json:"-"`
}

// membershipGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipGetResponseEnvelopeMessages]
type membershipGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipGetResponseEnvelopeSuccess bool

const (
	MembershipGetResponseEnvelopeSuccessTrue MembershipGetResponseEnvelopeSuccess = true
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

type MembershipUserSAccountMembershipsListMembershipsResponseEnvelope struct {
	Errors   []MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessages `json:"messages,required"`
	Result   []MembershipUserSAccountMembershipsListMembershipsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       membershipUserSAccountMembershipsListMembershipsResponseEnvelopeJSON       `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseEnvelope]
type membershipUserSAccountMembershipsListMembershipsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    membershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrorsJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrors]
type membershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    membershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessagesJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessages]
type membershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeSuccess bool

const (
	MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeSuccessTrue MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeSuccess = true
)

type MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                        `json:"total_count"`
	JSON       membershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfoJSON `json:"-"`
}

// membershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfo]
type membershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUserSAccountMembershipsListMembershipsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
