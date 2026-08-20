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

// AppFlagService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppFlagService] method instead.
type AppFlagService struct {
	Options   []option.RequestOption
	Changelog *AppFlagChangelogService
}

// NewAppFlagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppFlagService(opts ...option.RequestOption) (r *AppFlagService) {
	r = &AppFlagService{}
	r.Options = opts
	r.Changelog = NewAppFlagChangelogService(opts...)
	return
}

// Creates a flag. Returns 409 if the key already exists. `type` is inferred from
// variation values and may be omitted.
func (r *AppFlagService) New(ctx context.Context, appID string, params AppFlagNewParams, opts ...option.RequestOption) (res *AppFlagNewResponse, err error) {
	var env AppFlagNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replaces the entire flag definition. Omitted fields are dropped, not preserved —
// read before writing. Each update appends a changelog entry.
func (r *AppFlagService) Update(ctx context.Context, appID string, flagKey string, params AppFlagUpdateParams, opts ...option.RequestOption) (res *AppFlagUpdateResponse, err error) {
	var env AppFlagUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
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
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags/%s", params.AccountID, appID, flagKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists an app's flags ordered by key. Pass `cursor` from `result_info` to page
// forward; a null cursor indicates the last page.
func (r *AppFlagService) List(ctx context.Context, appID string, params AppFlagListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[AppFlagListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags", params.AccountID, appID)
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

// Lists an app's flags ordered by key. Pass `cursor` from `result_info` to page
// forward; a null cursor indicates the last page.
func (r *AppFlagService) ListAutoPaging(ctx context.Context, appID string, params AppFlagListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[AppFlagListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, appID, params, opts...))
}

// Permanently deletes a flag. Subsequent evaluations fall back to the
// caller-supplied default. Cannot be undone.
func (r *AppFlagService) Delete(ctx context.Context, appID string, flagKey string, body AppFlagDeleteParams, opts ...option.RequestOption) (res *AppFlagDeleteResponse, err error) {
	var env AppFlagDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
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
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags/%s", body.AccountID, appID, flagKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the full flag definition including rules, variations, and audit fields.
func (r *AppFlagService) Get(ctx context.Context, appID string, flagKey string, query AppFlagGetParams, opts ...option.RequestOption) (res *AppFlagGetResponse, err error) {
	var env AppFlagGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
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
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/flags/%s", query.AccountID, appID, flagKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AppFlagNewResponse struct {
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
	Rules []AppFlagNewResponseRule `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  map[string]AppFlagNewResponseVariationsUnion `json:"variations" api:"required"`
	Description string                                       `json:"description" api:"nullable"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type      AppFlagNewResponseType `json:"type"`
	UpdatedAt string                 `json:"updated_at"`
	UpdatedBy string                 `json:"updated_by"`
	JSON      appFlagNewResponseJSON `json:"-"`
}

// appFlagNewResponseJSON contains the JSON metadata for the struct
// [AppFlagNewResponse]
type appFlagNewResponseJSON struct {
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

func (r *AppFlagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseJSON) RawJSON() string {
	return r.raw
}

type AppFlagNewResponseRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions []AppFlagNewResponseRulesCondition `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation string                         `json:"serve_variation" api:"required"`
	Rollout        AppFlagNewResponseRulesRollout `json:"rollout"`
	JSON           appFlagNewResponseRuleJSON     `json:"-"`
}

// appFlagNewResponseRuleJSON contains the JSON metadata for the struct
// [AppFlagNewResponseRule]
type appFlagNewResponseRuleJSON struct {
	Conditions     apijson.Field
	Priority       apijson.Field
	ServeVariation apijson.Field
	Rollout        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AppFlagNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

type AppFlagNewResponseRulesCondition struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagNewResponseRulesConditionsObjectClause].
	Clauses         interface{}                                      `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of [interface{}].
	Value interface{}                          `json:"value"`
	JSON  appFlagNewResponseRulesConditionJSON `json:"-"`
	union AppFlagNewResponseRulesConditionsUnion
}

// appFlagNewResponseRulesConditionJSON contains the JSON metadata for the struct
// [AppFlagNewResponseRulesCondition]
type appFlagNewResponseRulesConditionJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagNewResponseRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagNewResponseRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagNewResponseRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagNewResponseRulesConditionsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagNewResponseRulesConditionsObject],
// [AppFlagNewResponseRulesConditionsObject].
func (r AppFlagNewResponseRulesCondition) AsUnion() AppFlagNewResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagNewResponseRulesConditionsObject] or
// [AppFlagNewResponseRulesConditionsObject].
type AppFlagNewResponseRulesConditionsUnion interface {
	implementsAppFlagNewResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsObject{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsObject struct {
	Attribute string                                          `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsObjectOperator `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value interface{}                                 `json:"value" api:"required"`
	JSON  appFlagNewResponseRulesConditionsObjectJSON `json:"-"`
}

// appFlagNewResponseRulesConditionsObjectJSON contains the JSON metadata for the
// struct [AppFlagNewResponseRulesConditionsObject]
type appFlagNewResponseRulesConditionsObjectJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsObject) implementsAppFlagNewResponseRulesCondition() {}

type AppFlagNewResponseRulesConditionsObjectOperator string

const (
	AppFlagNewResponseRulesConditionsObjectOperatorEquals              AppFlagNewResponseRulesConditionsObjectOperator = "equals"
	AppFlagNewResponseRulesConditionsObjectOperatorNotEquals           AppFlagNewResponseRulesConditionsObjectOperator = "not_equals"
	AppFlagNewResponseRulesConditionsObjectOperatorGreaterThan         AppFlagNewResponseRulesConditionsObjectOperator = "greater_than"
	AppFlagNewResponseRulesConditionsObjectOperatorLessThan            AppFlagNewResponseRulesConditionsObjectOperator = "less_than"
	AppFlagNewResponseRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsObjectOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsObjectOperatorContains            AppFlagNewResponseRulesConditionsObjectOperator = "contains"
	AppFlagNewResponseRulesConditionsObjectOperatorStartsWith          AppFlagNewResponseRulesConditionsObjectOperator = "starts_with"
	AppFlagNewResponseRulesConditionsObjectOperatorEndsWith            AppFlagNewResponseRulesConditionsObjectOperator = "ends_with"
	AppFlagNewResponseRulesConditionsObjectOperatorIn                  AppFlagNewResponseRulesConditionsObjectOperator = "in"
	AppFlagNewResponseRulesConditionsObjectOperatorNotIn               AppFlagNewResponseRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagNewResponseRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsObjectOperatorEquals, AppFlagNewResponseRulesConditionsObjectOperatorNotEquals, AppFlagNewResponseRulesConditionsObjectOperatorGreaterThan, AppFlagNewResponseRulesConditionsObjectOperatorLessThan, AppFlagNewResponseRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsObjectOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsObjectOperatorContains, AppFlagNewResponseRulesConditionsObjectOperatorStartsWith, AppFlagNewResponseRulesConditionsObjectOperatorEndsWith, AppFlagNewResponseRulesConditionsObjectOperatorIn, AppFlagNewResponseRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsLogicalOperatorAnd AppFlagNewResponseRulesConditionsLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsLogicalOperatorOr  AppFlagNewResponseRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsLogicalOperatorAnd, AppFlagNewResponseRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsOperator string

const (
	AppFlagNewResponseRulesConditionsOperatorEquals              AppFlagNewResponseRulesConditionsOperator = "equals"
	AppFlagNewResponseRulesConditionsOperatorNotEquals           AppFlagNewResponseRulesConditionsOperator = "not_equals"
	AppFlagNewResponseRulesConditionsOperatorGreaterThan         AppFlagNewResponseRulesConditionsOperator = "greater_than"
	AppFlagNewResponseRulesConditionsOperatorLessThan            AppFlagNewResponseRulesConditionsOperator = "less_than"
	AppFlagNewResponseRulesConditionsOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsOperatorContains            AppFlagNewResponseRulesConditionsOperator = "contains"
	AppFlagNewResponseRulesConditionsOperatorStartsWith          AppFlagNewResponseRulesConditionsOperator = "starts_with"
	AppFlagNewResponseRulesConditionsOperatorEndsWith            AppFlagNewResponseRulesConditionsOperator = "ends_with"
	AppFlagNewResponseRulesConditionsOperatorIn                  AppFlagNewResponseRulesConditionsOperator = "in"
	AppFlagNewResponseRulesConditionsOperatorNotIn               AppFlagNewResponseRulesConditionsOperator = "not_in"
)

func (r AppFlagNewResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsOperatorEquals, AppFlagNewResponseRulesConditionsOperatorNotEquals, AppFlagNewResponseRulesConditionsOperatorGreaterThan, AppFlagNewResponseRulesConditionsOperatorLessThan, AppFlagNewResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsOperatorContains, AppFlagNewResponseRulesConditionsOperatorStartsWith, AppFlagNewResponseRulesConditionsOperatorEndsWith, AppFlagNewResponseRulesConditionsOperatorIn, AppFlagNewResponseRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagNewResponseRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage float64 `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute string                             `json:"attribute"`
	JSON      appFlagNewResponseRulesRolloutJSON `json:"-"`
}

// appFlagNewResponseRulesRolloutJSON contains the JSON metadata for the struct
// [AppFlagNewResponseRulesRollout]
type appFlagNewResponseRulesRolloutJSON struct {
	Percentage  apijson.Field
	Attribute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesRollout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesRolloutJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppFlagNewResponseVariationsMap] or
// [AppFlagNewResponseVariationsArray].
type AppFlagNewResponseVariationsUnion interface {
	ImplementsAppFlagNewResponseVariationsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseVariationsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseVariationsMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseVariationsArray{}),
		},
	)
}

type AppFlagNewResponseVariationsMap map[string]interface{}

func (r AppFlagNewResponseVariationsMap) ImplementsAppFlagNewResponseVariationsUnion() {}

type AppFlagNewResponseVariationsArray []interface{}

func (r AppFlagNewResponseVariationsArray) ImplementsAppFlagNewResponseVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagNewResponseType string

const (
	AppFlagNewResponseTypeBoolean AppFlagNewResponseType = "boolean"
	AppFlagNewResponseTypeString  AppFlagNewResponseType = "string"
	AppFlagNewResponseTypeNumber  AppFlagNewResponseType = "number"
	AppFlagNewResponseTypeJson    AppFlagNewResponseType = "json"
)

func (r AppFlagNewResponseType) IsKnown() bool {
	switch r {
	case AppFlagNewResponseTypeBoolean, AppFlagNewResponseTypeString, AppFlagNewResponseTypeNumber, AppFlagNewResponseTypeJson:
		return true
	}
	return false
}

type AppFlagUpdateResponse struct {
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
	Rules []AppFlagUpdateResponseRule `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  map[string]AppFlagUpdateResponseVariationsUnion `json:"variations" api:"required"`
	Description string                                          `json:"description" api:"nullable"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type      AppFlagUpdateResponseType `json:"type"`
	UpdatedAt string                    `json:"updated_at"`
	UpdatedBy string                    `json:"updated_by"`
	JSON      appFlagUpdateResponseJSON `json:"-"`
}

// appFlagUpdateResponseJSON contains the JSON metadata for the struct
// [AppFlagUpdateResponse]
type appFlagUpdateResponseJSON struct {
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

func (r *AppFlagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateResponseRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions []AppFlagUpdateResponseRulesCondition `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation string                            `json:"serve_variation" api:"required"`
	Rollout        AppFlagUpdateResponseRulesRollout `json:"rollout"`
	JSON           appFlagUpdateResponseRuleJSON     `json:"-"`
}

// appFlagUpdateResponseRuleJSON contains the JSON metadata for the struct
// [AppFlagUpdateResponseRule]
type appFlagUpdateResponseRuleJSON struct {
	Conditions     apijson.Field
	Priority       apijson.Field
	ServeVariation apijson.Field
	Rollout        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateResponseRulesCondition struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagUpdateResponseRulesConditionsObjectClause].
	Clauses         interface{}                                         `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of [interface{}].
	Value interface{}                             `json:"value"`
	JSON  appFlagUpdateResponseRulesConditionJSON `json:"-"`
	union AppFlagUpdateResponseRulesConditionsUnion
}

// appFlagUpdateResponseRulesConditionJSON contains the JSON metadata for the
// struct [AppFlagUpdateResponseRulesCondition]
type appFlagUpdateResponseRulesConditionJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagUpdateResponseRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagUpdateResponseRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagUpdateResponseRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagUpdateResponseRulesConditionsUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagUpdateResponseRulesConditionsObject],
// [AppFlagUpdateResponseRulesConditionsObject].
func (r AppFlagUpdateResponseRulesCondition) AsUnion() AppFlagUpdateResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagUpdateResponseRulesConditionsObject] or
// [AppFlagUpdateResponseRulesConditionsObject].
type AppFlagUpdateResponseRulesConditionsUnion interface {
	implementsAppFlagUpdateResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsObject{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsObject struct {
	Attribute string                                             `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsObjectOperator `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value interface{}                                    `json:"value" api:"required"`
	JSON  appFlagUpdateResponseRulesConditionsObjectJSON `json:"-"`
}

// appFlagUpdateResponseRulesConditionsObjectJSON contains the JSON metadata for
// the struct [AppFlagUpdateResponseRulesConditionsObject]
type appFlagUpdateResponseRulesConditionsObjectJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsObject) implementsAppFlagUpdateResponseRulesCondition() {}

type AppFlagUpdateResponseRulesConditionsObjectOperator string

const (
	AppFlagUpdateResponseRulesConditionsObjectOperatorEquals              AppFlagUpdateResponseRulesConditionsObjectOperator = "equals"
	AppFlagUpdateResponseRulesConditionsObjectOperatorNotEquals           AppFlagUpdateResponseRulesConditionsObjectOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsObjectOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsObjectOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsObjectOperatorLessThan            AppFlagUpdateResponseRulesConditionsObjectOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsObjectOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsObjectOperatorContains            AppFlagUpdateResponseRulesConditionsObjectOperator = "contains"
	AppFlagUpdateResponseRulesConditionsObjectOperatorStartsWith          AppFlagUpdateResponseRulesConditionsObjectOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsObjectOperatorEndsWith            AppFlagUpdateResponseRulesConditionsObjectOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsObjectOperatorIn                  AppFlagUpdateResponseRulesConditionsObjectOperator = "in"
	AppFlagUpdateResponseRulesConditionsObjectOperatorNotIn               AppFlagUpdateResponseRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagUpdateResponseRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsObjectOperatorEquals, AppFlagUpdateResponseRulesConditionsObjectOperatorNotEquals, AppFlagUpdateResponseRulesConditionsObjectOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsObjectOperatorLessThan, AppFlagUpdateResponseRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsObjectOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsObjectOperatorContains, AppFlagUpdateResponseRulesConditionsObjectOperatorStartsWith, AppFlagUpdateResponseRulesConditionsObjectOperatorEndsWith, AppFlagUpdateResponseRulesConditionsObjectOperatorIn, AppFlagUpdateResponseRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsOperator string

const (
	AppFlagUpdateResponseRulesConditionsOperatorEquals              AppFlagUpdateResponseRulesConditionsOperator = "equals"
	AppFlagUpdateResponseRulesConditionsOperatorNotEquals           AppFlagUpdateResponseRulesConditionsOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsOperatorLessThan            AppFlagUpdateResponseRulesConditionsOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsOperatorContains            AppFlagUpdateResponseRulesConditionsOperator = "contains"
	AppFlagUpdateResponseRulesConditionsOperatorStartsWith          AppFlagUpdateResponseRulesConditionsOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsOperatorEndsWith            AppFlagUpdateResponseRulesConditionsOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsOperatorIn                  AppFlagUpdateResponseRulesConditionsOperator = "in"
	AppFlagUpdateResponseRulesConditionsOperatorNotIn               AppFlagUpdateResponseRulesConditionsOperator = "not_in"
)

func (r AppFlagUpdateResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsOperatorEquals, AppFlagUpdateResponseRulesConditionsOperatorNotEquals, AppFlagUpdateResponseRulesConditionsOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsOperatorLessThan, AppFlagUpdateResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsOperatorContains, AppFlagUpdateResponseRulesConditionsOperatorStartsWith, AppFlagUpdateResponseRulesConditionsOperatorEndsWith, AppFlagUpdateResponseRulesConditionsOperatorIn, AppFlagUpdateResponseRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage float64 `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute string                                `json:"attribute"`
	JSON      appFlagUpdateResponseRulesRolloutJSON `json:"-"`
}

// appFlagUpdateResponseRulesRolloutJSON contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesRollout]
type appFlagUpdateResponseRulesRolloutJSON struct {
	Percentage  apijson.Field
	Attribute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesRollout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesRolloutJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppFlagUpdateResponseVariationsMap] or
// [AppFlagUpdateResponseVariationsArray].
type AppFlagUpdateResponseVariationsUnion interface {
	ImplementsAppFlagUpdateResponseVariationsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseVariationsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseVariationsMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseVariationsArray{}),
		},
	)
}

