// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetRuleService] method
// instead.
type RulesetRuleService struct {
	Options []option.RequestOption
}

// NewRulesetRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRulesetRuleService(opts ...option.RequestOption) (r *RulesetRuleService) {
	r = &RulesetRuleService{}
	r.Options = opts
	return
}

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RulesetRuleService) New(ctx context.Context, rulesetID string, params RulesetRuleNewParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleNewResponseEnvelope
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
func (r *RulesetRuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RulesetRuleDeleteParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleDeleteResponseEnvelope
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
func (r *RulesetRuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RulesetRuleEditParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetRuleEditResponseEnvelope
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

type RulesetRuleNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RulesetRuleNewParamsPosition] `json:"position"`
}

func (r RulesetRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [RulesetRuleNewParamsPositionBeforePosition],
// [RulesetRuleNewParamsPositionAfterPosition],
// [RulesetRuleNewParamsPositionIndexPosition].
type RulesetRuleNewParamsPosition interface {
	implementsRulesetRuleNewParamsPosition()
}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RulesetRuleNewParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionBeforePosition) implementsRulesetRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RulesetRuleNewParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionAfterPosition) implementsRulesetRuleNewParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleNewParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RulesetRuleNewParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleNewParamsPositionIndexPosition) implementsRulesetRuleNewParamsPosition() {}

// A response object.
type RulesetRuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleNewResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleNewResponseEnvelope]
type rulesetRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleNewResponseEnvelopeErrors]
type rulesetRuleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                         `json:"pointer,required"`
	JSON    rulesetRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseEnvelopeErrorsSource]
type rulesetRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleNewResponseEnvelopeMessages]
type rulesetRuleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                           `json:"pointer,required"`
	JSON    rulesetRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetRuleNewResponseEnvelopeMessagesSource]
type rulesetRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleNewResponseEnvelopeSuccess bool

const (
	RulesetRuleNewResponseEnvelopeSuccessTrue RulesetRuleNewResponseEnvelopeSuccess = true
)

type RulesetRuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RulesetRuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleDeleteResponseEnvelope]
type rulesetRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleDeleteResponseEnvelopeErrors]
type rulesetRuleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleDeleteResponseEnvelopeErrorsSource]
type rulesetRuleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleDeleteResponseEnvelopeMessages]
type rulesetRuleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RulesetRuleDeleteResponseEnvelopeMessagesSource]
type rulesetRuleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleDeleteResponseEnvelopeSuccess bool

const (
	RulesetRuleDeleteResponseEnvelopeSuccessTrue RulesetRuleDeleteResponseEnvelopeSuccess = true
)

type RulesetRuleEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RulesetRuleEditParamsPosition] `json:"position"`
}

func (r RulesetRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
//
// Satisfied by [RulesetRuleEditParamsPositionBeforePosition],
// [RulesetRuleEditParamsPositionAfterPosition],
// [RulesetRuleEditParamsPositionIndexPosition].
type RulesetRuleEditParamsPosition interface {
	implementsRulesetRuleEditParamsPosition()
}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RulesetRuleEditParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionBeforePosition) implementsRulesetRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RulesetRuleEditParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionAfterPosition) implementsRulesetRuleEditParamsPosition() {}

// An object configuring where the rule will be placed.
type RulesetRuleEditParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RulesetRuleEditParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetRuleEditParamsPositionIndexPosition) implementsRulesetRuleEditParamsPosition() {}

// A response object.
type RulesetRuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetRuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetRuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetRuleEditResponseEnvelopeJSON    `json:"-"`
}

// rulesetRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RulesetRuleEditResponseEnvelope]
type rulesetRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetRuleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RulesetRuleEditResponseEnvelopeErrors]
type rulesetRuleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    rulesetRuleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetRuleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseEnvelopeErrorsSource]
type rulesetRuleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetRuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetRuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetRuleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RulesetRuleEditResponseEnvelopeMessages]
type rulesetRuleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetRuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    rulesetRuleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetRuleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [RulesetRuleEditResponseEnvelopeMessagesSource]
type rulesetRuleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetRuleEditResponseEnvelopeSuccess bool

const (
	RulesetRuleEditResponseEnvelopeSuccessTrue RulesetRuleEditResponseEnvelopeSuccess = true
)
