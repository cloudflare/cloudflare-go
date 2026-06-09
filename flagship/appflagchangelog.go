// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// AppFlagChangelogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppFlagChangelogService] method instead.
type AppFlagChangelogService struct {
	Options []option.RequestOption
}

// NewAppFlagChangelogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppFlagChangelogService(opts ...option.RequestOption) (r *AppFlagChangelogService) {
	r = &AppFlagChangelogService{}
	r.Options = opts
	return
}

// Returns the audit history for a flag, newest first. Each entry includes the
// event type and full flag state after the change; `update` entries include a
// field-level diff. Capped at 200 entries per flag.
func (r *AppFlagChangelogService) List(ctx context.Context, appID string, flagKey string, params AppFlagChangelogListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[AppFlagChangelogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	if flagKey == "" {
		err = errors.New("missing required flag_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags/%s/changelog", params.AccountID, appID, flagKey)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns the audit history for a flag, newest first. Each entry includes the
// event type and full flag state after the change; `update` entries include a
// field-level diff. Capped at 200 entries per flag.
func (r *AppFlagChangelogService) ListAutoPaging(ctx context.Context, appID string, flagKey string, params AppFlagChangelogListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[AppFlagChangelogListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, appID, flagKey, params, opts...))
}

type AppFlagChangelogListResponse struct {
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfter].
	After   interface{}                       `json:"after" api:"required"`
	Event   AppFlagChangelogListResponseEvent `json:"event" api:"required"`
	FlagKey string                            `json:"flag_key" api:"required"`
	// This field can have the runtime type of
	// [map[string]AppFlagChangelogListResponseObjectDiff].
	Diff  interface{}                      `json:"diff"`
	JSON  appFlagChangelogListResponseJSON `json:"-"`
	union AppFlagChangelogListResponseUnion
}

// appFlagChangelogListResponseJSON contains the JSON metadata for the struct
// [AppFlagChangelogListResponse]
type appFlagChangelogListResponseJSON struct {
	After       apijson.Field
	Event       apijson.Field
	FlagKey     apijson.Field
	Diff        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r appFlagChangelogListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagChangelogListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AppFlagChangelogListResponseObject],
// [AppFlagChangelogListResponseObject], [AppFlagChangelogListResponseObject].
func (r AppFlagChangelogListResponse) AsUnion() AppFlagChangelogListResponseUnion {
	return r.union
}

// Union satisfied by [AppFlagChangelogListResponseObject],
// [AppFlagChangelogListResponseObject] or [AppFlagChangelogListResponseObject].
type AppFlagChangelogListResponseUnion interface {
	implementsAppFlagChangelogListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObject{}),
		},
	)
}

type AppFlagChangelogListResponseObject struct {
	After   AppFlagChangelogListResponseObjectAfter `json:"after" api:"required"`
	Event   AppFlagChangelogListResponseObjectEvent `json:"event" api:"required"`
	FlagKey string                                  `json:"flag_key" api:"required"`
	JSON    appFlagChangelogListResponseObjectJSON  `json:"-"`
}

// appFlagChangelogListResponseObjectJSON contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObject]
type appFlagChangelogListResponseObjectJSON struct {
	After       apijson.Field
	Event       apijson.Field
	FlagKey     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObject) implementsAppFlagChangelogListResponse() {}

