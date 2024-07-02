// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// DevicePolicyIncludeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyIncludeService] method instead.
type DevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyIncludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyIncludeService(opts ...option.RequestOption) (r *DevicePolicyIncludeService) {
	r = &DevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Sets the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) Update(ctx context.Context, params DevicePolicyIncludeUpdateParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyIncludeUpdateResponseEnvelope
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
func (r *DevicePolicyIncludeService) List(ctx context.Context, query DevicePolicyIncludeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SplitTunnelInclude], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/policy/include", query.AccountID)
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

// Fetches the list of routes included in the WARP client's tunnel.
func (r *DevicePolicyIncludeService) ListAutoPaging(ctx context.Context, query DevicePolicyIncludeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SplitTunnelInclude] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Fetches the list of routes included in the WARP client's tunnel for a specific
// device settings profile.
func (r *DevicePolicyIncludeService) Get(ctx context.Context, policyID string, query DevicePolicyIncludeGetParams, opts ...option.RequestOption) (res *[]SplitTunnelInclude, err error) {
	var env DevicePolicyIncludeGetResponseEnvelope
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

type SplitTunnelInclude struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                 `json:"host"`
	JSON splitTunnelIncludeJSON `json:"-"`
}

// splitTunnelIncludeJSON contains the JSON metadata for the struct
// [SplitTunnelInclude]
type splitTunnelIncludeJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r splitTunnelIncludeJSON) RawJSON() string {
	return r.raw
}

type SplitTunnelIncludeParam struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r SplitTunnelIncludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevicePolicyIncludeUpdateParams struct {
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SplitTunnelIncludeParam `json:"body,required"`
}

func (r DevicePolicyIncludeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DevicePolicyIncludeUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeUpdateResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyIncludeUpdateResponseEnvelope]
type devicePolicyIncludeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyIncludeUpdateResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeUpdateResponseEnvelopeSuccessTrue DevicePolicyIncludeUpdateResponseEnvelopeSuccess = true
)

func (r DevicePolicyIncludeUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyIncludeUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyIncludeUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyIncludeUpdateResponseEnvelopeResultInfo]
type devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DevicePolicyIncludeListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyIncludeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyIncludeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []SplitTunnelInclude  `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyIncludeGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyIncludeGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyIncludeGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DevicePolicyIncludeGetResponseEnvelope]
type devicePolicyIncludeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyIncludeGetResponseEnvelopeSuccess bool

const (
	DevicePolicyIncludeGetResponseEnvelopeSuccessTrue DevicePolicyIncludeGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyIncludeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyIncludeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyIncludeGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       devicePolicyIncludeGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyIncludeGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DevicePolicyIncludeGetResponseEnvelopeResultInfo]
type devicePolicyIncludeGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyIncludeGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyIncludeGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
