// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RulesetPhaseVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetPhaseVersionService]
// method instead.
type RulesetPhaseVersionService struct {
	Options []option.RequestOption
}

// NewRulesetPhaseVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRulesetPhaseVersionService(opts ...option.RequestOption) (r *RulesetPhaseVersionService) {
	r = &RulesetPhaseVersionService{}
	r.Options = opts
	return
}

// Fetches the versions of an account or zone entry point ruleset.
func (r *RulesetPhaseVersionService) List(ctx context.Context, rulesetPhase RulesetPhaseVersionListParamsRulesetPhase, query RulesetPhaseVersionListParams, opts ...option.RequestOption) (res *[]RulesetPhaseVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseVersionListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a specific version of an account or zone entry point ruleset.
func (r *RulesetPhaseVersionService) Get(ctx context.Context, rulesetPhase RulesetPhaseVersionGetParamsRulesetPhase, rulesetVersion string, query RulesetPhaseVersionGetParams, opts ...option.RequestOption) (res *RulesetPhaseVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetPhaseVersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions/%s", accountOrZone, accountOrZoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RulesetPhaseVersionListResponse struct {
	// The kind of the ruleset.
	Kind RulesetPhaseVersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetPhaseVersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                              `json:"version"`
	JSON    rulesetPhaseVersionListResponseJSON `json:"-"`
}

// rulesetPhaseVersionListResponseJSON contains the JSON metadata for the struct
// [RulesetPhaseVersionListResponse]
type rulesetPhaseVersionListResponseJSON struct {
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	ID          apijson.Field
	Description apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetPhaseVersionListResponseKind string

const (
	RulesetPhaseVersionListResponseKindManaged RulesetPhaseVersionListResponseKind = "managed"
	RulesetPhaseVersionListResponseKindCustom  RulesetPhaseVersionListResponseKind = "custom"
	RulesetPhaseVersionListResponseKindRoot    RulesetPhaseVersionListResponseKind = "root"
	RulesetPhaseVersionListResponseKindZone    RulesetPhaseVersionListResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseVersionListResponsePhase string

const (
	RulesetPhaseVersionListResponsePhaseDDOSL4                         RulesetPhaseVersionListResponsePhase = "ddos_l4"
	RulesetPhaseVersionListResponsePhaseDDOSL7                         RulesetPhaseVersionListResponsePhase = "ddos_l7"
	RulesetPhaseVersionListResponsePhaseHTTPConfigSettings             RulesetPhaseVersionListResponsePhase = "http_config_settings"
	RulesetPhaseVersionListResponsePhaseHTTPCustomErrors               RulesetPhaseVersionListResponsePhase = "http_custom_errors"
	RulesetPhaseVersionListResponsePhaseHTTPLogCustomFields            RulesetPhaseVersionListResponsePhase = "http_log_custom_fields"
	RulesetPhaseVersionListResponsePhaseHTTPRatelimit                  RulesetPhaseVersionListResponsePhase = "http_ratelimit"
	RulesetPhaseVersionListResponsePhaseHTTPRequestCacheSettings       RulesetPhaseVersionListResponsePhase = "http_request_cache_settings"
	RulesetPhaseVersionListResponsePhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionListResponsePhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionListResponsePhaseHTTPRequestFirewallCustom      RulesetPhaseVersionListResponsePhase = "http_request_firewall_custom"
	RulesetPhaseVersionListResponsePhaseHTTPRequestFirewallManaged     RulesetPhaseVersionListResponsePhase = "http_request_firewall_managed"
	RulesetPhaseVersionListResponsePhaseHTTPRequestLateTransform       RulesetPhaseVersionListResponsePhase = "http_request_late_transform"
	RulesetPhaseVersionListResponsePhaseHTTPRequestOrigin              RulesetPhaseVersionListResponsePhase = "http_request_origin"
	RulesetPhaseVersionListResponsePhaseHTTPRequestRedirect            RulesetPhaseVersionListResponsePhase = "http_request_redirect"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSanitize            RulesetPhaseVersionListResponsePhase = "http_request_sanitize"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSbfm                RulesetPhaseVersionListResponsePhase = "http_request_sbfm"
	RulesetPhaseVersionListResponsePhaseHTTPRequestSelectConfiguration RulesetPhaseVersionListResponsePhase = "http_request_select_configuration"
	RulesetPhaseVersionListResponsePhaseHTTPRequestTransform           RulesetPhaseVersionListResponsePhase = "http_request_transform"
	RulesetPhaseVersionListResponsePhaseHTTPResponseCompression        RulesetPhaseVersionListResponsePhase = "http_response_compression"
	RulesetPhaseVersionListResponsePhaseHTTPResponseFirewallManaged    RulesetPhaseVersionListResponsePhase = "http_response_firewall_managed"
	RulesetPhaseVersionListResponsePhaseHTTPResponseHeadersTransform   RulesetPhaseVersionListResponsePhase = "http_response_headers_transform"
	RulesetPhaseVersionListResponsePhaseMagicTransit                   RulesetPhaseVersionListResponsePhase = "magic_transit"
	RulesetPhaseVersionListResponsePhaseMagicTransitIDsManaged         RulesetPhaseVersionListResponsePhase = "magic_transit_ids_managed"
	RulesetPhaseVersionListResponsePhaseMagicTransitManaged            RulesetPhaseVersionListResponsePhase = "magic_transit_managed"
)

// A result.
type RulesetPhaseVersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetPhaseVersionGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetPhaseVersionGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetPhaseVersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                             `json:"description"`
	JSON        rulesetPhaseVersionGetResponseJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseJSON contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponse]
type rulesetPhaseVersionGetResponseJSON struct {
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

func (r *RulesetPhaseVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetPhaseVersionGetResponseKind string

const (
	RulesetPhaseVersionGetResponseKindManaged RulesetPhaseVersionGetResponseKind = "managed"
	RulesetPhaseVersionGetResponseKindCustom  RulesetPhaseVersionGetResponseKind = "custom"
	RulesetPhaseVersionGetResponseKindRoot    RulesetPhaseVersionGetResponseKind = "root"
	RulesetPhaseVersionGetResponseKindZone    RulesetPhaseVersionGetResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetPhaseVersionGetResponsePhase string

const (
	RulesetPhaseVersionGetResponsePhaseDDOSL4                         RulesetPhaseVersionGetResponsePhase = "ddos_l4"
	RulesetPhaseVersionGetResponsePhaseDDOSL7                         RulesetPhaseVersionGetResponsePhase = "ddos_l7"
	RulesetPhaseVersionGetResponsePhaseHTTPConfigSettings             RulesetPhaseVersionGetResponsePhase = "http_config_settings"
	RulesetPhaseVersionGetResponsePhaseHTTPCustomErrors               RulesetPhaseVersionGetResponsePhase = "http_custom_errors"
	RulesetPhaseVersionGetResponsePhaseHTTPLogCustomFields            RulesetPhaseVersionGetResponsePhase = "http_log_custom_fields"
	RulesetPhaseVersionGetResponsePhaseHTTPRatelimit                  RulesetPhaseVersionGetResponsePhase = "http_ratelimit"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestCacheSettings       RulesetPhaseVersionGetResponsePhase = "http_request_cache_settings"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionGetResponsePhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestFirewallCustom      RulesetPhaseVersionGetResponsePhase = "http_request_firewall_custom"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestFirewallManaged     RulesetPhaseVersionGetResponsePhase = "http_request_firewall_managed"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestLateTransform       RulesetPhaseVersionGetResponsePhase = "http_request_late_transform"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestOrigin              RulesetPhaseVersionGetResponsePhase = "http_request_origin"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestRedirect            RulesetPhaseVersionGetResponsePhase = "http_request_redirect"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestSanitize            RulesetPhaseVersionGetResponsePhase = "http_request_sanitize"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestSbfm                RulesetPhaseVersionGetResponsePhase = "http_request_sbfm"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestSelectConfiguration RulesetPhaseVersionGetResponsePhase = "http_request_select_configuration"
	RulesetPhaseVersionGetResponsePhaseHTTPRequestTransform           RulesetPhaseVersionGetResponsePhase = "http_request_transform"
	RulesetPhaseVersionGetResponsePhaseHTTPResponseCompression        RulesetPhaseVersionGetResponsePhase = "http_response_compression"
	RulesetPhaseVersionGetResponsePhaseHTTPResponseFirewallManaged    RulesetPhaseVersionGetResponsePhase = "http_response_firewall_managed"
	RulesetPhaseVersionGetResponsePhaseHTTPResponseHeadersTransform   RulesetPhaseVersionGetResponsePhase = "http_response_headers_transform"
	RulesetPhaseVersionGetResponsePhaseMagicTransit                   RulesetPhaseVersionGetResponsePhase = "magic_transit"
	RulesetPhaseVersionGetResponsePhaseMagicTransitIDsManaged         RulesetPhaseVersionGetResponsePhase = "magic_transit_ids_managed"
	RulesetPhaseVersionGetResponsePhaseMagicTransitManaged            RulesetPhaseVersionGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetPhaseVersionGetResponseRulesRulesetsBlockRule],
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule],
// [RulesetPhaseVersionGetResponseRulesRulesetsLogRule] or
// [RulesetPhaseVersionGetResponseRulesRulesetsSkipRule].
type RulesetPhaseVersionGetResponseRule interface {
	implementsRulesetPhaseVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetPhaseVersionGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseVersionGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseVersionGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetPhaseVersionGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetPhaseVersionGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseRulesRulesetsBlockRule]
type rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseVersionGetResponseRulesRulesetsBlockRule) implementsRulesetPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionBlock RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                            `json:"status_code,required"`
	JSON       rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                            `json:"enabled,required"`
	JSON    rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLoggingJSON contains the
// JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLogging]
type rulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                     `json:"ref"`
	JSON rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseVersionGetResponseRulesRulesetsExecuteRule) implementsRulesetPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionExecute RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                                `json:"public_key,required"`
	JSON      rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                              `json:"enabled,required"`
	JSON    rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the
// JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLogging]
type rulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseVersionGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseVersionGetResponseRulesRulesetsLogRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseVersionGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetPhaseVersionGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsLogRuleJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseRulesRulesetsLogRule]
type rulesetPhaseVersionGetResponseRulesRulesetsLogRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseVersionGetResponseRulesRulesetsLogRule) implementsRulesetPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseVersionGetResponseRulesRulesetsLogRuleAction string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsLogRuleActionLog RulesetPhaseVersionGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseVersionGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                          `json:"enabled,required"`
	JSON    rulesetPhaseVersionGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsLogRuleLogging]
type rulesetPhaseVersionGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseVersionGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                  `json:"ref"`
	JSON rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseRulesRulesetsSkipRule]
type rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleJSON struct {
	LastUpdated      apijson.Field
	Version          apijson.Field
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetPhaseVersionGetResponseRulesRulesetsSkipRule) implementsRulesetPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionSkip RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                                `json:"rulesets"`
	JSON     rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                           `json:"enabled,required"`
	JSON    rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLogging]
type rulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetPhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type RulesetPhaseVersionListParamsRulesetPhase string

const (
	RulesetPhaseVersionListParamsRulesetPhaseDDOSL4                         RulesetPhaseVersionListParamsRulesetPhase = "ddos_l4"
	RulesetPhaseVersionListParamsRulesetPhaseDDOSL7                         RulesetPhaseVersionListParamsRulesetPhase = "ddos_l7"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseVersionListParamsRulesetPhase = "http_config_settings"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseVersionListParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseVersionListParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseVersionListParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseVersionListParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionListParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseVersionListParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseVersionListParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseVersionListParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseVersionListParamsRulesetPhase = "http_request_origin"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseVersionListParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseVersionListParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseVersionListParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseVersionListParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseVersionListParamsRulesetPhase = "http_request_transform"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseVersionListParamsRulesetPhase = "http_response_compression"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseVersionListParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseVersionListParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseVersionListParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransit                   RulesetPhaseVersionListParamsRulesetPhase = "magic_transit"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseVersionListParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseVersionListParamsRulesetPhaseMagicTransitManaged            RulesetPhaseVersionListParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseVersionListResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseVersionListResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseVersionListResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result []RulesetPhaseVersionListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseVersionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetPhaseVersionListResponseEnvelope]
type rulesetPhaseVersionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionListResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseVersionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetPhaseVersionListResponseEnvelopeErrors]
type rulesetPhaseVersionListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                  `json:"pointer,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionListResponseEnvelopeErrorsSource]
type rulesetPhaseVersionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionListResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseVersionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetPhaseVersionListResponseEnvelopeMessages]
type rulesetPhaseVersionListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                    `json:"pointer,required"`
	JSON    rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionListResponseEnvelopeMessagesSource]
type rulesetPhaseVersionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseVersionListResponseEnvelopeSuccess bool

const (
	RulesetPhaseVersionListResponseEnvelopeSuccessTrue RulesetPhaseVersionListResponseEnvelopeSuccess = true
)

type RulesetPhaseVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type RulesetPhaseVersionGetParamsRulesetPhase string

const (
	RulesetPhaseVersionGetParamsRulesetPhaseDDOSL4                         RulesetPhaseVersionGetParamsRulesetPhase = "ddos_l4"
	RulesetPhaseVersionGetParamsRulesetPhaseDDOSL7                         RulesetPhaseVersionGetParamsRulesetPhase = "ddos_l7"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPConfigSettings             RulesetPhaseVersionGetParamsRulesetPhase = "http_config_settings"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPCustomErrors               RulesetPhaseVersionGetParamsRulesetPhase = "http_custom_errors"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPLogCustomFields            RulesetPhaseVersionGetParamsRulesetPhase = "http_log_custom_fields"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRatelimit                  RulesetPhaseVersionGetParamsRulesetPhase = "http_ratelimit"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestCacheSettings       RulesetPhaseVersionGetParamsRulesetPhase = "http_request_cache_settings"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestDynamicRedirect     RulesetPhaseVersionGetParamsRulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallCustom      RulesetPhaseVersionGetParamsRulesetPhase = "http_request_firewall_custom"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestFirewallManaged     RulesetPhaseVersionGetParamsRulesetPhase = "http_request_firewall_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestLateTransform       RulesetPhaseVersionGetParamsRulesetPhase = "http_request_late_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestOrigin              RulesetPhaseVersionGetParamsRulesetPhase = "http_request_origin"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestRedirect            RulesetPhaseVersionGetParamsRulesetPhase = "http_request_redirect"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSanitize            RulesetPhaseVersionGetParamsRulesetPhase = "http_request_sanitize"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSbfm                RulesetPhaseVersionGetParamsRulesetPhase = "http_request_sbfm"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestSelectConfiguration RulesetPhaseVersionGetParamsRulesetPhase = "http_request_select_configuration"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPRequestTransform           RulesetPhaseVersionGetParamsRulesetPhase = "http_request_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseCompression        RulesetPhaseVersionGetParamsRulesetPhase = "http_response_compression"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseFirewallManaged    RulesetPhaseVersionGetParamsRulesetPhase = "http_response_firewall_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseHTTPResponseHeadersTransform   RulesetPhaseVersionGetParamsRulesetPhase = "http_response_headers_transform"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransit                   RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransitIDsManaged         RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseVersionGetParamsRulesetPhaseMagicTransitManaged            RulesetPhaseVersionGetParamsRulesetPhase = "magic_transit_managed"
)

// A response object.
type RulesetPhaseVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetPhaseVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetPhaseVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetPhaseVersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetPhaseVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetPhaseVersionGetResponseEnvelope]
type rulesetPhaseVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetPhaseVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetPhaseVersionGetResponseEnvelopeErrors]
type rulesetPhaseVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseEnvelopeErrorsSource]
type rulesetPhaseVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetPhaseVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetPhaseVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetPhaseVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetPhaseVersionGetResponseEnvelopeMessages]
type rulesetPhaseVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetPhaseVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetPhaseVersionGetResponseEnvelopeMessagesSource]
type rulesetPhaseVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetPhaseVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetPhaseVersionGetResponseEnvelopeSuccess bool

const (
	RulesetPhaseVersionGetResponseEnvelopeSuccessTrue RulesetPhaseVersionGetResponseEnvelopeSuccess = true
)
