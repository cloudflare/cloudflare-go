// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetService] method
// instead.
type ZoneRulesetService struct {
	Options  []option.RequestOption
	Phases   *ZoneRulesetPhaseService
	Rules    *ZoneRulesetRuleService
	Versions *ZoneRulesetVersionService
}

// NewZoneRulesetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetService(opts ...option.RequestOption) (r *ZoneRulesetService) {
	r = &ZoneRulesetService{}
	r.Options = opts
	r.Phases = NewZoneRulesetPhaseService(opts...)
	r.Rules = NewZoneRulesetRuleService(opts...)
	r.Versions = NewZoneRulesetVersionService(opts...)
	return
}

// Fetches the latest version of a zone ruleset.
func (r *ZoneRulesetService) Get(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a zone ruleset, creating a new version.
func (r *ZoneRulesetService) Update(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetUpdateParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes all versions of an existing zone ruleset.
func (r *ZoneRulesetService) Delete(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a ruleset at the zone level.
func (r *ZoneRulesetService) ZoneRulesetsNewAZoneRuleset(ctx context.Context, zoneID string, body ZoneRulesetZoneRulesetsNewAZoneRulesetParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the zone level.
func (r *ZoneRulesetService) ZoneRulesetsListZoneRulesets(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneRulesetUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetUpdateParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by [ZoneRulesetUpdateParamsRulesCreateUpdateRule],
// [ZoneRulesetUpdateParamsRulesObject].
type ZoneRulesetUpdateParamsRule interface {
	implementsZoneRulesetUpdateParamsRule()
}

// A rule object.
type ZoneRulesetUpdateParamsRulesCreateUpdateRule struct {
	// The action to perform when the rule matches.
	Action param.Field[string] `json:"action,required"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression,required"`
	// The parameters configuring the rule action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesCreateUpdateRule) implementsZoneRulesetUpdateParamsRule() {}

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetUpdateParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetZoneRulesetsNewAZoneRulesetParams struct {
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[string] `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind string

const (
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindCustom ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "custom"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindRoot   ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "root"
	ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKindZone   ZoneRulesetZoneRulesetsNewAZoneRulesetParamsKind = "zone"
)

// A rule object.
//
// Satisfied by
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule],
// [ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesObject].
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule interface {
	implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule()
}

// A rule object.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule struct {
	// The action to perform when the rule matches.
	Action param.Field[string] `json:"action,required"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression,required"`
	// The parameters configuring the rule action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRule) implementsZoneRulesetZoneRulesetsNewAZoneRulesetParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetZoneRulesetsNewAZoneRulesetParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
