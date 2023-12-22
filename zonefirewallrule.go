// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallRuleService] method
// instead.
type ZoneFirewallRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallRuleService(opts ...option.RequestOption) (r *ZoneFirewallRuleService) {
	r = &ZoneFirewallRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a firewall rule.
func (r *ZoneFirewallRuleService) Get(ctx context.Context, zoneIdentifier string, id string, query ZoneFirewallRuleGetParams, opts ...option.RequestOption) (res *FilterRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates an existing firewall rule.
func (r *ZoneFirewallRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallRuleUpdateParams, opts ...option.RequestOption) (res *FilterRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing firewall rule.
func (r *ZoneFirewallRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallRuleDeleteParams, opts ...option.RequestOption) (res *FilterRulesSingleResponseDelete, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create one or more firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesNewFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesNewFirewallRulesParams, opts ...option.RequestOption) (res *FilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *ZoneFirewallRuleService) FirewallRulesListFirewallRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallRuleFirewallRulesListFirewallRulesParams, opts ...option.RequestOption) (res *FilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates one or more existing firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesUpdateFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams, opts ...option.RequestOption) (res *FilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates the priority of existing firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesUpdatePriorityOfFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams, opts ...option.RequestOption) (res *FilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type FilterRulesResponseCollection struct {
	Errors     []FilterRulesResponseCollectionError    `json:"errors"`
	Messages   []FilterRulesResponseCollectionMessage  `json:"messages"`
	Result     []FilterRulesResponseCollectionResult   `json:"result"`
	ResultInfo FilterRulesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FilterRulesResponseCollectionSuccess `json:"success"`
	JSON    filterRulesResponseCollectionJSON    `json:"-"`
}

// filterRulesResponseCollectionJSON contains the JSON metadata for the struct
// [FilterRulesResponseCollection]
type filterRulesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    filterRulesResponseCollectionErrorJSON `json:"-"`
}

// filterRulesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [FilterRulesResponseCollectionError]
type filterRulesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterRulesResponseCollectionMessageJSON `json:"-"`
}

// filterRulesResponseCollectionMessageJSON contains the JSON metadata for the
// struct [FilterRulesResponseCollectionMessage]
type filterRulesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesResponseCollectionResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FilterRulesResponseCollectionResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                    `json:"description"`
	Filter      FilterRulesResponseCollectionResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                      `json:"priority"`
	Products []FilterRulesResponseCollectionResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                  `json:"ref"`
	JSON filterRulesResponseCollectionResultJSON `json:"-"`
}

// filterRulesResponseCollectionResultJSON contains the JSON metadata for the
// struct [FilterRulesResponseCollectionResult]
type filterRulesResponseCollectionResultJSON struct {
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

func (r *FilterRulesResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FilterRulesResponseCollectionResultAction string

const (
	FilterRulesResponseCollectionResultActionBlock            FilterRulesResponseCollectionResultAction = "block"
	FilterRulesResponseCollectionResultActionChallenge        FilterRulesResponseCollectionResultAction = "challenge"
	FilterRulesResponseCollectionResultActionJsChallenge      FilterRulesResponseCollectionResultAction = "js_challenge"
	FilterRulesResponseCollectionResultActionManagedChallenge FilterRulesResponseCollectionResultAction = "managed_challenge"
	FilterRulesResponseCollectionResultActionAllow            FilterRulesResponseCollectionResultAction = "allow"
	FilterRulesResponseCollectionResultActionLog              FilterRulesResponseCollectionResultAction = "log"
	FilterRulesResponseCollectionResultActionBypass           FilterRulesResponseCollectionResultAction = "bypass"
)

// Union satisfied by [FilterRulesResponseCollectionResultFilterFilterJRjYt5St] or
// [FilterRulesResponseCollectionResultFilterDeletedFilter].
type FilterRulesResponseCollectionResultFilter interface {
	implementsFilterRulesResponseCollectionResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FilterRulesResponseCollectionResultFilter)(nil)).Elem(), "")
}

type FilterRulesResponseCollectionResultFilterFilterJRjYt5St struct {
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
	Ref  string                                                      `json:"ref"`
	JSON filterRulesResponseCollectionResultFilterFilterJRjYt5StJSON `json:"-"`
}

