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
func (r *AccessSeatService) Edit(ctx context.Context, identifier string, body AccessSeatEditParams, opts ...option.RequestOption) (res *[]AccessSeatEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessSeatEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessSeatEditResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                     `json:"seat_uid"`
	UpdatedAt time.Time                  `json:"updated_at" format:"date-time"`
	JSON      accessSeatEditResponseJSON `json:"-"`
}

// accessSeatEditResponseJSON contains the JSON metadata for the struct
// [AccessSeatEditResponse]
type accessSeatEditResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatEditParams struct {
	Body param.Field[[]AccessSeatEditParamsBody] `json:"body,required"`
}

func (r AccessSeatEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessSeatEditParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r AccessSeatEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessSeatEditResponseEnvelope struct {
	Errors   []AccessSeatEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessSeatEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessSeatEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessSeatEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessSeatEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessSeatEditResponseEnvelopeJSON       `json:"-"`
}

// accessSeatEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessSeatEditResponseEnvelope]
type accessSeatEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatEditResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessSeatEditResponseEnvelopeErrorsJSON `json:"-"`
}

// accessSeatEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessSeatEditResponseEnvelopeErrors]
type accessSeatEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessSeatEditResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessSeatEditResponseEnvelopeMessagesJSON `json:"-"`
}

// accessSeatEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessSeatEditResponseEnvelopeMessages]
type accessSeatEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessSeatEditResponseEnvelopeSuccess bool

const (
	AccessSeatEditResponseEnvelopeSuccessTrue AccessSeatEditResponseEnvelopeSuccess = true
)

type AccessSeatEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       accessSeatEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessSeatEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessSeatEditResponseEnvelopeResultInfo]
type accessSeatEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSeatEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
