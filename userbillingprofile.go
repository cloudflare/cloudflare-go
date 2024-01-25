// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserBillingProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserBillingProfileService] method
// instead.
type UserBillingProfileService struct {
	Options []option.RequestOption
}

// NewUserBillingProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserBillingProfileService(opts ...option.RequestOption) (r *UserBillingProfileService) {
	r = &UserBillingProfileService{}
	r.Options = opts
	return
}

// Accesses your billing profile object.
func (r *UserBillingProfileService) UserBillingProfileBillingProfileDetails(ctx context.Context, opts ...option.RequestOption) (res *UserBillingProfileUserBillingProfileBillingProfileDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponse struct {
	Errors   []UserBillingProfileUserBillingProfileBillingProfileDetailsResponseError   `json:"errors"`
	Messages []UserBillingProfileUserBillingProfileBillingProfileDetailsResponseMessage `json:"messages"`
	Result   interface{}                                                                `json:"result"`
	// Whether the API call was successful
	Success UserBillingProfileUserBillingProfileBillingProfileDetailsResponseSuccess `json:"success"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseJSON    `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseJSON contains
// the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponse]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseErrorJSON `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseError]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    userBillingProfileUserBillingProfileBillingProfileDetailsResponseMessageJSON `json:"-"`
}

// userBillingProfileUserBillingProfileBillingProfileDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [UserBillingProfileUserBillingProfileBillingProfileDetailsResponseMessage]
type userBillingProfileUserBillingProfileBillingProfileDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingProfileUserBillingProfileBillingProfileDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserBillingProfileUserBillingProfileBillingProfileDetailsResponseSuccess bool

const (
	UserBillingProfileUserBillingProfileBillingProfileDetailsResponseSuccessTrue UserBillingProfileUserBillingProfileBillingProfileDetailsResponseSuccess = true
)
