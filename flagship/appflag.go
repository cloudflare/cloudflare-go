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

// Creates a flag. Returns 409 if the key already exists. `type` is always inferred
// from variation values; legacy request-side values are ignored. Optional
// `metadata` can be attached at creation.
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

// Replaces the entire flag definition. Omitted fields are cleared, not preserved —
// read before writing. The path key identifies the flag and cannot be renamed by
// changing the body `key`. A changelog entry is appended when the flag content
// actually changes; identical updates are no-ops.
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

// Deletes a flag permanently. After deletion propagates, direct evaluations return
// not found; typed binding accessors may return the caller-supplied default.
// Cannot be undone.
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

// Returns the full flag definition including rules, variations, metadata, and
// audit fields.
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
	Rules []AppFlagNewResponseRule `json:"rules" api:"required"`
	// Server-inferred value type shared by all of the flag's variations.
	Type AppFlagNewResponseType `json:"type" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations map[string]AppFlagNewResponseVariationsUnion `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description string `json:"description" api:"nullable"`
	// Key-value metadata associated with the flag. Returned only when the flag has
	// metadata.
	Metadata  map[string]AppFlagNewResponseMetadataUnion `json:"metadata"`
	UpdatedAt string                                     `json:"updated_at"`
	UpdatedBy string                                     `json:"updated_by"`
	JSON      appFlagNewResponseJSON                     `json:"-"`
}

// appFlagNewResponseJSON contains the JSON metadata for the struct
// [AppFlagNewResponse]
type appFlagNewResponseJSON struct {
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
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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
	// [[]AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause].
	Clauses         interface{}                                      `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion].
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
// [AppFlagNewResponseRulesConditionsFlagshipBaseCondition],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalCondition].
func (r AppFlagNewResponseRulesCondition) AsUnion() AppFlagNewResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagNewResponseRulesConditionsFlagshipBaseCondition] or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalCondition].
type AppFlagNewResponseRulesConditionsUnion interface {
	implementsAppFlagNewResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalCondition{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipBaseCondition struct {
	Attribute string                                                           `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipBaseConditionJSON contains the JSON
// metadata for the struct [AppFlagNewResponseRulesConditionsFlagshipBaseCondition]
type appFlagNewResponseRulesConditionsFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipBaseCondition) implementsAppFlagNewResponseRulesCondition() {
}

type AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagNewResponseRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueMap] or
// [AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalCondition struct {
	Clauses         []AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagNewResponseRulesConditionsFlagshipLogicalConditionJSON            `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionJSON contains the JSON
// metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalCondition]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionJSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalCondition) implementsAppFlagNewResponseRulesCondition() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
	Clauses         interface{}                                                                     `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                         `json:"value"`
	JSON  appFlagNewResponseRulesConditionsFlagshipLogicalConditionClauseJSON `json:"-"`
	union AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesUnion
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClauseJSON contains the
// JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause) AsUnion() AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesUnion interface {
	implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute string                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         []AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON            `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
	Clauses         interface{}                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                         `json:"value"`
	JSON  appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON `json:"-"`
	union AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) AsUnion() AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion interface {
	implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         []AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON            `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
	Clauses         interface{}                                                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                         `json:"value"`
	JSON  appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON `json:"-"`
	union AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) AsUnion() AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion interface {
	implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         []AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON            `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause].
	Clauses         interface{}                                                                                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                                         `json:"value"`
	JSON  appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON `json:"-"`
	union AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) AsUnion() AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion interface {
	implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         []AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON            `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute string                                                                                                                                                                                                     `json:"attribute" api:"required"`
	Operator  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator   `json:"operator" api:"required"`
	Value     AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion `json:"value" api:"required"`
	JSON      appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON        `json:"-"`
}

// appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]
type appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON) RawJSON() string {
	return r.raw
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap]
// or
// [AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray{}),
		},
	)
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagNewResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagNewResponseRulesConditionsOperatorHas                 AppFlagNewResponseRulesConditionsOperator = "has"
	AppFlagNewResponseRulesConditionsOperatorNotHas              AppFlagNewResponseRulesConditionsOperator = "not_has"
)

