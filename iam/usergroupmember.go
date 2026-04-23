// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

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
)

// UserGroupMemberService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserGroupMemberService] method instead.
type UserGroupMemberService struct {
	Options []option.RequestOption
}

// NewUserGroupMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserGroupMemberService(opts ...option.RequestOption) (r *UserGroupMemberService) {
	r = &UserGroupMemberService{}
	r.Options = opts
	return
}

// Add members to a User Group.
func (r *UserGroupMemberService) New(ctx context.Context, userGroupID string, params UserGroupMemberNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[UserGroupMemberNewResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Add members to a User Group.
func (r *UserGroupMemberService) NewAutoPaging(ctx context.Context, userGroupID string, params UserGroupMemberNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[UserGroupMemberNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, userGroupID, params, opts...))
}

// Replace the set of members attached to a User Group.
func (r *UserGroupMemberService) Update(ctx context.Context, userGroupID string, params UserGroupMemberUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[UserGroupMemberUpdateResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Replace the set of members attached to a User Group.
func (r *UserGroupMemberService) UpdateAutoPaging(ctx context.Context, userGroupID string, params UserGroupMemberUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[UserGroupMemberUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, userGroupID, params, opts...))
}

// List all the members attached to a user group.
func (r *UserGroupMemberService) List(ctx context.Context, userGroupID string, params UserGroupMemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserGroupMemberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
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

// List all the members attached to a user group.
func (r *UserGroupMemberService) ListAutoPaging(ctx context.Context, userGroupID string, params UserGroupMemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserGroupMemberListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, userGroupID, params, opts...))
}

// Remove a member from User Group
func (r *UserGroupMemberService) Delete(ctx context.Context, userGroupID string, memberID string, body UserGroupMemberDeleteParams, opts ...option.RequestOption) (res *UserGroupMemberDeleteResponse, err error) {
	var env UserGroupMemberDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members/%s", body.AccountID, userGroupID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Get details of a specific member in a user group.
func (r *UserGroupMemberService) Get(ctx context.Context, userGroupID string, memberID string, query UserGroupMemberGetParams, opts ...option.RequestOption) (res *UserGroupMemberGetResponse, err error) {
	var env UserGroupMemberGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return nil, err
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members/%s", query.AccountID, userGroupID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Member attached to a User Group.
type UserGroupMemberNewResponse struct {
	// Account member identifier.
	ID string `json:"id" api:"required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberNewResponseStatus `json:"status"`
	JSON   userGroupMemberNewResponseJSON   `json:"-"`
}

// userGroupMemberNewResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberNewResponse]
type userGroupMemberNewResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberNewResponseStatus string

const (
	UserGroupMemberNewResponseStatusAccepted UserGroupMemberNewResponseStatus = "accepted"
	UserGroupMemberNewResponseStatusPending  UserGroupMemberNewResponseStatus = "pending"
)

func (r UserGroupMemberNewResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberNewResponseStatusAccepted, UserGroupMemberNewResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberUpdateResponse struct {
	// Account member identifier.
	ID string `json:"id" api:"required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberUpdateResponseStatus `json:"status"`
	JSON   userGroupMemberUpdateResponseJSON   `json:"-"`
}

// userGroupMemberUpdateResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberUpdateResponse]
type userGroupMemberUpdateResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberUpdateResponseStatus string

const (
	UserGroupMemberUpdateResponseStatusAccepted UserGroupMemberUpdateResponseStatus = "accepted"
	UserGroupMemberUpdateResponseStatusPending  UserGroupMemberUpdateResponseStatus = "pending"
)

func (r UserGroupMemberUpdateResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberUpdateResponseStatusAccepted, UserGroupMemberUpdateResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberListResponse struct {
	// Account member identifier.
	ID string `json:"id" api:"required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberListResponseStatus `json:"status"`
	JSON   userGroupMemberListResponseJSON   `json:"-"`
}

// userGroupMemberListResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberListResponse]
type userGroupMemberListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberListResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberListResponseStatus string

const (
	UserGroupMemberListResponseStatusAccepted UserGroupMemberListResponseStatus = "accepted"
	UserGroupMemberListResponseStatusPending  UserGroupMemberListResponseStatus = "pending"
)

func (r UserGroupMemberListResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberListResponseStatusAccepted, UserGroupMemberListResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberDeleteResponse struct {
	// Account member identifier.
	ID string `json:"id" api:"required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberDeleteResponseStatus `json:"status"`
	JSON   userGroupMemberDeleteResponseJSON   `json:"-"`
}

// userGroupMemberDeleteResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberDeleteResponse]
type userGroupMemberDeleteResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberDeleteResponseStatus string

const (
	UserGroupMemberDeleteResponseStatusAccepted UserGroupMemberDeleteResponseStatus = "accepted"
	UserGroupMemberDeleteResponseStatusPending  UserGroupMemberDeleteResponseStatus = "pending"
)

func (r UserGroupMemberDeleteResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberDeleteResponseStatusAccepted, UserGroupMemberDeleteResponseStatusPending:
		return true
	}
	return false
}

// Detailed member information for a User Group member.
type UserGroupMemberGetResponse struct {
	// Account member identifier.
	ID string `json:"id" api:"required"`
	// When the member was added to the user group.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberGetResponseStatus `json:"status"`
	// Details of the user associated with this membership.
	User UserGroupMemberGetResponseUser `json:"user"`
	JSON userGroupMemberGetResponseJSON `json:"-"`
}

// userGroupMemberGetResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberGetResponse]
type userGroupMemberGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberGetResponseStatus string

const (
	UserGroupMemberGetResponseStatusAccepted UserGroupMemberGetResponseStatus = "accepted"
	UserGroupMemberGetResponseStatusPending  UserGroupMemberGetResponseStatus = "pending"
)

func (r UserGroupMemberGetResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberGetResponseStatusAccepted, UserGroupMemberGetResponseStatusPending:
		return true
	}
	return false
}

// Details of the user associated with this membership.
type UserGroupMemberGetResponseUser struct {
	// User identifier tag.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// User's first name.
	FirstName string `json:"first_name"`
	// User's last name.
	LastName string                             `json:"last_name"`
	JSON     userGroupMemberGetResponseUserJSON `json:"-"`
}

// userGroupMemberGetResponseUserJSON contains the JSON metadata for the struct
// [UserGroupMemberGetResponseUser]
type userGroupMemberGetResponseUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	FirstName   apijson.Field
	LastName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]              `path:"account_id" api:"required"`
	Members   []UserGroupMemberNewParamsMember `json:"members" api:"required"`
}

