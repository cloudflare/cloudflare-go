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

// AddressingAddressMapAccountService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressingAddressMapAccountService] method instead.
type AddressingAddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAddressingAddressMapAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingAddressMapAccountService(opts ...option.RequestOption) (r *AddressingAddressMapAccountService) {
	r = &AddressingAddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AddressingAddressMapAccountService) Update(ctx context.Context, addressMapID string, body AddressingAddressMapAccountUpdateParams, opts ...option.RequestOption) (res *AddressingAddressMapAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapAccountUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", body.AccountID, addressMapID, body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an account as a member of a particular address map.
func (r *AddressingAddressMapAccountService) Delete(ctx context.Context, addressMapID string, body AddressingAddressMapAccountDeleteParams, opts ...option.RequestOption) (res *AddressingAddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingAddressMapAccountDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", body.AccountID, addressMapID, body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressingAddressMapAccountUpdateResponseUnknown],
// [AddressingAddressMapAccountUpdateResponseArray] or [shared.UnionString].
type AddressingAddressMapAccountUpdateResponse interface {
	ImplementsAddressingAddressMapAccountUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapAccountUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapAccountUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapAccountUpdateResponseArray []interface{}

func (r AddressingAddressMapAccountUpdateResponseArray) ImplementsAddressingAddressMapAccountUpdateResponse() {
}

// Union satisfied by [AddressingAddressMapAccountDeleteResponseUnknown],
// [AddressingAddressMapAccountDeleteResponseArray] or [shared.UnionString].
type AddressingAddressMapAccountDeleteResponse interface {
	ImplementsAddressingAddressMapAccountDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressingAddressMapAccountDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressingAddressMapAccountDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressingAddressMapAccountDeleteResponseArray []interface{}

func (r AddressingAddressMapAccountDeleteResponseArray) ImplementsAddressingAddressMapAccountDeleteResponse() {
}

type AddressingAddressMapAccountUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapAccountUpdateResponseEnvelope struct {
	Errors   []AddressingAddressMapAccountUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapAccountUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapAccountUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapAccountUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapAccountUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapAccountUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapAccountUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingAddressMapAccountUpdateResponseEnvelope]
type addressingAddressMapAccountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapAccountUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    addressingAddressMapAccountUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapAccountUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapAccountUpdateResponseEnvelopeErrors]
type addressingAddressMapAccountUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapAccountUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    addressingAddressMapAccountUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapAccountUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapAccountUpdateResponseEnvelopeMessages]
type addressingAddressMapAccountUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapAccountUpdateResponseEnvelopeSuccess bool

const (
	AddressingAddressMapAccountUpdateResponseEnvelopeSuccessTrue AddressingAddressMapAccountUpdateResponseEnvelopeSuccess = true
)

type AddressingAddressMapAccountUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       addressingAddressMapAccountUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapAccountUpdateResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [AddressingAddressMapAccountUpdateResponseEnvelopeResultInfo]
type addressingAddressMapAccountUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapAccountDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressingAddressMapAccountDeleteResponseEnvelope struct {
	Errors   []AddressingAddressMapAccountDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingAddressMapAccountDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingAddressMapAccountDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressingAddressMapAccountDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressingAddressMapAccountDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressingAddressMapAccountDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressingAddressMapAccountDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressingAddressMapAccountDeleteResponseEnvelope]
type addressingAddressMapAccountDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapAccountDeleteResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    addressingAddressMapAccountDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingAddressMapAccountDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapAccountDeleteResponseEnvelopeErrors]
type addressingAddressMapAccountDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressingAddressMapAccountDeleteResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    addressingAddressMapAccountDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingAddressMapAccountDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AddressingAddressMapAccountDeleteResponseEnvelopeMessages]
type addressingAddressMapAccountDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressingAddressMapAccountDeleteResponseEnvelopeSuccess bool

const (
	AddressingAddressMapAccountDeleteResponseEnvelopeSuccessTrue AddressingAddressMapAccountDeleteResponseEnvelopeSuccess = true
)

type AddressingAddressMapAccountDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       addressingAddressMapAccountDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressingAddressMapAccountDeleteResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [AddressingAddressMapAccountDeleteResponseEnvelopeResultInfo]
type addressingAddressMapAccountDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingAddressMapAccountDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressingAddressMapAccountDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