type AppFlagUpdateResponseVariationsMap map[string]interface{}

func (r AppFlagUpdateResponseVariationsMap) ImplementsAppFlagUpdateResponseVariationsUnion() {}

type AppFlagUpdateResponseVariationsArray []interface{}

func (r AppFlagUpdateResponseVariationsArray) ImplementsAppFlagUpdateResponseVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagUpdateResponseType string

const (
	AppFlagUpdateResponseTypeBoolean AppFlagUpdateResponseType = "boolean"
	AppFlagUpdateResponseTypeString  AppFlagUpdateResponseType = "string"
	AppFlagUpdateResponseTypeNumber  AppFlagUpdateResponseType = "number"
	AppFlagUpdateResponseTypeJson    AppFlagUpdateResponseType = "json"
)

func (r AppFlagUpdateResponseType) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseTypeBoolean, AppFlagUpdateResponseTypeString, AppFlagUpdateResponseTypeNumber, AppFlagUpdateResponseTypeJson:
		return true
	}
	return false
}

type AppFlagListResponse struct {
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
	Rules []AppFlagListResponseRule `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  map[string]AppFlagListResponseVariationsUnion `json:"variations" api:"required"`
	Description string                                        `json:"description" api:"nullable"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type      AppFlagListResponseType `json:"type"`
	UpdatedAt string                  `json:"updated_at"`
	UpdatedBy string                  `json:"updated_by"`
	JSON      appFlagListResponseJSON `json:"-"`
}