func (r AppFlagNewResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagNewResponseRulesConditionsOperatorEquals, AppFlagNewResponseRulesConditionsOperatorNotEquals, AppFlagNewResponseRulesConditionsOperatorGreaterThan, AppFlagNewResponseRulesConditionsOperatorLessThan, AppFlagNewResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagNewResponseRulesConditionsOperatorLessThanOrEquals, AppFlagNewResponseRulesConditionsOperatorContains, AppFlagNewResponseRulesConditionsOperatorStartsWith, AppFlagNewResponseRulesConditionsOperatorEndsWith, AppFlagNewResponseRulesConditionsOperatorIn, AppFlagNewResponseRulesConditionsOperatorNotIn, AppFlagNewResponseRulesConditionsOperatorHas, AppFlagNewResponseRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewResponseRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Server-inferred value type shared by all of the flag's variations.
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

// Union satisfied by [shared.UnionBool], [shared.UnionString] or
// [shared.UnionFloat].
type AppFlagNewResponseMetadataUnion interface {
	ImplementsAppFlagNewResponseMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagNewResponseMetadataUnion)(nil)).Elem(),
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

type AppFlagUpdateResponse struct {
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
	Rules []AppFlagUpdateResponseRule `json:"rules" api:"required"`
	// Server-inferred value type shared by all of the flag's variations.
	Type AppFlagUpdateResponseType `json:"type" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations map[string]AppFlagUpdateResponseVariationsUnion `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description string `json:"description" api:"nullable"`
	// Key-value metadata associated with the flag. Returned only when the flag has
	// metadata.
	Metadata  map[string]AppFlagUpdateResponseMetadataUnion `json:"metadata"`
	UpdatedAt string                                        `json:"updated_at"`
	UpdatedBy string                                        `json:"updated_by"`
	JSON      appFlagUpdateResponseJSON                     `json:"-"`
}

