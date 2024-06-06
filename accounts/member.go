// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// MemberService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
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
func (r *MemberService) New(ctx context.Context, params MemberNewParams, opts ...option.RequestOption) (res *MemberNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Modify an account member.
func (r *MemberService) Update(ctx context.Context, memberID string, params MemberUpdateParams, opts ...option.RequestOption) (res *MemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List all members of an account.
func (r *MemberService) List(ctx context.Context, params MemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.Member], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/members", params.AccountID)
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
func (r *MemberService) ListAutoPaging(ctx context.Context, params MemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.Member] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a member from an account.
func (r *MemberService) Delete(ctx context.Context, memberID string, body MemberDeleteParams, opts ...option.RequestOption) (res *MemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberDeleteResponseEnvelope
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", body.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *MemberService) Get(ctx context.Context, memberID string, query MemberGetParams, opts ...option.RequestOption) (res *MemberGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Whether the user is a member of the organization or has an invitation pending.
type Status string

const (
	StatusMember  Status = "member"
	StatusInvited Status = "invited"
)

func (r Status) IsKnown() bool {
	switch r {
	case StatusMember, StatusInvited:
		return true
	}
	return false
}

type MemberNewResponse struct {
	Result interface{}           `json:"result,required"`
	JSON   memberNewResponseJSON `json:"-"`
	union  MemberNewResponseUnion
}

// memberNewResponseJSON contains the JSON metadata for the struct
// [MemberNewResponse]
type memberNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r memberNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MemberNewResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r MemberNewResponse) AsUnion() MemberNewResponseUnion {
	return r.union
}

// Union satisfied by [accounts.MemberNewResponseIamAPIResponseCommon] or
// [accounts.MemberNewResponseIamAPIResponseCommon].
type MemberNewResponseUnion interface {
	implementsAccountsMemberNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MemberNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberNewResponseIamAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberNewResponseIamAPIResponseCommon{}),
		},
	)
}

type MemberNewResponseIamAPIResponseCommon struct {
	Result shared.Member                             `json:"result"`
	JSON   memberNewResponseIamAPIResponseCommonJSON `json:"-"`
}

// memberNewResponseIamAPIResponseCommonJSON contains the JSON metadata for the
// struct [MemberNewResponseIamAPIResponseCommon]
type memberNewResponseIamAPIResponseCommonJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseIamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseIamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r MemberNewResponseIamAPIResponseCommon) implementsAccountsMemberNewResponse() {}

type MemberUpdateResponse struct {
	Result interface{}              `json:"result,required"`
	JSON   memberUpdateResponseJSON `json:"-"`
	union  MemberUpdateResponseUnion
}

// memberUpdateResponseJSON contains the JSON metadata for the struct
// [MemberUpdateResponse]
type memberUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r memberUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r MemberUpdateResponse) AsUnion() MemberUpdateResponseUnion {
	return r.union
}

// Union satisfied by [accounts.MemberUpdateResponseIamAPIResponseCommon] or
// [accounts.MemberUpdateResponseIamAPIResponseCommon].
type MemberUpdateResponseUnion interface {
	implementsAccountsMemberUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MemberUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberUpdateResponseIamAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberUpdateResponseIamAPIResponseCommon{}),
		},
	)
}

type MemberUpdateResponseIamAPIResponseCommon struct {
	Result shared.Member                                `json:"result"`
	JSON   memberUpdateResponseIamAPIResponseCommonJSON `json:"-"`
}

// memberUpdateResponseIamAPIResponseCommonJSON contains the JSON metadata for the
// struct [MemberUpdateResponseIamAPIResponseCommon]
type memberUpdateResponseIamAPIResponseCommonJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseIamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseIamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r MemberUpdateResponseIamAPIResponseCommon) implementsAccountsMemberUpdateResponse() {}

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

type MemberGetResponse struct {
	Result interface{}           `json:"result,required"`
	JSON   memberGetResponseJSON `json:"-"`
	union  MemberGetResponseUnion
}

