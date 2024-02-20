// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Responds to an invitation.
func (r *UserInviteService) Update(ctx context.Context, inviteID string, body UserInviteUpdateParams, opts ...option.RequestOption) (res *UserInviteUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserInviteUpdateResponseEnvelope
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all invitations associated with my user.
func (r *UserInviteService) List(ctx context.Context, opts ...option.RequestOption) (res *[]UserInviteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserInviteListResponseEnvelope
	path := "user/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of an invitation.
func (r *UserInviteService) Get(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *UserInviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserInviteGetResponseEnvelope
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserInviteUpdateResponseUnknown] or [shared.UnionString].
type UserInviteUpdateResponse interface {
	ImplementsUserInviteUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserInviteUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserInviteListResponse struct {
	// ID of the user to add to the organization.
	InvitedMemberID string `json:"invited_member_id,required,nullable"`
	// ID of the organization the user will be added to.
	OrganizationID string `json:"organization_id,required"`
	// Invite identifier tag.
	ID string `json:"id"`
	// When the invite is no longer active.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The email address of the user who created the invite.
	InvitedBy string `json:"invited_by"`
	// Email address of the user to add to the organization.
	InvitedMemberEmail string `json:"invited_member_email"`
	// When the invite was sent.
	InvitedOn time.Time `json:"invited_on" format:"date-time"`
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// Roles to be assigned to this user.
	Roles []UserInviteListResponseRole `json:"roles"`
	// Current status of the invitation.
	Status UserInviteListResponseStatus `json:"status"`
	JSON   userInviteListResponseJSON   `json:"-"`
}

// userInviteListResponseJSON contains the JSON metadata for the struct
// [UserInviteListResponse]
type userInviteListResponseJSON struct {
	InvitedMemberID    apijson.Field
	OrganizationID     apijson.Field
	ID                 apijson.Field
	ExpiresOn          apijson.Field
	InvitedBy          apijson.Field
	InvitedMemberEmail apijson.Field
	InvitedOn          apijson.Field
	OrganizationName   apijson.Field
	Roles              apijson.Field
	Status             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserInviteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteListResponseRole struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role Name.
	Name string `json:"name,required"`
	// Access permissions for this User.
	Permissions []string                       `json:"permissions,required"`
	JSON        userInviteListResponseRoleJSON `json:"-"`
}

// userInviteListResponseRoleJSON contains the JSON metadata for the struct
// [UserInviteListResponseRole]
type userInviteListResponseRoleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the invitation.
type UserInviteListResponseStatus string

const (
	UserInviteListResponseStatusPending  UserInviteListResponseStatus = "pending"
	UserInviteListResponseStatusAccepted UserInviteListResponseStatus = "accepted"
	UserInviteListResponseStatusRejected UserInviteListResponseStatus = "rejected"
	UserInviteListResponseStatusExpired  UserInviteListResponseStatus = "expired"
)

// Union satisfied by [UserInviteGetResponseUnknown] or [shared.UnionString].
type UserInviteGetResponse interface {
	ImplementsUserInviteGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserInviteGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

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

type UserInviteUpdateResponseEnvelope struct {
	Errors   []UserInviteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserInviteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserInviteUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserInviteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userInviteUpdateResponseEnvelopeJSON    `json:"-"`
}

// userInviteUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserInviteUpdateResponseEnvelope]
type userInviteUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    userInviteUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userInviteUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserInviteUpdateResponseEnvelopeErrors]
type userInviteUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    userInviteUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userInviteUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserInviteUpdateResponseEnvelopeMessages]
type userInviteUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteUpdateResponseEnvelopeSuccess bool

const (
	UserInviteUpdateResponseEnvelopeSuccessTrue UserInviteUpdateResponseEnvelopeSuccess = true
)

type UserInviteListResponseEnvelope struct {
	Errors   []UserInviteListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserInviteListResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserInviteListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserInviteListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserInviteListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userInviteListResponseEnvelopeJSON       `json:"-"`
}

// userInviteListResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserInviteListResponseEnvelope]
type userInviteListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    userInviteListResponseEnvelopeErrorsJSON `json:"-"`
}

// userInviteListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserInviteListResponseEnvelopeErrors]
type userInviteListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    userInviteListResponseEnvelopeMessagesJSON `json:"-"`
}

// userInviteListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserInviteListResponseEnvelopeMessages]
type userInviteListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteListResponseEnvelopeSuccess bool

const (
	UserInviteListResponseEnvelopeSuccessTrue UserInviteListResponseEnvelopeSuccess = true
)

type UserInviteListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       userInviteListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userInviteListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [UserInviteListResponseEnvelopeResultInfo]
type userInviteListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteGetResponseEnvelope struct {
	Errors   []UserInviteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserInviteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserInviteGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserInviteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userInviteGetResponseEnvelopeJSON    `json:"-"`
}

// userInviteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserInviteGetResponseEnvelope]
type userInviteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    userInviteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userInviteGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserInviteGetResponseEnvelopeErrors]
type userInviteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    userInviteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userInviteGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserInviteGetResponseEnvelopeMessages]
type userInviteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserInviteGetResponseEnvelopeSuccess bool

const (
	UserInviteGetResponseEnvelopeSuccessTrue UserInviteGetResponseEnvelopeSuccess = true
)
