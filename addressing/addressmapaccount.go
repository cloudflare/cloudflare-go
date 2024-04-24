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
func (r *AddressMapAccountService) Update(ctx context.Context, addressMapID string, params AddressMapAccountUpdateParams, opts ...option.RequestOption) (res *[]AddressMapAccountUpdateResponse, err error) {
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
func (r *AddressMapAccountService) Delete(ctx context.Context, addressMapID string, body AddressMapAccountDeleteParams, opts ...option.RequestOption) (res *[]AddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressMapAccountDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", body.AccountID, addressMapID, body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressMapAccountUpdateResponse = interface{}

type AddressMapAccountDeleteResponse = interface{}

type AddressMapAccountUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapAccountUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapAccountUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapAccountUpdateResponse                 `json:"result,nullable"`
	ResultInfo AddressMapAccountUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapAccountUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressMapAccountUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapAccountUpdateResponseEnvelope]
type addressMapAccountUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapAccountDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapAccountDeleteResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapAccountDeleteResponse                 `json:"result,nullable"`
	ResultInfo AddressMapAccountDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapAccountDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapAccountDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressMapAccountDeleteResponseEnvelope]
type addressMapAccountDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
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
