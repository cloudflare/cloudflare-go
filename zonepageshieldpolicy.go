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

// ZonePageShieldPolicyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePageShieldPolicyService]
// method instead.
type ZonePageShieldPolicyService struct {
	Options []option.RequestOption
}

// NewZonePageShieldPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageShieldPolicyService(opts ...option.RequestOption) (r *ZonePageShieldPolicyService) {
	r = &ZonePageShieldPolicyService{}
	r.Options = opts
	return
}

// Create a Page Shield policy.
func (r *ZonePageShieldPolicyService) New(ctx context.Context, zoneID string, body ZonePageShieldPolicyNewParams, opts ...option.RequestOption) (res *ZonePageShieldPolicyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Get(ctx context.Context, zoneID string, policyID string, opts ...option.RequestOption) (res *ZonePageShieldPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Update(ctx context.Context, zoneID string, policyID string, body ZonePageShieldPolicyUpdateParams, opts ...option.RequestOption) (res *ZonePageShieldPolicyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Page Shield policies.
func (r *ZonePageShieldPolicyService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZonePageShieldPolicyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Page Shield policy by ID.
func (r *ZonePageShieldPolicyService) Delete(ctx context.Context, zoneID string, policyID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", zoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZonePageShieldPolicyNewResponse struct {
	// The ID of the policy
	ID string `json:"id"`
	// The action to take if the expression matches
	Action ZonePageShieldPolicyNewResponseAction `json:"action"`
	// A description for the policy
	Description string `json:"description"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression"`
	// The policy which will be applied
	Value string                              `json:"value"`
	JSON  zonePageShieldPolicyNewResponseJSON `json:"-"`
}

// zonePageShieldPolicyNewResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyNewResponse]
type zonePageShieldPolicyNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyNewResponseAction string

const (
	ZonePageShieldPolicyNewResponseActionAllow ZonePageShieldPolicyNewResponseAction = "allow"
	ZonePageShieldPolicyNewResponseActionLog   ZonePageShieldPolicyNewResponseAction = "log"
)

type ZonePageShieldPolicyGetResponse struct {
	// The ID of the policy
	ID string `json:"id"`
	// The action to take if the expression matches
	Action ZonePageShieldPolicyGetResponseAction `json:"action"`
	// A description for the policy
	Description string `json:"description"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression"`
	// The policy which will be applied
	Value string                              `json:"value"`
	JSON  zonePageShieldPolicyGetResponseJSON `json:"-"`
}

// zonePageShieldPolicyGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyGetResponse]
type zonePageShieldPolicyGetResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyGetResponseAction string

const (
	ZonePageShieldPolicyGetResponseActionAllow ZonePageShieldPolicyGetResponseAction = "allow"
	ZonePageShieldPolicyGetResponseActionLog   ZonePageShieldPolicyGetResponseAction = "log"
)

type ZonePageShieldPolicyUpdateResponse struct {
	// The ID of the policy
	ID string `json:"id"`
	// The action to take if the expression matches
	Action ZonePageShieldPolicyUpdateResponseAction `json:"action"`
	// A description for the policy
	Description string `json:"description"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression"`
	// The policy which will be applied
	Value string                                 `json:"value"`
	JSON  zonePageShieldPolicyUpdateResponseJSON `json:"-"`
}

// zonePageShieldPolicyUpdateResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyUpdateResponse]
type zonePageShieldPolicyUpdateResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyUpdateResponseAction string

const (
	ZonePageShieldPolicyUpdateResponseActionAllow ZonePageShieldPolicyUpdateResponseAction = "allow"
	ZonePageShieldPolicyUpdateResponseActionLog   ZonePageShieldPolicyUpdateResponseAction = "log"
)

type ZonePageShieldPolicyListResponse struct {
	Errors     []ZonePageShieldPolicyListResponseError    `json:"errors"`
	Messages   []ZonePageShieldPolicyListResponseMessage  `json:"messages"`
	Result     []ZonePageShieldPolicyListResponseResult   `json:"result"`
	ResultInfo ZonePageShieldPolicyListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZonePageShieldPolicyListResponseSuccess `json:"success"`
	JSON    zonePageShieldPolicyListResponseJSON    `json:"-"`
}

// zonePageShieldPolicyListResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldPolicyListResponse]
type zonePageShieldPolicyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPolicyListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zonePageShieldPolicyListResponseErrorJSON `json:"-"`
}

// zonePageShieldPolicyListResponseErrorJSON contains the JSON metadata for the
// struct [ZonePageShieldPolicyListResponseError]
type zonePageShieldPolicyListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPolicyListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zonePageShieldPolicyListResponseMessageJSON `json:"-"`
}

// zonePageShieldPolicyListResponseMessageJSON contains the JSON metadata for the
// struct [ZonePageShieldPolicyListResponseMessage]
type zonePageShieldPolicyListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldPolicyListResponseResult struct {
	// The ID of the policy
	ID string `json:"id"`
	// The action to take if the expression matches
	Action ZonePageShieldPolicyListResponseResultAction `json:"action"`
	// A description for the policy
	Description string `json:"description"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression"`
	// The policy which will be applied
	Value string                                     `json:"value"`
	JSON  zonePageShieldPolicyListResponseResultJSON `json:"-"`
}

// zonePageShieldPolicyListResponseResultJSON contains the JSON metadata for the
// struct [ZonePageShieldPolicyListResponseResult]
type zonePageShieldPolicyListResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyListResponseResultAction string

const (
	ZonePageShieldPolicyListResponseResultActionAllow ZonePageShieldPolicyListResponseResultAction = "allow"
	ZonePageShieldPolicyListResponseResultActionLog   ZonePageShieldPolicyListResponseResultAction = "log"
)

type ZonePageShieldPolicyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       zonePageShieldPolicyListResponseResultInfoJSON `json:"-"`
}

// zonePageShieldPolicyListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZonePageShieldPolicyListResponseResultInfo]
type zonePageShieldPolicyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldPolicyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageShieldPolicyListResponseSuccess bool

const (
	ZonePageShieldPolicyListResponseSuccessTrue ZonePageShieldPolicyListResponseSuccess = true
)

type ZonePageShieldPolicyNewParams struct {
	// The action to take if the expression matches
	Action param.Field[ZonePageShieldPolicyNewParamsAction] `json:"action"`
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

func (r ZonePageShieldPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyNewParamsAction string

const (
	ZonePageShieldPolicyNewParamsActionAllow ZonePageShieldPolicyNewParamsAction = "allow"
	ZonePageShieldPolicyNewParamsActionLog   ZonePageShieldPolicyNewParamsAction = "log"
)

type ZonePageShieldPolicyUpdateParams struct {
	// The action to take if the expression matches
	Action param.Field[ZonePageShieldPolicyUpdateParamsAction] `json:"action"`
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

func (r ZonePageShieldPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type ZonePageShieldPolicyUpdateParamsAction string

const (
	ZonePageShieldPolicyUpdateParamsActionAllow ZonePageShieldPolicyUpdateParamsAction = "allow"
	ZonePageShieldPolicyUpdateParamsActionLog   ZonePageShieldPolicyUpdateParamsAction = "log"
)
