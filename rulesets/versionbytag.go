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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// VersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVersionByTagService] method
// instead.
type VersionByTagService struct {
	Options []option.RequestOption
}

// NewVersionByTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVersionByTagService(opts ...option.RequestOption) (r *VersionByTagService) {
	r = &VersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *VersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query VersionByTagGetParams, opts ...option.RequestOption) (res *VersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionByTagGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind VersionByTagGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionByTagGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []VersionByTagGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        versionByTagGetResponseJSON `json:"-"`
}

// versionByTagGetResponseJSON contains the JSON metadata for the struct
// [VersionByTagGetResponse]
type versionByTagGetResponseJSON struct {
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

func (r *VersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionByTagGetResponseKind string

const (
	VersionByTagGetResponseKindManaged VersionByTagGetResponseKind = "managed"
	VersionByTagGetResponseKindCustom  VersionByTagGetResponseKind = "custom"
	VersionByTagGetResponseKindRoot    VersionByTagGetResponseKind = "root"
	VersionByTagGetResponseKindZone    VersionByTagGetResponseKind = "zone"
)

func (r VersionByTagGetResponseKind) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseKindManaged, VersionByTagGetResponseKindCustom, VersionByTagGetResponseKindRoot, VersionByTagGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionByTagGetResponsePhase string

const (
	VersionByTagGetResponsePhaseDDoSL4                         VersionByTagGetResponsePhase = "ddos_l4"
	VersionByTagGetResponsePhaseDDoSL7                         VersionByTagGetResponsePhase = "ddos_l7"
	VersionByTagGetResponsePhaseHTTPConfigSettings             VersionByTagGetResponsePhase = "http_config_settings"
	VersionByTagGetResponsePhaseHTTPCustomErrors               VersionByTagGetResponsePhase = "http_custom_errors"
	VersionByTagGetResponsePhaseHTTPLogCustomFields            VersionByTagGetResponsePhase = "http_log_custom_fields"
	VersionByTagGetResponsePhaseHTTPRatelimit                  VersionByTagGetResponsePhase = "http_ratelimit"
	VersionByTagGetResponsePhaseHTTPRequestCacheSettings       VersionByTagGetResponsePhase = "http_request_cache_settings"
	VersionByTagGetResponsePhaseHTTPRequestDynamicRedirect     VersionByTagGetResponsePhase = "http_request_dynamic_redirect"
	VersionByTagGetResponsePhaseHTTPRequestFirewallCustom      VersionByTagGetResponsePhase = "http_request_firewall_custom"
	VersionByTagGetResponsePhaseHTTPRequestFirewallManaged     VersionByTagGetResponsePhase = "http_request_firewall_managed"
	VersionByTagGetResponsePhaseHTTPRequestLateTransform       VersionByTagGetResponsePhase = "http_request_late_transform"
	VersionByTagGetResponsePhaseHTTPRequestOrigin              VersionByTagGetResponsePhase = "http_request_origin"
	VersionByTagGetResponsePhaseHTTPRequestRedirect            VersionByTagGetResponsePhase = "http_request_redirect"
	VersionByTagGetResponsePhaseHTTPRequestSanitize            VersionByTagGetResponsePhase = "http_request_sanitize"
	VersionByTagGetResponsePhaseHTTPRequestSbfm                VersionByTagGetResponsePhase = "http_request_sbfm"
	VersionByTagGetResponsePhaseHTTPRequestSelectConfiguration VersionByTagGetResponsePhase = "http_request_select_configuration"
	VersionByTagGetResponsePhaseHTTPRequestTransform           VersionByTagGetResponsePhase = "http_request_transform"
	VersionByTagGetResponsePhaseHTTPResponseCompression        VersionByTagGetResponsePhase = "http_response_compression"
	VersionByTagGetResponsePhaseHTTPResponseFirewallManaged    VersionByTagGetResponsePhase = "http_response_firewall_managed"
	VersionByTagGetResponsePhaseHTTPResponseHeadersTransform   VersionByTagGetResponsePhase = "http_response_headers_transform"
	VersionByTagGetResponsePhaseMagicTransit                   VersionByTagGetResponsePhase = "magic_transit"
	VersionByTagGetResponsePhaseMagicTransitIDsManaged         VersionByTagGetResponsePhase = "magic_transit_ids_managed"
	VersionByTagGetResponsePhaseMagicTransitManaged            VersionByTagGetResponsePhase = "magic_transit_managed"
)

func (r VersionByTagGetResponsePhase) IsKnown() bool {
	switch r {
	case VersionByTagGetResponsePhaseDDoSL4, VersionByTagGetResponsePhaseDDoSL7, VersionByTagGetResponsePhaseHTTPConfigSettings, VersionByTagGetResponsePhaseHTTPCustomErrors, VersionByTagGetResponsePhaseHTTPLogCustomFields, VersionByTagGetResponsePhaseHTTPRatelimit, VersionByTagGetResponsePhaseHTTPRequestCacheSettings, VersionByTagGetResponsePhaseHTTPRequestDynamicRedirect, VersionByTagGetResponsePhaseHTTPRequestFirewallCustom, VersionByTagGetResponsePhaseHTTPRequestFirewallManaged, VersionByTagGetResponsePhaseHTTPRequestLateTransform, VersionByTagGetResponsePhaseHTTPRequestOrigin, VersionByTagGetResponsePhaseHTTPRequestRedirect, VersionByTagGetResponsePhaseHTTPRequestSanitize, VersionByTagGetResponsePhaseHTTPRequestSbfm, VersionByTagGetResponsePhaseHTTPRequestSelectConfiguration, VersionByTagGetResponsePhaseHTTPRequestTransform, VersionByTagGetResponsePhaseHTTPResponseCompression, VersionByTagGetResponsePhaseHTTPResponseFirewallManaged, VersionByTagGetResponsePhaseHTTPResponseHeadersTransform, VersionByTagGetResponsePhaseMagicTransit, VersionByTagGetResponsePhaseMagicTransitIDsManaged, VersionByTagGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type VersionByTagGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           VersionByTagGetResponseRulesAction `json:"action"`
	ActionParameters interface{}                        `json:"action_parameters,required"`
	Categories       interface{}                        `json:"categories,required"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                          `json:"version,required"`
	JSON    versionByTagGetResponseRuleJSON `json:"-"`
	union   VersionByTagGetResponseRulesUnion
}

// versionByTagGetResponseRuleJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseRule]
type versionByTagGetResponseRuleJSON struct {
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

func (r versionByTagGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r VersionByTagGetResponseRule) AsUnion() VersionByTagGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JsChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule] or [rulesets.SetCacheSettingsRule].
type VersionByTagGetResponseRulesUnion interface {
	implementsRulesetsVersionByTagGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesAction string

const (
	VersionByTagGetResponseRulesActionBlock            VersionByTagGetResponseRulesAction = "block"
	VersionByTagGetResponseRulesActionChallenge        VersionByTagGetResponseRulesAction = "challenge"
	VersionByTagGetResponseRulesActionCompressResponse VersionByTagGetResponseRulesAction = "compress_response"
	VersionByTagGetResponseRulesActionExecute          VersionByTagGetResponseRulesAction = "execute"
	VersionByTagGetResponseRulesActionJsChallenge      VersionByTagGetResponseRulesAction = "js_challenge"
	VersionByTagGetResponseRulesActionLog              VersionByTagGetResponseRulesAction = "log"
	VersionByTagGetResponseRulesActionManagedChallenge VersionByTagGetResponseRulesAction = "managed_challenge"
	VersionByTagGetResponseRulesActionRedirect         VersionByTagGetResponseRulesAction = "redirect"
	VersionByTagGetResponseRulesActionRewrite          VersionByTagGetResponseRulesAction = "rewrite"
	VersionByTagGetResponseRulesActionRoute            VersionByTagGetResponseRulesAction = "route"
	VersionByTagGetResponseRulesActionScore            VersionByTagGetResponseRulesAction = "score"
	VersionByTagGetResponseRulesActionServeError       VersionByTagGetResponseRulesAction = "serve_error"
	VersionByTagGetResponseRulesActionSetConfig        VersionByTagGetResponseRulesAction = "set_config"
	VersionByTagGetResponseRulesActionSkip             VersionByTagGetResponseRulesAction = "skip"
	VersionByTagGetResponseRulesActionSetCacheSettings VersionByTagGetResponseRulesAction = "set_cache_settings"
)

func (r VersionByTagGetResponseRulesAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesActionBlock, VersionByTagGetResponseRulesActionChallenge, VersionByTagGetResponseRulesActionCompressResponse, VersionByTagGetResponseRulesActionExecute, VersionByTagGetResponseRulesActionJsChallenge, VersionByTagGetResponseRulesActionLog, VersionByTagGetResponseRulesActionManagedChallenge, VersionByTagGetResponseRulesActionRedirect, VersionByTagGetResponseRulesActionRewrite, VersionByTagGetResponseRulesActionRoute, VersionByTagGetResponseRulesActionScore, VersionByTagGetResponseRulesActionServeError, VersionByTagGetResponseRulesActionSetConfig, VersionByTagGetResponseRulesActionSkip, VersionByTagGetResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

type VersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type VersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result VersionByTagGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// versionByTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseEnvelope]
type versionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeErrors]
type versionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeErrorsSource]
type versionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeMessages]
type versionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeMessagesSource]
type versionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionByTagGetResponseEnvelopeSuccess bool

const (
	VersionByTagGetResponseEnvelopeSuccessTrue VersionByTagGetResponseEnvelopeSuccess = true
)

func (r VersionByTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
