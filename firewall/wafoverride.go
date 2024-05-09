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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) New(ctx context.Context, zoneIdentifier string, body WAFOverrideNewParams, opts ...option.RequestOption) (res *Override, err error) {
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
func (r *WAFOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body WAFOverrideUpdateParams, opts ...option.RequestOption) (res *Override, err error) {
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
func (r *WAFOverrideService) List(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Override], err error) {
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
func (r *WAFOverrideService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Override] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *WAFOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *WAFOverrideDeleteResponse, err error) {
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
func (r *WAFOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *Override, err error) {
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

type Override struct {
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
	RewriteAction RewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules WAFRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []OverrideURL `json:"urls"`
	JSON overrideJSON  `json:"-"`
}

// overrideJSON contains the JSON metadata for the struct [Override]
type overrideJSON struct {
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

func (r *Override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideJSON) RawJSON() string {
	return r.raw
}

type OverrideURL = string

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type RewriteAction struct {
	// The WAF rule action to apply.
	Block     RewriteActionBlock `json:"block"`
	Challenge string             `json:"challenge"`
	Default   string             `json:"default"`
	// The WAF rule action to apply.
	Disable  RewriteActionDisable `json:"disable"`
	Simulate string               `json:"simulate"`
	JSON     rewriteActionJSON    `json:"-"`
}

// rewriteActionJSON contains the JSON metadata for the struct [RewriteAction]
type rewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteActionJSON) RawJSON() string {
	return r.raw
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

type WAFRule map[string]WAFRuleItemItem

// The WAF rule action to apply.
type WAFRuleItemItem string

const (
	WAFRuleItemItemChallenge WAFRuleItemItem = "challenge"
	WAFRuleItemItemBlock     WAFRuleItemItem = "block"
	WAFRuleItemItemSimulate  WAFRuleItemItem = "simulate"
	WAFRuleItemItemDisable   WAFRuleItemItem = "disable"
	WAFRuleItemItemDefault   WAFRuleItemItem = "default"
)

func (r WAFRuleItemItem) IsKnown() bool {
	switch r {
	case WAFRuleItemItemChallenge, WAFRuleItemItemBlock, WAFRuleItemItemSimulate, WAFRuleItemItemDisable, WAFRuleItemItemDefault:
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

type WAFOverrideNewParams struct {
	Body interface{} `json:"body,required"`
}

func (r WAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
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
	Body interface{} `json:"body,required"`
}

func (r WAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
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
