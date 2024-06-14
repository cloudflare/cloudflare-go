// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
	var env MemberNewResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *MemberService) Update(ctx context.Context, memberID string, params MemberUpdateParams, opts ...option.RequestOption) (res *MemberUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MemberUpdateResponseEnvelope
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *MemberService) List(ctx context.Context, params MemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[MemberListResponse], err error) {
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
func (r *MemberService) ListAutoPaging(ctx context.Context, params MemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[MemberListResponse] {
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
	var env MemberGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []MemberNewResponsePolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []shared.Role `json:"roles"`
	// A member's status in the account.
	Status MemberNewResponseStatus `json:"status"`
	// Details of the user associated to the membership.
	User MemberNewResponseUser `json:"user"`
	JSON memberNewResponseJSON `json:"-"`
}

// memberNewResponseJSON contains the JSON metadata for the struct
// [MemberNewResponse]
type memberNewResponseJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MemberNewResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MemberNewResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MemberNewResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           memberNewResponsePolicyJSON              `json:"-"`
}

// memberNewResponsePolicyJSON contains the JSON metadata for the struct
// [MemberNewResponsePolicy]
type memberNewResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberNewResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MemberNewResponsePoliciesAccess string

const (
	MemberNewResponsePoliciesAccessAllow MemberNewResponsePoliciesAccess = "allow"
	MemberNewResponsePoliciesAccessDeny  MemberNewResponsePoliciesAccess = "deny"
)

func (r MemberNewResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MemberNewResponsePoliciesAccessAllow, MemberNewResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MemberNewResponsePoliciesPermissionGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                                       `json:"name"`
	JSON memberNewResponsePoliciesPermissionGroupJSON `json:"-"`
}

// memberNewResponsePoliciesPermissionGroupJSON contains the JSON metadata for the
// struct [MemberNewResponsePoliciesPermissionGroup]
type memberNewResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MemberNewResponsePoliciesResourceGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MemberNewResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                                     `json:"name"`
	JSON memberNewResponsePoliciesResourceGroupJSON `json:"-"`
}

// memberNewResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [MemberNewResponsePoliciesResourceGroup]
type memberNewResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MemberNewResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MemberNewResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    memberNewResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// memberNewResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata for
// the struct [MemberNewResponsePoliciesResourceGroupsScope]
type memberNewResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MemberNewResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                 `json:"key,required"`
	JSON memberNewResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// memberNewResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [MemberNewResponsePoliciesResourceGroupsScopeObject]
type memberNewResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type MemberNewResponseStatus string

const (
	MemberNewResponseStatusAccepted MemberNewResponseStatus = "accepted"
	MemberNewResponseStatusPending  MemberNewResponseStatus = "pending"
)

func (r MemberNewResponseStatus) IsKnown() bool {
	switch r {
	case MemberNewResponseStatusAccepted, MemberNewResponseStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberNewResponseUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                      `json:"two_factor_authentication_enabled"`
	JSON                           memberNewResponseUserJSON `json:"-"`
}

// memberNewResponseUserJSON contains the JSON metadata for the struct
// [MemberNewResponseUser]
type memberNewResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberNewResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseUserJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponse struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []MemberUpdateResponsePolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []shared.Role `json:"roles"`
	// A member's status in the account.
	Status MemberUpdateResponseStatus `json:"status"`
	// Details of the user associated to the membership.
	User MemberUpdateResponseUser `json:"user"`
	JSON memberUpdateResponseJSON `json:"-"`
}

// memberUpdateResponseJSON contains the JSON metadata for the struct
// [MemberUpdateResponse]
type memberUpdateResponseJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MemberUpdateResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MemberUpdateResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MemberUpdateResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           memberUpdateResponsePolicyJSON              `json:"-"`
}

// memberUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [MemberUpdateResponsePolicy]
type memberUpdateResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MemberUpdateResponsePoliciesAccess string

const (
	MemberUpdateResponsePoliciesAccessAllow MemberUpdateResponsePoliciesAccess = "allow"
	MemberUpdateResponsePoliciesAccessDeny  MemberUpdateResponsePoliciesAccess = "deny"
)

