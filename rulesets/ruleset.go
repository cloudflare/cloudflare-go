// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
	opts = append(r.Options[:], opts...)
	var env RulesetNewResponseEnvelope
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
		return
	}
	res = &env.Result
	return
}

// Updates an account or zone ruleset, creating a new version.
func (r *RulesetService) Update(ctx context.Context, rulesetID string, params RulesetUpdateParams, opts ...option.RequestOption) (res *RulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetUpdateResponseEnvelope
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
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all rulesets.
func (r *RulesetService) List(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ruleset], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *RulesetService) ListAutoPaging(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ruleset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes all versions of an existing account or zone ruleset.
func (r *RulesetService) Delete(ctx context.Context, rulesetID string, body RulesetDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
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
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the latest version of an account or zone ruleset.
func (r *RulesetService) Get(ctx context.Context, rulesetID string, query RulesetGetParams, opts ...option.RequestOption) (res *RulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetGetResponseEnvelope
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
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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
	PhaseDDoSL4                         Phase = "ddos_l4"
	PhaseDDoSL7                         Phase = "ddos_l7"
	PhaseHTTPConfigSettings             Phase = "http_config_settings"
	PhaseHTTPCustomErrors               Phase = "http_custom_errors"
	PhaseHTTPLogCustomFields            Phase = "http_log_custom_fields"
	PhaseHTTPRatelimit                  Phase = "http_ratelimit"
	PhaseHTTPRequestCacheSettings       Phase = "http_request_cache_settings"
	PhaseHTTPRequestDynamicRedirect     Phase = "http_request_dynamic_redirect"
	PhaseHTTPRequestFirewallCustom      Phase = "http_request_firewall_custom"
	PhaseHTTPRequestFirewallManaged     Phase = "http_request_firewall_managed"
	PhaseHTTPRequestLateTransform       Phase = "http_request_late_transform"
	PhaseHTTPRequestOrigin              Phase = "http_request_origin"
	PhaseHTTPRequestRedirect            Phase = "http_request_redirect"
	PhaseHTTPRequestSanitize            Phase = "http_request_sanitize"
	PhaseHTTPRequestSBFM                Phase = "http_request_sbfm"
	PhaseHTTPRequestSelectConfiguration Phase = "http_request_select_configuration"
	PhaseHTTPRequestTransform           Phase = "http_request_transform"
	PhaseHTTPResponseCompression        Phase = "http_response_compression"
	PhaseHTTPResponseFirewallManaged    Phase = "http_response_firewall_managed"
	PhaseHTTPResponseHeadersTransform   Phase = "http_response_headers_transform"
	PhaseMagicTransit                   Phase = "magic_transit"
	PhaseMagicTransitIDsManaged         Phase = "magic_transit_ids_managed"
	PhaseMagicTransitManaged            Phase = "magic_transit_managed"
)

func (r Phase) IsKnown() bool {
	switch r {
	case PhaseDDoSL4, PhaseDDoSL7, PhaseHTTPConfigSettings, PhaseHTTPCustomErrors, PhaseHTTPLogCustomFields, PhaseHTTPRatelimit, PhaseHTTPRequestCacheSettings, PhaseHTTPRequestDynamicRedirect, PhaseHTTPRequestFirewallCustom, PhaseHTTPRequestFirewallManaged, PhaseHTTPRequestLateTransform, PhaseHTTPRequestOrigin, PhaseHTTPRequestRedirect, PhaseHTTPRequestSanitize, PhaseHTTPRequestSBFM, PhaseHTTPRequestSelectConfiguration, PhaseHTTPRequestTransform, PhaseHTTPResponseCompression, PhaseHTTPResponseFirewallManaged, PhaseHTTPResponseHeadersTransform, PhaseMagicTransit, PhaseMagicTransitIDsManaged, PhaseMagicTransitManaged:
		return true
	}
	return false
}

// A ruleset object.
type Ruleset struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind Kind `json:"kind"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase Phase       `json:"phase"`
	JSON  rulesetJSON `json:"-"`
}

// rulesetJSON contains the JSON metadata for the struct [Ruleset]
type rulesetJSON struct {
	ID          apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ruleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetJSON) RawJSON() string {
	return r.raw
}

// A ruleset object.
type RulesetNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetNewResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action           RulesetNewResponseRulesAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                     `json:"version,required"`
	JSON    rulesetNewResponseRuleJSON `json:"-"`
	union   RulesetNewResponseRulesUnion
}

// rulesetNewResponseRuleJSON contains the JSON metadata for the struct
// [RulesetNewResponseRule]
type rulesetNewResponseRuleJSON struct {
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

func (r rulesetNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRule) AsUnion() RulesetNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RulesetNewResponseRulesRulesetsLogCustomFieldRule].
type RulesetNewResponseRulesUnion interface {
	implementsRulesetsRulesetNewResponseRule()
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
			Type:               reflect.TypeOf(JSChallengeRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RulesetNewResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetNewResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsLogCustomFieldRule]
type rulesetNewResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRulesetNewResponseRule() {
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RulesetNewResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RulesetNewResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                           `json:"name,required"`
	JSON rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                            `json:"name,required"`
	JSON rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                             `json:"name,required"`
	JSON rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesAction string

const (
	RulesetNewResponseRulesActionBlock            RulesetNewResponseRulesAction = "block"
	RulesetNewResponseRulesActionChallenge        RulesetNewResponseRulesAction = "challenge"
	RulesetNewResponseRulesActionCompressResponse RulesetNewResponseRulesAction = "compress_response"
	RulesetNewResponseRulesActionExecute          RulesetNewResponseRulesAction = "execute"
	RulesetNewResponseRulesActionJSChallenge      RulesetNewResponseRulesAction = "js_challenge"
	RulesetNewResponseRulesActionLog              RulesetNewResponseRulesAction = "log"
	RulesetNewResponseRulesActionManagedChallenge RulesetNewResponseRulesAction = "managed_challenge"
	RulesetNewResponseRulesActionRedirect         RulesetNewResponseRulesAction = "redirect"
	RulesetNewResponseRulesActionRewrite          RulesetNewResponseRulesAction = "rewrite"
	RulesetNewResponseRulesActionRoute            RulesetNewResponseRulesAction = "route"
	RulesetNewResponseRulesActionScore            RulesetNewResponseRulesAction = "score"
	RulesetNewResponseRulesActionServeError       RulesetNewResponseRulesAction = "serve_error"
	RulesetNewResponseRulesActionSetConfig        RulesetNewResponseRulesAction = "set_config"
	RulesetNewResponseRulesActionSkip             RulesetNewResponseRulesAction = "skip"
	RulesetNewResponseRulesActionSetCacheSettings RulesetNewResponseRulesAction = "set_cache_settings"
	RulesetNewResponseRulesActionLogCustomField   RulesetNewResponseRulesAction = "log_custom_field"
)

func (r RulesetNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesActionBlock, RulesetNewResponseRulesActionChallenge, RulesetNewResponseRulesActionCompressResponse, RulesetNewResponseRulesActionExecute, RulesetNewResponseRulesActionJSChallenge, RulesetNewResponseRulesActionLog, RulesetNewResponseRulesActionManagedChallenge, RulesetNewResponseRulesActionRedirect, RulesetNewResponseRulesActionRewrite, RulesetNewResponseRulesActionRoute, RulesetNewResponseRulesActionScore, RulesetNewResponseRulesActionServeError, RulesetNewResponseRulesActionSetConfig, RulesetNewResponseRulesActionSkip, RulesetNewResponseRulesActionSetCacheSettings, RulesetNewResponseRulesActionLogCustomField:
		return true
	}
	return false
}

// A ruleset object.
type RulesetUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action           RulesetUpdateResponseRulesAction `json:"action"`
	ActionParameters interface{}                      `json:"action_parameters,required"`
	Categories       interface{}                      `json:"categories,required"`
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
	Version string                        `json:"version,required"`
	JSON    rulesetUpdateResponseRuleJSON `json:"-"`
	union   RulesetUpdateResponseRulesUnion
}

// rulesetUpdateResponseRuleJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseRule]
type rulesetUpdateResponseRuleJSON struct {
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

func (r rulesetUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRule) AsUnion() RulesetUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RulesetUpdateResponseRulesRulesetsLogCustomFieldRule].
type RulesetUpdateResponseRulesUnion interface {
	implementsRulesetsRulesetUpdateResponseRule()
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
			Type:               reflect.TypeOf(JSChallengeRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RulesetUpdateResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                   `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsLogCustomFieldRule]
type rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                              `json:"name,required"`
	JSON rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                               `json:"name,required"`
	JSON rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                                `json:"name,required"`
	JSON rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesAction string

const (
	RulesetUpdateResponseRulesActionBlock            RulesetUpdateResponseRulesAction = "block"
	RulesetUpdateResponseRulesActionChallenge        RulesetUpdateResponseRulesAction = "challenge"
	RulesetUpdateResponseRulesActionCompressResponse RulesetUpdateResponseRulesAction = "compress_response"
	RulesetUpdateResponseRulesActionExecute          RulesetUpdateResponseRulesAction = "execute"
	RulesetUpdateResponseRulesActionJSChallenge      RulesetUpdateResponseRulesAction = "js_challenge"
	RulesetUpdateResponseRulesActionLog              RulesetUpdateResponseRulesAction = "log"
	RulesetUpdateResponseRulesActionManagedChallenge RulesetUpdateResponseRulesAction = "managed_challenge"
	RulesetUpdateResponseRulesActionRedirect         RulesetUpdateResponseRulesAction = "redirect"
	RulesetUpdateResponseRulesActionRewrite          RulesetUpdateResponseRulesAction = "rewrite"
	RulesetUpdateResponseRulesActionRoute            RulesetUpdateResponseRulesAction = "route"
	RulesetUpdateResponseRulesActionScore            RulesetUpdateResponseRulesAction = "score"
	RulesetUpdateResponseRulesActionServeError       RulesetUpdateResponseRulesAction = "serve_error"
	RulesetUpdateResponseRulesActionSetConfig        RulesetUpdateResponseRulesAction = "set_config"
	RulesetUpdateResponseRulesActionSkip             RulesetUpdateResponseRulesAction = "skip"
	RulesetUpdateResponseRulesActionSetCacheSettings RulesetUpdateResponseRulesAction = "set_cache_settings"
	RulesetUpdateResponseRulesActionLogCustomField   RulesetUpdateResponseRulesAction = "log_custom_field"
)

func (r RulesetUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesActionBlock, RulesetUpdateResponseRulesActionChallenge, RulesetUpdateResponseRulesActionCompressResponse, RulesetUpdateResponseRulesActionExecute, RulesetUpdateResponseRulesActionJSChallenge, RulesetUpdateResponseRulesActionLog, RulesetUpdateResponseRulesActionManagedChallenge, RulesetUpdateResponseRulesActionRedirect, RulesetUpdateResponseRulesActionRewrite, RulesetUpdateResponseRulesActionRoute, RulesetUpdateResponseRulesActionScore, RulesetUpdateResponseRulesActionServeError, RulesetUpdateResponseRulesActionSetConfig, RulesetUpdateResponseRulesActionSkip, RulesetUpdateResponseRulesActionSetCacheSettings, RulesetUpdateResponseRulesActionLogCustomField:
		return true
	}
	return false
}