// filterRulesResponseCollectionResultFilterFilterJRjYt5StJSON contains the JSON
// metadata for the struct
// [FilterRulesResponseCollectionResultFilterFilterJRjYt5St]
type filterRulesResponseCollectionResultFilterFilterJRjYt5StJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollectionResultFilterFilterJRjYt5St) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesResponseCollectionResultFilterFilterJRjYt5St) implementsFilterRulesResponseCollectionResultFilter() {
}

type FilterRulesResponseCollectionResultFilterDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                       `json:"deleted,required"`
	JSON    filterRulesResponseCollectionResultFilterDeletedFilterJSON `json:"-"`
}

// filterRulesResponseCollectionResultFilterDeletedFilterJSON contains the JSON
// metadata for the struct [FilterRulesResponseCollectionResultFilterDeletedFilter]
type filterRulesResponseCollectionResultFilterDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollectionResultFilterDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesResponseCollectionResultFilterDeletedFilter) implementsFilterRulesResponseCollectionResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FilterRulesResponseCollectionResultProduct string

const (
	FilterRulesResponseCollectionResultProductZoneLockdown  FilterRulesResponseCollectionResultProduct = "zoneLockdown"
	FilterRulesResponseCollectionResultProductUaBlock       FilterRulesResponseCollectionResultProduct = "uaBlock"
	FilterRulesResponseCollectionResultProductBic           FilterRulesResponseCollectionResultProduct = "bic"
	FilterRulesResponseCollectionResultProductHot           FilterRulesResponseCollectionResultProduct = "hot"
	FilterRulesResponseCollectionResultProductSecurityLevel FilterRulesResponseCollectionResultProduct = "securityLevel"
	FilterRulesResponseCollectionResultProductRateLimit     FilterRulesResponseCollectionResultProduct = "rateLimit"
	FilterRulesResponseCollectionResultProductWaf           FilterRulesResponseCollectionResultProduct = "waf"
)

type FilterRulesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       filterRulesResponseCollectionResultInfoJSON `json:"-"`
}

// filterRulesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [FilterRulesResponseCollectionResultInfo]
type filterRulesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterRulesResponseCollectionSuccess bool

const (
	FilterRulesResponseCollectionSuccessTrue FilterRulesResponseCollectionSuccess = true
)

type FilterRulesSingleResponse struct {
	Errors   []FilterRulesSingleResponseError   `json:"errors"`
	Messages []FilterRulesSingleResponseMessage `json:"messages"`
	Result   FilterRulesSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success FilterRulesSingleResponseSuccess `json:"success"`
	JSON    filterRulesSingleResponseJSON    `json:"-"`
}

// filterRulesSingleResponseJSON contains the JSON metadata for the struct
// [FilterRulesSingleResponse]
type filterRulesSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    filterRulesSingleResponseErrorJSON `json:"-"`
}

// filterRulesSingleResponseErrorJSON contains the JSON metadata for the struct
// [FilterRulesSingleResponseError]
type filterRulesSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    filterRulesSingleResponseMessageJSON `json:"-"`
}

// filterRulesSingleResponseMessageJSON contains the JSON metadata for the struct
// [FilterRulesSingleResponseMessage]
type filterRulesSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FilterRulesSingleResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                `json:"description"`
	Filter      FilterRulesSingleResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                  `json:"priority"`
	Products []FilterRulesSingleResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                              `json:"ref"`
	JSON filterRulesSingleResponseResultJSON `json:"-"`
}

// filterRulesSingleResponseResultJSON contains the JSON metadata for the struct
// [FilterRulesSingleResponseResult]
type filterRulesSingleResponseResultJSON struct {
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

func (r *FilterRulesSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FilterRulesSingleResponseResultAction string

const (
	FilterRulesSingleResponseResultActionBlock            FilterRulesSingleResponseResultAction = "block"
	FilterRulesSingleResponseResultActionChallenge        FilterRulesSingleResponseResultAction = "challenge"
	FilterRulesSingleResponseResultActionJsChallenge      FilterRulesSingleResponseResultAction = "js_challenge"
	FilterRulesSingleResponseResultActionManagedChallenge FilterRulesSingleResponseResultAction = "managed_challenge"
	FilterRulesSingleResponseResultActionAllow            FilterRulesSingleResponseResultAction = "allow"
	FilterRulesSingleResponseResultActionLog              FilterRulesSingleResponseResultAction = "log"
	FilterRulesSingleResponseResultActionBypass           FilterRulesSingleResponseResultAction = "bypass"
)

// Union satisfied by [FilterRulesSingleResponseResultFilterFilterJRjYt5St] or
// [FilterRulesSingleResponseResultFilterDeletedFilter].
type FilterRulesSingleResponseResultFilter interface {
	implementsFilterRulesSingleResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FilterRulesSingleResponseResultFilter)(nil)).Elem(), "")
}

