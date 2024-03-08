// File generated from our OpenAPI spec by Stainless.

package email_routing

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

// RoutingRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingRuleService] method
// instead.
type RoutingRuleService struct {
	Options   []option.RequestOption
	CatchAlls *RoutingRuleCatchAllService
}

// NewRoutingRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingRuleService(opts ...option.RequestOption) (r *RoutingRuleService) {
	r = &RoutingRuleService{}
	r.Options = opts
	r.CatchAlls = NewRoutingRuleCatchAllService(opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *RoutingRuleService) New(ctx context.Context, zoneIdentifier string, body RoutingRuleNewParams, opts ...option.RequestOption) (res *RoutingRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *RoutingRuleService) Update(ctx context.Context, zoneIdentifier string, ruleIdentifier string, body RoutingRuleUpdateParams, opts ...option.RequestOption) (res *RoutingRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *RoutingRuleService) List(ctx context.Context, zoneIdentifier string, query RoutingRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RoutingRuleListResponse], err error) {
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
func (r *RoutingRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query RoutingRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RoutingRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Delete a specific routing rule.
func (r *RoutingRuleService) Delete(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RoutingRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a specific routing rule already created.
func (r *RoutingRuleService) Get(ctx context.Context, zoneIdentifier string, ruleIdentifier string, opts ...option.RequestOption) (res *RoutingRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoutingRuleNewResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RoutingRuleNewResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleNewResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RoutingRuleNewResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                     `json:"tag"`
	JSON routingRuleNewResponseJSON `json:"-"`
}

// routingRuleNewResponseJSON contains the JSON metadata for the struct
// [RoutingRuleNewResponse]
type routingRuleNewResponseJSON struct {
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

func (r *RoutingRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RoutingRuleNewResponseAction struct {
	// Type of supported action.
	Type  RoutingRuleNewResponseActionsType `json:"type,required"`
	Value []string                          `json:"value,required"`
	JSON  routingRuleNewResponseActionJSON  `json:"-"`
}

// routingRuleNewResponseActionJSON contains the JSON metadata for the struct
// [RoutingRuleNewResponseAction]
type routingRuleNewResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RoutingRuleNewResponseActionsType string

const (
	RoutingRuleNewResponseActionsTypeDrop    RoutingRuleNewResponseActionsType = "drop"
	RoutingRuleNewResponseActionsTypeForward RoutingRuleNewResponseActionsType = "forward"
	RoutingRuleNewResponseActionsTypeWorker  RoutingRuleNewResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleNewResponseEnabled bool

const (
	RoutingRuleNewResponseEnabledTrue  RoutingRuleNewResponseEnabled = true
	RoutingRuleNewResponseEnabledFalse RoutingRuleNewResponseEnabled = false
)

// Matching pattern to forward your actions.
type RoutingRuleNewResponseMatcher struct {
	// Field for type matcher.
	Field RoutingRuleNewResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RoutingRuleNewResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                            `json:"value,required"`
	JSON  routingRuleNewResponseMatcherJSON `json:"-"`
}

// routingRuleNewResponseMatcherJSON contains the JSON metadata for the struct
// [RoutingRuleNewResponseMatcher]
type routingRuleNewResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleNewResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RoutingRuleNewResponseMatchersField string

const (
	RoutingRuleNewResponseMatchersFieldTo RoutingRuleNewResponseMatchersField = "to"
)

// Type of matcher.
type RoutingRuleNewResponseMatchersType string

const (
	RoutingRuleNewResponseMatchersTypeLiteral RoutingRuleNewResponseMatchersType = "literal"
)

type RoutingRuleUpdateResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RoutingRuleUpdateResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleUpdateResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RoutingRuleUpdateResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                        `json:"tag"`
	JSON routingRuleUpdateResponseJSON `json:"-"`
}

// routingRuleUpdateResponseJSON contains the JSON metadata for the struct
// [RoutingRuleUpdateResponse]
type routingRuleUpdateResponseJSON struct {
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

func (r *RoutingRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RoutingRuleUpdateResponseAction struct {
	// Type of supported action.
	Type  RoutingRuleUpdateResponseActionsType `json:"type,required"`
	Value []string                             `json:"value,required"`
	JSON  routingRuleUpdateResponseActionJSON  `json:"-"`
}

// routingRuleUpdateResponseActionJSON contains the JSON metadata for the struct
// [RoutingRuleUpdateResponseAction]
type routingRuleUpdateResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleUpdateResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RoutingRuleUpdateResponseActionsType string

const (
	RoutingRuleUpdateResponseActionsTypeDrop    RoutingRuleUpdateResponseActionsType = "drop"
	RoutingRuleUpdateResponseActionsTypeForward RoutingRuleUpdateResponseActionsType = "forward"
	RoutingRuleUpdateResponseActionsTypeWorker  RoutingRuleUpdateResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleUpdateResponseEnabled bool

const (
	RoutingRuleUpdateResponseEnabledTrue  RoutingRuleUpdateResponseEnabled = true
	RoutingRuleUpdateResponseEnabledFalse RoutingRuleUpdateResponseEnabled = false
)

// Matching pattern to forward your actions.
type RoutingRuleUpdateResponseMatcher struct {
	// Field for type matcher.
	Field RoutingRuleUpdateResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RoutingRuleUpdateResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                               `json:"value,required"`
	JSON  routingRuleUpdateResponseMatcherJSON `json:"-"`
}

// routingRuleUpdateResponseMatcherJSON contains the JSON metadata for the struct
// [RoutingRuleUpdateResponseMatcher]
type routingRuleUpdateResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleUpdateResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RoutingRuleUpdateResponseMatchersField string

const (
	RoutingRuleUpdateResponseMatchersFieldTo RoutingRuleUpdateResponseMatchersField = "to"
)

// Type of matcher.
type RoutingRuleUpdateResponseMatchersType string

const (
	RoutingRuleUpdateResponseMatchersTypeLiteral RoutingRuleUpdateResponseMatchersType = "literal"
)

type RoutingRuleListResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RoutingRuleListResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleListResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RoutingRuleListResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                      `json:"tag"`
	JSON routingRuleListResponseJSON `json:"-"`
}

// routingRuleListResponseJSON contains the JSON metadata for the struct
// [RoutingRuleListResponse]
type routingRuleListResponseJSON struct {
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

func (r *RoutingRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RoutingRuleListResponseAction struct {
	// Type of supported action.
	Type  RoutingRuleListResponseActionsType `json:"type,required"`
	Value []string                           `json:"value,required"`
	JSON  routingRuleListResponseActionJSON  `json:"-"`
}

// routingRuleListResponseActionJSON contains the JSON metadata for the struct
// [RoutingRuleListResponseAction]
type routingRuleListResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleListResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RoutingRuleListResponseActionsType string

const (
	RoutingRuleListResponseActionsTypeDrop    RoutingRuleListResponseActionsType = "drop"
	RoutingRuleListResponseActionsTypeForward RoutingRuleListResponseActionsType = "forward"
	RoutingRuleListResponseActionsTypeWorker  RoutingRuleListResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleListResponseEnabled bool

const (
	RoutingRuleListResponseEnabledTrue  RoutingRuleListResponseEnabled = true
	RoutingRuleListResponseEnabledFalse RoutingRuleListResponseEnabled = false
)

// Matching pattern to forward your actions.
type RoutingRuleListResponseMatcher struct {
	// Field for type matcher.
	Field RoutingRuleListResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RoutingRuleListResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                             `json:"value,required"`
	JSON  routingRuleListResponseMatcherJSON `json:"-"`
}

// routingRuleListResponseMatcherJSON contains the JSON metadata for the struct
// [RoutingRuleListResponseMatcher]
type routingRuleListResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleListResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleListResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RoutingRuleListResponseMatchersField string

const (
	RoutingRuleListResponseMatchersFieldTo RoutingRuleListResponseMatchersField = "to"
)

// Type of matcher.
type RoutingRuleListResponseMatchersType string

const (
	RoutingRuleListResponseMatchersTypeLiteral RoutingRuleListResponseMatchersType = "literal"
)

type RoutingRuleDeleteResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RoutingRuleDeleteResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleDeleteResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RoutingRuleDeleteResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                        `json:"tag"`
	JSON routingRuleDeleteResponseJSON `json:"-"`
}

// routingRuleDeleteResponseJSON contains the JSON metadata for the struct
// [RoutingRuleDeleteResponse]
type routingRuleDeleteResponseJSON struct {
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

func (r *RoutingRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RoutingRuleDeleteResponseAction struct {
	// Type of supported action.
	Type  RoutingRuleDeleteResponseActionsType `json:"type,required"`
	Value []string                             `json:"value,required"`
	JSON  routingRuleDeleteResponseActionJSON  `json:"-"`
}

// routingRuleDeleteResponseActionJSON contains the JSON metadata for the struct
// [RoutingRuleDeleteResponseAction]
type routingRuleDeleteResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RoutingRuleDeleteResponseActionsType string

const (
	RoutingRuleDeleteResponseActionsTypeDrop    RoutingRuleDeleteResponseActionsType = "drop"
	RoutingRuleDeleteResponseActionsTypeForward RoutingRuleDeleteResponseActionsType = "forward"
	RoutingRuleDeleteResponseActionsTypeWorker  RoutingRuleDeleteResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleDeleteResponseEnabled bool

const (
	RoutingRuleDeleteResponseEnabledTrue  RoutingRuleDeleteResponseEnabled = true
	RoutingRuleDeleteResponseEnabledFalse RoutingRuleDeleteResponseEnabled = false
)

// Matching pattern to forward your actions.
type RoutingRuleDeleteResponseMatcher struct {
	// Field for type matcher.
	Field RoutingRuleDeleteResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RoutingRuleDeleteResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                               `json:"value,required"`
	JSON  routingRuleDeleteResponseMatcherJSON `json:"-"`
}

// routingRuleDeleteResponseMatcherJSON contains the JSON metadata for the struct
// [RoutingRuleDeleteResponseMatcher]
type routingRuleDeleteResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleDeleteResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RoutingRuleDeleteResponseMatchersField string

const (
	RoutingRuleDeleteResponseMatchersFieldTo RoutingRuleDeleteResponseMatchersField = "to"
)

// Type of matcher.
type RoutingRuleDeleteResponseMatchersType string

const (
	RoutingRuleDeleteResponseMatchersTypeLiteral RoutingRuleDeleteResponseMatchersType = "literal"
)

type RoutingRuleGetResponse struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []RoutingRuleGetResponseAction `json:"actions"`
	// Routing rule status.
	Enabled RoutingRuleGetResponseEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []RoutingRuleGetResponseMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string                     `json:"tag"`
	JSON routingRuleGetResponseJSON `json:"-"`
}

// routingRuleGetResponseJSON contains the JSON metadata for the struct
// [RoutingRuleGetResponse]
type routingRuleGetResponseJSON struct {
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

func (r *RoutingRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// Actions pattern.
type RoutingRuleGetResponseAction struct {
	// Type of supported action.
	Type  RoutingRuleGetResponseActionsType `json:"type,required"`
	Value []string                          `json:"value,required"`
	JSON  routingRuleGetResponseActionJSON  `json:"-"`
}

// routingRuleGetResponseActionJSON contains the JSON metadata for the struct
// [RoutingRuleGetResponseAction]
type routingRuleGetResponseActionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleGetResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseActionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type RoutingRuleGetResponseActionsType string

const (
	RoutingRuleGetResponseActionsTypeDrop    RoutingRuleGetResponseActionsType = "drop"
	RoutingRuleGetResponseActionsTypeForward RoutingRuleGetResponseActionsType = "forward"
	RoutingRuleGetResponseActionsTypeWorker  RoutingRuleGetResponseActionsType = "worker"
)

// Routing rule status.
type RoutingRuleGetResponseEnabled bool

const (
	RoutingRuleGetResponseEnabledTrue  RoutingRuleGetResponseEnabled = true
	RoutingRuleGetResponseEnabledFalse RoutingRuleGetResponseEnabled = false
)

// Matching pattern to forward your actions.
type RoutingRuleGetResponseMatcher struct {
	// Field for type matcher.
	Field RoutingRuleGetResponseMatchersField `json:"field,required"`
	// Type of matcher.
	Type RoutingRuleGetResponseMatchersType `json:"type,required"`
	// Value for matcher.
	Value string                            `json:"value,required"`
	JSON  routingRuleGetResponseMatcherJSON `json:"-"`
}

// routingRuleGetResponseMatcherJSON contains the JSON metadata for the struct
// [RoutingRuleGetResponseMatcher]
type routingRuleGetResponseMatcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleGetResponseMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseMatcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type RoutingRuleGetResponseMatchersField string

const (
	RoutingRuleGetResponseMatchersFieldTo RoutingRuleGetResponseMatchersField = "to"
)

// Type of matcher.
type RoutingRuleGetResponseMatchersType string

const (
	RoutingRuleGetResponseMatchersTypeLiteral RoutingRuleGetResponseMatchersType = "literal"
)

type RoutingRuleNewParams struct {
	// List actions patterns.
	Actions param.Field[[]RoutingRuleNewParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]RoutingRuleNewParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RoutingRuleNewParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RoutingRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type RoutingRuleNewParamsAction struct {
	// Type of supported action.
	Type  param.Field[RoutingRuleNewParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                        `json:"value,required"`
}

func (r RoutingRuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type RoutingRuleNewParamsActionsType string

const (
	RoutingRuleNewParamsActionsTypeDrop    RoutingRuleNewParamsActionsType = "drop"
	RoutingRuleNewParamsActionsTypeForward RoutingRuleNewParamsActionsType = "forward"
	RoutingRuleNewParamsActionsTypeWorker  RoutingRuleNewParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type RoutingRuleNewParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[RoutingRuleNewParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[RoutingRuleNewParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r RoutingRuleNewParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type RoutingRuleNewParamsMatchersField string

const (
	RoutingRuleNewParamsMatchersFieldTo RoutingRuleNewParamsMatchersField = "to"
)

// Type of matcher.
type RoutingRuleNewParamsMatchersType string

const (
	RoutingRuleNewParamsMatchersTypeLiteral RoutingRuleNewParamsMatchersType = "literal"
)

// Routing rule status.
type RoutingRuleNewParamsEnabled bool

const (
	RoutingRuleNewParamsEnabledTrue  RoutingRuleNewParamsEnabled = true
	RoutingRuleNewParamsEnabledFalse RoutingRuleNewParamsEnabled = false
)

type RoutingRuleNewResponseEnvelope struct {
	Errors   []RoutingRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleNewResponseEnvelopeJSON    `json:"-"`
}

// routingRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingRuleNewResponseEnvelope]
type routingRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    routingRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingRuleNewResponseEnvelopeErrors]
type routingRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    routingRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingRuleNewResponseEnvelopeMessages]
type routingRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleNewResponseEnvelopeSuccess bool

const (
	RoutingRuleNewResponseEnvelopeSuccessTrue RoutingRuleNewResponseEnvelopeSuccess = true
)

type RoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]RoutingRuleUpdateParamsAction] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]RoutingRuleUpdateParamsMatcher] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RoutingRuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Actions pattern.
type RoutingRuleUpdateParamsAction struct {
	// Type of supported action.
	Type  param.Field[RoutingRuleUpdateParamsActionsType] `json:"type,required"`
	Value param.Field[[]string]                           `json:"value,required"`
}

func (r RoutingRuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of supported action.
type RoutingRuleUpdateParamsActionsType string

const (
	RoutingRuleUpdateParamsActionsTypeDrop    RoutingRuleUpdateParamsActionsType = "drop"
	RoutingRuleUpdateParamsActionsTypeForward RoutingRuleUpdateParamsActionsType = "forward"
	RoutingRuleUpdateParamsActionsTypeWorker  RoutingRuleUpdateParamsActionsType = "worker"
)

// Matching pattern to forward your actions.
type RoutingRuleUpdateParamsMatcher struct {
	// Field for type matcher.
	Field param.Field[RoutingRuleUpdateParamsMatchersField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[RoutingRuleUpdateParamsMatchersType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r RoutingRuleUpdateParamsMatcher) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Field for type matcher.
type RoutingRuleUpdateParamsMatchersField string

const (
	RoutingRuleUpdateParamsMatchersFieldTo RoutingRuleUpdateParamsMatchersField = "to"
)

// Type of matcher.
type RoutingRuleUpdateParamsMatchersType string

const (
	RoutingRuleUpdateParamsMatchersTypeLiteral RoutingRuleUpdateParamsMatchersType = "literal"
)

// Routing rule status.
type RoutingRuleUpdateParamsEnabled bool

const (
	RoutingRuleUpdateParamsEnabledTrue  RoutingRuleUpdateParamsEnabled = true
	RoutingRuleUpdateParamsEnabledFalse RoutingRuleUpdateParamsEnabled = false
)

type RoutingRuleUpdateResponseEnvelope struct {
	Errors   []RoutingRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// routingRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingRuleUpdateResponseEnvelope]
type routingRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    routingRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingRuleUpdateResponseEnvelopeErrors]
type routingRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    routingRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingRuleUpdateResponseEnvelopeMessages]
type routingRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleUpdateResponseEnvelopeSuccess bool

const (
	RoutingRuleUpdateResponseEnvelopeSuccessTrue RoutingRuleUpdateResponseEnvelopeSuccess = true
)

type RoutingRuleListParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[RoutingRuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RoutingRuleListParams]'s query parameters as `url.Values`.
func (r RoutingRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type RoutingRuleListParamsEnabled bool

const (
	RoutingRuleListParamsEnabledTrue  RoutingRuleListParamsEnabled = true
	RoutingRuleListParamsEnabledFalse RoutingRuleListParamsEnabled = false
)

type RoutingRuleDeleteResponseEnvelope struct {
	Errors   []RoutingRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// routingRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingRuleDeleteResponseEnvelope]
type routingRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    routingRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingRuleDeleteResponseEnvelopeErrors]
type routingRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    routingRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingRuleDeleteResponseEnvelopeMessages]
type routingRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleDeleteResponseEnvelopeSuccess bool

const (
	RoutingRuleDeleteResponseEnvelopeSuccessTrue RoutingRuleDeleteResponseEnvelopeSuccess = true
)

type RoutingRuleGetResponseEnvelope struct {
	Errors   []RoutingRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingRuleGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingRuleGetResponseEnvelopeJSON    `json:"-"`
}

// routingRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingRuleGetResponseEnvelope]
type routingRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    routingRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routingRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingRuleGetResponseEnvelopeErrors]
type routingRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingRuleGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    routingRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routingRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingRuleGetResponseEnvelopeMessages]
type routingRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingRuleGetResponseEnvelopeSuccess bool

const (
	RoutingRuleGetResponseEnvelopeSuccessTrue RoutingRuleGetResponseEnvelopeSuccess = true
)
