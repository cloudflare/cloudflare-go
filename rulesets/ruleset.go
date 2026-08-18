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

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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
func (r *RulesetService) New(ctx context.Context, params RulesetNewParams, opts ...option.RequestOption) (res *RulesetNewResponseEnvelopeResult, err error) {
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
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetUpdateResponseEnvelopeResult, err error) {
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
func (r *RulesetService) Delete(ctx context.Context, rulesetID string, params RulesetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
		return err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
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
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters].
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
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck].
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
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit].
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
// [RulesetGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule],
// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRule].
func (r RulesetGetResponseRule) AsUnion() RulesetGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [RulesetGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [RulesetGetResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetGetResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [RulesetGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule]
// or [RulesetGetResponseRulesRulesetsTransformResponseHTMLRule].
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
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

type RulesetGetResponseRulesRulesetsTransformResponseHTMLRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                       `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRule]
type rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsTransformResponseHTMLRule) implementsRulesetGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleAction string

const (
	RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                  `json:"link_maze" api:"required"`
	JSON     rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters]
type rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                             `json:"username_expression" api:"required"`
	JSON               rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck]
type rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                `json:"score_response_header_name"`
	JSON                    rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit]
type rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesAction string

const (
	RulesetGetResponseRulesActionBlock                 RulesetGetResponseRulesAction = "block"
	RulesetGetResponseRulesActionChallenge             RulesetGetResponseRulesAction = "challenge"
	RulesetGetResponseRulesActionCompressResponse      RulesetGetResponseRulesAction = "compress_response"
	RulesetGetResponseRulesActionDDoSDynamic           RulesetGetResponseRulesAction = "ddos_dynamic"
	RulesetGetResponseRulesActionExecute               RulesetGetResponseRulesAction = "execute"
	RulesetGetResponseRulesActionForceConnectionClose  RulesetGetResponseRulesAction = "force_connection_close"
	RulesetGetResponseRulesActionJSChallenge           RulesetGetResponseRulesAction = "js_challenge"
	RulesetGetResponseRulesActionLog                   RulesetGetResponseRulesAction = "log"
	RulesetGetResponseRulesActionLogCustomField        RulesetGetResponseRulesAction = "log_custom_field"
	RulesetGetResponseRulesActionManagedChallenge      RulesetGetResponseRulesAction = "managed_challenge"
	RulesetGetResponseRulesActionRedirect              RulesetGetResponseRulesAction = "redirect"
	RulesetGetResponseRulesActionRewrite               RulesetGetResponseRulesAction = "rewrite"
	RulesetGetResponseRulesActionRoute                 RulesetGetResponseRulesAction = "route"
	RulesetGetResponseRulesActionScore                 RulesetGetResponseRulesAction = "score"
	RulesetGetResponseRulesActionServeError            RulesetGetResponseRulesAction = "serve_error"
	RulesetGetResponseRulesActionSetCacheControl       RulesetGetResponseRulesAction = "set_cache_control"
	RulesetGetResponseRulesActionSetCacheSettings      RulesetGetResponseRulesAction = "set_cache_settings"
	RulesetGetResponseRulesActionSetCacheTags          RulesetGetResponseRulesAction = "set_cache_tags"
	RulesetGetResponseRulesActionSetConfig             RulesetGetResponseRulesAction = "set_config"
	RulesetGetResponseRulesActionSkip                  RulesetGetResponseRulesAction = "skip"
	RulesetGetResponseRulesActionTransformResponseHTML RulesetGetResponseRulesAction = "transform_response_html"
)

func (r RulesetGetResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesActionBlock, RulesetGetResponseRulesActionChallenge, RulesetGetResponseRulesActionCompressResponse, RulesetGetResponseRulesActionDDoSDynamic, RulesetGetResponseRulesActionExecute, RulesetGetResponseRulesActionForceConnectionClose, RulesetGetResponseRulesActionJSChallenge, RulesetGetResponseRulesActionLog, RulesetGetResponseRulesActionLogCustomField, RulesetGetResponseRulesActionManagedChallenge, RulesetGetResponseRulesActionRedirect, RulesetGetResponseRulesActionRewrite, RulesetGetResponseRulesActionRoute, RulesetGetResponseRulesActionScore, RulesetGetResponseRulesActionServeError, RulesetGetResponseRulesActionSetCacheControl, RulesetGetResponseRulesActionSetCacheSettings, RulesetGetResponseRulesActionSetCacheTags, RulesetGetResponseRulesActionSetConfig, RulesetGetResponseRulesActionSkip, RulesetGetResponseRulesActionTransformResponseHTML:
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
	// Validates the request without persisting changes when set to `true`. Responses
	// that normally return 200 return `result: null`; endpoints that normally return
	// 204 continue to return 204.
	DryRun param.Field[bool] `query:"dry_run"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetNewParamsRuleUnion] `json:"rules"`
}

