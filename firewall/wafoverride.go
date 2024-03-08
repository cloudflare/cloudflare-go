// File generated from our OpenAPI spec by Stainless.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
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
func (r *WAFOverrideService) New(ctx context.Context, zoneIdentifier string, body WAFOverrideNewParams, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
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
func (r *WAFOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body WAFOverrideUpdateParams, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
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
func (r *WAFOverrideService) List(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[LegacyJhsOverride], err error) {
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
func (r *WAFOverrideService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query WAFOverrideListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[LegacyJhsOverride] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
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
func (r *WAFOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LegacyJhsOverride, err error) {
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

func (r legacyJhsOverrideJSON) RawJSON() string {
	return r.raw
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

func (r legacyJhsOverrideRewriteActionJSON) RawJSON() string {
	return r.raw
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
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideNewResponseEnvelope struct {
	Errors   []WAFOverrideNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WAFOverrideNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                        `json:"result,required,nullable"`
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

type WAFOverrideNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    wafOverrideNewResponseEnvelopeErrorsJSON `json:"-"`
}

// wafOverrideNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WAFOverrideNewResponseEnvelopeErrors]
type wafOverrideNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    wafOverrideNewResponseEnvelopeMessagesJSON `json:"-"`
}

// wafOverrideNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WAFOverrideNewResponseEnvelopeMessages]
type wafOverrideNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideNewResponseEnvelopeSuccess bool

const (
	WAFOverrideNewResponseEnvelopeSuccessTrue WAFOverrideNewResponseEnvelopeSuccess = true
)

type WAFOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r WAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WAFOverrideUpdateResponseEnvelope struct {
	Errors   []WAFOverrideUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WAFOverrideUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                           `json:"result,required,nullable"`
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

type WAFOverrideUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    wafOverrideUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// wafOverrideUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WAFOverrideUpdateResponseEnvelopeErrors]
type wafOverrideUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    wafOverrideUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// wafOverrideUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WAFOverrideUpdateResponseEnvelopeMessages]
type wafOverrideUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideUpdateResponseEnvelopeSuccess bool

const (
	WAFOverrideUpdateResponseEnvelopeSuccessTrue WAFOverrideUpdateResponseEnvelopeSuccess = true
)

type WAFOverrideListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFOverrideListParams]'s query parameters as `url.Values`.
func (r WAFOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
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
	Errors   []WAFOverrideGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WAFOverrideGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsOverride                        `json:"result,required,nullable"`
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

type WAFOverrideGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    wafOverrideGetResponseEnvelopeErrorsJSON `json:"-"`
}

// wafOverrideGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WAFOverrideGetResponseEnvelopeErrors]
type wafOverrideGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    wafOverrideGetResponseEnvelopeMessagesJSON `json:"-"`
}

// wafOverrideGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WAFOverrideGetResponseEnvelopeMessages]
type wafOverrideGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WAFOverrideGetResponseEnvelopeSuccess bool

const (
	WAFOverrideGetResponseEnvelopeSuccessTrue WAFOverrideGetResponseEnvelopeSuccess = true
)