// appFlagUpdateResponseJSON contains the JSON metadata for the struct
// [AppFlagUpdateResponse]
type appFlagUpdateResponseJSON struct {
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
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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
	// [[]AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause].
	Clauses         interface{}                                         `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion].
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
// [AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition].
func (r AppFlagUpdateResponseRulesCondition) AsUnion() AppFlagUpdateResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition]
// or [AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition].
type AppFlagUpdateResponseRulesConditionsUnion interface {
	implementsAppFlagUpdateResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition struct {
	Attribute string                                                              `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipBaseConditionJSON contains the JSON
// metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition]
type appFlagUpdateResponseRulesConditionsFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipBaseCondition) implementsAppFlagUpdateResponseRulesCondition() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueMap] or
// [AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition struct {
	Clauses         []AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionJSON            `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionJSON contains the
// JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionJSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalCondition) implementsAppFlagUpdateResponseRulesCondition() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
	Clauses         interface{}                                                                        `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                            `json:"value"`
	JSON  appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClauseJSON `json:"-"`
	union AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesUnion
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClauseJSON contains
// the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause) AsUnion() AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesUnion interface {
	implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute string                                                                                             `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         []AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON            `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
	Clauses         interface{}                                                                                                        `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                            `json:"value"`
	JSON  appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON `json:"-"`
	union AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) AsUnion() AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion interface {
	implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                             `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         []AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON            `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
	Clauses         interface{}                                                                                                                                        `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                            `json:"value"`
	JSON  appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON `json:"-"`
	union AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) AsUnion() AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion interface {
	implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                             `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         []AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON            `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause].
	Clauses         interface{}                                                                                                                                                                        `json:"clauses"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                                            `json:"value"`
	JSON  appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON `json:"-"`
	union AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) AsUnion() AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion interface {
	implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                                             `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         []AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON            `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute string                                                                                                                                                                                                        `json:"attribute" api:"required"`
	Operator  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator   `json:"operator" api:"required"`
	Value     AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion `json:"value" api:"required"`
	JSON      appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON        `json:"-"`
}

// appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]
type appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON) RawJSON() string {
	return r.raw
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap]
// or
// [AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray{}),
		},
	)
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagUpdateResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagUpdateResponseRulesConditionsOperatorHas                 AppFlagUpdateResponseRulesConditionsOperator = "has"
	AppFlagUpdateResponseRulesConditionsOperatorNotHas              AppFlagUpdateResponseRulesConditionsOperator = "not_has"
)

func (r AppFlagUpdateResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateResponseRulesConditionsOperatorEquals, AppFlagUpdateResponseRulesConditionsOperatorNotEquals, AppFlagUpdateResponseRulesConditionsOperatorGreaterThan, AppFlagUpdateResponseRulesConditionsOperatorLessThan, AppFlagUpdateResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagUpdateResponseRulesConditionsOperatorLessThanOrEquals, AppFlagUpdateResponseRulesConditionsOperatorContains, AppFlagUpdateResponseRulesConditionsOperatorStartsWith, AppFlagUpdateResponseRulesConditionsOperatorEndsWith, AppFlagUpdateResponseRulesConditionsOperatorIn, AppFlagUpdateResponseRulesConditionsOperatorNotIn, AppFlagUpdateResponseRulesConditionsOperatorHas, AppFlagUpdateResponseRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateResponseRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Server-inferred value type shared by all of the flag's variations.
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

// Union satisfied by [shared.UnionBool], [shared.UnionString] or
// [shared.UnionFloat].
type AppFlagUpdateResponseMetadataUnion interface {
	ImplementsAppFlagUpdateResponseMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagUpdateResponseMetadataUnion)(nil)).Elem(),
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

type AppFlagListResponse struct {
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
	Rules []AppFlagListResponseRule `json:"rules" api:"required"`
	// Server-inferred value type shared by all of the flag's variations.
	Type AppFlagListResponseType `json:"type" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations map[string]AppFlagListResponseVariationsUnion `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description string `json:"description" api:"nullable"`
	// Key-value metadata associated with the flag. Returned only when the flag has
	// metadata.
	Metadata  map[string]AppFlagListResponseMetadataUnion `json:"metadata"`
	UpdatedAt string                                      `json:"updated_at"`
	UpdatedBy string                                      `json:"updated_by"`
	JSON      appFlagListResponseJSON                     `json:"-"`
}

// appFlagListResponseJSON contains the JSON metadata for the struct
// [AppFlagListResponse]
type appFlagListResponseJSON struct {
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
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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
	// [[]AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause].
	Clauses         interface{}                                       `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion].
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
// [AppFlagListResponseRulesConditionsFlagshipBaseCondition],
// [AppFlagListResponseRulesConditionsFlagshipLogicalCondition].
func (r AppFlagListResponseRulesCondition) AsUnion() AppFlagListResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagListResponseRulesConditionsFlagshipBaseCondition] or
// [AppFlagListResponseRulesConditionsFlagshipLogicalCondition].
type AppFlagListResponseRulesConditionsUnion interface {
	implementsAppFlagListResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalCondition{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipBaseCondition struct {
	Attribute string                                                            `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipBaseConditionJSON contains the JSON
// metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipBaseCondition]
type appFlagListResponseRulesConditionsFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipBaseCondition) implementsAppFlagListResponseRulesCondition() {
}

type AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagListResponseRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipBaseConditionValueMap] or
// [AppFlagListResponseRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalCondition struct {
	Clauses         []AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagListResponseRulesConditionsFlagshipLogicalConditionJSON            `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionJSON contains the JSON
// metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalCondition]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionJSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalCondition) implementsAppFlagListResponseRulesCondition() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
	Clauses         interface{}                                                                      `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                          `json:"value"`
	JSON  appFlagListResponseRulesConditionsFlagshipLogicalConditionClauseJSON `json:"-"`
	union AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesUnion
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClauseJSON contains
// the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClause) AsUnion() AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesUnion interface {
	implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute string                                                                                           `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         []AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON            `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
	Clauses         interface{}                                                                                                      `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                          `json:"value"`
	JSON  appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON `json:"-"`
	union AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) AsUnion() AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion interface {
	implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                           `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         []AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON            `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
	Clauses         interface{}                                                                                                                                      `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                          `json:"value"`
	JSON  appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON `json:"-"`
	union AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) AsUnion() AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion interface {
	implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                           `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         []AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON            `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause].
	Clauses         interface{}                                                                                                                                                                      `json:"clauses"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                                          `json:"value"`
	JSON  appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON `json:"-"`
	union AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) AsUnion() AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion interface {
	implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                                           `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         []AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON            `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute string                                                                                                                                                                                                      `json:"attribute" api:"required"`
	Operator  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator   `json:"operator" api:"required"`
	Value     AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion `json:"value" api:"required"`
	JSON      appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON        `json:"-"`
}

// appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]
type appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON) RawJSON() string {
	return r.raw
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap]
// or
// [AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray{}),
		},
	)
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagListResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagListResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagListResponseRulesConditionsOperatorHas                 AppFlagListResponseRulesConditionsOperator = "has"
	AppFlagListResponseRulesConditionsOperatorNotHas              AppFlagListResponseRulesConditionsOperator = "not_has"
)

