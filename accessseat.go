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
func (r *AccessSeatService) ZeroTrustSeatsUpdateAUserSeat(ctx context.Context, identifier string, body AccessSeatZeroTrustSeatsUpdateAUserSeatParams, opts ...option.RequestOption) (res *AccessSeatZeroTrustSeatsUpdateAUserSeatResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponse struct {
	Errors     []AccessSeatZeroTrustSeatsUpdateAUserSeatResponseError    `json:"errors"`
	Messages   []AccessSeatZeroTrustSeatsUpdateAUserSeatResponseMessage  `json:"messages"`
	Result     []AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResult   `json:"result"`
	ResultInfo AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessSeatZeroTrustSeatsUpdateAUserSeatResponseSuccess `json:"success"`
	JSON    accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON    `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON contains the JSON metadata
// for the struct [AccessSeatZeroTrustSeatsUpdateAUserSeatResponse]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accessSeatZeroTrustSeatsUpdateAUserSeatResponseErrorJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseErrorJSON contains the JSON
// metadata for the struct [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseError]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accessSeatZeroTrustSeatsUpdateAUserSeatResponseMessageJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseMessageJSON contains the JSON
// metadata for the struct [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseMessage]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResult struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                                                    `json:"seat_uid"`
	UpdatedAt time.Time                                                 `json:"updated_at" format:"date-time"`
	JSON      accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultJSON contains the JSON
// metadata for the struct [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResult]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfoJSON `json:"-"`
}

// accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfo]
type accessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatZeroTrustSeatsUpdateAUserSeatResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessSeatZeroTrustSeatsUpdateAUserSeatResponseSuccess bool

const (
	AccessSeatZeroTrustSeatsUpdateAUserSeatResponseSuccessTrue AccessSeatZeroTrustSeatsUpdateAUserSeatResponseSuccess = true
)

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