type FilterRulesSingleResponseResultFilterFilterJRjYt5St struct {
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
	Ref  string                                                  `json:"ref"`
	JSON filterRulesSingleResponseResultFilterFilterJRjYt5StJSON `json:"-"`
}

// filterRulesSingleResponseResultFilterFilterJRjYt5StJSON contains the JSON
// metadata for the struct [FilterRulesSingleResponseResultFilterFilterJRjYt5St]
type filterRulesSingleResponseResultFilterFilterJRjYt5StJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseResultFilterFilterJRjYt5St) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesSingleResponseResultFilterFilterJRjYt5St) implementsFilterRulesSingleResponseResultFilter() {
}

type FilterRulesSingleResponseResultFilterDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                   `json:"deleted,required"`
	JSON    filterRulesSingleResponseResultFilterDeletedFilterJSON `json:"-"`
}

// filterRulesSingleResponseResultFilterDeletedFilterJSON contains the JSON
// metadata for the struct [FilterRulesSingleResponseResultFilterDeletedFilter]
type filterRulesSingleResponseResultFilterDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseResultFilterDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesSingleResponseResultFilterDeletedFilter) implementsFilterRulesSingleResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FilterRulesSingleResponseResultProduct string

const (
	FilterRulesSingleResponseResultProductZoneLockdown  FilterRulesSingleResponseResultProduct = "zoneLockdown"
	FilterRulesSingleResponseResultProductUaBlock       FilterRulesSingleResponseResultProduct = "uaBlock"
	FilterRulesSingleResponseResultProductBic           FilterRulesSingleResponseResultProduct = "bic"
	FilterRulesSingleResponseResultProductHot           FilterRulesSingleResponseResultProduct = "hot"
	FilterRulesSingleResponseResultProductSecurityLevel FilterRulesSingleResponseResultProduct = "securityLevel"
	FilterRulesSingleResponseResultProductRateLimit     FilterRulesSingleResponseResultProduct = "rateLimit"
	FilterRulesSingleResponseResultProductWaf           FilterRulesSingleResponseResultProduct = "waf"
)

// Whether the API call was successful
type FilterRulesSingleResponseSuccess bool

const (
	FilterRulesSingleResponseSuccessTrue FilterRulesSingleResponseSuccess = true
)

type FilterRulesSingleResponseDelete struct {
	Errors   []FilterRulesSingleResponseDeleteError   `json:"errors"`
	Messages []FilterRulesSingleResponseDeleteMessage `json:"messages"`
	Result   FilterRulesSingleResponseDeleteResult    `json:"result"`
	// Whether the API call was successful
	Success FilterRulesSingleResponseDeleteSuccess `json:"success"`
	JSON    filterRulesSingleResponseDeleteJSON    `json:"-"`
}

// filterRulesSingleResponseDeleteJSON contains the JSON metadata for the struct
// [FilterRulesSingleResponseDelete]
type filterRulesSingleResponseDeleteJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseDeleteError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterRulesSingleResponseDeleteErrorJSON `json:"-"`
}

// filterRulesSingleResponseDeleteErrorJSON contains the JSON metadata for the
// struct [FilterRulesSingleResponseDeleteError]
type filterRulesSingleResponseDeleteErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseDeleteError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseDeleteMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    filterRulesSingleResponseDeleteMessageJSON `json:"-"`
}

// filterRulesSingleResponseDeleteMessageJSON contains the JSON metadata for the
// struct [FilterRulesSingleResponseDeleteMessage]
type filterRulesSingleResponseDeleteMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseDeleteMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterRulesSingleResponseDeleteResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FilterRulesSingleResponseDeleteResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                      `json:"description"`
	Filter      FilterRulesSingleResponseDeleteResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                        `json:"priority"`
	Products []FilterRulesSingleResponseDeleteResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                    `json:"ref"`
	JSON filterRulesSingleResponseDeleteResultJSON `json:"-"`
}

