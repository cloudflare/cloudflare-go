// File generated from our OpenAPI spec by Stainless.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
func (r *RuleService) New(ctx context.Context, zoneIdentifier string, body RuleNewParams, opts ...option.RequestOption) (res *[]RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing firewall rule.
func (r *RuleService) Update(ctx context.Context, zoneIdentifier string, id string, body RuleUpdateParams, opts ...option.RequestOption) (res *RuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *RuleService) List(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *RuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing firewall rule.
func (r *RuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of an existing firewall rule.
func (r *RuleService) Edit(ctx context.Context, zoneIdentifier string, id string, body RuleEditParams, opts ...option.RequestOption) (res *[]RuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a firewall rule.
func (r *RuleService) Get(ctx context.Context, zoneIdentifier string, id string, query RuleGetParams, opts ...option.RequestOption) (res *RuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleNewResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleNewResponseAction `json:"action,required"`
	Filter RuleNewResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                  `json:"priority"`
	Products []RuleNewResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string              `json:"ref"`
	JSON ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleNewResponseAction string

const (
	RuleNewResponseActionBlock            RuleNewResponseAction = "block"
	RuleNewResponseActionChallenge        RuleNewResponseAction = "challenge"
	RuleNewResponseActionJsChallenge      RuleNewResponseAction = "js_challenge"
	RuleNewResponseActionManagedChallenge RuleNewResponseAction = "managed_challenge"
	RuleNewResponseActionAllow            RuleNewResponseAction = "allow"
	RuleNewResponseActionLog              RuleNewResponseAction = "log"
	RuleNewResponseActionBypass           RuleNewResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleNewResponseFilterLegacyJhsFilter] or
// [firewall.RuleNewResponseFilterLegacyJhsDeletedFilter].
type RuleNewResponseFilter interface {
	implementsFirewallRuleNewResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleNewResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleNewResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleNewResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                   `json:"ref"`
	JSON ruleNewResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleNewResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleNewResponseFilterLegacyJhsFilter]
type ruleNewResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseFilterLegacyJhsFilter) implementsFirewallRuleNewResponseFilter() {}

type RuleNewResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                            `json:"deleted,required"`
	JSON    ruleNewResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleNewResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata for
// the struct [RuleNewResponseFilterLegacyJhsDeletedFilter]
type ruleNewResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleNewResponseFilter() {}

// A list of products to bypass for a request when using the `bypass` action.
type RuleNewResponseProduct string

const (
	RuleNewResponseProductZoneLockdown  RuleNewResponseProduct = "zoneLockdown"
	RuleNewResponseProductUABlock       RuleNewResponseProduct = "uaBlock"
	RuleNewResponseProductBic           RuleNewResponseProduct = "bic"
	RuleNewResponseProductHot           RuleNewResponseProduct = "hot"
	RuleNewResponseProductSecurityLevel RuleNewResponseProduct = "securityLevel"
	RuleNewResponseProductRateLimit     RuleNewResponseProduct = "rateLimit"
	RuleNewResponseProductWAF           RuleNewResponseProduct = "waf"
)

type RuleUpdateResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleUpdateResponseAction `json:"action,required"`
	Filter RuleUpdateResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                     `json:"priority"`
	Products []RuleUpdateResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                 `json:"ref"`
	JSON ruleUpdateResponseJSON `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleUpdateResponseAction string

const (
	RuleUpdateResponseActionBlock            RuleUpdateResponseAction = "block"
	RuleUpdateResponseActionChallenge        RuleUpdateResponseAction = "challenge"
	RuleUpdateResponseActionJsChallenge      RuleUpdateResponseAction = "js_challenge"
	RuleUpdateResponseActionManagedChallenge RuleUpdateResponseAction = "managed_challenge"
	RuleUpdateResponseActionAllow            RuleUpdateResponseAction = "allow"
	RuleUpdateResponseActionLog              RuleUpdateResponseAction = "log"
	RuleUpdateResponseActionBypass           RuleUpdateResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleUpdateResponseFilterLegacyJhsFilter] or
// [firewall.RuleUpdateResponseFilterLegacyJhsDeletedFilter].
type RuleUpdateResponseFilter interface {
	implementsFirewallRuleUpdateResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleUpdateResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleUpdateResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleUpdateResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleUpdateResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                      `json:"ref"`
	JSON ruleUpdateResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleUpdateResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleUpdateResponseFilterLegacyJhsFilter]
type ruleUpdateResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleUpdateResponseFilterLegacyJhsFilter) implementsFirewallRuleUpdateResponseFilter() {}

type RuleUpdateResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                               `json:"deleted,required"`
	JSON    ruleUpdateResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleUpdateResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata
// for the struct [RuleUpdateResponseFilterLegacyJhsDeletedFilter]
type ruleUpdateResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleUpdateResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleUpdateResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type RuleUpdateResponseProduct string

const (
	RuleUpdateResponseProductZoneLockdown  RuleUpdateResponseProduct = "zoneLockdown"
	RuleUpdateResponseProductUABlock       RuleUpdateResponseProduct = "uaBlock"
	RuleUpdateResponseProductBic           RuleUpdateResponseProduct = "bic"
	RuleUpdateResponseProductHot           RuleUpdateResponseProduct = "hot"
	RuleUpdateResponseProductSecurityLevel RuleUpdateResponseProduct = "securityLevel"
	RuleUpdateResponseProductRateLimit     RuleUpdateResponseProduct = "rateLimit"
	RuleUpdateResponseProductWAF           RuleUpdateResponseProduct = "waf"
)

type RuleListResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleListResponseAction `json:"action,required"`
	Filter RuleListResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                   `json:"priority"`
	Products []RuleListResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string               `json:"ref"`
	JSON ruleListResponseJSON `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleListResponseAction string

const (
	RuleListResponseActionBlock            RuleListResponseAction = "block"
	RuleListResponseActionChallenge        RuleListResponseAction = "challenge"
	RuleListResponseActionJsChallenge      RuleListResponseAction = "js_challenge"
	RuleListResponseActionManagedChallenge RuleListResponseAction = "managed_challenge"
	RuleListResponseActionAllow            RuleListResponseAction = "allow"
	RuleListResponseActionLog              RuleListResponseAction = "log"
	RuleListResponseActionBypass           RuleListResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleListResponseFilterLegacyJhsFilter] or
// [firewall.RuleListResponseFilterLegacyJhsDeletedFilter].
type RuleListResponseFilter interface {
	implementsFirewallRuleListResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleListResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleListResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                    `json:"ref"`
	JSON ruleListResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleListResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleListResponseFilterLegacyJhsFilter]
type ruleListResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleListResponseFilterLegacyJhsFilter) implementsFirewallRuleListResponseFilter() {}

type RuleListResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                             `json:"deleted,required"`
	JSON    ruleListResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleListResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata for
// the struct [RuleListResponseFilterLegacyJhsDeletedFilter]
type ruleListResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleListResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleListResponseFilter() {}

// A list of products to bypass for a request when using the `bypass` action.
type RuleListResponseProduct string

const (
	RuleListResponseProductZoneLockdown  RuleListResponseProduct = "zoneLockdown"
	RuleListResponseProductUABlock       RuleListResponseProduct = "uaBlock"
	RuleListResponseProductBic           RuleListResponseProduct = "bic"
	RuleListResponseProductHot           RuleListResponseProduct = "hot"
	RuleListResponseProductSecurityLevel RuleListResponseProduct = "securityLevel"
	RuleListResponseProductRateLimit     RuleListResponseProduct = "rateLimit"
	RuleListResponseProductWAF           RuleListResponseProduct = "waf"
)

type RuleDeleteResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleDeleteResponseAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                   `json:"description"`
	Filter      RuleDeleteResponseFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                     `json:"priority"`
	Products []RuleDeleteResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                 `json:"ref"`
	JSON ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
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

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleDeleteResponseAction string

const (
	RuleDeleteResponseActionBlock            RuleDeleteResponseAction = "block"
	RuleDeleteResponseActionChallenge        RuleDeleteResponseAction = "challenge"
	RuleDeleteResponseActionJsChallenge      RuleDeleteResponseAction = "js_challenge"
	RuleDeleteResponseActionManagedChallenge RuleDeleteResponseAction = "managed_challenge"
	RuleDeleteResponseActionAllow            RuleDeleteResponseAction = "allow"
	RuleDeleteResponseActionLog              RuleDeleteResponseAction = "log"
	RuleDeleteResponseActionBypass           RuleDeleteResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleDeleteResponseFilterLegacyJhsFilter] or
// [firewall.RuleDeleteResponseFilterLegacyJhsDeletedFilter].
type RuleDeleteResponseFilter interface {
	implementsFirewallRuleDeleteResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleDeleteResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleDeleteResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleDeleteResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                      `json:"ref"`
	JSON ruleDeleteResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleDeleteResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleDeleteResponseFilterLegacyJhsFilter]
type ruleDeleteResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseFilterLegacyJhsFilter) implementsFirewallRuleDeleteResponseFilter() {}

type RuleDeleteResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                               `json:"deleted,required"`
	JSON    ruleDeleteResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleDeleteResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata
// for the struct [RuleDeleteResponseFilterLegacyJhsDeletedFilter]
type ruleDeleteResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleDeleteResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type RuleDeleteResponseProduct string

const (
	RuleDeleteResponseProductZoneLockdown  RuleDeleteResponseProduct = "zoneLockdown"
	RuleDeleteResponseProductUABlock       RuleDeleteResponseProduct = "uaBlock"
	RuleDeleteResponseProductBic           RuleDeleteResponseProduct = "bic"
	RuleDeleteResponseProductHot           RuleDeleteResponseProduct = "hot"
	RuleDeleteResponseProductSecurityLevel RuleDeleteResponseProduct = "securityLevel"
	RuleDeleteResponseProductRateLimit     RuleDeleteResponseProduct = "rateLimit"
	RuleDeleteResponseProductWAF           RuleDeleteResponseProduct = "waf"
)

type RuleEditResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleEditResponseAction `json:"action,required"`
	Filter RuleEditResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                   `json:"priority"`
	Products []RuleEditResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string               `json:"ref"`
	JSON ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleEditResponseAction string

const (
	RuleEditResponseActionBlock            RuleEditResponseAction = "block"
	RuleEditResponseActionChallenge        RuleEditResponseAction = "challenge"
	RuleEditResponseActionJsChallenge      RuleEditResponseAction = "js_challenge"
	RuleEditResponseActionManagedChallenge RuleEditResponseAction = "managed_challenge"
	RuleEditResponseActionAllow            RuleEditResponseAction = "allow"
	RuleEditResponseActionLog              RuleEditResponseAction = "log"
	RuleEditResponseActionBypass           RuleEditResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleEditResponseFilterLegacyJhsFilter] or
// [firewall.RuleEditResponseFilterLegacyJhsDeletedFilter].
type RuleEditResponseFilter interface {
	implementsFirewallRuleEditResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleEditResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleEditResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleEditResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                    `json:"ref"`
	JSON ruleEditResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleEditResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleEditResponseFilterLegacyJhsFilter]
type ruleEditResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseFilterLegacyJhsFilter) implementsFirewallRuleEditResponseFilter() {}

type RuleEditResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                             `json:"deleted,required"`
	JSON    ruleEditResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleEditResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata for
// the struct [RuleEditResponseFilterLegacyJhsDeletedFilter]
type ruleEditResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleEditResponseFilter() {}

// A list of products to bypass for a request when using the `bypass` action.
type RuleEditResponseProduct string

const (
	RuleEditResponseProductZoneLockdown  RuleEditResponseProduct = "zoneLockdown"
	RuleEditResponseProductUABlock       RuleEditResponseProduct = "uaBlock"
	RuleEditResponseProductBic           RuleEditResponseProduct = "bic"
	RuleEditResponseProductHot           RuleEditResponseProduct = "hot"
	RuleEditResponseProductSecurityLevel RuleEditResponseProduct = "securityLevel"
	RuleEditResponseProductRateLimit     RuleEditResponseProduct = "rateLimit"
	RuleEditResponseProductWAF           RuleEditResponseProduct = "waf"
)

type RuleGetResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action RuleGetResponseAction `json:"action,required"`
	Filter RuleGetResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                  `json:"priority"`
	Products []RuleGetResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string              `json:"ref"`
	JSON ruleGetResponseJSON `json:"-"`
}

