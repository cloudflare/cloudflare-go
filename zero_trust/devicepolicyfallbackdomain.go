// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DevicePolicyFallbackDomainService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyFallbackDomainService] method instead.
type DevicePolicyFallbackDomainService struct {
	Options []option.RequestOption
}

// NewDevicePolicyFallbackDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyFallbackDomainService(opts ...option.RequestOption) (r *DevicePolicyFallbackDomainService) {
	r = &DevicePolicyFallbackDomainService{}
	r.Options = opts
	return
}

// Sets the list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead. This will only apply to the
// specified device settings profile.
func (r *DevicePolicyFallbackDomainService) Update(ctx context.Context, policyID string, params DevicePolicyFallbackDomainUpdateParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyFallbackDomainService) List(ctx context.Context, query DevicePolicyFallbackDomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FallbackDomain], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/policy/fallback_domains", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches a list of domains to bypass Gateway DNS resolution. These domains will
// use the specified local DNS resolver instead.
func (r *DevicePolicyFallbackDomainService) ListAutoPaging(ctx context.Context, query DevicePolicyFallbackDomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FallbackDomain] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches the list of domains to bypass Gateway DNS resolution from a specified
// device settings profile. These domains will use the specified local DNS resolver
// instead.
func (r *DevicePolicyFallbackDomainService) Get(ctx context.Context, policyID string, query DevicePolicyFallbackDomainGetParams, opts ...option.RequestOption) (res *[]FallbackDomain, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyFallbackDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/fallback_domains", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FallbackDomain struct {
	// The domain suffix to match when resolving locally.
	Suffix string `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description string `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer []interface{}      `json:"dns_server"`
	JSON      fallbackDomainJSON `json:"-"`
}

// fallbackDomainJSON contains the JSON metadata for the struct [FallbackDomain]
type fallbackDomainJSON struct {
	Suffix      apijson.Field
	Description apijson.Field
	DNSServer   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FallbackDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fallbackDomainJSON) RawJSON() string {
	return r.raw
}

type FallbackDomainParam struct {
	// The domain suffix to match when resolving locally.
	Suffix param.Field[string] `json:"suffix,required"`
	// A description of the fallback domain, displayed in the client UI.
	Description param.Field[string] `json:"description"`
	// A list of IP addresses to handle domain resolution.
	DNSServer param.Field[[]interface{}] `json:"dns_server"`
}

func (r FallbackDomainParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyFallbackDomainUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id,required"`
	Body      []FallbackDomainParam `json:"body,required"`
}

func (r DevicePolicyFallbackDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyFallbackDomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyFallbackDomainUpdateResponseEnvelope]
type devicePolicyFallbackDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyFallbackDomainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyFallbackDomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyFallbackDomainUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyFallbackDomainListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyFallbackDomainGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyFallbackDomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []FallbackDomain      `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyFallbackDomainGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyFallbackDomainGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyFallbackDomainGetResponseEnvelope]
type devicePolicyFallbackDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyFallbackDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyFallbackDomainGetResponseEnvelopeSuccess bool

const (
	DevicePolicyFallbackDomainGetResponseEnvelopeSuccessTrue DevicePolicyFallbackDomainGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyFallbackDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyFallbackDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo]
type devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyFallbackDomainGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyFallbackDomainGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
