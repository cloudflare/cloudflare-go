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

// DevicePolicyCustomIncludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomIncludeService] method instead.
type DevicePolicyCustomIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCustomIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomIncludeService(opts ...option.RequestOption) (r *DevicePolicyCustomIncludeService) {
	r = &DevicePolicyCustomIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) Update(ctx context.Context, policyID string, params DevicePolicyCustomIncludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyCustomIncludeUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomIncludeService) Get(ctx context.Context, policyID string, query DevicePolicyCustomIncludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyCustomIncludeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/include", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyCustomIncludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelIncludeParam `json:"body,required"`
}

func (r DevicePolicyCustomIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyCustomIncludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomIncludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomIncludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomIncludeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyCustomIncludeUpdateResponseEnvelope]
type devicePolicyCustomIncludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomIncludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomIncludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccessTrue DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomIncludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomIncludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       devicePolicyCustomIncludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomIncludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyCustomIncludeUpdateResponseEnvelopeResultInfo]
type devicePolicyCustomIncludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomIncludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomIncludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCustomIncludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomIncludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomIncludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomIncludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomIncludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomIncludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyCustomIncludeGetResponseEnvelope]
type devicePolicyCustomIncludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomIncludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomIncludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomIncludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomIncludeGetResponseEnvelopeSuccessTrue DevicePolicyCustomIncludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomIncludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomIncludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomIncludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       devicePolicyCustomIncludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomIncludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyCustomIncludeGetResponseEnvelopeResultInfo]
type devicePolicyCustomIncludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomIncludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomIncludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
