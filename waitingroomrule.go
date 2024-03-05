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

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *WaitingRoomRuleService) New(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleNewParams, opts ...option.RequestOption) (res *[]WaitingroomRuleResult, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *WaitingRoomRuleService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleUpdateParams, opts ...option.RequestOption) (res *[]WaitingroomRuleResult, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists rules for a waiting room.
func (r *WaitingRoomRuleService) List(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *[]WaitingroomRuleResult, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleListResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a rule for a waiting room.
func (r *WaitingRoomRuleService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, opts ...option.RequestOption) (res *[]WaitingroomRuleResult, err error) {
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

// Patches a rule for a waiting room.
func (r *WaitingRoomRuleService) Edit(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, body WaitingRoomRuleEditParams, opts ...option.RequestOption) (res *[]WaitingroomRuleResult, err error) {
	opts = append(r.Options[:], opts...)
	var env WaitingRoomRuleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WaitingroomRuleResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingroomRuleResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                    `json:"version"`
	JSON    waitingroomRuleResultJSON `json:"-"`
}

// waitingroomRuleResultJSON contains the JSON metadata for the struct
// [WaitingroomRuleResult]
type waitingroomRuleResultJSON struct {
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

func (r *WaitingroomRuleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type WaitingroomRuleResultAction string

const (
	WaitingroomRuleResultActionBypassWaitingRoom WaitingroomRuleResultAction = "bypass_waiting_room"
)

type WaitingRoomRuleNewParams struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleNewParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WaitingRoomRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleNewParamsAction string

const (
	WaitingRoomRuleNewParamsActionBypassWaitingRoom WaitingRoomRuleNewParamsAction = "bypass_waiting_room"
)

type WaitingRoomRuleNewResponseEnvelope struct {
	Errors   []WaitingRoomRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomRuleResult                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleNewResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WaitingRoomRuleNewResponseEnvelope]
type waitingRoomRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleNewResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    waitingRoomRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WaitingRoomRuleNewResponseEnvelopeErrors]
type waitingRoomRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleNewResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    waitingRoomRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomRuleNewResponseEnvelopeMessages]
type waitingRoomRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleNewResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleNewResponseEnvelopeSuccessTrue WaitingRoomRuleNewResponseEnvelopeSuccess = true
)

type WaitingRoomRuleNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       waitingRoomRuleNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleNewResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WaitingRoomRuleNewResponseEnvelopeResultInfo]
type waitingRoomRuleNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleUpdateParams struct {
	Body param.Field[[]WaitingRoomRuleUpdateParamsBody] `json:"body,required"`
}

func (r WaitingRoomRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type WaitingRoomRuleUpdateParamsBody struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleUpdateParamsBodyAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r WaitingRoomRuleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleUpdateParamsBodyAction string

const (
	WaitingRoomRuleUpdateParamsBodyActionBypassWaitingRoom WaitingRoomRuleUpdateParamsBodyAction = "bypass_waiting_room"
)

type WaitingRoomRuleUpdateResponseEnvelope struct {
	Errors   []WaitingRoomRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomRuleResult                         `json:"result,required,nullable"`
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

type WaitingRoomRuleListResponseEnvelope struct {
	Errors   []WaitingRoomRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomRuleResult                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleListResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleListResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomRuleListResponseEnvelope]
type waitingRoomRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleListResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    waitingRoomRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WaitingRoomRuleListResponseEnvelopeErrors]
type waitingRoomRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleListResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    waitingRoomRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomRuleListResponseEnvelopeMessages]
type waitingRoomRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleListResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleListResponseEnvelopeSuccessTrue WaitingRoomRuleListResponseEnvelopeSuccess = true
)

type WaitingRoomRuleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       waitingRoomRuleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WaitingRoomRuleListResponseEnvelopeResultInfo]
type waitingRoomRuleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleDeleteResponseEnvelope struct {
	Errors   []WaitingRoomRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomRuleResult                         `json:"result,required,nullable"`
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

type WaitingRoomRuleEditParams struct {
	// The action to take when the expression matches.
	Action param.Field[WaitingRoomRuleEditParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[WaitingRoomRuleEditParamsPosition] `json:"position"`
}

func (r WaitingRoomRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type WaitingRoomRuleEditParamsAction string

const (
	WaitingRoomRuleEditParamsActionBypassWaitingRoom WaitingRoomRuleEditParamsAction = "bypass_waiting_room"
)

// Reorder the position of a rule
//
// Satisfied by [WaitingRoomRuleEditParamsPositionObject],
// [WaitingRoomRuleEditParamsPositionObject],
// [WaitingRoomRuleEditParamsPositionObject].
type WaitingRoomRuleEditParamsPosition interface {
	implementsWaitingRoomRuleEditParamsPosition()
}

type WaitingRoomRuleEditParamsPositionObject struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r WaitingRoomRuleEditParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r WaitingRoomRuleEditParamsPositionObject) implementsWaitingRoomRuleEditParamsPosition() {}

type WaitingRoomRuleEditResponseEnvelope struct {
	Errors   []WaitingRoomRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingroomRuleResult                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    WaitingRoomRuleEditResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo WaitingRoomRuleEditResponseEnvelopeResultInfo `json:"result_info"`
	JSON       waitingRoomRuleEditResponseEnvelopeJSON       `json:"-"`
}

// waitingRoomRuleEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [WaitingRoomRuleEditResponseEnvelope]
type waitingRoomRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleEditResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    waitingRoomRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// waitingRoomRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WaitingRoomRuleEditResponseEnvelopeErrors]
type waitingRoomRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomRuleEditResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    waitingRoomRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// waitingRoomRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WaitingRoomRuleEditResponseEnvelopeMessages]
type waitingRoomRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WaitingRoomRuleEditResponseEnvelopeSuccess bool

const (
	WaitingRoomRuleEditResponseEnvelopeSuccessTrue WaitingRoomRuleEditResponseEnvelopeSuccess = true
)

type WaitingRoomRuleEditResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       waitingRoomRuleEditResponseEnvelopeResultInfoJSON `json:"-"`
}

// waitingRoomRuleEditResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [WaitingRoomRuleEditResponseEnvelopeResultInfo]
type waitingRoomRuleEditResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomRuleEditResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
