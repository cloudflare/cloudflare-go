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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// EmailRoutingRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewEmailRoutingRuleService] method
// instead.
type EmailRoutingRuleService struct {
	Options   []option.RequestOption
	CatchAlls *EmailRoutingRuleCatchAllService
}

// NewEmailRoutingRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingRuleService(opts ...option.RequestOption) (r *EmailRoutingRuleService) {
	r = &EmailRoutingRuleService{}
	r.Options = opts
	r.CatchAlls = NewEmailRoutingRuleCatchAllService(opts...)
	return
}

// Get information for a specific routing rule already created.
func (r *EmailRoutingRuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *EmailRoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body EmailRoutingRuleUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a specific routing rule.
func (r *EmailRoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *EmailRoutingRuleService) EmailRoutingRoutingRulesNewRoutingRule(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams, opts ...option.RequestOption) (res *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists existing routing rules.
func (r *EmailRoutingRuleService) EmailRoutingRoutingRulesListRoutingRules(ctx context.Context, zoneIdentifier string, query EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams, opts ...option.RequestOption) (res *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type EmailRoutingRuleGetResponse struct {
	Errors   []EmailRoutingRuleGetResponseError   `json:"errors"`
	Messages []EmailRoutingRuleGetResponseMessage `json:"messages"`
	Result   EmailRoutingRuleGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleGetResponseSuccess `json:"success"`
	JSON    emailRoutingRuleGetResponseJSON    `json:"-"`
}

// emailRoutingRuleGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponse]
type emailRoutingRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    emailRoutingRuleGetResponseErrorJSON `json:"-"`
}

// emailRoutingRuleGetResponseErrorJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponseError]
type emailRoutingRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    emailRoutingRuleGetResponseMessageJSON `json:"-"`
}

// emailRoutingRuleGetResponseMessageJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponseMessage]
type emailRoutingRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleGetResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleGetResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleGetResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleGetResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                `json:"tag"`
	JSON emailRoutingRuleGetResponseResultJSON `json:"-"`
}

// emailRoutingRuleGetResponseResultJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponseResult]
type emailRoutingRuleGetResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleGetResponseResultAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleGetResponseResultActionsType `json:"type,required"`
	Value []string                                     `json:"value,required"`
	JSON  emailRoutingRuleGetResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleGetResponseResultActionJSON contains the JSON metadata for the
// struct [EmailRoutingRuleGetResponseResultAction]
type emailRoutingRuleGetResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleGetResponseResultActionsType string

const (
	EmailRoutingRuleGetResponseResultActionsTypeDrop    EmailRoutingRuleGetResponseResultActionsType = "drop"
	EmailRoutingRuleGetResponseResultActionsTypeForward EmailRoutingRuleGetResponseResultActionsType = "forward"
	EmailRoutingRuleGetResponseResultActionsTypeWorker  EmailRoutingRuleGetResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleGetResponseResultEnabled bool

const (
	EmailRoutingRuleGetResponseResultEnabledTrue  EmailRoutingRuleGetResponseResultEnabled = true
	EmailRoutingRuleGetResponseResultEnabledFalse EmailRoutingRuleGetResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleGetResponseResultMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleGetResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleGetResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                       `json:"value,required"`
	JSON  emailRoutingRuleGetResponseResultMatcherJSON `json:"-"`
}

// emailRoutingRuleGetResponseResultMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRuleGetResponseResultMatcher]
type emailRoutingRuleGetResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleGetResponseResultMatchersField string

const (
	EmailRoutingRuleGetResponseResultMatchersFieldTo EmailRoutingRuleGetResponseResultMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleGetResponseResultMatchersType string

const (
	EmailRoutingRuleGetResponseResultMatchersTypeLiteral EmailRoutingRuleGetResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type EmailRoutingRuleGetResponseSuccess bool

const (
	EmailRoutingRuleGetResponseSuccessTrue EmailRoutingRuleGetResponseSuccess = true
)

type EmailRoutingRuleUpdateResponse struct {
	Errors   []EmailRoutingRuleUpdateResponseError   `json:"errors"`
	Messages []EmailRoutingRuleUpdateResponseMessage `json:"messages"`
	Result   EmailRoutingRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleUpdateResponseSuccess `json:"success"`
	JSON    emailRoutingRuleUpdateResponseJSON    `json:"-"`
}

// emailRoutingRuleUpdateResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleUpdateResponse]
type emailRoutingRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    emailRoutingRuleUpdateResponseErrorJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseError]
type emailRoutingRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    emailRoutingRuleUpdateResponseMessageJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseMessage]
type emailRoutingRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleUpdateResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleUpdateResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleUpdateResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleUpdateResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                   `json:"tag"`
	JSON emailRoutingRuleUpdateResponseResultJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseResult]
type emailRoutingRuleUpdateResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleUpdateResponseResultAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleUpdateResponseResultActionsType `json:"type,required"`
	Value []string                                        `json:"value,required"`
	JSON  emailRoutingRuleUpdateResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleUpdateResponseResultActionJSON contains the JSON metadata for
// the struct [EmailRoutingRuleUpdateResponseResultAction]
type emailRoutingRuleUpdateResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleUpdateResponseResultActionsType string

const (
	EmailRoutingRuleUpdateResponseResultActionsTypeDrop    EmailRoutingRuleUpdateResponseResultActionsType = "drop"
	EmailRoutingRuleUpdateResponseResultActionsTypeForward EmailRoutingRuleUpdateResponseResultActionsType = "forward"
	EmailRoutingRuleUpdateResponseResultActionsTypeWorker  EmailRoutingRuleUpdateResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleUpdateResponseResultEnabled bool

const (
	EmailRoutingRuleUpdateResponseResultEnabledTrue  EmailRoutingRuleUpdateResponseResultEnabled = true
	EmailRoutingRuleUpdateResponseResultEnabledFalse EmailRoutingRuleUpdateResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleUpdateResponseResultMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleUpdateResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleUpdateResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                          `json:"value,required"`
	JSON  emailRoutingRuleUpdateResponseResultMatcherJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseResultMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRuleUpdateResponseResultMatcher]
type emailRoutingRuleUpdateResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleUpdateResponseResultMatchersField string

const (
	EmailRoutingRuleUpdateResponseResultMatchersFieldTo EmailRoutingRuleUpdateResponseResultMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleUpdateResponseResultMatchersType string

const (
	EmailRoutingRuleUpdateResponseResultMatchersTypeLiteral EmailRoutingRuleUpdateResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type EmailRoutingRuleUpdateResponseSuccess bool

const (
	EmailRoutingRuleUpdateResponseSuccessTrue EmailRoutingRuleUpdateResponseSuccess = true
)

type EmailRoutingRuleDeleteResponse struct {
	Errors   []EmailRoutingRuleDeleteResponseError   `json:"errors"`
	Messages []EmailRoutingRuleDeleteResponseMessage `json:"messages"`
	Result   EmailRoutingRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleDeleteResponseSuccess `json:"success"`
	JSON    emailRoutingRuleDeleteResponseJSON    `json:"-"`
}

// emailRoutingRuleDeleteResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleDeleteResponse]
type emailRoutingRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    emailRoutingRuleDeleteResponseErrorJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseError]
type emailRoutingRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    emailRoutingRuleDeleteResponseMessageJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseMessage]
type emailRoutingRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleDeleteResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleDeleteResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleDeleteResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleDeleteResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                   `json:"tag"`
	JSON emailRoutingRuleDeleteResponseResultJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseResult]
type emailRoutingRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleDeleteResponseResultAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleDeleteResponseResultActionsType `json:"type,required"`
	Value []string                                        `json:"value,required"`
	JSON  emailRoutingRuleDeleteResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleDeleteResponseResultActionJSON contains the JSON metadata for
// the struct [EmailRoutingRuleDeleteResponseResultAction]
type emailRoutingRuleDeleteResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleDeleteResponseResultActionsType string

