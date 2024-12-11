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

// DevicePolicyCustomExcludeService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomExcludeService] method instead.
type DevicePolicyCustomExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCustomExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomExcludeService(opts ...option.RequestOption) (r *DevicePolicyCustomExcludeService) {
	r = &DevicePolicyCustomExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomExcludeService) Update(ctx context.Context, policyID string, params DevicePolicyCustomExcludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	var env DevicePolicyCustomExcludeUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/exclude", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyCustomExcludeService) Get(ctx context.Context, policyID string, query DevicePolicyCustomExcludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	var env DevicePolicyCustomExcludeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/exclude", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyCustomExcludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelExcludeParam `json:"body,required"`
}

func (r DevicePolicyCustomExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyCustomExcludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomExcludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomExcludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomExcludeUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyCustomExcludeUpdateResponseEnvelope]
type devicePolicyCustomExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomExcludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccessTrue DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomExcludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomExcludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       devicePolicyCustomExcludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomExcludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyCustomExcludeUpdateResponseEnvelopeResultInfo]
type devicePolicyCustomExcludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomExcludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomExcludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCustomExcludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomExcludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomExcludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomExcludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomExcludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomExcludeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyCustomExcludeGetResponseEnvelope]
type devicePolicyCustomExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomExcludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomExcludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomExcludeGetResponseEnvelopeSuccessTrue DevicePolicyCustomExcludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomExcludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomExcludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomExcludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       devicePolicyCustomExcludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomExcludeGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyCustomExcludeGetResponseEnvelopeResultInfo]
type devicePolicyCustomExcludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomExcludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomExcludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
