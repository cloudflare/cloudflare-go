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

// PhaseService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPhaseService] method instead.
type PhaseService struct {
	Options  []option.RequestOption
	Versions *PhaseVersionService
}

// NewPhaseService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPhaseService(opts ...option.RequestOption) (r *PhaseService) {
	r = &PhaseService{}
	r.Options = opts
	r.Versions = NewPhaseVersionService(opts...)
	return
}

// Updates an account or zone entry point ruleset, creating a new version.
func (r *PhaseService) Update(ctx context.Context, rulesetPhase PhaseUpdateParamsRulesetPhase, params PhaseUpdateParams, opts ...option.RequestOption) (res *PhaseUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *PhaseService) Get(ctx context.Context, rulesetPhase PhaseGetParamsRulesetPhase, query PhaseGetParams, opts ...option.RequestOption) (res *PhaseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseUpdateResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind PhaseUpdateResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseUpdateResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []PhaseUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                  `json:"description"`
	JSON        phaseUpdateResponseJSON `json:"-"`
}

// phaseUpdateResponseJSON contains the JSON metadata for the struct
// [PhaseUpdateResponse]
type phaseUpdateResponseJSON struct {
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

func (r *PhaseUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseUpdateResponseKind string

const (
	PhaseUpdateResponseKindManaged PhaseUpdateResponseKind = "managed"
	PhaseUpdateResponseKindCustom  PhaseUpdateResponseKind = "custom"
	PhaseUpdateResponseKindRoot    PhaseUpdateResponseKind = "root"
	PhaseUpdateResponseKindZone    PhaseUpdateResponseKind = "zone"
)

func (r PhaseUpdateResponseKind) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseKindManaged, PhaseUpdateResponseKindCustom, PhaseUpdateResponseKindRoot, PhaseUpdateResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseUpdateResponsePhase string

const (
	PhaseUpdateResponsePhaseDDoSL4                         PhaseUpdateResponsePhase = "ddos_l4"
	PhaseUpdateResponsePhaseDDoSL7                         PhaseUpdateResponsePhase = "ddos_l7"
	PhaseUpdateResponsePhaseHTTPConfigSettings             PhaseUpdateResponsePhase = "http_config_settings"
	PhaseUpdateResponsePhaseHTTPCustomErrors               PhaseUpdateResponsePhase = "http_custom_errors"
	PhaseUpdateResponsePhaseHTTPLogCustomFields            PhaseUpdateResponsePhase = "http_log_custom_fields"
	PhaseUpdateResponsePhaseHTTPRatelimit                  PhaseUpdateResponsePhase = "http_ratelimit"
	PhaseUpdateResponsePhaseHTTPRequestCacheSettings       PhaseUpdateResponsePhase = "http_request_cache_settings"
	PhaseUpdateResponsePhaseHTTPRequestDynamicRedirect     PhaseUpdateResponsePhase = "http_request_dynamic_redirect"
	PhaseUpdateResponsePhaseHTTPRequestFirewallCustom      PhaseUpdateResponsePhase = "http_request_firewall_custom"
	PhaseUpdateResponsePhaseHTTPRequestFirewallManaged     PhaseUpdateResponsePhase = "http_request_firewall_managed"
	PhaseUpdateResponsePhaseHTTPRequestLateTransform       PhaseUpdateResponsePhase = "http_request_late_transform"
	PhaseUpdateResponsePhaseHTTPRequestOrigin              PhaseUpdateResponsePhase = "http_request_origin"
	PhaseUpdateResponsePhaseHTTPRequestRedirect            PhaseUpdateResponsePhase = "http_request_redirect"
	PhaseUpdateResponsePhaseHTTPRequestSanitize            PhaseUpdateResponsePhase = "http_request_sanitize"
	PhaseUpdateResponsePhaseHTTPRequestSbfm                PhaseUpdateResponsePhase = "http_request_sbfm"
	PhaseUpdateResponsePhaseHTTPRequestSelectConfiguration PhaseUpdateResponsePhase = "http_request_select_configuration"
	PhaseUpdateResponsePhaseHTTPRequestTransform           PhaseUpdateResponsePhase = "http_request_transform"
	PhaseUpdateResponsePhaseHTTPResponseCompression        PhaseUpdateResponsePhase = "http_response_compression"
	PhaseUpdateResponsePhaseHTTPResponseFirewallManaged    PhaseUpdateResponsePhase = "http_response_firewall_managed"
	PhaseUpdateResponsePhaseHTTPResponseHeadersTransform   PhaseUpdateResponsePhase = "http_response_headers_transform"
	PhaseUpdateResponsePhaseMagicTransit                   PhaseUpdateResponsePhase = "magic_transit"
	PhaseUpdateResponsePhaseMagicTransitIDsManaged         PhaseUpdateResponsePhase = "magic_transit_ids_managed"
	PhaseUpdateResponsePhaseMagicTransitManaged            PhaseUpdateResponsePhase = "magic_transit_managed"
)

func (r PhaseUpdateResponsePhase) IsKnown() bool {
	switch r {
	case PhaseUpdateResponsePhaseDDoSL4, PhaseUpdateResponsePhaseDDoSL7, PhaseUpdateResponsePhaseHTTPConfigSettings, PhaseUpdateResponsePhaseHTTPCustomErrors, PhaseUpdateResponsePhaseHTTPLogCustomFields, PhaseUpdateResponsePhaseHTTPRatelimit, PhaseUpdateResponsePhaseHTTPRequestCacheSettings, PhaseUpdateResponsePhaseHTTPRequestDynamicRedirect, PhaseUpdateResponsePhaseHTTPRequestFirewallCustom, PhaseUpdateResponsePhaseHTTPRequestFirewallManaged, PhaseUpdateResponsePhaseHTTPRequestLateTransform, PhaseUpdateResponsePhaseHTTPRequestOrigin, PhaseUpdateResponsePhaseHTTPRequestRedirect, PhaseUpdateResponsePhaseHTTPRequestSanitize, PhaseUpdateResponsePhaseHTTPRequestSbfm, PhaseUpdateResponsePhaseHTTPRequestSelectConfiguration, PhaseUpdateResponsePhaseHTTPRequestTransform, PhaseUpdateResponsePhaseHTTPResponseCompression, PhaseUpdateResponsePhaseHTTPResponseFirewallManaged, PhaseUpdateResponsePhaseHTTPResponseHeadersTransform, PhaseUpdateResponsePhaseMagicTransit, PhaseUpdateResponsePhaseMagicTransitIDsManaged, PhaseUpdateResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseUpdateResponseRule struct {
	// The action to perform when the rule matches.
	Action           PhaseUpdateResponseRulesAction `json:"action"`
	ActionParameters interface{}                    `json:"action_parameters,required"`
	Categories       interface{}                    `json:"categories,required"`
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
	Version string                      `json:"version,required"`
	JSON    phaseUpdateResponseRuleJSON `json:"-"`
	union   PhaseUpdateResponseRulesUnion
}

// phaseUpdateResponseRuleJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseRule]
type phaseUpdateResponseRuleJSON struct {
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

func (r phaseUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRule) AsUnion() PhaseUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ExecuteRule],
// [rulesets.LogRule] or [rulesets.SkipRule].
type PhaseUpdateResponseRulesUnion interface {
	implementsRulesetsPhaseUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesUnion)(nil)).Elem(),
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
type PhaseUpdateResponseRulesAction string

const (
	PhaseUpdateResponseRulesActionBlock   PhaseUpdateResponseRulesAction = "block"
	PhaseUpdateResponseRulesActionExecute PhaseUpdateResponseRulesAction = "execute"
	PhaseUpdateResponseRulesActionLog     PhaseUpdateResponseRulesAction = "log"
	PhaseUpdateResponseRulesActionSkip    PhaseUpdateResponseRulesAction = "skip"
)

func (r PhaseUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesActionBlock, PhaseUpdateResponseRulesActionExecute, PhaseUpdateResponseRulesActionLog, PhaseUpdateResponseRulesActionSkip:
		return true
	}
	return false
}

// A ruleset object.
type PhaseGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind PhaseGetResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase PhaseGetResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []PhaseGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        phaseGetResponseJSON `json:"-"`
}

// phaseGetResponseJSON contains the JSON metadata for the struct
// [PhaseGetResponse]
type phaseGetResponseJSON struct {
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

func (r *PhaseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type PhaseGetResponseKind string

const (
	PhaseGetResponseKindManaged PhaseGetResponseKind = "managed"
	PhaseGetResponseKindCustom  PhaseGetResponseKind = "custom"
	PhaseGetResponseKindRoot    PhaseGetResponseKind = "root"
	PhaseGetResponseKindZone    PhaseGetResponseKind = "zone"
)

func (r PhaseGetResponseKind) IsKnown() bool {
	switch r {
	case PhaseGetResponseKindManaged, PhaseGetResponseKindCustom, PhaseGetResponseKindRoot, PhaseGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseGetResponsePhase string

const (
	PhaseGetResponsePhaseDDoSL4                         PhaseGetResponsePhase = "ddos_l4"
	PhaseGetResponsePhaseDDoSL7                         PhaseGetResponsePhase = "ddos_l7"
	PhaseGetResponsePhaseHTTPConfigSettings             PhaseGetResponsePhase = "http_config_settings"
	PhaseGetResponsePhaseHTTPCustomErrors               PhaseGetResponsePhase = "http_custom_errors"
	PhaseGetResponsePhaseHTTPLogCustomFields            PhaseGetResponsePhase = "http_log_custom_fields"
	PhaseGetResponsePhaseHTTPRatelimit                  PhaseGetResponsePhase = "http_ratelimit"
	PhaseGetResponsePhaseHTTPRequestCacheSettings       PhaseGetResponsePhase = "http_request_cache_settings"
	PhaseGetResponsePhaseHTTPRequestDynamicRedirect     PhaseGetResponsePhase = "http_request_dynamic_redirect"
	PhaseGetResponsePhaseHTTPRequestFirewallCustom      PhaseGetResponsePhase = "http_request_firewall_custom"
	PhaseGetResponsePhaseHTTPRequestFirewallManaged     PhaseGetResponsePhase = "http_request_firewall_managed"
	PhaseGetResponsePhaseHTTPRequestLateTransform       PhaseGetResponsePhase = "http_request_late_transform"
	PhaseGetResponsePhaseHTTPRequestOrigin              PhaseGetResponsePhase = "http_request_origin"
	PhaseGetResponsePhaseHTTPRequestRedirect            PhaseGetResponsePhase = "http_request_redirect"
	PhaseGetResponsePhaseHTTPRequestSanitize            PhaseGetResponsePhase = "http_request_sanitize"
	PhaseGetResponsePhaseHTTPRequestSbfm                PhaseGetResponsePhase = "http_request_sbfm"
	PhaseGetResponsePhaseHTTPRequestSelectConfiguration PhaseGetResponsePhase = "http_request_select_configuration"
	PhaseGetResponsePhaseHTTPRequestTransform           PhaseGetResponsePhase = "http_request_transform"
	PhaseGetResponsePhaseHTTPResponseCompression        PhaseGetResponsePhase = "http_response_compression"
	PhaseGetResponsePhaseHTTPResponseFirewallManaged    PhaseGetResponsePhase = "http_response_firewall_managed"
	PhaseGetResponsePhaseHTTPResponseHeadersTransform   PhaseGetResponsePhase = "http_response_headers_transform"
	PhaseGetResponsePhaseMagicTransit                   PhaseGetResponsePhase = "magic_transit"
	PhaseGetResponsePhaseMagicTransitIDsManaged         PhaseGetResponsePhase = "magic_transit_ids_managed"
	PhaseGetResponsePhaseMagicTransitManaged            PhaseGetResponsePhase = "magic_transit_managed"
)

func (r PhaseGetResponsePhase) IsKnown() bool {
	switch r {
	case PhaseGetResponsePhaseDDoSL4, PhaseGetResponsePhaseDDoSL7, PhaseGetResponsePhaseHTTPConfigSettings, PhaseGetResponsePhaseHTTPCustomErrors, PhaseGetResponsePhaseHTTPLogCustomFields, PhaseGetResponsePhaseHTTPRatelimit, PhaseGetResponsePhaseHTTPRequestCacheSettings, PhaseGetResponsePhaseHTTPRequestDynamicRedirect, PhaseGetResponsePhaseHTTPRequestFirewallCustom, PhaseGetResponsePhaseHTTPRequestFirewallManaged, PhaseGetResponsePhaseHTTPRequestLateTransform, PhaseGetResponsePhaseHTTPRequestOrigin, PhaseGetResponsePhaseHTTPRequestRedirect, PhaseGetResponsePhaseHTTPRequestSanitize, PhaseGetResponsePhaseHTTPRequestSbfm, PhaseGetResponsePhaseHTTPRequestSelectConfiguration, PhaseGetResponsePhaseHTTPRequestTransform, PhaseGetResponsePhaseHTTPResponseCompression, PhaseGetResponsePhaseHTTPResponseFirewallManaged, PhaseGetResponsePhaseHTTPResponseHeadersTransform, PhaseGetResponsePhaseMagicTransit, PhaseGetResponsePhaseMagicTransitIDsManaged, PhaseGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           PhaseGetResponseRulesAction `json:"action"`
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
	JSON    phaseGetResponseRuleJSON `json:"-"`
	union   PhaseGetResponseRulesUnion
}

// phaseGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseGetResponseRule]
type phaseGetResponseRuleJSON struct {
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

func (r phaseGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ExecuteRule],
// [rulesets.LogRule] or [rulesets.SkipRule].
type PhaseGetResponseRulesUnion interface {
	implementsRulesetsPhaseGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesUnion)(nil)).Elem(),
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
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock   PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionExecute PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionLog     PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionSkip    PhaseGetResponseRulesAction = "skip"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionSkip:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRuleUnion] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[PhaseUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[PhaseUpdateParamsPhase] `json:"phase"`
}

func (r PhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase of the ruleset.
type PhaseUpdateParamsRulesetPhase string

const (
	PhaseUpdateParamsRulesetPhaseDDoSL4                         PhaseUpdateParamsRulesetPhase = "ddos_l4"
	PhaseUpdateParamsRulesetPhaseDDoSL7                         PhaseUpdateParamsRulesetPhase = "ddos_l7"
	PhaseUpdateParamsRulesetPhaseHTTPConfigSettings             PhaseUpdateParamsRulesetPhase = "http_config_settings"
	PhaseUpdateParamsRulesetPhaseHTTPCustomErrors               PhaseUpdateParamsRulesetPhase = "http_custom_errors"
	PhaseUpdateParamsRulesetPhaseHTTPLogCustomFields            PhaseUpdateParamsRulesetPhase = "http_log_custom_fields"
	PhaseUpdateParamsRulesetPhaseHTTPRatelimit                  PhaseUpdateParamsRulesetPhase = "http_ratelimit"
	PhaseUpdateParamsRulesetPhaseHTTPRequestCacheSettings       PhaseUpdateParamsRulesetPhase = "http_request_cache_settings"
	PhaseUpdateParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseUpdateParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseUpdateParamsRulesetPhase = "http_request_firewall_custom"
	PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseUpdateParamsRulesetPhase = "http_request_firewall_managed"
	PhaseUpdateParamsRulesetPhaseHTTPRequestLateTransform       PhaseUpdateParamsRulesetPhase = "http_request_late_transform"
	PhaseUpdateParamsRulesetPhaseHTTPRequestOrigin              PhaseUpdateParamsRulesetPhase = "http_request_origin"
	PhaseUpdateParamsRulesetPhaseHTTPRequestRedirect            PhaseUpdateParamsRulesetPhase = "http_request_redirect"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSanitize            PhaseUpdateParamsRulesetPhase = "http_request_sanitize"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSbfm                PhaseUpdateParamsRulesetPhase = "http_request_sbfm"
	PhaseUpdateParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseUpdateParamsRulesetPhase = "http_request_select_configuration"
	PhaseUpdateParamsRulesetPhaseHTTPRequestTransform           PhaseUpdateParamsRulesetPhase = "http_request_transform"
	PhaseUpdateParamsRulesetPhaseHTTPResponseCompression        PhaseUpdateParamsRulesetPhase = "http_response_compression"
	PhaseUpdateParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseUpdateParamsRulesetPhase = "http_response_firewall_managed"
	PhaseUpdateParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseUpdateParamsRulesetPhase = "http_response_headers_transform"
	PhaseUpdateParamsRulesetPhaseMagicTransit                   PhaseUpdateParamsRulesetPhase = "magic_transit"
	PhaseUpdateParamsRulesetPhaseMagicTransitIDsManaged         PhaseUpdateParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseUpdateParamsRulesetPhaseMagicTransitManaged            PhaseUpdateParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseUpdateParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesetPhaseDDoSL4, PhaseUpdateParamsRulesetPhaseDDoSL7, PhaseUpdateParamsRulesetPhaseHTTPConfigSettings, PhaseUpdateParamsRulesetPhaseHTTPCustomErrors, PhaseUpdateParamsRulesetPhaseHTTPLogCustomFields, PhaseUpdateParamsRulesetPhaseHTTPRatelimit, PhaseUpdateParamsRulesetPhaseHTTPRequestCacheSettings, PhaseUpdateParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseUpdateParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseUpdateParamsRulesetPhaseHTTPRequestLateTransform, PhaseUpdateParamsRulesetPhaseHTTPRequestOrigin, PhaseUpdateParamsRulesetPhaseHTTPRequestRedirect, PhaseUpdateParamsRulesetPhaseHTTPRequestSanitize, PhaseUpdateParamsRulesetPhaseHTTPRequestSbfm, PhaseUpdateParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseUpdateParamsRulesetPhaseHTTPRequestTransform, PhaseUpdateParamsRulesetPhaseHTTPResponseCompression, PhaseUpdateParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseUpdateParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseUpdateParamsRulesetPhaseMagicTransit, PhaseUpdateParamsRulesetPhaseMagicTransitIDsManaged, PhaseUpdateParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

type PhaseUpdateParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[PhaseUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                  `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                  `json:"categories,required"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[shared.UnnamedSchemaRef70f2c6ccd8a405358ac7ef8fc3d6751cParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.LogRuleParam], [rulesets.SkipRuleParam], [PhaseUpdateParamsRule].
type PhaseUpdateParamsRuleUnion interface {
	implementsRulesetsPhaseUpdateParamsRuleUnion()
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock   PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionExecute PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionLog     PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionSkip    PhaseUpdateParamsRulesAction = "skip"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionSkip:
		return true
	}
	return false
}

// The kind of the ruleset.
type PhaseUpdateParamsKind string

const (
	PhaseUpdateParamsKindManaged PhaseUpdateParamsKind = "managed"
	PhaseUpdateParamsKindCustom  PhaseUpdateParamsKind = "custom"
	PhaseUpdateParamsKindRoot    PhaseUpdateParamsKind = "root"
	PhaseUpdateParamsKindZone    PhaseUpdateParamsKind = "zone"
)

func (r PhaseUpdateParamsKind) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsKindManaged, PhaseUpdateParamsKindCustom, PhaseUpdateParamsKindRoot, PhaseUpdateParamsKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type PhaseUpdateParamsPhase string

const (
	PhaseUpdateParamsPhaseDDoSL4                         PhaseUpdateParamsPhase = "ddos_l4"
	PhaseUpdateParamsPhaseDDoSL7                         PhaseUpdateParamsPhase = "ddos_l7"
	PhaseUpdateParamsPhaseHTTPConfigSettings             PhaseUpdateParamsPhase = "http_config_settings"
	PhaseUpdateParamsPhaseHTTPCustomErrors               PhaseUpdateParamsPhase = "http_custom_errors"
	PhaseUpdateParamsPhaseHTTPLogCustomFields            PhaseUpdateParamsPhase = "http_log_custom_fields"
	PhaseUpdateParamsPhaseHTTPRatelimit                  PhaseUpdateParamsPhase = "http_ratelimit"
	PhaseUpdateParamsPhaseHTTPRequestCacheSettings       PhaseUpdateParamsPhase = "http_request_cache_settings"
	PhaseUpdateParamsPhaseHTTPRequestDynamicRedirect     PhaseUpdateParamsPhase = "http_request_dynamic_redirect"
	PhaseUpdateParamsPhaseHTTPRequestFirewallCustom      PhaseUpdateParamsPhase = "http_request_firewall_custom"
	PhaseUpdateParamsPhaseHTTPRequestFirewallManaged     PhaseUpdateParamsPhase = "http_request_firewall_managed"
	PhaseUpdateParamsPhaseHTTPRequestLateTransform       PhaseUpdateParamsPhase = "http_request_late_transform"
	PhaseUpdateParamsPhaseHTTPRequestOrigin              PhaseUpdateParamsPhase = "http_request_origin"
	PhaseUpdateParamsPhaseHTTPRequestRedirect            PhaseUpdateParamsPhase = "http_request_redirect"
	PhaseUpdateParamsPhaseHTTPRequestSanitize            PhaseUpdateParamsPhase = "http_request_sanitize"
	PhaseUpdateParamsPhaseHTTPRequestSbfm                PhaseUpdateParamsPhase = "http_request_sbfm"
	PhaseUpdateParamsPhaseHTTPRequestSelectConfiguration PhaseUpdateParamsPhase = "http_request_select_configuration"
	PhaseUpdateParamsPhaseHTTPRequestTransform           PhaseUpdateParamsPhase = "http_request_transform"
	PhaseUpdateParamsPhaseHTTPResponseCompression        PhaseUpdateParamsPhase = "http_response_compression"
	PhaseUpdateParamsPhaseHTTPResponseFirewallManaged    PhaseUpdateParamsPhase = "http_response_firewall_managed"
	PhaseUpdateParamsPhaseHTTPResponseHeadersTransform   PhaseUpdateParamsPhase = "http_response_headers_transform"
	PhaseUpdateParamsPhaseMagicTransit                   PhaseUpdateParamsPhase = "magic_transit"
	PhaseUpdateParamsPhaseMagicTransitIDsManaged         PhaseUpdateParamsPhase = "magic_transit_ids_managed"
	PhaseUpdateParamsPhaseMagicTransitManaged            PhaseUpdateParamsPhase = "magic_transit_managed"
)

func (r PhaseUpdateParamsPhase) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsPhaseDDoSL4, PhaseUpdateParamsPhaseDDoSL7, PhaseUpdateParamsPhaseHTTPConfigSettings, PhaseUpdateParamsPhaseHTTPCustomErrors, PhaseUpdateParamsPhaseHTTPLogCustomFields, PhaseUpdateParamsPhaseHTTPRatelimit, PhaseUpdateParamsPhaseHTTPRequestCacheSettings, PhaseUpdateParamsPhaseHTTPRequestDynamicRedirect, PhaseUpdateParamsPhaseHTTPRequestFirewallCustom, PhaseUpdateParamsPhaseHTTPRequestFirewallManaged, PhaseUpdateParamsPhaseHTTPRequestLateTransform, PhaseUpdateParamsPhaseHTTPRequestOrigin, PhaseUpdateParamsPhaseHTTPRequestRedirect, PhaseUpdateParamsPhaseHTTPRequestSanitize, PhaseUpdateParamsPhaseHTTPRequestSbfm, PhaseUpdateParamsPhaseHTTPRequestSelectConfiguration, PhaseUpdateParamsPhaseHTTPRequestTransform, PhaseUpdateParamsPhaseHTTPResponseCompression, PhaseUpdateParamsPhaseHTTPResponseFirewallManaged, PhaseUpdateParamsPhaseHTTPResponseHeadersTransform, PhaseUpdateParamsPhaseMagicTransit, PhaseUpdateParamsPhaseMagicTransitIDsManaged, PhaseUpdateParamsPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseUpdateResponseEnvelopeJSON    `json:"-"`
}

// phaseUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelope]
type phaseUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeErrors]
type phaseUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeErrorsSource]
type phaseUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessages]
type phaseUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer,required"`
	JSON    phaseUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeMessagesSource]
type phaseUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseUpdateResponseEnvelopeSuccess bool

const (
	PhaseUpdateResponseEnvelopeSuccessTrue PhaseUpdateResponseEnvelopeSuccess = true
)

func (r PhaseUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PhaseGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// The phase of the ruleset.
type PhaseGetParamsRulesetPhase string

const (
	PhaseGetParamsRulesetPhaseDDoSL4                         PhaseGetParamsRulesetPhase = "ddos_l4"
	PhaseGetParamsRulesetPhaseDDoSL7                         PhaseGetParamsRulesetPhase = "ddos_l7"
	PhaseGetParamsRulesetPhaseHTTPConfigSettings             PhaseGetParamsRulesetPhase = "http_config_settings"
	PhaseGetParamsRulesetPhaseHTTPCustomErrors               PhaseGetParamsRulesetPhase = "http_custom_errors"
	PhaseGetParamsRulesetPhaseHTTPLogCustomFields            PhaseGetParamsRulesetPhase = "http_log_custom_fields"
	PhaseGetParamsRulesetPhaseHTTPRatelimit                  PhaseGetParamsRulesetPhase = "http_ratelimit"
	PhaseGetParamsRulesetPhaseHTTPRequestCacheSettings       PhaseGetParamsRulesetPhase = "http_request_cache_settings"
	PhaseGetParamsRulesetPhaseHTTPRequestDynamicRedirect     PhaseGetParamsRulesetPhase = "http_request_dynamic_redirect"
	PhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom      PhaseGetParamsRulesetPhase = "http_request_firewall_custom"
	PhaseGetParamsRulesetPhaseHTTPRequestFirewallManaged     PhaseGetParamsRulesetPhase = "http_request_firewall_managed"
	PhaseGetParamsRulesetPhaseHTTPRequestLateTransform       PhaseGetParamsRulesetPhase = "http_request_late_transform"
	PhaseGetParamsRulesetPhaseHTTPRequestOrigin              PhaseGetParamsRulesetPhase = "http_request_origin"
	PhaseGetParamsRulesetPhaseHTTPRequestRedirect            PhaseGetParamsRulesetPhase = "http_request_redirect"
	PhaseGetParamsRulesetPhaseHTTPRequestSanitize            PhaseGetParamsRulesetPhase = "http_request_sanitize"
	PhaseGetParamsRulesetPhaseHTTPRequestSbfm                PhaseGetParamsRulesetPhase = "http_request_sbfm"
	PhaseGetParamsRulesetPhaseHTTPRequestSelectConfiguration PhaseGetParamsRulesetPhase = "http_request_select_configuration"
	PhaseGetParamsRulesetPhaseHTTPRequestTransform           PhaseGetParamsRulesetPhase = "http_request_transform"
	PhaseGetParamsRulesetPhaseHTTPResponseCompression        PhaseGetParamsRulesetPhase = "http_response_compression"
	PhaseGetParamsRulesetPhaseHTTPResponseFirewallManaged    PhaseGetParamsRulesetPhase = "http_response_firewall_managed"
	PhaseGetParamsRulesetPhaseHTTPResponseHeadersTransform   PhaseGetParamsRulesetPhase = "http_response_headers_transform"
	PhaseGetParamsRulesetPhaseMagicTransit                   PhaseGetParamsRulesetPhase = "magic_transit"
	PhaseGetParamsRulesetPhaseMagicTransitIDsManaged         PhaseGetParamsRulesetPhase = "magic_transit_ids_managed"
	PhaseGetParamsRulesetPhaseMagicTransitManaged            PhaseGetParamsRulesetPhase = "magic_transit_managed"
)

func (r PhaseGetParamsRulesetPhase) IsKnown() bool {
	switch r {
	case PhaseGetParamsRulesetPhaseDDoSL4, PhaseGetParamsRulesetPhaseDDoSL7, PhaseGetParamsRulesetPhaseHTTPConfigSettings, PhaseGetParamsRulesetPhaseHTTPCustomErrors, PhaseGetParamsRulesetPhaseHTTPLogCustomFields, PhaseGetParamsRulesetPhaseHTTPRatelimit, PhaseGetParamsRulesetPhaseHTTPRequestCacheSettings, PhaseGetParamsRulesetPhaseHTTPRequestDynamicRedirect, PhaseGetParamsRulesetPhaseHTTPRequestFirewallCustom, PhaseGetParamsRulesetPhaseHTTPRequestFirewallManaged, PhaseGetParamsRulesetPhaseHTTPRequestLateTransform, PhaseGetParamsRulesetPhaseHTTPRequestOrigin, PhaseGetParamsRulesetPhaseHTTPRequestRedirect, PhaseGetParamsRulesetPhaseHTTPRequestSanitize, PhaseGetParamsRulesetPhaseHTTPRequestSbfm, PhaseGetParamsRulesetPhaseHTTPRequestSelectConfiguration, PhaseGetParamsRulesetPhaseHTTPRequestTransform, PhaseGetParamsRulesetPhaseHTTPResponseCompression, PhaseGetParamsRulesetPhaseHTTPResponseFirewallManaged, PhaseGetParamsRulesetPhaseHTTPResponseHeadersTransform, PhaseGetParamsRulesetPhaseMagicTransit, PhaseGetParamsRulesetPhaseMagicTransitIDsManaged, PhaseGetParamsRulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A response object.
type PhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    phaseGetResponseEnvelopeJSON    `json:"-"`
}

// phaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelope]
type phaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeErrors]
type phaseGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeErrorsSource]
type phaseGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [PhaseGetResponseEnvelopeMessages]
type phaseGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    phaseGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [PhaseGetResponseEnvelopeMessagesSource]
type phaseGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseGetResponseEnvelopeSuccess bool

const (
	PhaseGetResponseEnvelopeSuccessTrue PhaseGetResponseEnvelopeSuccess = true
)

func (r PhaseGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
