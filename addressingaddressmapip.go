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

// AddressingAddressMapIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingAddressMapIPService]
// method instead.
type AddressingAddressMapIPService struct {
	Options []option.RequestOption
}

// NewAddressingAddressMapIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingAddressMapIPService(opts ...option.RequestOption) (r *AddressingAddressMapIPService) {
	r = &AddressingAddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AddressingAddressMapIPService) Update(ctx context.Context, addressMapID string, ipAddress string, body AddressingAddressMapIPUpdateParams, opts ...option.RequestOption) (res *AddressingAddressMapIPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapIPUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", body.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an IP from a particular address map.
func (r *AddressingAddressMapIPService) Delete(ctx context.Context, addressMapID string, ipAddress string, body AddressingAddressMapIPDeleteParams, opts ...option.RequestOption) (res *AddressingAddressMapIPDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapIPDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", body.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressingAddressMapIPUpdateResponseUnknown],
// [AddressingAddressMapIPUpdateResponseArray] or [shared.UnionString].
type AddressingAddressMapIPUpdateResponse interface {
	ImplementsAddressingAddressMapIPUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapIPUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapIPUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapIPUpdateResponseArray []interface{}

func (r AddressingAddressMapIPUpdateResponseArray) ImplementsAddressingAddressMapIPUpdateResponse() {}

// Union satisfied by [AddressingAddressMapIPDeleteResponseUnknown],
// [AddressingAddressMapIPDeleteResponseArray] or [shared.UnionString].
type AddressingAddressMapIPDeleteResponse interface {
	ImplementsAddressingAddressMapIPDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapIPDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapIPDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapIPDeleteResponseArray []interface{}

func (r AddressingAddressMapIPDeleteResponseArray) ImplementsAddressingAddressMapIPDeleteResponse() {}

type AddressingAddressMapIPUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapIPUpdateResponseEnvelope struct {
	Errors   []AddressingAddressMapIPUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapIPUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapIPUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapIPUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapIPUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapIPUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapIPUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingAddressMapIPUpdateResponseEnvelope]
type addressingAddressMapIPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapIPUpdateResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingAddressMapIPUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapIPUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPUpdateResponseEnvelopeErrors]
type addressingAddressMapIPUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapIPUpdateResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingAddressMapIPUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapIPUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPUpdateResponseEnvelopeMessages]
type addressingAddressMapIPUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapIPUpdateResponseEnvelopeSuccess bool

const (
	AddressingAddressMapIPUpdateResponseEnvelopeSuccessTrue AddressingAddressMapIPUpdateResponseEnvelopeSuccess = true
)

type AddressingAddressMapIPUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       addressingAddressMapIPUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapIPUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPUpdateResponseEnvelopeResultInfo]
type addressingAddressMapIPUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapIPDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapIPDeleteResponseEnvelope struct {
	Errors   []AddressingAddressMapIPDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapIPDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapIPDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapIPDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapIPDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapIPDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapIPDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressingAddressMapIPDeleteResponseEnvelope]
type addressingAddressMapIPDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapIPDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressingAddressMapIPDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapIPDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPDeleteResponseEnvelopeErrors]
type addressingAddressMapIPDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapIPDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressingAddressMapIPDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapIPDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPDeleteResponseEnvelopeMessages]
type addressingAddressMapIPDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapIPDeleteResponseEnvelopeSuccess bool

const (
	AddressingAddressMapIPDeleteResponseEnvelopeSuccessTrue AddressingAddressMapIPDeleteResponseEnvelopeSuccess = true
)

type AddressingAddressMapIPDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       addressingAddressMapIPDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapIPDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressingAddressMapIPDeleteResponseEnvelopeResultInfo]
type addressingAddressMapIPDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapIPDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapIPDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
