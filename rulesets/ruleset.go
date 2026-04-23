// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/tidwall/gjson"
)

// RulesetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRulesetService] method instead.
type RulesetService struct {
	Options  []option.RequestOption
	Phases   *PhaseService
	Rules    *RuleService
	Versions *VersionService
}

// NewRulesetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRulesetService(opts ...option.RequestOption) (r *RulesetService) {
	r = &RulesetService{}
	r.Options = opts
	r.Phases = NewPhaseService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Versions = NewVersionService(opts...)
	return
}

// Creates a ruleset.
func (r *RulesetService) New(ctx context.Context, params RulesetNewParams, opts ...option.RequestOption) (res *RulesetNewResponse, err error) {
	var env RulesetNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an account or zone ruleset, creating a new version.
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetUpdateResponse, err error) {
	var env RulesetUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches all rulesets.
func (r *RulesetService) List(ctx context.Context, params RulesetListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RulesetListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetches all rulesets.
func (r *RulesetService) ListAutoPaging(ctx context.Context, params RulesetListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RulesetListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes all versions of an existing account or zone ruleset.
func (r *RulesetService) Delete(ctx context.Context, rulesetID string, body RulesetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, rulesetID string, query RulesetGetParams, opts ...option.RequestOption) (res *RulesetGetResponse, err error) {
	var env RulesetGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// The kind of the ruleset.
type Kind string

const (
	KindManaged Kind = "managed"
	KindCustom  Kind = "custom"
	KindRoot    Kind = "root"
	KindZone    Kind = "zone"
)

func (r Kind) IsKnown() bool {
	switch r {
	case KindManaged, KindCustom, KindRoot, KindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type Phase string

const (
	PhaseDDoSL4                       Phase = "ddos_l4"
	PhaseDDoSL7                       Phase = "ddos_l7"
	PhaseHTTPConfigSettings           Phase = "http_config_settings"
	PhaseHTTPCustomErrors             Phase = "http_custom_errors"
	PhaseHTTPLogCustomFields          Phase = "http_log_custom_fields"
	PhaseHTTPRatelimit                Phase = "http_ratelimit"
	PhaseHTTPRequestCacheSettings     Phase = "http_request_cache_settings"
	PhaseHTTPRequestDynamicRedirect   Phase = "http_request_dynamic_redirect"
	PhaseHTTPRequestFirewallCustom    Phase = "http_request_firewall_custom"
	PhaseHTTPRequestFirewallManaged   Phase = "http_request_firewall_managed"
	PhaseHTTPRequestLateTransform     Phase = "http_request_late_transform"
	PhaseHTTPRequestOrigin            Phase = "http_request_origin"
	PhaseHTTPRequestRedirect          Phase = "http_request_redirect"
	PhaseHTTPRequestSanitize          Phase = "http_request_sanitize"
	PhaseHTTPRequestSBFM              Phase = "http_request_sbfm"
	PhaseHTTPRequestTransform         Phase = "http_request_transform"
	PhaseHTTPResponseCacheSettings    Phase = "http_response_cache_settings"
	PhaseHTTPResponseCompression      Phase = "http_response_compression"
	PhaseHTTPResponseFirewallManaged  Phase = "http_response_firewall_managed"
	PhaseHTTPResponseHeadersTransform Phase = "http_response_headers_transform"
	PhaseMagicTransit                 Phase = "magic_transit"
	PhaseMagicTransitIDsManaged       Phase = "magic_transit_ids_managed"
	PhaseMagicTransitManaged          Phase = "magic_transit_managed"
	PhaseMagicTransitRatelimit        Phase = "magic_transit_ratelimit"
)

func (r Phase) IsKnown() bool {
	switch r {
	case PhaseDDoSL4, PhaseDDoSL7, PhaseHTTPConfigSettings, PhaseHTTPCustomErrors, PhaseHTTPLogCustomFields, PhaseHTTPRatelimit, PhaseHTTPRequestCacheSettings, PhaseHTTPRequestDynamicRedirect, PhaseHTTPRequestFirewallCustom, PhaseHTTPRequestFirewallManaged, PhaseHTTPRequestLateTransform, PhaseHTTPRequestOrigin, PhaseHTTPRequestRedirect, PhaseHTTPRequestSanitize, PhaseHTTPRequestSBFM, PhaseHTTPRequestTransform, PhaseHTTPResponseCacheSettings, PhaseHTTPResponseCompression, PhaseHTTPResponseFirewallManaged, PhaseHTTPResponseHeadersTransform, PhaseMagicTransit, PhaseMagicTransitIDsManaged, PhaseMagicTransitManaged, PhaseMagicTransitRatelimit:
		return true
	}
	return false
}

// A ruleset object.
type RulesetNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The list of rules in the ruleset.
	Rules []RulesetNewResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetNewResponseJSON `json:"-"`
}

// rulesetNewResponseJSON contains the JSON metadata for the struct
// [RulesetNewResponse]
type rulesetNewResponseJSON struct {
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

func (r *RulesetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [RulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RulesetNewResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [LogCustomFieldRuleRatelimit], [ManagedChallengeRuleRatelimit],
	// [RedirectRuleRatelimit], [RewriteRuleRatelimit], [RouteRuleRatelimit],
	// [ScoreRuleRatelimit], [ServeErrorRuleRatelimit],
	// [RulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                     `json:"ref"`
	JSON  rulesetNewResponseRuleJSON `json:"-"`
	union RulesetNewResponseRulesUnion
}

// rulesetNewResponseRuleJSON contains the JSON metadata for the struct
// [RulesetNewResponseRule]
type rulesetNewResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r rulesetNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetNewResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetNewResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [RulesetNewResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetNewResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule].
func (r RulesetNewResponseRule) AsUnion() RulesetNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RulesetNewResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [RulesetNewResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetNewResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule] or
// [SkipRule].
type RulesetNewResponseRulesUnion interface {
	implementsRulesetNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
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
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
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
	)
}

type RulesetNewResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsChallengeRule]
type rulesetNewResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsChallengeRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge RulesetNewResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                    `json:"score_response_header_name"`
	JSON                    rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsJSChallengeRule]
type rulesetNewResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsJSChallengeRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetNewResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetNewResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                      `json:"score_response_header_name"`
	JSON                    rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsSetCacheControlRule]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetNewResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParameters]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
	Immutable            apijson.Field
	MaxAge               apijson.Field
	MustRevalidate       apijson.Field
	MustUnderstand       apijson.Field
	NoCache              apijson.Field
	NoStore              apijson.Field
	NoTransform          apijson.Field
	Private              apijson.Field
	ProxyRevalidate      apijson.Field
	Public               apijson.Field
	SMaxage              apijson.Field
	StaleIfError         apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                        `json:"value"`
	JSON  rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                   `json:"qualifiers"`
	JSON       rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                  `json:"qualifiers"`
	JSON       rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                   `json:"qualifiers"`
	JSON       rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                  `json:"qualifiers"`
	JSON       rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                         `json:"value"`
	JSON  rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                              `json:"value"`
	JSON  rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                      `json:"value"`
	JSON  rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetNewResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                       `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                          `json:"score_response_header_name"`
	JSON                    rulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimit]
type rulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsSetCacheTagsRule]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRule) implementsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetNewResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                         `json:"values"`
	JSON   rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetNewResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                       `json:"score_response_header_name"`
	JSON                    rulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type rulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesAction string

const (
	RulesetNewResponseRulesActionBlock                RulesetNewResponseRulesAction = "block"
	RulesetNewResponseRulesActionChallenge            RulesetNewResponseRulesAction = "challenge"
	RulesetNewResponseRulesActionCompressResponse     RulesetNewResponseRulesAction = "compress_response"
	RulesetNewResponseRulesActionDDoSDynamic          RulesetNewResponseRulesAction = "ddos_dynamic"
	RulesetNewResponseRulesActionExecute              RulesetNewResponseRulesAction = "execute"
	RulesetNewResponseRulesActionForceConnectionClose RulesetNewResponseRulesAction = "force_connection_close"
	RulesetNewResponseRulesActionJSChallenge          RulesetNewResponseRulesAction = "js_challenge"
	RulesetNewResponseRulesActionLog                  RulesetNewResponseRulesAction = "log"
	RulesetNewResponseRulesActionLogCustomField       RulesetNewResponseRulesAction = "log_custom_field"
	RulesetNewResponseRulesActionManagedChallenge     RulesetNewResponseRulesAction = "managed_challenge"
	RulesetNewResponseRulesActionRedirect             RulesetNewResponseRulesAction = "redirect"
	RulesetNewResponseRulesActionRewrite              RulesetNewResponseRulesAction = "rewrite"
	RulesetNewResponseRulesActionRoute                RulesetNewResponseRulesAction = "route"
	RulesetNewResponseRulesActionScore                RulesetNewResponseRulesAction = "score"
	RulesetNewResponseRulesActionServeError           RulesetNewResponseRulesAction = "serve_error"
	RulesetNewResponseRulesActionSetCacheControl      RulesetNewResponseRulesAction = "set_cache_control"
	RulesetNewResponseRulesActionSetCacheSettings     RulesetNewResponseRulesAction = "set_cache_settings"
	RulesetNewResponseRulesActionSetCacheTags         RulesetNewResponseRulesAction = "set_cache_tags"
	RulesetNewResponseRulesActionSetConfig            RulesetNewResponseRulesAction = "set_config"
	RulesetNewResponseRulesActionSkip                 RulesetNewResponseRulesAction = "skip"
)

func (r RulesetNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesActionBlock, RulesetNewResponseRulesActionChallenge, RulesetNewResponseRulesActionCompressResponse, RulesetNewResponseRulesActionDDoSDynamic, RulesetNewResponseRulesActionExecute, RulesetNewResponseRulesActionForceConnectionClose, RulesetNewResponseRulesActionJSChallenge, RulesetNewResponseRulesActionLog, RulesetNewResponseRulesActionLogCustomField, RulesetNewResponseRulesActionManagedChallenge, RulesetNewResponseRulesActionRedirect, RulesetNewResponseRulesActionRewrite, RulesetNewResponseRulesActionRoute, RulesetNewResponseRulesActionScore, RulesetNewResponseRulesActionServeError, RulesetNewResponseRulesActionSetCacheControl, RulesetNewResponseRulesActionSetCacheSettings, RulesetNewResponseRulesActionSetCacheTags, RulesetNewResponseRulesActionSetConfig, RulesetNewResponseRulesActionSkip:
		return true
	}
	return false
}

// A ruleset object.
type RulesetUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The list of rules in the ruleset.
	Rules []RulesetUpdateResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                    `json:"description"`
	JSON        rulesetUpdateResponseJSON `json:"-"`
}

// rulesetUpdateResponseJSON contains the JSON metadata for the struct
// [RulesetUpdateResponse]
type rulesetUpdateResponseJSON struct {
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

func (r *RulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                        `json:"ref"`
	JSON  rulesetUpdateResponseRuleJSON `json:"-"`
	union RulesetUpdateResponseRulesUnion
}

// rulesetUpdateResponseRuleJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseRule]
type rulesetUpdateResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r rulesetUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetUpdateResponseRulesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [RulesetUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule],
// [SkipRule].
func (r RulesetUpdateResponseRule) AsUnion() RulesetUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule],
// [RulesetUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [RulesetUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule] or
// [SkipRule].
type RulesetUpdateResponseRulesUnion interface {
	implementsRulesetUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
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
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
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
	)
}

type RulesetUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsChallengeRule]
type rulesetUpdateResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsChallengeRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge RulesetUpdateResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                       `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsJSChallengeRule]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsJSChallengeRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                      `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                         `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                    `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsSetCacheControlRule]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRule) implementsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetUpdateResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
	Immutable            apijson.Field
	MaxAge               apijson.Field
	MustRevalidate       apijson.Field
	MustUnderstand       apijson.Field
	NoCache              apijson.Field
	NoStore              apijson.Field
	NoTransform          apijson.Field
	Private              apijson.Field
	ProxyRevalidate      apijson.Field
	Public               apijson.Field
	SMaxage              apijson.Field
	StaleIfError         apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                           `json:"value"`
	JSON  rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                      `json:"qualifiers"`
	JSON       rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                     `json:"qualifiers"`
	JSON       rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                             `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                      `json:"qualifiers"`
	JSON       rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                     `json:"qualifiers"`
	JSON       rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                            `json:"value"`
	JSON  rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                 `json:"value"`
	JSON  rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                         `json:"value"`
	JSON  rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                          `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                             `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsSetCacheTagsRule]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRule) implementsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                            `json:"values"`
	JSON   rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                    `json:"values" api:"required"`
	JSON   rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                       `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                          `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesAction string

