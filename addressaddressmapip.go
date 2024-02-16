// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AddressAddressMapIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressAddressMapIPService]
// method instead.
type AddressAddressMapIPService struct {
	Options []option.RequestOption
}

// NewAddressAddressMapIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressAddressMapIPService(opts ...option.RequestOption) (r *AddressAddressMapIPService) {
	r = &AddressAddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AddressAddressMapIPService) Update(ctx context.Context, accountID string, addressMapID string, ipAddress string, opts ...option.RequestOption) (res *AddressAddressMapIPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapIPUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an IP from a particular address map.
func (r *AddressAddressMapIPService) Delete(ctx context.Context, accountID string, addressMapID string, ipAddress string, opts ...option.RequestOption) (res *AddressAddressMapIPDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapIPDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressAddressMapIPUpdateResponseUnknown],
// [AddressAddressMapIPUpdateResponseArray] or [shared.UnionString].
type AddressAddressMapIPUpdateResponse interface {
	ImplementsAddressAddressMapIPUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapIPUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapIPUpdateResponseArray []interface{}

func (r AddressAddressMapIPUpdateResponseArray) ImplementsAddressAddressMapIPUpdateResponse() {}

// Union satisfied by [AddressAddressMapIPDeleteResponseUnknown],
// [AddressAddressMapIPDeleteResponseArray] or [shared.UnionString].
type AddressAddressMapIPDeleteResponse interface {
	ImplementsAddressAddressMapIPDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapIPDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapIPDeleteResponseArray []interface{}

func (r AddressAddressMapIPDeleteResponseArray) ImplementsAddressAddressMapIPDeleteResponse() {}

type AddressAddressMapIPUpdateResponseEnvelope struct {
	Errors   []AddressAddressMapIPUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapIPUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapIPUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapIPUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapIPUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapIPUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapIPUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapIPUpdateResponseEnvelope]
type addressAddressMapIPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressAddressMapIPUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapIPUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressAddressMapIPUpdateResponseEnvelopeErrors]
type addressAddressMapIPUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressAddressMapIPUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapIPUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressAddressMapIPUpdateResponseEnvelopeMessages]
type addressAddressMapIPUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapIPUpdateResponseEnvelopeSuccess bool

const (
	AddressAddressMapIPUpdateResponseEnvelopeSuccessTrue AddressAddressMapIPUpdateResponseEnvelopeSuccess = true
)

type AddressAddressMapIPUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       addressAddressMapIPUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapIPUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressAddressMapIPUpdateResponseEnvelopeResultInfo]
type addressAddressMapIPUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPDeleteResponseEnvelope struct {
	Errors   []AddressAddressMapIPDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapIPDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapIPDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapIPDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapIPDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapIPDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapIPDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressAddressMapIPDeleteResponseEnvelope]
type addressAddressMapIPDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    addressAddressMapIPDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapIPDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressAddressMapIPDeleteResponseEnvelopeErrors]
type addressAddressMapIPDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressAddressMapIPDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapIPDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressAddressMapIPDeleteResponseEnvelopeMessages]
type addressAddressMapIPDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapIPDeleteResponseEnvelopeSuccess bool

const (
	AddressAddressMapIPDeleteResponseEnvelopeSuccessTrue AddressAddressMapIPDeleteResponseEnvelopeSuccess = true
)

type AddressAddressMapIPDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       addressAddressMapIPDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapIPDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressAddressMapIPDeleteResponseEnvelopeResultInfo]
type addressAddressMapIPDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
