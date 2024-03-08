// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// SeatService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSeatService] method instead.
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
func (r *SeatService) Edit(ctx context.Context, identifier string, body SeatEditParams, opts ...option.RequestOption) (res *[]SeatEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SeatEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SeatEditResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string               `json:"seat_uid"`
	UpdatedAt time.Time            `json:"updated_at" format:"date-time"`
	JSON      seatEditResponseJSON `json:"-"`
}

// seatEditResponseJSON contains the JSON metadata for the struct
// [SeatEditResponse]
type seatEditResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatEditResponseJSON) RawJSON() string {
	return r.raw
}

type SeatEditParams struct {
	Body param.Field[[]SeatEditParamsBody] `json:"body,required"`
}

func (r SeatEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SeatEditParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r SeatEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SeatEditResponseEnvelope struct {
	Errors   []SeatEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SeatEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []SeatEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    SeatEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo SeatEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       seatEditResponseEnvelopeJSON       `json:"-"`
}

// seatEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SeatEditResponseEnvelope]
type seatEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type SeatEditResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    seatEditResponseEnvelopeErrorsJSON `json:"-"`
}

// seatEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SeatEditResponseEnvelopeErrors]
type seatEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SeatEditResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    seatEditResponseEnvelopeMessagesJSON `json:"-"`
}

// seatEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SeatEditResponseEnvelopeMessages]
type seatEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SeatEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r seatEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SeatEditResponseEnvelopeSuccess bool

const (
	SeatEditResponseEnvelopeSuccessTrue SeatEditResponseEnvelopeSuccess = true
)

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
