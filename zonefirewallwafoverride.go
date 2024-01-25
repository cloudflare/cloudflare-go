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

// ZoneFirewallWafOverrideService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneFirewallWafOverrideService] method instead.
type ZoneFirewallWafOverrideService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallWafOverrideService(opts ...option.RequestOption) (r *ZoneFirewallWafOverrideService) {
	r = &ZoneFirewallWafOverrideService{}
	r.Options = opts
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallWafOverrideUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) WafOverridesNewAWafOverride(ctx context.Context, zoneIdentifier string, body ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) WafOverridesListWafOverrides(ctx context.Context, zoneIdentifier string, query ZoneFirewallWafOverrideWafOverridesListWafOverridesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallWafOverrideWafOverridesListWafOverridesResponse], err error) {
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

type ZoneFirewallWafOverrideGetResponse struct {
	Errors   []ZoneFirewallWafOverrideGetResponseError   `json:"errors"`
	Messages []ZoneFirewallWafOverrideGetResponseMessage `json:"messages"`
	Result   ZoneFirewallWafOverrideGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafOverrideGetResponseSuccess `json:"success"`
	JSON    zoneFirewallWafOverrideGetResponseJSON    `json:"-"`
}

// zoneFirewallWafOverrideGetResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallWafOverrideGetResponse]
type zoneFirewallWafOverrideGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneFirewallWafOverrideGetResponseErrorJSON `json:"-"`
}

// zoneFirewallWafOverrideGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideGetResponseError]
type zoneFirewallWafOverrideGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneFirewallWafOverrideGetResponseMessageJSON `json:"-"`
}

// zoneFirewallWafOverrideGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideGetResponseMessage]
type zoneFirewallWafOverrideGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideGetResponseResult struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction ZoneFirewallWafOverrideGetResponseResultRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                                     `json:"urls"`
	JSON zoneFirewallWafOverrideGetResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideGetResponseResult]
type zoneFirewallWafOverrideGetResponseResultJSON struct {
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

func (r *ZoneFirewallWafOverrideGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type ZoneFirewallWafOverrideGetResponseResultRewriteAction struct {
	// The WAF rule action to apply.
	Block     ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock `json:"block"`
	Challenge interface{}                                                `json:"challenge"`
	Default   interface{}                                                `json:"default"`
	// The WAF rule action to apply.
	Disable  ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                  `json:"simulate"`
	JSON     zoneFirewallWafOverrideGetResponseResultRewriteActionJSON    `json:"-"`
}

// zoneFirewallWafOverrideGetResponseResultRewriteActionJSON contains the JSON
// metadata for the struct [ZoneFirewallWafOverrideGetResponseResultRewriteAction]
type zoneFirewallWafOverrideGetResponseResultRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideGetResponseResultRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock string

const (
	ZoneFirewallWafOverrideGetResponseResultRewriteActionBlockChallenge ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock = "challenge"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionBlockBlock     ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock = "block"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionBlockSimulate  ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock = "simulate"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionBlockDisable   ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock = "disable"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionBlockDefault   ZoneFirewallWafOverrideGetResponseResultRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable string

const (
	ZoneFirewallWafOverrideGetResponseResultRewriteActionDisableChallenge ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable = "challenge"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionDisableBlock     ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable = "block"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionDisableSimulate  ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable = "simulate"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionDisableDisable   ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable = "disable"
	ZoneFirewallWafOverrideGetResponseResultRewriteActionDisableDefault   ZoneFirewallWafOverrideGetResponseResultRewriteActionDisable = "default"
)

// Whether the API call was successful
type ZoneFirewallWafOverrideGetResponseSuccess bool

const (
	ZoneFirewallWafOverrideGetResponseSuccessTrue ZoneFirewallWafOverrideGetResponseSuccess = true
)

type ZoneFirewallWafOverrideUpdateResponse struct {
	Errors   []ZoneFirewallWafOverrideUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallWafOverrideUpdateResponseMessage `json:"messages"`
	Result   ZoneFirewallWafOverrideUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafOverrideUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallWafOverrideUpdateResponseJSON    `json:"-"`
}

// zoneFirewallWafOverrideUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideUpdateResponse]
type zoneFirewallWafOverrideUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideUpdateResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneFirewallWafOverrideUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallWafOverrideUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideUpdateResponseError]
type zoneFirewallWafOverrideUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideUpdateResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneFirewallWafOverrideUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallWafOverrideUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideUpdateResponseMessage]
type zoneFirewallWafOverrideUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideUpdateResponseResult struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction ZoneFirewallWafOverrideUpdateResponseResultRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                                        `json:"urls"`
	JSON zoneFirewallWafOverrideUpdateResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideUpdateResponseResult]
