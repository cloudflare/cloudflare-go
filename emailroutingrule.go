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

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *EmailRoutingRuleService) New(ctx context.Context, zoneIdentifier string, body EmailRoutingRuleNewParams, opts ...option.RequestOption) (res *EmailRoutingRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *EmailRoutingRuleService) List(ctx context.Context, zoneIdentifier string, query EmailRoutingRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[EmailRoutingRuleListResponse], err error) {
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

// Lists existing routing rules.
func (r *EmailRoutingRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query EmailRoutingRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[EmailRoutingRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
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
func (r *EmailRoutingRuleService) Replace(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body EmailRoutingRuleReplaceParams, opts ...option.RequestOption) (res *EmailRoutingRuleReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRuleReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRuleNewResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleNewResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleNewResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleNewResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                          `json:"tag"`
	JSON emailRoutingRuleNewResponseJSON `json:"-"`
}

// emailRoutingRuleNewResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleNewResponse]
type emailRoutingRuleNewResponseJSON struct {
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

func (r *EmailRoutingRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleNewResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleNewResponseActionsType `json:"type,required"`
	Value []string                               `json:"value,required"`
	JSON  emailRoutingRuleNewResponseActionJSON  `json:"-"`
}

// emailRoutingRuleNewResponseActionJSON contains the JSON metadata for the struct
// [EmailRoutingRuleNewResponseAction]
type emailRoutingRuleNewResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleNewResponseActionsType string

const (
	EmailRoutingRuleNewResponseActionsTypeDrop    EmailRoutingRuleNewResponseActionsType = "drop"
	EmailRoutingRuleNewResponseActionsTypeForward EmailRoutingRuleNewResponseActionsType = "forward"
	EmailRoutingRuleNewResponseActionsTypeWorker  EmailRoutingRuleNewResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleNewResponseEnabled bool

const (
	EmailRoutingRuleNewResponseEnabledTrue  EmailRoutingRuleNewResponseEnabled = true
	EmailRoutingRuleNewResponseEnabledFalse EmailRoutingRuleNewResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleNewResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleNewResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleNewResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                 `json:"value,required"`
	JSON  emailRoutingRuleNewResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleNewResponseMatcherJSON contains the JSON metadata for the struct
// [EmailRoutingRuleNewResponseMatcher]
type emailRoutingRuleNewResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleNewResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleNewResponseMatchersField string

const (
	EmailRoutingRuleNewResponseMatchersFieldTo EmailRoutingRuleNewResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleNewResponseMatchersType string

const (
	EmailRoutingRuleNewResponseMatchersTypeLiteral EmailRoutingRuleNewResponseMatchersType = "literal"
)

type EmailRoutingRuleListResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleListResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleListResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleListResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                           `json:"tag"`
	JSON emailRoutingRuleListResponseJSON `json:"-"`
}

// emailRoutingRuleListResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleListResponse]
type emailRoutingRuleListResponseJSON struct {
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

func (r *EmailRoutingRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleListResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleListResponseActionsType `json:"type,required"`
	Value []string                                `json:"value,required"`
	JSON  emailRoutingRuleListResponseActionJSON  `json:"-"`
}

// emailRoutingRuleListResponseActionJSON contains the JSON metadata for the struct
// [EmailRoutingRuleListResponseAction]
type emailRoutingRuleListResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleListResponseActionsType string

const (
	EmailRoutingRuleListResponseActionsTypeDrop    EmailRoutingRuleListResponseActionsType = "drop"
	EmailRoutingRuleListResponseActionsTypeForward EmailRoutingRuleListResponseActionsType = "forward"
	EmailRoutingRuleListResponseActionsTypeWorker  EmailRoutingRuleListResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleListResponseEnabled bool

const (
	EmailRoutingRuleListResponseEnabledTrue  EmailRoutingRuleListResponseEnabled = true
	EmailRoutingRuleListResponseEnabledFalse EmailRoutingRuleListResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleListResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleListResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleListResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                  `json:"value,required"`
	JSON  emailRoutingRuleListResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleListResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRuleListResponseMatcher]
type emailRoutingRuleListResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleListResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleListResponseMatchersField string

const (
	EmailRoutingRuleListResponseMatchersFieldTo EmailRoutingRuleListResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleListResponseMatchersType string

const (
	EmailRoutingRuleListResponseMatchersTypeLiteral EmailRoutingRuleListResponseMatchersType = "literal"
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

type EmailRoutingRuleReplaceResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRuleReplaceResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleReplaceResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRuleReplaceResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                              `json:"tag"`
	JSON emailRoutingRuleReplaceResponseJSON `json:"-"`
}

// emailRoutingRuleReplaceResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRuleReplaceResponse]
type emailRoutingRuleReplaceResponseJSON struct {
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

func (r *EmailRoutingRuleReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions pattern.
type EmailRoutingRuleReplaceResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRuleReplaceResponseActionsType `json:"type,required"`
	Value []string                                   `json:"value,required"`
	JSON  emailRoutingRuleReplaceResponseActionJSON  `json:"-"`
}

// emailRoutingRuleReplaceResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRuleReplaceResponseAction]
type emailRoutingRuleReplaceResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleReplaceResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of supported action.
type EmailRoutingRuleReplaceResponseActionsType string

const (
	EmailRoutingRuleReplaceResponseActionsTypeDrop    EmailRoutingRuleReplaceResponseActionsType = "drop"
	EmailRoutingRuleReplaceResponseActionsTypeForward EmailRoutingRuleReplaceResponseActionsType = "forward"
	EmailRoutingRuleReplaceResponseActionsTypeWorker  EmailRoutingRuleReplaceResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRuleReplaceResponseEnabled bool

const (
	EmailRoutingRuleReplaceResponseEnabledTrue  EmailRoutingRuleReplaceResponseEnabled = true
	EmailRoutingRuleReplaceResponseEnabledFalse EmailRoutingRuleReplaceResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRuleReplaceResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRuleReplaceResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRuleReplaceResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                     `json:"value,required"`
	JSON  emailRoutingRuleReplaceResponseMatcherJSON `json:"-"`
}

// emailRoutingRuleReplaceResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRuleReplaceResponseMatcher]
type emailRoutingRuleReplaceResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleReplaceResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Field for type matcher.
type EmailRoutingRuleReplaceResponseMatchersField string

const (
	EmailRoutingRuleReplaceResponseMatchersFieldTo EmailRoutingRuleReplaceResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleReplaceResponseMatchersType string

const (
	EmailRoutingRuleReplaceResponseMatchersTypeLiteral EmailRoutingRuleReplaceResponseMatchersType = "literal"
)

type EmailRoutingRuleNewParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRuleNewParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRuleNewParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleNewParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRuleNewParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRuleNewParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                             `json:"value,required"`
}

func (r EmailRoutingRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRuleNewParamsActionsType string

const (
	EmailRoutingRuleNewParamsActionsTypeDrop    EmailRoutingRuleNewParamsActionsType = "drop"
	EmailRoutingRuleNewParamsActionsTypeForward EmailRoutingRuleNewParamsActionsType = "forward"
	EmailRoutingRuleNewParamsActionsTypeWorker  EmailRoutingRuleNewParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRuleNewParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRuleNewParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRuleNewParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRuleNewParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRuleNewParamsMatchersField string

const (
	EmailRoutingRuleNewParamsMatchersFieldTo EmailRoutingRuleNewParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleNewParamsMatchersType string

const (
	EmailRoutingRuleNewParamsMatchersTypeLiteral EmailRoutingRuleNewParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRuleNewParamsEnabled bool

const (
	EmailRoutingRuleNewParamsEnabledTrue  EmailRoutingRuleNewParamsEnabled = true
	EmailRoutingRuleNewParamsEnabledFalse EmailRoutingRuleNewParamsEnabled = false
)

type EmailRoutingRuleNewResponseEnvelope struct {
	Errors   []EmailRoutingRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRuleNewResponseEnvelope]
type emailRoutingRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    emailRoutingRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [EmailRoutingRuleNewResponseEnvelopeErrors]
type emailRoutingRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    emailRoutingRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [EmailRoutingRuleNewResponseEnvelopeMessages]
type emailRoutingRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleNewResponseEnvelopeSuccessTrue EmailRoutingRuleNewResponseEnvelopeSuccess = true
)

type EmailRoutingRuleListParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[EmailRoutingRuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [EmailRoutingRuleListParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type EmailRoutingRuleListParamsEnabled bool

const (
	EmailRoutingRuleListParamsEnabledTrue  EmailRoutingRuleListParamsEnabled = true
	EmailRoutingRuleListParamsEnabledFalse EmailRoutingRuleListParamsEnabled = false
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

type EmailRoutingRuleReplaceParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRuleReplaceParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRuleReplaceParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRuleReplaceParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRuleReplaceParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRuleReplaceParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                 `json:"value,required"`
}

func (r EmailRoutingRuleReplaceParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRuleReplaceParamsActionsType string

const (
	EmailRoutingRuleReplaceParamsActionsTypeDrop    EmailRoutingRuleReplaceParamsActionsType = "drop"
	EmailRoutingRuleReplaceParamsActionsTypeForward EmailRoutingRuleReplaceParamsActionsType = "forward"
	EmailRoutingRuleReplaceParamsActionsTypeWorker  EmailRoutingRuleReplaceParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRuleReplaceParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRuleReplaceParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRuleReplaceParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRuleReplaceParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRuleReplaceParamsMatchersField string

const (
	EmailRoutingRuleReplaceParamsMatchersFieldTo EmailRoutingRuleReplaceParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRuleReplaceParamsMatchersType string

const (
	EmailRoutingRuleReplaceParamsMatchersTypeLiteral EmailRoutingRuleReplaceParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRuleReplaceParamsEnabled bool

const (
	EmailRoutingRuleReplaceParamsEnabledTrue  EmailRoutingRuleReplaceParamsEnabled = true
	EmailRoutingRuleReplaceParamsEnabledFalse EmailRoutingRuleReplaceParamsEnabled = false
)

type EmailRoutingRuleReplaceResponseEnvelope struct {
	Errors   []EmailRoutingRuleReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRuleReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRuleReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRuleReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRuleReplaceResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRuleReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [EmailRoutingRuleReplaceResponseEnvelope]
type emailRoutingRuleReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleReplaceResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    emailRoutingRuleReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRuleReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EmailRoutingRuleReplaceResponseEnvelopeErrors]
type emailRoutingRuleReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmailRoutingRuleReplaceResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    emailRoutingRuleReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRuleReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EmailRoutingRuleReplaceResponseEnvelopeMessages]
type emailRoutingRuleReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRuleReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmailRoutingRuleReplaceResponseEnvelopeSuccess bool

const (
	EmailRoutingRuleReplaceResponseEnvelopeSuccessTrue EmailRoutingRuleReplaceResponseEnvelopeSuccess = true
)
