// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DevicePolicyDefaultExcludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultExcludeService] method instead.
type DevicePolicyDefaultExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultExcludeService(opts ...option.RequestOption) (r *DevicePolicyDefaultExcludeService) {
	r = &DevicePolicyDefaultExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) Update(ctx context.Context, params DevicePolicyDefaultExcludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	var env DevicePolicyDefaultExcludeUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyDefaultExcludeService) Get(ctx context.Context, query DevicePolicyDefaultExcludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	var env DevicePolicyDefaultExcludeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultExcludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelExcludeParam `json:"body,required"`
}

func (r DevicePolicyDefaultExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyDefaultExcludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultExcludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultExcludeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyDefaultExcludeUpdateResponseEnvelope]
type devicePolicyDefaultExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultExcludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccessTrue DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultExcludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       devicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfo]
type devicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultExcludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultExcludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultExcludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultExcludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultExcludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultExcludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultExcludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyDefaultExcludeGetResponseEnvelope]
type devicePolicyDefaultExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultExcludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultExcludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultExcludeGetResponseEnvelopeSuccessTrue DevicePolicyDefaultExcludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultExcludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultExcludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultExcludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       devicePolicyDefaultExcludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultExcludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyDefaultExcludeGetResponseEnvelopeResultInfo]
type devicePolicyDefaultExcludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultExcludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultExcludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
