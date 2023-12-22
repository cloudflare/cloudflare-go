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
	Options      []option.RequestOption
	FailedLogins *AccountAccessUserFailedLoginService
}

// NewAccountAccessUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessUserService(opts ...option.RequestOption) (r *AccountAccessUserService) {
	r = &AccountAccessUserService{}
	r.Options = opts
	r.FailedLogins = NewAccountAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccountAccessUserService) ZeroTrustUsersGetUsers(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *UsersResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UsersResponseCollection struct {
	Errors     []UsersResponseCollectionError    `json:"errors"`
	Messages   []UsersResponseCollectionMessage  `json:"messages"`
	Result     []UsersResponseCollectionResult   `json:"result"`
	ResultInfo UsersResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UsersResponseCollectionSuccess `json:"success"`
	JSON    usersResponseCollectionJSON    `json:"-"`
}

// usersResponseCollectionJSON contains the JSON metadata for the struct
// [UsersResponseCollection]
type usersResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsersResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsersResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    usersResponseCollectionErrorJSON `json:"-"`
}

// usersResponseCollectionErrorJSON contains the JSON metadata for the struct
// [UsersResponseCollectionError]
type usersResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsersResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsersResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    usersResponseCollectionMessageJSON `json:"-"`
}

// usersResponseCollectionMessageJSON contains the JSON metadata for the struct
// [UsersResponseCollectionMessage]
type usersResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsersResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsersResponseCollectionResult struct {
	// The ID of the user.
	ID interface{} `json:"id"`
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
	Uid       interface{}                       `json:"uid"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      usersResponseCollectionResultJSON `json:"-"`
}

// usersResponseCollectionResultJSON contains the JSON metadata for the struct
// [UsersResponseCollectionResult]
type usersResponseCollectionResultJSON struct {
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

func (r *UsersResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UsersResponseCollectionResultInfo struct {
	Count      interface{}                           `json:"count"`
	Page       interface{}                           `json:"page"`
	PerPage    interface{}                           `json:"per_page"`
	TotalCount interface{}                           `json:"total_count"`
	JSON       usersResponseCollectionResultInfoJSON `json:"-"`
}

// usersResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [UsersResponseCollectionResultInfo]
type usersResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UsersResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UsersResponseCollectionSuccess bool

const (
	UsersResponseCollectionSuccessTrue UsersResponseCollectionSuccess = true
)