func (r RulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [RulesetNewParams]'s query parameters as `url.Values`.
func (r RulesetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsTransformResponseHTMLRule],
// [RulesetNewParamsRule].
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

type RulesetNewParamsRulesRulesetsTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRule) implementsRulesetNewParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleAction string

const (
	RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze param.Field[interface{}] `json:"link_maze" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for exposed credential checking.
type RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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

func (r RulesetNewParamsRulesRulesetsTransformResponseHTMLRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesAction string

const (
	RulesetNewParamsRulesActionBlock                 RulesetNewParamsRulesAction = "block"
	RulesetNewParamsRulesActionChallenge             RulesetNewParamsRulesAction = "challenge"
	RulesetNewParamsRulesActionCompressResponse      RulesetNewParamsRulesAction = "compress_response"
	RulesetNewParamsRulesActionDDoSDynamic           RulesetNewParamsRulesAction = "ddos_dynamic"
	RulesetNewParamsRulesActionExecute               RulesetNewParamsRulesAction = "execute"
	RulesetNewParamsRulesActionForceConnectionClose  RulesetNewParamsRulesAction = "force_connection_close"
	RulesetNewParamsRulesActionJSChallenge           RulesetNewParamsRulesAction = "js_challenge"
	RulesetNewParamsRulesActionLog                   RulesetNewParamsRulesAction = "log"
	RulesetNewParamsRulesActionLogCustomField        RulesetNewParamsRulesAction = "log_custom_field"
	RulesetNewParamsRulesActionManagedChallenge      RulesetNewParamsRulesAction = "managed_challenge"
	RulesetNewParamsRulesActionRedirect              RulesetNewParamsRulesAction = "redirect"
	RulesetNewParamsRulesActionRewrite               RulesetNewParamsRulesAction = "rewrite"
	RulesetNewParamsRulesActionRoute                 RulesetNewParamsRulesAction = "route"
	RulesetNewParamsRulesActionScore                 RulesetNewParamsRulesAction = "score"
	RulesetNewParamsRulesActionServeError            RulesetNewParamsRulesAction = "serve_error"
	RulesetNewParamsRulesActionSetCacheControl       RulesetNewParamsRulesAction = "set_cache_control"
	RulesetNewParamsRulesActionSetCacheSettings      RulesetNewParamsRulesAction = "set_cache_settings"
	RulesetNewParamsRulesActionSetCacheTags          RulesetNewParamsRulesAction = "set_cache_tags"
	RulesetNewParamsRulesActionSetConfig             RulesetNewParamsRulesAction = "set_config"
	RulesetNewParamsRulesActionSkip                  RulesetNewParamsRulesAction = "skip"
	RulesetNewParamsRulesActionTransformResponseHTML RulesetNewParamsRulesAction = "transform_response_html"
)

func (r RulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesActionBlock, RulesetNewParamsRulesActionChallenge, RulesetNewParamsRulesActionCompressResponse, RulesetNewParamsRulesActionDDoSDynamic, RulesetNewParamsRulesActionExecute, RulesetNewParamsRulesActionForceConnectionClose, RulesetNewParamsRulesActionJSChallenge, RulesetNewParamsRulesActionLog, RulesetNewParamsRulesActionLogCustomField, RulesetNewParamsRulesActionManagedChallenge, RulesetNewParamsRulesActionRedirect, RulesetNewParamsRulesActionRewrite, RulesetNewParamsRulesActionRoute, RulesetNewParamsRulesActionScore, RulesetNewParamsRulesActionServeError, RulesetNewParamsRulesActionSetCacheControl, RulesetNewParamsRulesActionSetCacheSettings, RulesetNewParamsRulesActionSetCacheTags, RulesetNewParamsRulesActionSetConfig, RulesetNewParamsRulesActionSkip, RulesetNewParamsRulesActionTransformResponseHTML:
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
	// A result.
	Result RulesetNewResponseEnvelopeResult `json:"result" api:"required"`
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

// A ruleset object.
type RulesetNewResponseEnvelopeResult struct {
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
	Rules []RulesetNewResponseEnvelopeResultRules `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                               `json:"description"`
	JSON        rulesetNewResponseEnvelopeResultJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultJSON contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResult]
