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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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

// Create a membership that grants access to a specific Organization. (Currently in
// Closed Beta - see https://developers.cloudflare.com/fundamentals/organizations/)
func (r *MemberService) New(ctx context.Context, organizationID string, body MemberNewParams, opts ...option.RequestOption) (res *OrganizationMember, err error) {
	var env MemberNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/members", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List memberships for an Organization. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *MemberService) List(ctx context.Context, organizationID string, query MemberListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OrganizationMember], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
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

// List memberships for an Organization. (Currently in Closed Beta - see
// https://developers.cloudflare.com/fundamentals/organizations/)
func (r *MemberService) ListAutoPaging(ctx context.Context, organizationID string, query MemberListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OrganizationMember] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, organizationID, query, opts...))
}

// Delete a membership to a particular Organization. (Currently in Closed Beta -
// see https://developers.cloudflare.com/fundamentals/organizations/)
func (r *MemberService) Delete(ctx context.Context, organizationID string, memberID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/members/%s", organizationID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Retrieve a single membership from an Organization. (Currently in Closed Beta -
// see https://developers.cloudflare.com/fundamentals/organizations/)
func (r *MemberService) Get(ctx context.Context, organizationID string, memberID string, opts ...option.RequestOption) (res *OrganizationMember, err error) {
	var env MemberGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("organizations/%s/members/%s", organizationID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OrganizationMember struct {
	// Organization Member ID
	ID         string                   `json:"id,required"`
	CreateTime time.Time                `json:"create_time,required" format:"date-time"`
	Meta       map[string]interface{}   `json:"meta,required"`
	Status     OrganizationMemberStatus `json:"status,required"`
	UpdateTime time.Time                `json:"update_time,required" format:"date-time"`
	User       OrganizationMemberUser   `json:"user,required"`
	JSON       organizationMemberJSON   `json:"-"`
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
	OrganizationMemberStatusCanceled OrganizationMemberStatus = "canceled"
)

func (r OrganizationMemberStatus) IsKnown() bool {
	switch r {
	case OrganizationMemberStatusActive, OrganizationMemberStatusCanceled:
		return true
	}
	return false
}

type OrganizationMemberUser struct {
	ID                             string                     `json:"id,required"`
	Email                          string                     `json:"email,required"`
	Name                           string                     `json:"name,required"`
	TwoFactorAuthenticationEnabled bool                       `json:"two_factor_authentication_enabled,required"`
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

type MemberNewParams struct {
	Member param.Field[MemberNewParamsMember] `json:"member,required"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsMember struct {
	User   param.Field[MemberNewParamsMemberUser]   `json:"user,required"`
	Status param.Field[MemberNewParamsMemberStatus] `json:"status"`
}

func (r MemberNewParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsMemberUser struct {
	Email param.Field[string] `json:"email,required"`
}

func (r MemberNewParamsMemberUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsMemberStatus string

const (
	MemberNewParamsMemberStatusActive   MemberNewParamsMemberStatus = "active"
	MemberNewParamsMemberStatusCanceled MemberNewParamsMemberStatus = "canceled"
)

func (r MemberNewParamsMemberStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsMemberStatusActive, MemberNewParamsMemberStatusCanceled:
		return true
	}
	return false
}

type MemberNewResponseEnvelope struct {
	Errors   []interface{}                    `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   OrganizationMember               `json:"result,required"`
	Success  MemberNewResponseEnvelopeSuccess `json:"success,required"`
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
	// Filter the list of memberships by membership status.
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
	MemberListParamsStatusCanceled MemberListParamsStatus = "canceled"
)

func (r MemberListParamsStatus) IsKnown() bool {
	switch r {
	case MemberListParamsStatusActive, MemberListParamsStatusCanceled:
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

type MemberGetResponseEnvelope struct {
	Errors   []interface{}                    `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   OrganizationMember               `json:"result,required"`
	Success  MemberGetResponseEnvelopeSuccess `json:"success,required"`
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