// A ruleset object.
type RulesetGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RulesetGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action           RulesetGetResponseRulesAction `json:"action"`
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
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                     `json:"version,required"`
	JSON    rulesetGetResponseRuleJSON `json:"-"`
	union   RulesetGetResponseRulesUnion
}

// rulesetGetResponseRuleJSON contains the JSON metadata for the struct
// [RulesetGetResponseRule]
type rulesetGetResponseRuleJSON struct {
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

func (r rulesetGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRule) AsUnion() RulesetGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RulesetGetResponseRulesRulesetsLogCustomFieldRule].
type RulesetGetResponseRulesUnion interface {
	implementsRulesetsRulesetGetResponseRule()
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
			Type:               reflect.TypeOf(JSChallengeRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RulesetGetResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON rulesetGetResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsLogCustomFieldRule]
type rulesetGetResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRulesetGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RulesetGetResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RulesetGetResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                           `json:"name,required"`
	JSON rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                            `json:"name,required"`
	JSON rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                             `json:"name,required"`
	JSON rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesAction string

const (
	RulesetGetResponseRulesActionBlock            RulesetGetResponseRulesAction = "block"
	RulesetGetResponseRulesActionChallenge        RulesetGetResponseRulesAction = "challenge"
	RulesetGetResponseRulesActionCompressResponse RulesetGetResponseRulesAction = "compress_response"
	RulesetGetResponseRulesActionExecute          RulesetGetResponseRulesAction = "execute"
	RulesetGetResponseRulesActionJSChallenge      RulesetGetResponseRulesAction = "js_challenge"
	RulesetGetResponseRulesActionLog              RulesetGetResponseRulesAction = "log"
	RulesetGetResponseRulesActionManagedChallenge RulesetGetResponseRulesAction = "managed_challenge"
	RulesetGetResponseRulesActionRedirect         RulesetGetResponseRulesAction = "redirect"
	RulesetGetResponseRulesActionRewrite          RulesetGetResponseRulesAction = "rewrite"
	RulesetGetResponseRulesActionRoute            RulesetGetResponseRulesAction = "route"
	RulesetGetResponseRulesActionScore            RulesetGetResponseRulesAction = "score"
	RulesetGetResponseRulesActionServeError       RulesetGetResponseRulesAction = "serve_error"
	RulesetGetResponseRulesActionSetConfig        RulesetGetResponseRulesAction = "set_config"
	RulesetGetResponseRulesActionSkip             RulesetGetResponseRulesAction = "skip"
	RulesetGetResponseRulesActionSetCacheSettings RulesetGetResponseRulesAction = "set_cache_settings"
	RulesetGetResponseRulesActionLogCustomField   RulesetGetResponseRulesAction = "log_custom_field"
)

func (r RulesetGetResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesActionBlock, RulesetGetResponseRulesActionChallenge, RulesetGetResponseRulesActionCompressResponse, RulesetGetResponseRulesActionExecute, RulesetGetResponseRulesActionJSChallenge, RulesetGetResponseRulesActionLog, RulesetGetResponseRulesActionManagedChallenge, RulesetGetResponseRulesActionRedirect, RulesetGetResponseRulesActionRewrite, RulesetGetResponseRulesActionRoute, RulesetGetResponseRulesActionScore, RulesetGetResponseRulesActionServeError, RulesetGetResponseRulesActionSetConfig, RulesetGetResponseRulesActionSkip, RulesetGetResponseRulesActionSetCacheSettings, RulesetGetResponseRulesActionLogCustomField:
		return true
	}
	return false
}