type rulesetNewResponseEnvelopeResultJSON struct {
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

func (r *RulesetNewResponseEnvelopeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseEnvelopeResultRules struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                                    `json:"ref"`
	JSON  rulesetNewResponseEnvelopeResultRulesJSON `json:"-"`
	union RulesetNewResponseEnvelopeResultRulesUnion
}

// rulesetNewResponseEnvelopeResultRulesJSON contains the JSON metadata for the
// struct [RulesetNewResponseEnvelopeResultRules]
type rulesetNewResponseEnvelopeResultRulesJSON struct {
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

func (r rulesetNewResponseEnvelopeResultRulesJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRules) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRules{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetNewResponseEnvelopeResultRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
func (r RulesetNewResponseEnvelopeResultRules) AsUnion() RulesetNewResponseEnvelopeResultRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule] or
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
type RulesetNewResponseEnvelopeResultRulesUnion interface {
	implementsRulesetNewResponseEnvelopeResultRules()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                         `json:"ref"`
	JSON rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule]
type rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRule) implementsRulesetNewResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleAction string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                               `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                  `json:"score_response_header_name"`
	JSON                    rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit]
type rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                           `json:"ref"`
	JSON rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule]
type rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRule) implementsRulesetNewResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                 `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                    `json:"score_response_header_name"`
	JSON                    rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit]
type rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                               `json:"ref"`
	JSON rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRule) implementsRulesetNewResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                      `json:"value"`
	JSON  rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                           `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                              `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                           `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                              `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                 `json:"qualifiers"`
	JSON       rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                `json:"qualifiers"`
	JSON       rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                           `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                 `json:"qualifiers"`
	JSON       rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                `json:"qualifiers"`
	JSON       rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                               `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                       `json:"value"`
	JSON  rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                            `json:"value"`
	JSON  rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                                    `json:"value"`
	JSON  rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                    `json:"cloudflare_only"`
	JSON           rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                     `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                        `json:"score_response_header_name"`
	JSON                    rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                            `json:"ref"`
	JSON rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) implementsRulesetNewResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                                       `json:"values"`
	JSON   rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                            `json:"values" api:"required"`
	JSON   rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                               `json:"values" api:"required"`
	JSON   rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                            `json:"values" api:"required"`
	JSON   rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                  `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                     `json:"score_response_header_name"`
	JSON                    rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit]
type rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                                     `json:"ref"`
	JSON rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule]
type rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) implementsRulesetNewResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction string

const (
	RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                                `json:"link_maze" api:"required"`
	JSON     rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters]
type rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                           `json:"username_expression" api:"required"`
	JSON               rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck]
type rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                              `json:"score_response_header_name"`
	JSON                    rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit]
type rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *RulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetNewResponseEnvelopeResultRulesAction string

