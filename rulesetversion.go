// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// RulesetVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetVersionService] method
// instead.
type RulesetVersionService struct {
	Options []option.RequestOption
	ByTags  *RulesetVersionByTagService
}

// NewRulesetVersionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetVersionService(opts ...option.RequestOption) (r *RulesetVersionService) {
	r = &RulesetVersionService{}
	r.Options = opts
	r.ByTags = NewRulesetVersionByTagService(opts...)
	return
}

// Deletes an existing version of an account or zone ruleset.
func (r *RulesetVersionService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *RulesetVersionService) AccountRulesetsListAnAccountRulesetSVersions(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, opts ...option.RequestOption) (res *[]RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *RulesetVersionService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, rulesetID string, rulesetVersion string, opts ...option.RequestOption) (res *RulesetVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse struct {
	// The kind of the ruleset.
	Kind RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                                                                 `json:"version"`
	JSON    rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON contains
// the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseJSON struct {
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

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind string

const (
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKindManaged RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind = "managed"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKindCustom  RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind = "custom"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKindRoot    RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind = "root"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKindZone    RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase string

const (
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseDDOSL4                         RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "ddos_l4"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseDDOSL7                         RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "ddos_l7"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPConfigSettings             RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_config_settings"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPCustomErrors               RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_custom_errors"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPLogCustomFields            RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_log_custom_fields"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRatelimit                  RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_ratelimit"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestCacheSettings       RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_cache_settings"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestDynamicRedirect     RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_dynamic_redirect"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestFirewallCustom      RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_firewall_custom"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestFirewallManaged     RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_firewall_managed"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestLateTransform       RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_late_transform"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestOrigin              RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_origin"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestRedirect            RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_redirect"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestSanitize            RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_sanitize"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestSbfm                RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_sbfm"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestSelectConfiguration RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_select_configuration"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPRequestTransform           RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_request_transform"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPResponseCompression        RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_response_compression"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPResponseFirewallManaged    RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_response_firewall_managed"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseHTTPResponseHeadersTransform   RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "http_response_headers_transform"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseMagicTransit                   RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "magic_transit"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseMagicTransitIDsManaged         RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "magic_transit_ids_managed"
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhaseMagicTransitManaged            RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponsePhase = "magic_transit_managed"
)

// A result.
type RulesetVersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RulesetVersionGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RulesetVersionGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetVersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                        `json:"description"`
	JSON        rulesetVersionGetResponseJSON `json:"-"`
}

// rulesetVersionGetResponseJSON contains the JSON metadata for the struct
// [RulesetVersionGetResponse]
type rulesetVersionGetResponseJSON struct {
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

func (r *RulesetVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetVersionGetResponseKind string

const (
	RulesetVersionGetResponseKindManaged RulesetVersionGetResponseKind = "managed"
	RulesetVersionGetResponseKindCustom  RulesetVersionGetResponseKind = "custom"
	RulesetVersionGetResponseKindRoot    RulesetVersionGetResponseKind = "root"
	RulesetVersionGetResponseKindZone    RulesetVersionGetResponseKind = "zone"
)

// The phase of the ruleset.
type RulesetVersionGetResponsePhase string

const (
	RulesetVersionGetResponsePhaseDDOSL4                         RulesetVersionGetResponsePhase = "ddos_l4"
	RulesetVersionGetResponsePhaseDDOSL7                         RulesetVersionGetResponsePhase = "ddos_l7"
	RulesetVersionGetResponsePhaseHTTPConfigSettings             RulesetVersionGetResponsePhase = "http_config_settings"
	RulesetVersionGetResponsePhaseHTTPCustomErrors               RulesetVersionGetResponsePhase = "http_custom_errors"
	RulesetVersionGetResponsePhaseHTTPLogCustomFields            RulesetVersionGetResponsePhase = "http_log_custom_fields"
	RulesetVersionGetResponsePhaseHTTPRatelimit                  RulesetVersionGetResponsePhase = "http_ratelimit"
	RulesetVersionGetResponsePhaseHTTPRequestCacheSettings       RulesetVersionGetResponsePhase = "http_request_cache_settings"
	RulesetVersionGetResponsePhaseHTTPRequestDynamicRedirect     RulesetVersionGetResponsePhase = "http_request_dynamic_redirect"
	RulesetVersionGetResponsePhaseHTTPRequestFirewallCustom      RulesetVersionGetResponsePhase = "http_request_firewall_custom"
	RulesetVersionGetResponsePhaseHTTPRequestFirewallManaged     RulesetVersionGetResponsePhase = "http_request_firewall_managed"
	RulesetVersionGetResponsePhaseHTTPRequestLateTransform       RulesetVersionGetResponsePhase = "http_request_late_transform"
	RulesetVersionGetResponsePhaseHTTPRequestOrigin              RulesetVersionGetResponsePhase = "http_request_origin"
	RulesetVersionGetResponsePhaseHTTPRequestRedirect            RulesetVersionGetResponsePhase = "http_request_redirect"
	RulesetVersionGetResponsePhaseHTTPRequestSanitize            RulesetVersionGetResponsePhase = "http_request_sanitize"
	RulesetVersionGetResponsePhaseHTTPRequestSbfm                RulesetVersionGetResponsePhase = "http_request_sbfm"
	RulesetVersionGetResponsePhaseHTTPRequestSelectConfiguration RulesetVersionGetResponsePhase = "http_request_select_configuration"
	RulesetVersionGetResponsePhaseHTTPRequestTransform           RulesetVersionGetResponsePhase = "http_request_transform"
	RulesetVersionGetResponsePhaseHTTPResponseCompression        RulesetVersionGetResponsePhase = "http_response_compression"
	RulesetVersionGetResponsePhaseHTTPResponseFirewallManaged    RulesetVersionGetResponsePhase = "http_response_firewall_managed"
	RulesetVersionGetResponsePhaseHTTPResponseHeadersTransform   RulesetVersionGetResponsePhase = "http_response_headers_transform"
	RulesetVersionGetResponsePhaseMagicTransit                   RulesetVersionGetResponsePhase = "magic_transit"
	RulesetVersionGetResponsePhaseMagicTransitIDsManaged         RulesetVersionGetResponsePhase = "magic_transit_ids_managed"
	RulesetVersionGetResponsePhaseMagicTransitManaged            RulesetVersionGetResponsePhase = "magic_transit_managed"
)

// Union satisfied by [RulesetVersionGetResponseRulesRulesetsBlockRule],
// [RulesetVersionGetResponseRulesRulesetsExecuteRule],
// [RulesetVersionGetResponseRulesRulesetsLogRule] or
// [RulesetVersionGetResponseRulesRulesetsSkipRule].
type RulesetVersionGetResponseRule interface {
	implementsRulesetVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetVersionGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetVersionGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type RulesetVersionGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetVersionGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata
// for the struct [RulesetVersionGetResponseRulesRulesetsBlockRule]
type rulesetVersionGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *RulesetVersionGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetVersionGetResponseRulesRulesetsBlockRule) implementsRulesetVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetVersionGetResponseRulesRulesetsBlockRuleAction string

const (
	RulesetVersionGetResponseRulesRulesetsBlockRuleActionBlock RulesetVersionGetResponseRulesRulesetsBlockRuleAction = "block"
)

// The parameters configuring the rule's action.
type RulesetVersionGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response RulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsBlockRuleActionParameters]
type rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response to show when the block is applied.
type RulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                       `json:"status_code,required"`
	JSON       rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON
// contains the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type rulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetVersionGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                       `json:"enabled,required"`
	JSON    rulesetVersionGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetVersionGetResponseRulesRulesetsBlockRuleLogging]
type rulesetVersionGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetVersionGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetVersionGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata
// for the struct [RulesetVersionGetResponseRulesRulesetsExecuteRule]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetVersionGetResponseRulesRulesetsExecuteRule) implementsRulesetVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleAction string

const (
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionExecute RulesetVersionGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

// The parameters configuring the rule's action.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParameters]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration to use for matched data logging.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                           `json:"public_key,required"`
	JSON      rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A set of overrides to apply to the target ruleset.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON
// contains the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A category-level override
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for rules in the category.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

// A rule-level override
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The sensitivity level to use for the rule.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    RulesetVersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

// An object configuring the rule's logging behavior.
type RulesetVersionGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                         `json:"enabled,required"`
	JSON    rulesetVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON
// metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsExecuteRuleLogging]
type rulesetVersionGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetVersionGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging RulesetVersionGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                            `json:"ref"`
	JSON rulesetVersionGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for
// the struct [RulesetVersionGetResponseRulesRulesetsLogRule]
type rulesetVersionGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *RulesetVersionGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetVersionGetResponseRulesRulesetsLogRule) implementsRulesetVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetVersionGetResponseRulesRulesetsLogRuleAction string

const (
	RulesetVersionGetResponseRulesRulesetsLogRuleActionLog RulesetVersionGetResponseRulesRulesetsLogRuleAction = "log"
)

// An object configuring the rule's logging behavior.
type RulesetVersionGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                     `json:"enabled,required"`
	JSON    rulesetVersionGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetVersionGetResponseRulesRulesetsLogRuleLogging]
type rulesetVersionGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetVersionGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetVersionGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetVersionGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging RulesetVersionGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetVersionGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata
// for the struct [RulesetVersionGetResponseRulesRulesetsSkipRule]
type rulesetVersionGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *RulesetVersionGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetVersionGetResponseRulesRulesetsSkipRule) implementsRulesetVersionGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetVersionGetResponseRulesRulesetsSkipRuleAction string