func (r UserGroupMemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Members)
}

type UserGroupMemberNewParamsMember struct {
	// The identifier of an existing account Member.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r UserGroupMemberNewParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupMemberUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Set/Replace members to a user group.
	Members []UserGroupMemberUpdateParamsMember `json:"members" api:"required"`
}

func (r UserGroupMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Members)
}

type UserGroupMemberUpdateParamsMember struct {
	// The identifier of an existing account Member.
	ID param.Field[string] `json:"id" api:"required"`
}

func (r UserGroupMemberUpdateParamsMember) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupMemberListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sort order of returned user group members by email.
	Direction param.Field[UserGroupMemberListParamsDirection] `query:"direction"`
	// A string used for filtering members by partial email match.
	FuzzyEmail param.Field[string] `query:"fuzzyEmail"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserGroupMemberListParams]'s query parameters as
// `url.Values`.
func (r UserGroupMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sort order of returned user group members by email.
type UserGroupMemberListParamsDirection string

const (
	UserGroupMemberListParamsDirectionAsc  UserGroupMemberListParamsDirection = "asc"
	UserGroupMemberListParamsDirectionDesc UserGroupMemberListParamsDirection = "desc"
)

func (r UserGroupMemberListParamsDirection) IsKnown() bool {
	switch r {
	case UserGroupMemberListParamsDirectionAsc, UserGroupMemberListParamsDirectionDesc:
		return true
	}
	return false
}

type UserGroupMemberDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UserGroupMemberDeleteResponseEnvelope struct {
	Errors   []UserGroupMemberDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []UserGroupMemberDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success UserGroupMemberDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	// Member attached to a User Group.
	Result UserGroupMemberDeleteResponse             `json:"result"`
	JSON   userGroupMemberDeleteResponseEnvelopeJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserGroupMemberDeleteResponseEnvelope]
type userGroupMemberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           UserGroupMemberDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupMemberDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserGroupMemberDeleteResponseEnvelopeErrors]
type userGroupMemberDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberDeleteResponseEnvelopeErrorsSource]
type userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           UserGroupMemberDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupMemberDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserGroupMemberDeleteResponseEnvelopeMessages]
type userGroupMemberDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [UserGroupMemberDeleteResponseEnvelopeMessagesSource]
type userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupMemberDeleteResponseEnvelopeSuccess bool

const (
	UserGroupMemberDeleteResponseEnvelopeSuccessTrue UserGroupMemberDeleteResponseEnvelopeSuccess = true
)

func (r UserGroupMemberDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupMemberDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGroupMemberGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type UserGroupMemberGetResponseEnvelope struct {
	Errors   []UserGroupMemberGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []UserGroupMemberGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success UserGroupMemberGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Detailed member information for a User Group member.
	Result UserGroupMemberGetResponse             `json:"result"`
	JSON   userGroupMemberGetResponseEnvelopeJSON `json:"-"`
}

// userGroupMemberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupMemberGetResponseEnvelope]
type userGroupMemberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           UserGroupMemberGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupMemberGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupMemberGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserGroupMemberGetResponseEnvelopeErrors]
type userGroupMemberGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    userGroupMemberGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupMemberGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberGetResponseEnvelopeErrorsSource]
type userGroupMemberGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           UserGroupMemberGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupMemberGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupMemberGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserGroupMemberGetResponseEnvelopeMessages]
type userGroupMemberGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    userGroupMemberGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupMemberGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberGetResponseEnvelopeMessagesSource]
type userGroupMemberGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupMemberGetResponseEnvelopeSuccess bool

const (
	UserGroupMemberGetResponseEnvelopeSuccessTrue UserGroupMemberGetResponseEnvelopeSuccess = true
)

func (r UserGroupMemberGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupMemberGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
