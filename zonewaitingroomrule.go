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

// ZoneWaitingRoomRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneWaitingRoomRuleService]
// method instead.
type ZoneWaitingRoomRuleService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomRuleService(opts ...option.RequestOption) (r *ZoneWaitingRoomRuleService) {
	r = &ZoneWaitingRoomRuleService{}
	r.Options = opts
	return
}

// Patches a rule for a waiting room.
func (r *ZoneWaitingRoomRuleService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, body ZoneWaitingRoomRuleUpdateParams, opts ...option.RequestOption) (res *RulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a rule for a waiting room.
func (r *ZoneWaitingRoomRuleService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, opts ...option.RequestOption) (res *RulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomNewWaitingRoomRule(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams, opts ...option.RequestOption) (res *RulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists rules for a waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomListWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *RulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomReplaceWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams, opts ...option.RequestOption) (res *RulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type RulesResponseCollection struct {
	Errors     []RulesResponseCollectionError    `json:"errors"`
	Messages   []RulesResponseCollectionMessage  `json:"messages"`
	Result     []RulesResponseCollectionResult   `json:"result"`
	ResultInfo RulesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success RulesResponseCollectionSuccess `json:"success"`
	JSON    rulesResponseCollectionJSON    `json:"-"`
}

// rulesResponseCollectionJSON contains the JSON metadata for the struct
// [RulesResponseCollection]
type rulesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollectionError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    rulesResponseCollectionErrorJSON `json:"-"`
}

// rulesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [RulesResponseCollectionError]
type rulesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollectionMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    rulesResponseCollectionMessageJSON `json:"-"`
}

// rulesResponseCollectionMessageJSON contains the JSON metadata for the struct
// [RulesResponseCollectionMessage]
type rulesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesResponseCollectionResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RulesResponseCollectionResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                            `json:"version"`
	JSON    rulesResponseCollectionResultJSON `json:"-"`
}

// rulesResponseCollectionResultJSON contains the JSON metadata for the struct
// [RulesResponseCollectionResult]
type rulesResponseCollectionResultJSON struct {
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

func (r *RulesResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type RulesResponseCollectionResultAction string

const (
	RulesResponseCollectionResultActionBypassWaitingRoom RulesResponseCollectionResultAction = "bypass_waiting_room"
)

type RulesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       rulesResponseCollectionResultInfoJSON `json:"-"`
}

// rulesResponseCollectionResultInfoJSON contains the JSON metadata for the struct
// [RulesResponseCollectionResultInfo]
type rulesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RulesResponseCollectionSuccess bool

const (
	RulesResponseCollectionSuccessTrue RulesResponseCollectionSuccess = true
)

type ZoneWaitingRoomRuleUpdateParams struct {
	// The action to take when the expression matches.
	Action param.Field[ZoneWaitingRoomRuleUpdateParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[ZoneWaitingRoomRuleUpdateParamsPosition] `json:"position"`
}

func (r ZoneWaitingRoomRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleUpdateParamsAction string

const (
	ZoneWaitingRoomRuleUpdateParamsActionBypassWaitingRoom ZoneWaitingRoomRuleUpdateParamsAction = "bypass_waiting_room"
)

// Reorder the position of a rule
//
// Satisfied by [ZoneWaitingRoomRuleUpdateParamsPositionObject],
// [ZoneWaitingRoomRuleUpdateParamsPositionObject],
// [ZoneWaitingRoomRuleUpdateParamsPositionObject].
type ZoneWaitingRoomRuleUpdateParamsPosition interface {
	implementsZoneWaitingRoomRuleUpdateParamsPosition()
}

type ZoneWaitingRoomRuleUpdateParamsPositionObject struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r ZoneWaitingRoomRuleUpdateParamsPositionObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneWaitingRoomRuleUpdateParamsPositionObject) implementsZoneWaitingRoomRuleUpdateParamsPosition() {
}

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams struct {
	// The action to take when the expression matches.
	Action param.Field[ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction string

const (
	ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsActionBypassWaitingRoom ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParamsAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams struct {
	Body param.Field[[]ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody] `json:"body,required"`
}

func (r ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody struct {
	// The action to take when the expression matches.
	Action param.Field[ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction string

const (
	ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyActionBypassWaitingRoom ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParamsBodyAction = "bypass_waiting_room"
)
