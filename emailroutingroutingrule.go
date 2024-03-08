// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
)

// EmailRoutingRoutingRuleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingRuleService] method instead.
type EmailRoutingRoutingRuleService struct {
	Options   []option.RequestOption
	CatchAlls *EmailRoutingRoutingRuleCatchAllService
}

// NewEmailRoutingRoutingRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingRuleService(opts ...option.RequestOption) (r *EmailRoutingRoutingRuleService) {
	r = &EmailRoutingRoutingRuleService{}
	r.Options = opts
	r.CatchAlls = NewEmailRoutingRoutingRuleCatchAllService(opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *EmailRoutingRoutingRuleService) New(ctx context.Context, zoneIdentifier string, body EmailRoutingRoutingRuleNewParams, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *EmailRoutingRoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body EmailRoutingRoutingRuleUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *EmailRoutingRoutingRuleService) List(ctx context.Context, zoneIdentifier string, query EmailRoutingRoutingRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[EmailRoutingRoutingRuleListResponse], err error) {
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
func (r *EmailRoutingRoutingRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query EmailRoutingRoutingRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[EmailRoutingRoutingRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Delete a specific routing rule.
func (r *EmailRoutingRoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a specific routing rule already created.
func (r *EmailRoutingRoutingRuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingRuleNewResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRoutingRuleNewResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleNewResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRoutingRuleNewResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                 `json:"tag"`
	JSON emailRoutingRoutingRuleNewResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleNewResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRoutingRuleNewResponse]
type emailRoutingRoutingRuleNewResponseJSON struct {
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

func (r *EmailRoutingRoutingRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type EmailRoutingRoutingRuleNewResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRoutingRuleNewResponseActionsType `json:"type,required"`
	Value []string                                      `json:"value,required"`
	JSON  emailRoutingRoutingRuleNewResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleNewResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleNewResponseAction]
type emailRoutingRoutingRuleNewResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRoutingRoutingRuleNewResponseActionsType string

const (
	EmailRoutingRoutingRuleNewResponseActionsTypeDrop    EmailRoutingRoutingRuleNewResponseActionsType = "drop"
	EmailRoutingRoutingRuleNewResponseActionsTypeForward EmailRoutingRoutingRuleNewResponseActionsType = "forward"
	EmailRoutingRoutingRuleNewResponseActionsTypeWorker  EmailRoutingRoutingRuleNewResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleNewResponseEnabled bool

const (
	EmailRoutingRoutingRuleNewResponseEnabledTrue  EmailRoutingRoutingRuleNewResponseEnabled = true
	EmailRoutingRoutingRuleNewResponseEnabledFalse EmailRoutingRoutingRuleNewResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleNewResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRoutingRuleNewResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRoutingRuleNewResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                        `json:"value,required"`
	JSON  emailRoutingRoutingRuleNewResponseMatcherJSON `json:"-"`
}

// emailRoutingRoutingRuleNewResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleNewResponseMatcher]
type emailRoutingRoutingRuleNewResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleNewResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type EmailRoutingRoutingRuleNewResponseMatchersField string

const (
	EmailRoutingRoutingRuleNewResponseMatchersFieldTo EmailRoutingRoutingRuleNewResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleNewResponseMatchersType string

const (
	EmailRoutingRoutingRuleNewResponseMatchersTypeLiteral EmailRoutingRoutingRuleNewResponseMatchersType = "literal"
)

type EmailRoutingRoutingRuleUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRoutingRuleUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleUpdateResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRoutingRuleUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                    `json:"tag"`
	JSON emailRoutingRoutingRuleUpdateResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleUpdateResponse]
type emailRoutingRoutingRuleUpdateResponseJSON struct {
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

func (r *EmailRoutingRoutingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type EmailRoutingRoutingRuleUpdateResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRoutingRuleUpdateResponseActionsType `json:"type,required"`
	Value []string                                         `json:"value,required"`
	JSON  emailRoutingRoutingRuleUpdateResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseActionJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleUpdateResponseAction]
type emailRoutingRoutingRuleUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRoutingRoutingRuleUpdateResponseActionsType string

const (
	EmailRoutingRoutingRuleUpdateResponseActionsTypeDrop    EmailRoutingRoutingRuleUpdateResponseActionsType = "drop"
	EmailRoutingRoutingRuleUpdateResponseActionsTypeForward EmailRoutingRoutingRuleUpdateResponseActionsType = "forward"
	EmailRoutingRoutingRuleUpdateResponseActionsTypeWorker  EmailRoutingRoutingRuleUpdateResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleUpdateResponseEnabled bool

const (
	EmailRoutingRoutingRuleUpdateResponseEnabledTrue  EmailRoutingRoutingRuleUpdateResponseEnabled = true
	EmailRoutingRoutingRuleUpdateResponseEnabledFalse EmailRoutingRoutingRuleUpdateResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleUpdateResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRoutingRuleUpdateResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRoutingRuleUpdateResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                           `json:"value,required"`
	JSON  emailRoutingRoutingRuleUpdateResponseMatcherJSON `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleUpdateResponseMatcher]
type emailRoutingRoutingRuleUpdateResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type EmailRoutingRoutingRuleUpdateResponseMatchersField string

const (
	EmailRoutingRoutingRuleUpdateResponseMatchersFieldTo EmailRoutingRoutingRuleUpdateResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleUpdateResponseMatchersType string

const (
	EmailRoutingRoutingRuleUpdateResponseMatchersTypeLiteral EmailRoutingRoutingRuleUpdateResponseMatchersType = "literal"
)

type EmailRoutingRoutingRuleListResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRoutingRuleListResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleListResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRoutingRuleListResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                  `json:"tag"`
	JSON emailRoutingRoutingRuleListResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleListResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleListResponse]
type emailRoutingRoutingRuleListResponseJSON struct {
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

func (r *EmailRoutingRoutingRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type EmailRoutingRoutingRuleListResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRoutingRuleListResponseActionsType `json:"type,required"`
	Value []string                                       `json:"value,required"`
	JSON  emailRoutingRoutingRuleListResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleListResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleListResponseAction]
type emailRoutingRoutingRuleListResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleListResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRoutingRoutingRuleListResponseActionsType string

const (
	EmailRoutingRoutingRuleListResponseActionsTypeDrop    EmailRoutingRoutingRuleListResponseActionsType = "drop"
	EmailRoutingRoutingRuleListResponseActionsTypeForward EmailRoutingRoutingRuleListResponseActionsType = "forward"
	EmailRoutingRoutingRuleListResponseActionsTypeWorker  EmailRoutingRoutingRuleListResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleListResponseEnabled bool

const (
	EmailRoutingRoutingRuleListResponseEnabledTrue  EmailRoutingRoutingRuleListResponseEnabled = true
	EmailRoutingRoutingRuleListResponseEnabledFalse EmailRoutingRoutingRuleListResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleListResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRoutingRuleListResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRoutingRuleListResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                         `json:"value,required"`
	JSON  emailRoutingRoutingRuleListResponseMatcherJSON `json:"-"`
}

// emailRoutingRoutingRuleListResponseMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleListResponseMatcher]
type emailRoutingRoutingRuleListResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleListResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleListResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type EmailRoutingRoutingRuleListResponseMatchersField string

const (
	EmailRoutingRoutingRuleListResponseMatchersFieldTo EmailRoutingRoutingRuleListResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleListResponseMatchersType string

const (
	EmailRoutingRoutingRuleListResponseMatchersTypeLiteral EmailRoutingRoutingRuleListResponseMatchersType = "literal"
)

type EmailRoutingRoutingRuleDeleteResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRoutingRuleDeleteResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleDeleteResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRoutingRuleDeleteResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                    `json:"tag"`
	JSON emailRoutingRoutingRuleDeleteResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleDeleteResponse]
type emailRoutingRoutingRuleDeleteResponseJSON struct {
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

func (r *EmailRoutingRoutingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type EmailRoutingRoutingRuleDeleteResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRoutingRuleDeleteResponseActionsType `json:"type,required"`
	Value []string                                         `json:"value,required"`
	JSON  emailRoutingRoutingRuleDeleteResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseActionJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleDeleteResponseAction]
type emailRoutingRoutingRuleDeleteResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRoutingRoutingRuleDeleteResponseActionsType string

const (
	EmailRoutingRoutingRuleDeleteResponseActionsTypeDrop    EmailRoutingRoutingRuleDeleteResponseActionsType = "drop"
	EmailRoutingRoutingRuleDeleteResponseActionsTypeForward EmailRoutingRoutingRuleDeleteResponseActionsType = "forward"
	EmailRoutingRoutingRuleDeleteResponseActionsTypeWorker  EmailRoutingRoutingRuleDeleteResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleDeleteResponseEnabled bool

const (
	EmailRoutingRoutingRuleDeleteResponseEnabledTrue  EmailRoutingRoutingRuleDeleteResponseEnabled = true
	EmailRoutingRoutingRuleDeleteResponseEnabledFalse EmailRoutingRoutingRuleDeleteResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleDeleteResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRoutingRuleDeleteResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRoutingRuleDeleteResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                           `json:"value,required"`
	JSON  emailRoutingRoutingRuleDeleteResponseMatcherJSON `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseMatcherJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleDeleteResponseMatcher]
type emailRoutingRoutingRuleDeleteResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleDeleteResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type EmailRoutingRoutingRuleDeleteResponseMatchersField string

const (
	EmailRoutingRoutingRuleDeleteResponseMatchersFieldTo EmailRoutingRoutingRuleDeleteResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleDeleteResponseMatchersType string

const (
	EmailRoutingRoutingRuleDeleteResponseMatchersTypeLiteral EmailRoutingRoutingRuleDeleteResponseMatchersType = "literal"
)

type EmailRoutingRoutingRuleGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRoutingRoutingRuleGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRoutingRuleGetResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRoutingRoutingRuleGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                                 `json:"tag"`
	JSON emailRoutingRoutingRuleGetResponseJSON `json:"-"`
}

// emailRoutingRoutingRuleGetResponseJSON contains the JSON metadata for the struct
// [EmailRoutingRoutingRuleGetResponse]
type emailRoutingRoutingRuleGetResponseJSON struct {
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

func (r *EmailRoutingRoutingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type EmailRoutingRoutingRuleGetResponseAction struct {
	// Type of supported action.
	Type  EmailRoutingRoutingRuleGetResponseActionsType `json:"type,required"`
	Value []string                                      `json:"value,required"`
	JSON  emailRoutingRoutingRuleGetResponseActionJSON  `json:"-"`
}

// emailRoutingRoutingRuleGetResponseActionJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleGetResponseAction]
type emailRoutingRoutingRuleGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRoutingRoutingRuleGetResponseActionsType string

const (
	EmailRoutingRoutingRuleGetResponseActionsTypeDrop    EmailRoutingRoutingRuleGetResponseActionsType = "drop"
	EmailRoutingRoutingRuleGetResponseActionsTypeForward EmailRoutingRoutingRuleGetResponseActionsType = "forward"
	EmailRoutingRoutingRuleGetResponseActionsTypeWorker  EmailRoutingRoutingRuleGetResponseActionsType = "worker"
)

// Routing rule status.
type EmailRoutingRoutingRuleGetResponseEnabled bool

const (
	EmailRoutingRoutingRuleGetResponseEnabledTrue  EmailRoutingRoutingRuleGetResponseEnabled = true
	EmailRoutingRoutingRuleGetResponseEnabledFalse EmailRoutingRoutingRuleGetResponseEnabled = false
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleGetResponseMatcher struct {
	// Field for type matcher.
	Field EmailRoutingRoutingRuleGetResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type EmailRoutingRoutingRuleGetResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                                        `json:"value,required"`
	JSON  emailRoutingRoutingRuleGetResponseMatcherJSON `json:"-"`
}

// emailRoutingRoutingRuleGetResponseMatcherJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingRuleGetResponseMatcher]
type emailRoutingRoutingRuleGetResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type EmailRoutingRoutingRuleGetResponseMatchersField string

const (
	EmailRoutingRoutingRuleGetResponseMatchersFieldTo EmailRoutingRoutingRuleGetResponseMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleGetResponseMatchersType string

const (
	EmailRoutingRoutingRuleGetResponseMatchersTypeLiteral EmailRoutingRoutingRuleGetResponseMatchersType = "literal"
)

type EmailRoutingRoutingRuleNewParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRoutingRuleNewParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRoutingRuleNewParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRoutingRuleNewParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRoutingRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRoutingRuleNewParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRoutingRuleNewParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                    `json:"value,required"`
}

func (r EmailRoutingRoutingRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRoutingRuleNewParamsActionsType string

const (
	EmailRoutingRoutingRuleNewParamsActionsTypeDrop    EmailRoutingRoutingRuleNewParamsActionsType = "drop"
	EmailRoutingRoutingRuleNewParamsActionsTypeForward EmailRoutingRoutingRuleNewParamsActionsType = "forward"
	EmailRoutingRoutingRuleNewParamsActionsTypeWorker  EmailRoutingRoutingRuleNewParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleNewParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRoutingRuleNewParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRoutingRuleNewParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRoutingRuleNewParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRoutingRuleNewParamsMatchersField string

const (
	EmailRoutingRoutingRuleNewParamsMatchersFieldTo EmailRoutingRoutingRuleNewParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleNewParamsMatchersType string

const (
	EmailRoutingRoutingRuleNewParamsMatchersTypeLiteral EmailRoutingRoutingRuleNewParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRoutingRuleNewParamsEnabled bool

const (
	EmailRoutingRoutingRuleNewParamsEnabledTrue  EmailRoutingRoutingRuleNewParamsEnabled = true
	EmailRoutingRoutingRuleNewParamsEnabledFalse EmailRoutingRoutingRuleNewParamsEnabled = false
)

type EmailRoutingRoutingRuleNewResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingRuleNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleNewResponseEnvelope]
type emailRoutingRoutingRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleNewResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    emailRoutingRoutingRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingRuleNewResponseEnvelopeErrors]
type emailRoutingRoutingRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleNewResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    emailRoutingRoutingRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleNewResponseEnvelopeMessages]
type emailRoutingRoutingRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleNewResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleNewResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRoutingRoutingRuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRoutingRoutingRuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRoutingRoutingRuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r EmailRoutingRoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type EmailRoutingRoutingRuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[EmailRoutingRoutingRuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                                       `json:"value,required"`
}

func (r EmailRoutingRoutingRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type EmailRoutingRoutingRuleUpdateParamsActionsType string

const (
	EmailRoutingRoutingRuleUpdateParamsActionsTypeDrop    EmailRoutingRoutingRuleUpdateParamsActionsType = "drop"
	EmailRoutingRoutingRuleUpdateParamsActionsTypeForward EmailRoutingRoutingRuleUpdateParamsActionsType = "forward"
	EmailRoutingRoutingRuleUpdateParamsActionsTypeWorker  EmailRoutingRoutingRuleUpdateParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type EmailRoutingRoutingRuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[EmailRoutingRoutingRuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[EmailRoutingRoutingRuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r EmailRoutingRoutingRuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type EmailRoutingRoutingRuleUpdateParamsMatchersField string

const (
	EmailRoutingRoutingRuleUpdateParamsMatchersFieldTo EmailRoutingRoutingRuleUpdateParamsMatchersField = "to"
)

// Type of matcher.
type EmailRoutingRoutingRuleUpdateParamsMatchersType string

const (
	EmailRoutingRoutingRuleUpdateParamsMatchersTypeLiteral EmailRoutingRoutingRuleUpdateParamsMatchersType = "literal"
)

// Routing rule status.
type EmailRoutingRoutingRuleUpdateParamsEnabled bool

const (
	EmailRoutingRoutingRuleUpdateParamsEnabledTrue  EmailRoutingRoutingRuleUpdateParamsEnabled = true
	EmailRoutingRoutingRuleUpdateParamsEnabledFalse EmailRoutingRoutingRuleUpdateParamsEnabled = false
)

type EmailRoutingRoutingRuleUpdateResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingRuleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleUpdateResponseEnvelope]
type emailRoutingRoutingRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleUpdateResponseEnvelopeErrors]
type emailRoutingRoutingRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleUpdateResponseEnvelopeMessages]
type emailRoutingRoutingRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleUpdateResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleUpdateResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleUpdateResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingRuleListParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[EmailRoutingRoutingRuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [EmailRoutingRoutingRuleListParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingRoutingRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type EmailRoutingRoutingRuleListParamsEnabled bool

const (
	EmailRoutingRoutingRuleListParamsEnabledTrue  EmailRoutingRoutingRuleListParamsEnabled = true
	EmailRoutingRoutingRuleListParamsEnabledFalse EmailRoutingRoutingRuleListParamsEnabled = false
)

type EmailRoutingRoutingRuleDeleteResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingRuleDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleDeleteResponseEnvelope]
type emailRoutingRoutingRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleDeleteResponseEnvelopeErrors]
type emailRoutingRoutingRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleDeleteResponseEnvelopeMessages]
type emailRoutingRoutingRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleDeleteResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleDeleteResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleDeleteResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingRuleGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingRuleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingRuleGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingRuleGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingRuleGetResponseEnvelope]
type emailRoutingRoutingRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    emailRoutingRoutingRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingRuleGetResponseEnvelopeErrors]
type emailRoutingRoutingRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingRuleGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    emailRoutingRoutingRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingRuleGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingRuleGetResponseEnvelopeMessages]
type emailRoutingRoutingRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingRuleGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingRuleGetResponseEnvelopeSuccessTrue EmailRoutingRoutingRuleGetResponseEnvelopeSuccess = true
)