// appFlagListResponseJSON contains the JSON metadata for the struct
// [AppFlagListResponse]
type appFlagListResponseJSON struct {
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

func (r *AppFlagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseJSON) RawJSON() string {
	return r.raw
}

type AppFlagListResponseRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions []AppFlagListResponseRulesCondition `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation string                          `json:"serve_variation" api:"required"`
	Rollout        AppFlagListResponseRulesRollout `json:"rollout"`
	JSON           appFlagListResponseRuleJSON     `json:"-"`
}

// appFlagListResponseRuleJSON contains the JSON metadata for the struct
// [AppFlagListResponseRule]
type appFlagListResponseRuleJSON struct {
	Conditions     apijson.Field
	Priority       apijson.Field
	ServeVariation apijson.Field
	Rollout        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AppFlagListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRuleJSON) RawJSON() string {
	return r.raw
}

type AppFlagListResponseRulesCondition struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagListResponseRulesConditionsObjectClause].
	Clauses         interface{}                                       `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of [interface{}].
	Value interface{}                           `json:"value"`
	JSON  appFlagListResponseRulesConditionJSON `json:"-"`
	union AppFlagListResponseRulesConditionsUnion
}

// appFlagListResponseRulesConditionJSON contains the JSON metadata for the struct
// [AppFlagListResponseRulesCondition]
type appFlagListResponseRulesConditionJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagListResponseRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagListResponseRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagListResponseRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagListResponseRulesConditionsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagListResponseRulesConditionsObject],
// [AppFlagListResponseRulesConditionsObject].
func (r AppFlagListResponseRulesCondition) AsUnion() AppFlagListResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagListResponseRulesConditionsObject] or
// [AppFlagListResponseRulesConditionsObject].
type AppFlagListResponseRulesConditionsUnion interface {
	implementsAppFlagListResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsObject{}),
		},
	)
}

