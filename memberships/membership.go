// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package memberships

import (
	"context"
	"errors"
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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// MembershipService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
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
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
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
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
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
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Membership struct {
	// Membership identifier tag.
	ID      string           `json:"id"`
	Account accounts.Account `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
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
	Analytics    shared.PermissionGrant    `json:"analytics"`
	Billing      shared.PermissionGrant    `json:"billing"`
	CachePurge   shared.PermissionGrant    `json:"cache_purge"`
	DNS          shared.PermissionGrant    `json:"dns"`
	DNSRecords   shared.PermissionGrant    `json:"dns_records"`
	LB           shared.PermissionGrant    `json:"lb"`
	Logs         shared.PermissionGrant    `json:"logs"`
	Organization shared.PermissionGrant    `json:"organization"`
	SSL          shared.PermissionGrant    `json:"ssl"`
	WAF          shared.PermissionGrant    `json:"waf"`
	ZoneSettings shared.PermissionGrant    `json:"zone_settings"`
	Zones        shared.PermissionGrant    `json:"zones"`
	JSON         membershipPermissionsJSON `json:"-"`
}

// membershipPermissionsJSON contains the JSON metadata for the struct
// [MembershipPermissions]
type membershipPermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	LB           apijson.Field
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

type MembershipUpdateResponse struct {
	Result interface{}                  `json:"result,required"`
	JSON   membershipUpdateResponseJSON `json:"-"`
	union  MembershipUpdateResponseUnion
}

// membershipUpdateResponseJSON contains the JSON metadata for the struct
// [MembershipUpdateResponse]
type membershipUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r membershipUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MembershipUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r MembershipUpdateResponse) AsUnion() MembershipUpdateResponseUnion {
	return r.union
}

// Union satisfied by [memberships.MembershipUpdateResponseIamAPIResponseCommon] or
// [memberships.MembershipUpdateResponseIamAPIResponseCommon].
type MembershipUpdateResponseUnion interface {
	implementsMembershipsMembershipUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MembershipUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipUpdateResponseIamAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipUpdateResponseIamAPIResponseCommon{}),
		},
	)
}

type MembershipUpdateResponseIamAPIResponseCommon struct {
	Result Membership                                       `json:"result"`
	JSON   membershipUpdateResponseIamAPIResponseCommonJSON `json:"-"`
}

// membershipUpdateResponseIamAPIResponseCommonJSON contains the JSON metadata for
// the struct [MembershipUpdateResponseIamAPIResponseCommon]
type membershipUpdateResponseIamAPIResponseCommonJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseIamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseIamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r MembershipUpdateResponseIamAPIResponseCommon) implementsMembershipsMembershipUpdateResponse() {
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

type MembershipGetResponse struct {
	Result interface{}               `json:"result,required"`
	JSON   membershipGetResponseJSON `json:"-"`
	union  MembershipGetResponseUnion
}

// membershipGetResponseJSON contains the JSON metadata for the struct
// [MembershipGetResponse]
type membershipGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r membershipGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MembershipGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r MembershipGetResponse) AsUnion() MembershipGetResponseUnion {
	return r.union
}

// Union satisfied by [memberships.MembershipGetResponseIamAPIResponseCommon] or
// [memberships.MembershipGetResponseIamAPIResponseCommon].
type MembershipGetResponseUnion interface {
	implementsMembershipsMembershipGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MembershipGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipGetResponseIamAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MembershipGetResponseIamAPIResponseCommon{}),
		},
	)
}

type MembershipGetResponseIamAPIResponseCommon struct {
	Result Membership                                    `json:"result"`
	JSON   membershipGetResponseIamAPIResponseCommonJSON `json:"-"`
}

// membershipGetResponseIamAPIResponseCommonJSON contains the JSON metadata for the
// struct [MembershipGetResponseIamAPIResponseCommon]
type membershipGetResponseIamAPIResponseCommonJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseIamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseIamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r MembershipGetResponseIamAPIResponseCommon) implementsMembershipsMembershipGetResponse() {}

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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MembershipDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  MembershipDeleteResponse                `json:"result"`
	JSON    membershipDeleteResponseEnvelopeJSON    `json:"-"`
}

// membershipDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseEnvelope]
type membershipDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeJSON) RawJSON() string {
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