// ruleGetResponseJSON contains the JSON metadata for the struct [RuleGetResponse]
type ruleGetResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type RuleGetResponseAction string

const (
	RuleGetResponseActionBlock            RuleGetResponseAction = "block"
	RuleGetResponseActionChallenge        RuleGetResponseAction = "challenge"
	RuleGetResponseActionJsChallenge      RuleGetResponseAction = "js_challenge"
	RuleGetResponseActionManagedChallenge RuleGetResponseAction = "managed_challenge"
	RuleGetResponseActionAllow            RuleGetResponseAction = "allow"
	RuleGetResponseActionLog              RuleGetResponseAction = "log"
	RuleGetResponseActionBypass           RuleGetResponseAction = "bypass"
)

// Union satisfied by [firewall.RuleGetResponseFilterLegacyJhsFilter] or
// [firewall.RuleGetResponseFilterLegacyJhsDeletedFilter].
type RuleGetResponseFilter interface {
	implementsFirewallRuleGetResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleGetResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleGetResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleGetResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type RuleGetResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                   `json:"ref"`
	JSON ruleGetResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// ruleGetResponseFilterLegacyJhsFilterJSON contains the JSON metadata for the
// struct [RuleGetResponseFilterLegacyJhsFilter]
type ruleGetResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleGetResponseFilterLegacyJhsFilter) implementsFirewallRuleGetResponseFilter() {}

type RuleGetResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                            `json:"deleted,required"`
	JSON    ruleGetResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// ruleGetResponseFilterLegacyJhsDeletedFilterJSON contains the JSON metadata for
// the struct [RuleGetResponseFilterLegacyJhsDeletedFilter]
type ruleGetResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleGetResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleGetResponseFilter() {}

// A list of products to bypass for a request when using the `bypass` action.
type RuleGetResponseProduct string

const (
	RuleGetResponseProductZoneLockdown  RuleGetResponseProduct = "zoneLockdown"
	RuleGetResponseProductUABlock       RuleGetResponseProduct = "uaBlock"
	RuleGetResponseProductBic           RuleGetResponseProduct = "bic"
	RuleGetResponseProductHot           RuleGetResponseProduct = "hot"
	RuleGetResponseProductSecurityLevel RuleGetResponseProduct = "securityLevel"
	RuleGetResponseProductRateLimit     RuleGetResponseProduct = "rateLimit"
	RuleGetResponseProductWAF           RuleGetResponseProduct = "waf"
)

type RuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleNewResponseEnvelope struct {
	Errors   []RuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleNewResponseEnvelopeJSON       `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ruleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    ruleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

type RuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       ruleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeResultInfo]
type ruleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []RuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeErrors]
type ruleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeMessages]
type ruleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

type RuleListParams struct {
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

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r RuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleDeleteResponseEnvelope struct {
	Errors   []RuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

type RuleEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleEditResponseEnvelope struct {
	Errors   []RuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []RuleEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleEditResponseEnvelopeJSON       `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    ruleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleEditResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

type RuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       ruleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeResultInfo]
type ruleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleGetParams struct {
}

type RuleGetResponseEnvelope struct {
	Errors   []RuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ruleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelopeErrors]
type ruleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    ruleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelopeMessages]
type ruleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)
