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
func (r *ZoneWaitingRoomRuleService) Update(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, body ZoneWaitingRoomRuleUpdateParams, opts ...option.RequestOption) (res *ZoneWaitingRoomRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a rule for a waiting room.
func (r *ZoneWaitingRoomRuleService) Delete(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, ruleID string, opts ...option.RequestOption) (res *ZoneWaitingRoomRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules/%s", zoneIdentifier, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomNewWaitingRoomRule(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleParams, opts ...option.RequestOption) (res *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists rules for a waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomListWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, opts ...option.RequestOption) (res *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *ZoneWaitingRoomRuleService) WaitingRoomReplaceWaitingRoomRules(ctx context.Context, zoneIdentifier string, waitingRoomID interface{}, body ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesParams, opts ...option.RequestOption) (res *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/waiting_rooms/%v/rules", zoneIdentifier, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneWaitingRoomRuleUpdateResponse struct {
	Errors     []ZoneWaitingRoomRuleUpdateResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomRuleUpdateResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomRuleUpdateResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomRuleUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomRuleUpdateResponseJSON    `json:"-"`
}

// zoneWaitingRoomRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomRuleUpdateResponse]
type zoneWaitingRoomRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleUpdateResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneWaitingRoomRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleUpdateResponseError]
type zoneWaitingRoomRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleUpdateResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneWaitingRoomRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleUpdateResponseMessage]
type zoneWaitingRoomRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleUpdateResponseResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action ZoneWaitingRoomRuleUpdateResponseResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                      `json:"version"`
	JSON    zoneWaitingRoomRuleUpdateResponseResultJSON `json:"-"`
}

// zoneWaitingRoomRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleUpdateResponseResult]
type zoneWaitingRoomRuleUpdateResponseResultJSON struct {
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

func (r *ZoneWaitingRoomRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleUpdateResponseResultAction string

const (
	ZoneWaitingRoomRuleUpdateResponseResultActionBypassWaitingRoom ZoneWaitingRoomRuleUpdateResponseResultAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       zoneWaitingRoomRuleUpdateResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomRuleUpdateResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneWaitingRoomRuleUpdateResponseResultInfo]
type zoneWaitingRoomRuleUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomRuleUpdateResponseSuccess bool

const (
	ZoneWaitingRoomRuleUpdateResponseSuccessTrue ZoneWaitingRoomRuleUpdateResponseSuccess = true
)

type ZoneWaitingRoomRuleDeleteResponse struct {
	Errors     []ZoneWaitingRoomRuleDeleteResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomRuleDeleteResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomRuleDeleteResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomRuleDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomRuleDeleteResponseJSON    `json:"-"`
}

// zoneWaitingRoomRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneWaitingRoomRuleDeleteResponse]
type zoneWaitingRoomRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleDeleteResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneWaitingRoomRuleDeleteResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleDeleteResponseError]
type zoneWaitingRoomRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleDeleteResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneWaitingRoomRuleDeleteResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleDeleteResponseMessage]
type zoneWaitingRoomRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleDeleteResponseResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action ZoneWaitingRoomRuleDeleteResponseResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                      `json:"version"`
	JSON    zoneWaitingRoomRuleDeleteResponseResultJSON `json:"-"`
}

// zoneWaitingRoomRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneWaitingRoomRuleDeleteResponseResult]
type zoneWaitingRoomRuleDeleteResponseResultJSON struct {
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

func (r *ZoneWaitingRoomRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleDeleteResponseResultAction string

const (
	ZoneWaitingRoomRuleDeleteResponseResultActionBypassWaitingRoom ZoneWaitingRoomRuleDeleteResponseResultAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       zoneWaitingRoomRuleDeleteResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomRuleDeleteResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneWaitingRoomRuleDeleteResponseResultInfo]
type zoneWaitingRoomRuleDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomRuleDeleteResponseSuccess bool

const (
	ZoneWaitingRoomRuleDeleteResponseSuccessTrue ZoneWaitingRoomRuleDeleteResponseSuccess = true
)

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse struct {
	Errors     []ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON    `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse]
type zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseError]
type zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessage]
type zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                             `json:"version"`
	JSON    zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResult]
type zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultJSON struct {
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

func (r *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultAction string

const (
	ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultActionBypassWaitingRoom ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfoJSON contains
// the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfo]
type zoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseSuccess bool

const (
	ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseSuccessTrue ZoneWaitingRoomRuleWaitingRoomNewWaitingRoomRuleResponseSuccess = true
)

type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse struct {
	Errors     []ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON    `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON contains the JSON
// metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse]
type zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseError]
type zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessage]
type zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                               `json:"version"`
	JSON    zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResult]
type zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultJSON struct {
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

func (r *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultAction string

const (
	ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultActionBypassWaitingRoom ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfo]
type zoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseSuccess bool

const (
	ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseSuccessTrue ZoneWaitingRoomRuleWaitingRoomListWaitingRoomRulesResponseSuccess = true
)

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse struct {
	Errors     []ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseError    `json:"errors"`
	Messages   []ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessage  `json:"messages"`
	Result     []ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResult   `json:"result"`
	ResultInfo ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseSuccess `json:"success"`
	JSON    zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON    `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON contains the
// JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse]
type zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseError struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseErrorJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseError]
type zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessage struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessageJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessage]
type zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                                                  `json:"version"`
	JSON    zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResult]
type zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultJSON struct {
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

func (r *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to take when the expression matches.
type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultAction string

const (
	ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultActionBypassWaitingRoom ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultAction = "bypass_waiting_room"
)

type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                     `json:"total_count"`
	JSON       zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfoJSON `json:"-"`
}

// zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfo]
type zoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseSuccess bool

const (
	ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseSuccessTrue ZoneWaitingRoomRuleWaitingRoomReplaceWaitingRoomRulesResponseSuccess = true
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