const (
	RulesetNewResponseEnvelopeResultRulesActionBlock                 RulesetNewResponseEnvelopeResultRulesAction = "block"
	RulesetNewResponseEnvelopeResultRulesActionChallenge             RulesetNewResponseEnvelopeResultRulesAction = "challenge"
	RulesetNewResponseEnvelopeResultRulesActionCompressResponse      RulesetNewResponseEnvelopeResultRulesAction = "compress_response"
	RulesetNewResponseEnvelopeResultRulesActionDDoSDynamic           RulesetNewResponseEnvelopeResultRulesAction = "ddos_dynamic"
	RulesetNewResponseEnvelopeResultRulesActionExecute               RulesetNewResponseEnvelopeResultRulesAction = "execute"
	RulesetNewResponseEnvelopeResultRulesActionForceConnectionClose  RulesetNewResponseEnvelopeResultRulesAction = "force_connection_close"
	RulesetNewResponseEnvelopeResultRulesActionJSChallenge           RulesetNewResponseEnvelopeResultRulesAction = "js_challenge"
	RulesetNewResponseEnvelopeResultRulesActionLog                   RulesetNewResponseEnvelopeResultRulesAction = "log"
	RulesetNewResponseEnvelopeResultRulesActionLogCustomField        RulesetNewResponseEnvelopeResultRulesAction = "log_custom_field"
	RulesetNewResponseEnvelopeResultRulesActionManagedChallenge      RulesetNewResponseEnvelopeResultRulesAction = "managed_challenge"
	RulesetNewResponseEnvelopeResultRulesActionRedirect              RulesetNewResponseEnvelopeResultRulesAction = "redirect"
	RulesetNewResponseEnvelopeResultRulesActionRewrite               RulesetNewResponseEnvelopeResultRulesAction = "rewrite"
	RulesetNewResponseEnvelopeResultRulesActionRoute                 RulesetNewResponseEnvelopeResultRulesAction = "route"
	RulesetNewResponseEnvelopeResultRulesActionScore                 RulesetNewResponseEnvelopeResultRulesAction = "score"
	RulesetNewResponseEnvelopeResultRulesActionServeError            RulesetNewResponseEnvelopeResultRulesAction = "serve_error"
	RulesetNewResponseEnvelopeResultRulesActionSetCacheControl       RulesetNewResponseEnvelopeResultRulesAction = "set_cache_control"
	RulesetNewResponseEnvelopeResultRulesActionSetCacheSettings      RulesetNewResponseEnvelopeResultRulesAction = "set_cache_settings"
	RulesetNewResponseEnvelopeResultRulesActionSetCacheTags          RulesetNewResponseEnvelopeResultRulesAction = "set_cache_tags"
	RulesetNewResponseEnvelopeResultRulesActionSetConfig             RulesetNewResponseEnvelopeResultRulesAction = "set_config"
	RulesetNewResponseEnvelopeResultRulesActionSkip                  RulesetNewResponseEnvelopeResultRulesAction = "skip"
	RulesetNewResponseEnvelopeResultRulesActionTransformResponseHTML RulesetNewResponseEnvelopeResultRulesAction = "transform_response_html"
)

func (r RulesetNewResponseEnvelopeResultRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeResultRulesActionBlock, RulesetNewResponseEnvelopeResultRulesActionChallenge, RulesetNewResponseEnvelopeResultRulesActionCompressResponse, RulesetNewResponseEnvelopeResultRulesActionDDoSDynamic, RulesetNewResponseEnvelopeResultRulesActionExecute, RulesetNewResponseEnvelopeResultRulesActionForceConnectionClose, RulesetNewResponseEnvelopeResultRulesActionJSChallenge, RulesetNewResponseEnvelopeResultRulesActionLog, RulesetNewResponseEnvelopeResultRulesActionLogCustomField, RulesetNewResponseEnvelopeResultRulesActionManagedChallenge, RulesetNewResponseEnvelopeResultRulesActionRedirect, RulesetNewResponseEnvelopeResultRulesActionRewrite, RulesetNewResponseEnvelopeResultRulesActionRoute, RulesetNewResponseEnvelopeResultRulesActionScore, RulesetNewResponseEnvelopeResultRulesActionServeError, RulesetNewResponseEnvelopeResultRulesActionSetCacheControl, RulesetNewResponseEnvelopeResultRulesActionSetCacheSettings, RulesetNewResponseEnvelopeResultRulesActionSetCacheTags, RulesetNewResponseEnvelopeResultRulesActionSetConfig, RulesetNewResponseEnvelopeResultRulesActionSkip, RulesetNewResponseEnvelopeResultRulesActionTransformResponseHTML:
		return true
	}
	return false
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
	// Validates the request without persisting changes when set to `true`. Responses
	// that normally return 200 return `result: null`; endpoints that normally return
	// 204 continue to return 204.
	DryRun param.Field[bool] `query:"dry_run"`
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

