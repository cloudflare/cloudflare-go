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
func (r *UserInviteService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserInviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/invites/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Responds to an invitation.
func (r *UserInviteService) Update(ctx context.Context, identifier string, body UserInviteUpdateParams, opts ...option.RequestOption) (res *UserInviteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/invites/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists all invitations associated with my user.
func (r *UserInviteService) UserSInvitesListInvitations(ctx context.Context, opts ...option.RequestOption) (res *UserInviteUserSInvitesListInvitationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserInviteGetResponse struct {
	Errors   []UserInviteGetResponseError   `json:"errors"`
	Messages []UserInviteGetResponseMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success UserInviteGetResponseSuccess `json:"success"`
	JSON    userInviteGetResponseJSON    `json:"-"`
}

// userInviteGetResponseJSON contains the JSON metadata for the struct
// [UserInviteGetResponse]
type userInviteGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteGetResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    userInviteGetResponseErrorJSON `json:"-"`
}

// userInviteGetResponseErrorJSON contains the JSON metadata for the struct
// [UserInviteGetResponseError]
type userInviteGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteGetResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    userInviteGetResponseMessageJSON `json:"-"`
}

// userInviteGetResponseMessageJSON contains the JSON metadata for the struct
// [UserInviteGetResponseMessage]
type userInviteGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteGetResponseSuccess bool

const (
	UserInviteGetResponseSuccessTrue UserInviteGetResponseSuccess = true
)

type UserInviteUpdateResponse struct {
	Errors   []UserInviteUpdateResponseError   `json:"errors"`
	Messages []UserInviteUpdateResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success UserInviteUpdateResponseSuccess `json:"success"`
	JSON    userInviteUpdateResponseJSON    `json:"-"`
}

// userInviteUpdateResponseJSON contains the JSON metadata for the struct
// [UserInviteUpdateResponse]
type userInviteUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUpdateResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    userInviteUpdateResponseErrorJSON `json:"-"`
}

// userInviteUpdateResponseErrorJSON contains the JSON metadata for the struct
// [UserInviteUpdateResponseError]
type userInviteUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUpdateResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    userInviteUpdateResponseMessageJSON `json:"-"`
}

// userInviteUpdateResponseMessageJSON contains the JSON metadata for the struct
// [UserInviteUpdateResponseMessage]
type userInviteUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteUpdateResponseSuccess bool

const (
	UserInviteUpdateResponseSuccessTrue UserInviteUpdateResponseSuccess = true
)

type UserInviteUserSInvitesListInvitationsResponse struct {
	Errors     []UserInviteUserSInvitesListInvitationsResponseError    `json:"errors"`
	Messages   []UserInviteUserSInvitesListInvitationsResponseMessage  `json:"messages"`
	Result     []UserInviteUserSInvitesListInvitationsResponseResult   `json:"result"`
	ResultInfo UserInviteUserSInvitesListInvitationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserInviteUserSInvitesListInvitationsResponseSuccess `json:"success"`
	JSON    userInviteUserSInvitesListInvitationsResponseJSON    `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseJSON contains the JSON metadata for
// the struct [UserInviteUserSInvitesListInvitationsResponse]
type userInviteUserSInvitesListInvitationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUserSInvitesListInvitationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUserSInvitesListInvitationsResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userInviteUserSInvitesListInvitationsResponseErrorJSON `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseErrorJSON contains the JSON
// metadata for the struct [UserInviteUserSInvitesListInvitationsResponseError]
type userInviteUserSInvitesListInvitationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUserSInvitesListInvitationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUserSInvitesListInvitationsResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userInviteUserSInvitesListInvitationsResponseMessageJSON `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseMessageJSON contains the JSON
// metadata for the struct [UserInviteUserSInvitesListInvitationsResponseMessage]
type userInviteUserSInvitesListInvitationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUserSInvitesListInvitationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUserSInvitesListInvitationsResponseResult struct {
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
	Roles []UserInviteUserSInvitesListInvitationsResponseResultRole `json:"roles"`
	// Current status of the invitation.
	Status UserInviteUserSInvitesListInvitationsResponseResultStatus `json:"status"`
	JSON   userInviteUserSInvitesListInvitationsResponseResultJSON   `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseResultJSON contains the JSON
// metadata for the struct [UserInviteUserSInvitesListInvitationsResponseResult]
type userInviteUserSInvitesListInvitationsResponseResultJSON struct {
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

func (r *UserInviteUserSInvitesListInvitationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUserSInvitesListInvitationsResponseResultRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                                                    `json:"permissions,required"`
	JSON        userInviteUserSInvitesListInvitationsResponseResultRoleJSON `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseResultRoleJSON contains the JSON
// metadata for the struct
// [UserInviteUserSInvitesListInvitationsResponseResultRole]
type userInviteUserSInvitesListInvitationsResponseResultRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUserSInvitesListInvitationsResponseResultRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the invitation.
type UserInviteUserSInvitesListInvitationsResponseResultStatus string

const (
	UserInviteUserSInvitesListInvitationsResponseResultStatusPending  UserInviteUserSInvitesListInvitationsResponseResultStatus = "pending"
	UserInviteUserSInvitesListInvitationsResponseResultStatusAccepted UserInviteUserSInvitesListInvitationsResponseResultStatus = "accepted"
	UserInviteUserSInvitesListInvitationsResponseResultStatusRejected UserInviteUserSInvitesListInvitationsResponseResultStatus = "rejected"
	UserInviteUserSInvitesListInvitationsResponseResultStatusExpired  UserInviteUserSInvitesListInvitationsResponseResultStatus = "expired"
)

type UserInviteUserSInvitesListInvitationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       userInviteUserSInvitesListInvitationsResponseResultInfoJSON `json:"-"`
}

// userInviteUserSInvitesListInvitationsResponseResultInfoJSON contains the JSON
// metadata for the struct
// [UserInviteUserSInvitesListInvitationsResponseResultInfo]
type userInviteUserSInvitesListInvitationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUserSInvitesListInvitationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteUserSInvitesListInvitationsResponseSuccess bool

const (
	UserInviteUserSInvitesListInvitationsResponseSuccessTrue UserInviteUserSInvitesListInvitationsResponseSuccess = true
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
