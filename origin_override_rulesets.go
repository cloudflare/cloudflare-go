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
	// OriginOverrideRulesetKindRoot specifies a root Ruleset.
	OriginOverrideRulesetKindRoot = "root"

	// OriginOverrideRulesetPhaseOrigin specifies the Origin Override Ruleset phase.
	OriginOverrideRulesetPhaseOrigin = "http_request_origin"

	// OriginOverrideRulesetRuleActionRoute specifies a route action.
	OriginOverrideRulesetRuleActionRoute OriginOverrideRulesetRuleAction = "route"
)

// OriginOverrideRulesetRuleAction specifies the action for a Origin Override rule.
type OriginOverrideRulesetRuleAction string

// OriginOverrideRuleset contains information about a Origin Override Ruleset.
type OriginOverrideRuleset struct {
	ID          string                      `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Kind        string                      `json:"kind"`
	Version     string                      `json:"version,omitempty"`
	LastUpdated *time.Time                  `json:"last_updated,omitempty"`
	Phase       string                      `json:"phase"`
	Rules       []OriginOverrideRulesetRule `json:"rules"`
}

type OriginOverrideRulesetRuleActionOriginParameters struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

// OriginOverrideRulesetRuleActionParameters specifies the action parameters for a Origin Override rule.
type OriginOverrideRulesetRuleActionParameters struct {
	HostHeader string                                          `json:"host_header"`
	Origin     OriginOverrideRulesetRuleActionOriginParameters `json:"origin"`
}

// OriginOverrideRulesetRule contains information about a single Origin Override rule.
type OriginOverrideRulesetRule struct {
	ID               string                                     `json:"id,omitempty"`
	Version          string                                     `json:"version,omitempty"`
	Action           OriginOverrideRulesetRuleAction            `json:"action"`
	ActionParameters *OriginOverrideRulesetRuleActionParameters `json:"action_parameters,omitempty"`
	Expression       string                                     `json:"expression"`
	Description      string                                     `json:"description"`
	LastUpdated      *time.Time                                 `json:"last_updated,omitempty"`
	Ref              string                                     `json:"ref,omitempty"`
	Enabled          bool                                       `json:"enabled"`
}

// CreateOriginOverrideRulesetRequest contains data for a new Origin Override ruleset.
type CreateOriginOverrideRulesetRequest struct {
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Kind        string                      `json:"kind"`
	Phase       string                      `json:"phase"`
	Rules       []OriginOverrideRulesetRule `json:"rules"`
}

// UpdateOriginOverrideRulesetRequest contains data for a Origin Override ruleset update.
type UpdateOriginOverrideRulesetRequest struct {
	Description string                      `json:"description"`
	Rules       []OriginOverrideRulesetRule `json:"rules"`
}

// ListOriginOverrideRulesetResponse contains a list of Origin Override rulesets.
type ListOriginOverrideRulesetResponse struct {
	Response
	Result []OriginOverrideRuleset `json:"result"`
}

// GetOriginOverrideRulesetResponse contains a single Origin Override Ruleset.
type GetOriginOverrideRulesetResponse struct {
	Response
	Result OriginOverrideRuleset `json:"result"`
}

// CreateOriginOverrideRulesetResponse contains response data when creating a new Origin Override ruleset.
type CreateOriginOverrideRulesetResponse struct {
	Response
	Result OriginOverrideRuleset `json:"result"`
}

// UpdateOriginOverrideRulesetResponse contains response data when updating an existing Origin Override ruleset.
type UpdateOriginOverrideRulesetResponse struct {
	Response
	Result OriginOverrideRuleset `json:"result"`
}

// ListOriginOverrideRulesets lists all Rulesets for a given account
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) ListOriginOverrideRulesets(ctx context.Context, accountID string) ([]OriginOverrideRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets?phase=%s", accountID, OriginOverrideRulesetPhaseOrigin)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []OriginOverrideRuleset{}, err
	}

	result := ListOriginOverrideRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []OriginOverrideRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// GetOriginOverrideRuleset returns a specific Origin Override Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-get-a-ruleset
func (api *API) GetOriginOverrideRuleset(ctx context.Context, accountID, ID string) (OriginOverrideRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return OriginOverrideRuleset{}, err
	}

	result := GetOriginOverrideRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return OriginOverrideRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// CreateOriginOverrideRuleset creates a Origin Override ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
func (api *API) CreateOriginOverrideRuleset(ctx context.Context, accountID, name, description string, rules []OriginOverrideRulesetRule) (OriginOverrideRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		CreateOriginOverrideRulesetRequest{
			Name:        name,
			Description: description,
			Kind:        OriginOverrideRulesetKindRoot,
			Phase:       OriginOverrideRulesetPhaseOrigin,
			Rules:       rules})
	if err != nil {
		return OriginOverrideRuleset{}, err
	}

	result := CreateOriginOverrideRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return OriginOverrideRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// DeleteOriginOverrideRuleset deletes a Origin Override ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-delete-ruleset
func (api *API) DeleteOriginOverrideRuleset(ctx context.Context, accountID, ID string) error {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	// Firewall API is not implementing the standard response blob but returns an empty response (204) in case
	// of a success. So we are checking for the response body size here
	if len(res) > 0 {
		return errors.Wrap(errors.New(string(res)), errMakeRequestError)
	}

	return nil
}

// UpdateOriginOverrideRuleset updates a Origin Override ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-update-ruleset
func (api *API) UpdateOriginOverrideRuleset(ctx context.Context, accountID, ID string, description string, rules []OriginOverrideRulesetRule) (OriginOverrideRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri,
		UpdateOriginOverrideRulesetRequest{Description: description, Rules: rules})
	if err != nil {
		return OriginOverrideRuleset{}, err
	}

	result := UpdateOriginOverrideRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return OriginOverrideRuleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
