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
	// RulesetKindRoot is definition for an account level ruleset.
	RulesetKindRoot RulesetKind = "root"

	// RulesetKindCustom is the user defined rulesets.
	RulesetKindCustom RulesetKind = "custom"

	// RulesetKindManaged denotes Cloudflare managed rulesets.
	RulesetKindManaged RulesetKind = "managed"

	// RulesetKindSchema denotes a schema ruleset.
	RulesetKindSchema RulesetKind = "schema"

	// RulesetKindZone expresses a zone level ruleset.
	RulesetKindZone RulesetKind = "zone"

	// RulesetPhaseDDoSL7 phase runs during DDoS mitigation stage.
	RulesetPhaseDDoSL7 RulesetPhase = "ddos_l7"

	// RulesetPhaseMagicTransit phase is invoked when traffic is routed via Magic
	// Transit.
	RulesetPhaseMagicTransit RulesetPhase = "magic_transit"

	// RulesetPhaseHTTPRequestMain runs in the primary part of the HTTP request.
	RulesetPhaseHTTPRequestMain RulesetPhase = "http_request_main"

	// RulesetPhaseHTTPRequestFirewallCustom runs on custom firewall rulesets.
	RulesetPhaseHTTPRequestFirewallCustom RulesetPhase = "http_request_firewall_custom"

	// RulesetPhaseHTTPRequestFirewallManaged runs for Cloudflare managed rulesets.
	RulesetPhaseHTTPRequestFirewallManaged RulesetPhase = "http_request_firewall_managed"

	// RulesetPhaseHTTPRequestTransform is performed at the HTTP request
	// transformation phase.
	RulesetPhaseHTTPRequestTransform RulesetPhase = "http_request_transform"

	// RulesetPhaseHTTPRequestSanitize is run during the HTTP request sanitisation
	// phase.
	RulesetPhaseHTTPRequestSanitize RulesetPhase = "http_request_sanitize"

	// RulesetRuleActionSkip represents the "skip" action.
	RulesetRuleActionSkip RulesetRuleAction = "skip"

	// RulesetRuleActionBlock represents the "block" action.
	RulesetRuleActionBlock RulesetRuleAction = "block"

	// RulesetRuleActionJSChallenge represents the "js_challenge" action.
	RulesetRuleActionJSChallenge RulesetRuleAction = "js_challenge"

	// RulesetRuleActionChallenge represents the "challenge" action.
	RulesetRuleActionChallenge RulesetRuleAction = "challenge"

	// RulesetRuleActionLog represents the "log" action.
	RulesetRuleActionLog RulesetRuleAction = "log"
)

// RulesetRuleAction defines a custom type that is used to express allowed
// values for the rule action.
type RulesetRuleAction string

// RulesetKind is the custom type for allowed variances of rulesets.
type RulesetKind string

// RulesetPhase is the custom type for defining at what point the ruleset will
// be applied and limited to expected values.
type RulesetPhase string

