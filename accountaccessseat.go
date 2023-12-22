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

// AccountAccessSeatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAccessSeatService] method
// instead.
type AccountAccessSeatService struct {
	Options []option.RequestOption
}

// NewAccountAccessSeatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessSeatService(opts ...option.RequestOption) (r *AccountAccessSeatService) {
	r = &AccountAccessSeatService{}
	r.Options = opts
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *AccountAccessSeatService) ZeroTrustSeatsUpdateAUserSeat(ctx context.Context, identifier interface{}, body AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParams, opts ...option.RequestOption) (res *SeatsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SeatsResponseCollection struct {
	Errors     []SeatsResponseCollectionError    `json:"errors"`
	Messages   []SeatsResponseCollectionMessage  `json:"messages"`
	Result     []SeatsResponseCollectionResult   `json:"result"`
	ResultInfo SeatsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SeatsResponseCollectionSuccess `json:"success"`
	JSON    seatsResponseCollectionJSON    `json:"-"`
}

// seatsResponseCollectionJSON contains the JSON metadata for the struct
// [SeatsResponseCollection]
type seatsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SeatsResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    seatsResponseCollectionErrorJSON `json:"-"`
}

// seatsResponseCollectionErrorJSON contains the JSON metadata for the struct
// [SeatsResponseCollectionError]
type seatsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SeatsResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    seatsResponseCollectionMessageJSON `json:"-"`
}

// seatsResponseCollectionMessageJSON contains the JSON metadata for the struct
// [SeatsResponseCollectionMessage]
type seatsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SeatsResponseCollectionResult struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid   interface{}                       `json:"seat_uid"`
	UpdatedAt time.Time                         `json:"updated_at" format:"date-time"`
	JSON      seatsResponseCollectionResultJSON `json:"-"`
}

// seatsResponseCollectionResultJSON contains the JSON metadata for the struct
// [SeatsResponseCollectionResult]
type seatsResponseCollectionResultJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SeatsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       seatsResponseCollectionResultInfoJSON `json:"-"`
}

// seatsResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [SeatsResponseCollectionResultInfo]
type seatsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SeatsResponseCollectionSuccess bool

const (
	SeatsResponseCollectionSuccessTrue SeatsResponseCollectionSuccess = true
)

type AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParams struct {
	Body param.Field[[]AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody] `json:"body,required"`
}

func (r AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid param.Field[interface{}] `json:"seat_uid,required"`
}

func (r AccountAccessSeatZeroTrustSeatsUpdateAUserSeatParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
