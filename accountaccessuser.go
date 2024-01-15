// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAccessUserService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessUserService] method
// instead.
type AccountAccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *AccountAccessUserActiveSessionService
	LastSeenIdentity *AccountAccessUserLastSeenIdentityService
	FailedLogins     *AccountAccessUserFailedLoginService
}

// NewAccountAccessUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessUserService(opts ...option.RequestOption) (r *AccountAccessUserService) {
	r = &AccountAccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccountAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewAccountAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewAccountAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccountAccessUserService) ZeroTrustUsersGetUsers(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessUserZeroTrustUsersGetUsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessUserZeroTrustUsersGetUsersResponse struct {
	Errors     []AccountAccessUserZeroTrustUsersGetUsersResponseError    `json:"errors"`
	Messages   []AccountAccessUserZeroTrustUsersGetUsersResponseMessage  `json:"messages"`
	Result     []AccountAccessUserZeroTrustUsersGetUsersResponseResult   `json:"result"`
	ResultInfo AccountAccessUserZeroTrustUsersGetUsersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessUserZeroTrustUsersGetUsersResponseSuccess `json:"success"`
	JSON    accountAccessUserZeroTrustUsersGetUsersResponseJSON    `json:"-"`
}

// accountAccessUserZeroTrustUsersGetUsersResponseJSON contains the JSON metadata
// for the struct [AccountAccessUserZeroTrustUsersGetUsersResponse]
type accountAccessUserZeroTrustUsersGetUsersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserZeroTrustUsersGetUsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserZeroTrustUsersGetUsersResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAccessUserZeroTrustUsersGetUsersResponseErrorJSON `json:"-"`
}

// accountAccessUserZeroTrustUsersGetUsersResponseErrorJSON contains the JSON
// metadata for the struct [AccountAccessUserZeroTrustUsersGetUsersResponseError]
type accountAccessUserZeroTrustUsersGetUsersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserZeroTrustUsersGetUsersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserZeroTrustUsersGetUsersResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountAccessUserZeroTrustUsersGetUsersResponseMessageJSON `json:"-"`
}

// accountAccessUserZeroTrustUsersGetUsersResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessUserZeroTrustUsersGetUsersResponseMessage]
type accountAccessUserZeroTrustUsersGetUsersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserZeroTrustUsersGetUsersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserZeroTrustUsersGetUsersResponseResult struct {
	// UUID
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid interface{} `json:"seat_uid"`
	// The unique API identifier for the user.
	Uid       interface{}                                               `json:"uid"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      accountAccessUserZeroTrustUsersGetUsersResponseResultJSON `json:"-"`
}

// accountAccessUserZeroTrustUsersGetUsersResponseResultJSON contains the JSON
// metadata for the struct [AccountAccessUserZeroTrustUsersGetUsersResponseResult]
type accountAccessUserZeroTrustUsersGetUsersResponseResultJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUid             apijson.Field
	Uid                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountAccessUserZeroTrustUsersGetUsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessUserZeroTrustUsersGetUsersResponseResultInfo struct {
	Count      interface{}                                                   `json:"count"`
	Page       interface{}                                                   `json:"page"`
	PerPage    interface{}                                                   `json:"per_page"`
	TotalCount interface{}                                                   `json:"total_count"`
	JSON       accountAccessUserZeroTrustUsersGetUsersResponseResultInfoJSON `json:"-"`
}

// accountAccessUserZeroTrustUsersGetUsersResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAccessUserZeroTrustUsersGetUsersResponseResultInfo]
type accountAccessUserZeroTrustUsersGetUsersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserZeroTrustUsersGetUsersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessUserZeroTrustUsersGetUsersResponseSuccess bool

const (
	AccountAccessUserZeroTrustUsersGetUsersResponseSuccessTrue AccountAccessUserZeroTrustUsersGetUsersResponseSuccess = true
)