// filterRulesSingleResponseDeleteResultJSON contains the JSON metadata for the
// struct [FilterRulesSingleResponseDeleteResult]
type filterRulesSingleResponseDeleteResultJSON struct {
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

func (r *FilterRulesSingleResponseDeleteResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FilterRulesSingleResponseDeleteResultAction string

const (
	FilterRulesSingleResponseDeleteResultActionBlock            FilterRulesSingleResponseDeleteResultAction = "block"
	FilterRulesSingleResponseDeleteResultActionChallenge        FilterRulesSingleResponseDeleteResultAction = "challenge"
	FilterRulesSingleResponseDeleteResultActionJsChallenge      FilterRulesSingleResponseDeleteResultAction = "js_challenge"
	FilterRulesSingleResponseDeleteResultActionManagedChallenge FilterRulesSingleResponseDeleteResultAction = "managed_challenge"
	FilterRulesSingleResponseDeleteResultActionAllow            FilterRulesSingleResponseDeleteResultAction = "allow"
	FilterRulesSingleResponseDeleteResultActionLog              FilterRulesSingleResponseDeleteResultAction = "log"
	FilterRulesSingleResponseDeleteResultActionBypass           FilterRulesSingleResponseDeleteResultAction = "bypass"
)

// Union satisfied by [FilterRulesSingleResponseDeleteResultFilterFilterJRjYt5St]
// or [FilterRulesSingleResponseDeleteResultFilterDeletedFilter].
type FilterRulesSingleResponseDeleteResultFilter interface {
	implementsFilterRulesSingleResponseDeleteResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FilterRulesSingleResponseDeleteResultFilter)(nil)).Elem(), "")
}

type FilterRulesSingleResponseDeleteResultFilterFilterJRjYt5St struct {
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
	Ref  string                                                        `json:"ref"`
	JSON filterRulesSingleResponseDeleteResultFilterFilterJRjYt5StJSON `json:"-"`
}

// filterRulesSingleResponseDeleteResultFilterFilterJRjYt5StJSON contains the JSON
// metadata for the struct
// [FilterRulesSingleResponseDeleteResultFilterFilterJRjYt5St]
type filterRulesSingleResponseDeleteResultFilterFilterJRjYt5StJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseDeleteResultFilterFilterJRjYt5St) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesSingleResponseDeleteResultFilterFilterJRjYt5St) implementsFilterRulesSingleResponseDeleteResultFilter() {
}

type FilterRulesSingleResponseDeleteResultFilterDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                         `json:"deleted,required"`
	JSON    filterRulesSingleResponseDeleteResultFilterDeletedFilterJSON `json:"-"`
}

// filterRulesSingleResponseDeleteResultFilterDeletedFilterJSON contains the JSON
// metadata for the struct
// [FilterRulesSingleResponseDeleteResultFilterDeletedFilter]
type filterRulesSingleResponseDeleteResultFilterDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterRulesSingleResponseDeleteResultFilterDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FilterRulesSingleResponseDeleteResultFilterDeletedFilter) implementsFilterRulesSingleResponseDeleteResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FilterRulesSingleResponseDeleteResultProduct string

const (
	FilterRulesSingleResponseDeleteResultProductZoneLockdown  FilterRulesSingleResponseDeleteResultProduct = "zoneLockdown"
	FilterRulesSingleResponseDeleteResultProductUaBlock       FilterRulesSingleResponseDeleteResultProduct = "uaBlock"
	FilterRulesSingleResponseDeleteResultProductBic           FilterRulesSingleResponseDeleteResultProduct = "bic"
	FilterRulesSingleResponseDeleteResultProductHot           FilterRulesSingleResponseDeleteResultProduct = "hot"
	FilterRulesSingleResponseDeleteResultProductSecurityLevel FilterRulesSingleResponseDeleteResultProduct = "securityLevel"
	FilterRulesSingleResponseDeleteResultProductRateLimit     FilterRulesSingleResponseDeleteResultProduct = "rateLimit"
	FilterRulesSingleResponseDeleteResultProductWaf           FilterRulesSingleResponseDeleteResultProduct = "waf"
)

// Whether the API call was successful
type FilterRulesSingleResponseDeleteSuccess bool

const (
	FilterRulesSingleResponseDeleteSuccessTrue FilterRulesSingleResponseDeleteSuccess = true
)

type ZoneFirewallRuleGetParams struct {
}

type ZoneFirewallRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r ZoneFirewallRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesNewFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleFirewallRulesListFirewallRulesParams struct {
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallRuleFirewallRulesListFirewallRulesParams]'s
// query parameters as `url.Values`.
func (r ZoneFirewallRuleFirewallRulesListFirewallRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