func (r MemberUpdateResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MemberUpdateResponsePoliciesAccessAllow, MemberUpdateResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MemberUpdateResponsePoliciesPermissionGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                                          `json:"name"`
	JSON memberUpdateResponsePoliciesPermissionGroupJSON `json:"-"`
}

// memberUpdateResponsePoliciesPermissionGroupJSON contains the JSON metadata for
// the struct [MemberUpdateResponsePoliciesPermissionGroup]
type memberUpdateResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MemberUpdateResponsePoliciesResourceGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MemberUpdateResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                                        `json:"name"`
	JSON memberUpdateResponsePoliciesResourceGroupJSON `json:"-"`
}

// memberUpdateResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [MemberUpdateResponsePoliciesResourceGroup]
type memberUpdateResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MemberUpdateResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MemberUpdateResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    memberUpdateResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// memberUpdateResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata
// for the struct [MemberUpdateResponsePoliciesResourceGroupsScope]
type memberUpdateResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MemberUpdateResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                    `json:"key,required"`
	JSON memberUpdateResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// memberUpdateResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [MemberUpdateResponsePoliciesResourceGroupsScopeObject]
type memberUpdateResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type MemberUpdateResponseStatus string

const (
	MemberUpdateResponseStatusAccepted MemberUpdateResponseStatus = "accepted"
	MemberUpdateResponseStatusPending  MemberUpdateResponseStatus = "pending"
)

func (r MemberUpdateResponseStatus) IsKnown() bool {
	switch r {
	case MemberUpdateResponseStatusAccepted, MemberUpdateResponseStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberUpdateResponseUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                         `json:"two_factor_authentication_enabled"`
	JSON                           memberUpdateResponseUserJSON `json:"-"`
}

// memberUpdateResponseUserJSON contains the JSON metadata for the struct
// [MemberUpdateResponseUser]
type memberUpdateResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberUpdateResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseUserJSON) RawJSON() string {
	return r.raw
}

type MemberListResponse struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []MemberListResponsePolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []shared.Role `json:"roles"`
	// A member's status in the account.
	Status MemberListResponseStatus `json:"status"`
	// Details of the user associated to the membership.
	User MemberListResponseUser `json:"user"`
	JSON memberListResponseJSON `json:"-"`
}

// memberListResponseJSON contains the JSON metadata for the struct
// [MemberListResponse]
type memberListResponseJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponseJSON) RawJSON() string {
	return r.raw
}

type MemberListResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MemberListResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MemberListResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MemberListResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           memberListResponsePolicyJSON              `json:"-"`
}

// memberListResponsePolicyJSON contains the JSON metadata for the struct
// [MemberListResponsePolicy]
type memberListResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberListResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MemberListResponsePoliciesAccess string

const (
	MemberListResponsePoliciesAccessAllow MemberListResponsePoliciesAccess = "allow"
	MemberListResponsePoliciesAccessDeny  MemberListResponsePoliciesAccess = "deny"
)

func (r MemberListResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MemberListResponsePoliciesAccessAllow, MemberListResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MemberListResponsePoliciesPermissionGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                                        `json:"name"`
	JSON memberListResponsePoliciesPermissionGroupJSON `json:"-"`
}

// memberListResponsePoliciesPermissionGroupJSON contains the JSON metadata for the
// struct [MemberListResponsePoliciesPermissionGroup]
type memberListResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MemberListResponsePoliciesResourceGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MemberListResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                                      `json:"name"`
	JSON memberListResponsePoliciesResourceGroupJSON `json:"-"`
}

// memberListResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [MemberListResponsePoliciesResourceGroup]
type memberListResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MemberListResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MemberListResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    memberListResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// memberListResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata for
// the struct [MemberListResponsePoliciesResourceGroupsScope]
type memberListResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MemberListResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                  `json:"key,required"`
	JSON memberListResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// memberListResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [MemberListResponsePoliciesResourceGroupsScopeObject]
type memberListResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberListResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type MemberListResponseStatus string

const (
	MemberListResponseStatusAccepted MemberListResponseStatus = "accepted"
	MemberListResponseStatusPending  MemberListResponseStatus = "pending"
)

