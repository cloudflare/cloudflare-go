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

// DevicePolicyCustomFallbackDomainService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyCustomFallbackDomainService] method instead.
type DevicePolicyCustomFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyCustomFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyCustomFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyCustomFallbackDomainService) {
	r = &DevicePolicyCustomFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyCustomFallbackDomainService) Update(ctx context.Context, policyID string, params DevicePolicyCustomFallbackDomainUpdateParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	var env DevicePolicyCustomFallbackDomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyCustomFallbackDomainService) Get(ctx context.Context, policyID string, query DevicePolicyCustomFallbackDomainGetParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	var env DevicePolicyCustomFallbackDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyCustomFallbackDomainUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Domains   []FallbackDomainParam `json:"domains,required"`
}

func (r DevicePolicyCustomFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Domains)
}

type DevicePolicyCustomFallbackDomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomFallbackDomainUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomFallbackDomainUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DevicePolicyCustomFallbackDomainUpdateResponseEnvelope]
type devicePolicyCustomFallbackDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomFallbackDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomFallbackDomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccessTrue DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                              `json:"total_count"`
	JSON       devicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfo]
type devicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomFallbackDomainUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyCustomFallbackDomainGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyCustomFallbackDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyCustomFallbackDomainGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyCustomFallbackDomainGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DevicePolicyCustomFallbackDomainGetResponseEnvelope]
type devicePolicyCustomFallbackDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomFallbackDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomFallbackDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccess bool

const (
	DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccessTrue DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyCustomFallbackDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       devicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [DevicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfo]
type devicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyCustomFallbackDomainGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
