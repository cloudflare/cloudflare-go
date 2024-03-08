// File generated from our OpenAPI spec by Stainless.

package cloudflare

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

// Create one or more firewall rules.
func (r *FirewallRuleService) New(ctx context.Context, zoneIdentifier string, body FirewallRuleNewParams, opts ...option.RequestOption) (res *[]FirewallRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing firewall rule.
func (r *FirewallRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleUpdateParams, opts ...option.RequestOption) (res *FirewallRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleUpdateResponseEnvelope
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
func (r *FirewallRuleService) List(ctx context.Context, zoneIdentifier string, query FirewallRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallRuleListResponse], err error) {
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
func (r *FirewallRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing firewall rule.
func (r *FirewallRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleDeleteParams, opts ...option.RequestOption) (res *FirewallRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of an existing firewall rule.
func (r *FirewallRuleService) Edit(ctx context.Context, zoneIdentifier string, id string, body FirewallRuleEditParams, opts ...option.RequestOption) (res *[]FirewallRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a firewall rule.
func (r *FirewallRuleService) Get(ctx context.Context, zoneIdentifier string, id string, query FirewallRuleGetParams, opts ...option.RequestOption) (res *FirewallRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallRuleNewResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleNewResponseAction `json:"action,required"`
	Filter FirewallRuleNewResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                          `json:"priority"`
	Products []FirewallRuleNewResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                      `json:"ref"`
	JSON firewallRuleNewResponseJSON `json:"-"`
}

// firewallRuleNewResponseJSON contains the JSON metadata for the struct
// [FirewallRuleNewResponse]
type firewallRuleNewResponseJSON struct {
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

func (r *FirewallRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleNewResponseAction string

const (
	FirewallRuleNewResponseActionBlock            FirewallRuleNewResponseAction = "block"
	FirewallRuleNewResponseActionChallenge        FirewallRuleNewResponseAction = "challenge"
	FirewallRuleNewResponseActionJsChallenge      FirewallRuleNewResponseAction = "js_challenge"
	FirewallRuleNewResponseActionManagedChallenge FirewallRuleNewResponseAction = "managed_challenge"
	FirewallRuleNewResponseActionAllow            FirewallRuleNewResponseAction = "allow"
	FirewallRuleNewResponseActionLog              FirewallRuleNewResponseAction = "log"
	FirewallRuleNewResponseActionBypass           FirewallRuleNewResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleNewResponseFilterLegacyJhsFilter] or
// [FirewallRuleNewResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleNewResponseFilter interface {
	implementsFirewallRuleNewResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleNewResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleNewResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleNewResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleNewResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                           `json:"ref"`
	JSON firewallRuleNewResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleNewResponseFilterLegacyJhsFilterJSON contains the JSON metadata for
// the struct [FirewallRuleNewResponseFilterLegacyJhsFilter]
type firewallRuleNewResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleNewResponseFilterLegacyJhsFilter) implementsFirewallRuleNewResponseFilter() {}

type FirewallRuleNewResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                    `json:"deleted,required"`
	JSON    firewallRuleNewResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleNewResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleNewResponseFilterLegacyJhsDeletedFilter]
type firewallRuleNewResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleNewResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleNewResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleNewResponseProduct string

const (
	FirewallRuleNewResponseProductZoneLockdown  FirewallRuleNewResponseProduct = "zoneLockdown"
	FirewallRuleNewResponseProductUABlock       FirewallRuleNewResponseProduct = "uaBlock"
	FirewallRuleNewResponseProductBic           FirewallRuleNewResponseProduct = "bic"
	FirewallRuleNewResponseProductHot           FirewallRuleNewResponseProduct = "hot"
	FirewallRuleNewResponseProductSecurityLevel FirewallRuleNewResponseProduct = "securityLevel"
	FirewallRuleNewResponseProductRateLimit     FirewallRuleNewResponseProduct = "rateLimit"
	FirewallRuleNewResponseProductWAF           FirewallRuleNewResponseProduct = "waf"
)

type FirewallRuleUpdateResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleUpdateResponseAction `json:"action,required"`
	Filter FirewallRuleUpdateResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                             `json:"priority"`
	Products []FirewallRuleUpdateResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                         `json:"ref"`
	JSON firewallRuleUpdateResponseJSON `json:"-"`
}

// firewallRuleUpdateResponseJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponse]
type firewallRuleUpdateResponseJSON struct {
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

func (r *FirewallRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleUpdateResponseAction string

const (
	FirewallRuleUpdateResponseActionBlock            FirewallRuleUpdateResponseAction = "block"
	FirewallRuleUpdateResponseActionChallenge        FirewallRuleUpdateResponseAction = "challenge"
	FirewallRuleUpdateResponseActionJsChallenge      FirewallRuleUpdateResponseAction = "js_challenge"
	FirewallRuleUpdateResponseActionManagedChallenge FirewallRuleUpdateResponseAction = "managed_challenge"
	FirewallRuleUpdateResponseActionAllow            FirewallRuleUpdateResponseAction = "allow"
	FirewallRuleUpdateResponseActionLog              FirewallRuleUpdateResponseAction = "log"
	FirewallRuleUpdateResponseActionBypass           FirewallRuleUpdateResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleUpdateResponseFilterLegacyJhsFilter] or
// [FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleUpdateResponseFilter interface {
	implementsFirewallRuleUpdateResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleUpdateResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleUpdateResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleUpdateResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                              `json:"ref"`
	JSON firewallRuleUpdateResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleUpdateResponseFilterLegacyJhsFilterJSON contains the JSON metadata
// for the struct [FirewallRuleUpdateResponseFilterLegacyJhsFilter]
type firewallRuleUpdateResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleUpdateResponseFilterLegacyJhsFilter) implementsFirewallRuleUpdateResponseFilter() {
}

type FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                       `json:"deleted,required"`
	JSON    firewallRuleUpdateResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleUpdateResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter]
type firewallRuleUpdateResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleUpdateResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleUpdateResponseProduct string

const (
	FirewallRuleUpdateResponseProductZoneLockdown  FirewallRuleUpdateResponseProduct = "zoneLockdown"
	FirewallRuleUpdateResponseProductUABlock       FirewallRuleUpdateResponseProduct = "uaBlock"
	FirewallRuleUpdateResponseProductBic           FirewallRuleUpdateResponseProduct = "bic"
	FirewallRuleUpdateResponseProductHot           FirewallRuleUpdateResponseProduct = "hot"
	FirewallRuleUpdateResponseProductSecurityLevel FirewallRuleUpdateResponseProduct = "securityLevel"
	FirewallRuleUpdateResponseProductRateLimit     FirewallRuleUpdateResponseProduct = "rateLimit"
	FirewallRuleUpdateResponseProductWAF           FirewallRuleUpdateResponseProduct = "waf"
)

type FirewallRuleListResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleListResponseAction `json:"action,required"`
	Filter FirewallRuleListResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                           `json:"priority"`
	Products []FirewallRuleListResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                       `json:"ref"`
	JSON firewallRuleListResponseJSON `json:"-"`
}

// firewallRuleListResponseJSON contains the JSON metadata for the struct
// [FirewallRuleListResponse]
type firewallRuleListResponseJSON struct {
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

func (r *FirewallRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleListResponseAction string

const (
	FirewallRuleListResponseActionBlock            FirewallRuleListResponseAction = "block"
	FirewallRuleListResponseActionChallenge        FirewallRuleListResponseAction = "challenge"
	FirewallRuleListResponseActionJsChallenge      FirewallRuleListResponseAction = "js_challenge"
	FirewallRuleListResponseActionManagedChallenge FirewallRuleListResponseAction = "managed_challenge"
	FirewallRuleListResponseActionAllow            FirewallRuleListResponseAction = "allow"
	FirewallRuleListResponseActionLog              FirewallRuleListResponseAction = "log"
	FirewallRuleListResponseActionBypass           FirewallRuleListResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleListResponseFilterLegacyJhsFilter] or
// [FirewallRuleListResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleListResponseFilter interface {
	implementsFirewallRuleListResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleListResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleListResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleListResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleListResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                            `json:"ref"`
	JSON firewallRuleListResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleListResponseFilterLegacyJhsFilterJSON contains the JSON metadata for
// the struct [FirewallRuleListResponseFilterLegacyJhsFilter]
type firewallRuleListResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleListResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleListResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleListResponseFilterLegacyJhsFilter) implementsFirewallRuleListResponseFilter() {}

type FirewallRuleListResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                     `json:"deleted,required"`
	JSON    firewallRuleListResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleListResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleListResponseFilterLegacyJhsDeletedFilter]
type firewallRuleListResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleListResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleListResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleListResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleListResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleListResponseProduct string

const (
	FirewallRuleListResponseProductZoneLockdown  FirewallRuleListResponseProduct = "zoneLockdown"
	FirewallRuleListResponseProductUABlock       FirewallRuleListResponseProduct = "uaBlock"
	FirewallRuleListResponseProductBic           FirewallRuleListResponseProduct = "bic"
	FirewallRuleListResponseProductHot           FirewallRuleListResponseProduct = "hot"
	FirewallRuleListResponseProductSecurityLevel FirewallRuleListResponseProduct = "securityLevel"
	FirewallRuleListResponseProductRateLimit     FirewallRuleListResponseProduct = "rateLimit"
	FirewallRuleListResponseProductWAF           FirewallRuleListResponseProduct = "waf"
)

type FirewallRuleDeleteResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleDeleteResponseAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                           `json:"description"`
	Filter      FirewallRuleDeleteResponseFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                             `json:"priority"`
	Products []FirewallRuleDeleteResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                         `json:"ref"`
	JSON firewallRuleDeleteResponseJSON `json:"-"`
}

// firewallRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponse]
type firewallRuleDeleteResponseJSON struct {
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

func (r *FirewallRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleDeleteResponseAction string

const (
	FirewallRuleDeleteResponseActionBlock            FirewallRuleDeleteResponseAction = "block"
	FirewallRuleDeleteResponseActionChallenge        FirewallRuleDeleteResponseAction = "challenge"
	FirewallRuleDeleteResponseActionJsChallenge      FirewallRuleDeleteResponseAction = "js_challenge"
	FirewallRuleDeleteResponseActionManagedChallenge FirewallRuleDeleteResponseAction = "managed_challenge"
	FirewallRuleDeleteResponseActionAllow            FirewallRuleDeleteResponseAction = "allow"
	FirewallRuleDeleteResponseActionLog              FirewallRuleDeleteResponseAction = "log"
	FirewallRuleDeleteResponseActionBypass           FirewallRuleDeleteResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleDeleteResponseFilterLegacyJhsFilter] or
// [FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleDeleteResponseFilter interface {
	implementsFirewallRuleDeleteResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleDeleteResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleDeleteResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleDeleteResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                              `json:"ref"`
	JSON firewallRuleDeleteResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleDeleteResponseFilterLegacyJhsFilterJSON contains the JSON metadata
// for the struct [FirewallRuleDeleteResponseFilterLegacyJhsFilter]
type firewallRuleDeleteResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleDeleteResponseFilterLegacyJhsFilter) implementsFirewallRuleDeleteResponseFilter() {
}

type FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                       `json:"deleted,required"`
	JSON    firewallRuleDeleteResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleDeleteResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter]
type firewallRuleDeleteResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleDeleteResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleDeleteResponseProduct string

const (
	FirewallRuleDeleteResponseProductZoneLockdown  FirewallRuleDeleteResponseProduct = "zoneLockdown"
	FirewallRuleDeleteResponseProductUABlock       FirewallRuleDeleteResponseProduct = "uaBlock"
	FirewallRuleDeleteResponseProductBic           FirewallRuleDeleteResponseProduct = "bic"
	FirewallRuleDeleteResponseProductHot           FirewallRuleDeleteResponseProduct = "hot"
	FirewallRuleDeleteResponseProductSecurityLevel FirewallRuleDeleteResponseProduct = "securityLevel"
	FirewallRuleDeleteResponseProductRateLimit     FirewallRuleDeleteResponseProduct = "rateLimit"
	FirewallRuleDeleteResponseProductWAF           FirewallRuleDeleteResponseProduct = "waf"
)

type FirewallRuleEditResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleEditResponseAction `json:"action,required"`
	Filter FirewallRuleEditResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                           `json:"priority"`
	Products []FirewallRuleEditResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                       `json:"ref"`
	JSON firewallRuleEditResponseJSON `json:"-"`
}

// firewallRuleEditResponseJSON contains the JSON metadata for the struct
// [FirewallRuleEditResponse]
type firewallRuleEditResponseJSON struct {
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

func (r *FirewallRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleEditResponseAction string

const (
	FirewallRuleEditResponseActionBlock            FirewallRuleEditResponseAction = "block"
	FirewallRuleEditResponseActionChallenge        FirewallRuleEditResponseAction = "challenge"
	FirewallRuleEditResponseActionJsChallenge      FirewallRuleEditResponseAction = "js_challenge"
	FirewallRuleEditResponseActionManagedChallenge FirewallRuleEditResponseAction = "managed_challenge"
	FirewallRuleEditResponseActionAllow            FirewallRuleEditResponseAction = "allow"
	FirewallRuleEditResponseActionLog              FirewallRuleEditResponseAction = "log"
	FirewallRuleEditResponseActionBypass           FirewallRuleEditResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleEditResponseFilterLegacyJhsFilter] or
// [FirewallRuleEditResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleEditResponseFilter interface {
	implementsFirewallRuleEditResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleEditResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleEditResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleEditResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleEditResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                            `json:"ref"`
	JSON firewallRuleEditResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleEditResponseFilterLegacyJhsFilterJSON contains the JSON metadata for
// the struct [FirewallRuleEditResponseFilterLegacyJhsFilter]
type firewallRuleEditResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleEditResponseFilterLegacyJhsFilter) implementsFirewallRuleEditResponseFilter() {}

type FirewallRuleEditResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                     `json:"deleted,required"`
	JSON    firewallRuleEditResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleEditResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleEditResponseFilterLegacyJhsDeletedFilter]
type firewallRuleEditResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleEditResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleEditResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleEditResponseProduct string

const (
	FirewallRuleEditResponseProductZoneLockdown  FirewallRuleEditResponseProduct = "zoneLockdown"
	FirewallRuleEditResponseProductUABlock       FirewallRuleEditResponseProduct = "uaBlock"
	FirewallRuleEditResponseProductBic           FirewallRuleEditResponseProduct = "bic"
	FirewallRuleEditResponseProductHot           FirewallRuleEditResponseProduct = "hot"
	FirewallRuleEditResponseProductSecurityLevel FirewallRuleEditResponseProduct = "securityLevel"
	FirewallRuleEditResponseProductRateLimit     FirewallRuleEditResponseProduct = "rateLimit"
	FirewallRuleEditResponseProductWAF           FirewallRuleEditResponseProduct = "waf"
)

type FirewallRuleGetResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleGetResponseAction `json:"action,required"`
	Filter FirewallRuleGetResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                          `json:"priority"`
	Products []FirewallRuleGetResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                      `json:"ref"`
	JSON firewallRuleGetResponseJSON `json:"-"`
}