type AppFlagListResponseRulesConditionsObject struct {
	Attribute string                                           `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsObjectOperator `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value interface{}                                  `json:"value" api:"required"`
	JSON  appFlagListResponseRulesConditionsObjectJSON `json:"-"`
}

// appFlagListResponseRulesConditionsObjectJSON contains the JSON metadata for the
// struct [AppFlagListResponseRulesConditionsObject]
type appFlagListResponseRulesConditionsObjectJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsObject) implementsAppFlagListResponseRulesCondition() {}

type AppFlagListResponseRulesConditionsObjectOperator string

const (
	AppFlagListResponseRulesConditionsObjectOperatorEquals              AppFlagListResponseRulesConditionsObjectOperator = "equals"
	AppFlagListResponseRulesConditionsObjectOperatorNotEquals           AppFlagListResponseRulesConditionsObjectOperator = "not_equals"
	AppFlagListResponseRulesConditionsObjectOperatorGreaterThan         AppFlagListResponseRulesConditionsObjectOperator = "greater_than"
	AppFlagListResponseRulesConditionsObjectOperatorLessThan            AppFlagListResponseRulesConditionsObjectOperator = "less_than"
	AppFlagListResponseRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsObjectOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsObjectOperatorContains            AppFlagListResponseRulesConditionsObjectOperator = "contains"
	AppFlagListResponseRulesConditionsObjectOperatorStartsWith          AppFlagListResponseRulesConditionsObjectOperator = "starts_with"
	AppFlagListResponseRulesConditionsObjectOperatorEndsWith            AppFlagListResponseRulesConditionsObjectOperator = "ends_with"
	AppFlagListResponseRulesConditionsObjectOperatorIn                  AppFlagListResponseRulesConditionsObjectOperator = "in"
	AppFlagListResponseRulesConditionsObjectOperatorNotIn               AppFlagListResponseRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagListResponseRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsObjectOperatorEquals, AppFlagListResponseRulesConditionsObjectOperatorNotEquals, AppFlagListResponseRulesConditionsObjectOperatorGreaterThan, AppFlagListResponseRulesConditionsObjectOperatorLessThan, AppFlagListResponseRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsObjectOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsObjectOperatorContains, AppFlagListResponseRulesConditionsObjectOperatorStartsWith, AppFlagListResponseRulesConditionsObjectOperatorEndsWith, AppFlagListResponseRulesConditionsObjectOperatorIn, AppFlagListResponseRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsLogicalOperator string

const (
	AppFlagListResponseRulesConditionsLogicalOperatorAnd AppFlagListResponseRulesConditionsLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsLogicalOperatorOr  AppFlagListResponseRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsLogicalOperatorAnd, AppFlagListResponseRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsOperator string

const (
	AppFlagListResponseRulesConditionsOperatorEquals              AppFlagListResponseRulesConditionsOperator = "equals"
	AppFlagListResponseRulesConditionsOperatorNotEquals           AppFlagListResponseRulesConditionsOperator = "not_equals"
	AppFlagListResponseRulesConditionsOperatorGreaterThan         AppFlagListResponseRulesConditionsOperator = "greater_than"
	AppFlagListResponseRulesConditionsOperatorLessThan            AppFlagListResponseRulesConditionsOperator = "less_than"
	AppFlagListResponseRulesConditionsOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsOperatorContains            AppFlagListResponseRulesConditionsOperator = "contains"
	AppFlagListResponseRulesConditionsOperatorStartsWith          AppFlagListResponseRulesConditionsOperator = "starts_with"
	AppFlagListResponseRulesConditionsOperatorEndsWith            AppFlagListResponseRulesConditionsOperator = "ends_with"
	AppFlagListResponseRulesConditionsOperatorIn                  AppFlagListResponseRulesConditionsOperator = "in"
	AppFlagListResponseRulesConditionsOperatorNotIn               AppFlagListResponseRulesConditionsOperator = "not_in"
)

func (r AppFlagListResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsOperatorEquals, AppFlagListResponseRulesConditionsOperatorNotEquals, AppFlagListResponseRulesConditionsOperatorGreaterThan, AppFlagListResponseRulesConditionsOperatorLessThan, AppFlagListResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsOperatorContains, AppFlagListResponseRulesConditionsOperatorStartsWith, AppFlagListResponseRulesConditionsOperatorEndsWith, AppFlagListResponseRulesConditionsOperatorIn, AppFlagListResponseRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagListResponseRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage float64 `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute string                              `json:"attribute"`
	JSON      appFlagListResponseRulesRolloutJSON `json:"-"`
}

// appFlagListResponseRulesRolloutJSON contains the JSON metadata for the struct
// [AppFlagListResponseRulesRollout]
type appFlagListResponseRulesRolloutJSON struct {
	Percentage  apijson.Field
	Attribute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesRollout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesRolloutJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppFlagListResponseVariationsMap] or
// [AppFlagListResponseVariationsArray].
type AppFlagListResponseVariationsUnion interface {
	ImplementsAppFlagListResponseVariationsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseVariationsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseVariationsMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseVariationsArray{}),
		},
	)
}

type AppFlagListResponseVariationsMap map[string]interface{}

func (r AppFlagListResponseVariationsMap) ImplementsAppFlagListResponseVariationsUnion() {}

type AppFlagListResponseVariationsArray []interface{}

func (r AppFlagListResponseVariationsArray) ImplementsAppFlagListResponseVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagListResponseType string

const (
	AppFlagListResponseTypeBoolean AppFlagListResponseType = "boolean"
	AppFlagListResponseTypeString  AppFlagListResponseType = "string"
	AppFlagListResponseTypeNumber  AppFlagListResponseType = "number"
	AppFlagListResponseTypeJson    AppFlagListResponseType = "json"
)

func (r AppFlagListResponseType) IsKnown() bool {
	switch r {
	case AppFlagListResponseTypeBoolean, AppFlagListResponseTypeString, AppFlagListResponseTypeNumber, AppFlagListResponseTypeJson:
		return true
	}
	return false
}

type AppFlagDeleteResponse struct {
	Key  string                    `json:"key" api:"required"`
	JSON appFlagDeleteResponseJSON `json:"-"`
}

// appFlagDeleteResponseJSON contains the JSON metadata for the struct
// [AppFlagDeleteResponse]
type appFlagDeleteResponseJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponse struct {
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
	Rules []AppFlagGetResponseRule `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  map[string]AppFlagGetResponseVariationsUnion `json:"variations" api:"required"`
	Description string                                       `json:"description" api:"nullable"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type      AppFlagGetResponseType `json:"type"`
	UpdatedAt string                 `json:"updated_at"`
	UpdatedBy string                 `json:"updated_by"`
	JSON      appFlagGetResponseJSON `json:"-"`
}

// appFlagGetResponseJSON contains the JSON metadata for the struct
// [AppFlagGetResponse]
type appFlagGetResponseJSON struct {
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

func (r *AppFlagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponseRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions []AppFlagGetResponseRulesCondition `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation string                         `json:"serve_variation" api:"required"`
	Rollout        AppFlagGetResponseRulesRollout `json:"rollout"`
	JSON           appFlagGetResponseRuleJSON     `json:"-"`
}

// appFlagGetResponseRuleJSON contains the JSON metadata for the struct
// [AppFlagGetResponseRule]
type appFlagGetResponseRuleJSON struct {
	Conditions     apijson.Field
	Priority       apijson.Field
	ServeVariation apijson.Field
	Rollout        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AppFlagGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponseRulesCondition struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagGetResponseRulesConditionsObjectClause].
	Clauses         interface{}                                      `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of [interface{}].
	Value interface{}                          `json:"value"`
	JSON  appFlagGetResponseRulesConditionJSON `json:"-"`
	union AppFlagGetResponseRulesConditionsUnion
}

// appFlagGetResponseRulesConditionJSON contains the JSON metadata for the struct
// [AppFlagGetResponseRulesCondition]
type appFlagGetResponseRulesConditionJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagGetResponseRulesConditionJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagGetResponseRulesCondition) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagGetResponseRulesCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppFlagGetResponseRulesConditionsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagGetResponseRulesConditionsObject],
// [AppFlagGetResponseRulesConditionsObject].
func (r AppFlagGetResponseRulesCondition) AsUnion() AppFlagGetResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagGetResponseRulesConditionsObject] or
// [AppFlagGetResponseRulesConditionsObject].
type AppFlagGetResponseRulesConditionsUnion interface {
	implementsAppFlagGetResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsObject{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsObject struct {
	Attribute string                                          `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsObjectOperator `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value interface{}                                 `json:"value" api:"required"`
	JSON  appFlagGetResponseRulesConditionsObjectJSON `json:"-"`
}

// appFlagGetResponseRulesConditionsObjectJSON contains the JSON metadata for the
// struct [AppFlagGetResponseRulesConditionsObject]
type appFlagGetResponseRulesConditionsObjectJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsObjectJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsObject) implementsAppFlagGetResponseRulesCondition() {}

type AppFlagGetResponseRulesConditionsObjectOperator string

const (
	AppFlagGetResponseRulesConditionsObjectOperatorEquals              AppFlagGetResponseRulesConditionsObjectOperator = "equals"
	AppFlagGetResponseRulesConditionsObjectOperatorNotEquals           AppFlagGetResponseRulesConditionsObjectOperator = "not_equals"
	AppFlagGetResponseRulesConditionsObjectOperatorGreaterThan         AppFlagGetResponseRulesConditionsObjectOperator = "greater_than"
	AppFlagGetResponseRulesConditionsObjectOperatorLessThan            AppFlagGetResponseRulesConditionsObjectOperator = "less_than"
	AppFlagGetResponseRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsObjectOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsObjectOperatorContains            AppFlagGetResponseRulesConditionsObjectOperator = "contains"
	AppFlagGetResponseRulesConditionsObjectOperatorStartsWith          AppFlagGetResponseRulesConditionsObjectOperator = "starts_with"
	AppFlagGetResponseRulesConditionsObjectOperatorEndsWith            AppFlagGetResponseRulesConditionsObjectOperator = "ends_with"
	AppFlagGetResponseRulesConditionsObjectOperatorIn                  AppFlagGetResponseRulesConditionsObjectOperator = "in"
	AppFlagGetResponseRulesConditionsObjectOperatorNotIn               AppFlagGetResponseRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagGetResponseRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsObjectOperatorEquals, AppFlagGetResponseRulesConditionsObjectOperatorNotEquals, AppFlagGetResponseRulesConditionsObjectOperatorGreaterThan, AppFlagGetResponseRulesConditionsObjectOperatorLessThan, AppFlagGetResponseRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsObjectOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsObjectOperatorContains, AppFlagGetResponseRulesConditionsObjectOperatorStartsWith, AppFlagGetResponseRulesConditionsObjectOperatorEndsWith, AppFlagGetResponseRulesConditionsObjectOperatorIn, AppFlagGetResponseRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsLogicalOperatorAnd AppFlagGetResponseRulesConditionsLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsLogicalOperatorOr  AppFlagGetResponseRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsLogicalOperatorAnd, AppFlagGetResponseRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsOperator string

const (
	AppFlagGetResponseRulesConditionsOperatorEquals              AppFlagGetResponseRulesConditionsOperator = "equals"
	AppFlagGetResponseRulesConditionsOperatorNotEquals           AppFlagGetResponseRulesConditionsOperator = "not_equals"
	AppFlagGetResponseRulesConditionsOperatorGreaterThan         AppFlagGetResponseRulesConditionsOperator = "greater_than"
	AppFlagGetResponseRulesConditionsOperatorLessThan            AppFlagGetResponseRulesConditionsOperator = "less_than"
	AppFlagGetResponseRulesConditionsOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsOperatorContains            AppFlagGetResponseRulesConditionsOperator = "contains"
	AppFlagGetResponseRulesConditionsOperatorStartsWith          AppFlagGetResponseRulesConditionsOperator = "starts_with"
	AppFlagGetResponseRulesConditionsOperatorEndsWith            AppFlagGetResponseRulesConditionsOperator = "ends_with"
	AppFlagGetResponseRulesConditionsOperatorIn                  AppFlagGetResponseRulesConditionsOperator = "in"
	AppFlagGetResponseRulesConditionsOperatorNotIn               AppFlagGetResponseRulesConditionsOperator = "not_in"
)

func (r AppFlagGetResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsOperatorEquals, AppFlagGetResponseRulesConditionsOperatorNotEquals, AppFlagGetResponseRulesConditionsOperatorGreaterThan, AppFlagGetResponseRulesConditionsOperatorLessThan, AppFlagGetResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsOperatorContains, AppFlagGetResponseRulesConditionsOperatorStartsWith, AppFlagGetResponseRulesConditionsOperatorEndsWith, AppFlagGetResponseRulesConditionsOperatorIn, AppFlagGetResponseRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagGetResponseRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage float64 `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute string                             `json:"attribute"`
	JSON      appFlagGetResponseRulesRolloutJSON `json:"-"`
}

// appFlagGetResponseRulesRolloutJSON contains the JSON metadata for the struct
// [AppFlagGetResponseRulesRollout]
type appFlagGetResponseRulesRolloutJSON struct {
	Percentage  apijson.Field
	Attribute   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesRollout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesRolloutJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppFlagGetResponseVariationsMap] or
// [AppFlagGetResponseVariationsArray].
type AppFlagGetResponseVariationsUnion interface {
	ImplementsAppFlagGetResponseVariationsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseVariationsUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseVariationsMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseVariationsArray{}),
		},
	)
}

