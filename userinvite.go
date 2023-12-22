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

// UserInviteService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewUserInviteService] method instead.
type UserInviteService struct {
	Options []option.RequestOption
}

// NewUserInviteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserInviteService(opts ...option.RequestOption) (r *UserInviteService) {
	r = &UserInviteService{}
	r.Options = opts
	return
}

// Gets the details of an invitation.
func (r *UserInviteService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SingleInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/invites/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Responds to an invitation.
func (r *UserInviteService) Update(ctx context.Context, identifier string, body UserInviteUpdateParams, opts ...option.RequestOption) (res *SingleInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/invites/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists all invitations associated with my user.
func (r *UserInviteService) UserSInvitesListInvitations(ctx context.Context, opts ...option.RequestOption) (res *SchemasCollectionInviteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SchemasCollectionInviteResponse struct {
	Errors     []SchemasCollectionInviteResponseError    `json:"errors"`
	Messages   []SchemasCollectionInviteResponseMessage  `json:"messages"`
	Result     []SchemasCollectionInviteResponseResult   `json:"result"`
	ResultInfo SchemasCollectionInviteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasCollectionInviteResponseSuccess `json:"success"`
	JSON    schemasCollectionInviteResponseJSON    `json:"-"`
}

// schemasCollectionInviteResponseJSON contains the JSON metadata for the struct
// [SchemasCollectionInviteResponse]
type schemasCollectionInviteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionInviteResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasCollectionInviteResponseErrorJSON `json:"-"`
}

// schemasCollectionInviteResponseErrorJSON contains the JSON metadata for the
// struct [SchemasCollectionInviteResponseError]
type schemasCollectionInviteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionInviteResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasCollectionInviteResponseMessageJSON `json:"-"`
}

// schemasCollectionInviteResponseMessageJSON contains the JSON metadata for the
// struct [SchemasCollectionInviteResponseMessage]
type schemasCollectionInviteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionInviteResponseResult struct {
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
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// Roles to be assigned to this user.
	Roles []SchemasCollectionInviteResponseResultRole `json:"roles"`
	// Current status of the invitation.
	Status SchemasCollectionInviteResponseResultStatus `json:"status"`
	JSON   schemasCollectionInviteResponseResultJSON   `json:"-"`
}

// schemasCollectionInviteResponseResultJSON contains the JSON metadata for the
// struct [SchemasCollectionInviteResponseResult]
type schemasCollectionInviteResponseResultJSON struct {
	ID                 apijson.Field
	ExpiresOn          apijson.Field
	InvitedBy          apijson.Field
	InvitedMemberEmail apijson.Field
	InvitedMemberID    apijson.Field
	InvitedOn          apijson.Field
	OrganizationID     apijson.Field
	OrganizationName   apijson.Field
	Roles              apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionInviteResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                      `json:"permissions,required"`
	JSON        schemasCollectionInviteResponseResultRoleJSON `json:"-"`
}

// schemasCollectionInviteResponseResultRoleJSON contains the JSON metadata for the
// struct [SchemasCollectionInviteResponseResultRole]
type schemasCollectionInviteResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the invitation.
type SchemasCollectionInviteResponseResultStatus string

const (
	SchemasCollectionInviteResponseResultStatusPending  SchemasCollectionInviteResponseResultStatus = "pending"
	SchemasCollectionInviteResponseResultStatusAccepted SchemasCollectionInviteResponseResultStatus = "accepted"
	SchemasCollectionInviteResponseResultStatusRejected SchemasCollectionInviteResponseResultStatus = "rejected"
	SchemasCollectionInviteResponseResultStatusExpired  SchemasCollectionInviteResponseResultStatus = "expired"
)

type SchemasCollectionInviteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       schemasCollectionInviteResponseResultInfoJSON `json:"-"`
}

// schemasCollectionInviteResponseResultInfoJSON contains the JSON metadata for the
// struct [SchemasCollectionInviteResponseResultInfo]
type schemasCollectionInviteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionInviteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasCollectionInviteResponseSuccess bool

const (
	SchemasCollectionInviteResponseSuccessTrue SchemasCollectionInviteResponseSuccess = true
)

type UserInviteUpdateParams struct {
	// Status of your response to the invitation (rejected or accepted).
	Status param.Field[UserInviteUpdateParamsStatus] `json:"status,required"`
}

func (r UserInviteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of your response to the invitation (rejected or accepted).
type UserInviteUpdateParamsStatus string

const (
	UserInviteUpdateParamsStatusAccepted UserInviteUpdateParamsStatus = "accepted"
	UserInviteUpdateParamsStatusRejected UserInviteUpdateParamsStatus = "rejected"
)