type RulesetNewParams struct {
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetNewParamsRuleUnion] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r RulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[RulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                 `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                 `json:"categories,required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ChallengeRuleParam],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.JSChallengeRuleParam], [rulesets.LogRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsLogCustomFieldRule],
// [RulesetNewParamsRule].
type RulesetNewParamsRuleUnion interface {
	implementsRulesetsRulesetNewParamsRuleUnion()
}

type RulesetNewParamsRulesRulesetsLogCustomFieldRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsLogCustomFieldRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRule) implementsRulesetsRulesetNewParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsLogCustomFieldRuleAction string

const (
	RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField RulesetNewParamsRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields param.Field[[]RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField] `json:"cookie_fields"`
	// The request fields to log.
	RequestFields param.Field[[]RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField] `json:"request_fields"`
	// The response fields to log.
	ResponseFields param.Field[[]RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField] `json:"response_fields"`
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookie field to log.
type RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request field to log.
type RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response field to log.
type RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetNewParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesAction string

const (
	RulesetNewParamsRulesActionBlock            RulesetNewParamsRulesAction = "block"
	RulesetNewParamsRulesActionChallenge        RulesetNewParamsRulesAction = "challenge"
	RulesetNewParamsRulesActionCompressResponse RulesetNewParamsRulesAction = "compress_response"
	RulesetNewParamsRulesActionExecute          RulesetNewParamsRulesAction = "execute"
	RulesetNewParamsRulesActionJSChallenge      RulesetNewParamsRulesAction = "js_challenge"
	RulesetNewParamsRulesActionLog              RulesetNewParamsRulesAction = "log"
	RulesetNewParamsRulesActionManagedChallenge RulesetNewParamsRulesAction = "managed_challenge"
	RulesetNewParamsRulesActionRedirect         RulesetNewParamsRulesAction = "redirect"
	RulesetNewParamsRulesActionRewrite          RulesetNewParamsRulesAction = "rewrite"
	RulesetNewParamsRulesActionRoute            RulesetNewParamsRulesAction = "route"
	RulesetNewParamsRulesActionScore            RulesetNewParamsRulesAction = "score"
	RulesetNewParamsRulesActionServeError       RulesetNewParamsRulesAction = "serve_error"
	RulesetNewParamsRulesActionSetConfig        RulesetNewParamsRulesAction = "set_config"
	RulesetNewParamsRulesActionSkip             RulesetNewParamsRulesAction = "skip"
	RulesetNewParamsRulesActionSetCacheSettings RulesetNewParamsRulesAction = "set_cache_settings"
	RulesetNewParamsRulesActionLogCustomField   RulesetNewParamsRulesAction = "log_custom_field"
)

func (r RulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesActionBlock, RulesetNewParamsRulesActionChallenge, RulesetNewParamsRulesActionCompressResponse, RulesetNewParamsRulesActionExecute, RulesetNewParamsRulesActionJSChallenge, RulesetNewParamsRulesActionLog, RulesetNewParamsRulesActionManagedChallenge, RulesetNewParamsRulesActionRedirect, RulesetNewParamsRulesActionRewrite, RulesetNewParamsRulesActionRoute, RulesetNewParamsRulesActionScore, RulesetNewParamsRulesActionServeError, RulesetNewParamsRulesActionSetConfig, RulesetNewParamsRulesActionSkip, RulesetNewParamsRulesActionSetCacheSettings, RulesetNewParamsRulesActionLogCustomField:
		return true
	}
	return false
}

