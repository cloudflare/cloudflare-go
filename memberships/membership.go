// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
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
func (r *MembershipService) List(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Membership], err error) {
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
func (r *MembershipService) ListAutoPaging(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Membership] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
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
	Analytics    accounts.RolePermissionGrants `json:"analytics"`
	Billing      accounts.RolePermissionGrants `json:"billing"`
	CachePurge   accounts.RolePermissionGrants `json:"cache_purge"`
	DNS          accounts.RolePermissionGrants `json:"dns"`
	DNSRecords   accounts.RolePermissionGrants `json:"dns_records"`
	Lb           accounts.RolePermissionGrants `json:"lb"`
	Logs         accounts.RolePermissionGrants `json:"logs"`
	Organization accounts.RolePermissionGrants `json:"organization"`
	SSL          accounts.RolePermissionGrants `json:"ssl"`
	WAF          accounts.RolePermissionGrants `json:"waf"`
	ZoneSettings accounts.RolePermissionGrants `json:"zone_settings"`
	Zones        accounts.RolePermissionGrants `json:"zones"`
	JSON         membershipPermissionsJSON     `json:"-"`
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

// Status of this membership.
type MembershipStatus string

const (
	MembershipStatusAccepted MembershipStatus = "accepted"
	MembershipStatusPending  MembershipStatus = "pending"
	MembershipStatusRejected MembershipStatus = "rejected"
)

func (r MembershipStatus) IsKnown() bool {
	switch r {
	case MembershipStatusAccepted, MembershipStatusPending, MembershipStatusRejected:
		return true
	}
	return false
}

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

func (r MembershipUpdateParamsStatus) IsKnown() bool {
	switch r {
	case MembershipUpdateParamsStatusAccepted, MembershipUpdateParamsStatusRejected:
		return true
	}
	return false
}

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

func (r MembershipUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r MembershipListParamsDirection) IsKnown() bool {
	switch r {
	case MembershipListParamsDirectionAsc, MembershipListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order memberships by.
type MembershipListParamsOrder string

const (
	MembershipListParamsOrderID          MembershipListParamsOrder = "id"
	MembershipListParamsOrderAccountName MembershipListParamsOrder = "account.name"
	MembershipListParamsOrderStatus      MembershipListParamsOrder = "status"
)

func (r MembershipListParamsOrder) IsKnown() bool {
	switch r {
	case MembershipListParamsOrderID, MembershipListParamsOrderAccountName, MembershipListParamsOrderStatus:
		return true
	}
	return false
}

// Status of this membership.
type MembershipListParamsStatus string

const (
	MembershipListParamsStatusAccepted MembershipListParamsStatus = "accepted"
	MembershipListParamsStatusPending  MembershipListParamsStatus = "pending"
	MembershipListParamsStatusRejected MembershipListParamsStatus = "rejected"
)

func (r MembershipListParamsStatus) IsKnown() bool {
	switch r {
	case MembershipListParamsStatusAccepted, MembershipListParamsStatusPending, MembershipListParamsStatusRejected:
		return true
	}
	return false
}

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

func (r MembershipDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r MembershipGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