const (
	RulesetUpdateResponseRulesActionBlock                RulesetUpdateResponseRulesAction = "block"
	RulesetUpdateResponseRulesActionChallenge            RulesetUpdateResponseRulesAction = "challenge"
	RulesetUpdateResponseRulesActionCompressResponse     RulesetUpdateResponseRulesAction = "compress_response"
	RulesetUpdateResponseRulesActionDDoSDynamic          RulesetUpdateResponseRulesAction = "ddos_dynamic"
	RulesetUpdateResponseRulesActionExecute              RulesetUpdateResponseRulesAction = "execute"
	RulesetUpdateResponseRulesActionForceConnectionClose RulesetUpdateResponseRulesAction = "force_connection_close"
	RulesetUpdateResponseRulesActionJSChallenge          RulesetUpdateResponseRulesAction = "js_challenge"
	RulesetUpdateResponseRulesActionLog                  RulesetUpdateResponseRulesAction = "log"
	RulesetUpdateResponseRulesActionLogCustomField       RulesetUpdateResponseRulesAction = "log_custom_field"
	RulesetUpdateResponseRulesActionManagedChallenge     RulesetUpdateResponseRulesAction = "managed_challenge"
	RulesetUpdateResponseRulesActionRedirect             RulesetUpdateResponseRulesAction = "redirect"
	RulesetUpdateResponseRulesActionRewrite              RulesetUpdateResponseRulesAction = "rewrite"
	RulesetUpdateResponseRulesActionRoute                RulesetUpdateResponseRulesAction = "route"
	RulesetUpdateResponseRulesActionScore                RulesetUpdateResponseRulesAction = "score"
	RulesetUpdateResponseRulesActionServeError           RulesetUpdateResponseRulesAction = "serve_error"
	RulesetUpdateResponseRulesActionSetCacheControl      RulesetUpdateResponseRulesAction = "set_cache_control"
	RulesetUpdateResponseRulesActionSetCacheSettings     RulesetUpdateResponseRulesAction = "set_cache_settings"
	RulesetUpdateResponseRulesActionSetCacheTags         RulesetUpdateResponseRulesAction = "set_cache_tags"
	RulesetUpdateResponseRulesActionSetConfig            RulesetUpdateResponseRulesAction = "set_config"
	RulesetUpdateResponseRulesActionSkip                 RulesetUpdateResponseRulesAction = "skip"
)

