// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AddressMapAccountService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressMapAccountService] method instead.
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
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", params.AccountID, addressMapID, params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove an account as a member of a particular address map.
func (r *AddressMapAccountService) Delete(ctx context.Context, addressMapID string, body AddressMapAccountDeleteParams, opts ...option.RequestOption) (res *AddressMapAccountDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", body.AccountID, addressMapID, body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddressMapAccountUpdateResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapAccountUpdateResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapAccountUpdateResponseResultInfo `json:"result_info"`
	JSON       addressMapAccountUpdateResponseJSON       `json:"-"`
}

// addressMapAccountUpdateResponseJSON contains the JSON metadata for the struct
// [AddressMapAccountUpdateResponse]
type addressMapAccountUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapAccountUpdateResponseSuccess bool

const (
	AddressMapAccountUpdateResponseSuccessTrue AddressMapAccountUpdateResponseSuccess = true
)

func (r AddressMapAccountUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapAccountUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapAccountUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       addressMapAccountUpdateResponseResultInfoJSON `json:"-"`
}

// addressMapAccountUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapAccountUpdateResponseResultInfo]
type addressMapAccountUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapAccountDeleteResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapAccountDeleteResponseResultInfo `json:"result_info"`
	JSON       addressMapAccountDeleteResponseJSON       `json:"-"`
}

// addressMapAccountDeleteResponseJSON contains the JSON metadata for the struct
// [AddressMapAccountDeleteResponse]
type addressMapAccountDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapAccountDeleteResponseSuccess bool

const (
	AddressMapAccountDeleteResponseSuccessTrue AddressMapAccountDeleteResponseSuccess = true
)

func (r AddressMapAccountDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapAccountDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapAccountDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       addressMapAccountDeleteResponseResultInfoJSON `json:"-"`
}

// addressMapAccountDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapAccountDeleteResponseResultInfo]
type addressMapAccountDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapAccountDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapAccountDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapAccountUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapAccountDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