const (
	EmailRoutingRuleDeleteResponseResultActionsTypeDrop    EmailRoutingRuleDeleteResponseResultActionsType = "drop"
	EmailRoutingRuleDeleteResponseResultActionsTypeForward EmailRoutingRuleDeleteResponseResultActionsType = "forward"
	EmailRoutingRuleDeleteResponseResultActionsTypeWorker  EmailRoutingRuleDeleteResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleDeleteResponseResultEnabled bool

const (
	EmailRoutingRuleDeleteResponseResultEnabledTrue  EmailRoutingRuleDeleteResponseResultEnabled = true
	EmailRoutingRuleDeleteResponseResultEnabledFalse EmailRoutingRuleDeleteResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleDeleteResponseResultMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleDeleteResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleDeleteResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                          `json:"value,required"`
	JSON  emailRoutingRuleDeleteResponseResultMatcherJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseResultMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRuleDeleteResponseResultMatcher]
type emailRoutingRuleDeleteResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleDeleteResponseResultMatchersField string

const (
	EmailRoutingRuleDeleteResponseResultMatchersFieldTo EmailRoutingRuleDeleteResponseResultMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleDeleteResponseResultMatchersType string

const (
	EmailRoutingRuleDeleteResponseResultMatchersTypeLiteral EmailRoutingRuleDeleteResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type EmailRoutingRuleDeleteResponseSuccess bool

const (
	EmailRoutingRuleDeleteResponseSuccessTrue EmailRoutingRuleDeleteResponseSuccess = true
)

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse struct {
	Errors   []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError   `json:"errors"`
	Messages []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage `json:"messages"`
	Result   EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess `json:"success"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON    `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON contains the
// JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON contains
// the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                   `json:"tag"`
	JSON emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType `json:"type,required"`
	Value []string                                                                        `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeDrop    EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "drop"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeForward EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "forward"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsTypeWorker  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                          `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersFieldTo EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersTypeLiteral EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseResultMatchersType = "literal"
)

// Whether the API call was successful
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccessTrue EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseSuccess = true
)

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse struct {
	Errors     []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseError    `json:"errors"`
	Messages   []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessage  `json:"messages"`
	Result     []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResult   `json:"result"`
	ResultInfo EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseSuccess `json:"success"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON    `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseError struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseErrorJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseErrorJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseError]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessage struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessageJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessage]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResult struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                                     `json:"tag"`
	JSON emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResult]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsType `json:"type,required"`
	Value []string                                                                          `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionJSON  `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultAction]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsTypeDrop    EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsType = "drop"
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsTypeForward EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsType = "forward"
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsTypeWorker  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                            `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcherJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcher]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersField string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersFieldTo EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersTypeLiteral EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultMatchersType = "literal"
)

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfo struct {
	Count      interface{}                                                                    `json:"count"`
	Page       interface{}                                                                    `json:"page"`
	PerPage    interface{}                                                                    `json:"per_page"`
	TotalCount interface{}                                                                    `json:"total_count"`
	JSON       emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfoJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfo]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseSuccess bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseSuccessTrue EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseSuccess = true
)

type EmailRoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                `json:"value,required"`
}

func (r EmailRoutingRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRuleUpdateParamsActionsType string

const (
	EmailRoutingRuleUpdateParamsActionsTypeDrop    EmailRoutingRuleUpdateParamsActionsType = "drop"
	EmailRoutingRuleUpdateParamsActionsTypeForward EmailRoutingRuleUpdateParamsActionsType = "forward"
	EmailRoutingRuleUpdateParamsActionsTypeWorker  EmailRoutingRuleUpdateParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRuleUpdateParamsMatchersField string

const (
	EmailRoutingRuleUpdateParamsMatchersFieldTo EmailRoutingRuleUpdateParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleUpdateParamsMatchersType string

const (
	EmailRoutingRuleUpdateParamsMatchersTypeLiteral EmailRoutingRuleUpdateParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRuleUpdateParamsEnabled bool

const (
	EmailRoutingRuleUpdateParamsEnabledTrue  EmailRoutingRuleUpdateParamsEnabled = true
	EmailRoutingRuleUpdateParamsEnabledFalse EmailRoutingRuleUpdateParamsEnabled = false
)

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                                                `json:"value,required"`
}

func (r EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeDrop    EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "drop"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeForward EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "forward"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsTypeWorker  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersFieldTo EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersTypeLiteral EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParamsEnabled = false
)

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams]'s query
// parameters as `url.Values`.
func (r EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParamsEnabled = false
)