func (r MemberListResponseStatus) IsKnown() bool {
	switch r {
	case MemberListResponseStatusAccepted, MemberListResponseStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberListResponseUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                       `json:"two_factor_authentication_enabled"`
	JSON                           memberListResponseUserJSON `json:"-"`
}

// memberListResponseUserJSON contains the JSON metadata for the struct
// [MemberListResponseUser]
type memberListResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberListResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberListResponseUserJSON) RawJSON() string {
	return r.raw
}

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
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []MemberGetResponsePolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []shared.Role `json:"roles"`
	// A member's status in the account.
	Status MemberGetResponseStatus `json:"status"`
	// Details of the user associated to the membership.
	User MemberGetResponseUser `json:"user"`
	JSON memberGetResponseJSON `json:"-"`
}

// memberGetResponseJSON contains the JSON metadata for the struct
// [MemberGetResponse]
type memberGetResponseJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MemberGetResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MemberGetResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MemberGetResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           memberGetResponsePolicyJSON              `json:"-"`
}

// memberGetResponsePolicyJSON contains the JSON metadata for the struct
// [MemberGetResponsePolicy]
type memberGetResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberGetResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MemberGetResponsePoliciesAccess string

const (
	MemberGetResponsePoliciesAccessAllow MemberGetResponsePoliciesAccess = "allow"
	MemberGetResponsePoliciesAccessDeny  MemberGetResponsePoliciesAccess = "deny"
)

func (r MemberGetResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MemberGetResponsePoliciesAccessAllow, MemberGetResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MemberGetResponsePoliciesPermissionGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta interface{} `json:"meta"`
	// Name of the group.
	Name string                                       `json:"name"`
	JSON memberGetResponsePoliciesPermissionGroupJSON `json:"-"`
}

// memberGetResponsePoliciesPermissionGroupJSON contains the JSON metadata for the
// struct [MemberGetResponsePoliciesPermissionGroup]
type memberGetResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MemberGetResponsePoliciesResourceGroup struct {
	// Identifier of the group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MemberGetResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// Name of the resource group.
	Name string                                     `json:"name"`
	JSON memberGetResponsePoliciesResourceGroupJSON `json:"-"`
}

// memberGetResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [MemberGetResponsePoliciesResourceGroup]
type memberGetResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MemberGetResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MemberGetResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    memberGetResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// memberGetResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata for
// the struct [MemberGetResponsePoliciesResourceGroupsScope]
type memberGetResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MemberGetResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                 `json:"key,required"`
	JSON memberGetResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// memberGetResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [MemberGetResponsePoliciesResourceGroupsScopeObject]
type memberGetResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type MemberGetResponseStatus string

const (
	MemberGetResponseStatusAccepted MemberGetResponseStatus = "accepted"
	MemberGetResponseStatusPending  MemberGetResponseStatus = "pending"
)

func (r MemberGetResponseStatus) IsKnown() bool {
	switch r {
	case MemberGetResponseStatusAccepted, MemberGetResponseStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberGetResponseUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                      `json:"two_factor_authentication_enabled"`
	JSON                           memberGetResponseUserJSON `json:"-"`
}

// memberGetResponseUserJSON contains the JSON metadata for the struct
// [MemberGetResponseUser]
type memberGetResponseUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type MemberNewParams struct {
	// Account identifier tag.
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

type MemberNewResponseEnvelope struct {
	Result MemberNewResponse             `json:"result"`
	JSON   memberNewResponseEnvelopeJSON `json:"-"`
}

// memberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelope]
type memberNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateParams struct {
	// Account identifier tag.
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

type MemberUpdateResponseEnvelope struct {
	Result MemberUpdateResponse             `json:"result"`
	JSON   memberUpdateResponseEnvelopeJSON `json:"-"`
}

// memberUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelope]
type memberUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberListParams struct {
	// Account identifier tag.
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
	// Account identifier tag.
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
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MemberGetResponseEnvelope struct {
	Result MemberGetResponse             `json:"result"`
	JSON   memberGetResponseEnvelopeJSON `json:"-"`
}

// memberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelope]
type memberGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
