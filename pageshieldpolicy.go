// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PageShieldPolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageShieldPolicyService] method
// instead.
type PageShieldPolicyService struct {
	Options []option.RequestOption
}

// NewPageShieldPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageShieldPolicyService(opts ...option.RequestOption) (r *PageShieldPolicyService) {
	r = &PageShieldPolicyService{}
	r.Options = opts
	return
}

// Create a Page Shield policy.
func (r *PageShieldPolicyService) New(ctx context.Context, params PageShieldPolicyNewParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update a Page Shield policy by ID.
func (r *PageShieldPolicyService) Update(ctx context.Context, policyID string, params PageShieldPolicyUpdateParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", params.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists all Page Shield policies.
func (r *PageShieldPolicyService) List(ctx context.Context, query PageShieldPolicyListParams, opts ...option.RequestOption) (res *[]PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldPolicyListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/policies", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a Page Shield policy by ID.
func (r *PageShieldPolicyService) Delete(ctx context.Context, policyID string, body PageShieldPolicyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", body.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Fetches a Page Shield policy by ID.
func (r *PageShieldPolicyService) Get(ctx context.Context, policyID string, query PageShieldPolicyGetParams, opts ...option.RequestOption) (res *PageShieldPageshieldPolicy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", query.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// The action to take if the expression matches
type PageShieldPageshieldPolicyAction string

const (
	PageShieldPageshieldPolicyActionAllow PageShieldPageshieldPolicyAction = "allow"
	PageShieldPageshieldPolicyActionLog   PageShieldPageshieldPolicyAction = "log"
)

type PageShieldPolicyNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take if the expression matches
	Action param.Field[PageShieldPolicyNewParamsAction] `json:"action"`
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

func (r PageShieldPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PageShieldPolicyNewParamsAction string

const (
	PageShieldPolicyNewParamsActionAllow PageShieldPolicyNewParamsAction = "allow"
	PageShieldPolicyNewParamsActionLog   PageShieldPolicyNewParamsAction = "log"
)

type PageShieldPolicyUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take if the expression matches
	Action param.Field[PageShieldPolicyUpdateParamsAction] `json:"action"`
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

func (r PageShieldPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PageShieldPolicyUpdateParamsAction string

const (
	PageShieldPolicyUpdateParamsActionAllow PageShieldPolicyUpdateParamsAction = "allow"
	PageShieldPolicyUpdateParamsActionLog   PageShieldPolicyUpdateParamsAction = "log"
)

type PageShieldPolicyListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageShieldPolicyListResponseEnvelope struct {
	Errors   []PageShieldPolicyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldPolicyListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldPageshieldPolicy                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PageShieldPolicyListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageShieldPolicyListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageShieldPolicyListResponseEnvelopeJSON       `json:"-"`
}

// pageShieldPolicyListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageShieldPolicyListResponseEnvelope]
type pageShieldPolicyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPolicyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPolicyListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageShieldPolicyListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldPolicyListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageShieldPolicyListResponseEnvelopeErrors]
type pageShieldPolicyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPolicyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPolicyListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageShieldPolicyListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldPolicyListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageShieldPolicyListResponseEnvelopeMessages]
type pageShieldPolicyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPolicyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageShieldPolicyListResponseEnvelopeSuccess bool

const (
	PageShieldPolicyListResponseEnvelopeSuccessTrue PageShieldPolicyListResponseEnvelopeSuccess = true
)

type PageShieldPolicyListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       pageShieldPolicyListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageShieldPolicyListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [PageShieldPolicyListResponseEnvelopeResultInfo]
type pageShieldPolicyListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldPolicyListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldPolicyDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageShieldPolicyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
