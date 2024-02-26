// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AddressAddressMapAccountService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressAddressMapAccountService] method instead.
type AddressAddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAddressAddressMapAccountService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressAddressMapAccountService(opts ...option.RequestOption) (r *AddressAddressMapAccountService) {
	r = &AddressAddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AddressAddressMapAccountService) Update(ctx context.Context, addressMapID string, body AddressAddressMapAccountUpdateParams, opts ...option.RequestOption) (res *AddressAddressMapAccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapAccountUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/:account_id", body.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an account as a member of a particular address map.
func (r *AddressAddressMapAccountService) Delete(ctx context.Context, addressMapID string, body AddressAddressMapAccountDeleteParams, opts ...option.RequestOption) (res *AddressAddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapAccountDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/:account_id", body.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressAddressMapAccountUpdateResponseUnknown],
// [AddressAddressMapAccountUpdateResponseArray] or [shared.UnionString].
type AddressAddressMapAccountUpdateResponse interface {
	ImplementsAddressAddressMapAccountUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapAccountUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapAccountUpdateResponseArray []interface{}

func (r AddressAddressMapAccountUpdateResponseArray) ImplementsAddressAddressMapAccountUpdateResponse() {
}

// Union satisfied by [AddressAddressMapAccountDeleteResponseUnknown],
// [AddressAddressMapAccountDeleteResponseArray] or [shared.UnionString].
type AddressAddressMapAccountDeleteResponse interface {
	ImplementsAddressAddressMapAccountDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapAccountDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapAccountDeleteResponseArray []interface{}

func (r AddressAddressMapAccountDeleteResponseArray) ImplementsAddressAddressMapAccountDeleteResponse() {
}

type AddressAddressMapAccountUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressAddressMapAccountUpdateResponseEnvelope struct {
	Errors   []AddressAddressMapAccountUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapAccountUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapAccountUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapAccountUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapAccountUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapAccountUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapAccountUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressAddressMapAccountUpdateResponseEnvelope]
type addressAddressMapAccountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapAccountUpdateResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressAddressMapAccountUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapAccountUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressAddressMapAccountUpdateResponseEnvelopeErrors]
type addressAddressMapAccountUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapAccountUpdateResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressAddressMapAccountUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapAccountUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressAddressMapAccountUpdateResponseEnvelopeMessages]
type addressAddressMapAccountUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapAccountUpdateResponseEnvelopeSuccess bool

const (
	AddressAddressMapAccountUpdateResponseEnvelopeSuccessTrue AddressAddressMapAccountUpdateResponseEnvelopeSuccess = true
)

type AddressAddressMapAccountUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       addressAddressMapAccountUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapAccountUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressAddressMapAccountUpdateResponseEnvelopeResultInfo]
type addressAddressMapAccountUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapAccountDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressAddressMapAccountDeleteResponseEnvelope struct {
	Errors   []AddressAddressMapAccountDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapAccountDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapAccountDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapAccountDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapAccountDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapAccountDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapAccountDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [AddressAddressMapAccountDeleteResponseEnvelope]
type addressAddressMapAccountDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapAccountDeleteResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressAddressMapAccountDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapAccountDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressAddressMapAccountDeleteResponseEnvelopeErrors]
type addressAddressMapAccountDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapAccountDeleteResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    addressAddressMapAccountDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapAccountDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressAddressMapAccountDeleteResponseEnvelopeMessages]
type addressAddressMapAccountDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapAccountDeleteResponseEnvelopeSuccess bool

const (
	AddressAddressMapAccountDeleteResponseEnvelopeSuccessTrue AddressAddressMapAccountDeleteResponseEnvelopeSuccess = true
)

type AddressAddressMapAccountDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       addressAddressMapAccountDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapAccountDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AddressAddressMapAccountDeleteResponseEnvelopeResultInfo]
type addressAddressMapAccountDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapAccountDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
