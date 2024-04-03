// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AddressMapAccountService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressMapAccountService] method
// instead.
type AddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAddressMapAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressMapAccountService(opts ...option.RequestOption) (r *AddressMapAccountService) {
	r = &AddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AddressMapAccountService) Update(ctx context.Context, addressMapID string, params AddressMapAccountUpdateParams, opts ...option.RequestOption) (res *AddressMapAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapAccountUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", params.AccountID, addressMapID, params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an account as a member of a particular address map.
func (r *AddressMapAccountService) Delete(ctx context.Context, addressMapID string, params AddressMapAccountDeleteParams, opts ...option.RequestOption) (res *AddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapAccountDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", params.AccountID, addressMapID, params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [addressing.AddressMapAccountUpdateResponseUnknown],
// [addressing.AddressMapAccountUpdateResponseArray] or [shared.UnionString].
type AddressMapAccountUpdateResponse interface {
	ImplementsAddressingAddressMapAccountUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressMapAccountUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressMapAccountUpdateResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressMapAccountUpdateResponseArray []interface{}

func (r AddressMapAccountUpdateResponseArray) ImplementsAddressingAddressMapAccountUpdateResponse() {}

// Union satisfied by [addressing.AddressMapAccountDeleteResponseUnknown],
// [addressing.AddressMapAccountDeleteResponseArray] or [shared.UnionString].
type AddressMapAccountDeleteResponse interface {
	ImplementsAddressingAddressMapAccountDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressMapAccountDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AddressMapAccountDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressMapAccountDeleteResponseArray []interface{}

func (r AddressMapAccountDeleteResponseArray) ImplementsAddressingAddressMapAccountDeleteResponse() {}

type AddressMapAccountUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r AddressMapAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapAccountUpdateResponseEnvelope struct {
	Errors   []AddressMapAccountUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressMapAccountUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressMapAccountUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressMapAccountUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressMapAccountUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapAccountUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressMapAccountUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapAccountUpdateResponseEnvelope]
type addressMapAccountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressMapAccountUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressMapAccountUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressMapAccountUpdateResponseEnvelopeErrors]
type addressMapAccountUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressMapAccountUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressMapAccountUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressMapAccountUpdateResponseEnvelopeMessages]
type addressMapAccountUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapAccountUpdateResponseEnvelopeSuccess bool

const (
	AddressMapAccountUpdateResponseEnvelopeSuccessTrue AddressMapAccountUpdateResponseEnvelopeSuccess = true
)

func (r AddressMapAccountUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapAccountUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapAccountUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       addressMapAccountUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapAccountUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressMapAccountUpdateResponseEnvelopeResultInfo]
type addressMapAccountUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountDeleteParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r AddressMapAccountDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapAccountDeleteResponseEnvelope struct {
	Errors   []AddressMapAccountDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressMapAccountDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressMapAccountDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressMapAccountDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressMapAccountDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapAccountDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapAccountDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapAccountDeleteResponseEnvelope]
type addressMapAccountDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressMapAccountDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressMapAccountDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressMapAccountDeleteResponseEnvelopeErrors]
type addressMapAccountDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressMapAccountDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressMapAccountDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressMapAccountDeleteResponseEnvelopeMessages]
type addressMapAccountDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapAccountDeleteResponseEnvelopeSuccess bool

const (
	AddressMapAccountDeleteResponseEnvelopeSuccessTrue AddressMapAccountDeleteResponseEnvelopeSuccess = true
)

func (r AddressMapAccountDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapAccountDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapAccountDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       addressMapAccountDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapAccountDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressMapAccountDeleteResponseEnvelopeResultInfo]
type addressMapAccountDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