func (r AppFlagListResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagListResponseRulesConditionsOperatorEquals, AppFlagListResponseRulesConditionsOperatorNotEquals, AppFlagListResponseRulesConditionsOperatorGreaterThan, AppFlagListResponseRulesConditionsOperatorLessThan, AppFlagListResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagListResponseRulesConditionsOperatorLessThanOrEquals, AppFlagListResponseRulesConditionsOperatorContains, AppFlagListResponseRulesConditionsOperatorStartsWith, AppFlagListResponseRulesConditionsOperatorEndsWith, AppFlagListResponseRulesConditionsOperatorIn, AppFlagListResponseRulesConditionsOperatorNotIn, AppFlagListResponseRulesConditionsOperatorHas, AppFlagListResponseRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagListResponseRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Server-inferred value type shared by all of the flag's variations.
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

// Union satisfied by [shared.UnionBool], [shared.UnionString] or
// [shared.UnionFloat].
type AppFlagListResponseMetadataUnion interface {
	ImplementsAppFlagListResponseMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagListResponseMetadataUnion)(nil)).Elem(),
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
	Rules []AppFlagGetResponseRule `json:"rules" api:"required"`
	// Server-inferred value type shared by all of the flag's variations.
	Type AppFlagGetResponseType `json:"type" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations map[string]AppFlagGetResponseVariationsUnion `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description string `json:"description" api:"nullable"`
	// Key-value metadata associated with the flag. Returned only when the flag has
	// metadata.
	Metadata  map[string]AppFlagGetResponseMetadataUnion `json:"metadata"`
	UpdatedAt string                                     `json:"updated_at"`
	UpdatedBy string                                     `json:"updated_by"`
	JSON      appFlagGetResponseJSON                     `json:"-"`
}

// appFlagGetResponseJSON contains the JSON metadata for the struct
// [AppFlagGetResponse]
type appFlagGetResponseJSON struct {
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
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority int64 `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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
	// [[]AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause].
	Clauses         interface{}                                      `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion].
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
// [AppFlagGetResponseRulesConditionsFlagshipBaseCondition],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalCondition].
func (r AppFlagGetResponseRulesCondition) AsUnion() AppFlagGetResponseRulesConditionsUnion {
	return r.union
}

