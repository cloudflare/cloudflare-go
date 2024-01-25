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

// ZoneEmailRoutingRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneEmailRoutingRuleService]
// method instead.
type ZoneEmailRoutingRuleService struct {
	Options   []option.RequestOption
	CatchAlls *ZoneEmailRoutingRuleCatchAllService
}

// NewZoneEmailRoutingRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingRuleService(opts ...option.RequestOption) (r *ZoneEmailRoutingRuleService) {
	r = &ZoneEmailRoutingRuleService{}
	r.Options = opts
	r.CatchAlls = NewZoneEmailRoutingRuleCatchAllService(opts...)
	return
}

// Get information for a specific routing rule already created.
func (r *ZoneEmailRoutingRuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *ZoneEmailRoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body ZoneEmailRoutingRuleUpdateParams, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a specific routing rule.
func (r *ZoneEmailRoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *ZoneEmailRoutingRuleService) EmailRoutingRoutingRulesNewRoutingRule(ctx context.Context, zoneIdentifier string, body ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing routing rules.
func (r *ZoneEmailRoutingRuleService) EmailRoutingRoutingRulesListRoutingRules(ctx context.Context, zoneIdentifier string, query ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
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

type ZoneEmailRoutingRuleGetResponse struct {
	Errors   []ZoneEmailRoutingRuleGetResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleGetResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleGetResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleGetResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleGetResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleGetResponse]
type zoneEmailRoutingRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneEmailRoutingRuleGetResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleGetResponseError]
type zoneEmailRoutingRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneEmailRoutingRuleGetResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleGetResponseMessage]
type zoneEmailRoutingRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleGetResponseResult struct {
	// List actions patterns.
	Actions []ZoneEmailRoutingRuleGetResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleGetResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []ZoneEmailRoutingRuleGetResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                    `json:"tag"`
	JSON zoneEmailRoutingRuleGetResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleGetResponseResult]
type zoneEmailRoutingRuleGetResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type ZoneEmailRoutingRuleGetResponseResultAction struct {
	// Type of supported action.
	Type  ZoneEmailRoutingRuleGetResponseResultActionsType `json:"type,required"`
	Value []string                                         `json:"value,required"`
	JSON  zoneEmailRoutingRuleGetResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleGetResponseResultActionJSON contains the JSON metadata for
// the struct [ZoneEmailRoutingRuleGetResponseResultAction]
type zoneEmailRoutingRuleGetResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type ZoneEmailRoutingRuleGetResponseResultActionsType string

const (
	ZoneEmailRoutingRuleGetResponseResultActionsTypeDrop    ZoneEmailRoutingRuleGetResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleGetResponseResultActionsTypeForward ZoneEmailRoutingRuleGetResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleGetResponseResultActionsTypeWorker  ZoneEmailRoutingRuleGetResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleGetResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleGetResponseResultEnabledTrue  ZoneEmailRoutingRuleGetResponseResultEnabled = true
	ZoneEmailRoutingRuleGetResponseResultEnabledFalse ZoneEmailRoutingRuleGetResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleGetResponseResultMatcher struct {
	// Field for type matcher.
	Field ZoneEmailRoutingRuleGetResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type ZoneEmailRoutingRuleGetResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                           `json:"value,required"`
	JSON  zoneEmailRoutingRuleGetResponseResultMatcherJSON `json:"-"`
}

// zoneEmailRoutingRuleGetResponseResultMatcherJSON contains the JSON metadata for
// the struct [ZoneEmailRoutingRuleGetResponseResultMatcher]
type zoneEmailRoutingRuleGetResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleGetResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleGetResponseResultMatchersField string

const (
	ZoneEmailRoutingRuleGetResponseResultMatchersFieldTo ZoneEmailRoutingRuleGetResponseResultMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleGetResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleGetResponseResultMatchersTypeLiteral ZoneEmailRoutingRuleGetResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleGetResponseSuccess bool

const (
	ZoneEmailRoutingRuleGetResponseSuccessTrue ZoneEmailRoutingRuleGetResponseSuccess = true
)

type ZoneEmailRoutingRuleUpdateResponse struct {
	Errors   []ZoneEmailRoutingRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleUpdateResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleUpdateResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleUpdateResponse]
type zoneEmailRoutingRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneEmailRoutingRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleUpdateResponseError]
type zoneEmailRoutingRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneEmailRoutingRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleUpdateResponseMessage]
type zoneEmailRoutingRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleUpdateResponseResult struct {
	// List actions patterns.
	Actions []ZoneEmailRoutingRuleUpdateResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleUpdateResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []ZoneEmailRoutingRuleUpdateResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                       `json:"tag"`
	JSON zoneEmailRoutingRuleUpdateResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleUpdateResponseResult]
type zoneEmailRoutingRuleUpdateResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type ZoneEmailRoutingRuleUpdateResponseResultAction struct {
	// Type of supported action.
	Type  ZoneEmailRoutingRuleUpdateResponseResultActionsType `json:"type,required"`
	Value []string                                            `json:"value,required"`
	JSON  zoneEmailRoutingRuleUpdateResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseResultActionJSON contains the JSON metadata
// for the struct [ZoneEmailRoutingRuleUpdateResponseResultAction]
type zoneEmailRoutingRuleUpdateResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type ZoneEmailRoutingRuleUpdateResponseResultActionsType string

const (
	ZoneEmailRoutingRuleUpdateResponseResultActionsTypeDrop    ZoneEmailRoutingRuleUpdateResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleUpdateResponseResultActionsTypeForward ZoneEmailRoutingRuleUpdateResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleUpdateResponseResultActionsTypeWorker  ZoneEmailRoutingRuleUpdateResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleUpdateResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleUpdateResponseResultEnabledTrue  ZoneEmailRoutingRuleUpdateResponseResultEnabled = true
	ZoneEmailRoutingRuleUpdateResponseResultEnabledFalse ZoneEmailRoutingRuleUpdateResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleUpdateResponseResultMatcher struct {
	// Field for type matcher.
	Field ZoneEmailRoutingRuleUpdateResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type ZoneEmailRoutingRuleUpdateResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                              `json:"value,required"`
	JSON  zoneEmailRoutingRuleUpdateResponseResultMatcherJSON `json:"-"`
}

// zoneEmailRoutingRuleUpdateResponseResultMatcherJSON contains the JSON metadata
// for the struct [ZoneEmailRoutingRuleUpdateResponseResultMatcher]
type zoneEmailRoutingRuleUpdateResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleUpdateResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleUpdateResponseResultMatchersField string

const (
	ZoneEmailRoutingRuleUpdateResponseResultMatchersFieldTo ZoneEmailRoutingRuleUpdateResponseResultMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleUpdateResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleUpdateResponseResultMatchersTypeLiteral ZoneEmailRoutingRuleUpdateResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleUpdateResponseSuccess bool

const (
	ZoneEmailRoutingRuleUpdateResponseSuccessTrue ZoneEmailRoutingRuleUpdateResponseSuccess = true
)

type ZoneEmailRoutingRuleDeleteResponse struct {
	Errors   []ZoneEmailRoutingRuleDeleteResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleDeleteResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleDeleteResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleDeleteResponse]
type zoneEmailRoutingRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleDeleteResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneEmailRoutingRuleDeleteResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleDeleteResponseError]
type zoneEmailRoutingRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleDeleteResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneEmailRoutingRuleDeleteResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleDeleteResponseMessage]
type zoneEmailRoutingRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleDeleteResponseResult struct {
	// List actions patterns.
	Actions []ZoneEmailRoutingRuleDeleteResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleDeleteResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []ZoneEmailRoutingRuleDeleteResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                       `json:"tag"`
	JSON zoneEmailRoutingRuleDeleteResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneEmailRoutingRuleDeleteResponseResult]
type zoneEmailRoutingRuleDeleteResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type ZoneEmailRoutingRuleDeleteResponseResultAction struct {
	// Type of supported action.
	Type  ZoneEmailRoutingRuleDeleteResponseResultActionsType `json:"type,required"`
	Value []string                                            `json:"value,required"`
	JSON  zoneEmailRoutingRuleDeleteResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseResultActionJSON contains the JSON metadata
// for the struct [ZoneEmailRoutingRuleDeleteResponseResultAction]
type zoneEmailRoutingRuleDeleteResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type ZoneEmailRoutingRuleDeleteResponseResultActionsType string

const (
	ZoneEmailRoutingRuleDeleteResponseResultActionsTypeDrop    ZoneEmailRoutingRuleDeleteResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleDeleteResponseResultActionsTypeForward ZoneEmailRoutingRuleDeleteResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleDeleteResponseResultActionsTypeWorker  ZoneEmailRoutingRuleDeleteResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleDeleteResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleDeleteResponseResultEnabledTrue  ZoneEmailRoutingRuleDeleteResponseResultEnabled = true
	ZoneEmailRoutingRuleDeleteResponseResultEnabledFalse ZoneEmailRoutingRuleDeleteResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleDeleteResponseResultMatcher struct {
	// Field for type matcher.
	Field ZoneEmailRoutingRuleDeleteResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type ZoneEmailRoutingRuleDeleteResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                              `json:"value,required"`
	JSON  zoneEmailRoutingRuleDeleteResponseResultMatcherJSON `json:"-"`
}

// zoneEmailRoutingRuleDeleteResponseResultMatcherJSON contains the JSON metadata
// for the struct [ZoneEmailRoutingRuleDeleteResponseResultMatcher]
type zoneEmailRoutingRuleDeleteResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleDeleteResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleDeleteResponseResultMatchersField string

const (
	ZoneEmailRoutingRuleDeleteResponseResultMatchersFieldTo ZoneEmailRoutingRuleDeleteResponseResultMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleDeleteResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleDeleteResponseResultMatchersTypeLiteral ZoneEmailRoutingRuleDeleteResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleDeleteResponseSuccess bool

const (
	ZoneEmailRoutingRuleDeleteResponseSuccessTrue ZoneEmailRoutingRuleDeleteResponseSuccess = true
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse struct {
	Errors   []ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError   `json:"errors"`
	Messages []ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage `json:"messages"`
	Result   ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess `json:"success"`
	JSON    zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON    `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON contains
// the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult struct {
	// List actions patterns.
	Actions []ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                                                       `json:"tag"`
	JSON zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction struct {
	// Type of supported action.
	Type  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                            `json:"value,required"`
	JSON  zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeDrop    ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "drop"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeForward ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "forward"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeWorker  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher struct {
	// Field for type matcher.
	Field ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                              `json:"value,required"`
	JSON  zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersFieldTo ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersTypeLiteral ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccessTrue ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess = true
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse struct {
	// List actions patterns.
	Actions []ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction `json:"actions"`
	// Routing rule status.
	Enabled ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule identifier.
	Tag  string                                                                   `json:"tag"`
	JSON zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON struct {
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction struct {
	// Type of supported action.
	Type  ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType `json:"type,required"`
	Value []string                                                                        `json:"value,required"`
	JSON  zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON  `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeDrop    ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "drop"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeForward ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "forward"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeWorker  ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "worker"
)

// Routing rule status.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled = false
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher struct {
	// Field for type matcher.
	Field ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                          `json:"value,required"`
	JSON  zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON `json:"-"`
}

// zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON
// contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher]
type zoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersFieldTo ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersTypeLiteral ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType = "literal"
)

type ZoneEmailRoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]ZoneEmailRoutingRuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]ZoneEmailRoutingRuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[ZoneEmailRoutingRuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type ZoneEmailRoutingRuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[ZoneEmailRoutingRuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                    `json:"value,required"`
}

func (r ZoneEmailRoutingRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type ZoneEmailRoutingRuleUpdateParamsActionsType string

const (
	ZoneEmailRoutingRuleUpdateParamsActionsTypeDrop    ZoneEmailRoutingRuleUpdateParamsActionsType = "drop"
	ZoneEmailRoutingRuleUpdateParamsActionsTypeForward ZoneEmailRoutingRuleUpdateParamsActionsType = "forward"
	ZoneEmailRoutingRuleUpdateParamsActionsTypeWorker  ZoneEmailRoutingRuleUpdateParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[ZoneEmailRoutingRuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[ZoneEmailRoutingRuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r ZoneEmailRoutingRuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleUpdateParamsMatchersField string

const (
	ZoneEmailRoutingRuleUpdateParamsMatchersFieldTo ZoneEmailRoutingRuleUpdateParamsMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleUpdateParamsMatchersType string

const (
	ZoneEmailRoutingRuleUpdateParamsMatchersTypeLiteral ZoneEmailRoutingRuleUpdateParamsMatchersType = "literal"
)

// Routing rule status.
type ZoneEmailRoutingRuleUpdateParamsEnabled bool

const (
	ZoneEmailRoutingRuleUpdateParamsEnabledTrue  ZoneEmailRoutingRuleUpdateParamsEnabled = true
	ZoneEmailRoutingRuleUpdateParamsEnabledFalse ZoneEmailRoutingRuleUpdateParamsEnabled = false
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams struct {
	// List actions patterns.
	Actions param.Field[[]ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction struct {
	// Type of supported action.
	Type  param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                                                    `json:"value,required"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeDrop    ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "drop"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "forward"
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeWorker  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField = "to"
)

// Type of matcher.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType string

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType = "literal"
)

// Routing rule status.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = false
)

type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams]'s query
// parameters as `url.Values`.
func (r ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled bool

const (
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledTrue  ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = true
	ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledFalse ZoneEmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = false
)
