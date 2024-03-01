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

// ZeroTrustAccessSeatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustAccessSeatService]
// method instead.
type ZeroTrustAccessSeatService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessSeatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustAccessSeatService(opts ...option.RequestOption) (r *ZeroTrustAccessSeatService) {
	r = &ZeroTrustAccessSeatService{}
	r.Options = opts
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *ZeroTrustAccessSeatService) Edit(ctx context.Context, identifier string, body ZeroTrustAccessSeatEditParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessSeatEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessSeatEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessSeatEditResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                              `json:"seat_uid"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessSeatEditResponseJSON `json:"-"`
}

// zeroTrustAccessSeatEditResponseJSON contains the JSON metadata for the struct
// [ZeroTrustAccessSeatEditResponse]
type zeroTrustAccessSeatEditResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessSeatEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessSeatEditParams struct {
	Body param.Field[[]ZeroTrustAccessSeatEditParamsBody] `json:"body,required"`
}

func (r ZeroTrustAccessSeatEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustAccessSeatEditParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r ZeroTrustAccessSeatEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessSeatEditResponseEnvelope struct {
	Errors   []ZeroTrustAccessSeatEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessSeatEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessSeatEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessSeatEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessSeatEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessSeatEditResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessSeatEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustAccessSeatEditResponseEnvelope]
type zeroTrustAccessSeatEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessSeatEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessSeatEditResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustAccessSeatEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessSeatEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustAccessSeatEditResponseEnvelopeErrors]
type zeroTrustAccessSeatEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessSeatEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessSeatEditResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustAccessSeatEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessSeatEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustAccessSeatEditResponseEnvelopeMessages]
type zeroTrustAccessSeatEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessSeatEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessSeatEditResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessSeatEditResponseEnvelopeSuccessTrue ZeroTrustAccessSeatEditResponseEnvelopeSuccess = true
)

type ZeroTrustAccessSeatEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       zeroTrustAccessSeatEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessSeatEditResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustAccessSeatEditResponseEnvelopeResultInfo]
type zeroTrustAccessSeatEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessSeatEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
