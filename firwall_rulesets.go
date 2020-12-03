package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	// FirewallRulesetKindRoot specifies a root Ruleset
	FirewallRulesetKindRoot = "root"

	// FirewallRulesetPhaseMagicTransit specifies the Magic Transit Ruleset phase
	FirewallRulesetPhaseMagicTransit = "magic_transit"

	// FirewallRulesetRuleActionAllow specifies an allow action
	FirewallRulesetRuleActionAllow FirewallRuleAction = "allow"

	// FirewallRulesetRuleActionBlock specifies a block action
	FirewallRulesetRuleActionBlock FirewallRuleAction = "block"
)

// FirewallRuleAction specifies the action for a Firewall rule
type FirewallRuleAction string

// FirewallRuleset contains information about a Firewall Ruleset
type FirewallRuleset struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Kind        string                `json:"kind"`
	Version     string                `json:"version"`
	LastUpdated *time.Time            `json:"last_updated"`
	Phase       string                `json:"phase"`
	Rules       []FirewallRulesetRule `json:"rules"`
}

// FirewallRulesetRule contains information about a single Firewall rule
type FirewallRulesetRule struct {
	ID          string             `json:"id"`
	Version     string             `json:"version"`
	Action      FirewallRuleAction `json:"action"`
	Expression  string             `json:"expression"`
	Description string             `json:"description"`
	LastUpdated *time.Time         `json:"last_updated"`
	Ref         string             `json:"ref"`
	Enabled     bool               `json:"enabled"`
}

// CreateFirewallRulesetRequest contains data for a new Firewall ruleset
type CreateFirewallRulesetRequest struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Kind        string                `json:"kind"`
	Phase       string                `json:"phase"`
	Rules       []FirewallRulesetRule `json:"rules"`
}

// UpdateFirewallRulesetRequest contains data for a Firewall ruleset update
type UpdateFirewallRulesetRequest struct {
	Description string                `json:"description"`
	Rules       []FirewallRulesetRule `json:"rules"`
}

// ListFirewallRulesetResponse contains a list of Firewall rulesets
type ListFirewallRulesetResponse struct {
	Response
	Result []FirewallRuleset `json:"result"`
}

// CreateFirewallRulesetResponse contains response data when creating a new Firewall ruleset
type CreateFirewallRulesetResponse struct {
	Response
	Result FirewallRuleset `json:"result"`
}

// UpdateFirewallRulesetResponse contains response data when updating an existing Firewall ruleset
type UpdateFirewallRulesetResponse struct {
	Response
	Result FirewallRuleset `json:"result"`
}

// ListFirewallRulesets lists all Rulesets for a given account
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) ListFirewallRulesets(ctx context.Context) ([]FirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []FirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []FirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// CreateFirewallRuleset creates a Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) CreateFirewallRuleset(ctx context.Context, name string, description string, kind string, phase string, rules []FirewallRulesetRule) (FirewallRuleset, error) {
	// if kind is not provided, we will use the root type
	if kind == "" {
		kind = FirewallRulesetKindRoot
	}

	// if phase is not provided, we are using the magic_transit default phase
	if phase == "" {
		phase = FirewallRulesetPhaseMagicTransit
	}

	uri := fmt.Sprintf("/accounts/%s/rulesets", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		CreateFirewallRulesetRequest{Name: name, Description: description, Kind: kind, Phase: phase, Rules: rules})
	if err != nil {
		return FirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := CreateFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return FirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// DeleteFirewallRuleset deletes a Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-delete-ruleset
func (api *API) DeleteFirewallRuleset(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	// Firewall API is not implementing the standard response blob but returns an empty response (204) in case
	// of a success. So we are checking for the response body size here
	if len(res) > 0 {
		return errors.Wrap(errors.New(string(res)), errMakeRequestError)
	}

	return nil
}

// UpdateFirewallRuleset updates a Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-update-ruleset
func (api *API) UpdateFirewallRuleset(ctx context.Context, id string, description string, rules []FirewallRulesetRule) (FirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri,
		UpdateFirewallRulesetRequest{Description: description, Rules: rules})
	if err != nil {
		return FirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := UpdateFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return FirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
