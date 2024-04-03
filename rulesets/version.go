// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
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

// VersionService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
	ByTag   *VersionByTagService
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r *VersionService) {
	r = &VersionService{}
	r.Options = opts
	r.ByTag = NewVersionByTagService(opts...)
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[VersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
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

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) ListAutoPaging(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[VersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetID, query, opts...))
}

// Deletes an existing version of an account or zone ruleset.
func (r *VersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, body VersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *VersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionListResponse struct {
	// The kind of the ruleset.
	Kind VersionListResponseKind `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionListResponsePhase `json:"phase,required"`
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the ruleset.
	Version string                  `json:"version"`
	JSON    versionListResponseJSON `json:"-"`
}

// versionListResponseJSON contains the JSON metadata for the struct
// [VersionListResponse]
type versionListResponseJSON struct {
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

func (r *VersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionListResponseKind string

const (
	VersionListResponseKindManaged VersionListResponseKind = "managed"
	VersionListResponseKindCustom  VersionListResponseKind = "custom"
	VersionListResponseKindRoot    VersionListResponseKind = "root"
	VersionListResponseKindZone    VersionListResponseKind = "zone"
)

func (r VersionListResponseKind) IsKnown() bool {
	switch r {
	case VersionListResponseKindManaged, VersionListResponseKindCustom, VersionListResponseKindRoot, VersionListResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionListResponsePhase string

const (
	VersionListResponsePhaseDDoSL4                         VersionListResponsePhase = "ddos_l4"
	VersionListResponsePhaseDDoSL7                         VersionListResponsePhase = "ddos_l7"
	VersionListResponsePhaseHTTPConfigSettings             VersionListResponsePhase = "http_config_settings"
	VersionListResponsePhaseHTTPCustomErrors               VersionListResponsePhase = "http_custom_errors"
	VersionListResponsePhaseHTTPLogCustomFields            VersionListResponsePhase = "http_log_custom_fields"
	VersionListResponsePhaseHTTPRatelimit                  VersionListResponsePhase = "http_ratelimit"
	VersionListResponsePhaseHTTPRequestCacheSettings       VersionListResponsePhase = "http_request_cache_settings"
	VersionListResponsePhaseHTTPRequestDynamicRedirect     VersionListResponsePhase = "http_request_dynamic_redirect"
	VersionListResponsePhaseHTTPRequestFirewallCustom      VersionListResponsePhase = "http_request_firewall_custom"
	VersionListResponsePhaseHTTPRequestFirewallManaged     VersionListResponsePhase = "http_request_firewall_managed"
	VersionListResponsePhaseHTTPRequestLateTransform       VersionListResponsePhase = "http_request_late_transform"
	VersionListResponsePhaseHTTPRequestOrigin              VersionListResponsePhase = "http_request_origin"
	VersionListResponsePhaseHTTPRequestRedirect            VersionListResponsePhase = "http_request_redirect"
	VersionListResponsePhaseHTTPRequestSanitize            VersionListResponsePhase = "http_request_sanitize"
	VersionListResponsePhaseHTTPRequestSbfm                VersionListResponsePhase = "http_request_sbfm"
	VersionListResponsePhaseHTTPRequestSelectConfiguration VersionListResponsePhase = "http_request_select_configuration"
	VersionListResponsePhaseHTTPRequestTransform           VersionListResponsePhase = "http_request_transform"
	VersionListResponsePhaseHTTPResponseCompression        VersionListResponsePhase = "http_response_compression"
	VersionListResponsePhaseHTTPResponseFirewallManaged    VersionListResponsePhase = "http_response_firewall_managed"
	VersionListResponsePhaseHTTPResponseHeadersTransform   VersionListResponsePhase = "http_response_headers_transform"
	VersionListResponsePhaseMagicTransit                   VersionListResponsePhase = "magic_transit"
	VersionListResponsePhaseMagicTransitIDsManaged         VersionListResponsePhase = "magic_transit_ids_managed"
	VersionListResponsePhaseMagicTransitManaged            VersionListResponsePhase = "magic_transit_managed"
)

func (r VersionListResponsePhase) IsKnown() bool {
	switch r {
	case VersionListResponsePhaseDDoSL4, VersionListResponsePhaseDDoSL7, VersionListResponsePhaseHTTPConfigSettings, VersionListResponsePhaseHTTPCustomErrors, VersionListResponsePhaseHTTPLogCustomFields, VersionListResponsePhaseHTTPRatelimit, VersionListResponsePhaseHTTPRequestCacheSettings, VersionListResponsePhaseHTTPRequestDynamicRedirect, VersionListResponsePhaseHTTPRequestFirewallCustom, VersionListResponsePhaseHTTPRequestFirewallManaged, VersionListResponsePhaseHTTPRequestLateTransform, VersionListResponsePhaseHTTPRequestOrigin, VersionListResponsePhaseHTTPRequestRedirect, VersionListResponsePhaseHTTPRequestSanitize, VersionListResponsePhaseHTTPRequestSbfm, VersionListResponsePhaseHTTPRequestSelectConfiguration, VersionListResponsePhaseHTTPRequestTransform, VersionListResponsePhaseHTTPResponseCompression, VersionListResponsePhaseHTTPResponseFirewallManaged, VersionListResponsePhaseHTTPResponseHeadersTransform, VersionListResponsePhaseMagicTransit, VersionListResponsePhaseMagicTransitIDsManaged, VersionListResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// A ruleset object.
type VersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind VersionGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase VersionGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []VersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        versionGetResponseJSON `json:"-"`
}

// versionGetResponseJSON contains the JSON metadata for the struct
// [VersionGetResponse]
type versionGetResponseJSON struct {
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

func (r *VersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type VersionGetResponseKind string

const (
	VersionGetResponseKindManaged VersionGetResponseKind = "managed"
	VersionGetResponseKindCustom  VersionGetResponseKind = "custom"
	VersionGetResponseKindRoot    VersionGetResponseKind = "root"
	VersionGetResponseKindZone    VersionGetResponseKind = "zone"
)

func (r VersionGetResponseKind) IsKnown() bool {
	switch r {
	case VersionGetResponseKindManaged, VersionGetResponseKindCustom, VersionGetResponseKindRoot, VersionGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type VersionGetResponsePhase string

const (
	VersionGetResponsePhaseDDoSL4                         VersionGetResponsePhase = "ddos_l4"
	VersionGetResponsePhaseDDoSL7                         VersionGetResponsePhase = "ddos_l7"
	VersionGetResponsePhaseHTTPConfigSettings             VersionGetResponsePhase = "http_config_settings"
	VersionGetResponsePhaseHTTPCustomErrors               VersionGetResponsePhase = "http_custom_errors"
	VersionGetResponsePhaseHTTPLogCustomFields            VersionGetResponsePhase = "http_log_custom_fields"
	VersionGetResponsePhaseHTTPRatelimit                  VersionGetResponsePhase = "http_ratelimit"
	VersionGetResponsePhaseHTTPRequestCacheSettings       VersionGetResponsePhase = "http_request_cache_settings"
	VersionGetResponsePhaseHTTPRequestDynamicRedirect     VersionGetResponsePhase = "http_request_dynamic_redirect"
	VersionGetResponsePhaseHTTPRequestFirewallCustom      VersionGetResponsePhase = "http_request_firewall_custom"
	VersionGetResponsePhaseHTTPRequestFirewallManaged     VersionGetResponsePhase = "http_request_firewall_managed"
	VersionGetResponsePhaseHTTPRequestLateTransform       VersionGetResponsePhase = "http_request_late_transform"
	VersionGetResponsePhaseHTTPRequestOrigin              VersionGetResponsePhase = "http_request_origin"
	VersionGetResponsePhaseHTTPRequestRedirect            VersionGetResponsePhase = "http_request_redirect"
	VersionGetResponsePhaseHTTPRequestSanitize            VersionGetResponsePhase = "http_request_sanitize"
	VersionGetResponsePhaseHTTPRequestSbfm                VersionGetResponsePhase = "http_request_sbfm"
	VersionGetResponsePhaseHTTPRequestSelectConfiguration VersionGetResponsePhase = "http_request_select_configuration"
	VersionGetResponsePhaseHTTPRequestTransform           VersionGetResponsePhase = "http_request_transform"
	VersionGetResponsePhaseHTTPResponseCompression        VersionGetResponsePhase = "http_response_compression"
	VersionGetResponsePhaseHTTPResponseFirewallManaged    VersionGetResponsePhase = "http_response_firewall_managed"
	VersionGetResponsePhaseHTTPResponseHeadersTransform   VersionGetResponsePhase = "http_response_headers_transform"
	VersionGetResponsePhaseMagicTransit                   VersionGetResponsePhase = "magic_transit"
	VersionGetResponsePhaseMagicTransitIDsManaged         VersionGetResponsePhase = "magic_transit_ids_managed"
	VersionGetResponsePhaseMagicTransitManaged            VersionGetResponsePhase = "magic_transit_managed"
)

func (r VersionGetResponsePhase) IsKnown() bool {
	switch r {
	case VersionGetResponsePhaseDDoSL4, VersionGetResponsePhaseDDoSL7, VersionGetResponsePhaseHTTPConfigSettings, VersionGetResponsePhaseHTTPCustomErrors, VersionGetResponsePhaseHTTPLogCustomFields, VersionGetResponsePhaseHTTPRatelimit, VersionGetResponsePhaseHTTPRequestCacheSettings, VersionGetResponsePhaseHTTPRequestDynamicRedirect, VersionGetResponsePhaseHTTPRequestFirewallCustom, VersionGetResponsePhaseHTTPRequestFirewallManaged, VersionGetResponsePhaseHTTPRequestLateTransform, VersionGetResponsePhaseHTTPRequestOrigin, VersionGetResponsePhaseHTTPRequestRedirect, VersionGetResponsePhaseHTTPRequestSanitize, VersionGetResponsePhaseHTTPRequestSbfm, VersionGetResponsePhaseHTTPRequestSelectConfiguration, VersionGetResponsePhaseHTTPRequestTransform, VersionGetResponsePhaseHTTPResponseCompression, VersionGetResponsePhaseHTTPResponseFirewallManaged, VersionGetResponsePhaseHTTPResponseHeadersTransform, VersionGetResponsePhaseMagicTransit, VersionGetResponsePhaseMagicTransitIDsManaged, VersionGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// Union satisfied by [rulesets.VersionGetResponseRulesRulesetsBlockRule],
// [rulesets.VersionGetResponseRulesRulesetsExecuteRule],
// [rulesets.VersionGetResponseRulesRulesetsLogRule] or
// [rulesets.VersionGetResponseRulesRulesetsSkipRule].
type VersionGetResponseRule interface {
	implementsRulesetsVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRule)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsSkipRule{}),
			DiscriminatorValue: "skip",
		},
	)
}

type VersionGetResponseRulesRulesetsBlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsBlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsBlockRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionGetResponseRulesRulesetsBlockRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                       `json:"ref"`
	JSON versionGetResponseRulesRulesetsBlockRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsBlockRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRulesetsBlockRule]
type versionGetResponseRulesRulesetsBlockRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsBlockRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsBlockRuleAction string

const (
	VersionGetResponseRulesRulesetsBlockRuleActionBlock VersionGetResponseRulesRulesetsBlockRuleAction = "block"
)

func (r VersionGetResponseRulesRulesetsBlockRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsBlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsBlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response VersionGetResponseRulesRulesetsBlockRuleActionParametersResponse `json:"response"`
	JSON     versionGetResponseRulesRulesetsBlockRuleActionParametersJSON     `json:"-"`
}

// versionGetResponseRulesRulesetsBlockRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsBlockRuleActionParameters]
type versionGetResponseRulesRulesetsBlockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsBlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsBlockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type VersionGetResponseRulesRulesetsBlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                                                `json:"status_code,required"`
	JSON       versionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON `json:"-"`
}

// versionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsBlockRuleActionParametersResponse]
type versionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsBlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsBlockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type VersionGetResponseRulesRulesetsBlockRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                `json:"enabled,required"`
	JSON    versionGetResponseRulesRulesetsBlockRuleLoggingJSON `json:"-"`
}

// versionGetResponseRulesRulesetsBlockRuleLoggingJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesRulesetsBlockRuleLogging]
type versionGetResponseRulesRulesetsBlockRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsBlockRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsBlockRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesRulesetsExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsExecuteRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionGetResponseRulesRulesetsExecuteRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                         `json:"ref"`
	JSON versionGetResponseRulesRulesetsExecuteRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsExecuteRule]
type versionGetResponseRulesRulesetsExecuteRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsExecuteRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsExecuteRuleAction string

const (
	VersionGetResponseRulesRulesetsExecuteRuleActionExecute VersionGetResponseRulesRulesetsExecuteRuleAction = "execute"
)

func (r VersionGetResponseRulesRulesetsExecuteRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData VersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      versionGetResponseRulesRulesetsExecuteRuleActionParametersJSON      `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsExecuteRuleActionParameters]
type versionGetResponseRulesRulesetsExecuteRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                                                    `json:"public_key,required"`
	JSON      versionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData]
type versionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON             `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides]
type versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory]
type versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule]
type versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium  VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow     VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff    VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelDefault, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelMedium, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelLow, VersionGetResponseRulesRulesetsExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type VersionGetResponseRulesRulesetsExecuteRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                                  `json:"enabled,required"`
	JSON    versionGetResponseRulesRulesetsExecuteRuleLoggingJSON `json:"-"`
}

// versionGetResponseRulesRulesetsExecuteRuleLoggingJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesRulesetsExecuteRuleLogging]
type versionGetResponseRulesRulesetsExecuteRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsExecuteRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsExecuteRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesRulesetsLogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsLogRuleAction `json:"action"`
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
	Logging VersionGetResponseRulesRulesetsLogRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                     `json:"ref"`
	JSON versionGetResponseRulesRulesetsLogRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsLogRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRulesetsLogRule]
type versionGetResponseRulesRulesetsLogRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsLogRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsLogRuleAction string

const (
	VersionGetResponseRulesRulesetsLogRuleActionLog VersionGetResponseRulesRulesetsLogRuleAction = "log"
)

func (r VersionGetResponseRulesRulesetsLogRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsLogRuleActionLog:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type VersionGetResponseRulesRulesetsLogRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                              `json:"enabled,required"`
	JSON    versionGetResponseRulesRulesetsLogRuleLoggingJSON `json:"-"`
}

// versionGetResponseRulesRulesetsLogRuleLoggingJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesRulesetsLogRuleLogging]
type versionGetResponseRulesRulesetsLogRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsLogRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsLogRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesRulesetsSkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsSkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesRulesetsSkipRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging VersionGetResponseRulesRulesetsSkipRuleLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                      `json:"ref"`
	JSON versionGetResponseRulesRulesetsSkipRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSkipRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRulesetsSkipRule]
type versionGetResponseRulesRulesetsSkipRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsSkipRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsSkipRuleAction string

const (
	VersionGetResponseRulesRulesetsSkipRuleActionSkip VersionGetResponseRulesRulesetsSkipRuleAction = "skip"
)

func (r VersionGetResponseRulesRulesetsSkipRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesRulesetsSkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset VersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                                                    `json:"rulesets"`
	JSON     versionGetResponseRulesRulesetsSkipRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSkipRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesRulesetsSkipRuleActionParameters]
type versionGetResponseRulesRulesetsSkipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSkipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase string

const (
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4                         VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l4"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7                         VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "ddos_l7"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings             VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_config_settings"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors               VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_custom_errors"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields            VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_log_custom_fields"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit                  VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_ratelimit"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings       VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_cache_settings"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom      VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_custom"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged     VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_firewall_managed"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform       VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_late_transform"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin              VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_origin"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect            VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_redirect"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize            VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sanitize"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm                VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_sbfm"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_select_configuration"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform           VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_request_transform"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression        VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_compression"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged    VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_firewall_managed"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform   VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "http_response_headers_transform"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit                   VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged         VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_ids_managed"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged            VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r VersionGetResponseRulesRulesetsSkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL4, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseDDoSL7, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPConfigSettings, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPCustomErrors, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPLogCustomFields, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRatelimit, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestCacheSettings, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallCustom, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestFirewallManaged, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestLateTransform, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestOrigin, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestRedirect, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSanitize, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSbfm, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPRequestTransform, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseCompression, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseFirewallManaged, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseHTTPResponseHeadersTransform, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransit, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitIDsManaged, VersionGetResponseRulesRulesetsSkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct string

const (
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic           VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "bic"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot           VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "hot"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit     VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "rateLimit"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "securityLevel"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock       VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "uaBlock"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF           VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "waf"
	VersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown  VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct = "zoneLockdown"
)

func (r VersionGetResponseRulesRulesetsSkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSkipRuleActionParametersProductBic, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductHot, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductRateLimit, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductSecurityLevel, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductUABlock, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductWAF, VersionGetResponseRulesRulesetsSkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type VersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset string

const (
	VersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent VersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset = "current"
)

func (r VersionGetResponseRulesRulesetsSkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsSkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

// An object configuring the rule's logging behavior.
type VersionGetResponseRulesRulesetsSkipRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                               `json:"enabled,required"`
	JSON    versionGetResponseRulesRulesetsSkipRuleLoggingJSON `json:"-"`
}

// versionGetResponseRulesRulesetsSkipRuleLoggingJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesRulesetsSkipRuleLogging]
type versionGetResponseRulesRulesetsSkipRuleLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRulesetsSkipRuleLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsSkipRuleLoggingJSON) RawJSON() string {
	return r.raw
}

type VersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type VersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type VersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result VersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionGetResponseEnvelopeJSON    `json:"-"`
}

// versionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelope]
type versionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeErrors]
type versionGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeErrorsSource]
type versionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [VersionGetResponseEnvelopeMessages]
type versionGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    versionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [VersionGetResponseEnvelopeMessagesSource]
type versionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionGetResponseEnvelopeSuccess bool

const (
	VersionGetResponseEnvelopeSuccessTrue VersionGetResponseEnvelopeSuccess = true
)

func (r VersionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
