// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AddressMapZoneService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressMapZoneService] method
// instead.
type AddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAddressMapZoneService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapZoneService(opts ...option.RequestOption) (r *AddressMapZoneService) {
	r = &AddressMapZoneService{}
	r.Options = opts
	return
}

// Add a zone as a member of a particular address map.
func (r *AddressMapZoneService) Update(ctx context.Context, addressMapID string, params AddressMapZoneUpdateParams, opts ...option.RequestOption) (res *[]AddressMapZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapZoneUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", params.AccountID, addressMapID, params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove a zone as a member of a particular address map.
func (r *AddressMapZoneService) Delete(ctx context.Context, addressMapID string, body AddressMapZoneDeleteParams, opts ...option.RequestOption) (res *[]AddressMapZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapZoneDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", body.AccountID, addressMapID, body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressMapZoneUpdateResponse = interface{}

type AddressMapZoneDeleteResponse = interface{}

type AddressMapZoneUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapZoneUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapZoneUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapZoneUpdateResponse                 `json:"result,nullable"`
	ResultInfo AddressMapZoneUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapZoneUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressMapZoneUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapZoneUpdateResponseEnvelope]
type addressMapZoneUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapZoneUpdateResponseEnvelopeSuccess bool

const (
	AddressMapZoneUpdateResponseEnvelopeSuccessTrue AddressMapZoneUpdateResponseEnvelopeSuccess = true
)

func (r AddressMapZoneUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapZoneUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapZoneUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       addressMapZoneUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapZoneUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressMapZoneUpdateResponseEnvelopeResultInfo]
type addressMapZoneUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapZoneDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapZoneDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapZoneDeleteResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapZoneDeleteResponse                 `json:"result,nullable"`
	ResultInfo AddressMapZoneDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapZoneDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapZoneDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapZoneDeleteResponseEnvelope]
type addressMapZoneDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapZoneDeleteResponseEnvelopeSuccess bool

const (
	AddressMapZoneDeleteResponseEnvelopeSuccessTrue AddressMapZoneDeleteResponseEnvelopeSuccess = true
)

func (r AddressMapZoneDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapZoneDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapZoneDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       addressMapZoneDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapZoneDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AddressMapZoneDeleteResponseEnvelopeResultInfo]
type addressMapZoneDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
