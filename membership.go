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
func (r *MembershipService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SingleMembershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("memberships/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Accept or reject this account invitation.
func (r *MembershipService) Update(ctx context.Context, identifier string, body MembershipUpdateParams, opts ...option.RequestOption) (res *SingleMembershipResponse, err error) {
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
func (r *MembershipService) UserSAccountMembershipsListMemberships(ctx context.Context, query MembershipUserSAccountMembershipsListMembershipsParams, opts ...option.RequestOption) (res *CollectionMembershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "memberships"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CollectionMembershipResponse struct {
	Errors     []CollectionMembershipResponseError    `json:"errors"`
	Messages   []CollectionMembershipResponseMessage  `json:"messages"`
	Result     []CollectionMembershipResponseResult   `json:"result"`
	ResultInfo CollectionMembershipResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CollectionMembershipResponseSuccess `json:"success"`
	JSON    collectionMembershipResponseJSON    `json:"-"`
}

// collectionMembershipResponseJSON contains the JSON metadata for the struct
// [CollectionMembershipResponse]
type collectionMembershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    collectionMembershipResponseErrorJSON `json:"-"`
}

// collectionMembershipResponseErrorJSON contains the JSON metadata for the struct
// [CollectionMembershipResponseError]
type collectionMembershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    collectionMembershipResponseMessageJSON `json:"-"`
}

// collectionMembershipResponseMessageJSON contains the JSON metadata for the
// struct [CollectionMembershipResponseMessage]
type collectionMembershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResult struct {
	// Membership identifier tag.
	ID      string                                    `json:"id"`
	Account CollectionMembershipResponseResultAccount `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// The unique activation code for the account membership.
	Code        string                                        `json:"code"`
	Permissions CollectionMembershipResponseResultPermissions `json:"permissions"`
	// List of role names for the user at the account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status CollectionMembershipResponseResultStatus `json:"status"`
	JSON   collectionMembershipResponseResultJSON   `json:"-"`
}

// collectionMembershipResponseResultJSON contains the JSON metadata for the struct
// [CollectionMembershipResponseResult]
type collectionMembershipResponseResultJSON struct {
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

func (r *CollectionMembershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultAccount struct {
	// Identifier
	ID string `json:"id,required"`
	// Account name
	Name string `json:"name,required"`
	// Timestamp for the creation of the account
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Account settings
	Settings CollectionMembershipResponseResultAccountSettings `json:"settings"`
	JSON     collectionMembershipResponseResultAccountJSON     `json:"-"`
}

// collectionMembershipResponseResultAccountJSON contains the JSON metadata for the
// struct [CollectionMembershipResponseResultAccount]
type collectionMembershipResponseResultAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	CreatedOn   apijson.Field
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Account settings
type CollectionMembershipResponseResultAccountSettings struct {
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor bool `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default
	UseAccountCustomNsByDefault bool                                                  `json:"use_account_custom_ns_by_default"`
	JSON                        collectionMembershipResponseResultAccountSettingsJSON `json:"-"`
}

// collectionMembershipResponseResultAccountSettingsJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultAccountSettings]
type collectionMembershipResponseResultAccountSettingsJSON struct {
	EnforceTwofactor            apijson.Field
	UseAccountCustomNsByDefault apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissions struct {
	Analytics    CollectionMembershipResponseResultPermissionsAnalytics    `json:"analytics"`
	Billing      CollectionMembershipResponseResultPermissionsBilling      `json:"billing"`
	CachePurge   CollectionMembershipResponseResultPermissionsCachePurge   `json:"cache_purge"`
	DNS          CollectionMembershipResponseResultPermissionsDNS          `json:"dns"`
	DNSRecords   CollectionMembershipResponseResultPermissionsDNSRecords   `json:"dns_records"`
	Lb           CollectionMembershipResponseResultPermissionsLb           `json:"lb"`
	Logs         CollectionMembershipResponseResultPermissionsLogs         `json:"logs"`
	Organization CollectionMembershipResponseResultPermissionsOrganization `json:"organization"`
	Ssl          CollectionMembershipResponseResultPermissionsSsl          `json:"ssl"`
	Waf          CollectionMembershipResponseResultPermissionsWaf          `json:"waf"`
	ZoneSettings CollectionMembershipResponseResultPermissionsZoneSettings `json:"zone_settings"`
	Zones        CollectionMembershipResponseResultPermissionsZones        `json:"zones"`
	JSON         collectionMembershipResponseResultPermissionsJSON         `json:"-"`
}

// collectionMembershipResponseResultPermissionsJSON contains the JSON metadata for
// the struct [CollectionMembershipResponseResultPermissions]
type collectionMembershipResponseResultPermissionsJSON struct {
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

func (r *CollectionMembershipResponseResultPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsAnalytics struct {
	Read  bool                                                       `json:"read"`
	Write bool                                                       `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsAnalyticsJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsAnalyticsJSON contains the JSON
// metadata for the struct [CollectionMembershipResponseResultPermissionsAnalytics]
type collectionMembershipResponseResultPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsBilling struct {
	Read  bool                                                     `json:"read"`
	Write bool                                                     `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsBillingJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsBillingJSON contains the JSON
// metadata for the struct [CollectionMembershipResponseResultPermissionsBilling]
type collectionMembershipResponseResultPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsCachePurge struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsCachePurgeJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsCachePurgeJSON contains the JSON
// metadata for the struct
// [CollectionMembershipResponseResultPermissionsCachePurge]
type collectionMembershipResponseResultPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsDNS struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsDNSJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsDNSJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultPermissionsDNS]
type collectionMembershipResponseResultPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsDNSRecords struct {
	Read  bool                                                        `json:"read"`
	Write bool                                                        `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsDNSRecordsJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsDNSRecordsJSON contains the JSON
// metadata for the struct
// [CollectionMembershipResponseResultPermissionsDNSRecords]
type collectionMembershipResponseResultPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsLb struct {
	Read  bool                                                `json:"read"`
	Write bool                                                `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsLbJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsLbJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultPermissionsLb]
type collectionMembershipResponseResultPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsLogs struct {
	Read  bool                                                  `json:"read"`
	Write bool                                                  `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsLogsJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsLogsJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultPermissionsLogs]
type collectionMembershipResponseResultPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsOrganization struct {
	Read  bool                                                          `json:"read"`
	Write bool                                                          `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsOrganizationJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsOrganizationJSON contains the JSON
// metadata for the struct
// [CollectionMembershipResponseResultPermissionsOrganization]
type collectionMembershipResponseResultPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsSsl struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsSslJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsSslJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultPermissionsSsl]
type collectionMembershipResponseResultPermissionsSslJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsWaf struct {
	Read  bool                                                 `json:"read"`
	Write bool                                                 `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsWafJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsWafJSON contains the JSON metadata
// for the struct [CollectionMembershipResponseResultPermissionsWaf]
type collectionMembershipResponseResultPermissionsWafJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsZoneSettings struct {
	Read  bool                                                          `json:"read"`
	Write bool                                                          `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsZoneSettingsJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsZoneSettingsJSON contains the JSON
// metadata for the struct
// [CollectionMembershipResponseResultPermissionsZoneSettings]
type collectionMembershipResponseResultPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMembershipResponseResultPermissionsZones struct {
	Read  bool                                                   `json:"read"`
	Write bool                                                   `json:"write"`
	JSON  collectionMembershipResponseResultPermissionsZonesJSON `json:"-"`
}

// collectionMembershipResponseResultPermissionsZonesJSON contains the JSON
// metadata for the struct [CollectionMembershipResponseResultPermissionsZones]
type collectionMembershipResponseResultPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of this membership.
type CollectionMembershipResponseResultStatus string

const (
	CollectionMembershipResponseResultStatusAccepted CollectionMembershipResponseResultStatus = "accepted"
	CollectionMembershipResponseResultStatusPending  CollectionMembershipResponseResultStatus = "pending"
	CollectionMembershipResponseResultStatusRejected CollectionMembershipResponseResultStatus = "rejected"
)

type CollectionMembershipResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       collectionMembershipResponseResultInfoJSON `json:"-"`
}

// collectionMembershipResponseResultInfoJSON contains the JSON metadata for the
// struct [CollectionMembershipResponseResultInfo]
type collectionMembershipResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMembershipResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CollectionMembershipResponseSuccess bool

const (
	CollectionMembershipResponseSuccessTrue CollectionMembershipResponseSuccess = true
)

type SingleMembershipResponse struct {
	Errors   []SingleMembershipResponseError   `json:"errors"`
	Messages []SingleMembershipResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success SingleMembershipResponseSuccess `json:"success"`
	JSON    singleMembershipResponseJSON    `json:"-"`
}

// singleMembershipResponseJSON contains the JSON metadata for the struct
// [SingleMembershipResponse]
type singleMembershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleMembershipResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    singleMembershipResponseErrorJSON `json:"-"`
}

// singleMembershipResponseErrorJSON contains the JSON metadata for the struct
// [SingleMembershipResponseError]
type singleMembershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleMembershipResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    singleMembershipResponseMessageJSON `json:"-"`
}

// singleMembershipResponseMessageJSON contains the JSON metadata for the struct
// [SingleMembershipResponseMessage]
type singleMembershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMembershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleMembershipResponseSuccess bool

const (
	SingleMembershipResponseSuccessTrue SingleMembershipResponseSuccess = true
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
