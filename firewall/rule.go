// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *RuleService) List(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RuleListResponse], err error) {
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
func (r *RuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
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
func (r *RuleService) Get(ctx context.Context, zoneIdentifier string, params RuleGetParams, opts ...option.RequestOption) (res *RuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, params.PathID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
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

func (r RuleNewResponseAction) IsKnown() bool {
	switch r {
	case RuleNewResponseActionBlock, RuleNewResponseActionChallenge, RuleNewResponseActionJsChallenge, RuleNewResponseActionManagedChallenge, RuleNewResponseActionAllow, RuleNewResponseActionLog, RuleNewResponseActionBypass:
		return true
	}
	return false
}

type RuleNewResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                      `json:"deleted"`
	JSON    ruleNewResponseFilterJSON `json:"-"`
	union   RuleNewResponseFilterUnion
}

// ruleNewResponseFilterJSON contains the JSON metadata for the struct
// [RuleNewResponseFilter]
type ruleNewResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleNewResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleNewResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleNewResponseFilter) AsUnion() RuleNewResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleNewResponseFilterFirewallFilter] or
// [firewall.RuleNewResponseFilterFirewallDeletedFilter].
type RuleNewResponseFilterUnion interface {
	implementsFirewallRuleNewResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleNewResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleNewResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleNewResponseFilterFirewallFilter struct {
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
	Ref  string                                  `json:"ref"`
	JSON ruleNewResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleNewResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleNewResponseFilterFirewallFilter]
type ruleNewResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseFilterFirewallFilter) implementsFirewallRuleNewResponseFilter() {}

type RuleNewResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                           `json:"deleted,required"`
	JSON    ruleNewResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleNewResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleNewResponseFilterFirewallDeletedFilter]
type ruleNewResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseFilterFirewallDeletedFilter) implementsFirewallRuleNewResponseFilter() {}

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

func (r RuleNewResponseProduct) IsKnown() bool {
	switch r {
	case RuleNewResponseProductZoneLockdown, RuleNewResponseProductUABlock, RuleNewResponseProductBic, RuleNewResponseProductHot, RuleNewResponseProductSecurityLevel, RuleNewResponseProductRateLimit, RuleNewResponseProductWAF:
		return true
	}
	return false
}

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

func (r RuleUpdateResponseAction) IsKnown() bool {
	switch r {
	case RuleUpdateResponseActionBlock, RuleUpdateResponseActionChallenge, RuleUpdateResponseActionJsChallenge, RuleUpdateResponseActionManagedChallenge, RuleUpdateResponseActionAllow, RuleUpdateResponseActionLog, RuleUpdateResponseActionBypass:
		return true
	}
	return false
}

type RuleUpdateResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                         `json:"deleted"`
	JSON    ruleUpdateResponseFilterJSON `json:"-"`
	union   RuleUpdateResponseFilterUnion
}

// ruleUpdateResponseFilterJSON contains the JSON metadata for the struct
// [RuleUpdateResponseFilter]
type ruleUpdateResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleUpdateResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleUpdateResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleUpdateResponseFilter) AsUnion() RuleUpdateResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleUpdateResponseFilterFirewallFilter] or
// [firewall.RuleUpdateResponseFilterFirewallDeletedFilter].
type RuleUpdateResponseFilterUnion interface {
	implementsFirewallRuleUpdateResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleUpdateResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleUpdateResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleUpdateResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleUpdateResponseFilterFirewallFilter struct {
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
	Ref  string                                     `json:"ref"`
	JSON ruleUpdateResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleUpdateResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleUpdateResponseFilterFirewallFilter]
type ruleUpdateResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleUpdateResponseFilterFirewallFilter) implementsFirewallRuleUpdateResponseFilter() {}

type RuleUpdateResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                              `json:"deleted,required"`
	JSON    ruleUpdateResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleUpdateResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleUpdateResponseFilterFirewallDeletedFilter]
type ruleUpdateResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleUpdateResponseFilterFirewallDeletedFilter) implementsFirewallRuleUpdateResponseFilter() {}

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

func (r RuleUpdateResponseProduct) IsKnown() bool {
	switch r {
	case RuleUpdateResponseProductZoneLockdown, RuleUpdateResponseProductUABlock, RuleUpdateResponseProductBic, RuleUpdateResponseProductHot, RuleUpdateResponseProductSecurityLevel, RuleUpdateResponseProductRateLimit, RuleUpdateResponseProductWAF:
		return true
	}
	return false
}

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

