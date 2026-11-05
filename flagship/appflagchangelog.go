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
	// Variation the API serves when the flag is off, or when it's on but no rule
	// matches the context. Must be a key in `variations`.
	DefaultVariation string `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled bool `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key string `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules []AppFlagChangelogListResponseObjectAfterRule `json:"rules" api:"required"`
	// Server-inferred value type shared by all of the flag's variations.
	Type AppFlagChangelogListResponseObjectAfterType `json:"type" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations map[string]AppFlagChangelogListResponseObjectAfterVariationsUnion `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description string `json:"description" api:"nullable"`
	// Key-value metadata associated with the flag. Returned only when the flag has
	// metadata.
	Metadata  map[string]AppFlagChangelogListResponseObjectAfterMetadataUnion `json:"metadata"`
	UpdatedAt string                                                          `json:"updated_at"`
	UpdatedBy string                                                          `json:"updated_by"`
	JSON      appFlagChangelogListResponseObjectAfterJSON                     `json:"-"`
}

// appFlagChangelogListResponseObjectAfterJSON contains the JSON metadata for the
// struct [AppFlagChangelogListResponseObjectAfter]
type appFlagChangelogListResponseObjectAfterJSON struct {
	DefaultVariation apijson.Field
	Enabled          apijson.Field
	Key              apijson.Field
	Rules            apijson.Field
	Type             apijson.Field
	Variations       apijson.Field
	Description      apijson.Field
	Metadata         apijson.Field
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
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause].
	Clauses         interface{}                                                           `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion].
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
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition].
func (r AppFlagChangelogListResponseObjectAfterRulesCondition) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition] or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition].
type AppFlagChangelogListResponseObjectAfterRulesConditionsUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition struct {
	Attribute string                                                                                `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseCondition) implementsAppFlagChangelogListResponseObjectAfterRulesCondition() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition struct {
	Clauses         []AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionJSON            `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionJSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalCondition) implementsAppFlagChangelogListResponseObjectAfterRulesCondition() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
	Clauses         interface{}                                                                                          `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                              `json:"value"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClauseJSON `json:"-"`
	union AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesUnion
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClauseJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                               `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         []AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON            `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
	Clauses         interface{}                                                                                                                          `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                              `json:"value"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON `json:"-"`
	union AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                               `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         []AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON            `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
	Clauses         interface{}                                                                                                                                                          `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                              `json:"value"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON `json:"-"`
	union AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                               `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         []AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON            `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause].
	Clauses         interface{}                                                                                                                                                                                          `json:"clauses"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                                                              `json:"value"`
	JSON  appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON `json:"-"`
	union AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) AsUnion() AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion interface {
	implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                                                               `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         []AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON            `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute string                                                                                                                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator   `json:"operator" api:"required"`
	Value     AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion `json:"value" api:"required"`
	JSON      appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON        `json:"-"`
}

// appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]
type appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON) RawJSON() string {
	return r.raw
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap]
// or
// [AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray{}),
		},
	)
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagChangelogListResponseObjectAfterRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorHas                 AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "has"
	AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotHas              AppFlagChangelogListResponseObjectAfterRulesConditionsOperator = "not_has"
)

func (r AppFlagChangelogListResponseObjectAfterRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThan, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThan, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorGreaterThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorLessThanOrEquals, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorContains, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorStartsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorEndsWith, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorIn, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotIn, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorHas, AppFlagChangelogListResponseObjectAfterRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagChangelogListResponseObjectAfterRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Server-inferred value type shared by all of the flag's variations.
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

// Union satisfied by [shared.UnionBool], [shared.UnionString] or
// [shared.UnionFloat].
type AppFlagChangelogListResponseObjectAfterMetadataUnion interface {
	ImplementsAppFlagChangelogListResponseObjectAfterMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagChangelogListResponseObjectAfterMetadataUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
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
	// Cloudflare account ID that owns the Flagship app.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Max items to return (1–200).
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AppFlagChangelogListParams]'s query parameters as
// `url.Values`.
func (r AppFlagChangelogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
