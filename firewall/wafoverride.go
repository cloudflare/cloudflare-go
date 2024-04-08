// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WAFOverrideService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWAFOverrideService] method
// instead.
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

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) New(ctx context.Context, zoneIdentifier string, body WAFOverrideNewParams, opts ...option.RequestOption) (res *WAFOverrideNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFOverrideNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body WAFOverrideUpdateParams, opts ...option.RequestOption) (res *WAFOverrideUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFOverrideUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) List(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFOverrideListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFOverrideListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, body WAFOverrideDeleteParams, opts ...option.RequestOption) (res *WAFOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFOverrideDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *WAFOverrideGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WAFOverrideGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WAFOverrideNewResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction WAFOverrideNewResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]WAFOverrideNewResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                   `json:"urls"`
	JSON wafOverrideNewResponseJSON `json:"-"`
}

// wafOverrideNewResponseJSON contains the JSON metadata for the struct
// [WAFOverrideNewResponse]
type wafOverrideNewResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFOverrideNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type WAFOverrideNewResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     WAFOverrideNewResponseRewriteActionBlock `json:"block"`
	Challenge string                                   `json:"challenge"`
	Default   string                                   `json:"default"`
	// The WAF rule action to apply.
	Disable  WAFOverrideNewResponseRewriteActionDisable `json:"disable"`
	Simulate string                                     `json:"simulate"`
	JSON     wafOverrideNewResponseRewriteActionJSON    `json:"-"`
}

// wafOverrideNewResponseRewriteActionJSON contains the JSON metadata for the
// struct [WAFOverrideNewResponseRewriteAction]
type wafOverrideNewResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideNewResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseRewriteActionJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type WAFOverrideNewResponseRewriteActionBlock string

const (
	WAFOverrideNewResponseRewriteActionBlockChallenge WAFOverrideNewResponseRewriteActionBlock = "challenge"
	WAFOverrideNewResponseRewriteActionBlockBlock     WAFOverrideNewResponseRewriteActionBlock = "block"
	WAFOverrideNewResponseRewriteActionBlockSimulate  WAFOverrideNewResponseRewriteActionBlock = "simulate"
	WAFOverrideNewResponseRewriteActionBlockDisable   WAFOverrideNewResponseRewriteActionBlock = "disable"
	WAFOverrideNewResponseRewriteActionBlockDefault   WAFOverrideNewResponseRewriteActionBlock = "default"
)