// Union satisfied by [AppFlagGetResponseRulesConditionsFlagshipBaseCondition] or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalCondition].
type AppFlagGetResponseRulesConditionsUnion interface {
	implementsAppFlagGetResponseRulesCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalCondition{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipBaseCondition struct {
	Attribute string                                                           `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipBaseConditionJSON contains the JSON
// metadata for the struct [AppFlagGetResponseRulesConditionsFlagshipBaseCondition]
type appFlagGetResponseRulesConditionsFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipBaseCondition) implementsAppFlagGetResponseRulesCondition() {
}

type AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagGetResponseRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueMap] or
// [AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalCondition struct {
	Clauses         []AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagGetResponseRulesConditionsFlagshipLogicalConditionJSON            `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionJSON contains the JSON
// metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalCondition]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionJSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalCondition) implementsAppFlagGetResponseRulesCondition() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
	Clauses         interface{}                                                                     `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                         `json:"value"`
	JSON  appFlagGetResponseRulesConditionsFlagshipLogicalConditionClauseJSON `json:"-"`
	union AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesUnion
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClauseJSON contains the
// JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause) AsUnion() AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesUnion interface {
	implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute string                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         []AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON            `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
	Clauses         interface{}                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                         `json:"value"`
	JSON  appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON `json:"-"`
	union AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) AsUnion() AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion interface {
	implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         []AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON            `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
	Clauses         interface{}                                                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                         `json:"value"`
	JSON  appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON `json:"-"`
	union AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) AsUnion() AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion interface {
	implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         []AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON            `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute string `json:"attribute"`
	// This field can have the runtime type of
	// [[]AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause].
	Clauses         interface{}                                                                                                                                                                     `json:"clauses"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator `json:"logical_operator"`
	Operator        AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator        `json:"operator"`
	// This field can have the runtime type of
	// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion].
	Value interface{}                                                                                                                                                         `json:"value"`
	JSON  appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON `json:"-"`
	union AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON struct {
	Attribute       apijson.Field
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	Operator        apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseJSON) RawJSON() string {
	return r.raw
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) UnmarshalJSON(data []byte) (err error) {
	*r = AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) AsUnion() AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion {
	return r.union
}

// Union satisfied by
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion interface {
	implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute string                                                                                                                                                                                          `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON       `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionJSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         []AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause        `json:"clauses" api:"required"`
	LogicalOperator AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator `json:"logical_operator" api:"required"`
	JSON            appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON            `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON struct {
	Clauses         apijson.Field
	LogicalOperator apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4JSON) RawJSON() string {
	return r.raw
}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute string                                                                                                                                                                                                     `json:"attribute" api:"required"`
	Operator  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator   `json:"operator" api:"required"`
	Value     AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion `json:"value" api:"required"`
	JSON      appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON        `json:"-"`
}

// appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON
// contains the JSON metadata for the struct
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]
type appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON struct {
	Attribute   apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClauseJSON) RawJSON() string {
	return r.raw
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool],
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap]
// or
// [AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion)(nil)).Elem(),
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
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray{}),
		},
	)
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagGetResponseRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagGetResponseRulesConditionsOperatorHas                 AppFlagGetResponseRulesConditionsOperator = "has"
	AppFlagGetResponseRulesConditionsOperatorNotHas              AppFlagGetResponseRulesConditionsOperator = "not_has"
)

func (r AppFlagGetResponseRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagGetResponseRulesConditionsOperatorEquals, AppFlagGetResponseRulesConditionsOperatorNotEquals, AppFlagGetResponseRulesConditionsOperatorGreaterThan, AppFlagGetResponseRulesConditionsOperatorLessThan, AppFlagGetResponseRulesConditionsOperatorGreaterThanOrEquals, AppFlagGetResponseRulesConditionsOperatorLessThanOrEquals, AppFlagGetResponseRulesConditionsOperatorContains, AppFlagGetResponseRulesConditionsOperatorStartsWith, AppFlagGetResponseRulesConditionsOperatorEndsWith, AppFlagGetResponseRulesConditionsOperatorIn, AppFlagGetResponseRulesConditionsOperatorNotIn, AppFlagGetResponseRulesConditionsOperatorHas, AppFlagGetResponseRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagGetResponseRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Server-inferred value type shared by all of the flag's variations.
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

// Union satisfied by [shared.UnionBool], [shared.UnionString] or
// [shared.UnionFloat].
type AppFlagGetResponseMetadataUnion interface {
	ImplementsAppFlagGetResponseMetadataUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppFlagGetResponseMetadataUnion)(nil)).Elem(),
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