type AppFlagChangelogListResponseObjectAfter struct {
	// Variation served when no rule matches or the flag is disabled. Must be a key in
	// `variations`.
	DefaultVariation string `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled bool `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key string `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules []AppFlagChangelogListResponseObjectAfterRule `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  map[string]AppFlagChangelogListResponseObjectAfterVariationsUnion `json:"variations" api:"required"`
	Description string                                                            `json:"description" api:"nullable"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type      AppFlagChangelogListResponseObjectAfterType `json:"type"`
	UpdatedAt string                                      `json:"updated_at"`
	UpdatedBy string                                      `json:"updated_by"`
	JSON      appFlagChangelogListResponseObjectAfterJSON `json:"-"`
}

// appFlagChangelogListResponseObjectAfterJSON contains the JSON metadata for the
// struct [AppFlagChangelogListResponseObjectAfter]
type appFlagChangelogListResponseObjectAfterJSON struct {
	DefaultVariation apijson.Field
	Enabled          apijson.Field
	Key              apijson.Field
	Rules            apijson.Field
	Variations       apijson.Field
	Description      apijson.Field
	Type             apijson.Field
	UpdatedAt        apijson.Field
	UpdatedBy        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterJSON) RawJSON() string {
	return r.raw
}

type AppFlagChangelogListResponseObjectAfterRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions []AppFlagChangelogListResponseObjectAfterRulesCondition `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation string                                              `json:"serve_variation" api:"required"`
	Rollout        AppFlagChangelogListResponseObjectAfterRulesRollout `json:"rollout"`
	JSON           appFlagChangelogListResponseObjectAfterRuleJSON     `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRuleJSON contains the JSON metadata for
// the struct [AppFlagChangelogListResponseObjectAfterRule]
type appFlagChangelogListResponseObjectAfterRuleJSON struct {
	Conditions     apijson.Field
	Priority       apijson.Field
	ServeVariation apijson.Field
	Rollout        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRuleJSON) RawJSON() string {
	return r.raw
}

type AppFlagChangelogListResponseObjectAfterRulesCondition struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsObjectClause].
	Clauses         interface{}                                                           `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of [interface{}].
	Value interface{}                                               `json:"value"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionJSON `json:"-"`
	union AppFlagChangelogListResponseObjectAfterRulesConditionsUnion
}

// appFlagChangelogListResponseObjectAfterRulesConditionJSON contains the JSON
// metadata for the struct [AppFlagChangelogListResponseObjectAfterRulesCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponseObjectAfterRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponseObjectAfterRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagChangelogListResponseObjectAfterRulesConditionsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagChangelogListResponseObjectAfterRulesConditionsObject],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsObject].
func (r AppFlagChangelogListResponseObjectAfterRulesCondition) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsObject] or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsObject].
type AppFlagChangelogListResponseObjectAfterRulesConditionsUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsObject{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsObject struct {
	Attribute string                                                               `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value interface{}                                                      `json:"value" api:"required"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionsObjectJSON `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsObjectJSON contains the
// JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsObject]
type appFlagChangelogListResponseObjectAfterRulesConditionsObjectJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsObject) implementsAppFlagChangelogListResponseObjectAfterRulesCondition() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "not_in"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage float64 `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute string                                                  `json:"attribute"`
	JSON      appFlagChangelogListResponseObjectAfterRulesRolloutJSON `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesRolloutJSON contains the JSON
// metadata for the struct [AppFlagChangelogListResponseObjectAfterRulesRollout]
type appFlagChangelogListResponseObjectAfterRulesRolloutJSON struct {
	Percentage  apijson.Field
	Attribute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesRollout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesRolloutJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppFlagChangelogListResponseObjectAfterVariationsMap] or
// [AppFlagChangelogListResponseObjectAfterVariationsArray].
type AppFlagChangelogListResponseObjectAfterVariationsUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterVariationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterVariationsMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterVariationsArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterVariationsMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterVariationsMap) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion() {
}

type AppFlagChangelogListResponseObjectAfterVariationsArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterVariationsArray) ImplementsAppFlagChangelogListResponseObjectAfterVariationsUnion() {
}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagChangelogListResponseObjectAfterType string

const (
	AppFlagChangelogListResponseObjectAfterTypeBoolean AppFlagChangelogListResponseObjectAfterType = "boolean"
	AppFlagChangelogListResponseObjectAfterTypeString  AppFlagChangelogListResponseObjectAfterType = "string"
	AppFlagChangelogListResponseObjectAfterTypeNumber  AppFlagChangelogListResponseObjectAfterType = "number"
	AppFlagChangelogListResponseObjectAfterTypeJson    AppFlagChangelogListResponseObjectAfterType = "json"
)

func (r AppFlagChangelogListResponseObjectAfterType) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterTypeBoolean, AppFlagChangelogListResponseObjectAfterTypeString, AppFlagChangelogListResponseObjectAfterTypeNumber, AppFlagChangelogListResponseObjectAfterTypeJson:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectEvent string

const (
	AppFlagChangelogListResponseObjectEventCreate AppFlagChangelogListResponseObjectEvent = "create"
)

func (r AppFlagChangelogListResponseObjectEvent) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectEventCreate:
		return true
	}
	return false
}

type AppFlagChangelogListResponseEvent string

const (
	AppFlagChangelogListResponseEventCreate AppFlagChangelogListResponseEvent = "create"
	AppFlagChangelogListResponseEventDelete AppFlagChangelogListResponseEvent = "delete"
	AppFlagChangelogListResponseEventUpdate AppFlagChangelogListResponseEvent = "update"
)

func (r AppFlagChangelogListResponseEvent) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseEventCreate, AppFlagChangelogListResponseEventDelete, AppFlagChangelogListResponseEventUpdate:
		return true
	}
	return false
}

type AppFlagChangelogListParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Max items to return (1–200).
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [AppFlagChangelogListParams]'s query parameters as
// `url.Values`.
func (r AppFlagChangelogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