type zoneFirewallWafOverrideUpdateResponseResultJSON struct {
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

func (r *ZoneFirewallWafOverrideUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type ZoneFirewallWafOverrideUpdateResponseResultRewriteAction struct {
	// The WAF rule action to apply.
	Block     ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock `json:"block"`
	Challenge interface{}                                                   `json:"challenge"`
	Default   interface{}                                                   `json:"default"`
	// The WAF rule action to apply.
	Disable  ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                     `json:"simulate"`
	JSON     zoneFirewallWafOverrideUpdateResponseResultRewriteActionJSON    `json:"-"`
}

// zoneFirewallWafOverrideUpdateResponseResultRewriteActionJSON contains the JSON
// metadata for the struct
// [ZoneFirewallWafOverrideUpdateResponseResultRewriteAction]
type zoneFirewallWafOverrideUpdateResponseResultRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideUpdateResponseResultRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock string

const (
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlockChallenge ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock = "challenge"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlockBlock     ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock = "block"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlockSimulate  ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock = "simulate"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlockDisable   ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock = "disable"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlockDefault   ZoneFirewallWafOverrideUpdateResponseResultRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable string

const (
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisableChallenge ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable = "challenge"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisableBlock     ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable = "block"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisableSimulate  ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable = "simulate"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisableDisable   ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable = "disable"
	ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisableDefault   ZoneFirewallWafOverrideUpdateResponseResultRewriteActionDisable = "default"
)

// Whether the API call was successful
type ZoneFirewallWafOverrideUpdateResponseSuccess bool

const (
	ZoneFirewallWafOverrideUpdateResponseSuccessTrue ZoneFirewallWafOverrideUpdateResponseSuccess = true
)

type ZoneFirewallWafOverrideDeleteResponse struct {
	Result ZoneFirewallWafOverrideDeleteResponseResult `json:"result"`
	JSON   zoneFirewallWafOverrideDeleteResponseJSON   `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideDeleteResponse]
type zoneFirewallWafOverrideDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideDeleteResponseResult struct {
	// The unique identifier of the WAF override.
	ID   string                                          `json:"id"`
	JSON zoneFirewallWafOverrideDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideDeleteResponseResult]
type zoneFirewallWafOverrideDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponse struct {
	Errors   []ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseError   `json:"errors"`
	Messages []ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessage `json:"messages"`
	Result   ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseSuccess `json:"success"`
	JSON    zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseJSON    `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseJSON contains the JSON
// metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponse]
type zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseErrorJSON `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseError]
type zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessageJSON `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessage]
type zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResult struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                                                             `json:"urls"`
	JSON zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResult]
type zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultJSON struct {
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

func (r *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteAction struct {
	// The WAF rule action to apply.
	Block     ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock `json:"block"`
	Challenge interface{}                                                                        `json:"challenge"`
	Default   interface{}                                                                        `json:"default"`
	// The WAF rule action to apply.
	Disable  ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                                          `json:"simulate"`
	JSON     zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionJSON    `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteAction]
type zoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock string

const (
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlockChallenge ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock = "challenge"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlockBlock     ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock = "block"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlockSimulate  ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock = "simulate"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlockDisable   ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock = "disable"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlockDefault   ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable string

const (
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisableChallenge ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable = "challenge"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisableBlock     ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable = "block"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisableSimulate  ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable = "simulate"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisableDisable   ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable = "disable"
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisableDefault   ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseResultRewriteActionDisable = "default"
)

// Whether the API call was successful
type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseSuccess bool

const (
	ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseSuccessTrue ZoneFirewallWafOverrideWafOverridesNewAWafOverrideResponseSuccess = true
)

type ZoneFirewallWafOverrideWafOverridesListWafOverridesResponse struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules interface{} `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string                                                        `json:"urls"`
	JSON zoneFirewallWafOverrideWafOverridesListWafOverridesResponseJSON `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesListWafOverridesResponseJSON contains the
// JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesListWafOverridesResponse]
type zoneFirewallWafOverrideWafOverridesListWafOverridesResponseJSON struct {
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

func (r *ZoneFirewallWafOverrideWafOverridesListWafOverridesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteAction struct {
	// The WAF rule action to apply.
	Block     ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock `json:"block"`
	Challenge interface{}                                                                   `json:"challenge"`
	Default   interface{}                                                                   `json:"default"`
	// The WAF rule action to apply.
	Disable  ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable `json:"disable"`
	Simulate interface{}                                                                     `json:"simulate"`
	JSON     zoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionJSON    `json:"-"`
}

// zoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionJSON
// contains the JSON metadata for the struct
// [ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteAction]
type zoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The WAF rule action to apply.
type ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock string

const (
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlockChallenge ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock = "challenge"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlockBlock     ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock = "block"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlockSimulate  ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock = "simulate"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlockDisable   ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock = "disable"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlockDefault   ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionBlock = "default"
)

// The WAF rule action to apply.
type ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable string

const (
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisableChallenge ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable = "challenge"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisableBlock     ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable = "block"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisableSimulate  ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable = "simulate"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisableDisable   ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable = "disable"
	ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisableDefault   ZoneFirewallWafOverrideWafOverridesListWafOverridesResponseRewriteActionDisable = "default"
)

type ZoneFirewallWafOverrideUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallWafOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallWafOverrideWafOverridesNewAWafOverrideParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallWafOverrideWafOverridesListWafOverridesParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneFirewallWafOverrideWafOverridesListWafOverridesParams]'s query parameters
// as `url.Values`.
func (r ZoneFirewallWafOverrideWafOverridesListWafOverridesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