type AppFlagNewParams struct {
	// Cloudflare account ID that owns the Flagship app.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Variation the API serves when the flag is off, or when it's on but no rule
	// matches the context. Must be a key in `variations`.
	DefaultVariation param.Field[string] `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key param.Field[string] `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules param.Field[[]AppFlagNewParamsRule] `json:"rules" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations param.Field[map[string]AppFlagNewParamsVariationsUnion] `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description param.Field[string] `json:"description"`
	// Key-value metadata associated with the flag. `null` or an empty object clears
	// existing metadata; omission semantics depend on the endpoint (PUT clears it,
	// PATCH preserves it).
	Metadata param.Field[map[string]AppFlagNewParamsMetadataUnion] `json:"metadata"`
	// Deprecated compatibility field. Omit it; the API ignores this value and infers
	// the type from the flag's variations.
	Type param.Field[AppFlagNewParamsType] `json:"type"`
}

func (r AppFlagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagNewParamsRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions param.Field[[]AppFlagNewParamsRulesConditionUnion] `json:"conditions" api:"required"`
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority param.Field[int64] `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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

// Satisfied by [flagship.AppFlagNewParamsRulesConditionsFlagshipBaseCondition],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalCondition],
// [AppFlagNewParamsRulesCondition].
type AppFlagNewParamsRulesConditionUnion interface {
	implementsAppFlagNewParamsRulesConditionUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                         `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipBaseCondition) implementsAppFlagNewParamsRulesConditionUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagNewParamsRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalCondition struct {
	Clauses         param.Field[[]AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalCondition) implementsAppFlagNewParamsRulesConditionUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute       param.Field[string]                                                                        `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                   `json:"clauses"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                   `json:"value"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClause) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1],
// [AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClause].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion interface {
	implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                        `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         param.Field[[]AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute       param.Field[string]                                                                                                        `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                   `json:"clauses"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                   `json:"value"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2],
// [AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion interface {
	implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                        `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         param.Field[[]AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute       param.Field[string]                                                                                                                                        `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                                                   `json:"clauses"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                                                   `json:"value"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3],
// [AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion interface {
	implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                                                        `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         param.Field[[]AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute       param.Field[string]                                                                                                                                                                        `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                                                                                   `json:"clauses"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                                                                                   `json:"value"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4],
// [AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion interface {
	implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                                                                                        `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         param.Field[[]AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]        `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute param.Field[string]                                                                                                                                                                                                   `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion] `json:"value" api:"required"`
}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap],
// [flagship.AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagNewParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagNewParamsRulesConditionsOperatorHas                 AppFlagNewParamsRulesConditionsOperator = "has"
	AppFlagNewParamsRulesConditionsOperatorNotHas              AppFlagNewParamsRulesConditionsOperator = "not_has"
)

