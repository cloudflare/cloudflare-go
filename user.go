// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	AuditLogs     *UserAuditLogService
	Billing       *UserBillingService
	Firewall      *UserFirewallService
	Invites       *UserInviteService
	LoadBalancers *UserLoadBalancerService
	Organizations *UserOrganizationService
	Subscriptions *UserSubscriptionService
	Tokens        *UserTokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.AuditLogs = NewUserAuditLogService(opts...)
	r.Billing = NewUserBillingService(opts...)
	r.Firewall = NewUserFirewallService(opts...)
	r.Invites = NewUserInviteService(opts...)
	r.LoadBalancers = NewUserLoadBalancerService(opts...)
	r.Organizations = NewUserOrganizationService(opts...)
	r.Subscriptions = NewUserSubscriptionService(opts...)
	r.Tokens = NewUserTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Edit(ctx context.Context, body UserEditParams, opts ...option.RequestOption) (res *UserEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserEditResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserGetResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserEditResponseUnknown] or [shared.UnionString].
type UserEditResponse interface {
	ImplementsUserEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [UserGetResponseUnknown] or [shared.UnionString].
type UserGetResponse interface {
	ImplementsUserGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserEditParams struct {
	// The country in which the user lives.
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
	// User's telephone number
	Telephone param.Field[string] `json:"telephone"`
	// The zipcode or postal code where the user lives.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r UserEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserEditResponseEnvelope struct {
	Errors   []UserEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserEditResponseEnvelopeMessages `json:"messages,required"`
	Result   UserEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userEditResponseEnvelopeJSON    `json:"-"`
}

// userEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelope]
type userEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    userEditResponseEnvelopeErrorsJSON `json:"-"`
}

// userEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelopeErrors]
type userEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    userEditResponseEnvelopeMessagesJSON `json:"-"`
}

// userEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelopeMessages]
type userEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserEditResponseEnvelopeSuccess bool

const (
	UserEditResponseEnvelopeSuccessTrue UserEditResponseEnvelopeSuccess = true
)

type UserGetResponseEnvelope struct {
	Errors   []UserGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGetResponseEnvelopeMessages `json:"messages,required"`
	Result   UserGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userGetResponseEnvelopeJSON    `json:"-"`
}

// userGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelope]
type userGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    userGetResponseEnvelopeErrorsJSON `json:"-"`
}

// userGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelopeErrors]
type userGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    userGetResponseEnvelopeMessagesJSON `json:"-"`
}

// userGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelopeMessages]
type userGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserGetResponseEnvelopeSuccess bool

const (
	UserGetResponseEnvelopeSuccessTrue UserGetResponseEnvelopeSuccess = true
)
