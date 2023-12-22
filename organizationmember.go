// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationMemberService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationMemberService] method
// instead.
type OrganizationMemberService struct {
	Options []option.RequestOption
}

// NewOrganizationMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationMemberService(opts ...option.RequestOption) (r *OrganizationMemberService) {
	r = &OrganizationMemberService{}
	r.Options = opts
	return
}

// Get information about a specific member of an organization.
func (r *OrganizationMemberService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *SingleMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change the Roles of an Organization's Member.
func (r *OrganizationMemberService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationMemberUpdateParams, opts ...option.RequestOption) (res *SingleMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Remove a member from an organization.
func (r *OrganizationMemberService) Delete(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationMemberDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all members of a organization.
func (r *OrganizationMemberService) OrganizationMembersListMembers(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *CollectionMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/members", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionMemberResponse struct {
	Errors     []CollectionMemberResponseError    `json:"errors"`
	Messages   []CollectionMemberResponseMessage  `json:"messages"`
	Result     []CollectionMemberResponseResult   `json:"result"`
	ResultInfo CollectionMemberResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CollectionMemberResponseSuccess `json:"success"`
	JSON    collectionMemberResponseJSON    `json:"-"`
}

// collectionMemberResponseJSON contains the JSON metadata for the struct
// [CollectionMemberResponse]
type collectionMemberResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMemberResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    collectionMemberResponseErrorJSON `json:"-"`
}

// collectionMemberResponseErrorJSON contains the JSON metadata for the struct
// [CollectionMemberResponseError]
type collectionMemberResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMemberResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    collectionMemberResponseMessageJSON `json:"-"`
}

// collectionMemberResponseMessageJSON contains the JSON metadata for the struct
// [CollectionMemberResponseMessage]
type collectionMemberResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMemberResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Member Name.
	Name string `json:"name,required,nullable"`
	// Roles assigned to this Member.
	Roles []CollectionMemberResponseResultRole `json:"roles,required"`
	// A member's status in the organization.
	Status CollectionMemberResponseResultStatus `json:"status,required"`
	JSON   collectionMemberResponseResultJSON   `json:"-"`
}

// collectionMemberResponseResultJSON contains the JSON metadata for the struct
// [CollectionMemberResponseResult]
type collectionMemberResponseResultJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionMemberResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                               `json:"permissions,required"`
	JSON        collectionMemberResponseResultRoleJSON `json:"-"`
}

// collectionMemberResponseResultRoleJSON contains the JSON metadata for the struct
// [CollectionMemberResponseResultRole]
type collectionMemberResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A member's status in the organization.
type CollectionMemberResponseResultStatus string

const (
	CollectionMemberResponseResultStatusAccepted CollectionMemberResponseResultStatus = "accepted"
	CollectionMemberResponseResultStatusInvited  CollectionMemberResponseResultStatus = "invited"
)

type CollectionMemberResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       collectionMemberResponseResultInfoJSON `json:"-"`
}

// collectionMemberResponseResultInfoJSON contains the JSON metadata for the struct
// [CollectionMemberResponseResultInfo]
type collectionMemberResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionMemberResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CollectionMemberResponseSuccess bool

const (
	CollectionMemberResponseSuccessTrue CollectionMemberResponseSuccess = true
)

type SingleMemberResponse struct {
	Errors   []SingleMemberResponseError   `json:"errors"`
	Messages []SingleMemberResponseMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success SingleMemberResponseSuccess `json:"success"`
	JSON    singleMemberResponseJSON    `json:"-"`
}

// singleMemberResponseJSON contains the JSON metadata for the struct
// [SingleMemberResponse]
type singleMemberResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleMemberResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    singleMemberResponseErrorJSON `json:"-"`
}

// singleMemberResponseErrorJSON contains the JSON metadata for the struct
// [SingleMemberResponseError]
type singleMemberResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMemberResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleMemberResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleMemberResponseMessageJSON `json:"-"`
}

// singleMemberResponseMessageJSON contains the JSON metadata for the struct
// [SingleMemberResponseMessage]
type singleMemberResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleMemberResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleMemberResponseSuccess bool

const (
	SingleMemberResponseSuccessTrue SingleMemberResponseSuccess = true
)

type OrganizationMemberDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON organizationMemberDeleteResponseJSON `json:"-"`
}

// organizationMemberDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationMemberDeleteResponse]
type organizationMemberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationMemberUpdateParams struct {
	// Array of Roles associated with this Member.
	Roles param.Field[[]string] `json:"roles"`
}

func (r OrganizationMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