type AppFlagGetResponseVariationsMap map[string]interface{}

func (r AppFlagGetResponseVariationsMap) ImplementsAppFlagGetResponseVariationsUnion() {}

type AppFlagGetResponseVariationsArray []interface{}

func (r AppFlagGetResponseVariationsArray) ImplementsAppFlagGetResponseVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagGetResponseType string

const (
	AppFlagGetResponseTypeBoolean AppFlagGetResponseType = "boolean"
	AppFlagGetResponseTypeString  AppFlagGetResponseType = "string"
	AppFlagGetResponseTypeNumber  AppFlagGetResponseType = "number"
	AppFlagGetResponseTypeJson    AppFlagGetResponseType = "json"
)

func (r AppFlagGetResponseType) IsKnown() bool {
	switch r {
	case AppFlagGetResponseTypeBoolean, AppFlagGetResponseTypeString, AppFlagGetResponseTypeNumber, AppFlagGetResponseTypeJson:
		return true
	}
	return false
}

type AppFlagNewParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Variation served when no rule matches or the flag is disabled. Must be a key in
	// `variations`.
	DefaultVariation param.Field[string] `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key param.Field[string] `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules param.Field[[]AppFlagNewParamsRule] `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  param.Field[map[string]AppFlagNewParamsVariationsUnion] `json:"variations" api:"required"`
	Description param.Field[string]                                     `json:"description"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type param.Field[AppFlagNewParamsType] `json:"type"`
}

