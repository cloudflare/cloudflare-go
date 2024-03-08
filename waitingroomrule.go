// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *WaitingRoomRuleService) New(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleNewParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleNewResponse, err error) {
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
func (r *WaitingRoomRuleService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body WaitingRoomRuleUpdateParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleUpdateResponse, err error) {
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
func (r *WaitingRoomRuleService) List(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *[]WaitingRoomRuleListResponse, err error) {
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

// Patches a rule for a waiting room.
func (r *WaitingRoomRuleService) Edit(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, body WaitingRoomRuleEditParams, opts ...option.RequestOption) (res *[]WaitingRoomRuleEditResponse, err error) {
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

type WaitingRoomRuleNewResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleNewResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                         `json:"version"`
	JSON    waitingRoomRuleNewResponseJSON `json:"-"`
}

// waitingRoomRuleNewResponseJSON contains the JSON metadata for the struct
// [WaitingRoomRuleNewResponse]
type waitingRoomRuleNewResponseJSON struct {
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

func (r *WaitingRoomRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleNewResponseAction string

const (
	WaitingRoomRuleNewResponseActionBypassWaitingRoom WaitingRoomRuleNewResponseAction = "bypass_waiting_room"
)

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

func (r waitingRoomRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleUpdateResponseAction string

const (
	WaitingRoomRuleUpdateResponseActionBypassWaitingRoom WaitingRoomRuleUpdateResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleListResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleListResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                          `json:"version"`
	JSON    waitingRoomRuleListResponseJSON `json:"-"`
}

// waitingRoomRuleListResponseJSON contains the JSON metadata for the struct
// [WaitingRoomRuleListResponse]
type waitingRoomRuleListResponseJSON struct {
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

func (r *WaitingRoomRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleListResponseAction string

const (
	WaitingRoomRuleListResponseActionBypassWaitingRoom WaitingRoomRuleListResponseAction = "bypass_waiting_room"
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

func (r waitingRoomRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleDeleteResponseAction string

const (
	WaitingRoomRuleDeleteResponseActionBypassWaitingRoom WaitingRoomRuleDeleteResponseAction = "bypass_waiting_room"
)

type WaitingRoomRuleEditResponse struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action WaitingRoomRuleEditResponseAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                          `json:"version"`
	JSON    waitingRoomRuleEditResponseJSON `json:"-"`
}

// waitingRoomRuleEditResponseJSON contains the JSON metadata for the struct
// [WaitingRoomRuleEditResponse]
type waitingRoomRuleEditResponseJSON struct {
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

func (r *WaitingRoomRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r waitingRoomRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type WaitingRoomRuleEditResponseAction string

const (
	WaitingRoomRuleEditResponseActionBypassWaitingRoom WaitingRoomRuleEditResponseAction = "bypass_waiting_room"
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
	Result   []WaitingRoomRuleNewResponse                 `json:"result,required,nullable"`
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

func (r waitingRoomRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type WaitingRoomRuleListResponseEnvelope struct {
	Errors   []WaitingRoomRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WaitingRoomRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WaitingRoomRuleListResponse                 `json:"result,required,nullable"`
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

func (r waitingRoomRuleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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
	Result   []WaitingRoomRuleEditResponse                 `json:"result,required,nullable"`
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

func (r waitingRoomRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r waitingRoomRuleEditResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
