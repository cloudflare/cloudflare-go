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

// AccessUserService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessUserService] method instead.
type AccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *AccessUserActiveSessionService
	LastSeenIdentity *AccessUserLastSeenIdentityService
	FailedLogins     *AccessUserFailedLoginService
}

// NewAccessUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessUserService(opts ...option.RequestOption) (r *AccessUserService) {
	r = &AccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccessUserService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccessUserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessUserListResponse struct {
	Errors     []AccessUserListResponseError    `json:"errors"`
	Messages   []AccessUserListResponseMessage  `json:"messages"`
	Result     []AccessUserListResponseResult   `json:"result"`
	ResultInfo AccessUserListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessUserListResponseSuccess `json:"success"`
	JSON    accessUserListResponseJSON    `json:"-"`
}

// accessUserListResponseJSON contains the JSON metadata for the struct
// [AccessUserListResponse]
type accessUserListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserListResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    accessUserListResponseErrorJSON `json:"-"`
}

// accessUserListResponseErrorJSON contains the JSON metadata for the struct
// [AccessUserListResponseError]
type accessUserListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserListResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    accessUserListResponseMessageJSON `json:"-"`
}

// accessUserListResponseMessageJSON contains the JSON metadata for the struct
// [AccessUserListResponseMessage]
type accessUserListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserListResponseResult struct {
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
	Uid       interface{}                      `json:"uid"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessUserListResponseResultJSON `json:"-"`
}

// accessUserListResponseResultJSON contains the JSON metadata for the struct
// [AccessUserListResponseResult]
type accessUserListResponseResultJSON struct {
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

func (r *AccessUserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessUserListResponseResultInfo struct {
	Count      interface{}                          `json:"count"`
	Page       interface{}                          `json:"page"`
	PerPage    interface{}                          `json:"per_page"`
	TotalCount interface{}                          `json:"total_count"`
	JSON       accessUserListResponseResultInfoJSON `json:"-"`
}

// accessUserListResponseResultInfoJSON contains the JSON metadata for the struct
// [AccessUserListResponseResultInfo]
type accessUserListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessUserListResponseSuccess bool

const (
	AccessUserListResponseSuccessTrue AccessUserListResponseSuccess = true
)
