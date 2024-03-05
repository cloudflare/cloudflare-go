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

// ZeroTrustAccessUserService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessUserService]
// method instead.
type ZeroTrustAccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *ZeroTrustAccessUserActiveSessionService
	LastSeenIdentity *ZeroTrustAccessUserLastSeenIdentityService
	FailedLogins     *ZeroTrustAccessUserFailedLoginService
}

// NewZeroTrustAccessUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessUserService(opts ...option.RequestOption) (r *ZeroTrustAccessUserService) {
	r = &ZeroTrustAccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewZeroTrustAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewZeroTrustAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewZeroTrustAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *ZeroTrustAccessUserService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessUsers, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessUserListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUsers struct {
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
	Uid       interface{}     `json:"uid"`
	UpdatedAt time.Time       `json:"updated_at" format:"date-time"`
	JSON      accessUsersJSON `json:"-"`
}

// accessUsersJSON contains the JSON metadata for the struct [AccessUsers]
type accessUsersJSON struct {
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

func (r *AccessUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserListResponseEnvelope struct {
	Errors   []ZeroTrustAccessUserListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessUserListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessUsers                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessUserListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessUserListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessUserListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessUserListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessUserListResponseEnvelope]
type zeroTrustAccessUserListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessUserListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessUserListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessUserListResponseEnvelopeErrors]
type zeroTrustAccessUserListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessUserListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustAccessUserListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessUserListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserListResponseEnvelopeMessages]
type zeroTrustAccessUserListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessUserListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessUserListResponseEnvelopeSuccessTrue ZeroTrustAccessUserListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessUserListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       zeroTrustAccessUserListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessUserListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustAccessUserListResponseEnvelopeResultInfo]
type zeroTrustAccessUserListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessUserListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
