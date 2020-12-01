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
	// MagicFirewallRulesetTypeRoot specifies a Ruleset used for static mitigations for Magic Transit
	MagicFirewallRulesetTypeRoot = "root"

	// MagicFirewallRulesetPhaseMagicTransit specifies a Ruleset phase
	MagicFirewallRulesetPhaseMagicTransit = "magic_transit"

	// MagicFirewallRuleActionAllow specifies an allow action
	MagicFirewallRuleActionAllow MagicFirewallRuleAction = "allow"

	// MagicFirewallRuleActionBlock specifies a block action
	MagicFirewallRuleActionBlock MagicFirewallRuleAction = "block"
)

// MagicFirewallRuleAction specifies the action for a Magic Firewall rule
type MagicFirewallRuleAction string

// MagicFirewallRuleset contains information about a Magic Firewall Ruleset
type MagicFirewallRuleset struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Kind        string              `json:"kind"`
	Version     string              `json:"version"`
	LastUpdated *time.Time          `json:"last_updated"`
	Phase       string              `json:"phase"`
	Rules       []MagicFirewallRule `json:"rules"`
}

// MagicFirewallRule contains information about a single Magic Firewall rule
type MagicFirewallRule struct {
	ID          string                  `json:"id"`
	Version     string                  `json:"version"`
	Action      MagicFirewallRuleAction `json:"action"`
	Expression  string                  `json:"expression"`
	Description string                  `json:"description"`
	LastUpdated *time.Time              `json:"last_updated"`
	Ref         string                  `json:"ref"`
	Enabled     bool                    `json:"enabled"`
}

// CreateMagicFirewallRulesetRequest contains data for a new Magic Firewall ruleset
type CreateMagicFirewallRulesetRequest struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Kind        string              `json:"kind"`
	Phase       string              `json:"phase"`
	Rules       []MagicFirewallRule `json:"rules"`
}

// UpdateMagicFirewallRulesetRequest contains data for a Magic Firewall ruleset update
type UpdateMagicFirewallRulesetRequest struct {
	Description string              `json:"description"`
	Rules       []MagicFirewallRule `json:"rules"`
}

// ListFirewallRulesetResponse contains a list of Magic Firewall rulesets
type ListFirewallRulesetResponse struct {
	Response
	Result []MagicFirewallRuleset `json:"result"`
}

// CreateMagicFirewallRulesetResponse contains response data when creating a new Magic Firewall ruleset
type CreateMagicFirewallRulesetResponse struct {
	Response
	Result MagicFirewallRuleset `json:"result"`
}

// UpdateMagicFirewallRulesetResponse contains response data when updating an existing Magic Firewall ruleset
type UpdateMagicFirewallRulesetResponse struct {
	Response
	Result MagicFirewallRuleset `json:"result"`
}

// ListMagicFirewallRulesets lists all Rulesets for a given account
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) ListMagicFirewallRulesets(ctx context.Context) ([]MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MagicFirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := ListFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicFirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// CreateMagicFirewallRuleset creates a Magic Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) CreateMagicFirewallRuleset(ctx context.Context, name string, description string, kind string, phase string, rules []MagicFirewallRule) (MagicFirewallRuleset, error) {
	// if kind is not provided, we will use the root type
	if kind == "" {
		kind = MagicFirewallRulesetTypeRoot
	}

	// if phase is not provided, we are using the magic_transit default phase
	if phase == "" {
		phase = MagicFirewallRulesetPhaseMagicTransit
	}

	uri := fmt.Sprintf("/accounts/%s/rulesets", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		CreateMagicFirewallRulesetRequest{Name: name, Description: description, Kind: kind, Phase: phase, Rules: rules})
	if err != nil {
		return MagicFirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := CreateMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicFirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// DeleteMagicFirewallRuleset creates a Magic Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-delete-ruleset
func (api *API) DeleteMagicFirewallRuleset(ctx context.Context, id string) error {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	// Magic Firewall API is not implementing the standard response blob but returns an empty response (204) in case
	// of a success. So we are checking for the response body size here
	if len(res) > 0 {
		return errors.Wrap(errors.New(string(res)), errMakeRequestError)
	}

	return nil
}

// UpdateMagicFirewallRuleset updates a Magic Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-update-ruleset
func (api *API) UpdateMagicFirewallRuleset(ctx context.Context, id string, description string, rules []MagicFirewallRule) (MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", api.AccountID, id)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri,
		UpdateMagicFirewallRulesetRequest{Description: description, Rules: rules})
	if err != nil {
		return MagicFirewallRuleset{}, errors.Wrap(err, errMakeRequestError)
	}

	result := CreateMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicFirewallRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