// A response object.
type RulesetNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetNewResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                     `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                       `json:"pointer,required"`
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
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetUpdateParamsRuleUnion] `json:"rules,required"`
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
}

func (r RulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[RulesetUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                    `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                    `json:"categories,required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ChallengeRuleParam],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.JSChallengeRuleParam], [rulesets.LogRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsLogCustomFieldRule],
// [RulesetUpdateParamsRule].
type RulesetUpdateParamsRuleUnion interface {
	implementsRulesetsRulesetUpdateParamsRuleUnion()
}

type RulesetUpdateParamsRulesRulesetsLogCustomFieldRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields param.Field[[]RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField] `json:"cookie_fields"`
	// The request fields to log.
	RequestFields param.Field[[]RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField] `json:"request_fields"`
	// The response fields to log.
	ResponseFields param.Field[[]RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField] `json:"response_fields"`
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookie field to log.
type RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request field to log.
type RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response field to log.
type RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r RulesetUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesAction string

const (
	RulesetUpdateParamsRulesActionBlock            RulesetUpdateParamsRulesAction = "block"
	RulesetUpdateParamsRulesActionChallenge        RulesetUpdateParamsRulesAction = "challenge"
	RulesetUpdateParamsRulesActionCompressResponse RulesetUpdateParamsRulesAction = "compress_response"
	RulesetUpdateParamsRulesActionExecute          RulesetUpdateParamsRulesAction = "execute"
	RulesetUpdateParamsRulesActionJSChallenge      RulesetUpdateParamsRulesAction = "js_challenge"
	RulesetUpdateParamsRulesActionLog              RulesetUpdateParamsRulesAction = "log"
	RulesetUpdateParamsRulesActionManagedChallenge RulesetUpdateParamsRulesAction = "managed_challenge"
	RulesetUpdateParamsRulesActionRedirect         RulesetUpdateParamsRulesAction = "redirect"
	RulesetUpdateParamsRulesActionRewrite          RulesetUpdateParamsRulesAction = "rewrite"
	RulesetUpdateParamsRulesActionRoute            RulesetUpdateParamsRulesAction = "route"
	RulesetUpdateParamsRulesActionScore            RulesetUpdateParamsRulesAction = "score"
	RulesetUpdateParamsRulesActionServeError       RulesetUpdateParamsRulesAction = "serve_error"
	RulesetUpdateParamsRulesActionSetConfig        RulesetUpdateParamsRulesAction = "set_config"
	RulesetUpdateParamsRulesActionSkip             RulesetUpdateParamsRulesAction = "skip"
	RulesetUpdateParamsRulesActionSetCacheSettings RulesetUpdateParamsRulesAction = "set_cache_settings"
	RulesetUpdateParamsRulesActionLogCustomField   RulesetUpdateParamsRulesAction = "log_custom_field"
)

func (r RulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesActionBlock, RulesetUpdateParamsRulesActionChallenge, RulesetUpdateParamsRulesActionCompressResponse, RulesetUpdateParamsRulesActionExecute, RulesetUpdateParamsRulesActionJSChallenge, RulesetUpdateParamsRulesActionLog, RulesetUpdateParamsRulesActionManagedChallenge, RulesetUpdateParamsRulesActionRedirect, RulesetUpdateParamsRulesActionRewrite, RulesetUpdateParamsRulesActionRoute, RulesetUpdateParamsRulesActionScore, RulesetUpdateParamsRulesActionServeError, RulesetUpdateParamsRulesActionSetConfig, RulesetUpdateParamsRulesActionSkip, RulesetUpdateParamsRulesActionSetCacheSettings, RulesetUpdateParamsRulesActionLogCustomField:
		return true
	}
	return false
}

// A response object.
type RulesetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetUpdateResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                        `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                          `json:"pointer,required"`
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
	Errors []RulesetGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RulesetGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetGetResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                     `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                       `json:"pointer,required"`
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