func (r RulesetUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesActionBlock, RulesetUpdateResponseRulesActionChallenge, RulesetUpdateResponseRulesActionCompressResponse, RulesetUpdateResponseRulesActionDDoSDynamic, RulesetUpdateResponseRulesActionExecute, RulesetUpdateResponseRulesActionForceConnectionClose, RulesetUpdateResponseRulesActionJSChallenge, RulesetUpdateResponseRulesActionLog, RulesetUpdateResponseRulesActionLogCustomField, RulesetUpdateResponseRulesActionManagedChallenge, RulesetUpdateResponseRulesActionRedirect, RulesetUpdateResponseRulesActionRewrite, RulesetUpdateResponseRulesActionRoute, RulesetUpdateResponseRulesActionScore, RulesetUpdateResponseRulesActionServeError, RulesetUpdateResponseRulesActionSetCacheControl, RulesetUpdateResponseRulesActionSetCacheSettings, RulesetUpdateResponseRulesActionSetCacheTags, RulesetUpdateResponseRulesActionSetConfig, RulesetUpdateResponseRulesActionSkip:
		return true
	}
	return false
}

// A ruleset object.
type RulesetListResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                  `json:"description"`
	JSON        rulesetListResponseJSON `json:"-"`
}

// rulesetListResponseJSON contains the JSON metadata for the struct
// [RulesetListResponse]
type rulesetListResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetListResponseJSON) RawJSON() string {
	return r.raw
}