// URLQuery serializes [RulesetUpdateParams]'s query parameters as `url.Values`.
func (r RulesetUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
// [rulesets.RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRule],
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

type RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRule) implementsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze param.Field[interface{}] `json:"link_maze" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for exposed credential checking.
type RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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

func (r RulesetUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesAction string

const (
	RulesetUpdateParamsRulesActionBlock                 RulesetUpdateParamsRulesAction = "block"
	RulesetUpdateParamsRulesActionChallenge             RulesetUpdateParamsRulesAction = "challenge"
	RulesetUpdateParamsRulesActionCompressResponse      RulesetUpdateParamsRulesAction = "compress_response"
	RulesetUpdateParamsRulesActionDDoSDynamic           RulesetUpdateParamsRulesAction = "ddos_dynamic"
	RulesetUpdateParamsRulesActionExecute               RulesetUpdateParamsRulesAction = "execute"
	RulesetUpdateParamsRulesActionForceConnectionClose  RulesetUpdateParamsRulesAction = "force_connection_close"
	RulesetUpdateParamsRulesActionJSChallenge           RulesetUpdateParamsRulesAction = "js_challenge"
	RulesetUpdateParamsRulesActionLog                   RulesetUpdateParamsRulesAction = "log"
	RulesetUpdateParamsRulesActionLogCustomField        RulesetUpdateParamsRulesAction = "log_custom_field"
	RulesetUpdateParamsRulesActionManagedChallenge      RulesetUpdateParamsRulesAction = "managed_challenge"
	RulesetUpdateParamsRulesActionRedirect              RulesetUpdateParamsRulesAction = "redirect"
	RulesetUpdateParamsRulesActionRewrite               RulesetUpdateParamsRulesAction = "rewrite"
	RulesetUpdateParamsRulesActionRoute                 RulesetUpdateParamsRulesAction = "route"
	RulesetUpdateParamsRulesActionScore                 RulesetUpdateParamsRulesAction = "score"
	RulesetUpdateParamsRulesActionServeError            RulesetUpdateParamsRulesAction = "serve_error"
	RulesetUpdateParamsRulesActionSetCacheControl       RulesetUpdateParamsRulesAction = "set_cache_control"
	RulesetUpdateParamsRulesActionSetCacheSettings      RulesetUpdateParamsRulesAction = "set_cache_settings"
	RulesetUpdateParamsRulesActionSetCacheTags          RulesetUpdateParamsRulesAction = "set_cache_tags"
	RulesetUpdateParamsRulesActionSetConfig             RulesetUpdateParamsRulesAction = "set_config"
	RulesetUpdateParamsRulesActionSkip                  RulesetUpdateParamsRulesAction = "skip"
	RulesetUpdateParamsRulesActionTransformResponseHTML RulesetUpdateParamsRulesAction = "transform_response_html"
)

func (r RulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesActionBlock, RulesetUpdateParamsRulesActionChallenge, RulesetUpdateParamsRulesActionCompressResponse, RulesetUpdateParamsRulesActionDDoSDynamic, RulesetUpdateParamsRulesActionExecute, RulesetUpdateParamsRulesActionForceConnectionClose, RulesetUpdateParamsRulesActionJSChallenge, RulesetUpdateParamsRulesActionLog, RulesetUpdateParamsRulesActionLogCustomField, RulesetUpdateParamsRulesActionManagedChallenge, RulesetUpdateParamsRulesActionRedirect, RulesetUpdateParamsRulesActionRewrite, RulesetUpdateParamsRulesActionRoute, RulesetUpdateParamsRulesActionScore, RulesetUpdateParamsRulesActionServeError, RulesetUpdateParamsRulesActionSetCacheControl, RulesetUpdateParamsRulesActionSetCacheSettings, RulesetUpdateParamsRulesActionSetCacheTags, RulesetUpdateParamsRulesActionSetConfig, RulesetUpdateParamsRulesActionSkip, RulesetUpdateParamsRulesActionTransformResponseHTML:
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
	// A result.
	Result RulesetUpdateResponseEnvelopeResult `json:"result" api:"required"`
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

// A ruleset object.
type RulesetUpdateResponseEnvelopeResult struct {
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
	Rules []RulesetUpdateResponseEnvelopeResultRules `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                                  `json:"description"`
	JSON        rulesetUpdateResponseEnvelopeResultJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeResult]
type rulesetUpdateResponseEnvelopeResultJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseEnvelopeResultRules struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                                       `json:"ref"`
	JSON  rulesetUpdateResponseEnvelopeResultRulesJSON `json:"-"`
	union RulesetUpdateResponseEnvelopeResultRulesUnion
}

// rulesetUpdateResponseEnvelopeResultRulesJSON contains the JSON metadata for the
// struct [RulesetUpdateResponseEnvelopeResultRules]
type rulesetUpdateResponseEnvelopeResultRulesJSON struct {
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

func (r rulesetUpdateResponseEnvelopeResultRulesJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRules) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRules{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [RulesetUpdateResponseEnvelopeResultRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
func (r RulesetUpdateResponseEnvelopeResultRules) AsUnion() RulesetUpdateResponseEnvelopeResultRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule] or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
type RulesetUpdateResponseEnvelopeResultRulesUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRules()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule{}),
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
			Type:               reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                            `json:"ref"`
	JSON rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRule) implementsRulesetUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                  `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                     `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                              `json:"ref"`
	JSON rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule) implementsRulesetUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                    `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                       `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                                  `json:"ref"`
	JSON rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule) implementsRulesetUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                         `json:"value"`
	JSON  rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                              `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                              `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                 `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                    `json:"qualifiers"`
	JSON       rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                   `json:"qualifiers"`
	JSON       rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                           `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                              `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                    `json:"qualifiers"`
	JSON       rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                   `json:"qualifiers"`
	JSON       rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                               `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                  `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                          `json:"value"`
	JSON  rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                               `json:"value"`
	JSON  rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                               `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                                       `json:"value"`
	JSON  rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                    `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                       `json:"cloudflare_only"`
	JSON           rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                        `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                           `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                               `json:"ref"`
	JSON rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) implementsRulesetUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                                          `json:"values"`
	JSON   rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                               `json:"values" api:"required"`
	JSON   rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                                  `json:"values" api:"required"`
	JSON   rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                               `json:"values" api:"required"`
	JSON   rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsRulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                     `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                        `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                                        `json:"ref"`
	JSON rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) implementsRulesetUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                                   `json:"link_maze" api:"required"`
	JSON     rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                              `json:"username_expression" api:"required"`
	JSON               rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                                 `json:"score_response_header_name"`
	JSON                    rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit]
type rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *RulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetUpdateResponseEnvelopeResultRulesAction string

const (
	RulesetUpdateResponseEnvelopeResultRulesActionBlock                 RulesetUpdateResponseEnvelopeResultRulesAction = "block"
	RulesetUpdateResponseEnvelopeResultRulesActionChallenge             RulesetUpdateResponseEnvelopeResultRulesAction = "challenge"
	RulesetUpdateResponseEnvelopeResultRulesActionCompressResponse      RulesetUpdateResponseEnvelopeResultRulesAction = "compress_response"
	RulesetUpdateResponseEnvelopeResultRulesActionDDoSDynamic           RulesetUpdateResponseEnvelopeResultRulesAction = "ddos_dynamic"
	RulesetUpdateResponseEnvelopeResultRulesActionExecute               RulesetUpdateResponseEnvelopeResultRulesAction = "execute"
	RulesetUpdateResponseEnvelopeResultRulesActionForceConnectionClose  RulesetUpdateResponseEnvelopeResultRulesAction = "force_connection_close"
	RulesetUpdateResponseEnvelopeResultRulesActionJSChallenge           RulesetUpdateResponseEnvelopeResultRulesAction = "js_challenge"
	RulesetUpdateResponseEnvelopeResultRulesActionLog                   RulesetUpdateResponseEnvelopeResultRulesAction = "log"
	RulesetUpdateResponseEnvelopeResultRulesActionLogCustomField        RulesetUpdateResponseEnvelopeResultRulesAction = "log_custom_field"
	RulesetUpdateResponseEnvelopeResultRulesActionManagedChallenge      RulesetUpdateResponseEnvelopeResultRulesAction = "managed_challenge"
	RulesetUpdateResponseEnvelopeResultRulesActionRedirect              RulesetUpdateResponseEnvelopeResultRulesAction = "redirect"
	RulesetUpdateResponseEnvelopeResultRulesActionRewrite               RulesetUpdateResponseEnvelopeResultRulesAction = "rewrite"
	RulesetUpdateResponseEnvelopeResultRulesActionRoute                 RulesetUpdateResponseEnvelopeResultRulesAction = "route"
	RulesetUpdateResponseEnvelopeResultRulesActionScore                 RulesetUpdateResponseEnvelopeResultRulesAction = "score"
	RulesetUpdateResponseEnvelopeResultRulesActionServeError            RulesetUpdateResponseEnvelopeResultRulesAction = "serve_error"
	RulesetUpdateResponseEnvelopeResultRulesActionSetCacheControl       RulesetUpdateResponseEnvelopeResultRulesAction = "set_cache_control"
	RulesetUpdateResponseEnvelopeResultRulesActionSetCacheSettings      RulesetUpdateResponseEnvelopeResultRulesAction = "set_cache_settings"
	RulesetUpdateResponseEnvelopeResultRulesActionSetCacheTags          RulesetUpdateResponseEnvelopeResultRulesAction = "set_cache_tags"
	RulesetUpdateResponseEnvelopeResultRulesActionSetConfig             RulesetUpdateResponseEnvelopeResultRulesAction = "set_config"
	RulesetUpdateResponseEnvelopeResultRulesActionSkip                  RulesetUpdateResponseEnvelopeResultRulesAction = "skip"
	RulesetUpdateResponseEnvelopeResultRulesActionTransformResponseHTML RulesetUpdateResponseEnvelopeResultRulesAction = "transform_response_html"
)

func (r RulesetUpdateResponseEnvelopeResultRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseEnvelopeResultRulesActionBlock, RulesetUpdateResponseEnvelopeResultRulesActionChallenge, RulesetUpdateResponseEnvelopeResultRulesActionCompressResponse, RulesetUpdateResponseEnvelopeResultRulesActionDDoSDynamic, RulesetUpdateResponseEnvelopeResultRulesActionExecute, RulesetUpdateResponseEnvelopeResultRulesActionForceConnectionClose, RulesetUpdateResponseEnvelopeResultRulesActionJSChallenge, RulesetUpdateResponseEnvelopeResultRulesActionLog, RulesetUpdateResponseEnvelopeResultRulesActionLogCustomField, RulesetUpdateResponseEnvelopeResultRulesActionManagedChallenge, RulesetUpdateResponseEnvelopeResultRulesActionRedirect, RulesetUpdateResponseEnvelopeResultRulesActionRewrite, RulesetUpdateResponseEnvelopeResultRulesActionRoute, RulesetUpdateResponseEnvelopeResultRulesActionScore, RulesetUpdateResponseEnvelopeResultRulesActionServeError, RulesetUpdateResponseEnvelopeResultRulesActionSetCacheControl, RulesetUpdateResponseEnvelopeResultRulesActionSetCacheSettings, RulesetUpdateResponseEnvelopeResultRulesActionSetCacheTags, RulesetUpdateResponseEnvelopeResultRulesActionSetConfig, RulesetUpdateResponseEnvelopeResultRulesActionSkip, RulesetUpdateResponseEnvelopeResultRulesActionTransformResponseHTML:
		return true
	}
	return false
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
	// Validates the request without persisting changes when set to `true`. Responses
	// that normally return 200 return `result: null`; endpoints that normally return
	// 204 continue to return 204.
	DryRun param.Field[bool] `query:"dry_run"`
}

// URLQuery serializes [RulesetDeleteParams]'s query parameters as `url.Values`.
func (r RulesetDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
