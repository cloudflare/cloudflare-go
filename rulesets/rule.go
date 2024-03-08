// File generated from our OpenAPI spec by Stainless.

package rulesets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsPosition] `json:"position"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsPositionBeforePosition],
// [rulesets.RuleNewParamsPositionAfterPosition],
// [rulesets.RuleNewParamsPositionIndexPosition].
type RuleNewParamsPosition interface {
	implementsRulesetsRuleNewParamsPosition()
}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionBeforePosition) implementsRulesetsRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionAfterPosition) implementsRulesetsRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionIndexPosition) implementsRulesetsRuleNewParamsPosition() {}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                  `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeErrorsSource]
type ruleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeMessagesSource]
type ruleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

type RuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeErrorsSource]
type ruleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeMessagesSource]
type ruleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

type RuleEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsPosition] `json:"position"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsPositionBeforePosition],
// [rulesets.RuleEditParamsPositionAfterPosition],
// [rulesets.RuleEditParamsPositionIndexPosition].
type RuleEditParamsPosition interface {
	implementsRulesetsRuleEditParamsPosition()
}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionBeforePosition) implementsRulesetsRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAfterPosition) implementsRulesetsRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionIndexPosition) implementsRulesetsRuleEditParamsPosition() {}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleEditResponseEnvelopeJSON    `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeErrorsSource]
type ruleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeMessagesSource]
type ruleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)
