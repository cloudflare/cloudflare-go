// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallWAFOverrideService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFOverrideService]
// method instead.
type FirewallWAFOverrideService struct {
	Options []option.RequestOption
}

// NewFirewallWAFOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallWAFOverrideService(opts ...option.RequestOption) (r *FirewallWAFOverrideService) {
	r = &FirewallWAFOverrideService{}
	r.Options = opts
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) New(ctx context.Context, zoneIdentifier string, body FirewallWAFOverrideNewParams, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideNewResponseEnvelope
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
func (r *FirewallWAFOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallWAFOverrideUpdateParams, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideUpdateResponseEnvelope
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
func (r *FirewallWAFOverrideService) List(ctx context.Context, zoneIdentifier string, query FirewallWAFOverrideListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[LegacyJhsOverride], err error) {
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
func (r *FirewallWAFOverrideService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallWAFOverrideListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[LegacyJhsOverride] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *FirewallWAFOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallWAFOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideDeleteResponseEnvelope
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
func (r *FirewallWAFOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallWAFOverrideGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LegacyJhsOverride struct {
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
	RewriteAction LegacyJhsOverrideRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]LegacyJhsOverrideRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string              `json:"urls"`
	JSON legacyJhsOverrideJSON `json:"-"`
}

// legacyJhsOverrideJSON contains the JSON metadata for the struct
// [LegacyJhsOverride]
type legacyJhsOverrideJSON struct {
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

func (r *LegacyJhsOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type LegacyJhsOverrideRewriteAction struct {
	// The WAF rule action to apply.
	Block     LegacyJhsOverrideRewriteActionBlock `json:"block"`
	Challenge interface{}                         `json:"challenge"`
	Default   interface{}                         `json:"default"`
	// The WAF rule action to apply.
	Disable  LegacyJhsOverrideRewriteActionDisable `json:"disable"`
	Simulate interface{}                           `json:"simulate"`
	JSON     legacyJhsOverrideRewriteActionJSON    `json:"-"`
}

// legacyJhsOverrideRewriteActionJSON contains the JSON metadata for the struct
// [LegacyJhsOverrideRewriteAction]
type legacyJhsOverrideRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsOverrideRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type LegacyJhsOverrideRewriteActionBlock string

const (
	LegacyJhsOverrideRewriteActionBlockChallenge LegacyJhsOverrideRewriteActionBlock = "challenge"
	LegacyJhsOverrideRewriteActionBlockBlock     LegacyJhsOverrideRewriteActionBlock = "block"
	LegacyJhsOverrideRewriteActionBlockSimulate  LegacyJhsOverrideRewriteActionBlock = "simulate"
	LegacyJhsOverrideRewriteActionBlockDisable   LegacyJhsOverrideRewriteActionBlock = "disable"
	LegacyJhsOverrideRewriteActionBlockDefault   LegacyJhsOverrideRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type LegacyJhsOverrideRewriteActionDisable string

const (
	LegacyJhsOverrideRewriteActionDisableChallenge LegacyJhsOverrideRewriteActionDisable = "challenge"
	LegacyJhsOverrideRewriteActionDisableBlock     LegacyJhsOverrideRewriteActionDisable = "block"
	LegacyJhsOverrideRewriteActionDisableSimulate  LegacyJhsOverrideRewriteActionDisable = "simulate"
	LegacyJhsOverrideRewriteActionDisableDisable   LegacyJhsOverrideRewriteActionDisable = "disable"
	LegacyJhsOverrideRewriteActionDisableDefault   LegacyJhsOverrideRewriteActionDisable = "default"
)

// The WAF rule action to apply.
type LegacyJhsOverrideRule string

const (
	LegacyJhsOverrideRuleChallenge LegacyJhsOverrideRule = "challenge"
	LegacyJhsOverrideRuleBlock     LegacyJhsOverrideRule = "block"
	LegacyJhsOverrideRuleSimulate  LegacyJhsOverrideRule = "simulate"
	LegacyJhsOverrideRuleDisable   LegacyJhsOverrideRule = "disable"
	LegacyJhsOverrideRuleDefault   LegacyJhsOverrideRule = "default"
)

type FirewallWAFOverrideDeleteResponse struct {
	// The unique identifier of the WAF override.
	ID   string                                `json:"id"`
	JSON firewallWAFOverrideDeleteResponseJSON `json:"-"`
}

// firewallWAFOverrideDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallWAFOverrideDeleteResponse]
type firewallWAFOverrideDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallWAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallWAFOverrideNewResponseEnvelope struct {
	Errors   []FirewallWAFOverrideNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideNewResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideNewResponseEnvelope]
type firewallWAFOverrideNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallWAFOverrideNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallWAFOverrideNewResponseEnvelopeErrors]
type firewallWAFOverrideNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallWAFOverrideNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideNewResponseEnvelopeMessages]
type firewallWAFOverrideNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideNewResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideNewResponseEnvelopeSuccessTrue FirewallWAFOverrideNewResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallWAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallWAFOverrideUpdateResponseEnvelope struct {
	Errors   []FirewallWAFOverrideUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideUpdateResponseEnvelope]
type firewallWAFOverrideUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideUpdateResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideUpdateResponseEnvelopeErrors]
type firewallWAFOverrideUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideUpdateResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideUpdateResponseEnvelopeMessages]
type firewallWAFOverrideUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideUpdateResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideUpdateResponseEnvelopeSuccessTrue FirewallWAFOverrideUpdateResponseEnvelopeSuccess = true
)

type FirewallWAFOverrideListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [FirewallWAFOverrideListParams]'s query parameters as
// `url.Values`.
func (r FirewallWAFOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallWAFOverrideDeleteResponseEnvelope struct {
	Result FirewallWAFOverrideDeleteResponse             `json:"result"`
	JSON   firewallWAFOverrideDeleteResponseEnvelopeJSON `json:"-"`
}

// firewallWAFOverrideDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideDeleteResponseEnvelope]
type firewallWAFOverrideDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelope struct {
	Errors   []FirewallWAFOverrideGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallWAFOverrideGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallWAFOverrideGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeJSON    `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallWAFOverrideGetResponseEnvelope]
type firewallWAFOverrideGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallWAFOverrideGetResponseEnvelopeErrors]
type firewallWAFOverrideGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallWAFOverrideGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallWAFOverrideGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallWAFOverrideGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallWAFOverrideGetResponseEnvelopeMessages]
type firewallWAFOverrideGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallWAFOverrideGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallWAFOverrideGetResponseEnvelopeSuccess bool

const (
	FirewallWAFOverrideGetResponseEnvelopeSuccessTrue FirewallWAFOverrideGetResponseEnvelopeSuccess = true
)