// memberGetResponseJSON contains the JSON metadata for the struct
// [MemberGetResponse]
type memberGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r memberGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *MemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r MemberGetResponse) AsUnion() MemberGetResponseUnion {
	return r.union
}

// Union satisfied by [accounts.MemberGetResponseIamAPIResponseCommon] or
// [accounts.MemberGetResponseIamAPIResponseCommon].
type MemberGetResponseUnion interface {
	implementsAccountsMemberGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MemberGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberGetResponseIamAPIResponseCommon{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MemberGetResponseIamAPIResponseCommon{}),
		},
	)
}

type MemberGetResponseIamAPIResponseCommon struct {
	Result shared.Member                             `json:"result"`
	JSON   memberGetResponseIamAPIResponseCommonJSON `json:"-"`
}

// memberGetResponseIamAPIResponseCommonJSON contains the JSON metadata for the
// struct [MemberGetResponseIamAPIResponseCommon]
type memberGetResponseIamAPIResponseCommonJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseIamAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseIamAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

func (r MemberGetResponseIamAPIResponseCommon) implementsAccountsMemberGetResponse() {}

type MemberNewParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      MemberNewParamsBodyUnion `json:"body,required"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MemberNewParamsBody struct {
	// The contact email address of the user.
	Email    param.Field[string]                    `json:"email,required"`
	Roles    param.Field[interface{}]               `json:"roles,required"`
	Status   param.Field[MemberNewParamsBodyStatus] `json:"status"`
	Policies param.Field[interface{}]               `json:"policies,required"`
}

func (r MemberNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBody) implementsAccountsMemberNewParamsBodyUnion() {}

// Satisfied by [accounts.MemberNewParamsBodyIamCreateMemberWithRoles],
// [accounts.MemberNewParamsBodyIamCreateMemberWithPolicies],
// [MemberNewParamsBody].
type MemberNewParamsBodyUnion interface {
	implementsAccountsMemberNewParamsBodyUnion()
}

type MemberNewParamsBodyIamCreateMemberWithRoles struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                                          `json:"roles,required"`
	Status param.Field[MemberNewParamsBodyIamCreateMemberWithRolesStatus] `json:"status"`
}

func (r MemberNewParamsBodyIamCreateMemberWithRoles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBodyIamCreateMemberWithRoles) implementsAccountsMemberNewParamsBodyUnion() {}

type MemberNewParamsBodyIamCreateMemberWithRolesStatus string

const (
	MemberNewParamsBodyIamCreateMemberWithRolesStatusAccepted MemberNewParamsBodyIamCreateMemberWithRolesStatus = "accepted"
	MemberNewParamsBodyIamCreateMemberWithRolesStatusPending  MemberNewParamsBodyIamCreateMemberWithRolesStatus = "pending"
)

func (r MemberNewParamsBodyIamCreateMemberWithRolesStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIamCreateMemberWithRolesStatusAccepted, MemberNewParamsBodyIamCreateMemberWithRolesStatusPending:
		return true
	}
	return false
}

type MemberNewParamsBodyIamCreateMemberWithPolicies struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of policies associated with this member.
	Policies param.Field[[]MemberNewParamsBodyIamCreateMemberWithPoliciesPolicy] `json:"policies,required"`
	Status   param.Field[MemberNewParamsBodyIamCreateMemberWithPoliciesStatus]   `json:"status"`
}

func (r MemberNewParamsBodyIamCreateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBodyIamCreateMemberWithPolicies) implementsAccountsMemberNewParamsBodyUnion() {
}

type MemberNewParamsBodyIamCreateMemberWithPoliciesPolicy struct {
	// Allow or deny operations against the resources.
	Access param.Field[MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource groups that the policy applies to.
	ResourceGroups param.Field[[]MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r MemberNewParamsBodyIamCreateMemberWithPoliciesPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccess string

const (
	MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccessAllow MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccess = "allow"
	MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccessDeny  MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccess = "deny"
)

func (r MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccess) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccessAllow, MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesAccessDeny:
		return true
	}
	return false
}

// A group of permissions.
type MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesPermissionGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesResourceGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberNewParamsBodyIamCreateMemberWithPoliciesPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsBodyIamCreateMemberWithPoliciesStatus string

const (
	MemberNewParamsBodyIamCreateMemberWithPoliciesStatusAccepted MemberNewParamsBodyIamCreateMemberWithPoliciesStatus = "accepted"
	MemberNewParamsBodyIamCreateMemberWithPoliciesStatusPending  MemberNewParamsBodyIamCreateMemberWithPoliciesStatus = "pending"
)

func (r MemberNewParamsBodyIamCreateMemberWithPoliciesStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIamCreateMemberWithPoliciesStatusAccepted, MemberNewParamsBodyIamCreateMemberWithPoliciesStatusPending:
		return true
	}
	return false
}

type MemberNewParamsBodyStatus string

const (
	MemberNewParamsBodyStatusAccepted MemberNewParamsBodyStatus = "accepted"
	MemberNewParamsBodyStatusPending  MemberNewParamsBodyStatus = "pending"
)

func (r MemberNewParamsBodyStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyStatusAccepted, MemberNewParamsBodyStatusPending:
		return true
	}
	return false
}

type MemberUpdateParams struct {
	AccountID param.Field[string]         `path:"account_id,required"`
	Body      MemberUpdateParamsBodyUnion `json:"body,required"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MemberUpdateParamsBody struct {
	Roles    param.Field[interface{}] `json:"roles,required"`
	User     param.Field[interface{}] `json:"user,required"`
	Policies param.Field[interface{}] `json:"policies,required"`
}