// Ruleset contains the structure of a Ruleset.
type Ruleset struct {
	ID          string        `json:"id,omitempty"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Kind        string        `json:"kind"`
	Version     string        `json:"version,omitempty"`
	LastUpdated *time.Time    `json:"last_updated,omitempty"`
	Phase       RulesetPhase  `json:"phase"`
	Rules       []RulesetRule `json:"rules"`
}

// RulesetRuleActionParameters specifies the action parameters for a Ruleset
// rule.
type RulesetRuleActionParameters struct {
	Ruleset string `json:"ruleset,omitempty"`
}

// RulesetRule contains information about a single Ruleset Rule.
type RulesetRule struct {
	ID               string                       `json:"id,omitempty"`
	Version          string                       `json:"version,omitempty"`
	Action           RulesetRuleAction            `json:"action"`
	ActionParameters *RulesetRuleActionParameters `json:"action_parameters,omitempty"`
	Expression       string                       `json:"expression"`
	Description      string                       `json:"description"`
	LastUpdated      *time.Time                   `json:"last_updated,omitempty"`
	Ref              string                       `json:"ref,omitempty"`
	Enabled          bool                         `json:"enabled"`
}

// UpdateRulesetRequest is the representation of a Ruleset update.
type UpdateRulesetRequest struct {
	Description string        `json:"description"`
	Rules       []RulesetRule `json:"rules"`
}

// ListRulesetResponse contains all Rulesets.
type ListRulesetResponse struct {
	Response
	Result []Ruleset `json:"result"`
}

// GetRulesetResponse contains a single Ruleset.
type GetRulesetResponse struct {
	Response
	Result Ruleset `json:"result"`
}

// CreateRulesetResponse contains response data when creating a new Ruleset.
type CreateRulesetResponse struct {
	Response
	Result Ruleset `json:"result"`
}

// UpdateRulesetResponse contains response data when updating an existing
// Ruleset.
type UpdateRulesetResponse struct {
	Response
	Result Ruleset `json:"result"`
}

// ListZoneRulesets fetches all rulesets for a zone.
//
// API reference: https://api.cloudflare.com/#zone-rulesets-list-zone-rulesets
func (api *API) ListZoneRulesets(ctx context.Context, zoneID string) ([]Ruleset, error) {
	return api.listRulesets(ctx, ZoneRouteRoot, zoneID)
}

// ListAccountRulesets fetches all rulesets for an account.
//
// API reference: https://api.cloudflare.com/#account-rulesets-list-account-rulesets
func (api *API) ListAccountRulesets(ctx context.Context, accountID string) ([]Ruleset, error) {
	return api.listRulesets(ctx, AccountRouteRoot, accountID)
}

// listRulesets lists all Rulesets for a given zone or account depending on the
// identifier type provided.
func (api *API) listRulesets(ctx context.Context, identifierType RouteRoot, identifier string) ([]Ruleset, error) {
	uri := fmt.Sprintf("/%s/%s/rulesets", identifierType, identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Ruleset{}, err
	}

	result := ListRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []Ruleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// GetZoneRuleset fetches a single ruleset for a zone.
//
// API reference: https://api.cloudflare.com/#zone-rulesets-get-a-zone-ruleset
func (api *API) GetZoneRuleset(ctx context.Context, zoneID, rulesetID string) (Ruleset, error) {
	return api.getRuleset(ctx, ZoneRouteRoot, zoneID, rulesetID)
}

// GetAccountRuleset fetches a single ruleset for an account.
//
// API reference: https://api.cloudflare.com/#account-rulesets-get-an-account-ruleset
func (api *API) GetAccountRuleset(ctx context.Context, accountID, rulesetID string) (Ruleset, error) {
	return api.getRuleset(ctx, AccountRouteRoot, accountID, rulesetID)
}

// getRuleset fetches a single ruleset based on the zone or account, the
// identifer and the ruleset ID.
func (api *API) getRuleset(ctx context.Context, identifierType RouteRoot, identifier, rulesetID string) (Ruleset, error) {
	uri := fmt.Sprintf("/%s/%s/rulesets/%s", identifierType, identifier, rulesetID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Ruleset{}, err
	}

	result := GetRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return Ruleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// CreateZoneRuleset creates a new ruleset for a zone.
//
// API reference: https://api.cloudflare.com/#zone-rulesets-create-zone-ruleset
func (api *API) CreateZoneRuleset(ctx context.Context, zoneID string, ruleset Ruleset) (Ruleset, error) {
	return api.createRuleset(ctx, ZoneRouteRoot, zoneID, ruleset)
}

// CreateAccountRuleset creates a new ruleset for an account.
//
// API reference: https://api.cloudflare.com/#account-rulesets-create-account-ruleset
func (api *API) CreateAccountRuleset(ctx context.Context, accountID string, ruleset Ruleset) (Ruleset, error) {
	return api.createRuleset(ctx, AccountRouteRoot, accountID, ruleset)
}

func (api *API) createRuleset(ctx context.Context, identifierType RouteRoot, identifier string, ruleset Ruleset) (Ruleset, error) {
	uri := fmt.Sprintf("/%s/%s/rulesets", identifierType, identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, ruleset)

	if err != nil {
		return Ruleset{}, err
	}

	result := CreateRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return Ruleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}

// DeleteZoneRuleset deletes a single ruleset for a zone.
//
// API reference: https://api.cloudflare.com/#zone-rulesets-delete-zone-ruleset
func (api *API) DeleteZoneRuleset(ctx context.Context, zoneID, rulesetID string) error {
	return api.deleteRuleset(ctx, ZoneRouteRoot, zoneID, rulesetID)
}

// DeleteAccountRuleset deletes a single ruleset for an account.
//
// API reference: https://api.cloudflare.com/#account-rulesets-delete-account-ruleset
func (api *API) DeleteAccountRuleset(ctx context.Context, accountID, rulesetID string) error {
	return api.deleteRuleset(ctx, AccountRouteRoot, accountID, rulesetID)
}

// deleteRuleset removes a ruleset based on the ruleset ID.
func (api *API) deleteRuleset(ctx context.Context, identifierType RouteRoot, identifier, rulesetID string) error {
	uri := fmt.Sprintf("/%s/%s/rulesets/%s", identifierType, identifier, rulesetID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	// The API is not implementing the standard response blob but returns an
	// empty response (204) in case of a success. So we are checking for the
	// response body size here.
	if len(res) > 0 {
		return errors.Wrap(errors.New(string(res)), errMakeRequestError)
	}

	return nil
}

// UpdateZoneRuleset updates a single ruleset for a zone.
//
// API reference: https://api.cloudflare.com/#zone-rulesets-update-a-zone-ruleset
func (api *API) UpdateZoneRuleset(ctx context.Context, zoneID, rulesetID, description string, rules []RulesetRule) (Ruleset, error) {
	return api.updateRuleset(ctx, ZoneRouteRoot, zoneID, rulesetID, description, rules)
}

// UpdateAccountRuleset updates a single ruleset for an account.
//
// API reference: https://api.cloudflare.com/#account-rulesets-update-account-ruleset
func (api *API) UpdateAccountRuleset(ctx context.Context, accountID, rulesetID, description string, rules []RulesetRule) (Ruleset, error) {
	return api.updateRuleset(ctx, AccountRouteRoot, accountID, rulesetID, description, rules)
}

// updateRuleset updates a ruleset based on the ruleset ID.
func (api *API) updateRuleset(ctx context.Context, identifierType RouteRoot, identifier, rulesetID, description string, rules []RulesetRule) (Ruleset, error) {
	uri := fmt.Sprintf("/%s/%s/rulesets/%s", identifierType, identifier, rulesetID)
	payload := UpdateRulesetRequest{Description: description, Rules: rules}
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, payload)
	if err != nil {
		return Ruleset{}, err
	}

	result := UpdateRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return Ruleset{}, errors.Wrap(err, errUnmarshalError)
	}

	return result.Result, nil
}
