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
func (r *AccessSeatService) Update(ctx context.Context, identifier string, body AccessSeatUpdateParams, opts ...option.RequestOption) (res *[]AccessSeatUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessSeatUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessSeatUpdateResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                       `json:"seat_uid"`
	UpdatedAt time.Time                    `json:"updated_at" format:"date-time"`
	JSON      accessSeatUpdateResponseJSON `json:"-"`
}

// accessSeatUpdateResponseJSON contains the JSON metadata for the struct
// [AccessSeatUpdateResponse]
type accessSeatUpdateResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatUpdateParams struct {
	Body param.Field[[]AccessSeatUpdateParamsBody] `json:"body,required"`
}

func (r AccessSeatUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessSeatUpdateParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r AccessSeatUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessSeatUpdateResponseEnvelope struct {
	Errors   []AccessSeatUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessSeatUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSeatUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessSeatUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessSeatUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessSeatUpdateResponseEnvelopeJSON       `json:"-"`
}

// accessSeatUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessSeatUpdateResponseEnvelope]
type accessSeatUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatUpdateResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessSeatUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessSeatUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessSeatUpdateResponseEnvelopeErrors]
type accessSeatUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatUpdateResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accessSeatUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessSeatUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessSeatUpdateResponseEnvelopeMessages]
type accessSeatUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessSeatUpdateResponseEnvelopeSuccess bool

const (
	AccessSeatUpdateResponseEnvelopeSuccessTrue AccessSeatUpdateResponseEnvelopeSuccess = true
)

type AccessSeatUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       accessSeatUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessSeatUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AccessSeatUpdateResponseEnvelopeResultInfo]
type accessSeatUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
