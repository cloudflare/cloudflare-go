// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *AccessUserService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessUserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessUserListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/users", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserListResponse struct {
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
	Uid       interface{}                `json:"uid"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessUserListResponseJSON `json:"-"`
}

// accessUserListResponseJSON contains the JSON metadata for the struct
// [AccessUserListResponse]
type accessUserListResponseJSON struct {
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

func (r *AccessUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserListResponseEnvelope struct {
	Errors   []AccessUserListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessUserListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessUserListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessUserListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessUserListResponseEnvelopeJSON       `json:"-"`
}

// accessUserListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessUserListResponseEnvelope]
type accessUserListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessUserListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessUserListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessUserListResponseEnvelopeErrors]
type accessUserListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessUserListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessUserListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessUserListResponseEnvelopeMessages]
type accessUserListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessUserListResponseEnvelopeSuccess bool

const (
	AccessUserListResponseEnvelopeSuccessTrue AccessUserListResponseEnvelopeSuccess = true
)

type AccessUserListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       accessUserListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessUserListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessUserListResponseEnvelopeResultInfo]
type accessUserListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
