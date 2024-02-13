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

// Create one or more firewall rules.
func (r *FirewallRuleService) FirewallRulesNewFirewallRules(ctx context.Context, zoneIdentifier string, body FirewallRuleFirewallRulesNewFirewallRulesParams, opts ...option.RequestOption) (res *[]FirewallRuleFirewallRulesNewFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *FirewallRuleService) FirewallRulesListFirewallRules(ctx context.Context, zoneIdentifier string, query FirewallRuleFirewallRulesListFirewallRulesParams, opts ...option.RequestOption) (res *[]FirewallRuleFirewallRulesListFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleFirewallRulesListFirewallRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more existing firewall rules.
func (r *FirewallRuleService) FirewallRulesUpdateFirewallRules(ctx context.Context, zoneIdentifier string, body FirewallRuleFirewallRulesUpdateFirewallRulesParams, opts ...option.RequestOption) (res *[]FirewallRuleFirewallRulesUpdateFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the priority of existing firewall rules.
func (r *FirewallRuleService) FirewallRulesUpdatePriorityOfFirewallRules(ctx context.Context, zoneIdentifier string, body FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams, opts ...option.RequestOption) (res *[]FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
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
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleUpdateResponseFilter)(nil)).Elem(), "")
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

func (r FirewallRuleUpdateResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleUpdateResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleUpdateResponseProduct string

const (
	FirewallRuleUpdateResponseProductZoneLockdown  FirewallRuleUpdateResponseProduct = "zoneLockdown"
	FirewallRuleUpdateResponseProductUaBlock       FirewallRuleUpdateResponseProduct = "uaBlock"
	FirewallRuleUpdateResponseProductBic           FirewallRuleUpdateResponseProduct = "bic"
	FirewallRuleUpdateResponseProductHot           FirewallRuleUpdateResponseProduct = "hot"
	FirewallRuleUpdateResponseProductSecurityLevel FirewallRuleUpdateResponseProduct = "securityLevel"
	FirewallRuleUpdateResponseProductRateLimit     FirewallRuleUpdateResponseProduct = "rateLimit"
	FirewallRuleUpdateResponseProductWAF           FirewallRuleUpdateResponseProduct = "waf"
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
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleDeleteResponseFilter)(nil)).Elem(), "")
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

func (r FirewallRuleDeleteResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleDeleteResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleDeleteResponseProduct string

const (
	FirewallRuleDeleteResponseProductZoneLockdown  FirewallRuleDeleteResponseProduct = "zoneLockdown"
	FirewallRuleDeleteResponseProductUaBlock       FirewallRuleDeleteResponseProduct = "uaBlock"
	FirewallRuleDeleteResponseProductBic           FirewallRuleDeleteResponseProduct = "bic"
	FirewallRuleDeleteResponseProductHot           FirewallRuleDeleteResponseProduct = "hot"
	FirewallRuleDeleteResponseProductSecurityLevel FirewallRuleDeleteResponseProduct = "securityLevel"
	FirewallRuleDeleteResponseProductRateLimit     FirewallRuleDeleteResponseProduct = "rateLimit"
	FirewallRuleDeleteResponseProductWAF           FirewallRuleDeleteResponseProduct = "waf"
)

type FirewallRuleFirewallRulesNewFirewallRulesResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleFirewallRulesNewFirewallRulesResponseAction `json:"action,required"`
	Filter FirewallRuleFirewallRulesNewFirewallRulesResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                    `json:"priority"`
	Products []FirewallRuleFirewallRulesNewFirewallRulesResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                `json:"ref"`
	JSON firewallRuleFirewallRulesNewFirewallRulesResponseJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseJSON contains the JSON metadata
// for the struct [FirewallRuleFirewallRulesNewFirewallRulesResponse]
type firewallRuleFirewallRulesNewFirewallRulesResponseJSON struct {
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

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleFirewallRulesNewFirewallRulesResponseAction string

const (
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionBlock            FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "block"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionChallenge        FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "challenge"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionJsChallenge      FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "js_challenge"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionManagedChallenge FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "managed_challenge"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionAllow            FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "allow"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionLog              FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "log"
	FirewallRuleFirewallRulesNewFirewallRulesResponseActionBypass           FirewallRuleFirewallRulesNewFirewallRulesResponseAction = "bypass"
)

// Union satisfied by
// [FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilter] or
// [FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleFirewallRulesNewFirewallRulesResponseFilter interface {
	implementsFirewallRuleFirewallRulesNewFirewallRulesResponseFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleFirewallRulesNewFirewallRulesResponseFilter)(nil)).Elem(), "")
}

type FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                                                     `json:"ref"`
	JSON firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilter]
type firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsFilter) implementsFirewallRuleFirewallRulesNewFirewallRulesResponseFilter() {
}

type FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                              `json:"deleted,required"`
	JSON    firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilter]
type firewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesNewFirewallRulesResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleFirewallRulesNewFirewallRulesResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleFirewallRulesNewFirewallRulesResponseProduct string

const (
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductZoneLockdown  FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "zoneLockdown"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductUaBlock       FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "uaBlock"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductBic           FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "bic"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductHot           FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "hot"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductSecurityLevel FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "securityLevel"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductRateLimit     FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "rateLimit"
	FirewallRuleFirewallRulesNewFirewallRulesResponseProductWAF           FirewallRuleFirewallRulesNewFirewallRulesResponseProduct = "waf"
)

type FirewallRuleFirewallRulesListFirewallRulesResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleFirewallRulesListFirewallRulesResponseAction `json:"action,required"`
	Filter FirewallRuleFirewallRulesListFirewallRulesResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                     `json:"priority"`
	Products []FirewallRuleFirewallRulesListFirewallRulesResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                 `json:"ref"`
	JSON firewallRuleFirewallRulesListFirewallRulesResponseJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseJSON contains the JSON
// metadata for the struct [FirewallRuleFirewallRulesListFirewallRulesResponse]
type firewallRuleFirewallRulesListFirewallRulesResponseJSON struct {
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

func (r *FirewallRuleFirewallRulesListFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleFirewallRulesListFirewallRulesResponseAction string

const (
	FirewallRuleFirewallRulesListFirewallRulesResponseActionBlock            FirewallRuleFirewallRulesListFirewallRulesResponseAction = "block"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionChallenge        FirewallRuleFirewallRulesListFirewallRulesResponseAction = "challenge"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionJsChallenge      FirewallRuleFirewallRulesListFirewallRulesResponseAction = "js_challenge"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionManagedChallenge FirewallRuleFirewallRulesListFirewallRulesResponseAction = "managed_challenge"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionAllow            FirewallRuleFirewallRulesListFirewallRulesResponseAction = "allow"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionLog              FirewallRuleFirewallRulesListFirewallRulesResponseAction = "log"
	FirewallRuleFirewallRulesListFirewallRulesResponseActionBypass           FirewallRuleFirewallRulesListFirewallRulesResponseAction = "bypass"
)

// Union satisfied by
// [FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilter] or
// [FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleFirewallRulesListFirewallRulesResponseFilter interface {
	implementsFirewallRuleFirewallRulesListFirewallRulesResponseFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleFirewallRulesListFirewallRulesResponseFilter)(nil)).Elem(), "")
}

type FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                                                      `json:"ref"`
	JSON firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilter]
type firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsFilter) implementsFirewallRuleFirewallRulesListFirewallRulesResponseFilter() {
}

type FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                               `json:"deleted,required"`
	JSON    firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilter]
type firewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesListFirewallRulesResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleFirewallRulesListFirewallRulesResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleFirewallRulesListFirewallRulesResponseProduct string

const (
	FirewallRuleFirewallRulesListFirewallRulesResponseProductZoneLockdown  FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "zoneLockdown"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductUaBlock       FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "uaBlock"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductBic           FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "bic"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductHot           FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "hot"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductSecurityLevel FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "securityLevel"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductRateLimit     FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "rateLimit"
	FirewallRuleFirewallRulesListFirewallRulesResponseProductWAF           FirewallRuleFirewallRulesListFirewallRulesResponseProduct = "waf"
)

type FirewallRuleFirewallRulesUpdateFirewallRulesResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction `json:"action,required"`
	Filter FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                       `json:"priority"`
	Products []FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                   `json:"ref"`
	JSON firewallRuleFirewallRulesUpdateFirewallRulesResponseJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseJSON contains the JSON
// metadata for the struct [FirewallRuleFirewallRulesUpdateFirewallRulesResponse]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseJSON struct {
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

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction string

const (
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionBlock            FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "block"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionChallenge        FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "challenge"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionJsChallenge      FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "js_challenge"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionManagedChallenge FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "managed_challenge"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionAllow            FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "allow"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionLog              FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "log"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseActionBypass           FirewallRuleFirewallRulesUpdateFirewallRulesResponseAction = "bypass"
)

// Union satisfied by
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilter] or
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter interface {
	implementsFirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter)(nil)).Elem(), "")
}

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                                                        `json:"ref"`
	JSON firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilter]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsFilter) implementsFirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter() {
}

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                 `json:"deleted,required"`
	JSON    firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilter]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesUpdateFirewallRulesResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleFirewallRulesUpdateFirewallRulesResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct string

