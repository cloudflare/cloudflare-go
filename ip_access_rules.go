package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ListIpAccessRulesOrderOption string
type ListIpAccessRulesMatchOption string
type IPAccessRulesModeOption string

const (
	IpAccessRulesConfigurationTarget  ListIpAccessRulesOrderOption = "configuration.target"
	IpAccessRulesConfigurationValue   ListIpAccessRulesOrderOption = "configuration.value"
	IpAccessRulesMode                 ListIpAccessRulesOrderOption = "mode"
	MatchOptionAll                    ListIpAccessRulesMatchOption = "all"
	MatchOptionAny                    ListIpAccessRulesMatchOption = "any"
	IpAccessRulesModeBlock            IPAccessRulesModeOption      = "block"
	IpAccessRulesModeChallenge        IPAccessRulesModeOption      = "challenge"
	IpAccessRulesModeJsChallenge      IPAccessRulesModeOption      = "js_challenge"
	IpAccessRulesModeManagedChallenge IPAccessRulesModeOption      = "managed_challenge"
	IpAccessRulesModeWhitelist        IPAccessRulesModeOption      = "whitelist"
)

type ListIpAccessRulesFilters struct {
	Configuration IPAccessRuleConfiguration    `json:"configuration,omitempty"`
	Match         ListIpAccessRulesMatchOption `json:"match,omitempty"`
	Mode          IPAccessRulesModeOption      `json:"mode,omitempty"`
	Notes         string                       `json:"notes,omitempty"`
}

type ListIpAccessRulesParams struct {
	Direction string                       `url:"direction,omitempty"`
	Filters   ListIpAccessRulesFilters     `url:"filters,omitempty"`
	Order     ListIpAccessRulesOrderOption `url:"order,omitempty"`
	PaginationOptions
}

type IPAccessRuleConfiguration struct {
	Target string `json:"target,omitempty"`
	Value  string `json:"value,omitempty"`
}

type IpAccessRule struct {
	AllowedModes  []IPAccessRulesModeOption `json:"allowed_modes"`
	Configuration IPAccessRuleConfiguration `json:"configuration"`
	CreatedOn     string                    `json:"created_on"`
	ID            string                    `json:"id"`
	Mode          IPAccessRulesModeOption   `json:"mode"`
	ModifiedOn    string                    `json:"modified_on"`
	Notes         string                    `json:"notes"`
}

type ListIpAccessRulesResponse struct {
	Result     []IpAccessRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// ListIpAccessRules
//
// Fetches IP Access rules of a zone/user/account. You can filter the results using several optional parameters.
//
// API references:
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-a-user-list-ip-access-rules
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-a-zone-list-ip-access-rules
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-an-account-list-ip-access-rules
func (api *API) ListIpAccessRules(ctx context.Context, rc *ResourceContainer, params ListIpAccessRulesParams) ([]IpAccessRule, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []IpAccessRule{}, &ResultInfo{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/firewall/access_rules/rules", rc.Level, rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []IpAccessRule{}, &ResultInfo{}, err
	}

	result := ListIpAccessRulesResponse{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return []IpAccessRule{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, &result.ResultInfo, nil
}
