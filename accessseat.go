// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessSeatService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessSeatService] method instead.
type AccessSeatService struct {
	Options []option.RequestOption
}

// NewAccessSeatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessSeatService(opts ...option.RequestOption) (r *AccessSeatService) {
	r = &AccessSeatService{}
	r.Options = opts
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *AccessSeatService) ZeroTrustSeatsUpdateAUserSeat(ctx context.Context, identifier string, body AccessSeatZeroTrustSeatsUpdateAUserSeatParams, opts ...option.RequestOption) (res *[]AccessSeatZeroTrustSeatsUpdateAUserSeatResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                                              `json:"seat_uid"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON contains the JSON metadata
// for the struct [AccessSeatZeroTrustSeatsUpdateAUserSeatResponse]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatParams struct {
	Body param.Field[[]AccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody] `json:"body,required"`
}

func (r AccessSeatZeroTrustSeatsUpdateAUserSeatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r AccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelope struct {
	Errors   []AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSeatZeroTrustSeatsUpdateAUserSeatResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeJSON       `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelope]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrors struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrorsJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrors]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessages struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessagesJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessages]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeSuccess bool

const (
	AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeSuccessTrue AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeSuccess = true
)

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfo]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