const (
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductZoneLockdown  FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "zoneLockdown"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductUaBlock       FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "uaBlock"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductBic           FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "bic"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductHot           FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "hot"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductSecurityLevel FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "securityLevel"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductRateLimit     FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "rateLimit"
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseProductWAF           FirewallRuleFirewallRulesUpdateFirewallRulesResponseProduct = "waf"
)

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id,required"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction `json:"action,required"`
	Filter FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter `json:"filter,required"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the firewall rule.
	Description string `json:"description"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                                 `json:"priority"`
	Products []FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                             `json:"ref"`
	JSON firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON contains the
// JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON struct {
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

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction string

const (
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionBlock            FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "block"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionChallenge        FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "challenge"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionJsChallenge      FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "js_challenge"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionManagedChallenge FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "managed_challenge"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionAllow            FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "allow"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionLog              FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "log"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseActionBypass           FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseAction = "bypass"
)

// Union satisfied by
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilter]
// or
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilter].
type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter interface {
	implementsFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter)(nil)).Elem(), "")
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilter struct {
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
	Ref  string                                                                                  `json:"ref"`
	JSON firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilter]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsFilter) implementsFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter() {
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                           `json:"deleted,required"`
	JSON    firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilter]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct string

const (
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductZoneLockdown  FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "zoneLockdown"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductUaBlock       FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "uaBlock"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductBic           FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "bic"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductHot           FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "hot"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductSecurityLevel FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "securityLevel"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductRateLimit     FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "rateLimit"
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProductWAF           FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseProduct = "waf"
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
	apijson.RegisterUnion(reflect.TypeOf((*FirewallRuleGetResponseFilter)(nil)).Elem(), "")
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

func (r FirewallRuleGetResponseFilterLegacyJhsDeletedFilter) implementsFirewallRuleGetResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallRuleGetResponseProduct string

const (
	FirewallRuleGetResponseProductZoneLockdown  FirewallRuleGetResponseProduct = "zoneLockdown"
	FirewallRuleGetResponseProductUaBlock       FirewallRuleGetResponseProduct = "uaBlock"
	FirewallRuleGetResponseProductBic           FirewallRuleGetResponseProduct = "bic"
	FirewallRuleGetResponseProductHot           FirewallRuleGetResponseProduct = "hot"
	FirewallRuleGetResponseProductSecurityLevel FirewallRuleGetResponseProduct = "securityLevel"
	FirewallRuleGetResponseProductRateLimit     FirewallRuleGetResponseProduct = "rateLimit"
	FirewallRuleGetResponseProductWAF           FirewallRuleGetResponseProduct = "waf"
)

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

// Whether the API call was successful
type FirewallRuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallRuleUpdateResponseEnvelopeSuccessTrue FirewallRuleUpdateResponseEnvelopeSuccess = true
)

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

// Whether the API call was successful
type FirewallRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallRuleDeleteResponseEnvelopeSuccessTrue FirewallRuleDeleteResponseEnvelopeSuccess = true
)

type FirewallRuleFirewallRulesNewFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleFirewallRulesNewFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelope struct {
	Errors   []FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleFirewallRulesNewFirewallRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelope]
type firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrors struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrors]
type firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessages struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessages]
type firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeSuccess bool

const (
	FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeSuccessTrue FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeSuccess = true
)

type FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                 `json:"total_count"`
	JSON       firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfoJSON contains
// the JSON metadata for the struct
// [FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfo]
type firewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesNewFirewallRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesListFirewallRulesParams struct {
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

// URLQuery serializes [FirewallRuleFirewallRulesListFirewallRulesParams]'s query
// parameters as `url.Values`.
func (r FirewallRuleFirewallRulesListFirewallRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallRuleFirewallRulesListFirewallRulesResponseEnvelope struct {
	Errors   []FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleFirewallRulesListFirewallRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseEnvelope]
type firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrors]
type firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessages]
type firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeSuccess bool

const (
	FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeSuccessTrue FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeSuccess = true
)

type FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfo]
type firewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesListFirewallRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdateFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleFirewallRulesUpdateFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelope struct {
	Errors   []FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleFirewallRulesUpdateFirewallRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelope]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrors]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessages]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeSuccess bool

const (
	FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeSuccessTrue FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeSuccess = true
)

type FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfo]
type firewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdateFirewallRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelope struct {
	Errors   []FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelope]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrors]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessages]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeSuccess bool

const (
	FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeSuccessTrue FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeSuccess = true
)

type FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                              `json:"total_count"`
	JSON       firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfo]
type firewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

// Whether the API call was successful
type FirewallRuleGetResponseEnvelopeSuccess bool

const (
	FirewallRuleGetResponseEnvelopeSuccessTrue FirewallRuleGetResponseEnvelopeSuccess = true
)