// A ruleset object.
type RulesetGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The list of rules in the ruleset.
	Rules []RulesetGetResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        rulesetGetResponseJSON `json:"-"`
}

// rulesetGetResponseJSON contains the JSON metadata for the struct
// [RulesetGetResponse]
type rulesetGetResponseJSON struct {
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

func (r *RulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [LogCustomFieldRuleRatelimit], [ManagedChallengeRuleRatelimit],
	// [RedirectRuleRatelimit], [RewriteRuleRatelimit], [RouteRuleRatelimit],
	// [ScoreRuleRatelimit], [ServeErrorRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                     `json:"ref"`
	JSON  rulesetGetResponseRuleJSON `json:"-"`
	union RulesetGetResponseRulesUnion
}

// rulesetGetResponseRuleJSON contains the JSON metadata for the struct
// [RulesetGetResponseRule]
type rulesetGetResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r rulesetGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetGetResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [RulesetGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetGetResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule].
func (r RulesetGetResponseRule) AsUnion() RulesetGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RulesetGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [RulesetGetResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetGetResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule] or
// [SkipRule].
type RulesetGetResponseRulesUnion interface {
	implementsRulesetGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
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
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
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
	)
}

type RulesetGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                           `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsChallengeRule]
type rulesetGetResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsChallengeRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge RulesetGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression" api:"required"`
	JSON               rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetGetResponseRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                    `json:"score_response_header_name"`
	JSON                    rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsChallengeRuleRatelimit]
type rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                             `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsJSChallengeRule]
type rulesetGetResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsJSChallengeRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge RulesetGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression" api:"required"`
	JSON               rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                      `json:"score_response_header_name"`
	JSON                    rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit]
type rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                 `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsSetCacheControlRule]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetGetResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParameters]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
	Immutable            apijson.Field
	MaxAge               apijson.Field
	MustRevalidate       apijson.Field
	MustUnderstand       apijson.Field
	NoCache              apijson.Field
	NoStore              apijson.Field
	NoTransform          apijson.Field
	Private              apijson.Field
	ProxyRevalidate      apijson.Field
	Public               apijson.Field
	SMaxage              apijson.Field
	StaleIfError         apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                        `json:"value"`
	JSON  rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                   `json:"qualifiers"`
	JSON       rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                  `json:"qualifiers"`
	JSON       rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                   `json:"qualifiers"`
	JSON       rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                  `json:"qualifiers"`
	JSON       rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                         `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                         `json:"value"`
	JSON  rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                              `json:"value"`
	JSON  rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                      `json:"value"`
	JSON  rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                       `json:"username_expression" api:"required"`
	JSON               rulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                          `json:"score_response_header_name"`
	JSON                    rulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimit]
type rulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                              `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsSetCacheTagsRule]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRule) implementsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetGetResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                         `json:"values"`
	JSON   rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               rulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                       `json:"score_response_header_name"`
	JSON                    rulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type rulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesAction string

const (
	RulesetGetResponseRulesActionBlock                RulesetGetResponseRulesAction = "block"
	RulesetGetResponseRulesActionChallenge            RulesetGetResponseRulesAction = "challenge"
	RulesetGetResponseRulesActionCompressResponse     RulesetGetResponseRulesAction = "compress_response"
	RulesetGetResponseRulesActionDDoSDynamic          RulesetGetResponseRulesAction = "ddos_dynamic"
	RulesetGetResponseRulesActionExecute              RulesetGetResponseRulesAction = "execute"
	RulesetGetResponseRulesActionForceConnectionClose RulesetGetResponseRulesAction = "force_connection_close"
	RulesetGetResponseRulesActionJSChallenge          RulesetGetResponseRulesAction = "js_challenge"
	RulesetGetResponseRulesActionLog                  RulesetGetResponseRulesAction = "log"
	RulesetGetResponseRulesActionLogCustomField       RulesetGetResponseRulesAction = "log_custom_field"
	RulesetGetResponseRulesActionManagedChallenge     RulesetGetResponseRulesAction = "managed_challenge"
	RulesetGetResponseRulesActionRedirect             RulesetGetResponseRulesAction = "redirect"
	RulesetGetResponseRulesActionRewrite              RulesetGetResponseRulesAction = "rewrite"
	RulesetGetResponseRulesActionRoute                RulesetGetResponseRulesAction = "route"
	RulesetGetResponseRulesActionScore                RulesetGetResponseRulesAction = "score"
	RulesetGetResponseRulesActionServeError           RulesetGetResponseRulesAction = "serve_error"
	RulesetGetResponseRulesActionSetCacheControl      RulesetGetResponseRulesAction = "set_cache_control"
	RulesetGetResponseRulesActionSetCacheSettings     RulesetGetResponseRulesAction = "set_cache_settings"
	RulesetGetResponseRulesActionSetCacheTags         RulesetGetResponseRulesAction = "set_cache_tags"
	RulesetGetResponseRulesActionSetConfig            RulesetGetResponseRulesAction = "set_config"
	RulesetGetResponseRulesActionSkip                 RulesetGetResponseRulesAction = "skip"
)

func (r RulesetGetResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesActionBlock, RulesetGetResponseRulesActionChallenge, RulesetGetResponseRulesActionCompressResponse, RulesetGetResponseRulesActionDDoSDynamic, RulesetGetResponseRulesActionExecute, RulesetGetResponseRulesActionForceConnectionClose, RulesetGetResponseRulesActionJSChallenge, RulesetGetResponseRulesActionLog, RulesetGetResponseRulesActionLogCustomField, RulesetGetResponseRulesActionManagedChallenge, RulesetGetResponseRulesActionRedirect, RulesetGetResponseRulesActionRewrite, RulesetGetResponseRulesActionRoute, RulesetGetResponseRulesActionScore, RulesetGetResponseRulesActionServeError, RulesetGetResponseRulesActionSetCacheControl, RulesetGetResponseRulesActionSetCacheSettings, RulesetGetResponseRulesActionSetCacheTags, RulesetGetResponseRulesActionSetConfig, RulesetGetResponseRulesActionSkip:
		return true
	}
	return false
}

type RulesetNewParams struct {
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind" api:"required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase" api:"required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetNewParamsRuleUnion] `json:"rules"`
}

func (r RulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                 `json:"action_parameters"`
	Categories       param.Field[interface{}]                 `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRule) implementsRulesetNewParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ExecuteRuleParam], [rulesets.ForceConnectionCloseRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.LogCustomFieldRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRule],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRule],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam], [RulesetNewParamsRule].
