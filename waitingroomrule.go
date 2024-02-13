// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WaitingRoomRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWaitingRoomRuleService] method
// instead.
type WaitingRoomRuleService struct {
	Options []option.RequestOption
}

// NewWaitingRoomRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWaitingRoomRuleService(opts ...option.RequestOption) (r *WaitingRoomRuleService) {
	r = &WaitingRoomRuleService{}
	r.Options = opts
	return
}

// Patches a rule for a waiting room.
func (r *WaitingRoomRuleService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, body WaitingRoomRuleUpdateParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a rule for a waiting room.
func (r *WaitingRoomRuleService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, opts ...option.RequestOption) (res *[]WaitingRoomRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *WaitingRoomRuleService) WaitingRoomNewWaitingRoomRule(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists rules for a waiting room.
func (r *WaitingRoomRuleService) WaitingRoomListWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *[]WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *WaitingRoomRuleService) WaitingRoomReplaceWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingRoomRuleUpdateResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleUpdateResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                            `json:"version"`
	JSON    waitingRoomRuleUpdateResponseJSON `json:"-"`
}

// waitingRoomRuleUpdateResponseJSON contains the JSON metadata for the struct
// [WaitingRoomRuleUpdateResponse]
type waitingRoomRuleUpdateResponseJSON struct {
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

func (r *WaitingRoomRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingRoomRuleUpdateResponseAction string

const (
	WaitingRoomRuleUpdateResponseActionBypassWaitingRoom WaitingRoomRuleUpdateResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleDeleteResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleDeleteResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                            `json:"version"`
	JSON    waitingRoomRuleDeleteResponseJSON `json:"-"`
}

// waitingRoomRuleDeleteResponseJSON contains the JSON metadata for the struct
// [WaitingRoomRuleDeleteResponse]
type waitingRoomRuleDeleteResponseJSON struct {
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

func (r *WaitingRoomRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingRoomRuleDeleteResponseAction string

const (
	WaitingRoomRuleDeleteResponseActionBypassWaitingRoom WaitingRoomRuleDeleteResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                   `json:"version"`
	JSON    waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON contains the JSON
// metadata for the struct [WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse]
type waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON struct {
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

func (r *WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseAction string

const (
	WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseActionBypassWaitingRoom WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                     `json:"version"`
	JSON    waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON contains the JSON
// metadata for the struct [WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse]
type waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON struct {
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

func (r *WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseAction string

const (
	WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseActionBypassWaitingRoom WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                        `json:"version"`
	JSON    waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON contains the JSON
// metadata for the struct
// [WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse]
type waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON struct {
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

func (r *WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseAction string

const (
	WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseActionBypassWaitingRoom WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleUpdateParams struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleUpdateParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[WaitingRoomRuleUpdateParamsPosition] `json:"position"`
}

func (r WaitingRoomRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleUpdateParamsAction string

const (
	WaitingRoomRuleUpdateParamsActionBypassWaitingRoom WaitingRoomRuleUpdateParamsAction = "bypass_waiting_room"
)

// Reorder the position of a rule
//
// Satisfied by [WaitingRoomRuleUpdateParamsPositionObject],
// [WaitingRoomRuleUpdateParamsPositionObject],
// [WaitingRoomRuleUpdateParamsPositionObject].
type WaitingRoomRuleUpdateParamsPosition interface {
	implementsWaitingRoomRuleUpdateParamsPosition()
}

type WaitingRoomRuleUpdateParamsPositionObject struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r WaitingRoomRuleUpdateParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WaitingRoomRuleUpdateParamsPositionObject) implementsWaitingRoomRuleUpdateParamsPosition() {}

type WaitingRoomRuleUpdateResponseEnvelope struct {
	Errors   []WaitingRoomRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleUpdateResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleUpdateResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleUpdateResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomRuleUpdateResponseEnvelope]
type waitingRoomRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    waitingRoomRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WaitingRoomRuleUpdateResponseEnvelopeErrors]
type waitingRoomRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    waitingRoomRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomRuleUpdateResponseEnvelopeMessages]
type waitingRoomRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleUpdateResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleUpdateResponseEnvelopeSuccessTrue WaitingRoomRuleUpdateResponseEnvelopeSuccess = true
)

type WaitingRoomRuleUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       waitingRoomRuleUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [WaitingRoomRuleUpdateResponseEnvelopeResultInfo]
type waitingRoomRuleUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleDeleteResponseEnvelope struct {
	Errors   []WaitingRoomRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleDeleteResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleDeleteResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomRuleDeleteResponseEnvelope]
type waitingRoomRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    waitingRoomRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [WaitingRoomRuleDeleteResponseEnvelopeErrors]
type waitingRoomRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    waitingRoomRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomRuleDeleteResponseEnvelopeMessages]
type waitingRoomRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleDeleteResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleDeleteResponseEnvelopeSuccessTrue WaitingRoomRuleDeleteResponseEnvelopeSuccess = true
)

type WaitingRoomRuleDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       waitingRoomRuleDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleDeleteResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [WaitingRoomRuleDeleteResponseEnvelopeResultInfo]
type waitingRoomRuleDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction string

const (
	WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsActionBypassWaitingRoom WaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction = "bypass_waiting_room"
)

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelope struct {
	Errors   []WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelope]
type waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrors struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrors]
type waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessages struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessages]
type waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeSuccessTrue WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeSuccess = true
)

type WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                    `json:"total_count"`
	JSON       waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfo]
type waitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelope struct {
	Errors   []WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelope]
type waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrors struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrors]
type waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessages struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessages]
type waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeSuccessTrue WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeSuccess = true
)

type WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                      `json:"total_count"`
	JSON       waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfo]
type waitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams struct {
	Body param.Field[[]WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody] `json:"body,required"`
}

func (r WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction string

const (
	WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyActionBypassWaitingRoom WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction = "bypass_waiting_room"
)

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelope struct {
	Errors   []WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelope]
type waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrors]
type waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessages]
type waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeSuccessTrue WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeSuccess = true
)

type WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfo]
type waitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
