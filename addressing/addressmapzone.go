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

// AddressMapZoneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressMapZoneService] method instead.
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
func (r *AddressMapZoneService) Update(ctx context.Context, addressMapID string, params AddressMapZoneUpdateParams, opts ...option.RequestOption) (res *AddressMapZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", params.AccountID, addressMapID, params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Remove a zone as a member of a particular address map.
func (r *AddressMapZoneService) Delete(ctx context.Context, addressMapID string, body AddressMapZoneDeleteParams, opts ...option.RequestOption) (res *AddressMapZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", body.AccountID, addressMapID, body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AddressMapZoneUpdateResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapZoneUpdateResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapZoneUpdateResponseResultInfo `json:"result_info"`
	JSON       addressMapZoneUpdateResponseJSON       `json:"-"`
}

// addressMapZoneUpdateResponseJSON contains the JSON metadata for the struct
// [AddressMapZoneUpdateResponse]
type addressMapZoneUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapZoneUpdateResponseSuccess bool

const (
	AddressMapZoneUpdateResponseSuccessTrue AddressMapZoneUpdateResponseSuccess = true
)

func (r AddressMapZoneUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapZoneUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapZoneUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       addressMapZoneUpdateResponseResultInfoJSON `json:"-"`
}

// addressMapZoneUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapZoneUpdateResponseResultInfo]
type addressMapZoneUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapZoneDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    AddressMapZoneDeleteResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapZoneDeleteResponseResultInfo `json:"result_info"`
	JSON       addressMapZoneDeleteResponseJSON       `json:"-"`
}

// addressMapZoneDeleteResponseJSON contains the JSON metadata for the struct
// [AddressMapZoneDeleteResponse]
type addressMapZoneDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressMapZoneDeleteResponseSuccess bool

const (
	AddressMapZoneDeleteResponseSuccessTrue AddressMapZoneDeleteResponseSuccess = true
)

func (r AddressMapZoneDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapZoneDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapZoneDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       addressMapZoneDeleteResponseResultInfoJSON `json:"-"`
}

// addressMapZoneDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AddressMapZoneDeleteResponseResultInfo]
type addressMapZoneDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapZoneDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapZoneDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

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

type AddressMapZoneDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