func (r RuleListResponseAction) IsKnown() bool {
	switch r {
	case RuleListResponseActionBlock, RuleListResponseActionChallenge, RuleListResponseActionJsChallenge, RuleListResponseActionManagedChallenge, RuleListResponseActionAllow, RuleListResponseActionLog, RuleListResponseActionBypass:
		return true
	}
	return false
}

type RuleListResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                       `json:"deleted"`
	JSON    ruleListResponseFilterJSON `json:"-"`
	union   RuleListResponseFilterUnion
}

// ruleListResponseFilterJSON contains the JSON metadata for the struct
// [RuleListResponseFilter]
type ruleListResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleListResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleListResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleListResponseFilter) AsUnion() RuleListResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleListResponseFilterFirewallFilter] or
// [firewall.RuleListResponseFilterFirewallDeletedFilter].
type RuleListResponseFilterUnion interface {
	implementsFirewallRuleListResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleListResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleListResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleListResponseFilterFirewallFilter struct {
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
	JSON ruleListResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleListResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleListResponseFilterFirewallFilter]
type ruleListResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleListResponseFilterFirewallFilter) implementsFirewallRuleListResponseFilter() {}

type RuleListResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                            `json:"deleted,required"`
	JSON    ruleListResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleListResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleListResponseFilterFirewallDeletedFilter]
type ruleListResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleListResponseFilterFirewallDeletedFilter) implementsFirewallRuleListResponseFilter() {}

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

func (r RuleListResponseProduct) IsKnown() bool {
	switch r {
	case RuleListResponseProductZoneLockdown, RuleListResponseProductUABlock, RuleListResponseProductBic, RuleListResponseProductHot, RuleListResponseProductSecurityLevel, RuleListResponseProductRateLimit, RuleListResponseProductWAF:
		return true
	}
	return false
}

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

func (r RuleDeleteResponseAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseActionBlock, RuleDeleteResponseActionChallenge, RuleDeleteResponseActionJsChallenge, RuleDeleteResponseActionManagedChallenge, RuleDeleteResponseActionAllow, RuleDeleteResponseActionLog, RuleDeleteResponseActionBypass:
		return true
	}
	return false
}

type RuleDeleteResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                         `json:"deleted"`
	JSON    ruleDeleteResponseFilterJSON `json:"-"`
	union   RuleDeleteResponseFilterUnion
}

// ruleDeleteResponseFilterJSON contains the JSON metadata for the struct
// [RuleDeleteResponseFilter]
type ruleDeleteResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleDeleteResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleDeleteResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleDeleteResponseFilter) AsUnion() RuleDeleteResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleDeleteResponseFilterFirewallFilter] or
// [firewall.RuleDeleteResponseFilterFirewallDeletedFilter].
type RuleDeleteResponseFilterUnion interface {
	implementsFirewallRuleDeleteResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleDeleteResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleDeleteResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleDeleteResponseFilterFirewallFilter struct {
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
	Ref  string                                     `json:"ref"`
	JSON ruleDeleteResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleDeleteResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleDeleteResponseFilterFirewallFilter]
type ruleDeleteResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseFilterFirewallFilter) implementsFirewallRuleDeleteResponseFilter() {}

type RuleDeleteResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                              `json:"deleted,required"`
	JSON    ruleDeleteResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleDeleteResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleDeleteResponseFilterFirewallDeletedFilter]
type ruleDeleteResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseFilterFirewallDeletedFilter) implementsFirewallRuleDeleteResponseFilter() {}

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

func (r RuleDeleteResponseProduct) IsKnown() bool {
	switch r {
	case RuleDeleteResponseProductZoneLockdown, RuleDeleteResponseProductUABlock, RuleDeleteResponseProductBic, RuleDeleteResponseProductHot, RuleDeleteResponseProductSecurityLevel, RuleDeleteResponseProductRateLimit, RuleDeleteResponseProductWAF:
		return true
	}
	return false
}

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

func (r RuleEditResponseAction) IsKnown() bool {
	switch r {
	case RuleEditResponseActionBlock, RuleEditResponseActionChallenge, RuleEditResponseActionJsChallenge, RuleEditResponseActionManagedChallenge, RuleEditResponseActionAllow, RuleEditResponseActionLog, RuleEditResponseActionBypass:
		return true
	}
	return false
}

type RuleEditResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                       `json:"deleted"`
	JSON    ruleEditResponseFilterJSON `json:"-"`
	union   RuleEditResponseFilterUnion
}

