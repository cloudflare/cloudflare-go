// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallRuleService] method
// instead.
type FirewallRuleService struct {
	Options []option.RequestOption
}

// NewFirewallRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallRuleService(opts ...option.RequestOption) (r *FirewallRuleService) {
	r = &FirewallRuleService{}
	r.Options = opts
	return
}

// Updates the priority of an existing firewall rule.
func (r *FirewallRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleUpdateParams, opts ...option.RequestOption) (res *FirewallRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes existing firewall rules.
func (r *FirewallRuleService) Delete(ctx context.Context, zoneIdentifier string, body FirewallRuleDeleteParams, opts ...option.RequestOption) (res *FirewallRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type FirewallRuleUpdateResponse struct {
	Errors     []FirewallRuleUpdateResponseError    `json:"errors"`
	Messages   []FirewallRuleUpdateResponseMessage  `json:"messages"`
	Result     []FirewallRuleUpdateResponseResult   `json:"result"`
	ResultInfo FirewallRuleUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FirewallRuleUpdateResponseSuccess `json:"success"`
	JSON    firewallRuleUpdateResponseJSON    `json:"-"`
}

// firewallRuleUpdateResponseJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponse]
type firewallRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    firewallRuleUpdateResponseErrorJSON `json:"-"`
}

// firewallRuleUpdateResponseErrorJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponseError]
type firewallRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    firewallRuleUpdateResponseMessageJSON `json:"-"`
}

// firewallRuleUpdateResponseMessageJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponseMessage]
type firewallRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleUpdateResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleUpdateResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                 `json:"description"`
	Filter      FirewallRuleUpdateResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                   `json:"priority"`
	Products []FirewallRuleUpdateResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                               `json:"ref"`
	JSON firewallRuleUpdateResponseResultJSON `json:"-"`
}

// firewallRuleUpdateResponseResultJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponseResult]
type firewallRuleUpdateResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleUpdateResponseResultAction string

const (
	FirewallRuleUpdateResponseResultActionBlock            FirewallRuleUpdateResponseResultAction = "block"
	FirewallRuleUpdateResponseResultActionChallenge        FirewallRuleUpdateResponseResultAction = "challenge"
	FirewallRuleUpdateResponseResultActionJsChallenge      FirewallRuleUpdateResponseResultAction = "js_challenge"
	FirewallRuleUpdateResponseResultActionManagedChallenge FirewallRuleUpdateResponseResultAction = "managed_challenge"
	FirewallRuleUpdateResponseResultActionAllow            FirewallRuleUpdateResponseResultAction = "allow"
	FirewallRuleUpdateResponseResultActionLog              FirewallRuleUpdateResponseResultAction = "log"
	FirewallRuleUpdateResponseResultActionBypass           FirewallRuleUpdateResponseResultAction = "bypass"
)

// Union satisfied by [FirewallRuleUpdateResponseResultFilterDZw70ubTFilter] or
// [FirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter].
type FirewallRuleUpdateResponseResultFilter interface {
	implementsFirewallRuleUpdateResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleUpdateResponseResultFilter)(nil)).Elem(), "")
}

type FirewallRuleUpdateResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                   `json:"ref"`
	JSON firewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// firewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON contains the JSON
// metadata for the struct [FirewallRuleUpdateResponseResultFilterDZw70ubTFilter]
type firewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleUpdateResponseResultFilterDZw70ubTFilter) implementsFirewallRuleUpdateResponseResultFilter() {
}

type FirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                            `json:"deleted,required"`
	JSON    firewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// firewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON contains the
// JSON metadata for the struct
// [FirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter]
type firewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter) implementsFirewallRuleUpdateResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleUpdateResponseResultProduct string

const (
	FirewallRuleUpdateResponseResultProductZoneLockdown  FirewallRuleUpdateResponseResultProduct = "zoneLockdown"
	FirewallRuleUpdateResponseResultProductUaBlock       FirewallRuleUpdateResponseResultProduct = "uaBlock"
	FirewallRuleUpdateResponseResultProductBic           FirewallRuleUpdateResponseResultProduct = "bic"
	FirewallRuleUpdateResponseResultProductHot           FirewallRuleUpdateResponseResultProduct = "hot"
	FirewallRuleUpdateResponseResultProductSecurityLevel FirewallRuleUpdateResponseResultProduct = "securityLevel"
	FirewallRuleUpdateResponseResultProductRateLimit     FirewallRuleUpdateResponseResultProduct = "rateLimit"
	FirewallRuleUpdateResponseResultProductWaf           FirewallRuleUpdateResponseResultProduct = "waf"
)

type FirewallRuleUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       firewallRuleUpdateResponseResultInfoJSON `json:"-"`
}

