// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options   []option.RequestOption
	CatchAlls *RuleCatchAllService
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	r.CatchAlls = NewRuleCatchAllService(opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *EmailRoutingRule, err error) {
	var env RuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *RuleService) Update(ctx context.Context, ruleIdentifier string, params RuleUpdateParams, opts ...option.RequestOption) (res *EmailRoutingRule, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", params.ZoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing routing rules.
func (r *RuleService) List(ctx context.Context, params RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[EmailRoutingRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules", params.ZoneID)
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

// Lists existing routing rules.
func (r *RuleService) ListAutoPaging(ctx context.Context, params RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[EmailRoutingRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a specific routing rule.
func (r *RuleService) Delete(ctx context.Context, ruleIdentifier string, body RuleDeleteParams, opts ...option.RequestOption) (res *EmailRoutingRule, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", body.ZoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information for a specific routing rule already created.
func (r *RuleService) Get(ctx context.Context, ruleIdentifier string, query RuleGetParams, opts ...option.RequestOption) (res *EmailRoutingRule, err error) {
	var env RuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", query.ZoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Actions pattern.
type Action struct {
	// Type of supported action.
	Type  ActionType `json:"type,required"`
	Value []string   `json:"value,required"`
	JSON  actionJSON `json:"-"`
}

// actionJSON contains the JSON metadata for the struct [Action]
type actionJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Action) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type ActionType string

const (
	ActionTypeDrop    ActionType = "drop"
	ActionTypeForward ActionType = "forward"
	ActionTypeWorker  ActionType = "worker"
)

func (r ActionType) IsKnown() bool {
	switch r {
	case ActionTypeDrop, ActionTypeForward, ActionTypeWorker:
		return true
	}
	return false
}

// Actions pattern.
type ActionParam struct {
	// Type of supported action.
	Type  param.Field[ActionType] `json:"type,required"`
	Value param.Field[[]string]   `json:"value,required"`
}

func (r ActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRoutingRule struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []Action `json:"actions"`
	// Routing rule status.
	Enabled EmailRoutingRuleEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []Matcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	Tag  string               `json:"tag"`
	JSON emailRoutingRuleJSON `json:"-"`
}

// emailRoutingRuleJSON contains the JSON metadata for the struct
// [EmailRoutingRule]
type emailRoutingRuleJSON struct {
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

func (r *EmailRoutingRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRuleJSON) RawJSON() string {
	return r.raw
}

// Routing rule status.
type EmailRoutingRuleEnabled bool

const (
	EmailRoutingRuleEnabledTrue  EmailRoutingRuleEnabled = true
	EmailRoutingRuleEnabledFalse EmailRoutingRuleEnabled = false
)

func (r EmailRoutingRuleEnabled) IsKnown() bool {
	switch r {
	case EmailRoutingRuleEnabledTrue, EmailRoutingRuleEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type Matcher struct {
	// Field for type matcher.
	Field MatcherField `json:"field,required"`
	// Type of matcher.
	Type MatcherType `json:"type,required"`
	// Value for matcher.
	Value string      `json:"value,required"`
	JSON  matcherJSON `json:"-"`
}

// matcherJSON contains the JSON metadata for the struct [Matcher]
type matcherJSON struct {
	Field       apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Matcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matcherJSON) RawJSON() string {
	return r.raw
}

// Field for type matcher.
type MatcherField string

const (
	MatcherFieldTo MatcherField = "to"
)

func (r MatcherField) IsKnown() bool {
	switch r {
	case MatcherFieldTo:
		return true
	}
	return false
}

// Type of matcher.
type MatcherType string

const (
	MatcherTypeLiteral MatcherType = "literal"
)

func (r MatcherType) IsKnown() bool {
	switch r {
	case MatcherTypeLiteral:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type MatcherParam struct {
	// Field for type matcher.
	Field param.Field[MatcherField] `json:"field,required"`
	// Type of matcher.
	Type param.Field[MatcherType] `json:"type,required"`
	// Value for matcher.
	Value param.Field[string] `json:"value,required"`
}

func (r MatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// List actions patterns.
	Actions param.Field[[]ActionParam] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]MatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleNewParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Routing rule status.
type RuleNewParamsEnabled bool

const (
	RuleNewParamsEnabledTrue  RuleNewParamsEnabled = true
	RuleNewParamsEnabledFalse RuleNewParamsEnabled = false
)

func (r RuleNewParamsEnabled) IsKnown() bool {
	switch r {
	case RuleNewParamsEnabledTrue, RuleNewParamsEnabledFalse:
		return true
	}
	return false
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingRule               `json:"result"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// List actions patterns.
	Actions param.Field[[]ActionParam] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]MatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[RuleUpdateParamsEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Routing rule status.
type RuleUpdateParamsEnabled bool

const (
	RuleUpdateParamsEnabledTrue  RuleUpdateParamsEnabled = true
	RuleUpdateParamsEnabledFalse RuleUpdateParamsEnabled = false
)

func (r RuleUpdateParamsEnabled) IsKnown() bool {
	switch r {
	case RuleUpdateParamsEnabledTrue, RuleUpdateParamsEnabledFalse:
		return true
	}
	return false
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingRule                  `json:"result"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Filter by enabled routing rules.
	Enabled param.Field[RuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by enabled routing rules.
type RuleListParamsEnabled bool

const (
	RuleListParamsEnabledTrue  RuleListParamsEnabled = true
	RuleListParamsEnabledFalse RuleListParamsEnabled = false
)

func (r RuleListParamsEnabled) IsKnown() bool {
	switch r {
	case RuleListParamsEnabledTrue, RuleListParamsEnabledFalse:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingRule                  `json:"result"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  EmailRoutingRule               `json:"result"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
