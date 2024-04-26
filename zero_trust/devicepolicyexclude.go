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

// DevicePolicyExcludeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyExcludeService]
// method instead.
type DevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyExcludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyExcludeService(opts ...option.RequestOption) (r *DevicePolicyExcludeService) {
	r = &DevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Sets the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) Update(ctx context.Context, params DevicePolicyExcludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) List(ctx context.Context, query DevicePolicyExcludeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelExclude], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/devices/policy/exclude", query.AccountID)
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

// Fetches the list of routes excluded from the WARP client's tunnel.
func (r *DevicePolicyExcludeService) ListAutoPaging(ctx context.Context, query DevicePolicyExcludeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelExclude] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches the list of routes excluded from the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyExcludeService) Get(ctx context.Context, policyID string, query DevicePolicyExcludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelExclude, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyExcludeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/policy/%s/exclude", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SplitTunnelExclude struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address string `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host string                 `json:"host"`
	JSON splitTunnelExcludeJSON `json:"-"`
}

// splitTunnelExcludeJSON contains the JSON metadata for the struct
// [SplitTunnelExclude]
type splitTunnelExcludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelExcludeJSON) RawJSON() string {
	return r.raw
}

type SplitTunnelExcludeParam struct {
	// The address in CIDR format to exclude from the tunnel. If `address` is present,
	// `host` must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the Split Tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If `host` is present, `address` must
	// not be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyExcludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelExcludeParam `json:"body,required"`
}

func (r DevicePolicyExcludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyExcludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeUpdateResponseEnvelope]
type devicePolicyExcludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyExcludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyExcludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeUpdateResponseEnvelopeSuccessTrue DevicePolicyExcludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyExcludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyExcludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyExcludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyExcludeUpdateResponseEnvelopeResultInfo]
type devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyExcludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyExcludeListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyExcludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyExcludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelExclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyExcludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyExcludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyExcludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyExcludeGetResponseEnvelope]
type devicePolicyExcludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyExcludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyExcludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyExcludeGetResponseEnvelopeSuccessTrue DevicePolicyExcludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyExcludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyExcludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyExcludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       devicePolicyExcludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyExcludeGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyExcludeGetResponseEnvelopeResultInfo]
type devicePolicyExcludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyExcludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyExcludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
