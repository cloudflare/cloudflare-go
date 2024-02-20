// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUserService] method instead.
type UserService struct {
	Options                []option.RequestOption
	AuditLogs              *UserAuditLogService
	Billings               *UserBillingService
	Firewall               *UserFirewallService
	Invites                *UserInviteService
	LoadBalancers          *UserLoadBalancerService
	LoadBalancingAnalytics *UserLoadBalancingAnalyticService
	Organizations          *UserOrganizationService
	Subscriptions          *UserSubscriptionService
	Tokens                 *UserTokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.AuditLogs = NewUserAuditLogService(opts...)
	r.Billings = NewUserBillingService(opts...)
	r.Firewall = NewUserFirewallService(opts...)
	r.Invites = NewUserInviteService(opts...)
	r.LoadBalancers = NewUserLoadBalancerService(opts...)
	r.LoadBalancingAnalytics = NewUserLoadBalancingAnalyticService(opts...)
	r.Organizations = NewUserOrganizationService(opts...)
	r.Subscriptions = NewUserSubscriptionService(opts...)
	r.Tokens = NewUserTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *UserUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserUpdateResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) List(ctx context.Context, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserListResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserUpdateResponseUnknown] or [shared.UnionString].
type UserUpdateResponse interface {
	ImplementsUserUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [UserListResponseUnknown] or [shared.UnionString].
type UserListResponse interface {
	ImplementsUserListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserUpdateParams struct {
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

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUpdateResponseEnvelope struct {
	Errors   []UserUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   UserUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    userUpdateResponseEnvelopeJSON    `json:"-"`
}

// userUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserUpdateResponseEnvelope]
type userUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    userUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// userUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserUpdateResponseEnvelopeErrors]
type userUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    userUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// userUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserUpdateResponseEnvelopeMessages]
type userUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserUpdateResponseEnvelopeSuccess bool

const (
	UserUpdateResponseEnvelopeSuccessTrue UserUpdateResponseEnvelopeSuccess = true
)

type UserListResponseEnvelope struct {
	Errors   []UserListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserListResponseEnvelopeMessages `json:"messages,required"`
	Result   UserListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserListResponseEnvelopeSuccess `json:"success,required"`
	JSON    userListResponseEnvelopeJSON    `json:"-"`
}

// userListResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserListResponseEnvelope]
type userListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    userListResponseEnvelopeErrorsJSON `json:"-"`
}

// userListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserListResponseEnvelopeErrors]
type userListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    userListResponseEnvelopeMessagesJSON `json:"-"`
}

// userListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserListResponseEnvelopeMessages]
type userListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserListResponseEnvelopeSuccess bool

const (
	UserListResponseEnvelopeSuccessTrue UserListResponseEnvelopeSuccess = true
)