func (r AppFlagNewParamsRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagNewParamsRulesConditionsOperatorEquals, AppFlagNewParamsRulesConditionsOperatorNotEquals, AppFlagNewParamsRulesConditionsOperatorGreaterThan, AppFlagNewParamsRulesConditionsOperatorLessThan, AppFlagNewParamsRulesConditionsOperatorGreaterThanOrEquals, AppFlagNewParamsRulesConditionsOperatorLessThanOrEquals, AppFlagNewParamsRulesConditionsOperatorContains, AppFlagNewParamsRulesConditionsOperatorStartsWith, AppFlagNewParamsRulesConditionsOperatorEndsWith, AppFlagNewParamsRulesConditionsOperatorIn, AppFlagNewParamsRulesConditionsOperatorNotIn, AppFlagNewParamsRulesConditionsOperatorHas, AppFlagNewParamsRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagNewParamsRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Satisfied by [shared.UnionBool], [shared.UnionString], [shared.UnionFloat].
type AppFlagNewParamsMetadataUnion interface {
	ImplementsAppFlagNewParamsMetadataUnion()
}

// Deprecated compatibility field. Omit it; the API ignores this value and infers
// the type from the flag's variations.
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
	// Cloudflare account ID that owns the Flagship app.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Variation the API serves when the flag is off, or when it's on but no rule
	// matches the context. Must be a key in `variations`.
	DefaultVariation param.Field[string] `json:"default_variation" api:"required"`
	// When false, the flag bypasses all rules and always serves `default_variation`.
	Enabled param.Field[bool] `json:"enabled" api:"required"`
	// Unique identifier for the flag within an app. Used in all evaluation and SDK
	// calls.
	Key param.Field[string] `json:"key" api:"required"`
	// Targeting rules evaluated in ascending `priority`; the first matching rule wins.
	// An empty array means the flag always serves `default_variation`.
	Rules param.Field[[]AppFlagUpdateParamsRule] `json:"rules" api:"required"`
	// Map of variation name to value. All values share the same type (boolean, string,
	// number, or JSON object/array), and each serialized value stays within 10KB.
	Variations param.Field[map[string]AppFlagUpdateParamsVariationsUnion] `json:"variations" api:"required"`
	// Optional operator-facing description. It does not affect flag evaluation.
	Description param.Field[string] `json:"description"`
	// Key-value metadata associated with the flag. `null` or an empty object clears
	// existing metadata; omission semantics depend on the endpoint (PUT clears it,
	// PATCH preserves it).
	Metadata param.Field[map[string]AppFlagUpdateParamsMetadataUnion] `json:"metadata"`
	// Deprecated compatibility field. Omit it; the API ignores this value and infers
	// the type from the flag's variations.
	Type param.Field[AppFlagUpdateParamsType] `json:"type"`
}

func (r AppFlagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagUpdateParamsRule struct {
	// Conditions the context must satisfy for this rule to match. An empty array
	// matches all contexts.
	Conditions param.Field[[]AppFlagUpdateParamsRulesConditionUnion] `json:"conditions" api:"required"`
	// Evaluation order: the API evaluates rules with lower numbers first. Must be
	// unique across the flag's rules.
	Priority param.Field[int64] `json:"priority" api:"required"`
	// Variation the API serves when this rule matches. Must be a key in `variations`.
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

// Satisfied by [flagship.AppFlagUpdateParamsRulesConditionsFlagshipBaseCondition],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalCondition],
// [AppFlagUpdateParamsRulesCondition].
type AppFlagUpdateParamsRulesConditionUnion interface {
	implementsAppFlagUpdateParamsRulesConditionUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                            `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipBaseCondition) implementsAppFlagUpdateParamsRulesConditionUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalCondition struct {
	Clauses         param.Field[[]AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalCondition) implementsAppFlagUpdateParamsRulesConditionUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClause struct {
	Attribute       param.Field[string]                                                                           `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                      `json:"clauses"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                      `json:"value"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClause) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1],
// [AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClause].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion interface {
	implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                           `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseCondition) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1 struct {
	Clauses         param.Field[[]AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause struct {
	Attribute       param.Field[string]                                                                                                           `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                      `json:"clauses"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                      `json:"value"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2],
// [AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1Clause].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion interface {
	implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                           `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseCondition) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2 struct {
	Clauses         param.Field[[]AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause struct {
	Attribute       param.Field[string]                                                                                                                                           `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                                                      `json:"clauses"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                                                      `json:"value"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3],