type RulesetNewParamsRuleUnion interface {
	implementsRulesetNewParamsRuleUnion()
}

type RulesetNewParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsChallengeRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge RulesetNewParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetNewParamsRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetNewParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsJSChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsJSChallengeRuleActionJSChallenge RulesetNewParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetNewParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsSetCacheControlRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetNewParamsRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion] `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion] `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion] `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion] `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion] `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion] `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion] `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion] `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion] `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion] `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion] `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion] `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion] `json:"stale-while-revalidate"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion()
}

// Set the directive with a duration value in seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion()
}

// Set the directive with optional qualifiers.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion()
}

// Set the directive with optional qualifiers.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublic].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion()
}

// Set the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion()
}

// Set the directive with a duration value in seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion()
}

// Set the directive with a duration value in seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective],
// [RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate].
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion()
}

// Set the directive with a duration value in seconds.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetNewParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetNewParamsRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheControlRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRule) implementsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetNewParamsRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation] `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string]      `json:"expression"`
	Values     param.Field[interface{}] `json:"values"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParameters) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The parameters configuring the rule's action.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression],
// [RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParameters].
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion()
}

// Add cache tags using a list of values.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetNewParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetNewParamsRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheTagsRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesAction string

const (
	RulesetNewParamsRulesActionBlock                RulesetNewParamsRulesAction = "block"
	RulesetNewParamsRulesActionChallenge            RulesetNewParamsRulesAction = "challenge"
	RulesetNewParamsRulesActionCompressResponse     RulesetNewParamsRulesAction = "compress_response"
	RulesetNewParamsRulesActionDDoSDynamic          RulesetNewParamsRulesAction = "ddos_dynamic"
	RulesetNewParamsRulesActionExecute              RulesetNewParamsRulesAction = "execute"
	RulesetNewParamsRulesActionForceConnectionClose RulesetNewParamsRulesAction = "force_connection_close"
	RulesetNewParamsRulesActionJSChallenge          RulesetNewParamsRulesAction = "js_challenge"
	RulesetNewParamsRulesActionLog                  RulesetNewParamsRulesAction = "log"
	RulesetNewParamsRulesActionLogCustomField       RulesetNewParamsRulesAction = "log_custom_field"
	RulesetNewParamsRulesActionManagedChallenge     RulesetNewParamsRulesAction = "managed_challenge"
	RulesetNewParamsRulesActionRedirect             RulesetNewParamsRulesAction = "redirect"
	RulesetNewParamsRulesActionRewrite              RulesetNewParamsRulesAction = "rewrite"
	RulesetNewParamsRulesActionRoute                RulesetNewParamsRulesAction = "route"
	RulesetNewParamsRulesActionScore                RulesetNewParamsRulesAction = "score"
	RulesetNewParamsRulesActionServeError           RulesetNewParamsRulesAction = "serve_error"
	RulesetNewParamsRulesActionSetCacheControl      RulesetNewParamsRulesAction = "set_cache_control"
	RulesetNewParamsRulesActionSetCacheSettings     RulesetNewParamsRulesAction = "set_cache_settings"
	RulesetNewParamsRulesActionSetCacheTags         RulesetNewParamsRulesAction = "set_cache_tags"
	RulesetNewParamsRulesActionSetConfig            RulesetNewParamsRulesAction = "set_config"
	RulesetNewParamsRulesActionSkip                 RulesetNewParamsRulesAction = "skip"
)

func (r RulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesActionBlock, RulesetNewParamsRulesActionChallenge, RulesetNewParamsRulesActionCompressResponse, RulesetNewParamsRulesActionDDoSDynamic, RulesetNewParamsRulesActionExecute, RulesetNewParamsRulesActionForceConnectionClose, RulesetNewParamsRulesActionJSChallenge, RulesetNewParamsRulesActionLog, RulesetNewParamsRulesActionLogCustomField, RulesetNewParamsRulesActionManagedChallenge, RulesetNewParamsRulesActionRedirect, RulesetNewParamsRulesActionRewrite, RulesetNewParamsRulesActionRoute, RulesetNewParamsRulesActionScore, RulesetNewParamsRulesActionServeError, RulesetNewParamsRulesActionSetCacheControl, RulesetNewParamsRulesActionSetCacheSettings, RulesetNewParamsRulesActionSetCacheTags, RulesetNewParamsRulesActionSetConfig, RulesetNewParamsRulesActionSkip:
		return true
	}
	return false
}

// A response object.
type RulesetNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetNewResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []RulesetNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result RulesetNewResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RulesetNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    rulesetNewResponseEnvelopeJSON    `json:"-"`
}

// rulesetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelope]
type rulesetNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeErrors]
type rulesetNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer" api:"required"`
	JSON    rulesetNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeErrorsSource]
type rulesetNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeMessages]
type rulesetNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer" api:"required"`
	JSON    rulesetNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeMessagesSource]
type rulesetNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetNewResponseEnvelopeSuccess bool

const (
	RulesetNewResponseEnvelopeSuccessTrue RulesetNewResponseEnvelopeSuccess = true
)

func (r RulesetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RulesetUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetUpdateParamsRuleUnion] `json:"rules"`
}

func (r RulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[RulesetUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                    `json:"action_parameters"`
	Categories       param.Field[interface{}]                    `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRule) implementsRulesetUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ExecuteRuleParam], [rulesets.ForceConnectionCloseRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.LogCustomFieldRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRule],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRule],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [RulesetUpdateParamsRule].
type RulesetUpdateParamsRuleUnion interface {
	implementsRulesetUpdateParamsRuleUnion()
}

type RulesetUpdateParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) implementsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge RulesetUpdateParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRule) implementsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetUpdateParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsSetCacheControlRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRule) implementsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetUpdateParamsRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion] `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion] `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion] `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion] `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion] `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion] `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion] `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion] `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion] `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion] `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion] `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion] `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion] `json:"stale-while-revalidate"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion()
}

// Set the directive with a duration value in seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion()
}

// Set the directive with optional qualifiers.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion()
}

// Set the directive with optional qualifiers.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion()
}

// Set the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion()
}

// Set the directive with a duration value in seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion()
}

// Set the directive with a duration value in seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective],
// [RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate].
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion()
}

