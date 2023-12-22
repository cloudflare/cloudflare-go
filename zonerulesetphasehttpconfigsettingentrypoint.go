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

// ZoneRulesetPhaseHTTPConfigSettingEntrypointService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseHTTPConfigSettingEntrypointService] method instead.
type ZoneRulesetPhaseHTTPConfigSettingEntrypointService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseHTTPConfigSettingEntrypointService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneRulesetPhaseHTTPConfigSettingEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseHTTPConfigSettingEntrypointService) {
	r = &ZoneRulesetPhaseHTTPConfigSettingEntrypointService{}
	r.Options = opts
	return
}

// Fetches all Config Rules in a zone.
func (r *ZoneRulesetPhaseHTTPConfigSettingEntrypointService) ConfigRulesListConfigRules(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ConfigRulesRuleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_config_settings/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Config Rules of a zone.
func (r *ZoneRulesetPhaseHTTPConfigSettingEntrypointService) ConfigRulesUpdateConfigRules(ctx context.Context, zoneID string, body ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_config_settings/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConfigRulesRuleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []ConfigRulesRulesetRule `json:"rules"`
	JSON  configRulesRulesetJSON   `json:"-"`
}

// configRulesRulesetJSON contains the JSON metadata for the struct
// [ConfigRulesRuleset]
type configRulesRulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigRulesRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigRulesRulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The parameters configuring the action.
	ActionParameters ConfigRulesRulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                             `json:"description"`
	Expression       interface{}                             `json:"expression"`
	Version          interface{}                             `json:"version"`
	JSON             configRulesRulesetRuleJSON              `json:"-"`
}

// configRulesRulesetRuleJSON contains the JSON metadata for the struct
// [ConfigRulesRulesetRule]
type configRulesRulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ConfigRulesRulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters configuring the action.
type ConfigRulesRulesetRulesActionParameters struct {
	// Enable or disable Automatic HTTPS Rewrites for matching requests
	AutomaticHTTPsRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify ConfigRulesRulesetRulesActionParametersAutominify `json:"autominify"`
	// Enable or disable Browser Integrity Check
	Bic bool `json:"bic"`
	// Disable all active Cloudflare Apps
	DisableApps bool `json:"disable_apps"`
	// Disable Cloudflare Railgun
	DisableRailgun bool `json:"disable_railgun"`
	// Disable Cloudflare Railgun
	DisableZaraz bool `json:"disable_zaraz"`
	// Enable or disable Email Obfuscation
	EmailObfuscation bool `json:"email_obfuscation"`
	// Enable or disable Hotlink Protection
	HotlinkProtection bool `json:"hotlink_protection"`
	// Enable or disable Mirage
	Mirage bool `json:"mirage"`
	// Enable or disableOpportunistic Encryption
	OpportunisticEncryption bool `json:"opportunistic_encryption"`
	// Set Polish compression options
	Polish string `json:"polish"`
	// Enable or disable Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Set the Security Level
	SecurityLevel string `json:"security_level"`
	// Enable or disable Server Side Excludes
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Select the SSL encryption mode
	Ssl string `json:"ssl"`
	// Enable or disable Signed Exchangesn(SXG)
	Sxg  bool                                        `json:"sxg"`
	JSON configRulesRulesetRulesActionParametersJSON `json:"-"`
}

// configRulesRulesetRulesActionParametersJSON contains the JSON metadata for the
// struct [ConfigRulesRulesetRulesActionParameters]
type configRulesRulesetRulesActionParametersJSON struct {
	AutomaticHTTPsRewrites  apijson.Field
	Autominify              apijson.Field
	Bic                     apijson.Field
	DisableApps             apijson.Field
	DisableRailgun          apijson.Field
	DisableZaraz            apijson.Field
	EmailObfuscation        apijson.Field
	HotlinkProtection       apijson.Field
	Mirage                  apijson.Field
	OpportunisticEncryption apijson.Field
	Polish                  apijson.Field
	RocketLoader            apijson.Field
	SecurityLevel           apijson.Field
	ServerSideExcludes      apijson.Field
	Ssl                     apijson.Field
	Sxg                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ConfigRulesRulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Select which file extensions to minify automatically.
type ConfigRulesRulesetRulesActionParametersAutominify struct {
	Css  bool                                                  `json:"css"`
	HTML bool                                                  `json:"html"`
	Js   bool                                                  `json:"js"`
	JSON configRulesRulesetRulesActionParametersAutominifyJSON `json:"-"`
}

// configRulesRulesetRulesActionParametersAutominifyJSON contains the JSON metadata
// for the struct [ConfigRulesRulesetRulesActionParametersAutominify]
type configRulesRulesetRulesActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigRulesRulesetRulesActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesObject].
type ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRule interface {
	implementsZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseHTTPConfigSettingEntrypointConfigRulesUpdateConfigRulesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
