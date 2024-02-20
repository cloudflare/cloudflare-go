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

// List memberships of accounts the user can access.
func (r *MembershipService) List(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[MembershipListResponse], err error) {
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

// List memberships of accounts the user can access.
func (r *MembershipService) ListAutoPaging(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[MembershipListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
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

type MembershipListResponse struct {
	// Membership identifier tag.
	ID      string                        `json:"id"`
	Account MembershipListResponseAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// The unique activation code for the account membership.
	Code string `json:"code"`
	// All access permissions for the user at the account.
	Permissions MembershipListResponsePermissions `json:"permissions"`
	// List of role names for the user at the account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipListResponseStatus `json:"status"`
	JSON   membershipListResponseJSON   `json:"-"`
}

// membershipListResponseJSON contains the JSON metadata for the struct
// [MembershipListResponse]
type membershipListResponseJSON struct {
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

func (r *MembershipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponseAccount struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Account settings
	Settings MembershipListResponseAccountSettings `json:"settings"`
	JSON     membershipListResponseAccountJSON     `json:"-"`
}

// membershipListResponseAccountJSON contains the JSON metadata for the struct
// [MembershipListResponseAccount]
type membershipListResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	CreatedOn   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Account settings
type MembershipListResponseAccountSettings struct {
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
	DefaultNameservers MembershipListResponseAccountSettingsDefaultNameservers `json:"default_nameservers"`
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default.
	//
	// Deprecated in favor of `default_nameservers`.
	UseAccountCustomNsByDefault bool                                      `json:"use_account_custom_ns_by_default"`
	JSON                        membershipListResponseAccountSettingsJSON `json:"-"`
}

// membershipListResponseAccountSettingsJSON contains the JSON metadata for the
// struct [MembershipListResponseAccountSettings]
type membershipListResponseAccountSettingsJSON struct {
	DefaultNameservers          apijson.Field
	EnforceTwofactor            apijson.Field
	UseAccountCustomNsByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *MembershipListResponseAccountSettings) UnmarshalJSON(data []byte) (err error) {
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
type MembershipListResponseAccountSettingsDefaultNameservers string

const (
	MembershipListResponseAccountSettingsDefaultNameserversCloudflareStandard MembershipListResponseAccountSettingsDefaultNameservers = "cloudflare.standard"
	MembershipListResponseAccountSettingsDefaultNameserversCustomAccount      MembershipListResponseAccountSettingsDefaultNameservers = "custom.account"
	MembershipListResponseAccountSettingsDefaultNameserversCustomTenant       MembershipListResponseAccountSettingsDefaultNameservers = "custom.tenant"
)

// All access permissions for the user at the account.
type MembershipListResponsePermissions struct {
	Analytics    MembershipListResponsePermissionsAnalytics    `json:"analytics"`
	Billing      MembershipListResponsePermissionsBilling      `json:"billing"`
	CachePurge   MembershipListResponsePermissionsCachePurge   `json:"cache_purge"`
	DNS          MembershipListResponsePermissionsDNS          `json:"dns"`
	DNSRecords   MembershipListResponsePermissionsDNSRecords   `json:"dns_records"`
	Lb           MembershipListResponsePermissionsLb           `json:"lb"`
	Logs         MembershipListResponsePermissionsLogs         `json:"logs"`
	Organization MembershipListResponsePermissionsOrganization `json:"organization"`
	SSL          MembershipListResponsePermissionsSSL          `json:"ssl"`
	WAF          MembershipListResponsePermissionsWAF          `json:"waf"`
	ZoneSettings MembershipListResponsePermissionsZoneSettings `json:"zone_settings"`
	Zones        MembershipListResponsePermissionsZones        `json:"zones"`
	JSON         membershipListResponsePermissionsJSON         `json:"-"`
}

// membershipListResponsePermissionsJSON contains the JSON metadata for the struct
// [MembershipListResponsePermissions]
type membershipListResponsePermissionsJSON struct {
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

func (r *MembershipListResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsAnalytics struct {
	Read  bool                                           `json:"read"`
	Write bool                                           `json:"write"`
	JSON  membershipListResponsePermissionsAnalyticsJSON `json:"-"`
}

// membershipListResponsePermissionsAnalyticsJSON contains the JSON metadata for
// the struct [MembershipListResponsePermissionsAnalytics]
type membershipListResponsePermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsBilling struct {
	Read  bool                                         `json:"read"`
	Write bool                                         `json:"write"`
	JSON  membershipListResponsePermissionsBillingJSON `json:"-"`
}

// membershipListResponsePermissionsBillingJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsBilling]
type membershipListResponsePermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsCachePurge struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  membershipListResponsePermissionsCachePurgeJSON `json:"-"`
}

// membershipListResponsePermissionsCachePurgeJSON contains the JSON metadata for
// the struct [MembershipListResponsePermissionsCachePurge]
type membershipListResponsePermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsDNS struct {
	Read  bool                                     `json:"read"`
	Write bool                                     `json:"write"`
	JSON  membershipListResponsePermissionsDNSJSON `json:"-"`
}

// membershipListResponsePermissionsDNSJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsDNS]
type membershipListResponsePermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsDNSRecords struct {
	Read  bool                                            `json:"read"`
	Write bool                                            `json:"write"`
	JSON  membershipListResponsePermissionsDNSRecordsJSON `json:"-"`
}

// membershipListResponsePermissionsDNSRecordsJSON contains the JSON metadata for
// the struct [MembershipListResponsePermissionsDNSRecords]
type membershipListResponsePermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsLb struct {
	Read  bool                                    `json:"read"`
	Write bool                                    `json:"write"`
	JSON  membershipListResponsePermissionsLbJSON `json:"-"`
}

// membershipListResponsePermissionsLbJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsLb]
type membershipListResponsePermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsLogs struct {
	Read  bool                                      `json:"read"`
	Write bool                                      `json:"write"`
	JSON  membershipListResponsePermissionsLogsJSON `json:"-"`
}

// membershipListResponsePermissionsLogsJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsLogs]
type membershipListResponsePermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsOrganization struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  membershipListResponsePermissionsOrganizationJSON `json:"-"`
}

// membershipListResponsePermissionsOrganizationJSON contains the JSON metadata for
// the struct [MembershipListResponsePermissionsOrganization]
type membershipListResponsePermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsSSL struct {
	Read  bool                                     `json:"read"`
	Write bool                                     `json:"write"`
	JSON  membershipListResponsePermissionsSSLJSON `json:"-"`
}

// membershipListResponsePermissionsSSLJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsSSL]
type membershipListResponsePermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsWAF struct {
	Read  bool                                     `json:"read"`
	Write bool                                     `json:"write"`
	JSON  membershipListResponsePermissionsWAFJSON `json:"-"`
}

// membershipListResponsePermissionsWAFJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsWAF]
type membershipListResponsePermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsZoneSettings struct {
	Read  bool                                              `json:"read"`
	Write bool                                              `json:"write"`
	JSON  membershipListResponsePermissionsZoneSettingsJSON `json:"-"`
}

// membershipListResponsePermissionsZoneSettingsJSON contains the JSON metadata for
// the struct [MembershipListResponsePermissionsZoneSettings]
type membershipListResponsePermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MembershipListResponsePermissionsZones struct {
	Read  bool                                       `json:"read"`
	Write bool                                       `json:"write"`
	JSON  membershipListResponsePermissionsZonesJSON `json:"-"`
}

// membershipListResponsePermissionsZonesJSON contains the JSON metadata for the
// struct [MembershipListResponsePermissionsZones]
type membershipListResponsePermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipListResponsePermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of this membership.
type MembershipListResponseStatus string

const (
	MembershipListResponseStatusAccepted MembershipListResponseStatus = "accepted"
	MembershipListResponseStatusPending  MembershipListResponseStatus = "pending"
	MembershipListResponseStatusRejected MembershipListResponseStatus = "rejected"
)

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

type MembershipListParams struct {
	Account param.Field[MembershipListParamsAccount] `query:"account"`
	// Direction to order memberships.
	Direction param.Field[MembershipListParamsDirection] `query:"direction"`
	// Account name
	Name param.Field[string] `query:"name"`
	// Field to order memberships by.
	Order param.Field[MembershipListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of memberships per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of this membership.
	Status param.Field[MembershipListParamsStatus] `query:"status"`
}

// URLQuery serializes [MembershipListParams]'s query parameters as `url.Values`.
func (r MembershipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MembershipListParamsAccount struct {
	// Account name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [MembershipListParamsAccount]'s query parameters as
// `url.Values`.
func (r MembershipListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order memberships.
type MembershipListParamsDirection string

const (
	MembershipListParamsDirectionAsc  MembershipListParamsDirection = "asc"
	MembershipListParamsDirectionDesc MembershipListParamsDirection = "desc"
)

// Field to order memberships by.
type MembershipListParamsOrder string

const (
	MembershipListParamsOrderID          MembershipListParamsOrder = "id"
	MembershipListParamsOrderAccountName MembershipListParamsOrder = "account.name"
	MembershipListParamsOrderStatus      MembershipListParamsOrder = "status"
)

// Status of this membership.
type MembershipListParamsStatus string

const (
	MembershipListParamsStatusAccepted MembershipListParamsStatus = "accepted"
	MembershipListParamsStatusPending  MembershipListParamsStatus = "pending"
	MembershipListParamsStatusRejected MembershipListParamsStatus = "rejected"
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
