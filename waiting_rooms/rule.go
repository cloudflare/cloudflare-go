// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_rooms

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *RuleService) New(ctx context.Context, waitingRoomID string, params RuleNewParams, opts ...option.RequestOption) (res *[]RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", params.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *RuleService) Update(ctx context.Context, waitingRoomID string, params RuleUpdateParams, opts ...option.RequestOption) (res *[]RuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", params.ZoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists rules for a waiting room.
func (r *RuleService) List(ctx context.Context, waitingRoomID string, query RuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", query.ZoneID, waitingRoomID)
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

// Lists rules for a waiting room.
func (r *RuleService) ListAutoPaging(ctx context.Context, waitingRoomID string, query RuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, waitingRoomID, query, opts...))
}

// Deletes a rule for a waiting room.
func (r *RuleService) Delete(ctx context.Context, waitingRoomID string, ruleID string, params RuleDeleteParams, opts ...option.RequestOption) (res *[]RuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", params.ZoneID, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches a rule for a waiting room.
func (r *RuleService) Edit(ctx context.Context, waitingRoomID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *[]RuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", params.ZoneID, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Rule struct {
	// The action to take when the expression matches.
	Action RuleAction `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression string `json:"expression,required"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool     `json:"enabled"`
	JSON    ruleJSON `json:"-"`
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
	Action      apijson.Field
	Expression  apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleAction string

const (
	RuleActionBypassWaitingRoom RuleAction = "bypass_waiting_room"
)

func (r RuleAction) IsKnown() bool {
	switch r {
	case RuleActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleParam struct {
	// The action to take when the expression matches.
	Action param.Field[RuleAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleNewResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string              `json:"version"`
	JSON    ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleNewResponseAction string

const (
	RuleNewResponseActionBypassWaitingRoom RuleNewResponseAction = "bypass_waiting_room"
)

func (r RuleNewResponseAction) IsKnown() bool {
	switch r {
	case RuleNewResponseActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleUpdateResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleUpdateResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                 `json:"version"`
	JSON    ruleUpdateResponseJSON `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleUpdateResponseAction string

const (
	RuleUpdateResponseActionBypassWaitingRoom RuleUpdateResponseAction = "bypass_waiting_room"
)

func (r RuleUpdateResponseAction) IsKnown() bool {
	switch r {
	case RuleUpdateResponseActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleListResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleListResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string               `json:"version"`
	JSON    ruleListResponseJSON `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleListResponseAction string

const (
	RuleListResponseActionBypassWaitingRoom RuleListResponseAction = "bypass_waiting_room"
)

func (r RuleListResponseAction) IsKnown() bool {
	switch r {
	case RuleListResponseActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleDeleteResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleDeleteResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                 `json:"version"`
	JSON    ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleDeleteResponseAction string

const (
	RuleDeleteResponseActionBypassWaitingRoom RuleDeleteResponseAction = "bypass_waiting_room"
)

func (r RuleDeleteResponseAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleEditResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleEditResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string               `json:"version"`
	JSON    ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleEditResponseAction string

const (
	RuleEditResponseActionBypassWaitingRoom RuleEditResponseAction = "bypass_waiting_room"
)

func (r RuleEditResponseAction) IsKnown() bool {
	switch r {
	case RuleEditResponseActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take when the expression matches.
	Action param.Field[RuleNewParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type RuleNewParamsAction string

const (
	RuleNewParamsActionBypassWaitingRoom RuleNewParamsAction = "bypass_waiting_room"
)

func (r RuleNewParamsAction) IsKnown() bool {
	switch r {
	case RuleNewParamsActionBypassWaitingRoom:
		return true
	}
	return false
}

type RuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleNewResponse                                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleNewResponseEnvelopeJSON       `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
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

type RuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       ruleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeResultInfo]
type ruleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[[]RuleParam] `json:"body,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleUpdateResponse                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleUpdateResponseEnvelopeJSON       `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
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

type RuleUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       ruleUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RuleUpdateResponseEnvelopeResultInfo]
type ruleUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r RuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleDeleteResponse                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleDeleteResponseEnvelopeJSON       `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
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

type RuleDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       ruleDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeResultInfo]
type ruleDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type RuleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take when the expression matches.
	Action param.Field[RuleEditParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[RuleEditParamsPositionUnion] `json:"position"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type RuleEditParamsAction string

const (
	RuleEditParamsActionBypassWaitingRoom RuleEditParamsAction = "bypass_waiting_room"
)

func (r RuleEditParamsAction) IsKnown() bool {
	switch r {
	case RuleEditParamsActionBypassWaitingRoom:
		return true
	}
	return false
}

// Reorder the position of a rule
type RuleEditParamsPosition struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
	// Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the first rule in the ruleset.
	Before param.Field[string] `json:"before"`
	// Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the last rule in the ruleset.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPosition) implementsWaitingRoomsRuleEditParamsPositionUnion() {}

// Reorder the position of a rule
//
// Satisfied by [waiting_rooms.RuleEditParamsPositionObject],
// [waiting_rooms.RuleEditParamsPositionObject],
// [waiting_rooms.RuleEditParamsPositionObject], [RuleEditParamsPosition].
type RuleEditParamsPositionUnion interface {
	implementsWaitingRoomsRuleEditParamsPositionUnion()
}

type RuleEditParamsPositionObject struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r RuleEditParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionObject) implementsWaitingRoomsRuleEditParamsPositionUnion() {}

type RuleEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []RuleEditResponse                                        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    RuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo RuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       ruleEditResponseEnvelopeJSON       `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                `json:"total_count"`
	JSON       ruleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// ruleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeResultInfo]
type ruleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
