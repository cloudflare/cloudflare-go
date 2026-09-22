// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// WAFOverrideService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFOverrideService] method instead.
type WAFOverrideService struct {
	Options []option.RequestOption
}

// NewWAFOverrideService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFOverrideService(opts ...option.RequestOption) (r *WAFOverrideService) {
	r = &WAFOverrideService{}
	r.Options = opts
	return
}

// **This endpoint has been deprecated and returns 410 Gone. Please use the
// [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**
//
// Previously created a URI-based WAF override for a zone.
//
// Deprecated: deprecated
func (r *WAFOverrideService) New(ctx context.Context, params WAFOverrideNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// **This endpoint has been deprecated and returns 410 Gone. Please use the
// [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**
//
// Previously updated an existing URI-based WAF override.
//
// Deprecated: deprecated
func (r *WAFOverrideService) Update(ctx context.Context, overridesID string, params WAFOverrideUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", params.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// **This endpoint has been deprecated and returns 410 Gone. Please use the
// [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**
//
// Previously fetched the URI-based WAF overrides in a zone.
//
// Deprecated: deprecated
func (r *WAFOverrideService) List(ctx context.Context, params WAFOverrideListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return err
}

// **This endpoint has been deprecated and returns 410 Gone. Please use the
// [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**
//
// Previously deleted an existing URI-based WAF override.
//
// Deprecated: deprecated
func (r *WAFOverrideService) Delete(ctx context.Context, overridesID string, body WAFOverrideDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", body.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// **This endpoint has been deprecated and returns 410 Gone. Please use the
// [Rulesets API](https://developers.cloudflare.com/ruleset-engine/) instead.**
//
// Previously fetched the details of a URI-based WAF override.
//
// Deprecated: deprecated
func (r *WAFOverrideService) Get(ctx context.Context, overridesID string, query WAFOverrideGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return err
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return err
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", query.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type OverrideURLParam = string

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type RewriteActionParam struct {
	// The WAF rule action to apply.
	Block param.Field[RewriteActionBlock] `json:"block"`
	// The WAF rule action to apply.
	Challenge param.Field[RewriteActionChallenge] `json:"challenge"`
	// The WAF rule action to apply.
	Default param.Field[RewriteActionDefault] `json:"default"`
	// The WAF rule action to apply.
	Disable param.Field[RewriteActionDisable] `json:"disable"`
	// The WAF rule action to apply.
	Simulate param.Field[RewriteActionSimulate] `json:"simulate"`
}

func (r RewriteActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The WAF rule action to apply.
type RewriteActionBlock string

const (
	RewriteActionBlockChallenge RewriteActionBlock = "challenge"
	RewriteActionBlockBlock     RewriteActionBlock = "block"
	RewriteActionBlockSimulate  RewriteActionBlock = "simulate"
	RewriteActionBlockDisable   RewriteActionBlock = "disable"
	RewriteActionBlockDefault   RewriteActionBlock = "default"
)

func (r RewriteActionBlock) IsKnown() bool {
	switch r {
	case RewriteActionBlockChallenge, RewriteActionBlockBlock, RewriteActionBlockSimulate, RewriteActionBlockDisable, RewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionChallenge string

const (
	RewriteActionChallengeChallenge RewriteActionChallenge = "challenge"
	RewriteActionChallengeBlock     RewriteActionChallenge = "block"
	RewriteActionChallengeSimulate  RewriteActionChallenge = "simulate"
	RewriteActionChallengeDisable   RewriteActionChallenge = "disable"
	RewriteActionChallengeDefault   RewriteActionChallenge = "default"
)

func (r RewriteActionChallenge) IsKnown() bool {
	switch r {
	case RewriteActionChallengeChallenge, RewriteActionChallengeBlock, RewriteActionChallengeSimulate, RewriteActionChallengeDisable, RewriteActionChallengeDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionDefault string

const (
	RewriteActionDefaultChallenge RewriteActionDefault = "challenge"
	RewriteActionDefaultBlock     RewriteActionDefault = "block"
	RewriteActionDefaultSimulate  RewriteActionDefault = "simulate"
	RewriteActionDefaultDisable   RewriteActionDefault = "disable"
	RewriteActionDefaultDefault   RewriteActionDefault = "default"
)

func (r RewriteActionDefault) IsKnown() bool {
	switch r {
	case RewriteActionDefaultChallenge, RewriteActionDefaultBlock, RewriteActionDefaultSimulate, RewriteActionDefaultDisable, RewriteActionDefaultDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionDisable string

const (
	RewriteActionDisableChallenge RewriteActionDisable = "challenge"
	RewriteActionDisableBlock     RewriteActionDisable = "block"
	RewriteActionDisableSimulate  RewriteActionDisable = "simulate"
	RewriteActionDisableDisable   RewriteActionDisable = "disable"
	RewriteActionDisableDefault   RewriteActionDisable = "default"
)

func (r RewriteActionDisable) IsKnown() bool {
	switch r {
	case RewriteActionDisableChallenge, RewriteActionDisableBlock, RewriteActionDisableSimulate, RewriteActionDisableDisable, RewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionSimulate string

const (
	RewriteActionSimulateChallenge RewriteActionSimulate = "challenge"
	RewriteActionSimulateBlock     RewriteActionSimulate = "block"
	RewriteActionSimulateSimulate  RewriteActionSimulate = "simulate"
	RewriteActionSimulateDisable   RewriteActionSimulate = "disable"
	RewriteActionSimulateDefault   RewriteActionSimulate = "default"
)

func (r RewriteActionSimulate) IsKnown() bool {
	switch r {
	case RewriteActionSimulateChallenge, RewriteActionSimulateBlock, RewriteActionSimulateSimulate, RewriteActionSimulateDisable, RewriteActionSimulateDefault:
		return true
	}
	return false
}

type WAFRuleParam map[string]WAFRuleItem

// The WAF rule action to apply.
type WAFRuleItem string

const (
	WAFRuleItemChallenge WAFRuleItem = "challenge"
	WAFRuleItemBlock     WAFRuleItem = "block"
	WAFRuleItemSimulate  WAFRuleItem = "simulate"
	WAFRuleItemDisable   WAFRuleItem = "disable"
	WAFRuleItemDefault   WAFRuleItem = "default"
)

func (r WAFRuleItem) IsKnown() bool {
	switch r {
	case WAFRuleItemChallenge, WAFRuleItemBlock, WAFRuleItemSimulate, WAFRuleItemDisable, WAFRuleItemDefault:
		return true
	}
	return false
}

type WAFOverrideNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls" api:"required"`
}

func (r WAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WAFOverrideUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Defines an identifier.
	ID param.Field[string] `json:"id" api:"required"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction param.Field[RewriteActionParam] `json:"rewrite_action" api:"required"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules param.Field[WAFRuleParam] `json:"rules" api:"required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls" api:"required"`
}

func (r WAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WAFOverrideListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFOverrideListParams]'s query parameters as `url.Values`.
func (r WAFOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WAFOverrideDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type WAFOverrideGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}
