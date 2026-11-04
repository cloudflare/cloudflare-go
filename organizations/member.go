// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package organizations

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
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

// Create a membership that grants access to a specific Organization.
func (r *MemberService) New(ctx context.Context, organizationID string, body MemberNewParams, opts ...option.RequestOption) (res *OrganizationMember, err error) {
	var env MemberNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/members", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List memberships for an Organization.
func (r *MemberService) List(ctx context.Context, organizationID string, query MemberListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OrganizationMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/members", organizationID)
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

// List memberships for an Organization.
func (r *MemberService) ListAutoPaging(ctx context.Context, organizationID string, query MemberListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OrganizationMember] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, organizationID, query, opts...))
}

// Delete a membership to a particular Organization.
func (r *MemberService) Delete(ctx context.Context, organizationID string, memberID string, opts ...option.RequestOption) (res *MemberDeleteResponse, err error) {
	var env MemberDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/members/%s", organizationID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a single membership from an Organization.
func (r *MemberService) Get(ctx context.Context, organizationID string, memberID string, opts ...option.RequestOption) (res *OrganizationMember, err error) {
	var env MemberGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("organizations/%s/members/%s", organizationID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type OrganizationMember struct {
	// Organization Member ID
	ID         string                   `json:"id" api:"required"`
	CreateTime time.Time                `json:"create_time" api:"required" format:"date-time"`
	Meta       map[string]interface{}   `json:"meta" api:"required"`
	Status     OrganizationMemberStatus `json:"status" api:"required"`
	UpdateTime time.Time                `json:"update_time" api:"required" format:"date-time"`
	User       OrganizationMemberUser   `json:"user" api:"required"`
	// Direct policies of the membership. Present only for DSR members enrolled in
	// Organization Roles.
	Policies []OrganizationMemberPolicy `json:"policies"`
	JSON     organizationMemberJSON     `json:"-"`
}

// organizationMemberJSON contains the JSON metadata for the struct
// [OrganizationMember]
type organizationMemberJSON struct {
	ID          apijson.Field
	CreateTime  apijson.Field
	Meta        apijson.Field
	Status      apijson.Field
	UpdateTime  apijson.Field
	User        apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberStatus string

const (
	OrganizationMemberStatusActive   OrganizationMemberStatus = "active"
	OrganizationMemberStatusPending  OrganizationMemberStatus = "pending"
	OrganizationMemberStatusRejected OrganizationMemberStatus = "rejected"
	OrganizationMemberStatusCanceled OrganizationMemberStatus = "canceled"
)

func (r OrganizationMemberStatus) IsKnown() bool {
	switch r {
	case OrganizationMemberStatusActive, OrganizationMemberStatusPending, OrganizationMemberStatusRejected, OrganizationMemberStatusCanceled:
		return true
	}
	return false
}

type OrganizationMemberUser struct {
	ID                             string                     `json:"id" api:"required"`
	Email                          string                     `json:"email" api:"required"`
	Name                           string                     `json:"name" api:"required"`
	TwoFactorAuthenticationEnabled bool                       `json:"two_factor_authentication_enabled" api:"required"`
	JSON                           organizationMemberUserJSON `json:"-"`
}

// organizationMemberUserJSON contains the JSON metadata for the struct
// [OrganizationMemberUser]
type organizationMemberUserJSON struct {
	ID                             apijson.Field
	Email                          apijson.Field
	Name                           apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *OrganizationMemberUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberUserJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberPolicy struct {
	Roles  []OrganizationMemberPoliciesRole  `json:"roles" api:"required"`
	Scopes []OrganizationMemberPoliciesScope `json:"scopes" api:"required"`
	// Omit to create a policy. An unchanged policy keeps its ID; a changed policy is
	// replaced under a new ID.
	ID   string                       `json:"id"`
	JSON organizationMemberPolicyJSON `json:"-"`
}

// organizationMemberPolicyJSON contains the JSON metadata for the struct
// [OrganizationMemberPolicy]
type organizationMemberPolicyJSON struct {
	Roles       apijson.Field
	Scopes      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberPolicyJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberPoliciesRole struct {
	ID   string                             `json:"id" api:"required"`
	JSON organizationMemberPoliciesRoleJSON `json:"-"`
}

// organizationMemberPoliciesRoleJSON contains the JSON metadata for the struct
// [OrganizationMemberPoliciesRole]
type organizationMemberPoliciesRoleJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberPoliciesRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberPoliciesRoleJSON) RawJSON() string {
	return r.raw
}

type OrganizationMemberPoliciesScope struct {
	// An Organization (`com.cloudflare.api.tenant.unit.<tag>`) or Account
	// (`com.cloudflare.api.account.<tag>`) key.
	Key string `json:"key" api:"required"`
	// The route Organization for Account scopes. Optional in requests and always
	// returned for Account scopes.
	SubsetOf []interface{}                       `json:"subset_of"`
	JSON     organizationMemberPoliciesScopeJSON `json:"-"`
}

// organizationMemberPoliciesScopeJSON contains the JSON metadata for the struct
// [OrganizationMemberPoliciesScope]
type organizationMemberPoliciesScopeJSON struct {
	Key         apijson.Field
	SubsetOf    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberPoliciesScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationMemberPoliciesScopeJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponse struct {
	ID   string                   `json:"id" api:"required"`
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

type MemberNewParams struct {
	Member param.Field[MemberNewParamsMember] `json:"member" api:"required"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsMember struct {
	User param.Field[MemberNewParamsMemberUser] `json:"user" api:"required"`
}

func (r MemberNewParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsMemberUser struct {
	Email param.Field[string] `json:"email" api:"required"`
}

func (r MemberNewParamsMemberUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewResponseEnvelope struct {
	Errors   []interface{}                    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo            `json:"messages" api:"required"`
	Result   OrganizationMember               `json:"result" api:"required"`
	Success  MemberNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     memberNewResponseEnvelopeJSON    `json:"-"`
}

// memberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelope]
type memberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeSuccess bool

const (
	MemberNewResponseEnvelopeSuccessTrue MemberNewResponseEnvelopeSuccess = true
)

func (r MemberNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberListParams struct {
	// The amount of items to return. Defaults to 10.
	PageSize param.Field[int64] `query:"page_size"`
	// An opaque token returned from the last list response that when provided will
	// retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent requests
	// with a page token.
	PageToken param.Field[string] `query:"page_token"`
	// Filter the list of memberships by membership status. Defaults to `active`,
	// `pending`, and `rejected`; pass `canceled` explicitly to include canceled
	// memberships.
	Status param.Field[[]MemberListParamsStatus] `query:"status"`
	User   param.Field[MemberListParamsUser]     `query:"user"`
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MemberListParamsStatus string

const (
	MemberListParamsStatusActive   MemberListParamsStatus = "active"
	MemberListParamsStatusPending  MemberListParamsStatus = "pending"
	MemberListParamsStatusRejected MemberListParamsStatus = "rejected"
	MemberListParamsStatusCanceled MemberListParamsStatus = "canceled"
)

func (r MemberListParamsStatus) IsKnown() bool {
	switch r {
	case MemberListParamsStatusActive, MemberListParamsStatusPending, MemberListParamsStatusRejected, MemberListParamsStatusCanceled:
		return true
	}
	return false
}

type MemberListParamsUser struct {
	// Filter the list of memberships for a specific email that ends with a substring.
	Email param.Field[string] `query:"email"`
}

// URLQuery serializes [MemberListParamsUser]'s query parameters as `url.Values`.
func (r MemberListParamsUser) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MemberDeleteResponseEnvelope struct {
	Errors   []interface{}                       `json:"errors" api:"required"`
	Messages []shared.ResponseInfo               `json:"messages" api:"required"`
	Result   MemberDeleteResponse                `json:"result" api:"required"`
	Success  MemberDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     memberDeleteResponseEnvelopeJSON    `json:"-"`
}

// memberDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelope]
type memberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

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

type MemberGetResponseEnvelope struct {
	Errors   []interface{}                    `json:"errors" api:"required"`
	Messages []shared.ResponseInfo            `json:"messages" api:"required"`
	Result   OrganizationMember               `json:"result" api:"required"`
	Success  MemberGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON     memberGetResponseEnvelopeJSON    `json:"-"`
}

// memberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelope]
type memberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeSuccess bool

const (
	MemberGetResponseEnvelopeSuccessTrue MemberGetResponseEnvelopeSuccess = true
)

func (r MemberGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
