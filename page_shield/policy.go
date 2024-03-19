// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PolicyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r *PolicyService) {
	r = &PolicyService{}
	r.Options = opts
	return
}

// Create a Page Shield policy.
func (r *PolicyService) New(ctx context.Context, params PolicyNewParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a Page Shield policy by ID.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", params.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists all Page Shield policies.
func (r *PolicyService) List(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) (res *[]PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env PolicyListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/policies", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Page Shield policy by ID.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", body.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a Page Shield policy by ID.
func (r *PolicyService) Get(ctx context.Context, policyID string, query PolicyGetParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", query.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PageShieldPageshieldPolicy struct {
	// The ID of the policy
	ID string `json:"id"`
	// The action to take if the expression matches
	Action PageShieldPageshieldPolicyAction `json:"action"`
	// A description for the policy
	Description string `json:"description"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression"`
	// The policy which will be applied
	Value string                         `json:"value"`
	JSON  pageShieldPageshieldPolicyJSON `json:"-"`
}

// pageShieldPageshieldPolicyJSON contains the JSON metadata for the struct
// [PageShieldPageshieldPolicy]
type pageShieldPageshieldPolicyJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPageshieldPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldPageshieldPolicyJSON) RawJSON() string {
	return r.raw
}

// The action to take if the expression matches
type PageShieldPageshieldPolicyAction string

const (
	PageShieldPageshieldPolicyActionAllow PageShieldPageshieldPolicyAction = "allow"
	PageShieldPageshieldPolicyActionLog   PageShieldPageshieldPolicyAction = "log"
)

type PolicyNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take if the expression matches
	Action param.Field[PolicyNewParamsAction] `json:"action"`
	// A description for the policy
	Description param.Field[string] `json:"description"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression"`
	// The policy which will be applied
	Value param.Field[string] `json:"value"`
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PolicyNewParamsAction string

const (
	PolicyNewParamsActionAllow PolicyNewParamsAction = "allow"
	PolicyNewParamsActionLog   PolicyNewParamsAction = "log"
)

type PolicyUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take if the expression matches
	Action param.Field[PolicyUpdateParamsAction] `json:"action"`
	// A description for the policy
	Description param.Field[string] `json:"description"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression"`
	// The policy which will be applied
	Value param.Field[string] `json:"value"`
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PolicyUpdateParamsAction string

const (
	PolicyUpdateParamsActionAllow PolicyUpdateParamsAction = "allow"
	PolicyUpdateParamsActionLog   PolicyUpdateParamsAction = "log"
)

type PolicyListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PolicyListResponseEnvelope struct {
	Errors   []PolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldPageshieldPolicy         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       policyListResponseEnvelopeJSON       `json:"-"`
}

// policyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyListResponseEnvelope]
type policyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PolicyListResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    policyListResponseEnvelopeErrorsJSON `json:"-"`
}

// policyListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PolicyListResponseEnvelopeErrors]
type policyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PolicyListResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    policyListResponseEnvelopeMessagesJSON `json:"-"`
}

// policyListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PolicyListResponseEnvelopeMessages]
type policyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyListResponseEnvelopeSuccess bool

const (
	PolicyListResponseEnvelopeSuccessTrue PolicyListResponseEnvelopeSuccess = true
)

type PolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       policyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// policyListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [PolicyListResponseEnvelopeResultInfo]
type policyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PolicyDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PolicyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