// firewallRuleUpdateResponseResultInfoJSON contains the JSON metadata for the
// struct [FirewallRuleUpdateResponseResultInfo]
type firewallRuleUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleUpdateResponseSuccess bool

const (
	FirewallRuleUpdateResponseSuccessTrue FirewallRuleUpdateResponseSuccess = true
)

type FirewallRuleDeleteResponse struct {
	Errors     []FirewallRuleDeleteResponseError    `json:"errors"`
	Messages   []FirewallRuleDeleteResponseMessage  `json:"messages"`
	Result     []FirewallRuleDeleteResponseResult   `json:"result"`
	ResultInfo FirewallRuleDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FirewallRuleDeleteResponseSuccess `json:"success"`
	JSON    firewallRuleDeleteResponseJSON    `json:"-"`
}

// firewallRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponse]
type firewallRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleDeleteResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    firewallRuleDeleteResponseErrorJSON `json:"-"`
}

// firewallRuleDeleteResponseErrorJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponseError]
type firewallRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleDeleteResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    firewallRuleDeleteResponseMessageJSON `json:"-"`
}

// firewallRuleDeleteResponseMessageJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponseMessage]
type firewallRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleDeleteResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleDeleteResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                 `json:"description"`
	Filter      FirewallRuleDeleteResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                   `json:"priority"`
	Products []FirewallRuleDeleteResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                               `json:"ref"`
	JSON firewallRuleDeleteResponseResultJSON `json:"-"`
}

// firewallRuleDeleteResponseResultJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponseResult]
type firewallRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleDeleteResponseResultAction string

const (
	FirewallRuleDeleteResponseResultActionBlock            FirewallRuleDeleteResponseResultAction = "block"
	FirewallRuleDeleteResponseResultActionChallenge        FirewallRuleDeleteResponseResultAction = "challenge"
	FirewallRuleDeleteResponseResultActionJsChallenge      FirewallRuleDeleteResponseResultAction = "js_challenge"
	FirewallRuleDeleteResponseResultActionManagedChallenge FirewallRuleDeleteResponseResultAction = "managed_challenge"
	FirewallRuleDeleteResponseResultActionAllow            FirewallRuleDeleteResponseResultAction = "allow"
	FirewallRuleDeleteResponseResultActionLog              FirewallRuleDeleteResponseResultAction = "log"
	FirewallRuleDeleteResponseResultActionBypass           FirewallRuleDeleteResponseResultAction = "bypass"
)

// Union satisfied by [FirewallRuleDeleteResponseResultFilterDZw70ubTFilter] or
// [FirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter].
type FirewallRuleDeleteResponseResultFilter interface {
	implementsFirewallRuleDeleteResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleDeleteResponseResultFilter)(nil)).Elem(), "")
}

type FirewallRuleDeleteResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                   `json:"ref"`
	JSON firewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// firewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON contains the JSON
// metadata for the struct [FirewallRuleDeleteResponseResultFilterDZw70ubTFilter]
type firewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleDeleteResponseResultFilterDZw70ubTFilter) implementsFirewallRuleDeleteResponseResultFilter() {
}

type FirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                            `json:"deleted,required"`
	JSON    firewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// firewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON contains the
// JSON metadata for the struct
// [FirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter]
type firewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter) implementsFirewallRuleDeleteResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleDeleteResponseResultProduct string

const (
	FirewallRuleDeleteResponseResultProductZoneLockdown  FirewallRuleDeleteResponseResultProduct = "zoneLockdown"
	FirewallRuleDeleteResponseResultProductUaBlock       FirewallRuleDeleteResponseResultProduct = "uaBlock"
	FirewallRuleDeleteResponseResultProductBic           FirewallRuleDeleteResponseResultProduct = "bic"
	FirewallRuleDeleteResponseResultProductHot           FirewallRuleDeleteResponseResultProduct = "hot"
	FirewallRuleDeleteResponseResultProductSecurityLevel FirewallRuleDeleteResponseResultProduct = "securityLevel"
	FirewallRuleDeleteResponseResultProductRateLimit     FirewallRuleDeleteResponseResultProduct = "rateLimit"
	FirewallRuleDeleteResponseResultProductWaf           FirewallRuleDeleteResponseResultProduct = "waf"
)

type FirewallRuleDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       firewallRuleDeleteResponseResultInfoJSON `json:"-"`
}

// firewallRuleDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [FirewallRuleDeleteResponseResultInfo]
type firewallRuleDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleDeleteResponseSuccess bool

const (
	FirewallRuleDeleteResponseSuccessTrue FirewallRuleDeleteResponseSuccess = true
)

type FirewallRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleDeleteParams struct {
}
