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
func (r *OrganizationInviteService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *OrganizationInviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Change the Roles of a Pending Invite.
func (r *OrganizationInviteService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationInviteUpdateParams, opts ...option.RequestOption) (res *OrganizationInviteUpdateResponse, err error) {
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
func (r *OrganizationInviteService) OrganizationInvitesNewInvitation(ctx context.Context, organizationIdentifier string, body OrganizationInviteOrganizationInvitesNewInvitationParams, opts ...option.RequestOption) (res *OrganizationInviteOrganizationInvitesNewInvitationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all invitations associated with an organization.
func (r *OrganizationInviteService) OrganizationInvitesListInvitations(ctx context.Context, organizationIdentifier string, opts ...option.RequestOption) (res *OrganizationInviteOrganizationInvitesListInvitationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/invites", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OrganizationInviteGetResponse struct {
	Errors   []OrganizationInviteGetResponseError   `json:"errors"`
	Messages []OrganizationInviteGetResponseMessage `json:"messages"`
	Result   interface{}                            `json:"result"`
	// Whether the API call was successful
	Success OrganizationInviteGetResponseSuccess `json:"success"`
	JSON    organizationInviteGetResponseJSON    `json:"-"`
}

// organizationInviteGetResponseJSON contains the JSON metadata for the struct
// [OrganizationInviteGetResponse]
type organizationInviteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    organizationInviteGetResponseErrorJSON `json:"-"`
}

// organizationInviteGetResponseErrorJSON contains the JSON metadata for the struct
// [OrganizationInviteGetResponseError]
type organizationInviteGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    organizationInviteGetResponseMessageJSON `json:"-"`
}

// organizationInviteGetResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationInviteGetResponseMessage]
type organizationInviteGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationInviteGetResponseSuccess bool

const (
	OrganizationInviteGetResponseSuccessTrue OrganizationInviteGetResponseSuccess = true
)

type OrganizationInviteUpdateResponse struct {
	Errors   []OrganizationInviteUpdateResponseError   `json:"errors"`
	Messages []OrganizationInviteUpdateResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success OrganizationInviteUpdateResponseSuccess `json:"success"`
	JSON    organizationInviteUpdateResponseJSON    `json:"-"`
}

// organizationInviteUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationInviteUpdateResponse]
type organizationInviteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    organizationInviteUpdateResponseErrorJSON `json:"-"`
}

// organizationInviteUpdateResponseErrorJSON contains the JSON metadata for the
// struct [OrganizationInviteUpdateResponseError]
type organizationInviteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    organizationInviteUpdateResponseMessageJSON `json:"-"`
}

// organizationInviteUpdateResponseMessageJSON contains the JSON metadata for the
// struct [OrganizationInviteUpdateResponseMessage]
type organizationInviteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationInviteUpdateResponseSuccess bool

const (
	OrganizationInviteUpdateResponseSuccessTrue OrganizationInviteUpdateResponseSuccess = true
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

type OrganizationInviteOrganizationInvitesNewInvitationResponse struct {
	Errors   []OrganizationInviteOrganizationInvitesNewInvitationResponseError   `json:"errors"`
	Messages []OrganizationInviteOrganizationInvitesNewInvitationResponseMessage `json:"messages"`
	Result   interface{}                                                         `json:"result"`
	// Whether the API call was successful
	Success OrganizationInviteOrganizationInvitesNewInvitationResponseSuccess `json:"success"`
	JSON    organizationInviteOrganizationInvitesNewInvitationResponseJSON    `json:"-"`
}

// organizationInviteOrganizationInvitesNewInvitationResponseJSON contains the JSON
// metadata for the struct
// [OrganizationInviteOrganizationInvitesNewInvitationResponse]
type organizationInviteOrganizationInvitesNewInvitationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesNewInvitationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesNewInvitationResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    organizationInviteOrganizationInvitesNewInvitationResponseErrorJSON `json:"-"`
}

// organizationInviteOrganizationInvitesNewInvitationResponseErrorJSON contains the
// JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesNewInvitationResponseError]
type organizationInviteOrganizationInvitesNewInvitationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesNewInvitationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesNewInvitationResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    organizationInviteOrganizationInvitesNewInvitationResponseMessageJSON `json:"-"`
}

// organizationInviteOrganizationInvitesNewInvitationResponseMessageJSON contains
// the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesNewInvitationResponseMessage]
type organizationInviteOrganizationInvitesNewInvitationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesNewInvitationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationInviteOrganizationInvitesNewInvitationResponseSuccess bool

const (
	OrganizationInviteOrganizationInvitesNewInvitationResponseSuccessTrue OrganizationInviteOrganizationInvitesNewInvitationResponseSuccess = true
)

type OrganizationInviteOrganizationInvitesListInvitationsResponse struct {
	Errors     []OrganizationInviteOrganizationInvitesListInvitationsResponseError    `json:"errors"`
	Messages   []OrganizationInviteOrganizationInvitesListInvitationsResponseMessage  `json:"messages"`
	Result     []OrganizationInviteOrganizationInvitesListInvitationsResponseResult   `json:"result"`
	ResultInfo OrganizationInviteOrganizationInvitesListInvitationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OrganizationInviteOrganizationInvitesListInvitationsResponseSuccess `json:"success"`
	JSON    organizationInviteOrganizationInvitesListInvitationsResponseJSON    `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseJSON contains the
// JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponse]
type organizationInviteOrganizationInvitesListInvitationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesListInvitationsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    organizationInviteOrganizationInvitesListInvitationsResponseErrorJSON `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseErrorJSON contains
// the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponseError]
type organizationInviteOrganizationInvitesListInvitationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesListInvitationsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    organizationInviteOrganizationInvitesListInvitationsResponseMessageJSON `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseMessageJSON contains
// the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponseMessage]
type organizationInviteOrganizationInvitesListInvitationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesListInvitationsResponseResult struct {
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
	Roles []OrganizationInviteOrganizationInvitesListInvitationsResponseResultRole `json:"roles"`
	// Current status of the invitation.
	Status OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus `json:"status"`
	JSON   organizationInviteOrganizationInvitesListInvitationsResponseResultJSON   `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseResultJSON contains
// the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponseResult]
type organizationInviteOrganizationInvitesListInvitationsResponseResultJSON struct {
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

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationInviteOrganizationInvitesListInvitationsResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                                                   `json:"permissions,required"`
	JSON        organizationInviteOrganizationInvitesListInvitationsResponseResultRoleJSON `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseResultRoleJSON
// contains the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponseResultRole]
type organizationInviteOrganizationInvitesListInvitationsResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the invitation.
type OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus string

const (
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusPending  OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "pending"
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusAccepted OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "accepted"
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusRejected OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "rejected"
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusCanceled OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "canceled"
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusLeft     OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "left"
	OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatusExpired  OrganizationInviteOrganizationInvitesListInvitationsResponseResultStatus = "expired"
)

type OrganizationInviteOrganizationInvitesListInvitationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       organizationInviteOrganizationInvitesListInvitationsResponseResultInfoJSON `json:"-"`
}

// organizationInviteOrganizationInvitesListInvitationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [OrganizationInviteOrganizationInvitesListInvitationsResponseResultInfo]
type organizationInviteOrganizationInvitesListInvitationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationInviteOrganizationInvitesListInvitationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OrganizationInviteOrganizationInvitesListInvitationsResponseSuccess bool

const (
	OrganizationInviteOrganizationInvitesListInvitationsResponseSuccessTrue OrganizationInviteOrganizationInvitesListInvitationsResponseSuccess = true
)

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
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r OrganizationInviteOrganizationInvitesNewInvitationParamsRole) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
