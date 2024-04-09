// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
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

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account or zone ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *RuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RuleNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleNewResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleNewResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string              `json:"description"`
	JSON        ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleNewResponseKind string

const (
	RuleNewResponseKindManaged RuleNewResponseKind = "managed"
	RuleNewResponseKindCustom  RuleNewResponseKind = "custom"
	RuleNewResponseKindRoot    RuleNewResponseKind = "root"
	RuleNewResponseKindZone    RuleNewResponseKind = "zone"
)

func (r RuleNewResponseKind) IsKnown() bool {
	switch r {
	case RuleNewResponseKindManaged, RuleNewResponseKindCustom, RuleNewResponseKindRoot, RuleNewResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleNewResponsePhase string

const (
	RuleNewResponsePhaseDDoSL4                         RuleNewResponsePhase = "ddos_l4"
	RuleNewResponsePhaseDDoSL7                         RuleNewResponsePhase = "ddos_l7"
	RuleNewResponsePhaseHTTPConfigSettings             RuleNewResponsePhase = "http_config_settings"
	RuleNewResponsePhaseHTTPCustomErrors               RuleNewResponsePhase = "http_custom_errors"
	RuleNewResponsePhaseHTTPLogCustomFields            RuleNewResponsePhase = "http_log_custom_fields"
	RuleNewResponsePhaseHTTPRatelimit                  RuleNewResponsePhase = "http_ratelimit"
	RuleNewResponsePhaseHTTPRequestCacheSettings       RuleNewResponsePhase = "http_request_cache_settings"
	RuleNewResponsePhaseHTTPRequestDynamicRedirect     RuleNewResponsePhase = "http_request_dynamic_redirect"
	RuleNewResponsePhaseHTTPRequestFirewallCustom      RuleNewResponsePhase = "http_request_firewall_custom"
	RuleNewResponsePhaseHTTPRequestFirewallManaged     RuleNewResponsePhase = "http_request_firewall_managed"
	RuleNewResponsePhaseHTTPRequestLateTransform       RuleNewResponsePhase = "http_request_late_transform"
	RuleNewResponsePhaseHTTPRequestOrigin              RuleNewResponsePhase = "http_request_origin"
	RuleNewResponsePhaseHTTPRequestRedirect            RuleNewResponsePhase = "http_request_redirect"
	RuleNewResponsePhaseHTTPRequestSanitize            RuleNewResponsePhase = "http_request_sanitize"
	RuleNewResponsePhaseHTTPRequestSbfm                RuleNewResponsePhase = "http_request_sbfm"
	RuleNewResponsePhaseHTTPRequestSelectConfiguration RuleNewResponsePhase = "http_request_select_configuration"
	RuleNewResponsePhaseHTTPRequestTransform           RuleNewResponsePhase = "http_request_transform"
	RuleNewResponsePhaseHTTPResponseCompression        RuleNewResponsePhase = "http_response_compression"
	RuleNewResponsePhaseHTTPResponseFirewallManaged    RuleNewResponsePhase = "http_response_firewall_managed"
	RuleNewResponsePhaseHTTPResponseHeadersTransform   RuleNewResponsePhase = "http_response_headers_transform"
	RuleNewResponsePhaseMagicTransit                   RuleNewResponsePhase = "magic_transit"
	RuleNewResponsePhaseMagicTransitIDsManaged         RuleNewResponsePhase = "magic_transit_ids_managed"
	RuleNewResponsePhaseMagicTransitManaged            RuleNewResponsePhase = "magic_transit_managed"
)

func (r RuleNewResponsePhase) IsKnown() bool {
	switch r {
	case RuleNewResponsePhaseDDoSL4, RuleNewResponsePhaseDDoSL7, RuleNewResponsePhaseHTTPConfigSettings, RuleNewResponsePhaseHTTPCustomErrors, RuleNewResponsePhaseHTTPLogCustomFields, RuleNewResponsePhaseHTTPRatelimit, RuleNewResponsePhaseHTTPRequestCacheSettings, RuleNewResponsePhaseHTTPRequestDynamicRedirect, RuleNewResponsePhaseHTTPRequestFirewallCustom, RuleNewResponsePhaseHTTPRequestFirewallManaged, RuleNewResponsePhaseHTTPRequestLateTransform, RuleNewResponsePhaseHTTPRequestOrigin, RuleNewResponsePhaseHTTPRequestRedirect, RuleNewResponsePhaseHTTPRequestSanitize, RuleNewResponsePhaseHTTPRequestSbfm, RuleNewResponsePhaseHTTPRequestSelectConfiguration, RuleNewResponsePhaseHTTPRequestTransform, RuleNewResponsePhaseHTTPResponseCompression, RuleNewResponsePhaseHTTPResponseFirewallManaged, RuleNewResponsePhaseHTTPResponseHeadersTransform, RuleNewResponsePhaseMagicTransit, RuleNewResponsePhaseMagicTransitIDsManaged, RuleNewResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RuleNewResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleNewResponseRulesAction `json:"action"`
	ActionParameters interface{}                `json:"action_parameters,required"`
	Categories       interface{}                `json:"categories,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                  `json:"version,required"`
	JSON    ruleNewResponseRuleJSON `json:"-"`
	union   RuleNewResponseRulesUnion
}

// ruleNewResponseRuleJSON contains the JSON metadata for the struct
// [RuleNewResponseRule]
type ruleNewResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	ID               apijson.Field
	LastUpdated      apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r ruleNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleNewResponseRule) AsUnion() RuleNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ExecuteRule],
// [rulesets.LogRule] or [rulesets.SkipRule].
type RuleNewResponseRulesUnion interface {
	implementsRulesetsRuleNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

// The action to perform when the rule matches.
type RuleNewResponseRulesAction string

const (
	RuleNewResponseRulesActionBlock   RuleNewResponseRulesAction = "block"
	RuleNewResponseRulesActionExecute RuleNewResponseRulesAction = "execute"
	RuleNewResponseRulesActionLog     RuleNewResponseRulesAction = "log"
	RuleNewResponseRulesActionSkip    RuleNewResponseRulesAction = "skip"
)

func (r RuleNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesActionBlock, RuleNewResponseRulesActionExecute, RuleNewResponseRulesActionLog, RuleNewResponseRulesActionSkip:
		return true
	}
	return false
}

// A ruleset object.
type RuleDeleteResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleDeleteResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleDeleteResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleDeleteResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleDeleteResponseKind string

const (
	RuleDeleteResponseKindManaged RuleDeleteResponseKind = "managed"
	RuleDeleteResponseKindCustom  RuleDeleteResponseKind = "custom"
	RuleDeleteResponseKindRoot    RuleDeleteResponseKind = "root"
	RuleDeleteResponseKindZone    RuleDeleteResponseKind = "zone"
)

func (r RuleDeleteResponseKind) IsKnown() bool {
	switch r {
	case RuleDeleteResponseKindManaged, RuleDeleteResponseKindCustom, RuleDeleteResponseKindRoot, RuleDeleteResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleDeleteResponsePhase string

const (
	RuleDeleteResponsePhaseDDoSL4                         RuleDeleteResponsePhase = "ddos_l4"
	RuleDeleteResponsePhaseDDoSL7                         RuleDeleteResponsePhase = "ddos_l7"
	RuleDeleteResponsePhaseHTTPConfigSettings             RuleDeleteResponsePhase = "http_config_settings"
	RuleDeleteResponsePhaseHTTPCustomErrors               RuleDeleteResponsePhase = "http_custom_errors"
	RuleDeleteResponsePhaseHTTPLogCustomFields            RuleDeleteResponsePhase = "http_log_custom_fields"
	RuleDeleteResponsePhaseHTTPRatelimit                  RuleDeleteResponsePhase = "http_ratelimit"
	RuleDeleteResponsePhaseHTTPRequestCacheSettings       RuleDeleteResponsePhase = "http_request_cache_settings"
	RuleDeleteResponsePhaseHTTPRequestDynamicRedirect     RuleDeleteResponsePhase = "http_request_dynamic_redirect"
	RuleDeleteResponsePhaseHTTPRequestFirewallCustom      RuleDeleteResponsePhase = "http_request_firewall_custom"
	RuleDeleteResponsePhaseHTTPRequestFirewallManaged     RuleDeleteResponsePhase = "http_request_firewall_managed"
	RuleDeleteResponsePhaseHTTPRequestLateTransform       RuleDeleteResponsePhase = "http_request_late_transform"
	RuleDeleteResponsePhaseHTTPRequestOrigin              RuleDeleteResponsePhase = "http_request_origin"
	RuleDeleteResponsePhaseHTTPRequestRedirect            RuleDeleteResponsePhase = "http_request_redirect"
	RuleDeleteResponsePhaseHTTPRequestSanitize            RuleDeleteResponsePhase = "http_request_sanitize"
	RuleDeleteResponsePhaseHTTPRequestSbfm                RuleDeleteResponsePhase = "http_request_sbfm"
	RuleDeleteResponsePhaseHTTPRequestSelectConfiguration RuleDeleteResponsePhase = "http_request_select_configuration"
	RuleDeleteResponsePhaseHTTPRequestTransform           RuleDeleteResponsePhase = "http_request_transform"
	RuleDeleteResponsePhaseHTTPResponseCompression        RuleDeleteResponsePhase = "http_response_compression"
	RuleDeleteResponsePhaseHTTPResponseFirewallManaged    RuleDeleteResponsePhase = "http_response_firewall_managed"
	RuleDeleteResponsePhaseHTTPResponseHeadersTransform   RuleDeleteResponsePhase = "http_response_headers_transform"
	RuleDeleteResponsePhaseMagicTransit                   RuleDeleteResponsePhase = "magic_transit"
	RuleDeleteResponsePhaseMagicTransitIDsManaged         RuleDeleteResponsePhase = "magic_transit_ids_managed"
	RuleDeleteResponsePhaseMagicTransitManaged            RuleDeleteResponsePhase = "magic_transit_managed"
)

func (r RuleDeleteResponsePhase) IsKnown() bool {
	switch r {
	case RuleDeleteResponsePhaseDDoSL4, RuleDeleteResponsePhaseDDoSL7, RuleDeleteResponsePhaseHTTPConfigSettings, RuleDeleteResponsePhaseHTTPCustomErrors, RuleDeleteResponsePhaseHTTPLogCustomFields, RuleDeleteResponsePhaseHTTPRatelimit, RuleDeleteResponsePhaseHTTPRequestCacheSettings, RuleDeleteResponsePhaseHTTPRequestDynamicRedirect, RuleDeleteResponsePhaseHTTPRequestFirewallCustom, RuleDeleteResponsePhaseHTTPRequestFirewallManaged, RuleDeleteResponsePhaseHTTPRequestLateTransform, RuleDeleteResponsePhaseHTTPRequestOrigin, RuleDeleteResponsePhaseHTTPRequestRedirect, RuleDeleteResponsePhaseHTTPRequestSanitize, RuleDeleteResponsePhaseHTTPRequestSbfm, RuleDeleteResponsePhaseHTTPRequestSelectConfiguration, RuleDeleteResponsePhaseHTTPRequestTransform, RuleDeleteResponsePhaseHTTPResponseCompression, RuleDeleteResponsePhaseHTTPResponseFirewallManaged, RuleDeleteResponsePhaseHTTPResponseHeadersTransform, RuleDeleteResponsePhaseMagicTransit, RuleDeleteResponsePhaseMagicTransitIDsManaged, RuleDeleteResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RuleDeleteResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleDeleteResponseRulesAction `json:"action"`
	ActionParameters interface{}                   `json:"action_parameters,required"`
	Categories       interface{}                   `json:"categories,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                     `json:"version,required"`
	JSON    ruleDeleteResponseRuleJSON `json:"-"`
	union   RuleDeleteResponseRulesUnion
}

// ruleDeleteResponseRuleJSON contains the JSON metadata for the struct
// [RuleDeleteResponseRule]
type ruleDeleteResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	ID               apijson.Field
	LastUpdated      apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r ruleDeleteResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleDeleteResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleDeleteResponseRule) AsUnion() RuleDeleteResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ExecuteRule],
// [rulesets.LogRule] or [rulesets.SkipRule].
type RuleDeleteResponseRulesUnion interface {
	implementsRulesetsRuleDeleteResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesAction string

const (
	RuleDeleteResponseRulesActionBlock   RuleDeleteResponseRulesAction = "block"
	RuleDeleteResponseRulesActionExecute RuleDeleteResponseRulesAction = "execute"
	RuleDeleteResponseRulesActionLog     RuleDeleteResponseRulesAction = "log"
	RuleDeleteResponseRulesActionSkip    RuleDeleteResponseRulesAction = "skip"
)

func (r RuleDeleteResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesActionBlock, RuleDeleteResponseRulesActionExecute, RuleDeleteResponseRulesActionLog, RuleDeleteResponseRulesActionSkip:
		return true
	}
	return false
}

// A ruleset object.
type RuleEditResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleEditResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleEditResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleEditResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleEditResponseKind string

const (
	RuleEditResponseKindManaged RuleEditResponseKind = "managed"
	RuleEditResponseKindCustom  RuleEditResponseKind = "custom"
	RuleEditResponseKindRoot    RuleEditResponseKind = "root"
	RuleEditResponseKindZone    RuleEditResponseKind = "zone"
)

func (r RuleEditResponseKind) IsKnown() bool {
	switch r {
	case RuleEditResponseKindManaged, RuleEditResponseKindCustom, RuleEditResponseKindRoot, RuleEditResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleEditResponsePhase string

const (
	RuleEditResponsePhaseDDoSL4                         RuleEditResponsePhase = "ddos_l4"
	RuleEditResponsePhaseDDoSL7                         RuleEditResponsePhase = "ddos_l7"
	RuleEditResponsePhaseHTTPConfigSettings             RuleEditResponsePhase = "http_config_settings"
	RuleEditResponsePhaseHTTPCustomErrors               RuleEditResponsePhase = "http_custom_errors"
	RuleEditResponsePhaseHTTPLogCustomFields            RuleEditResponsePhase = "http_log_custom_fields"
	RuleEditResponsePhaseHTTPRatelimit                  RuleEditResponsePhase = "http_ratelimit"
	RuleEditResponsePhaseHTTPRequestCacheSettings       RuleEditResponsePhase = "http_request_cache_settings"
	RuleEditResponsePhaseHTTPRequestDynamicRedirect     RuleEditResponsePhase = "http_request_dynamic_redirect"
	RuleEditResponsePhaseHTTPRequestFirewallCustom      RuleEditResponsePhase = "http_request_firewall_custom"
	RuleEditResponsePhaseHTTPRequestFirewallManaged     RuleEditResponsePhase = "http_request_firewall_managed"
	RuleEditResponsePhaseHTTPRequestLateTransform       RuleEditResponsePhase = "http_request_late_transform"
	RuleEditResponsePhaseHTTPRequestOrigin              RuleEditResponsePhase = "http_request_origin"
	RuleEditResponsePhaseHTTPRequestRedirect            RuleEditResponsePhase = "http_request_redirect"
	RuleEditResponsePhaseHTTPRequestSanitize            RuleEditResponsePhase = "http_request_sanitize"
	RuleEditResponsePhaseHTTPRequestSbfm                RuleEditResponsePhase = "http_request_sbfm"
	RuleEditResponsePhaseHTTPRequestSelectConfiguration RuleEditResponsePhase = "http_request_select_configuration"
	RuleEditResponsePhaseHTTPRequestTransform           RuleEditResponsePhase = "http_request_transform"
	RuleEditResponsePhaseHTTPResponseCompression        RuleEditResponsePhase = "http_response_compression"
	RuleEditResponsePhaseHTTPResponseFirewallManaged    RuleEditResponsePhase = "http_response_firewall_managed"
	RuleEditResponsePhaseHTTPResponseHeadersTransform   RuleEditResponsePhase = "http_response_headers_transform"
	RuleEditResponsePhaseMagicTransit                   RuleEditResponsePhase = "magic_transit"
	RuleEditResponsePhaseMagicTransitIDsManaged         RuleEditResponsePhase = "magic_transit_ids_managed"
	RuleEditResponsePhaseMagicTransitManaged            RuleEditResponsePhase = "magic_transit_managed"
)

func (r RuleEditResponsePhase) IsKnown() bool {
	switch r {
	case RuleEditResponsePhaseDDoSL4, RuleEditResponsePhaseDDoSL7, RuleEditResponsePhaseHTTPConfigSettings, RuleEditResponsePhaseHTTPCustomErrors, RuleEditResponsePhaseHTTPLogCustomFields, RuleEditResponsePhaseHTTPRatelimit, RuleEditResponsePhaseHTTPRequestCacheSettings, RuleEditResponsePhaseHTTPRequestDynamicRedirect, RuleEditResponsePhaseHTTPRequestFirewallCustom, RuleEditResponsePhaseHTTPRequestFirewallManaged, RuleEditResponsePhaseHTTPRequestLateTransform, RuleEditResponsePhaseHTTPRequestOrigin, RuleEditResponsePhaseHTTPRequestRedirect, RuleEditResponsePhaseHTTPRequestSanitize, RuleEditResponsePhaseHTTPRequestSbfm, RuleEditResponsePhaseHTTPRequestSelectConfiguration, RuleEditResponsePhaseHTTPRequestTransform, RuleEditResponsePhaseHTTPResponseCompression, RuleEditResponsePhaseHTTPResponseFirewallManaged, RuleEditResponsePhaseHTTPResponseHeadersTransform, RuleEditResponsePhaseMagicTransit, RuleEditResponsePhaseMagicTransitIDsManaged, RuleEditResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RuleEditResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleEditResponseRulesAction `json:"action"`
	ActionParameters interface{}                 `json:"action_parameters,required"`
	Categories       interface{}                 `json:"categories,required"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751c `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                   `json:"version,required"`
	JSON    ruleEditResponseRuleJSON `json:"-"`
	union   RuleEditResponseRulesUnion
}

// ruleEditResponseRuleJSON contains the JSON metadata for the struct
// [RuleEditResponseRule]
type ruleEditResponseRuleJSON struct {
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	ID               apijson.Field
	LastUpdated      apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r ruleEditResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleEditResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleEditResponseRule) AsUnion() RuleEditResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ExecuteRule],
// [rulesets.LogRule] or [rulesets.SkipRule].
type RuleEditResponseRulesUnion interface {
	implementsRulesetsRuleEditResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

// The action to perform when the rule matches.
type RuleEditResponseRulesAction string

const (
	RuleEditResponseRulesActionBlock   RuleEditResponseRulesAction = "block"
	RuleEditResponseRulesActionExecute RuleEditResponseRulesAction = "execute"
	RuleEditResponseRulesActionLog     RuleEditResponseRulesAction = "log"
	RuleEditResponseRulesActionSkip    RuleEditResponseRulesAction = "skip"
)

func (r RuleEditResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesActionBlock, RuleEditResponseRulesActionExecute, RuleEditResponseRulesActionLog, RuleEditResponseRulesActionSkip:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [RuleNewParamsBlockRule], [RuleNewParamsExecuteRule], [RuleNewParamsLogRule],
// [RuleNewParamsSkipRule].
type RuleNewParams interface {
	ImplementsRuleNewParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type RuleNewParamsBlockRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsBlockRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsBlockRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsBlockRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsBlockRuleAction string

const (
	RuleNewParamsBlockRuleActionBlock RuleNewParamsBlockRuleAction = "block"
)

func (r RuleNewParamsBlockRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RuleNewParamsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RuleNewParamsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RuleNewParamsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RuleNewParamsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsExecuteRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsExecuteRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsExecuteRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsExecuteRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsExecuteRuleAction string

const (
	RuleNewParamsExecuteRuleActionExecute RuleNewParamsExecuteRuleAction = "execute"
)

func (r RuleNewParamsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RuleNewParamsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RuleNewParamsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RuleNewParamsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RuleNewParamsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RuleNewParamsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RuleNewParamsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RuleNewParamsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RuleNewParamsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RuleNewParamsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RuleNewParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type RuleNewParamsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleNewParamsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RuleNewParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelLow, RuleNewParamsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type RuleNewParamsLogRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsLogRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsLogRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsLogRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsLogRuleAction string

const (
	RuleNewParamsLogRuleActionLog RuleNewParamsLogRuleAction = "log"
)

func (r RuleNewParamsLogRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsLogRuleActionLog:
		return true
	}
	return false
}

type RuleNewParamsSkipRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleNewParamsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleNewParamsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleNewParamsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsSkipRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleNewParamsSkipRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleNewParamsSkipRule) ImplementsRuleNewParams() {

}

// The action to perform when the rule matches.
type RuleNewParamsSkipRuleAction string

const (
	RuleNewParamsSkipRuleActionSkip RuleNewParamsSkipRuleAction = "skip"
)

func (r RuleNewParamsSkipRuleAction) IsKnown() bool {
	switch r {
	case RuleNewParamsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewParamsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RuleNewParamsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RuleNewParamsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RuleNewParamsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RuleNewParamsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RuleNewParamsSkipRuleActionParametersPhase string

const (
	RuleNewParamsSkipRuleActionParametersPhaseDDoSL4                         RuleNewParamsSkipRuleActionParametersPhase = "ddos_l4"
	RuleNewParamsSkipRuleActionParametersPhaseDDoSL7                         RuleNewParamsSkipRuleActionParametersPhase = "ddos_l7"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleNewParamsSkipRuleActionParametersPhase = "http_config_settings"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleNewParamsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleNewParamsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleNewParamsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleNewParamsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleNewParamsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleNewParamsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleNewParamsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleNewParamsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleNewParamsSkipRuleActionParametersPhase = "http_request_origin"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleNewParamsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleNewParamsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleNewParamsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleNewParamsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleNewParamsSkipRuleActionParametersPhase = "http_request_transform"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleNewParamsSkipRuleActionParametersPhase = "http_response_compression"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleNewParamsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleNewParamsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleNewParamsSkipRuleActionParametersPhaseMagicTransit                   RuleNewParamsSkipRuleActionParametersPhase = "magic_transit"
	RuleNewParamsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleNewParamsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleNewParamsSkipRuleActionParametersPhaseMagicTransitManaged            RuleNewParamsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r RuleNewParamsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RuleNewParamsSkipRuleActionParametersPhaseDDoSL4, RuleNewParamsSkipRuleActionParametersPhaseDDoSL7, RuleNewParamsSkipRuleActionParametersPhaseHTTPConfigSettings, RuleNewParamsSkipRuleActionParametersPhaseHTTPCustomErrors, RuleNewParamsSkipRuleActionParametersPhaseHTTPLogCustomFields, RuleNewParamsSkipRuleActionParametersPhaseHTTPRatelimit, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestOrigin, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestRedirect, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSanitize, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSbfm, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RuleNewParamsSkipRuleActionParametersPhaseHTTPRequestTransform, RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseCompression, RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RuleNewParamsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RuleNewParamsSkipRuleActionParametersPhaseMagicTransit, RuleNewParamsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RuleNewParamsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type RuleNewParamsSkipRuleActionParametersProduct string

const (
	RuleNewParamsSkipRuleActionParametersProductBic           RuleNewParamsSkipRuleActionParametersProduct = "bic"
	RuleNewParamsSkipRuleActionParametersProductHot           RuleNewParamsSkipRuleActionParametersProduct = "hot"
	RuleNewParamsSkipRuleActionParametersProductRateLimit     RuleNewParamsSkipRuleActionParametersProduct = "rateLimit"
	RuleNewParamsSkipRuleActionParametersProductSecurityLevel RuleNewParamsSkipRuleActionParametersProduct = "securityLevel"
	RuleNewParamsSkipRuleActionParametersProductUABlock       RuleNewParamsSkipRuleActionParametersProduct = "uaBlock"
	RuleNewParamsSkipRuleActionParametersProductWAF           RuleNewParamsSkipRuleActionParametersProduct = "waf"
	RuleNewParamsSkipRuleActionParametersProductZoneLockdown  RuleNewParamsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r RuleNewParamsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RuleNewParamsSkipRuleActionParametersProductBic, RuleNewParamsSkipRuleActionParametersProductHot, RuleNewParamsSkipRuleActionParametersProductRateLimit, RuleNewParamsSkipRuleActionParametersProductSecurityLevel, RuleNewParamsSkipRuleActionParametersProductUABlock, RuleNewParamsSkipRuleActionParametersProductWAF, RuleNewParamsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleNewParamsSkipRuleActionParametersRuleset string

const (
	RuleNewParamsSkipRuleActionParametersRulesetCurrent RuleNewParamsSkipRuleActionParametersRuleset = "current"
)

func (r RuleNewParamsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RuleNewParamsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                  `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeErrorsSource]
type ruleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeMessagesSource]
type ruleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

type RuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
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

// A message.
type RuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeErrorsSource]
type ruleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeMessagesSource]
type ruleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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

// This interface is a union satisfied by one of the following:
// [RuleEditParamsBlockRule], [RuleEditParamsExecuteRule], [RuleEditParamsLogRule],
// [RuleEditParamsSkipRule].
type RuleEditParams interface {
	ImplementsRuleEditParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type RuleEditParamsBlockRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsBlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsBlockRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsBlockRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsBlockRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsBlockRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsBlockRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsBlockRuleAction string

const (
	RuleEditParamsBlockRuleActionBlock RuleEditParamsBlockRuleAction = "block"
)

func (r RuleEditParamsBlockRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[RuleEditParamsBlockRuleActionParametersResponse] `json:"response"`
}

func (r RuleEditParamsBlockRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type RuleEditParamsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r RuleEditParamsBlockRuleActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsExecuteRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsExecuteRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsExecuteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsExecuteRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsExecuteRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsExecuteRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsExecuteRuleAction string

const (
	RuleEditParamsExecuteRuleActionExecute RuleEditParamsExecuteRuleAction = "execute"
)

func (r RuleEditParamsExecuteRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[RuleEditParamsExecuteRuleActionParametersMatchedData] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[RuleEditParamsExecuteRuleActionParametersOverrides] `json:"overrides"`
}

func (r RuleEditParamsExecuteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type RuleEditParamsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r RuleEditParamsExecuteRuleActionParametersMatchedData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type RuleEditParamsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]RuleEditParamsExecuteRuleActionParametersOverridesCategory] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]RuleEditParamsExecuteRuleActionParametersOverridesRule] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsExecuteRuleActionParametersOverrides) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type RuleEditParamsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsExecuteRuleActionParametersOverridesCategory) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for rules in the category.
type RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, RuleEditParamsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type RuleEditParamsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r RuleEditParamsExecuteRuleActionParametersOverridesRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The sensitivity level to use for the rule.
type RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, RuleEditParamsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelDefault RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelLow     RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelDefault, RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelMedium, RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelLow, RuleEditParamsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type RuleEditParamsLogRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsLogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsLogRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsLogRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsLogRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsLogRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsLogRuleAction string

const (
	RuleEditParamsLogRuleActionLog RuleEditParamsLogRuleAction = "log"
)

func (r RuleEditParamsLogRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsLogRuleActionLog:
		return true
	}
	return false
}

type RuleEditParamsSkipRule struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RuleEditParamsSkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RuleEditParamsSkipRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RuleEditParamsSkipRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsSkipRule) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r RuleEditParamsSkipRule) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (RuleEditParamsSkipRule) ImplementsRuleEditParams() {

}

// The action to perform when the rule matches.
type RuleEditParamsSkipRuleAction string

const (
	RuleEditParamsSkipRuleActionSkip RuleEditParamsSkipRuleAction = "skip"
)

func (r RuleEditParamsSkipRuleAction) IsKnown() bool {
	switch r {
	case RuleEditParamsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditParamsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]RuleEditParamsSkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]RuleEditParamsSkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[RuleEditParamsSkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r RuleEditParamsSkipRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A phase to skip the execution of.
type RuleEditParamsSkipRuleActionParametersPhase string

const (
	RuleEditParamsSkipRuleActionParametersPhaseDDoSL4                         RuleEditParamsSkipRuleActionParametersPhase = "ddos_l4"
	RuleEditParamsSkipRuleActionParametersPhaseDDoSL7                         RuleEditParamsSkipRuleActionParametersPhase = "ddos_l7"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPConfigSettings             RuleEditParamsSkipRuleActionParametersPhase = "http_config_settings"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPCustomErrors               RuleEditParamsSkipRuleActionParametersPhase = "http_custom_errors"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPLogCustomFields            RuleEditParamsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRatelimit                  RuleEditParamsSkipRuleActionParametersPhase = "http_ratelimit"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RuleEditParamsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RuleEditParamsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RuleEditParamsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RuleEditParamsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RuleEditParamsSkipRuleActionParametersPhase = "http_request_late_transform"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestOrigin              RuleEditParamsSkipRuleActionParametersPhase = "http_request_origin"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestRedirect            RuleEditParamsSkipRuleActionParametersPhase = "http_request_redirect"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSanitize            RuleEditParamsSkipRuleActionParametersPhase = "http_request_sanitize"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSbfm                RuleEditParamsSkipRuleActionParametersPhase = "http_request_sbfm"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RuleEditParamsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestTransform           RuleEditParamsSkipRuleActionParametersPhase = "http_request_transform"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseCompression        RuleEditParamsSkipRuleActionParametersPhase = "http_response_compression"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RuleEditParamsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RuleEditParamsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RuleEditParamsSkipRuleActionParametersPhaseMagicTransit                   RuleEditParamsSkipRuleActionParametersPhase = "magic_transit"
	RuleEditParamsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RuleEditParamsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RuleEditParamsSkipRuleActionParametersPhaseMagicTransitManaged            RuleEditParamsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r RuleEditParamsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case RuleEditParamsSkipRuleActionParametersPhaseDDoSL4, RuleEditParamsSkipRuleActionParametersPhaseDDoSL7, RuleEditParamsSkipRuleActionParametersPhaseHTTPConfigSettings, RuleEditParamsSkipRuleActionParametersPhaseHTTPCustomErrors, RuleEditParamsSkipRuleActionParametersPhaseHTTPLogCustomFields, RuleEditParamsSkipRuleActionParametersPhaseHTTPRatelimit, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestLateTransform, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestOrigin, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestRedirect, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSanitize, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSbfm, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, RuleEditParamsSkipRuleActionParametersPhaseHTTPRequestTransform, RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseCompression, RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, RuleEditParamsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, RuleEditParamsSkipRuleActionParametersPhaseMagicTransit, RuleEditParamsSkipRuleActionParametersPhaseMagicTransitIDsManaged, RuleEditParamsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type RuleEditParamsSkipRuleActionParametersProduct string

const (
	RuleEditParamsSkipRuleActionParametersProductBic           RuleEditParamsSkipRuleActionParametersProduct = "bic"
	RuleEditParamsSkipRuleActionParametersProductHot           RuleEditParamsSkipRuleActionParametersProduct = "hot"
	RuleEditParamsSkipRuleActionParametersProductRateLimit     RuleEditParamsSkipRuleActionParametersProduct = "rateLimit"
	RuleEditParamsSkipRuleActionParametersProductSecurityLevel RuleEditParamsSkipRuleActionParametersProduct = "securityLevel"
	RuleEditParamsSkipRuleActionParametersProductUABlock       RuleEditParamsSkipRuleActionParametersProduct = "uaBlock"
	RuleEditParamsSkipRuleActionParametersProductWAF           RuleEditParamsSkipRuleActionParametersProduct = "waf"
	RuleEditParamsSkipRuleActionParametersProductZoneLockdown  RuleEditParamsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r RuleEditParamsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case RuleEditParamsSkipRuleActionParametersProductBic, RuleEditParamsSkipRuleActionParametersProductHot, RuleEditParamsSkipRuleActionParametersProductRateLimit, RuleEditParamsSkipRuleActionParametersProductSecurityLevel, RuleEditParamsSkipRuleActionParametersProductUABlock, RuleEditParamsSkipRuleActionParametersProductWAF, RuleEditParamsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RuleEditParamsSkipRuleActionParametersRuleset string

const (
	RuleEditParamsSkipRuleActionParametersRulesetCurrent RuleEditParamsSkipRuleActionParametersRuleset = "current"
)

func (r RuleEditParamsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case RuleEditParamsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleEditResponseEnvelopeJSON    `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeErrorsSource]
type ruleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeMessagesSource]
type ruleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
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
