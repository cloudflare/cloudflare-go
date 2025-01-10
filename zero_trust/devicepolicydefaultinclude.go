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

// DevicePolicyDefaultIncludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultIncludeService] method instead.
type DevicePolicyDefaultIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultIncludeService(opts ...option.RequestOption) (r *DevicePolicyDefaultIncludeService) {
	r = &DevicePolicyDefaultIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) Update(ctx context.Context, params DevicePolicyDefaultIncludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyDefaultIncludeUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyDefaultIncludeService) Get(ctx context.Context, query DevicePolicyDefaultIncludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyDefaultIncludeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultIncludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelIncludeParam `json:"body,required"`
}

func (r DevicePolicyDefaultIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyDefaultIncludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultIncludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultIncludeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyDefaultIncludeUpdateResponseEnvelope]
type devicePolicyDefaultIncludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultIncludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultIncludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccessTrue DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultIncludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       devicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfo]
type devicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultIncludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultIncludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultIncludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultIncludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultIncludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultIncludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultIncludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyDefaultIncludeGetResponseEnvelope]
type devicePolicyDefaultIncludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultIncludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultIncludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultIncludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultIncludeGetResponseEnvelopeSuccessTrue DevicePolicyDefaultIncludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultIncludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultIncludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultIncludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       devicePolicyDefaultIncludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultIncludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyDefaultIncludeGetResponseEnvelopeResultInfo]
type devicePolicyDefaultIncludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultIncludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultIncludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
