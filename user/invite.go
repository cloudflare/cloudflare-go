// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// InviteService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewInviteService] method instead.
type InviteService struct {
	Options []option.RequestOption
}

// NewInviteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInviteService(opts ...option.RequestOption) (r *InviteService) {
	r = &InviteService{}
	r.Options = opts
	return
}

// Lists all invitations associated with my user.
func (r *InviteService) List(ctx context.Context, opts ...option.RequestOption) (res *[]InviteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env InviteListResponseEnvelope
	path := "user/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Responds to an invitation.
func (r *InviteService) Edit(ctx context.Context, inviteID string, body InviteEditParams, opts ...option.RequestOption) (res *InviteEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env InviteEditResponseEnvelope
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of an invitation.
func (r *InviteService) Get(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *InviteGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env InviteGetResponseEnvelope
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InviteListResponse struct {
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
	Roles []accounts.Role `json:"roles"`
	// Current status of the invitation.
	Status InviteListResponseStatus `json:"status"`
	JSON   inviteListResponseJSON   `json:"-"`
}

// inviteListResponseJSON contains the JSON metadata for the struct
// [InviteListResponse]
type inviteListResponseJSON struct {
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

func (r *InviteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteListResponseJSON) RawJSON() string {
	return r.raw
}

// Current status of the invitation.
type InviteListResponseStatus string

const (
	InviteListResponseStatusPending  InviteListResponseStatus = "pending"
	InviteListResponseStatusAccepted InviteListResponseStatus = "accepted"
	InviteListResponseStatusRejected InviteListResponseStatus = "rejected"
	InviteListResponseStatusExpired  InviteListResponseStatus = "expired"
)

func (r InviteListResponseStatus) IsKnown() bool {
	switch r {
	case InviteListResponseStatusPending, InviteListResponseStatusAccepted, InviteListResponseStatusRejected, InviteListResponseStatusExpired:
		return true
	}
	return false
}

// Union satisfied by [user.InviteEditResponseUnknown] or [shared.UnionString].
type InviteEditResponse interface {
	ImplementsUserInviteEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InviteEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [user.InviteGetResponseUnknown] or [shared.UnionString].
type InviteGetResponse interface {
	ImplementsUserInviteGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InviteGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type InviteListResponseEnvelope struct {
	Errors   []InviteListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InviteListResponseEnvelopeMessages `json:"messages,required"`
	Result   []InviteListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    InviteListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo InviteListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       inviteListResponseEnvelopeJSON       `json:"-"`
}

// inviteListResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteListResponseEnvelope]
type inviteListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteListResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    inviteListResponseEnvelopeErrorsJSON `json:"-"`
}

// inviteListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InviteListResponseEnvelopeErrors]
type inviteListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InviteListResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    inviteListResponseEnvelopeMessagesJSON `json:"-"`
}

// inviteListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [InviteListResponseEnvelopeMessages]
type inviteListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InviteListResponseEnvelopeSuccess bool

const (
	InviteListResponseEnvelopeSuccessTrue InviteListResponseEnvelopeSuccess = true
)

func (r InviteListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InviteListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InviteListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       inviteListResponseEnvelopeResultInfoJSON `json:"-"`
}

// inviteListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [InviteListResponseEnvelopeResultInfo]
type inviteListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type InviteEditParams struct {
	// Status of your response to the invitation (rejected or accepted).
	Status param.Field[InviteEditParamsStatus] `json:"status,required"`
}

func (r InviteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of your response to the invitation (rejected or accepted).
type InviteEditParamsStatus string

const (
	InviteEditParamsStatusAccepted InviteEditParamsStatus = "accepted"
	InviteEditParamsStatusRejected InviteEditParamsStatus = "rejected"
)

func (r InviteEditParamsStatus) IsKnown() bool {
	switch r {
	case InviteEditParamsStatusAccepted, InviteEditParamsStatusRejected:
		return true
	}
	return false
}

type InviteEditResponseEnvelope struct {
	Errors   []InviteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InviteEditResponseEnvelopeMessages `json:"messages,required"`
	Result   InviteEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success InviteEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    inviteEditResponseEnvelopeJSON    `json:"-"`
}

// inviteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelope]
type inviteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    inviteEditResponseEnvelopeErrorsJSON `json:"-"`
}

// inviteEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelopeErrors]
type inviteEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    inviteEditResponseEnvelopeMessagesJSON `json:"-"`
}

// inviteEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelopeMessages]
type inviteEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InviteEditResponseEnvelopeSuccess bool

const (
	InviteEditResponseEnvelopeSuccessTrue InviteEditResponseEnvelopeSuccess = true
)

func (r InviteEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InviteEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InviteGetResponseEnvelope struct {
	Errors   []InviteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InviteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   InviteGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success InviteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    inviteGetResponseEnvelopeJSON    `json:"-"`
}

// inviteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelope]
type inviteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    inviteGetResponseEnvelopeErrorsJSON `json:"-"`
}

// inviteGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelopeErrors]
type inviteGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    inviteGetResponseEnvelopeMessagesJSON `json:"-"`
}

// inviteGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelopeMessages]
type inviteGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type InviteGetResponseEnvelopeSuccess bool

const (
	InviteGetResponseEnvelopeSuccessTrue InviteGetResponseEnvelopeSuccess = true
)

func (r InviteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InviteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