const (
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionSkip RulesetVersionGetResponseRulesRulesetsSkipRuleAction = "skip"
)

// The parameters configuring the rule's action.
type RulesetVersionGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                           `json:"rulesets"`
	JSON     rulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetVersionGetResponseRulesRulesetsSkipRuleActionParameters]
type rulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A phase to skip the execution of.
type RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL4                         RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDOSL7                         RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

// The name of a legacy security product to skip the execution of.
type RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic           RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot           RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductUaBlock       RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent RulesetVersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

// An object configuring the rule's logging behavior.
type RulesetVersionGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                      `json:"enabled,required"`
	JSON    rulesetVersionGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// rulesetVersionGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON
// metadata for the struct [RulesetVersionGetResponseRulesRulesetsSkipRuleLogging]
type rulesetVersionGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A response object.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result []RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelope]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrors]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                     `json:"pointer,required"`
	JSON    rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSourceJSON
// contains the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSource]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessages]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                                                       `json:"pointer,required"`
	JSON    rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSourceJSON
// contains the JSON metadata for the struct
// [RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSource]
type rulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeSuccess bool

const (
	RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeSuccessTrue RulesetVersionAccountRulesetsListAnAccountRulesetSVersionsResponseEnvelopeSuccess = true
)

// A response object.
type RulesetVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetVersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetVersionGetResponseEnvelope]
type rulesetVersionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetVersionGetResponseEnvelopeErrors]
type rulesetVersionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetVersionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetVersionGetResponseEnvelopeErrorsSource]
type rulesetVersionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetVersionGetResponseEnvelopeMessages]
type rulesetVersionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    rulesetVersionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetVersionGetResponseEnvelopeMessagesSource]
type rulesetVersionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetVersionGetResponseEnvelopeSuccess bool

const (
	RulesetVersionGetResponseEnvelopeSuccessTrue RulesetVersionGetResponseEnvelopeSuccess = true
)
