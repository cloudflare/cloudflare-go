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
func (r *UserService) UserEditUser(ctx context.Context, body UserUserEditUserParams, opts ...option.RequestOption) (res *SingleUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// User Details
func (r *UserService) UserUserDetails(ctx context.Context, opts ...option.RequestOption) (res *SingleUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SingleUserResponse struct {
	Errors   []SingleUserResponseError   `json:"errors"`
	Messages []SingleUserResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success SingleUserResponseSuccess `json:"success"`
	JSON    singleUserResponseJSON    `json:"-"`
}

// singleUserResponseJSON contains the JSON metadata for the struct
// [SingleUserResponse]
type singleUserResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleUserResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    singleUserResponseErrorJSON `json:"-"`
}

// singleUserResponseErrorJSON contains the JSON metadata for the struct
// [SingleUserResponseError]
type singleUserResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleUserResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleUserResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    singleUserResponseMessageJSON `json:"-"`
}

// singleUserResponseMessageJSON contains the JSON metadata for the struct
// [SingleUserResponseMessage]
type singleUserResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleUserResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleUserResponseSuccess bool

const (
	SingleUserResponseSuccessTrue SingleUserResponseSuccess = true
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
