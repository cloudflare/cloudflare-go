// File generated from our OpenAPI spec by Stainless.

package memberships

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *MembershipService) List(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[Membership], err error) {
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
func (r *MembershipService) ListAutoPaging(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[Membership] {
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

type Membership struct {
	// Membership identifier tag.
	ID      string           `json:"id"`
	Account accounts.Account `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// The unique activation code for the account membership.
	Code string `json:"code"`
	// All access permissions for the user at the account.
	Permissions MembershipPermissions `json:"permissions"`
	// List of role names for the user at the account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipStatus `json:"status"`
	JSON   membershipJSON   `json:"-"`
}

// membershipJSON contains the JSON metadata for the struct [Membership]
type membershipJSON struct {
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

func (r *Membership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipJSON) RawJSON() string {
	return r.raw
}

// All access permissions for the user at the account.
type MembershipPermissions struct {
	Analytics    MembershipPermissionsAnalytics    `json:"analytics"`
	Billing      MembershipPermissionsBilling      `json:"billing"`
	CachePurge   MembershipPermissionsCachePurge   `json:"cache_purge"`
	DNS          MembershipPermissionsDNS          `json:"dns"`
	DNSRecords   MembershipPermissionsDNSRecords   `json:"dns_records"`
	Lb           MembershipPermissionsLb           `json:"lb"`
	Logs         MembershipPermissionsLogs         `json:"logs"`
	Organization MembershipPermissionsOrganization `json:"organization"`
	SSL          MembershipPermissionsSSL          `json:"ssl"`
	WAF          MembershipPermissionsWAF          `json:"waf"`
	ZoneSettings MembershipPermissionsZoneSettings `json:"zone_settings"`
	Zones        MembershipPermissionsZones        `json:"zones"`
	JSON         membershipPermissionsJSON         `json:"-"`
}

// membershipPermissionsJSON contains the JSON metadata for the struct
// [MembershipPermissions]
type membershipPermissionsJSON struct {
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

func (r *MembershipPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsAnalytics struct {
	Read  bool                               `json:"read"`
	Write bool                               `json:"write"`
	JSON  membershipPermissionsAnalyticsJSON `json:"-"`
}

// membershipPermissionsAnalyticsJSON contains the JSON metadata for the struct
// [MembershipPermissionsAnalytics]
type membershipPermissionsAnalyticsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsAnalyticsJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsBilling struct {
	Read  bool                             `json:"read"`
	Write bool                             `json:"write"`
	JSON  membershipPermissionsBillingJSON `json:"-"`
}

// membershipPermissionsBillingJSON contains the JSON metadata for the struct
// [MembershipPermissionsBilling]
type membershipPermissionsBillingJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsBillingJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsCachePurge struct {
	Read  bool                                `json:"read"`
	Write bool                                `json:"write"`
	JSON  membershipPermissionsCachePurgeJSON `json:"-"`
}

// membershipPermissionsCachePurgeJSON contains the JSON metadata for the struct
// [MembershipPermissionsCachePurge]
type membershipPermissionsCachePurgeJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsCachePurge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsCachePurgeJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsDNS struct {
	Read  bool                         `json:"read"`
	Write bool                         `json:"write"`
	JSON  membershipPermissionsDNSJSON `json:"-"`
}

// membershipPermissionsDNSJSON contains the JSON metadata for the struct
// [MembershipPermissionsDNS]
type membershipPermissionsDNSJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsDNSJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsDNSRecords struct {
	Read  bool                                `json:"read"`
	Write bool                                `json:"write"`
	JSON  membershipPermissionsDNSRecordsJSON `json:"-"`
}

// membershipPermissionsDNSRecordsJSON contains the JSON metadata for the struct
// [MembershipPermissionsDNSRecords]
type membershipPermissionsDNSRecordsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsDNSRecords) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsDNSRecordsJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsLb struct {
	Read  bool                        `json:"read"`
	Write bool                        `json:"write"`
	JSON  membershipPermissionsLbJSON `json:"-"`
}

// membershipPermissionsLbJSON contains the JSON metadata for the struct
// [MembershipPermissionsLb]
type membershipPermissionsLbJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsLb) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsLbJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsLogs struct {
	Read  bool                          `json:"read"`
	Write bool                          `json:"write"`
	JSON  membershipPermissionsLogsJSON `json:"-"`
}

// membershipPermissionsLogsJSON contains the JSON metadata for the struct
// [MembershipPermissionsLogs]
type membershipPermissionsLogsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsLogsJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsOrganization struct {
	Read  bool                                  `json:"read"`
	Write bool                                  `json:"write"`
	JSON  membershipPermissionsOrganizationJSON `json:"-"`
}

// membershipPermissionsOrganizationJSON contains the JSON metadata for the struct
// [MembershipPermissionsOrganization]
type membershipPermissionsOrganizationJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsOrganization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsOrganizationJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsSSL struct {
	Read  bool                         `json:"read"`
	Write bool                         `json:"write"`
	JSON  membershipPermissionsSSLJSON `json:"-"`
}

// membershipPermissionsSSLJSON contains the JSON metadata for the struct
// [MembershipPermissionsSSL]
type membershipPermissionsSSLJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsSSLJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsWAF struct {
	Read  bool                         `json:"read"`
	Write bool                         `json:"write"`
	JSON  membershipPermissionsWAFJSON `json:"-"`
}

// membershipPermissionsWAFJSON contains the JSON metadata for the struct
// [MembershipPermissionsWAF]
type membershipPermissionsWAFJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsWAFJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsZoneSettings struct {
	Read  bool                                  `json:"read"`
	Write bool                                  `json:"write"`
	JSON  membershipPermissionsZoneSettingsJSON `json:"-"`
}

// membershipPermissionsZoneSettingsJSON contains the JSON metadata for the struct
// [MembershipPermissionsZoneSettings]
type membershipPermissionsZoneSettingsJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsZoneSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsZoneSettingsJSON) RawJSON() string {
	return r.raw
}

type MembershipPermissionsZones struct {
	Read  bool                           `json:"read"`
	Write bool                           `json:"write"`
	JSON  membershipPermissionsZonesJSON `json:"-"`
}

// membershipPermissionsZonesJSON contains the JSON metadata for the struct
// [MembershipPermissionsZones]
type membershipPermissionsZonesJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipPermissionsZones) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsZonesJSON) RawJSON() string {
	return r.raw
}

// Status of this membership.
type MembershipStatus string

const (
	MembershipStatusAccepted MembershipStatus = "accepted"
	MembershipStatusPending  MembershipStatus = "pending"
	MembershipStatusRejected MembershipStatus = "rejected"
)

// Union satisfied by [memberships.MembershipUpdateResponseUnknown] or
// [shared.UnionString].
type MembershipUpdateResponse interface {
	ImplementsMembershipsMembershipUpdateResponse()
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

func (r membershipDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [memberships.MembershipGetResponseUnknown] or
// [shared.UnionString].
type MembershipGetResponse interface {
	ImplementsMembershipsMembershipGetResponse()
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

func (r membershipUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r membershipUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r membershipUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r membershipDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r membershipDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r membershipDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r membershipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r membershipGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r membershipGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MembershipGetResponseEnvelopeSuccess bool

const (
	MembershipGetResponseEnvelopeSuccessTrue MembershipGetResponseEnvelopeSuccess = true
)
