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

// DevicePolicyDefaultFallbackDomainService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultFallbackDomainService] method instead.
type DevicePolicyDefaultFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultFallbackDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyDefaultFallbackDomainService) {
	r = &DevicePolicyDefaultFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyDefaultFallbackDomainService) Update(ctx context.Context, params DevicePolicyDefaultFallbackDomainUpdateParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	var env DevicePolicyDefaultFallbackDomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/fallback_domains", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyDefaultFallbackDomainService) Get(ctx context.Context, query DevicePolicyDefaultFallbackDomainGetParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	var env DevicePolicyDefaultFallbackDomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/fallback_domains", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultFallbackDomainUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Domains   []FallbackDomainParam `json:"domains,required"`
}

func (r DevicePolicyDefaultFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Domains)
}

type DevicePolicyDefaultFallbackDomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [DevicePolicyDefaultFallbackDomainUpdateResponseEnvelope]
type devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultFallbackDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccessTrue DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                               `json:"total_count"`
	JSON       devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfo]
type devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultFallbackDomainUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyDefaultFallbackDomainGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultFallbackDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultFallbackDomainGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultFallbackDomainGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DevicePolicyDefaultFallbackDomainGetResponseEnvelope]
type devicePolicyDefaultFallbackDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultFallbackDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultFallbackDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccessTrue DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultFallbackDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                            `json:"total_count"`
	JSON       devicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [DevicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfo]
type devicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultFallbackDomainGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