func (r WAFOverrideNewResponseRewriteActionBlock) IsKnown() bool {
	switch r {
	case WAFOverrideNewResponseRewriteActionBlockChallenge, WAFOverrideNewResponseRewriteActionBlockBlock, WAFOverrideNewResponseRewriteActionBlockSimulate, WAFOverrideNewResponseRewriteActionBlockDisable, WAFOverrideNewResponseRewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideNewResponseRewriteActionDisable string

const (
	WAFOverrideNewResponseRewriteActionDisableChallenge WAFOverrideNewResponseRewriteActionDisable = "challenge"
	WAFOverrideNewResponseRewriteActionDisableBlock     WAFOverrideNewResponseRewriteActionDisable = "block"
	WAFOverrideNewResponseRewriteActionDisableSimulate  WAFOverrideNewResponseRewriteActionDisable = "simulate"
	WAFOverrideNewResponseRewriteActionDisableDisable   WAFOverrideNewResponseRewriteActionDisable = "disable"
	WAFOverrideNewResponseRewriteActionDisableDefault   WAFOverrideNewResponseRewriteActionDisable = "default"
)

func (r WAFOverrideNewResponseRewriteActionDisable) IsKnown() bool {
	switch r {
	case WAFOverrideNewResponseRewriteActionDisableChallenge, WAFOverrideNewResponseRewriteActionDisableBlock, WAFOverrideNewResponseRewriteActionDisableSimulate, WAFOverrideNewResponseRewriteActionDisableDisable, WAFOverrideNewResponseRewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideNewResponseRule string

const (
	WAFOverrideNewResponseRuleChallenge WAFOverrideNewResponseRule = "challenge"
	WAFOverrideNewResponseRuleBlock     WAFOverrideNewResponseRule = "block"
	WAFOverrideNewResponseRuleSimulate  WAFOverrideNewResponseRule = "simulate"
	WAFOverrideNewResponseRuleDisable   WAFOverrideNewResponseRule = "disable"
	WAFOverrideNewResponseRuleDefault   WAFOverrideNewResponseRule = "default"
)

func (r WAFOverrideNewResponseRule) IsKnown() bool {
	switch r {
	case WAFOverrideNewResponseRuleChallenge, WAFOverrideNewResponseRuleBlock, WAFOverrideNewResponseRuleSimulate, WAFOverrideNewResponseRuleDisable, WAFOverrideNewResponseRuleDefault:
		return true
	}
	return false
}

type WAFOverrideUpdateResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction WAFOverrideUpdateResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]WAFOverrideUpdateResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                      `json:"urls"`
	JSON wafOverrideUpdateResponseJSON `json:"-"`
}

// wafOverrideUpdateResponseJSON contains the JSON metadata for the struct
// [WAFOverrideUpdateResponse]
type wafOverrideUpdateResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type WAFOverrideUpdateResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     WAFOverrideUpdateResponseRewriteActionBlock `json:"block"`
	Challenge string                                      `json:"challenge"`
	Default   string                                      `json:"default"`
	// The WAF rule action to apply.
	Disable  WAFOverrideUpdateResponseRewriteActionDisable `json:"disable"`
	Simulate string                                        `json:"simulate"`
	JSON     wafOverrideUpdateResponseRewriteActionJSON    `json:"-"`
}

// wafOverrideUpdateResponseRewriteActionJSON contains the JSON metadata for the
// struct [WAFOverrideUpdateResponseRewriteAction]
type wafOverrideUpdateResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseRewriteActionJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type WAFOverrideUpdateResponseRewriteActionBlock string

const (
	WAFOverrideUpdateResponseRewriteActionBlockChallenge WAFOverrideUpdateResponseRewriteActionBlock = "challenge"
	WAFOverrideUpdateResponseRewriteActionBlockBlock     WAFOverrideUpdateResponseRewriteActionBlock = "block"
	WAFOverrideUpdateResponseRewriteActionBlockSimulate  WAFOverrideUpdateResponseRewriteActionBlock = "simulate"
	WAFOverrideUpdateResponseRewriteActionBlockDisable   WAFOverrideUpdateResponseRewriteActionBlock = "disable"
	WAFOverrideUpdateResponseRewriteActionBlockDefault   WAFOverrideUpdateResponseRewriteActionBlock = "default"
)

func (r WAFOverrideUpdateResponseRewriteActionBlock) IsKnown() bool {
	switch r {
	case WAFOverrideUpdateResponseRewriteActionBlockChallenge, WAFOverrideUpdateResponseRewriteActionBlockBlock, WAFOverrideUpdateResponseRewriteActionBlockSimulate, WAFOverrideUpdateResponseRewriteActionBlockDisable, WAFOverrideUpdateResponseRewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideUpdateResponseRewriteActionDisable string

const (
	WAFOverrideUpdateResponseRewriteActionDisableChallenge WAFOverrideUpdateResponseRewriteActionDisable = "challenge"
	WAFOverrideUpdateResponseRewriteActionDisableBlock     WAFOverrideUpdateResponseRewriteActionDisable = "block"
	WAFOverrideUpdateResponseRewriteActionDisableSimulate  WAFOverrideUpdateResponseRewriteActionDisable = "simulate"
	WAFOverrideUpdateResponseRewriteActionDisableDisable   WAFOverrideUpdateResponseRewriteActionDisable = "disable"
	WAFOverrideUpdateResponseRewriteActionDisableDefault   WAFOverrideUpdateResponseRewriteActionDisable = "default"
)

func (r WAFOverrideUpdateResponseRewriteActionDisable) IsKnown() bool {
	switch r {
	case WAFOverrideUpdateResponseRewriteActionDisableChallenge, WAFOverrideUpdateResponseRewriteActionDisableBlock, WAFOverrideUpdateResponseRewriteActionDisableSimulate, WAFOverrideUpdateResponseRewriteActionDisableDisable, WAFOverrideUpdateResponseRewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideUpdateResponseRule string

const (
	WAFOverrideUpdateResponseRuleChallenge WAFOverrideUpdateResponseRule = "challenge"
	WAFOverrideUpdateResponseRuleBlock     WAFOverrideUpdateResponseRule = "block"
	WAFOverrideUpdateResponseRuleSimulate  WAFOverrideUpdateResponseRule = "simulate"
	WAFOverrideUpdateResponseRuleDisable   WAFOverrideUpdateResponseRule = "disable"
	WAFOverrideUpdateResponseRuleDefault   WAFOverrideUpdateResponseRule = "default"
)

func (r WAFOverrideUpdateResponseRule) IsKnown() bool {
	switch r {
	case WAFOverrideUpdateResponseRuleChallenge, WAFOverrideUpdateResponseRuleBlock, WAFOverrideUpdateResponseRuleSimulate, WAFOverrideUpdateResponseRuleDisable, WAFOverrideUpdateResponseRuleDefault:
		return true
	}
	return false
}

type WAFOverrideListResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id,required"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused,required"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string `json:"urls,required"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction WAFOverrideListResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]WAFOverrideListResponseRule `json:"rules"`
	JSON  wafOverrideListResponseJSON            `json:"-"`
}

// wafOverrideListResponseJSON contains the JSON metadata for the struct
// [WAFOverrideListResponse]
type wafOverrideListResponseJSON struct {
	ID            apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	URLs          apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFOverrideListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideListResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type WAFOverrideListResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     WAFOverrideListResponseRewriteActionBlock `json:"block"`
	Challenge string                                    `json:"challenge"`
	Default   string                                    `json:"default"`
	// The WAF rule action to apply.
	Disable  WAFOverrideListResponseRewriteActionDisable `json:"disable"`
	Simulate string                                      `json:"simulate"`
	JSON     wafOverrideListResponseRewriteActionJSON    `json:"-"`
}

// wafOverrideListResponseRewriteActionJSON contains the JSON metadata for the
// struct [WAFOverrideListResponseRewriteAction]
type wafOverrideListResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideListResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideListResponseRewriteActionJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type WAFOverrideListResponseRewriteActionBlock string

const (
	WAFOverrideListResponseRewriteActionBlockChallenge WAFOverrideListResponseRewriteActionBlock = "challenge"
	WAFOverrideListResponseRewriteActionBlockBlock     WAFOverrideListResponseRewriteActionBlock = "block"
	WAFOverrideListResponseRewriteActionBlockSimulate  WAFOverrideListResponseRewriteActionBlock = "simulate"
	WAFOverrideListResponseRewriteActionBlockDisable   WAFOverrideListResponseRewriteActionBlock = "disable"
	WAFOverrideListResponseRewriteActionBlockDefault   WAFOverrideListResponseRewriteActionBlock = "default"
)

func (r WAFOverrideListResponseRewriteActionBlock) IsKnown() bool {
	switch r {
	case WAFOverrideListResponseRewriteActionBlockChallenge, WAFOverrideListResponseRewriteActionBlockBlock, WAFOverrideListResponseRewriteActionBlockSimulate, WAFOverrideListResponseRewriteActionBlockDisable, WAFOverrideListResponseRewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideListResponseRewriteActionDisable string

const (
	WAFOverrideListResponseRewriteActionDisableChallenge WAFOverrideListResponseRewriteActionDisable = "challenge"
	WAFOverrideListResponseRewriteActionDisableBlock     WAFOverrideListResponseRewriteActionDisable = "block"
	WAFOverrideListResponseRewriteActionDisableSimulate  WAFOverrideListResponseRewriteActionDisable = "simulate"
	WAFOverrideListResponseRewriteActionDisableDisable   WAFOverrideListResponseRewriteActionDisable = "disable"
	WAFOverrideListResponseRewriteActionDisableDefault   WAFOverrideListResponseRewriteActionDisable = "default"
)

func (r WAFOverrideListResponseRewriteActionDisable) IsKnown() bool {
	switch r {
	case WAFOverrideListResponseRewriteActionDisableChallenge, WAFOverrideListResponseRewriteActionDisableBlock, WAFOverrideListResponseRewriteActionDisableSimulate, WAFOverrideListResponseRewriteActionDisableDisable, WAFOverrideListResponseRewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideListResponseRule string

const (
	WAFOverrideListResponseRuleChallenge WAFOverrideListResponseRule = "challenge"
	WAFOverrideListResponseRuleBlock     WAFOverrideListResponseRule = "block"
	WAFOverrideListResponseRuleSimulate  WAFOverrideListResponseRule = "simulate"
	WAFOverrideListResponseRuleDisable   WAFOverrideListResponseRule = "disable"
	WAFOverrideListResponseRuleDefault   WAFOverrideListResponseRule = "default"
)

func (r WAFOverrideListResponseRule) IsKnown() bool {
	switch r {
	case WAFOverrideListResponseRuleChallenge, WAFOverrideListResponseRuleBlock, WAFOverrideListResponseRuleSimulate, WAFOverrideListResponseRuleDisable, WAFOverrideListResponseRuleDefault:
		return true
	}
	return false
}

type WAFOverrideDeleteResponse struct {
	// The unique identifier of the WAF override.
	ID   string                        `json:"id"`
	JSON wafOverrideDeleteResponseJSON `json:"-"`
}

// wafOverrideDeleteResponseJSON contains the JSON metadata for the struct
// [WAFOverrideDeleteResponse]
type wafOverrideDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideGetResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction WAFOverrideGetResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]WAFOverrideGetResponseRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                   `json:"urls"`
	JSON wafOverrideGetResponseJSON `json:"-"`
}

// wafOverrideGetResponseJSON contains the JSON metadata for the struct
// [WAFOverrideGetResponse]
type wafOverrideGetResponseJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *WAFOverrideGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type WAFOverrideGetResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     WAFOverrideGetResponseRewriteActionBlock `json:"block"`
	Challenge string                                   `json:"challenge"`
	Default   string                                   `json:"default"`
	// The WAF rule action to apply.
	Disable  WAFOverrideGetResponseRewriteActionDisable `json:"disable"`
	Simulate string                                     `json:"simulate"`
	JSON     wafOverrideGetResponseRewriteActionJSON    `json:"-"`
}

// wafOverrideGetResponseRewriteActionJSON contains the JSON metadata for the
// struct [WAFOverrideGetResponseRewriteAction]
type wafOverrideGetResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideGetResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseRewriteActionJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type WAFOverrideGetResponseRewriteActionBlock string

const (
	WAFOverrideGetResponseRewriteActionBlockChallenge WAFOverrideGetResponseRewriteActionBlock = "challenge"
	WAFOverrideGetResponseRewriteActionBlockBlock     WAFOverrideGetResponseRewriteActionBlock = "block"
	WAFOverrideGetResponseRewriteActionBlockSimulate  WAFOverrideGetResponseRewriteActionBlock = "simulate"
	WAFOverrideGetResponseRewriteActionBlockDisable   WAFOverrideGetResponseRewriteActionBlock = "disable"
	WAFOverrideGetResponseRewriteActionBlockDefault   WAFOverrideGetResponseRewriteActionBlock = "default"
)

func (r WAFOverrideGetResponseRewriteActionBlock) IsKnown() bool {
	switch r {
	case WAFOverrideGetResponseRewriteActionBlockChallenge, WAFOverrideGetResponseRewriteActionBlockBlock, WAFOverrideGetResponseRewriteActionBlockSimulate, WAFOverrideGetResponseRewriteActionBlockDisable, WAFOverrideGetResponseRewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideGetResponseRewriteActionDisable string

const (
	WAFOverrideGetResponseRewriteActionDisableChallenge WAFOverrideGetResponseRewriteActionDisable = "challenge"
	WAFOverrideGetResponseRewriteActionDisableBlock     WAFOverrideGetResponseRewriteActionDisable = "block"
	WAFOverrideGetResponseRewriteActionDisableSimulate  WAFOverrideGetResponseRewriteActionDisable = "simulate"
	WAFOverrideGetResponseRewriteActionDisableDisable   WAFOverrideGetResponseRewriteActionDisable = "disable"
	WAFOverrideGetResponseRewriteActionDisableDefault   WAFOverrideGetResponseRewriteActionDisable = "default"
)

func (r WAFOverrideGetResponseRewriteActionDisable) IsKnown() bool {
	switch r {
	case WAFOverrideGetResponseRewriteActionDisableChallenge, WAFOverrideGetResponseRewriteActionDisableBlock, WAFOverrideGetResponseRewriteActionDisableSimulate, WAFOverrideGetResponseRewriteActionDisableDisable, WAFOverrideGetResponseRewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type WAFOverrideGetResponseRule string

const (
	WAFOverrideGetResponseRuleChallenge WAFOverrideGetResponseRule = "challenge"
	WAFOverrideGetResponseRuleBlock     WAFOverrideGetResponseRule = "block"
	WAFOverrideGetResponseRuleSimulate  WAFOverrideGetResponseRule = "simulate"
	WAFOverrideGetResponseRuleDisable   WAFOverrideGetResponseRule = "disable"
	WAFOverrideGetResponseRuleDefault   WAFOverrideGetResponseRule = "default"
)

func (r WAFOverrideGetResponseRule) IsKnown() bool {
	switch r {
	case WAFOverrideGetResponseRuleChallenge, WAFOverrideGetResponseRuleBlock, WAFOverrideGetResponseRuleSimulate, WAFOverrideGetResponseRuleDisable, WAFOverrideGetResponseRuleDefault:
		return true
	}
	return false
}

type WAFOverrideNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   WAFOverrideNewResponse                                    `json:"result,required"`
	// Whether the API call was successful
	Success WAFOverrideNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideNewResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideNewResponseEnvelope]
type wafOverrideNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideNewResponseEnvelopeSuccess bool

const (
	WAFOverrideNewResponseEnvelopeSuccessTrue WAFOverrideNewResponseEnvelopeSuccess = true
)

func (r WAFOverrideNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   WAFOverrideUpdateResponse                                 `json:"result,required"`
	// Whether the API call was successful
	Success WAFOverrideUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideUpdateResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideUpdateResponseEnvelope]
type wafOverrideUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideUpdateResponseEnvelopeSuccess bool

const (
	WAFOverrideUpdateResponseEnvelopeSuccessTrue WAFOverrideUpdateResponseEnvelopeSuccess = true
)

func (r WAFOverrideUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFOverrideListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFOverrideListParams]'s query parameters as `url.Values`.
func (r WAFOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WAFOverrideDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WAFOverrideDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideDeleteResponseEnvelope struct {
	Result WAFOverrideDeleteResponse             `json:"result"`
	JSON   wafOverrideDeleteResponseEnvelopeJSON `json:"-"`
}

// wafOverrideDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideDeleteResponseEnvelope]
type wafOverrideDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   WAFOverrideGetResponse                                    `json:"result,required"`
	// Whether the API call was successful
	Success WAFOverrideGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideGetResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideGetResponseEnvelope]
type wafOverrideGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideGetResponseEnvelopeSuccess bool

const (
	WAFOverrideGetResponseEnvelopeSuccessTrue WAFOverrideGetResponseEnvelopeSuccess = true
)

func (r WAFOverrideGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
