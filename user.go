// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewUserService] method instead.
type UserService struct {
	Options                []option.RequestOption
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
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// User Details
func (r *UserService) UserUserDetails(ctx context.Context, opts ...option.RequestOption) (res *UserUserUserDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserUserEditUserResponse struct {
	Errors   []UserUserEditUserResponseError   `json:"errors"`
	Messages []UserUserEditUserResponseMessage `json:"messages"`
	Result   interface{}                       `json:"result"`
	// Whether the API call was successful
	Success UserUserEditUserResponseSuccess `json:"success"`
	JSON    userUserEditUserResponseJSON    `json:"-"`
}

// userUserEditUserResponseJSON contains the JSON metadata for the struct
// [UserUserEditUserResponse]
type userUserEditUserResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserEditUserResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    userUserEditUserResponseErrorJSON `json:"-"`
}

// userUserEditUserResponseErrorJSON contains the JSON metadata for the struct
// [UserUserEditUserResponseError]
type userUserEditUserResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserEditUserResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    userUserEditUserResponseMessageJSON `json:"-"`
}

// userUserEditUserResponseMessageJSON contains the JSON metadata for the struct
// [UserUserEditUserResponseMessage]
type userUserEditUserResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserEditUserResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserUserEditUserResponseSuccess bool

const (
	UserUserEditUserResponseSuccessTrue UserUserEditUserResponseSuccess = true
)

type UserUserUserDetailsResponse struct {
	Errors   []UserUserUserDetailsResponseError   `json:"errors"`
	Messages []UserUserUserDetailsResponseMessage `json:"messages"`
	Result   interface{}                          `json:"result"`
	// Whether the API call was successful
	Success UserUserUserDetailsResponseSuccess `json:"success"`
	JSON    userUserUserDetailsResponseJSON    `json:"-"`
}

// userUserUserDetailsResponseJSON contains the JSON metadata for the struct
// [UserUserUserDetailsResponse]
type userUserUserDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserUserDetailsResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    userUserUserDetailsResponseErrorJSON `json:"-"`
}

// userUserUserDetailsResponseErrorJSON contains the JSON metadata for the struct
// [UserUserUserDetailsResponseError]
type userUserUserDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserUserUserDetailsResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    userUserUserDetailsResponseMessageJSON `json:"-"`
}

// userUserUserDetailsResponseMessageJSON contains the JSON metadata for the struct
// [UserUserUserDetailsResponseMessage]
type userUserUserDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserUserUserDetailsResponseSuccess bool

const (
	UserUserUserDetailsResponseSuccessTrue UserUserUserDetailsResponseSuccess = true
)

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
