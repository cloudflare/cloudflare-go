// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// AddressingAddressMapZoneService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingAddressMapZoneService] method instead.
type AddressingAddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAddressingAddressMapZoneService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingAddressMapZoneService(opts ...option.RequestOption) (r *AddressingAddressMapZoneService) {
	r = &AddressingAddressMapZoneService{}
	r.Options = opts
	return
}

// Add a zone as a member of a particular address map.
func (r *AddressingAddressMapZoneService) Update(ctx context.Context, addressMapID string, body AddressingAddressMapZoneUpdateParams, opts ...option.RequestOption) (res *AddressingAddressMapZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapZoneUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountOrZone, addressMapID, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a zone as a member of a particular address map.
func (r *AddressingAddressMapZoneService) Delete(ctx context.Context, addressMapID string, body AddressingAddressMapZoneDeleteParams, opts ...option.RequestOption) (res *AddressingAddressMapZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapZoneDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountOrZone, addressMapID, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressingAddressMapZoneUpdateResponseUnknown],
// [AddressingAddressMapZoneUpdateResponseArray] or [shared.UnionString].
type AddressingAddressMapZoneUpdateResponse interface {
	ImplementsAddressingAddressMapZoneUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapZoneUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapZoneUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapZoneUpdateResponseArray []interface{}

func (r AddressingAddressMapZoneUpdateResponseArray) ImplementsAddressingAddressMapZoneUpdateResponse() {
}

// Union satisfied by [AddressingAddressMapZoneDeleteResponseUnknown],
// [AddressingAddressMapZoneDeleteResponseArray] or [shared.UnionString].
type AddressingAddressMapZoneDeleteResponse interface {
	ImplementsAddressingAddressMapZoneDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapZoneDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapZoneDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapZoneDeleteResponseArray []interface{}

func (r AddressingAddressMapZoneDeleteResponseArray) ImplementsAddressingAddressMapZoneDeleteResponse() {
}

type AddressingAddressMapZoneUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapZoneUpdateResponseEnvelope struct {
	Errors   []AddressingAddressMapZoneUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapZoneUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapZoneUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapZoneUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapZoneUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapZoneUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapZoneUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingAddressMapZoneUpdateResponseEnvelope]
type addressingAddressMapZoneUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapZoneUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingAddressMapZoneUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapZoneUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingAddressMapZoneUpdateResponseEnvelopeErrors]
type addressingAddressMapZoneUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapZoneUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingAddressMapZoneUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapZoneUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingAddressMapZoneUpdateResponseEnvelopeMessages]
type addressingAddressMapZoneUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapZoneUpdateResponseEnvelopeSuccess bool

const (
	AddressingAddressMapZoneUpdateResponseEnvelopeSuccessTrue AddressingAddressMapZoneUpdateResponseEnvelopeSuccess = true
)

type AddressingAddressMapZoneUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       addressingAddressMapZoneUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapZoneUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapZoneUpdateResponseEnvelopeResultInfo]
type addressingAddressMapZoneUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapZoneDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapZoneDeleteResponseEnvelope struct {
	Errors   []AddressingAddressMapZoneDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapZoneDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapZoneDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapZoneDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapZoneDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapZoneDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapZoneDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingAddressMapZoneDeleteResponseEnvelope]
type addressingAddressMapZoneDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapZoneDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingAddressMapZoneDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapZoneDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingAddressMapZoneDeleteResponseEnvelopeErrors]
type addressingAddressMapZoneDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapZoneDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressingAddressMapZoneDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapZoneDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingAddressMapZoneDeleteResponseEnvelopeMessages]
type addressingAddressMapZoneDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapZoneDeleteResponseEnvelopeSuccess bool

const (
	AddressingAddressMapZoneDeleteResponseEnvelopeSuccessTrue AddressingAddressMapZoneDeleteResponseEnvelopeSuccess = true
)

type AddressingAddressMapZoneDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       addressingAddressMapZoneDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapZoneDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapZoneDeleteResponseEnvelopeResultInfo]
type addressingAddressMapZoneDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapZoneDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapZoneDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
