package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type TeamsRuleSettings struct {
	// list of ipv4 or ipv6 ips to override with, when action is set to dns override
	OverrideIPs []string `json:"override_ips"`

	// show this string at block page caused by this rule
	BlockReason string `json:"block_reason"`

	// host name to override with when action is set to dns override. Can not be used with OverrideIPs
	OverrideHost string `json:"override_host"`

	// settings for browser isolation actions
	BISOAdminControls *TeamsBISOAdminControlSettings `json:"biso_admin_controls"`

	// settings for l4(network) level overrides
	L4Override *TeamsL4OverrideSettings `json:"l4override"`

	// settings for adding headers to http requests
	AddHeaders http.Header `json:"add_headers"`

	// settings for session check in allow action
	CheckSession *TeamsCheckSessionSettings `json:"check_session"`

	// Enable block page on rules with action block
	BlockPageEnabled bool `json:"block_page_enabled"`

	// whether to disable dnssec validation for allow action
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
}

// TeamsL4OverrideSettings used in l4 filter type rule with action set to override.
type TeamsL4OverrideSettings struct {
	IP   string `json:"ip,omitempty"`
	Port int    `json:"port,omitempty"`
}

type TeamsBISOAdminControlSettings struct {
	DisablePrinting  bool `json:"dp"`
	DisableCopyPaste bool `json:"dcp"`
	DisableDownload  bool `json:"dd"`
	DisableUpload    bool `json:"du"`
	DisableKeyboard  bool `json:"dk"`
}

type TeamsCheckSessionSettings struct {
	Enforce  bool     `json:"enforce"`
	Duration Duration `json:"duration"`
}

type TeamsFilterType string

type TeamsGatewayAction string

const (
	HTTPFilter TeamsFilterType = "http"
	DNSFilter  TeamsFilterType = "dns"
	L4Filter   TeamsFilterType = "l4"
)

const (
	Allow        TeamsGatewayAction = "allow"
	Block        TeamsGatewayAction = "block"
	SafeSearch   TeamsGatewayAction = "safesearch"
	YTRestricted TeamsGatewayAction = "ytrestricted"
	On           TeamsGatewayAction = "on"
	Off          TeamsGatewayAction = "off"
	Scan         TeamsGatewayAction = "scan"
	NoScan       TeamsGatewayAction = "noscan"
	Isolate      TeamsGatewayAction = "isolate"
	NoIsolate    TeamsGatewayAction = "noisolate"
	Override     TeamsGatewayAction = "override"
	L4Override   TeamsGatewayAction = "l4_override"
)

func TeamsRulesActionValues() []string {
	return []string{
		string(Allow),
		string(Block),
		string(SafeSearch),
		string(YTRestricted),
		string(On),
		string(Off),
		string(Scan),
		string(NoScan),
		string(Isolate),
		string(NoIsolate),
		string(Override),
		string(L4Override),
	}
}

// TeamsRule represents an Teams wirefilter rule.
type TeamsRule struct {
	ID            string             `json:"id,omitempty"`
	CreatedAt     *time.Time         `json:"created_at,omitempty"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty"`
	DeletedAt     *time.Time         `json:"deleted_at,omitempty"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Precedence    uint64             `json:"precedence"`
	Enabled       bool               `json:"enabled"`
	Action        TeamsGatewayAction `json:"action"`
	Filters       []TeamsFilterType  `json:"filters"`
	Traffic       string             `json:"traffic"`
	Identity      string             `json:"identity"`
	DevicePosture string             `json:"device_posture"`
	Version       uint64             `json:"version"`
	RuleSettings  TeamsRuleSettings  `json:"rule_settings,omitempty"`
}

// TeamsRuleResponse is the API response, containing a single rule.
type TeamsRuleResponse struct {
	Response
	Result TeamsRule `json:"result"`
}

// TeamsRulesResponse is the API response, containing an array of rules.
type TeamsRulesResponse struct {
	Response
	Result []TeamsRule `json:"result"`
}

// TeamsRulePatchRequest is used to patch an existing rule.
type TeamsRulePatchRequest struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Precedence   uint64             `json:"precedence"`
	Enabled      bool               `json:"enabled"`
	Action       TeamsGatewayAction `json:"action"`
	RuleSettings TeamsRuleSettings  `json:"rule_settings,omitempty"`
}

// TeamsRules returns all rules within an account.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsRules(ctx context.Context, accountID string) ([]TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsRule{}, err
	}

	var teamsRulesResponse TeamsRulesResponse
	err = json.Unmarshal(res, &teamsRulesResponse)
	if err != nil {
		return []TeamsRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return teamsRulesResponse.Result, nil
}

// TeamsRule returns the rule with rule ID in the URL.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsRule(ctx context.Context, accountID string, ruleID string) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsCreateRule creates a rule with wirefilter expression.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsCreateRule(ctx context.Context, accountID string, rule TeamsRule) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsUpdateRule updates a rule with wirefilter expression.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsUpdateRule(ctx context.Context, accountID string, ruleID string, rule TeamsRule) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsPatchRule patches a rule associated values.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsPatchRule(ctx context.Context, accountID string, ruleID string, rule TeamsRulePatchRequest) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsDeleteRule deletes a rule.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsDeleteRule(ctx context.Context, accountID string, ruleID string) error {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