// firewallRuleGetResponseJSON contains the JSON metadata for the struct
// [FirewallRuleGetResponse]
type firewallRuleGetResponseJSON struct {
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

func (r *FirewallRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleGetResponseAction string

const (
	FirewallRuleGetResponseActionBlock            FirewallRuleGetResponseAction = "block"
	FirewallRuleGetResponseActionChallenge        FirewallRuleGetResponseAction = "challenge"
	FirewallRuleGetResponseActionJsChallenge      FirewallRuleGetResponseAction = "js_challenge"
	FirewallRuleGetResponseActionManagedChallenge FirewallRuleGetResponseAction = "managed_challenge"
	FirewallRuleGetResponseActionAllow            FirewallRuleGetResponseAction = "allow"
	FirewallRuleGetResponseActionLog              FirewallRuleGetResponseAction = "log"
	FirewallRuleGetResponseActionBypass           FirewallRuleGetResponseAction = "bypass"
)

// Union satisfied by [FirewallRuleGetResponseFilterLegacyJhsFilter] or
// [FirewallRuleGetResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleGetResponseFilter interface {
	implementsFirewallRuleGetResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleGetResponseFilter)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleGetResponseFilterLegacyJhsFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleGetResponseFilterLegacyJhsDeletedFilter{}),
		},
	)
}

type FirewallRuleGetResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                           `json:"ref"`
	JSON firewallRuleGetResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleGetResponseFilterLegacyJhsFilterJSON contains the JSON metadata for
// the struct [FirewallRuleGetResponseFilterLegacyJhsFilter]
type firewallRuleGetResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseFilterLegacyJhsFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleGetResponseFilterLegacyJhsFilter) implementsFirewallRuleGetResponseFilter() {}

type FirewallRuleGetResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                    `json:"deleted,required"`
	JSON    firewallRuleGetResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleGetResponseFilterLegacyJhsDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallRuleGetResponseFilterLegacyJhsDeletedFilter]
type firewallRuleGetResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseFilterLegacyJhsDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleGetResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleGetResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleGetResponseProduct string

const (
	FirewallRuleGetResponseProductZoneLockdown  FirewallRuleGetResponseProduct = "zoneLockdown"
	FirewallRuleGetResponseProductUABlock       FirewallRuleGetResponseProduct = "uaBlock"
	FirewallRuleGetResponseProductBic           FirewallRuleGetResponseProduct = "bic"
	FirewallRuleGetResponseProductHot           FirewallRuleGetResponseProduct = "hot"
	FirewallRuleGetResponseProductSecurityLevel FirewallRuleGetResponseProduct = "securityLevel"
	FirewallRuleGetResponseProductRateLimit     FirewallRuleGetResponseProduct = "rateLimit"
	FirewallRuleGetResponseProductWAF           FirewallRuleGetResponseProduct = "waf"
)

type FirewallRuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleNewResponseEnvelope struct {
	Errors   []FirewallRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleNewResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleNewResponseEnvelope]
type firewallRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    firewallRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeErrors]
type firewallRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeMessages]
type firewallRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallRuleNewResponseEnvelopeSuccess bool

const (
	FirewallRuleNewResponseEnvelopeSuccessTrue FirewallRuleNewResponseEnvelopeSuccess = true
)

type FirewallRuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       firewallRuleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [FirewallRuleNewResponseEnvelopeResultInfo]
type firewallRuleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleUpdateResponseEnvelope struct {
	Errors   []FirewallRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleUpdateResponseEnvelope]
type firewallRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleUpdateResponseEnvelopeErrors]
type firewallRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallRuleUpdateResponseEnvelopeMessages]
type firewallRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallRuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallRuleUpdateResponseEnvelopeSuccessTrue FirewallRuleUpdateResponseEnvelopeSuccess = true
)

type FirewallRuleListParams struct {
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

// URLQuery serializes [FirewallRuleListParams]'s query parameters as `url.Values`.
func (r FirewallRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallRuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r FirewallRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallRuleDeleteResponseEnvelope struct {
	Errors   []FirewallRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleDeleteResponseEnvelope]
type firewallRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleDeleteResponseEnvelopeErrors]
type firewallRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallRuleDeleteResponseEnvelopeMessages]
type firewallRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallRuleDeleteResponseEnvelopeSuccessTrue FirewallRuleDeleteResponseEnvelopeSuccess = true
)

type FirewallRuleEditParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleEditResponseEnvelope struct {
	Errors   []FirewallRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleEditResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleEditResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleEditResponseEnvelope]
type firewallRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    firewallRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleEditResponseEnvelopeErrors]
type firewallRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewallRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleEditResponseEnvelopeMessages]
type firewallRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallRuleEditResponseEnvelopeSuccess bool

const (
	FirewallRuleEditResponseEnvelopeSuccessTrue FirewallRuleEditResponseEnvelopeSuccess = true
)

type FirewallRuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       firewallRuleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [FirewallRuleEditResponseEnvelopeResultInfo]
type firewallRuleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleGetParams struct {
}

type FirewallRuleGetResponseEnvelope struct {
	Errors   []FirewallRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallRuleGetResponseEnvelope]
type firewallRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    firewallRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallRuleGetResponseEnvelopeErrors]
type firewallRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallRuleGetResponseEnvelopeMessages]
type firewallRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallRuleGetResponseEnvelopeSuccess bool

const (
	FirewallRuleGetResponseEnvelopeSuccessTrue FirewallRuleGetResponseEnvelopeSuccess = true
)