// ruleEditResponseFilterJSON contains the JSON metadata for the struct
// [RuleEditResponseFilter]
type ruleEditResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleEditResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleEditResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleEditResponseFilter) AsUnion() RuleEditResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleEditResponseFilterFirewallFilter] or
// [firewall.RuleEditResponseFilterFirewallDeletedFilter].
type RuleEditResponseFilterUnion interface {
	implementsFirewallRuleEditResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleEditResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleEditResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleEditResponseFilterFirewallFilter struct {
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
	JSON ruleEditResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleEditResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleEditResponseFilterFirewallFilter]
type ruleEditResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseFilterFirewallFilter) implementsFirewallRuleEditResponseFilter() {}

type RuleEditResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                            `json:"deleted,required"`
	JSON    ruleEditResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleEditResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleEditResponseFilterFirewallDeletedFilter]
type ruleEditResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseFilterFirewallDeletedFilter) implementsFirewallRuleEditResponseFilter() {}

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

func (r RuleEditResponseProduct) IsKnown() bool {
	switch r {
	case RuleEditResponseProductZoneLockdown, RuleEditResponseProductUABlock, RuleEditResponseProductBic, RuleEditResponseProductHot, RuleEditResponseProductSecurityLevel, RuleEditResponseProductRateLimit, RuleEditResponseProductWAF:
		return true
	}
	return false
}

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

func (r RuleGetResponseAction) IsKnown() bool {
	switch r {
	case RuleGetResponseActionBlock, RuleGetResponseActionChallenge, RuleGetResponseActionJsChallenge, RuleGetResponseActionManagedChallenge, RuleGetResponseActionAllow, RuleGetResponseActionLog, RuleGetResponseActionBypass:
		return true
	}
	return false
}

type RuleGetResponseFilter struct {
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref string `json:"ref"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                      `json:"deleted"`
	JSON    ruleGetResponseFilterJSON `json:"-"`
	union   RuleGetResponseFilterUnion
}

// ruleGetResponseFilterJSON contains the JSON metadata for the struct
// [RuleGetResponseFilter]
type ruleGetResponseFilterJSON struct {
	Description apijson.Field
	Expression  apijson.Field
	ID          apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r ruleGetResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *RuleGetResponseFilter) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleGetResponseFilter) AsUnion() RuleGetResponseFilterUnion {
	return r.union
}

// Union satisfied by [firewall.RuleGetResponseFilterFirewallFilter] or
// [firewall.RuleGetResponseFilterFirewallDeletedFilter].
type RuleGetResponseFilterUnion interface {
	implementsFirewallRuleGetResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleGetResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleGetResponseFilterFirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RuleGetResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type RuleGetResponseFilterFirewallFilter struct {
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
	Ref  string                                  `json:"ref"`
	JSON ruleGetResponseFilterFirewallFilterJSON `json:"-"`
}

// ruleGetResponseFilterFirewallFilterJSON contains the JSON metadata for the
// struct [RuleGetResponseFilterFirewallFilter]
type ruleGetResponseFilterFirewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseFilterFirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseFilterFirewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleGetResponseFilterFirewallFilter) implementsFirewallRuleGetResponseFilter() {}

type RuleGetResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                           `json:"deleted,required"`
	JSON    ruleGetResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// ruleGetResponseFilterFirewallDeletedFilterJSON contains the JSON metadata for
// the struct [RuleGetResponseFilterFirewallDeletedFilter]
type ruleGetResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r RuleGetResponseFilterFirewallDeletedFilter) implementsFirewallRuleGetResponseFilter() {}

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

func (r RuleGetResponseProduct) IsKnown() bool {
	switch r {
	case RuleGetResponseProductZoneLockdown, RuleGetResponseProductUABlock, RuleGetResponseProductBic, RuleGetResponseProductHot, RuleGetResponseProductSecurityLevel, RuleGetResponseProductRateLimit, RuleGetResponseProductWAF:
		return true
	}
	return false
}

type RuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleNewResponse                                         `json:"result,required,nullable"`
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

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleUpdateResponse                                        `json:"result,required"`
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

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleDeleteResponse                                        `json:"result,required"`
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

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleEditResponse                                        `json:"result,required,nullable"`
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

// Whether the API call was successful
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	// The unique identifier of the firewall rule.
	PathID param.Field[string] `path:"id,required"`
	// The unique identifier of the firewall rule.
	QueryID param.Field[string] `query:"id"`
}

// URLQuery serializes [RuleGetParams]'s query parameters as `url.Values`.
func (r RuleGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   RuleGetResponse                                           `json:"result,required"`
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

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