// Set the directive with a duration value in seconds.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRule) implementsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation] `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string]      `json:"expression"`
	Values     param.Field[interface{}] `json:"values"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The parameters configuring the rule's action.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression],
// [RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters].
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion()
}

// Add cache tags using a list of values.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesAction string

const (
	RulesetUpdateParamsRulesActionBlock                RulesetUpdateParamsRulesAction = "block"
	RulesetUpdateParamsRulesActionChallenge            RulesetUpdateParamsRulesAction = "challenge"
	RulesetUpdateParamsRulesActionCompressResponse     RulesetUpdateParamsRulesAction = "compress_response"
	RulesetUpdateParamsRulesActionDDoSDynamic          RulesetUpdateParamsRulesAction = "ddos_dynamic"
	RulesetUpdateParamsRulesActionExecute              RulesetUpdateParamsRulesAction = "execute"
	RulesetUpdateParamsRulesActionForceConnectionClose RulesetUpdateParamsRulesAction = "force_connection_close"
	RulesetUpdateParamsRulesActionJSChallenge          RulesetUpdateParamsRulesAction = "js_challenge"
	RulesetUpdateParamsRulesActionLog                  RulesetUpdateParamsRulesAction = "log"
	RulesetUpdateParamsRulesActionLogCustomField       RulesetUpdateParamsRulesAction = "log_custom_field"
	RulesetUpdateParamsRulesActionManagedChallenge     RulesetUpdateParamsRulesAction = "managed_challenge"
	RulesetUpdateParamsRulesActionRedirect             RulesetUpdateParamsRulesAction = "redirect"
	RulesetUpdateParamsRulesActionRewrite              RulesetUpdateParamsRulesAction = "rewrite"
	RulesetUpdateParamsRulesActionRoute                RulesetUpdateParamsRulesAction = "route"
	RulesetUpdateParamsRulesActionScore                RulesetUpdateParamsRulesAction = "score"
	RulesetUpdateParamsRulesActionServeError           RulesetUpdateParamsRulesAction = "serve_error"
	RulesetUpdateParamsRulesActionSetCacheControl      RulesetUpdateParamsRulesAction = "set_cache_control"
	RulesetUpdateParamsRulesActionSetCacheSettings     RulesetUpdateParamsRulesAction = "set_cache_settings"
	RulesetUpdateParamsRulesActionSetCacheTags         RulesetUpdateParamsRulesAction = "set_cache_tags"
	RulesetUpdateParamsRulesActionSetConfig            RulesetUpdateParamsRulesAction = "set_config"
	RulesetUpdateParamsRulesActionSkip                 RulesetUpdateParamsRulesAction = "skip"
)

func (r RulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesActionBlock, RulesetUpdateParamsRulesActionChallenge, RulesetUpdateParamsRulesActionCompressResponse, RulesetUpdateParamsRulesActionDDoSDynamic, RulesetUpdateParamsRulesActionExecute, RulesetUpdateParamsRulesActionForceConnectionClose, RulesetUpdateParamsRulesActionJSChallenge, RulesetUpdateParamsRulesActionLog, RulesetUpdateParamsRulesActionLogCustomField, RulesetUpdateParamsRulesActionManagedChallenge, RulesetUpdateParamsRulesActionRedirect, RulesetUpdateParamsRulesActionRewrite, RulesetUpdateParamsRulesActionRoute, RulesetUpdateParamsRulesActionScore, RulesetUpdateParamsRulesActionServeError, RulesetUpdateParamsRulesActionSetCacheControl, RulesetUpdateParamsRulesActionSetCacheSettings, RulesetUpdateParamsRulesActionSetCacheTags, RulesetUpdateParamsRulesActionSetConfig, RulesetUpdateParamsRulesActionSkip:
		return true
	}
	return false
}

// A response object.
type RulesetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []RulesetUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result RulesetUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RulesetUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    rulesetUpdateResponseEnvelopeJSON    `json:"-"`
}

// rulesetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelope]
type rulesetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrors]
type rulesetUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer" api:"required"`
	JSON    rulesetUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeErrorsSource]
type rulesetUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeMessages]
type rulesetUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer" api:"required"`
	JSON    rulesetUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseEnvelopeMessagesSource]
type rulesetUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetUpdateResponseEnvelopeSuccess bool

const (
	RulesetUpdateResponseEnvelopeSuccessTrue RulesetUpdateResponseEnvelopeSuccess = true
)

func (r RulesetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RulesetListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The cursor to use for the next page.
	Cursor param.Field[string] `query:"cursor"`
	// The number of rulesets to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [RulesetListParams]'s query parameters as `url.Values`.
func (r RulesetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RulesetDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type RulesetGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RulesetGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []RulesetGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result RulesetGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success RulesetGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    rulesetGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelope]
type rulesetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeErrors]
type rulesetGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer" api:"required"`
	JSON    rulesetGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeErrorsSource]
type rulesetGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RulesetGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RulesetGetResponseEnvelopeMessages]
type rulesetGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RulesetGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer" api:"required"`
	JSON    rulesetGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RulesetGetResponseEnvelopeMessagesSource]
type rulesetGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RulesetGetResponseEnvelopeSuccess bool

const (
	RulesetGetResponseEnvelopeSuccessTrue RulesetGetResponseEnvelopeSuccess = true
)

func (r RulesetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
