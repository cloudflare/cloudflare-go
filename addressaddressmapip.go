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

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AddressAddressMapIPService) Replace(ctx context.Context, accountID string, addressMapID string, ipAddress string, opts ...option.RequestOption) (res *AddressAddressMapIPReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapIPReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

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

// Union satisfied by [AddressAddressMapIPReplaceResponseUnknown],
// [AddressAddressMapIPReplaceResponseArray] or [shared.UnionString].
type AddressAddressMapIPReplaceResponse interface {
	ImplementsAddressAddressMapIPReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapIPReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapIPReplaceResponseArray []interface{}

func (r AddressAddressMapIPReplaceResponseArray) ImplementsAddressAddressMapIPReplaceResponse() {}

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

type AddressAddressMapIPReplaceResponseEnvelope struct {
	Errors   []AddressAddressMapIPReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapIPReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapIPReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapIPReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapIPReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapIPReplaceResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapIPReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressAddressMapIPReplaceResponseEnvelope]
type addressAddressMapIPReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPReplaceResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressAddressMapIPReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapIPReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressAddressMapIPReplaceResponseEnvelopeErrors]
type addressAddressMapIPReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapIPReplaceResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressAddressMapIPReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapIPReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressAddressMapIPReplaceResponseEnvelopeMessages]
type addressAddressMapIPReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapIPReplaceResponseEnvelopeSuccess bool

const (
	AddressAddressMapIPReplaceResponseEnvelopeSuccessTrue AddressAddressMapIPReplaceResponseEnvelopeSuccess = true
)

type AddressAddressMapIPReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       addressAddressMapIPReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapIPReplaceResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressAddressMapIPReplaceResponseEnvelopeResultInfo]
type addressAddressMapIPReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapIPReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
