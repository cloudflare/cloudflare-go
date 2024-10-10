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
func (r *AddressMapIPService) Update(ctx context.Context, addressMapID string, ipAddress string, params AddressMapIPUpdateParams, opts ...option.RequestOption) (res *AddressMapIPUpdateResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove an IP from a particular address map.
func (r *AddressMapIPService) Delete(ctx context.Context, addressMapID string, ipAddress string, body AddressMapIPDeleteParams, opts ...option.RequestOption) (res *AddressMapIPDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddressMapIPUpdateResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapIPUpdateResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapIPUpdateResponseResultInfo `json:"result_info"`
	JSON       addressMapIPUpdateResponseJSON       `json:"-"`
}

// addressMapIPUpdateResponseJSON contains the JSON metadata for the struct
// [AddressMapIPUpdateResponse]
type addressMapIPUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPUpdateResponseSuccess bool

const (
	AddressMapIPUpdateResponseSuccessTrue AddressMapIPUpdateResponseSuccess = true
)

func (r AddressMapIPUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       addressMapIPUpdateResponseResultInfoJSON `json:"-"`
}

// addressMapIPUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapIPUpdateResponseResultInfo]
type addressMapIPUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapIPDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapIPDeleteResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapIPDeleteResponseResultInfo `json:"result_info"`
	JSON       addressMapIPDeleteResponseJSON       `json:"-"`
}

// addressMapIPDeleteResponseJSON contains the JSON metadata for the struct
// [AddressMapIPDeleteResponse]
type addressMapIPDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapIPDeleteResponseSuccess bool

const (
	AddressMapIPDeleteResponseSuccessTrue AddressMapIPDeleteResponseSuccess = true
)

func (r AddressMapIPDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapIPDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapIPDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       addressMapIPDeleteResponseResultInfoJSON `json:"-"`
}

// addressMapIPDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapIPDeleteResponseResultInfo]
type addressMapIPDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapIPDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapIPDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapIPUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AddressMapIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AddressMapIPDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