func (r MemberUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberUpdateParamsBody) ImplementsAccountsMemberUpdateParamsBodyUnion() {}

// Satisfied by [shared.MemberParam],
// [accounts.MemberUpdateParamsBodyIamUpdateMemberWithPolicies],
// [MemberUpdateParamsBody].
type MemberUpdateParamsBodyUnion interface {
	ImplementsAccountsMemberUpdateParamsBodyUnion()
}

type MemberUpdateParamsBodyIamUpdateMemberWithPolicies struct {
	// Array of policies associated with this member.
	Policies param.Field[[]MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPolicy] `json:"policies,required"`
}

func (r MemberUpdateParamsBodyIamUpdateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberUpdateParamsBodyIamUpdateMemberWithPolicies) ImplementsAccountsMemberUpdateParamsBodyUnion() {
}

type MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPolicy struct {
	// Allow or deny operations against the resources.
	Access param.Field[MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource groups that the policy applies to.
	ResourceGroups param.Field[[]MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccess string

const (
	MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccessAllow MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccess = "allow"
	MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccessDeny  MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccess = "deny"
)

func (r MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccess) IsKnown() bool {
	switch r {
	case MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccessAllow, MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesAccessDeny:
		return true
	}
	return false
}

// A group of permissions.
type MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesPermissionGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesResourceGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsBodyIamUpdateMemberWithPoliciesPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A member's status in the account.
type MemberUpdateParamsBodyStatus string

const (
	MemberUpdateParamsBodyStatusAccepted MemberUpdateParamsBodyStatus = "accepted"
	MemberUpdateParamsBodyStatusPending  MemberUpdateParamsBodyStatus = "pending"
)

func (r MemberUpdateParamsBodyStatus) IsKnown() bool {
	switch r {
	case MemberUpdateParamsBodyStatusAccepted, MemberUpdateParamsBodyStatusPending:
		return true
	}
	return false
}

type MemberListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type MemberDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  MemberDeleteResponse                `json:"result,nullable"`
	JSON    memberDeleteResponseEnvelopeJSON    `json:"-"`
}

// memberDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelope]
type memberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeJSON) RawJSON() string {
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
	AccountID param.Field[string] `path:"account_id,required"`
}
