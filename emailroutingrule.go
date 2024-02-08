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
	var env EmailRoutingRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *EmailRoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body EmailRoutingRuleUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a specific routing rule.
func (r *EmailRoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *EmailRoutingRuleService) EmailRoutingRoutingRulesNewRoutingRule(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleParams, opts ...option.RequestOption) (res *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *EmailRoutingRuleService) EmailRoutingRoutingRulesListRoutingRules(ctx context.Context, zoneIdentifier string, query EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesParams, opts ...option.RequestOption) (res *[]EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRuleGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleGetResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                          `json:"tag"`
	JSON emailRoutingRuleGetResponseJSON `json:"-"`
}

// emailRoutingRuleGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponse]
type emailRoutingRuleGetResponseJSON struct {
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

func (r *EmailRoutingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleGetResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleGetResponseActionsType `json:"type,required"`
	Value []string                               `json:"value,required"`
	JSON  emailRoutingRuleGetResponseActionJSON  `json:"-"`
}

// emailRoutingRuleGetResponseActionJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponseAction]
type emailRoutingRuleGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleGetResponseActionsType string

const (
	EmailRoutingRuleGetResponseActionsTypeDrop    EmailRoutingRuleGetResponseActionsType = "drop"
	EmailRoutingRuleGetResponseActionsTypeForward EmailRoutingRuleGetResponseActionsType = "forward"
	EmailRoutingRuleGetResponseActionsTypeWorker  EmailRoutingRuleGetResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleGetResponseEnabled bool

const (
	EmailRoutingRuleGetResponseEnabledTrue  EmailRoutingRuleGetResponseEnabled = true
	EmailRoutingRuleGetResponseEnabledFalse EmailRoutingRuleGetResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleGetResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleGetResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleGetResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                 `json:"value,required"`
	JSON  emailRoutingRuleGetResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleGetResponseMatcherJSON contains the JSON metadata for the struct
// [EmailRoutingRuleGetResponseMatcher]
type emailRoutingRuleGetResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleGetResponseMatchersField string

const (
	EmailRoutingRuleGetResponseMatchersFieldTo EmailRoutingRuleGetResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleGetResponseMatchersType string

const (
	EmailRoutingRuleGetResponseMatchersTypeLiteral EmailRoutingRuleGetResponseMatchersType = "literal"
)

type EmailRoutingRuleUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleUpdateResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                             `json:"tag"`
	JSON emailRoutingRuleUpdateResponseJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleUpdateResponse]
type emailRoutingRuleUpdateResponseJSON struct {
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

func (r *EmailRoutingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleUpdateResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleUpdateResponseActionsType `json:"type,required"`
	Value []string                                  `json:"value,required"`
	JSON  emailRoutingRuleUpdateResponseActionJSON  `json:"-"`
}

// emailRoutingRuleUpdateResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseAction]
type emailRoutingRuleUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleUpdateResponseActionsType string

const (
	EmailRoutingRuleUpdateResponseActionsTypeDrop    EmailRoutingRuleUpdateResponseActionsType = "drop"
	EmailRoutingRuleUpdateResponseActionsTypeForward EmailRoutingRuleUpdateResponseActionsType = "forward"
	EmailRoutingRuleUpdateResponseActionsTypeWorker  EmailRoutingRuleUpdateResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleUpdateResponseEnabled bool

const (
	EmailRoutingRuleUpdateResponseEnabledTrue  EmailRoutingRuleUpdateResponseEnabled = true
	EmailRoutingRuleUpdateResponseEnabledFalse EmailRoutingRuleUpdateResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleUpdateResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleUpdateResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleUpdateResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                    `json:"value,required"`
	JSON  emailRoutingRuleUpdateResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseMatcher]
type emailRoutingRuleUpdateResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleUpdateResponseMatchersField string

const (
	EmailRoutingRuleUpdateResponseMatchersFieldTo EmailRoutingRuleUpdateResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleUpdateResponseMatchersType string

const (
	EmailRoutingRuleUpdateResponseMatchersTypeLiteral EmailRoutingRuleUpdateResponseMatchersType = "literal"
)

type EmailRoutingRuleDeleteResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleDeleteResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleDeleteResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleDeleteResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                             `json:"tag"`
	JSON emailRoutingRuleDeleteResponseJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleDeleteResponse]
type emailRoutingRuleDeleteResponseJSON struct {
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

func (r *EmailRoutingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleDeleteResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleDeleteResponseActionsType `json:"type,required"`
	Value []string                                  `json:"value,required"`
	JSON  emailRoutingRuleDeleteResponseActionJSON  `json:"-"`
}

// emailRoutingRuleDeleteResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseAction]
type emailRoutingRuleDeleteResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleDeleteResponseActionsType string

const (
	EmailRoutingRuleDeleteResponseActionsTypeDrop    EmailRoutingRuleDeleteResponseActionsType = "drop"
	EmailRoutingRuleDeleteResponseActionsTypeForward EmailRoutingRuleDeleteResponseActionsType = "forward"
	EmailRoutingRuleDeleteResponseActionsTypeWorker  EmailRoutingRuleDeleteResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleDeleteResponseEnabled bool

const (
	EmailRoutingRuleDeleteResponseEnabledTrue  EmailRoutingRuleDeleteResponseEnabled = true
	EmailRoutingRuleDeleteResponseEnabledFalse EmailRoutingRuleDeleteResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleDeleteResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleDeleteResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleDeleteResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                    `json:"value,required"`
	JSON  emailRoutingRuleDeleteResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseMatcher]
type emailRoutingRuleDeleteResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleDeleteResponseMatchersField string

const (
	EmailRoutingRuleDeleteResponseMatchersFieldTo EmailRoutingRuleDeleteResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleDeleteResponseMatchersType string

const (
	EmailRoutingRuleDeleteResponseMatchersTypeLiteral EmailRoutingRuleDeleteResponseMatchersType = "literal"
)

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                             `json:"tag"`
	JSON emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON contains the
// JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseJSON struct {
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

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsType `json:"type,required"`
	Value []string                                                                  `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionJSON  `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseAction]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsTypeDrop    EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsType = "drop"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsTypeForward EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsType = "forward"
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsTypeWorker  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                    `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcher]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersField string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersFieldTo EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersTypeLiteral EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseMatchersType = "literal"
)

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                                               `json:"tag"`
	JSON emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON contains
// the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseJSON struct {
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

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType `json:"type,required"`
	Value []string                                                                    `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON  `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeDrop    EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "drop"
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeForward EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "forward"
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsTypeWorker  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabledTrue  EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled = true
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabledFalse EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                                                      `json:"value,required"`
	JSON  emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersFieldTo EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType string

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersTypeLiteral EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseMatchersType = "literal"
)

type EmailRoutingRuleGetResponseEnvelope struct {
	Errors   []EmailRoutingRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRuleGetResponseEnvelope]
type emailRoutingRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    emailRoutingRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailRoutingRuleGetResponseEnvelopeErrors]
type emailRoutingRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    emailRoutingRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [EmailRoutingRuleGetResponseEnvelopeMessages]
type emailRoutingRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleGetResponseEnvelopeSuccessTrue EmailRoutingRuleGetResponseEnvelopeSuccess = true
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

type EmailRoutingRuleUpdateResponseEnvelope struct {
	Errors   []EmailRoutingRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRuleUpdateResponseEnvelope]
type emailRoutingRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    emailRoutingRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingRuleUpdateResponseEnvelopeErrors]
type emailRoutingRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    emailRoutingRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingRuleUpdateResponseEnvelopeMessages]
type emailRoutingRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleUpdateResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleUpdateResponseEnvelopeSuccessTrue EmailRoutingRuleUpdateResponseEnvelopeSuccess = true
)

type EmailRoutingRuleDeleteResponseEnvelope struct {
	Errors   []EmailRoutingRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRuleDeleteResponseEnvelope]
type emailRoutingRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    emailRoutingRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingRuleDeleteResponseEnvelopeErrors]
type emailRoutingRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    emailRoutingRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingRuleDeleteResponseEnvelopeMessages]
type emailRoutingRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleDeleteResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleDeleteResponseEnvelopeSuccessTrue EmailRoutingRuleDeleteResponseEnvelopeSuccess = true
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

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelope struct {
	Errors   []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelope]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrors struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrors]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessages struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessages]
type emailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeSuccessTrue EmailRoutingRuleEmailRoutingRoutingRulesNewRoutingRuleResponseEnvelopeSuccess = true
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

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelope struct {
	Errors   []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeJSON       `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelope]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrors struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrors]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessages struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessages]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeSuccessTrue EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeSuccess = true
)

type EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                `json:"total_count"`
	JSON       emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfo]
type emailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleEmailRoutingRoutingRulesListRoutingRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