// [AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2Clause].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion interface {
	implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                                                           `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseCondition) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3 struct {
	Clauses         param.Field[[]AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion]   `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause struct {
	Attribute       param.Field[string]                                                                                                                                                                           `json:"attribute"`
	Clauses         param.Field[interface{}]                                                                                                                                                                      `json:"clauses"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator] `json:"logical_operator"`
	Operator        param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator]        `json:"operator"`
	Value           param.Field[interface{}]                                                                                                                                                                      `json:"value"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

// Satisfied by
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4],
// [AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3Clause].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion interface {
	implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition struct {
	Attribute param.Field[string]                                                                                                                                                                                           `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseCondition) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipBaseConditionValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4 struct {
	Clauses         param.Field[[]AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause]        `json:"clauses" api:"required"`
	LogicalOperator param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator] `json:"logical_operator" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4) implementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClauseUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause struct {
	Attribute param.Field[string]                                                                                                                                                                                                      `json:"attribute" api:"required"`
	Operator  param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator]   `json:"operator" api:"required"`
	Value     param.Field[AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion] `json:"value" api:"required"`
}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4Clause) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesOperatorNotHas:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap],
// [flagship.AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray].
type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion interface {
	ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion()
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap map[string]interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueMap) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray []interface{}

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueArray) ImplementsAppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4ClausesValueUnion() {
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesFlagshipLogicalCondition4LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesFlagshipLogicalCondition3LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesFlagshipLogicalCondition2LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1ClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesFlagshipLogicalCondition1LogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesLogicalOperatorOr:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEquals              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals           AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan         AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "greater_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals    AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "less_than_or_equals"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorContains            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "contains"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith          AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "starts_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith            AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "ends_with"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorIn                  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn               AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_in"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorHas                 AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "has"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas              AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThan, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorContains, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorStartsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorEndsWith, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotIn, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorHas, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionClausesOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperator string

const (
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperator = "AND"
	AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorOr  AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperator = "OR"
)

func (r AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorAnd, AppFlagUpdateParamsRulesConditionsFlagshipLogicalConditionLogicalOperatorOr:
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
	AppFlagUpdateParamsRulesConditionsOperatorHas                 AppFlagUpdateParamsRulesConditionsOperator = "has"
	AppFlagUpdateParamsRulesConditionsOperatorNotHas              AppFlagUpdateParamsRulesConditionsOperator = "not_has"
)

func (r AppFlagUpdateParamsRulesConditionsOperator) IsKnown() bool {
	switch r {
	case AppFlagUpdateParamsRulesConditionsOperatorEquals, AppFlagUpdateParamsRulesConditionsOperatorNotEquals, AppFlagUpdateParamsRulesConditionsOperatorGreaterThan, AppFlagUpdateParamsRulesConditionsOperatorLessThan, AppFlagUpdateParamsRulesConditionsOperatorGreaterThanOrEquals, AppFlagUpdateParamsRulesConditionsOperatorLessThanOrEquals, AppFlagUpdateParamsRulesConditionsOperatorContains, AppFlagUpdateParamsRulesConditionsOperatorStartsWith, AppFlagUpdateParamsRulesConditionsOperatorEndsWith, AppFlagUpdateParamsRulesConditionsOperatorIn, AppFlagUpdateParamsRulesConditionsOperatorNotIn, AppFlagUpdateParamsRulesConditionsOperatorHas, AppFlagUpdateParamsRulesConditionsOperatorNotHas:
		return true
	}
	return false
}

type AppFlagUpdateParamsRulesRollout struct {
	// Percentage of matching traffic (0–100, up to 2 decimal places) served this
	// variation. For multi-way splits, use cumulative upper bounds across rules (e.g.
	// 30, 70, 100).
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

// Satisfied by [shared.UnionBool], [shared.UnionString], [shared.UnionFloat].
type AppFlagUpdateParamsMetadataUnion interface {
	ImplementsAppFlagUpdateParamsMetadataUnion()
}

// Deprecated compatibility field. Omit it; the API ignores this value and infers
// the type from the flag's variations.
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
	// Cloudflare account ID that owns the Flagship app.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Pagination cursor from a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Max items to return (1–200).
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AppFlagListParams]'s query parameters as `url.Values`.
func (r AppFlagListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AppFlagDeleteParams struct {
	// Cloudflare account ID that owns the Flagship app.
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
	// Cloudflare account ID that owns the Flagship app.
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
