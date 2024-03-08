// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// ZeroTrustSeatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustSeatService] method
// instead.
type ZeroTrustSeatService struct {
	Options []option.RequestOption
}

// NewZeroTrustSeatService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustSeatService(opts ...option.RequestOption) (r *ZeroTrustSeatService) {
	r = &ZeroTrustSeatService{}
	r.Options = opts
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *ZeroTrustSeatService) Edit(ctx context.Context, identifier string, body ZeroTrustSeatEditParams, opts ...option.RequestOption) (res *[]ZeroTrustSeatEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustSeatEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/seats", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustSeatEditResponse struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// Identifier
	SeatUid   string                        `json:"seat_uid"`
	UpdatedAt time.Time                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustSeatEditResponseJSON `json:"-"`
}

// zeroTrustSeatEditResponseJSON contains the JSON metadata for the struct
// [ZeroTrustSeatEditResponse]
type zeroTrustSeatEditResponseJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustSeatEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustSeatEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustSeatEditParams struct {
	Body param.Field[[]ZeroTrustSeatEditParamsBody] `json:"body,required"`
}

func (r ZeroTrustSeatEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZeroTrustSeatEditParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
}

func (r ZeroTrustSeatEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustSeatEditResponseEnvelope struct {
	Errors   []ZeroTrustSeatEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustSeatEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustSeatEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustSeatEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustSeatEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustSeatEditResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustSeatEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZeroTrustSeatEditResponseEnvelope]
type zeroTrustSeatEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustSeatEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustSeatEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustSeatEditResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zeroTrustSeatEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustSeatEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustSeatEditResponseEnvelopeErrors]
type zeroTrustSeatEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustSeatEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustSeatEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustSeatEditResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zeroTrustSeatEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustSeatEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ZeroTrustSeatEditResponseEnvelopeMessages]
type zeroTrustSeatEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustSeatEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustSeatEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustSeatEditResponseEnvelopeSuccess bool

const (
	ZeroTrustSeatEditResponseEnvelopeSuccessTrue ZeroTrustSeatEditResponseEnvelopeSuccess = true
)

type ZeroTrustSeatEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       zeroTrustSeatEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustSeatEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [ZeroTrustSeatEditResponseEnvelopeResultInfo]
type zeroTrustSeatEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustSeatEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustSeatEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