func (r AppFlagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagNewParamsRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions param.Field[[]AppFlagNewParamsRulesConditionUnion] `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority param.Field[int64] `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation param.Field[string]                       `json:"serve_variation" api:"required"`
	Rollout        param.Field[AppFlagNewParamsRulesRollout] `json:"rollout"`
}

func (r AppFlagNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagNewParamsRulesCondition struct {
	Attribute       param.Field[string]                                         `json:"attribute"`
	Clauses         param.Field[interface{}]                                    `json:"clauses"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagNewParamsRulesConditionsOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                    `json:"value"`
}

func (r AppFlagNewParamsRulesCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesCondition) implementsAppFlagNewParamsRulesConditionUnion() {}

// Satisfied by [flagship.AppFlagNewParamsRulesConditionsObject],
// [flagship.AppFlagNewParamsRulesConditionsObject],
// [AppFlagNewParamsRulesCondition].
type AppFlagNewParamsRulesConditionUnion interface {
	implementsAppFlagNewParamsRulesConditionUnion()
}

type AppFlagNewParamsRulesConditionsObject struct {
	Attribute param.Field[string]                                        `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsObjectOperator] `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value param.Field[interface{}] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsObject) implementsAppFlagNewParamsRulesConditionUnion() {}

type AppFlagNewParamsRulesConditionsObjectOperator string

const (
	AppFlagNewParamsRulesConditionsObjectOperatorEquals              AppFlagNewParamsRulesConditionsObjectOperator = "equals"
	AppFlagNewParamsRulesConditionsObjectOperatorNotEquals           AppFlagNewParamsRulesConditionsObjectOperator = "not_equals"
	AppFlagNewParamsRulesConditionsObjectOperatorGreaterThan         AppFlagNewParamsRulesConditionsObjectOperator = "greater_than"
	AppFlagNewParamsRulesConditionsObjectOperatorLessThan            AppFlagNewParamsRulesConditionsObjectOperator = "less_than"
	AppFlagNewParamsRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsObjectOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsObjectOperatorContains            AppFlagNewParamsRulesConditionsObjectOperator = "contains"
	AppFlagNewParamsRulesConditionsObjectOperatorStartsWith          AppFlagNewParamsRulesConditionsObjectOperator = "starts_with"
	AppFlagNewParamsRulesConditionsObjectOperatorEndsWith            AppFlagNewParamsRulesConditionsObjectOperator = "ends_with"
	AppFlagNewParamsRulesConditionsObjectOperatorIn                  AppFlagNewParamsRulesConditionsObjectOperator = "in"
	AppFlagNewParamsRulesConditionsObjectOperatorNotIn               AppFlagNewParamsRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagNewParamsRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsObjectOperatorEquals, AppFlagNewParamsRulesConditionsObjectOperatorNotEquals, AppFlagNewParamsRulesConditionsObjectOperatorGreaterThan, AppFlagNewParamsRulesConditionsObjectOperatorLessThan, AppFlagNewParamsRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsObjectOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsObjectOperatorContains, AppFlagNewParamsRulesConditionsObjectOperatorStartsWith, AppFlagNewParamsRulesConditionsObjectOperatorEndsWith, AppFlagNewParamsRulesConditionsObjectOperatorIn, AppFlagNewParamsRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsLogicalOperatorAnd AppFlagNewParamsRulesConditionsLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsLogicalOperatorOr  AppFlagNewParamsRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsLogicalOperatorAnd, AppFlagNewParamsRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsOperator string

const (
	AppFlagNewParamsRulesConditionsOperatorEquals              AppFlagNewParamsRulesConditionsOperator = "equals"
	AppFlagNewParamsRulesConditionsOperatorNotEquals           AppFlagNewParamsRulesConditionsOperator = "not_equals"
	AppFlagNewParamsRulesConditionsOperatorGreaterThan         AppFlagNewParamsRulesConditionsOperator = "greater_than"
	AppFlagNewParamsRulesConditionsOperatorLessThan            AppFlagNewParamsRulesConditionsOperator = "less_than"
	AppFlagNewParamsRulesConditionsOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsOperatorContains            AppFlagNewParamsRulesConditionsOperator = "contains"
	AppFlagNewParamsRulesConditionsOperatorStartsWith          AppFlagNewParamsRulesConditionsOperator = "starts_with"
	AppFlagNewParamsRulesConditionsOperatorEndsWith            AppFlagNewParamsRulesConditionsOperator = "ends_with"
	AppFlagNewParamsRulesConditionsOperatorIn                  AppFlagNewParamsRulesConditionsOperator = "in"
	AppFlagNewParamsRulesConditionsOperatorNotIn               AppFlagNewParamsRulesConditionsOperator = "not_in"
)

func (r AppFlagNewParamsRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsOperatorEquals, AppFlagNewParamsRulesConditionsOperatorNotEquals, AppFlagNewParamsRulesConditionsOperatorGreaterThan, AppFlagNewParamsRulesConditionsOperatorLessThan, AppFlagNewParamsRulesConditionsOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsOperatorContains, AppFlagNewParamsRulesConditionsOperatorStartsWith, AppFlagNewParamsRulesConditionsOperatorEndsWith, AppFlagNewParamsRulesConditionsOperatorIn, AppFlagNewParamsRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagNewParamsRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute param.Field[string] `json:"attribute"`
}

func (r AppFlagNewParamsRulesRollout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsVariationsMap],
// [flagship.AppFlagNewParamsVariationsArray].
type AppFlagNewParamsVariationsUnion interface {
	ImplementsAppFlagNewParamsVariationsUnion()
}

type AppFlagNewParamsVariationsMap map[string]interface{}

func (r AppFlagNewParamsVariationsMap) ImplementsAppFlagNewParamsVariationsUnion() {}

type AppFlagNewParamsVariationsArray []interface{}

func (r AppFlagNewParamsVariationsArray) ImplementsAppFlagNewParamsVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagNewParamsType string

const (
	AppFlagNewParamsTypeBoolean AppFlagNewParamsType = "boolean"
	AppFlagNewParamsTypeString  AppFlagNewParamsType = "string"
	AppFlagNewParamsTypeNumber  AppFlagNewParamsType = "number"
	AppFlagNewParamsTypeJson    AppFlagNewParamsType = "json"
)

func (r AppFlagNewParamsType) IsKnown() bool {
	switch r {
	case AppFlagNewParamsTypeBoolean, AppFlagNewParamsTypeString, AppFlagNewParamsTypeNumber, AppFlagNewParamsTypeJson:
		return true
	}
	return false
}

type AppFlagNewResponseEnvelope struct {
	Errors   []AppFlagNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppFlagNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppFlagNewResponse                   `json:"result" api:"required"`
	Success  bool                                 `json:"success" api:"required"`
	JSON     appFlagNewResponseEnvelopeJSON       `json:"-"`
}

// appFlagNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppFlagNewResponseEnvelope]
type appFlagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppFlagNewResponseEnvelopeErrors struct {
	Message string                               `json:"message" api:"required"`
	JSON    appFlagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// appFlagNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppFlagNewResponseEnvelopeErrors]
type appFlagNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppFlagNewResponseEnvelopeMessages struct {
	Message string                                 `json:"message" api:"required"`
	JSON    appFlagNewResponseEnvelopeMessagesJSON `json:"-"`
}

// appFlagNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppFlagNewResponseEnvelopeMessages]
type appFlagNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Variation served when no rule matches or the flag is disabled. Must be a key in
	// `variations`.
	DefaultVariation param.Field[string] `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key param.Field[string] `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules param.Field[[]AppFlagUpdateParamsRule] `json:"rules" api:"required"`
	// Map of variation name to value. All values must be the same type (boolean,
	// string, number, or JSON object/array). Each serialized value must be 10KB or
	// smaller.
	Variations  param.Field[map[string]AppFlagUpdateParamsVariationsUnion] `json:"variations" api:"required"`
	Description param.Field[string]                                        `json:"description"`
	// Value type of the flag's variations. Inferred from the variation values on
	// write, so it may be omitted in requests.
	Type param.Field[AppFlagUpdateParamsType] `json:"type"`
}

func (r AppFlagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagUpdateParamsRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions param.Field[[]AppFlagUpdateParamsRulesConditionUnion] `json:"conditions" api:"required"`
	// Evaluation order; lower numbers are evaluated first. Must be unique across the
	// flag's rules.
	Priority param.Field[int64] `json:"priority" api:"required"`
	// Variation served when this rule matches. Must be a key in `variations`.
	ServeVariation param.Field[string]                          `json:"serve_variation" api:"required"`
	Rollout        param.Field[AppFlagUpdateParamsRulesRollout] `json:"rollout"`
}

func (r AppFlagUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagUpdateParamsRulesCondition struct {
	Attribute       param.Field[string]                                            `json:"attribute"`
	Clauses         param.Field[interface{}]                                       `json:"clauses"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagUpdateParamsRulesConditionsOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                       `json:"value"`
}

func (r AppFlagUpdateParamsRulesCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesCondition) implementsAppFlagUpdateParamsRulesConditionUnion() {}

// Satisfied by [flagship.AppFlagUpdateParamsRulesConditionsObject],
// [flagship.AppFlagUpdateParamsRulesConditionsObject],
// [AppFlagUpdateParamsRulesCondition].
type AppFlagUpdateParamsRulesConditionUnion interface {
	implementsAppFlagUpdateParamsRulesConditionUnion()
}

type AppFlagUpdateParamsRulesConditionsObject struct {
	Attribute param.Field[string]                                           `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsObjectOperator] `json:"operator" api:"required"`
	// Value to compare against the context attribute. Must be an array for `in` and
	// `not_in`; numeric and ISO-8601 datetime strings are accepted by the ordering
	// operators.
	Value param.Field[interface{}] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsObject) implementsAppFlagUpdateParamsRulesConditionUnion() {
}

type AppFlagUpdateParamsRulesConditionsObjectOperator string

const (
	AppFlagUpdateParamsRulesConditionsObjectOperatorEquals              AppFlagUpdateParamsRulesConditionsObjectOperator = "equals"
	AppFlagUpdateParamsRulesConditionsObjectOperatorNotEquals           AppFlagUpdateParamsRulesConditionsObjectOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsObjectOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsObjectOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsObjectOperatorLessThan            AppFlagUpdateParamsRulesConditionsObjectOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsObjectOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsObjectOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsObjectOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsObjectOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsObjectOperatorContains            AppFlagUpdateParamsRulesConditionsObjectOperator = "contains"
	AppFlagUpdateParamsRulesConditionsObjectOperatorStartsWith          AppFlagUpdateParamsRulesConditionsObjectOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsObjectOperatorEndsWith            AppFlagUpdateParamsRulesConditionsObjectOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsObjectOperatorIn                  AppFlagUpdateParamsRulesConditionsObjectOperator = "in"
	AppFlagUpdateParamsRulesConditionsObjectOperatorNotIn               AppFlagUpdateParamsRulesConditionsObjectOperator = "not_in"
)

func (r AppFlagUpdateParamsRulesConditionsObjectOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsObjectOperatorEquals, AppFlagUpdateParamsRulesConditionsObjectOperatorNotEquals, AppFlagUpdateParamsRulesConditionsObjectOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsObjectOperatorLessThan, AppFlagUpdateParamsRulesConditionsObjectOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsObjectOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsObjectOperatorContains, AppFlagUpdateParamsRulesConditionsObjectOperatorStartsWith, AppFlagUpdateParamsRulesConditionsObjectOperatorEndsWith, AppFlagUpdateParamsRulesConditionsObjectOperatorIn, AppFlagUpdateParamsRulesConditionsObjectOperatorNotIn:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsOperator string

const (
	AppFlagUpdateParamsRulesConditionsOperatorEquals              AppFlagUpdateParamsRulesConditionsOperator = "equals"
	AppFlagUpdateParamsRulesConditionsOperatorNotEquals           AppFlagUpdateParamsRulesConditionsOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsOperatorLessThan            AppFlagUpdateParamsRulesConditionsOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsOperatorContains            AppFlagUpdateParamsRulesConditionsOperator = "contains"
	AppFlagUpdateParamsRulesConditionsOperatorStartsWith          AppFlagUpdateParamsRulesConditionsOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsOperatorEndsWith            AppFlagUpdateParamsRulesConditionsOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsOperatorIn                  AppFlagUpdateParamsRulesConditionsOperator = "in"
	AppFlagUpdateParamsRulesConditionsOperatorNotIn               AppFlagUpdateParamsRulesConditionsOperator = "not_in"
)

func (r AppFlagUpdateParamsRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsOperatorEquals, AppFlagUpdateParamsRulesConditionsOperatorNotEquals, AppFlagUpdateParamsRulesConditionsOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsOperatorLessThan, AppFlagUpdateParamsRulesConditionsOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsOperatorContains, AppFlagUpdateParamsRulesConditionsOperatorStartsWith, AppFlagUpdateParamsRulesConditionsOperatorEndsWith, AppFlagUpdateParamsRulesConditionsOperatorIn, AppFlagUpdateParamsRulesConditionsOperatorNotIn:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesRollout struct {
	// Percentage of matching traffic (0–100) served this variation. For multi-way
	// splits, use cumulative upper bounds across rules (e.g. 30, 70, 100).
	Percentage param.Field[float64] `json:"percentage" api:"required"`
	// Context attribute used for sticky bucketing. Defaults to `targetingKey`. If
	// absent at evaluation time, bucketing is random per request.
	Attribute param.Field[string] `json:"attribute"`
}

func (r AppFlagUpdateParamsRulesRollout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsVariationsMap],
// [flagship.AppFlagUpdateParamsVariationsArray].
type AppFlagUpdateParamsVariationsUnion interface {
	ImplementsAppFlagUpdateParamsVariationsUnion()
}

type AppFlagUpdateParamsVariationsMap map[string]interface{}

func (r AppFlagUpdateParamsVariationsMap) ImplementsAppFlagUpdateParamsVariationsUnion() {}

type AppFlagUpdateParamsVariationsArray []interface{}

func (r AppFlagUpdateParamsVariationsArray) ImplementsAppFlagUpdateParamsVariationsUnion() {}

// Value type of the flag's variations. Inferred from the variation values on
// write, so it may be omitted in requests.
type AppFlagUpdateParamsType string

const (
	AppFlagUpdateParamsTypeBoolean AppFlagUpdateParamsType = "boolean"
	AppFlagUpdateParamsTypeString  AppFlagUpdateParamsType = "string"
	AppFlagUpdateParamsTypeNumber  AppFlagUpdateParamsType = "number"
	AppFlagUpdateParamsTypeJson    AppFlagUpdateParamsType = "json"
)

func (r AppFlagUpdateParamsType) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsTypeBoolean, AppFlagUpdateParamsTypeString, AppFlagUpdateParamsTypeNumber, AppFlagUpdateParamsTypeJson:
		return true
	}
	return false
}

type AppFlagUpdateResponseEnvelope struct {
	Errors   []AppFlagUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppFlagUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppFlagUpdateResponse                   `json:"result" api:"required"`
	Success  bool                                    `json:"success" api:"required"`
	JSON     appFlagUpdateResponseEnvelopeJSON       `json:"-"`
}

// appFlagUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppFlagUpdateResponseEnvelope]
type appFlagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateResponseEnvelopeErrors struct {
	Message string                                  `json:"message" api:"required"`
	JSON    appFlagUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// appFlagUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AppFlagUpdateResponseEnvelopeErrors]
type appFlagUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateResponseEnvelopeMessages struct {
	Message string                                    `json:"message" api:"required"`
	JSON    appFlagUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// appFlagUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AppFlagUpdateResponseEnvelopeMessages]
type appFlagUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppFlagListParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Max items to return (1–200).
	Limit param.Field[string] `query:"limit"`
}

// URLQuery serializes [AppFlagListParams]'s query parameters as `url.Values`.
func (r AppFlagListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AppFlagDeleteParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppFlagDeleteResponseEnvelope struct {
	Errors   []AppFlagDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppFlagDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppFlagDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                    `json:"success" api:"required"`
	JSON     appFlagDeleteResponseEnvelopeJSON       `json:"-"`
}

// appFlagDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppFlagDeleteResponseEnvelope]
type appFlagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppFlagDeleteResponseEnvelopeErrors struct {
	Message string                                  `json:"message" api:"required"`
	JSON    appFlagDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// appFlagDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AppFlagDeleteResponseEnvelopeErrors]
type appFlagDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppFlagDeleteResponseEnvelopeMessages struct {
	Message string                                    `json:"message" api:"required"`
	JSON    appFlagDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// appFlagDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AppFlagDeleteResponseEnvelopeMessages]
type appFlagDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AppFlagGetResponseEnvelope struct {
	Errors   []AppFlagGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AppFlagGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   AppFlagGetResponse                   `json:"result" api:"required"`
	Success  bool                                 `json:"success" api:"required"`
	JSON     appFlagGetResponseEnvelopeJSON       `json:"-"`
}

// appFlagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppFlagGetResponseEnvelope]
type appFlagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponseEnvelopeErrors struct {
	Message string                               `json:"message" api:"required"`
	JSON    appFlagGetResponseEnvelopeErrorsJSON `json:"-"`
}

// appFlagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppFlagGetResponseEnvelopeErrors]
type appFlagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponseEnvelopeMessages struct {
	Message string                                 `json:"message" api:"required"`
	JSON    appFlagGetResponseEnvelopeMessagesJSON `json:"-"`
}

// appFlagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppFlagGetResponseEnvelopeMessages]
type appFlagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
