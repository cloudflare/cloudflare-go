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
	Firewalls              *UserFirewallService
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
	r.Firewalls = NewUserFirewallService(opts...)
	r.Invites = NewUserInviteService(opts...)
	r.LoadBalancers = NewUserLoadBalancerService(opts...)
	r.LoadBalancingAnalytics = NewUserLoadBalancingAnalyticService(opts...)
	r.Organizations = NewUserOrganizationService(opts...)
	r.Subscriptions = NewUserSubscriptionService(opts...)
	r.Tokens = NewUserTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) UserEditUser(ctx context.Context, body UserUserEditUserParams, opts ...option.RequestOption) (res *UserUserEditUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserUserEditUserResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) UserUserDetails(ctx context.Context, opts ...option.RequestOption) (res *UserUserUserDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserUserUserDetailsResponseEnvelope
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [UserUserEditUserResponseUnknown] or [shared.UnionString].
type UserUserEditUserResponse interface {
	ImplementsUserUserEditUserResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserUserEditUserResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [UserUserUserDetailsResponseUnknown] or [shared.UnionString].
type UserUserUserDetailsResponse interface {
	ImplementsUserUserUserDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserUserUserDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserUserEditUserParams struct {
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

func (r UserUserEditUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserUserEditUserResponseEnvelope struct {
	Errors   []UserUserEditUserResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserUserEditUserResponseEnvelopeMessages `json:"messages,required"`
	Result   UserUserEditUserResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserUserEditUserResponseEnvelopeSuccess `json:"success,required"`
	JSON    userUserEditUserResponseEnvelopeJSON    `json:"-"`
}

// userUserEditUserResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserUserEditUserResponseEnvelope]
type userUserEditUserResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserEditUserResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    userUserEditUserResponseEnvelopeErrorsJSON `json:"-"`
}

// userUserEditUserResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserUserEditUserResponseEnvelopeErrors]
type userUserEditUserResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserEditUserResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    userUserEditUserResponseEnvelopeMessagesJSON `json:"-"`
}

// userUserEditUserResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserUserEditUserResponseEnvelopeMessages]
type userUserEditUserResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserUserEditUserResponseEnvelopeSuccess bool

const (
	UserUserEditUserResponseEnvelopeSuccessTrue UserUserEditUserResponseEnvelopeSuccess = true
)

type UserUserUserDetailsResponseEnvelope struct {
	Errors   []UserUserUserDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserUserUserDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   UserUserUserDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success UserUserUserDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    userUserUserDetailsResponseEnvelopeJSON    `json:"-"`
}

// userUserUserDetailsResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserUserUserDetailsResponseEnvelope]
type userUserUserDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserUserDetailsResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userUserUserDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// userUserUserDetailsResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserUserUserDetailsResponseEnvelopeErrors]
type userUserUserDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserUserDetailsResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    userUserUserDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// userUserUserDetailsResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserUserUserDetailsResponseEnvelopeMessages]
type userUserUserDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserUserUserDetailsResponseEnvelopeSuccess bool

const (
	UserUserUserDetailsResponseEnvelopeSuccessTrue UserUserUserDetailsResponseEnvelopeSuccess = true
)
