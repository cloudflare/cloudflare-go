// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SeatService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSeatService] method instead.
type SeatService struct {
	Options []option.RequestOption
}

// NewSeatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSeatService(opts ...option.RequestOption) (r *SeatService) {
	r = &SeatService{}
	r.Options = opts
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *SeatService) Edit(ctx context.Context, params SeatEditParams, opts ...option.RequestOption) (res *[]Seat, err error) {
	var env SeatEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/seats", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Seat struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID   string    `json:"seat_uid"`
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	JSON      seatJSON  `json:"-"`
}

// seatJSON contains the JSON metadata for the struct [Seat]
type seatJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUID     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Seat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatJSON) RawJSON() string {
	return r.raw
}

type SeatEditParams struct {
	// Identifier
	AccountID param.Field[string]  `path:"account_id,required"`
	Body      []SeatEditParamsBody `json:"body,required"`
}

func (r SeatEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SeatEditParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID param.Field[string] `json:"seat_uid,required"`
}

func (r SeatEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SeatEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    SeatEditResponseEnvelopeSuccess    `json:"success,required"`
	Result     []Seat                             `json:"result"`
	ResultInfo SeatEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       seatEditResponseEnvelopeJSON       `json:"-"`
}

// seatEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SeatEditResponseEnvelope]
type seatEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SeatEditResponseEnvelopeSuccess bool

const (
	SeatEditResponseEnvelopeSuccessTrue SeatEditResponseEnvelopeSuccess = true
)

func (r SeatEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SeatEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SeatEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       seatEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// seatEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [SeatEditResponseEnvelopeResultInfo]
type seatEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
