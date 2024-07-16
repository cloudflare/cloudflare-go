// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// AddressMapIPService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressMapIPService] method instead.
type AddressMapIPService struct {
	Options []option.RequestOption
}

// NewAddressMapIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapIPService(opts ...option.RequestOption) (r *AddressMapIPService) {
	r = &AddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AddressMapIPService) Update(ctx context.Context, addressMapID string, ipAddress string, params AddressMapIPUpdateParams, opts ...option.RequestOption) (res *[]AddressMapIPUpdateResponse, err error) {
	var env AddressMapIPUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if ipAddress == "" {
		err = errors.New("missing required ip_address parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", params.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Remove an IP from a particular address map.
func (r *AddressMapIPService) Delete(ctx context.Context, addressMapID string, ipAddress string, body AddressMapIPDeleteParams, opts ...option.RequestOption) (res *[]AddressMapIPDeleteResponse, err error) {
	var env AddressMapIPDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	if ipAddress == "" {
		err = errors.New("missing required ip_address parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", body.AccountID, addressMapID, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressMapIPUpdateResponse = interface{}

type AddressMapIPDeleteResponse = interface{}

type AddressMapIPUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapIPUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapIPUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapIPUpdateResponse                 `json:"result,nullable"`
	ResultInfo AddressMapIPUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapIPUpdateResponseEnvelopeJSON       `json:"-"`
}

// addressMapIPUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapIPUpdateResponseEnvelope]
type addressMapIPUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPUpdateResponseEnvelopeSuccess bool

const (
	AddressMapIPUpdateResponseEnvelopeSuccessTrue AddressMapIPUpdateResponseEnvelopeSuccess = true
)

func (r AddressMapIPUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       addressMapIPUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapIPUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressMapIPUpdateResponseEnvelopeResultInfo]
type addressMapIPUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapIPDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapIPDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapIPDeleteResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AddressMapIPDeleteResponse                 `json:"result,nullable"`
	ResultInfo AddressMapIPDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       addressMapIPDeleteResponseEnvelopeJSON       `json:"-"`
}

// addressMapIPDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapIPDeleteResponseEnvelope]
type addressMapIPDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPDeleteResponseEnvelopeSuccess bool

const (
	AddressMapIPDeleteResponseEnvelopeSuccessTrue AddressMapIPDeleteResponseEnvelopeSuccess = true
)

func (r AddressMapIPDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       addressMapIPDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// addressMapIPDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AddressMapIPDeleteResponseEnvelopeResultInfo]
type addressMapIPDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
