// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationInviteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationInviteService] method
// instead.
type OrganizationInviteService struct {
	Options []option.RequestOption
}

// NewOrganizationInviteService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInviteService(opts ...option.RequestOption) (r *OrganizationInviteService) {
	r = &OrganizationInviteService{}
	r.Options = opts
	return
}

// Get the details of an invitation.
func (r *OrganizationInviteService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *SingleInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change the Roles of a Pending Invite.
func (r *OrganizationInviteService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationInviteUpdateParams, opts ...option.RequestOption) (res *SingleInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Cancel an existing invitation.
func (r *OrganizationInviteService) Delete(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationInviteDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Invite a User to become a Member of an Organization.
func (r *OrganizationInviteService) OrganizationInvitesNewInvitation(ctx context.Context, organizationIdentifier string, body OrganizationInviteOrganizationInvitesNewInvitationParams, opts ...option.RequestOption) (res *SingleInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all invitations associated with an organization.
func (r *OrganizationInviteService) OrganizationInvitesListInvitations(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *CollectionInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CollectionInviteResponse struct {
	Errors     []CollectionInviteResponseError    `json:"errors"`
	Messages   []CollectionInviteResponseMessage  `json:"messages"`
	Result     []CollectionInviteResponseResult   `json:"result"`
	ResultInfo CollectionInviteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CollectionInviteResponseSuccess `json:"success"`
	JSON    collectionInviteResponseJSON    `json:"-"`
}

// collectionInviteResponseJSON contains the JSON metadata for the struct
// [CollectionInviteResponse]
type collectionInviteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionInviteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInviteResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    collectionInviteResponseErrorJSON `json:"-"`
}

// collectionInviteResponseErrorJSON contains the JSON metadata for the struct
// [CollectionInviteResponseError]
type collectionInviteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionInviteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInviteResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    collectionInviteResponseMessageJSON `json:"-"`
}

// collectionInviteResponseMessageJSON contains the JSON metadata for the struct
// [CollectionInviteResponseMessage]
type collectionInviteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionInviteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInviteResponseResult struct {
	// Invite identifier tag.
	ID string `json:"id"`
	// When the invite is no longer active.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The email address of the user who created the invite.
	InvitedBy string `json:"invited_by"`
	// Email address of the user to add to the organization.
	InvitedMemberEmail string `json:"invited_member_email"`
	// ID of the user to add to the organization.
	InvitedMemberID string `json:"invited_member_id,nullable"`
	// When the invite was sent.
	InvitedOn time.Time `json:"invited_on" format:"date-time"`
	// ID of the organization the user will be added to.
	OrganizationID string `json:"organization_id"`
	// Current status of two-factor enforcement on the organization.
	OrganizationIsEnforcingTwofactor bool `json:"organization_is_enforcing_twofactor"`
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// Roles to be assigned to this user.
	Roles []CollectionInviteResponseResultRole `json:"roles"`
	// Current status of the invitation.
	Status CollectionInviteResponseResultStatus `json:"status"`
	JSON   collectionInviteResponseResultJSON   `json:"-"`
}

// collectionInviteResponseResultJSON contains the JSON metadata for the struct
// [CollectionInviteResponseResult]
type collectionInviteResponseResultJSON struct {
	ID                               apijson.Field
	ExpiresOn                        apijson.Field
	InvitedBy                        apijson.Field
	InvitedMemberEmail               apijson.Field
	InvitedMemberID                  apijson.Field
	InvitedOn                        apijson.Field
	OrganizationID                   apijson.Field
	OrganizationIsEnforcingTwofactor apijson.Field
	OrganizationName                 apijson.Field
	Roles                            apijson.Field
	Status                           apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *CollectionInviteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CollectionInviteResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                               `json:"permissions,required"`
	JSON        collectionInviteResponseResultRoleJSON `json:"-"`
}

// collectionInviteResponseResultRoleJSON contains the JSON metadata for the struct
// [CollectionInviteResponseResultRole]
type collectionInviteResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionInviteResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the invitation.
type CollectionInviteResponseResultStatus string

const (
	CollectionInviteResponseResultStatusPending  CollectionInviteResponseResultStatus = "pending"
	CollectionInviteResponseResultStatusAccepted CollectionInviteResponseResultStatus = "accepted"
	CollectionInviteResponseResultStatusRejected CollectionInviteResponseResultStatus = "rejected"
	CollectionInviteResponseResultStatusCanceled CollectionInviteResponseResultStatus = "canceled"
	CollectionInviteResponseResultStatusLeft     CollectionInviteResponseResultStatus = "left"
	CollectionInviteResponseResultStatusExpired  CollectionInviteResponseResultStatus = "expired"
)

type CollectionInviteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       collectionInviteResponseResultInfoJSON `json:"-"`
}

// collectionInviteResponseResultInfoJSON contains the JSON metadata for the struct
// [CollectionInviteResponseResultInfo]
type collectionInviteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CollectionInviteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CollectionInviteResponseSuccess bool

const (
	CollectionInviteResponseSuccessTrue CollectionInviteResponseSuccess = true
)

type SingleInviteResponse struct {
	Errors   []SingleInviteResponseError   `json:"errors"`
	Messages []SingleInviteResponseMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success SingleInviteResponseSuccess `json:"success"`
	JSON    singleInviteResponseJSON    `json:"-"`
}

// singleInviteResponseJSON contains the JSON metadata for the struct
// [SingleInviteResponse]
type singleInviteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleInviteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleInviteResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    singleInviteResponseErrorJSON `json:"-"`
}

// singleInviteResponseErrorJSON contains the JSON metadata for the struct
// [SingleInviteResponseError]
type singleInviteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleInviteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleInviteResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    singleInviteResponseMessageJSON `json:"-"`
}

// singleInviteResponseMessageJSON contains the JSON metadata for the struct
// [SingleInviteResponseMessage]
type singleInviteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleInviteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleInviteResponseSuccess bool

const (
	SingleInviteResponseSuccessTrue SingleInviteResponseSuccess = true
)

type OrganizationInviteDeleteResponse struct {
	// Invite identifier tag.
	ID   string                               `json:"id"`
	JSON organizationInviteDeleteResponseJSON `json:"-"`
}

// organizationInviteDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationInviteDeleteResponse]
type organizationInviteDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteUpdateParams struct {
	// Array of Roles associated with the invited user.
	Roles param.Field[[]string] `json:"roles"`
}

func (r OrganizationInviteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationInviteOrganizationInvitesNewInvitationParams struct {
	// Email address of the user to add to the organization.
	InvitedMemberEmail param.Field[string] `json:"invited_member_email,required"`
	// Array of Roles associated with the invited user.
	Roles param.Field[[]OrganizationInviteOrganizationInvitesNewInvitationParamsRole] `json:"roles,required"`
	// When present and set to true, allows for the invited user to be automatically
	// accepted to the organization. No invitation is sent.
	AutoAccept param.Field[bool] `json:"auto_accept"`
}

func (r OrganizationInviteOrganizationInvitesNewInvitationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationInviteOrganizationInvitesNewInvitationParamsRole struct {
}

func (r OrganizationInviteOrganizationInvitesNewInvitationParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
