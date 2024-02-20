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

// AddressAddressMapZoneService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressAddressMapZoneService]
// method instead.
type AddressAddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAddressAddressMapZoneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressAddressMapZoneService(opts ...option.RequestOption) (r *AddressAddressMapZoneService) {
	r = &AddressAddressMapZoneService{}
	r.Options = opts
	return
}

// Remove a zone as a member of a particular address map.
func (r *AddressAddressMapZoneService) Delete(ctx context.Context, accountID string, addressMapID string, zoneID string, opts ...option.RequestOption) (res *AddressAddressMapZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapZoneDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountID, addressMapID, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Add a zone as a member of a particular address map.
func (r *AddressAddressMapZoneService) Replace(ctx context.Context, accountID string, addressMapID string, zoneID string, opts ...option.RequestOption) (res *AddressAddressMapZoneReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressAddressMapZoneReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountID, addressMapID, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AddressAddressMapZoneDeleteResponseUnknown],
// [AddressAddressMapZoneDeleteResponseArray] or [shared.UnionString].
type AddressAddressMapZoneDeleteResponse interface {
	ImplementsAddressAddressMapZoneDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapZoneDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapZoneDeleteResponseArray []interface{}

func (r AddressAddressMapZoneDeleteResponseArray) ImplementsAddressAddressMapZoneDeleteResponse() {}

// Union satisfied by [AddressAddressMapZoneReplaceResponseUnknown],
// [AddressAddressMapZoneReplaceResponseArray] or [shared.UnionString].
type AddressAddressMapZoneReplaceResponse interface {
	ImplementsAddressAddressMapZoneReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AddressAddressMapZoneReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AddressAddressMapZoneReplaceResponseArray []interface{}

func (r AddressAddressMapZoneReplaceResponseArray) ImplementsAddressAddressMapZoneReplaceResponse() {}

type AddressAddressMapZoneDeleteResponseEnvelope struct {
	Errors   []AddressAddressMapZoneDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapZoneDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapZoneDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapZoneDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapZoneDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapZoneDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapZoneDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressAddressMapZoneDeleteResponseEnvelope]
type addressAddressMapZoneDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapZoneDeleteResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    addressAddressMapZoneDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapZoneDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressAddressMapZoneDeleteResponseEnvelopeErrors]
type addressAddressMapZoneDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapZoneDeleteResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    addressAddressMapZoneDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapZoneDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressAddressMapZoneDeleteResponseEnvelopeMessages]
type addressAddressMapZoneDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapZoneDeleteResponseEnvelopeSuccess bool

const (
	AddressAddressMapZoneDeleteResponseEnvelopeSuccessTrue AddressAddressMapZoneDeleteResponseEnvelopeSuccess = true
)

type AddressAddressMapZoneDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       addressAddressMapZoneDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapZoneDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressAddressMapZoneDeleteResponseEnvelopeResultInfo]
type addressAddressMapZoneDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapZoneReplaceResponseEnvelope struct {
	Errors   []AddressAddressMapZoneReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressAddressMapZoneReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressAddressMapZoneReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AddressAddressMapZoneReplaceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AddressAddressMapZoneReplaceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressAddressMapZoneReplaceResponseEnvelopeJSON       `json:"-"`
}

// addressAddressMapZoneReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [AddressAddressMapZoneReplaceResponseEnvelope]
type addressAddressMapZoneReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapZoneReplaceResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    addressAddressMapZoneReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// addressAddressMapZoneReplaceResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AddressAddressMapZoneReplaceResponseEnvelopeErrors]
type addressAddressMapZoneReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressAddressMapZoneReplaceResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    addressAddressMapZoneReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// addressAddressMapZoneReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AddressAddressMapZoneReplaceResponseEnvelopeMessages]
type addressAddressMapZoneReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressAddressMapZoneReplaceResponseEnvelopeSuccess bool

const (
	AddressAddressMapZoneReplaceResponseEnvelopeSuccessTrue AddressAddressMapZoneReplaceResponseEnvelopeSuccess = true
)

type AddressAddressMapZoneReplaceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       addressAddressMapZoneReplaceResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressAddressMapZoneReplaceResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [AddressAddressMapZoneReplaceResponseEnvelopeResultInfo]
type addressAddressMapZoneReplaceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressAddressMapZoneReplaceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
