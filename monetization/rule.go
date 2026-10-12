// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package monetization

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
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

// Replaces the zone's Payment Required ruleset with the submitted desired state.
// Submit an empty rules array to clear all payment rules.
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *MonetizationRuleCollection, err error) {
	var env RuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes every Payment Required rule from the zone. Any rules the service does
// not own are preserved. Deleting a zone with no payment rules is a no-op. Returns
// the resulting (empty) rule collection.
func (r *RuleService) Delete(ctx context.Context, body RuleDeleteParams, opts ...option.RequestOption) (res *MonetizationRuleCollection, err error) {
	var env RuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes a single Payment Required rule identified by its ID. Every other rule is
// preserved. Returns the resulting rule collection.
func (r *RuleService) DeleteRule(ctx context.Context, ruleID string, body RuleDeleteRuleParams, opts ...option.RequestOption) (res *MonetizationRuleCollection, err error) {
	var env RuleDeleteRuleResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules/%s", body.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Applies a partial update to a single Payment Required rule. Only the fields
// present in the request body are changed; omitted fields keep their current
// values. Returns the full resulting rule collection.
func (r *RuleService) EditRule(ctx context.Context, ruleID string, params RuleEditRuleParams, opts ...option.RequestOption) (res *MonetizationRuleCollection, err error) {
	var env RuleEditRuleResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules/%s", params.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the currently deployed Payment Required rules for a zone. A zone with no
// Payment Required ruleset returns 404.
func (r *RuleService) Get(ctx context.Context, query RuleGetParams, opts ...option.RequestOption) (res *MonetizationRuleCollection, err error) {
	var env RuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a single Payment Required rule identified by its ID.
func (r *RuleService) GetRule(ctx context.Context, ruleID string, query RuleGetRuleParams, opts ...option.RequestOption) (res *MonetizationRule, err error) {
	var env RuleGetRuleResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/monetization/rules/%s", query.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Payment rule with a fixed price set at configuration time.
type MonetizationRule struct {
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID string `json:"id"`
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address     string `json:"address"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression string `json:"expression"`
	// Price in the smallest indivisible unit of the configured payment token, encoded
	// as a decimal string. Must be a canonical decimal integer in [1000, 100000000]:
	// no leading zeros, and a bare JSON number is rejected. The price is required for
	// the fixed-price schemes ("exact" and "upto") and must be at least 1000, the
	// smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token
	// such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the
	// scheme is "origin_controlled" the origin server sets pricing dynamically and the
	// rule carries no price: the field must be omitted — any provided value, including
	// "0", is rejected because it would not be enforced. Responses always encode this
	// as a string and omit it for "origin_controlled" rules; the pattern matches
	// exactly the set of values the server accepts.
	Price string `json:"price"`
	// X402 payment scheme. "exact" requires the specified payment amount; "upto"
	// permits a payment up to the specified amount. Both fixed-price schemes require
	// the price field.
	Scheme MonetizationRuleScheme `json:"scheme"`
	JSON   monetizationRuleJSON   `json:"-"`
	union  MonetizationRuleUnion
}

// monetizationRuleJSON contains the JSON metadata for the struct
// [MonetizationRule]
type monetizationRuleJSON struct {
	ID          apijson.Field
	Address     apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Price       apijson.Field
	Scheme      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r monetizationRuleJSON) RawJSON() string {
	return r.raw
}

func (r *MonetizationRule) UnmarshalJSON(data []byte) (err error) {
	*r = MonetizationRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MonetizationRuleUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice],
// [MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled].
func (r MonetizationRule) AsUnion() MonetizationRuleUnion {
	return r.union
}

// Payment rule with a fixed price set at configuration time.
//
// Union satisfied by
// [MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice] or
// [MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled].
type MonetizationRuleUnion interface {
	implementsMonetizationRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MonetizationRuleUnion)(nil)).Elem(),
		"scheme",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice{}),
			DiscriminatorValue: "exact",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice{}),
			DiscriminatorValue: "upto",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled{}),
			DiscriminatorValue: "origin_controlled",
		},
	)
}

// Payment rule with a fixed price set at configuration time.
type MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice struct {
	// The server-assigned unique ID of the payment rule. Stable across full-ruleset
	// replacements.
	ID   string                                                               `json:"id" api:"required"`
	JSON monetizationRuleMonetizationRulesMonetizationRuleInputFixedPriceJSON `json:"-"`
	MonetizationRuleInputFixedPrice
}

// monetizationRuleMonetizationRulesMonetizationRuleInputFixedPriceJSON contains
// the JSON metadata for the struct
// [MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice]
type monetizationRuleMonetizationRulesMonetizationRuleInputFixedPriceJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetizationRuleMonetizationRulesMonetizationRuleInputFixedPriceJSON) RawJSON() string {
	return r.raw
}

func (r MonetizationRuleMonetizationRulesMonetizationRuleInputFixedPrice) implementsMonetizationRule() {
}

// Payment rule whose price the origin server sets dynamically. The rule carries no
// price and the price field must be omitted — any provided value, including "0",
// is rejected because it would not be enforced.
type MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled struct {
	// The server-assigned unique ID of the payment rule. Stable across full-ruleset
	// replacements.
	ID   string                                                                     `json:"id" api:"required"`
	JSON monetizationRuleMonetizationRulesMonetizationRuleInputOriginControlledJSON `json:"-"`
	MonetizationRuleInputOriginControlled
}

// monetizationRuleMonetizationRulesMonetizationRuleInputOriginControlledJSON
// contains the JSON metadata for the struct
// [MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled]
type monetizationRuleMonetizationRulesMonetizationRuleInputOriginControlledJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetizationRuleMonetizationRulesMonetizationRuleInputOriginControlledJSON) RawJSON() string {
	return r.raw
}

func (r MonetizationRuleMonetizationRulesMonetizationRuleInputOriginControlled) implementsMonetizationRule() {
}

// X402 payment scheme. "exact" requires the specified payment amount; "upto"
// permits a payment up to the specified amount. Both fixed-price schemes require
// the price field.
type MonetizationRuleScheme string

const (
	MonetizationRuleSchemeExact            MonetizationRuleScheme = "exact"
	MonetizationRuleSchemeUpto             MonetizationRuleScheme = "upto"
	MonetizationRuleSchemeOriginControlled MonetizationRuleScheme = "origin_controlled"
)

func (r MonetizationRuleScheme) IsKnown() bool {
	switch r {
	case MonetizationRuleSchemeExact, MonetizationRuleSchemeUpto, MonetizationRuleSchemeOriginControlled:
		return true
	}
	return false
}

// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
// deploy endpoint, so the response can be read back as the desired state.
type MonetizationRuleCollection struct {
	// The zone's payment rules, in the order they are evaluated. Empty when the zone
	// has no payment rules.
	Rules []MonetizationRule             `json:"rules" api:"required"`
	JSON  monetizationRuleCollectionJSON `json:"-"`
}

// monetizationRuleCollectionJSON contains the JSON metadata for the struct
// [MonetizationRuleCollection]
type monetizationRuleCollectionJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonetizationRuleCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetizationRuleCollectionJSON) RawJSON() string {
	return r.raw
}

// Payment rule with a fixed price set at configuration time.
type MonetizationRuleInputParam struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address param.Field[string] `json:"address" api:"required"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression param.Field[string] `json:"expression" api:"required"`
	// X402 payment scheme. "exact" requires the specified payment amount; "upto"
	// permits a payment up to the specified amount. Both fixed-price schemes require
	// the price field.
	Scheme param.Field[MonetizationRuleInputScheme] `json:"scheme" api:"required"`
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	// Price in the smallest indivisible unit of the configured payment token, encoded
	// as a decimal string. Must be a canonical decimal integer in [1000, 100000000]:
	// no leading zeros, and a bare JSON number is rejected. The price is required for
	// the fixed-price schemes ("exact" and "upto") and must be at least 1000, the
	// smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token
	// such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the
	// scheme is "origin_controlled" the origin server sets pricing dynamically and the
	// rule carries no price: the field must be omitted — any provided value, including
	// "0", is rejected because it would not be enforced. Responses always encode this
	// as a string and omit it for "origin_controlled" rules; the pattern matches
	// exactly the set of values the server accepts.
	Price param.Field[string] `json:"price"`
}

func (r MonetizationRuleInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MonetizationRuleInputParam) implementsMonetizationRuleInputUnionParam() {}

// Payment rule with a fixed price set at configuration time.
//
// Satisfied by [monetization.MonetizationRuleInputFixedPriceParam],
// [monetization.MonetizationRuleInputOriginControlledParam],
// [MonetizationRuleInputParam].
type MonetizationRuleInputUnionParam interface {
	implementsMonetizationRuleInputUnionParam()
}

// X402 payment scheme. "exact" requires the specified payment amount; "upto"
// permits a payment up to the specified amount. Both fixed-price schemes require
// the price field.
type MonetizationRuleInputScheme string

const (
	MonetizationRuleInputSchemeExact            MonetizationRuleInputScheme = "exact"
	MonetizationRuleInputSchemeUpto             MonetizationRuleInputScheme = "upto"
	MonetizationRuleInputSchemeOriginControlled MonetizationRuleInputScheme = "origin_controlled"
)

func (r MonetizationRuleInputScheme) IsKnown() bool {
	switch r {
	case MonetizationRuleInputSchemeExact, MonetizationRuleInputSchemeUpto, MonetizationRuleInputSchemeOriginControlled:
		return true
	}
	return false
}

// Payment rule with a fixed price set at configuration time.
type MonetizationRuleInputFixedPrice struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address string `json:"address" api:"required"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression string `json:"expression" api:"required"`
	// Price in the smallest indivisible unit of the configured payment token, encoded
	// as a decimal string. Must be a canonical decimal integer in [1000, 100000000]:
	// no leading zeros, and a bare JSON number is rejected. The price is required for
	// the fixed-price schemes ("exact" and "upto") and must be at least 1000, the
	// smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token
	// such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the
	// scheme is "origin_controlled" the origin server sets pricing dynamically and the
	// rule carries no price: the field must be omitted — any provided value, including
	// "0", is rejected because it would not be enforced. Responses always encode this
	// as a string and omit it for "origin_controlled" rules; the pattern matches
	// exactly the set of values the server accepts.
	Price string `json:"price" api:"required"`
	// X402 payment scheme. "exact" requires the specified payment amount; "upto"
	// permits a payment up to the specified amount. Both fixed-price schemes require
	// the price field.
	Scheme MonetizationRuleInputFixedPriceScheme `json:"scheme" api:"required"`
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID          string                              `json:"id"`
	Description string                              `json:"description"`
	Enabled     bool                                `json:"enabled"`
	JSON        monetizationRuleInputFixedPriceJSON `json:"-"`
}

// monetizationRuleInputFixedPriceJSON contains the JSON metadata for the struct
// [MonetizationRuleInputFixedPrice]
type monetizationRuleInputFixedPriceJSON struct {
	Address     apijson.Field
	Expression  apijson.Field
	Price       apijson.Field
	Scheme      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonetizationRuleInputFixedPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetizationRuleInputFixedPriceJSON) RawJSON() string {
	return r.raw
}

// X402 payment scheme. "exact" requires the specified payment amount; "upto"
// permits a payment up to the specified amount. Both fixed-price schemes require
// the price field.
type MonetizationRuleInputFixedPriceScheme string

const (
	MonetizationRuleInputFixedPriceSchemeExact MonetizationRuleInputFixedPriceScheme = "exact"
	MonetizationRuleInputFixedPriceSchemeUpto  MonetizationRuleInputFixedPriceScheme = "upto"
)

func (r MonetizationRuleInputFixedPriceScheme) IsKnown() bool {
	switch r {
	case MonetizationRuleInputFixedPriceSchemeExact, MonetizationRuleInputFixedPriceSchemeUpto:
		return true
	}
	return false
}

// Payment rule with a fixed price set at configuration time.
type MonetizationRuleInputFixedPriceParam struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address param.Field[string] `json:"address" api:"required"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression param.Field[string] `json:"expression" api:"required"`
	// Price in the smallest indivisible unit of the configured payment token, encoded
	// as a decimal string. Must be a canonical decimal integer in [1000, 100000000]:
	// no leading zeros, and a bare JSON number is rejected. The price is required for
	// the fixed-price schemes ("exact" and "upto") and must be at least 1000, the
	// smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token
	// such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the
	// scheme is "origin_controlled" the origin server sets pricing dynamically and the
	// rule carries no price: the field must be omitted — any provided value, including
	// "0", is rejected because it would not be enforced. Responses always encode this
	// as a string and omit it for "origin_controlled" rules; the pattern matches
	// exactly the set of values the server accepts.
	Price param.Field[string] `json:"price" api:"required"`
	// X402 payment scheme. "exact" requires the specified payment amount; "upto"
	// permits a payment up to the specified amount. Both fixed-price schemes require
	// the price field.
	Scheme param.Field[MonetizationRuleInputFixedPriceScheme] `json:"scheme" api:"required"`
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r MonetizationRuleInputFixedPriceParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MonetizationRuleInputFixedPriceParam) implementsMonetizationRuleInputUnionParam() {}

// Payment rule whose price the origin server sets dynamically. The rule carries no
// price and the price field must be omitted — any provided value, including "0",
// is rejected because it would not be enforced.
type MonetizationRuleInputOriginControlled struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address string `json:"address" api:"required"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression string `json:"expression" api:"required"`
	// X402 payment scheme. "origin_controlled" lets the origin server set pricing
	// dynamically; the rule carries no price and the price field must be omitted.
	Scheme MonetizationRuleInputOriginControlledScheme `json:"scheme" api:"required"`
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID          string                                    `json:"id"`
	Description string                                    `json:"description"`
	Enabled     bool                                      `json:"enabled"`
	JSON        monetizationRuleInputOriginControlledJSON `json:"-"`
}

// monetizationRuleInputOriginControlledJSON contains the JSON metadata for the
// struct [MonetizationRuleInputOriginControlled]
type monetizationRuleInputOriginControlledJSON struct {
	Address     apijson.Field
	Expression  apijson.Field
	Scheme      apijson.Field
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonetizationRuleInputOriginControlled) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monetizationRuleInputOriginControlledJSON) RawJSON() string {
	return r.raw
}

// X402 payment scheme. "origin_controlled" lets the origin server set pricing
// dynamically; the rule carries no price and the price field must be omitted.
type MonetizationRuleInputOriginControlledScheme string

const (
	MonetizationRuleInputOriginControlledSchemeOriginControlled MonetizationRuleInputOriginControlledScheme = "origin_controlled"
)

func (r MonetizationRuleInputOriginControlledScheme) IsKnown() bool {
	switch r {
	case MonetizationRuleInputOriginControlledSchemeOriginControlled:
		return true
	}
	return false
}

// Payment rule whose price the origin server sets dynamically. The rule carries no
// price and the price field must be omitted — any provided value, including "0",
// is rejected because it would not be enforced.
type MonetizationRuleInputOriginControlledParam struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The address must pass wallet screening whenever its rule is deployed
	// or patched.
	Address param.Field[string] `json:"address" api:"required"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression param.Field[string] `json:"expression" api:"required"`
	// X402 payment scheme. "origin_controlled" lets the origin server set pricing
	// dynamically; the rule carries no price and the price field must be omitted.
	Scheme param.Field[MonetizationRuleInputOriginControlledScheme] `json:"scheme" api:"required"`
	// Optional on input. Include the ID of an existing payment rule to update it in
	// place, preserving its stable identity across the full-ruleset replacement; omit
	// it to create a new rule. An ID that does not match an existing payment rule in
	// the zone is rejected. Rules omitted from the request are deleted.
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
}

func (r MonetizationRuleInputOriginControlledParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MonetizationRuleInputOriginControlledParam) implementsMonetizationRuleInputUnionParam() {}

// Partial update to a single payment rule. Only the fields present are modified;
// omitted fields keep their existing values.
type MonetizationRulePatchParam struct {
	// 0x-prefixed 20-byte hexadecimal Ethereum address. Mixed-case addresses must
	// carry a valid EIP-55 checksum; all-lowercase or all-uppercase addresses are also
	// accepted. The effective address must pass wallet screening whenever the rule is
	// patched.
	Address     param.Field[string] `json:"address"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	// Wirefilter expression identifying the requests that require payment. Forwarded
	// verbatim to the Rulesets API, which validates its syntax.
	Expression param.Field[string] `json:"expression"`
	// Price in the smallest indivisible unit of the configured payment token, encoded
	// as a decimal string. Must be a canonical decimal integer in [1000, 100000000]:
	// no leading zeros, and a bare JSON number is rejected. The price is required for
	// the fixed-price schemes ("exact" and "upto") and must be at least 1000, the
	// smallest amount the payment facilitator can settle ($0.001 for a 6-decimal token
	// such as USDC), and at most 100000000 ($100 for a 6-decimal token). When the
	// scheme is "origin_controlled" the origin server sets pricing dynamically and the
	// rule carries no price: the field must be omitted — any provided value, including
	// "0", is rejected because it would not be enforced. Responses always encode this
	// as a string and omit it for "origin_controlled" rules; the pattern matches
	// exactly the set of values the server accepts.
	Price param.Field[string] `json:"price"`
	// X402 payment scheme. "exact" requires the specified payment amount; "upto"
	// permits a payment up to the specified amount; "origin_controlled" lets the
	// origin server set pricing dynamically, in which case the rule carries no price
	// and the price field must be omitted.
	Scheme param.Field[MonetizationRulePatchScheme] `json:"scheme"`
}

func (r MonetizationRulePatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// X402 payment scheme. "exact" requires the specified payment amount; "upto"
// permits a payment up to the specified amount; "origin_controlled" lets the
// origin server set pricing dynamically, in which case the rule carries no price
// and the price field must be omitted.
type MonetizationRulePatchScheme string

const (
	MonetizationRulePatchSchemeExact            MonetizationRulePatchScheme = "exact"
	MonetizationRulePatchSchemeUpto             MonetizationRulePatchScheme = "upto"
	MonetizationRulePatchSchemeOriginControlled MonetizationRulePatchScheme = "origin_controlled"
)

func (r MonetizationRulePatchScheme) IsKnown() bool {
	switch r {
	case MonetizationRulePatchSchemeExact, MonetizationRulePatchSchemeUpto, MonetizationRulePatchSchemeOriginControlled:
		return true
	}
	return false
}

type MonetizationRulesetInputParam struct {
	// Full desired Payment Required ruleset. An empty array clears all payment rules.
	// The ruleset may contain at most 40 unique wallet addresses; address comparison
	// is case-insensitive. Every address must pass wallet screening before deployment.
	Rules param.Field[[]MonetizationRuleInputUnionParam] `json:"rules" api:"required"`
}

func (r MonetizationRulesetInputParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleUpdateParams struct {
	// The unique ID of the zone.
	ZoneID                   param.Field[string]           `path:"zone_id" api:"required"`
	MonetizationRulesetInput MonetizationRulesetInputParam `json:"monetization_ruleset_input" api:"required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MonetizationRulesetInput)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []RuleUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
	// deploy endpoint, so the response can be read back as the desired state.
	Result  MonetizationRuleCollection        `json:"result" api:"required"`
	Success RuleUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code    int64                                `json:"code"`
	Message string                               `json:"message"`
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
	Code    int64                                  `json:"code"`
	Message string                                 `json:"message"`
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

type RuleDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []RuleDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
	// deploy endpoint, so the response can be read back as the desired state.
	Result  MonetizationRuleCollection        `json:"result" api:"required"`
	Success RuleDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code    int64                                `json:"code"`
	Message string                               `json:"message"`
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
	Code    int64                                  `json:"code"`
	Message string                                 `json:"message"`
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

type RuleDeleteRuleParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleDeleteRuleResponseEnvelope struct {
	Errors   []RuleDeleteRuleResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleDeleteRuleResponseEnvelopeMessages `json:"messages" api:"required"`
	// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
	// deploy endpoint, so the response can be read back as the desired state.
	Result  MonetizationRuleCollection            `json:"result" api:"required"`
	Success RuleDeleteRuleResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    ruleDeleteRuleResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteRuleResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteRuleResponseEnvelope]
type ruleDeleteRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteRuleResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteRuleResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code"`
	Message string                                   `json:"message"`
	JSON    ruleDeleteRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleDeleteRuleResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RuleDeleteRuleResponseEnvelopeErrors]
type ruleDeleteRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteRuleResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteRuleResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code"`
	Message string                                     `json:"message"`
	JSON    ruleDeleteRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleDeleteRuleResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleDeleteRuleResponseEnvelopeMessages]
type ruleDeleteRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteRuleResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteRuleResponseEnvelopeSuccess bool

const (
	RuleDeleteRuleResponseEnvelopeSuccessTrue RuleDeleteRuleResponseEnvelopeSuccess = true
)

func (r RuleDeleteRuleResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteRuleResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditRuleParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Partial update to a single payment rule. Only the fields present are modified;
	// omitted fields keep their existing values.
	MonetizationRulePatch MonetizationRulePatchParam `json:"monetization_rule_patch" api:"required"`
}

func (r RuleEditRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MonetizationRulePatch)
}

type RuleEditRuleResponseEnvelope struct {
	Errors   []RuleEditRuleResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleEditRuleResponseEnvelopeMessages `json:"messages" api:"required"`
	// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
	// deploy endpoint, so the response can be read back as the desired state.
	Result  MonetizationRuleCollection          `json:"result" api:"required"`
	Success RuleEditRuleResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    ruleEditRuleResponseEnvelopeJSON    `json:"-"`
}

// ruleEditRuleResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditRuleResponseEnvelope]
type ruleEditRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditRuleResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleEditRuleResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code"`
	Message string                                 `json:"message"`
	JSON    ruleEditRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleEditRuleResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditRuleResponseEnvelopeErrors]
type ruleEditRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditRuleResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleEditRuleResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code"`
	Message string                                   `json:"message"`
	JSON    ruleEditRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleEditRuleResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleEditRuleResponseEnvelopeMessages]
type ruleEditRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditRuleResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleEditRuleResponseEnvelopeSuccess bool

const (
	RuleEditRuleResponseEnvelopeSuccessTrue RuleEditRuleResponseEnvelopeSuccess = true
)

func (r RuleEditRuleResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditRuleResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleGetParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleGetResponseEnvelope struct {
	Errors   []RuleGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// The zone's payment rules. Mirrors the shape of the ruleset submitted to the
	// deploy endpoint, so the response can be read back as the desired state.
	Result  MonetizationRuleCollection     `json:"result" api:"required"`
	Success RuleGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Code    int64                             `json:"code"`
	Message string                            `json:"message"`
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
	Code    int64                               `json:"code"`
	Message string                              `json:"message"`
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

type RuleGetRuleParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type RuleGetRuleResponseEnvelope struct {
	Errors   []RuleGetRuleResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []RuleGetRuleResponseEnvelopeMessages `json:"messages" api:"required"`
	// Payment rule with a fixed price set at configuration time.
	Result  MonetizationRule                   `json:"result" api:"required"`
	Success RuleGetRuleResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    ruleGetRuleResponseEnvelopeJSON    `json:"-"`
}

// ruleGetRuleResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetRuleResponseEnvelope]
type ruleGetRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetRuleResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleGetRuleResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code"`
	Message string                                `json:"message"`
	JSON    ruleGetRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleGetRuleResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleGetRuleResponseEnvelopeErrors]
type ruleGetRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetRuleResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleGetRuleResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code"`
	Message string                                  `json:"message"`
	JSON    ruleGetRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleGetRuleResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleGetRuleResponseEnvelopeMessages]
type ruleGetRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetRuleResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleGetRuleResponseEnvelopeSuccess bool

const (
	RuleGetRuleResponseEnvelopeSuccessTrue RuleGetRuleResponseEnvelopeSuccess = true
)

func (r RuleGetRuleResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetRuleResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
