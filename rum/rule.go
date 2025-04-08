// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rum

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
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

// Creates a new rule in a Web Analytics ruleset.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RUMRule, err error) {
	var env RuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", params.AccountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *RuleService) Update(ctx context.Context, rulesetID string, ruleID string, params RuleUpdateParams, opts ...option.RequestOption) (res *RUMRule, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", params.AccountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *RuleService) List(ctx context.Context, rulesetID string, query RuleListParams, opts ...option.RequestOption) (res *RuleListResponse, err error) {
	var env RuleListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", query.AccountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", body.AccountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modifies one or more rules in a Web Analytics ruleset with a single request.
func (r *RuleService) BulkNew(ctx context.Context, rulesetID string, params RuleBulkNewParams, opts ...option.RequestOption) (res *RuleBulkNewResponse, err error) {
	var env RuleBulkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", params.AccountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RUMRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string    `json:"paths"`
	Priority float64     `json:"priority"`
	JSON     rumRuleJSON `json:"-"`
}

// rumRuleJSON contains the JSON metadata for the struct [RUMRule]
type rumRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RUMRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rumRuleJSON) RawJSON() string {
	return r.raw
}

type RuleListResponse struct {
	// A list of rules.
	Rules   []RUMRule               `json:"rules"`
	Ruleset RuleListResponseRuleset `json:"ruleset"`
	JSON    ruleListResponseJSON    `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

type RuleListResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                      `json:"zone_tag"`
	JSON    ruleListResponseRulesetJSON `json:"-"`
}

// ruleListResponseRulesetJSON contains the JSON metadata for the struct
// [RuleListResponseRuleset]
type ruleListResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponse struct {
	// The Web Analytics rule identifier.
	ID   string                 `json:"id"`
	JSON ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RuleBulkNewResponse struct {
	// A list of rules.
	Rules   []RUMRule                  `json:"rules"`
	Ruleset RuleBulkNewResponseRuleset `json:"ruleset"`
	JSON    ruleBulkNewResponseJSON    `json:"-"`
}

// ruleBulkNewResponseJSON contains the JSON metadata for the struct
// [RuleBulkNewResponse]
type ruleBulkNewResponseJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type RuleBulkNewResponseRuleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string                         `json:"zone_tag"`
	JSON    ruleBulkNewResponseRulesetJSON `json:"-"`
}

// ruleBulkNewResponseRulesetJSON contains the JSON metadata for the struct
// [RuleBulkNewResponseRuleset]
type ruleBulkNewResponseRulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkNewResponseRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkNewResponseRulesetJSON) RawJSON() string {
	return r.raw
}

type RuleNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Host      param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewResponseEnvelope struct {
	Errors   []RuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                        `json:"success,required"`
	Result  RUMRule                     `json:"result"`
	JSON    ruleNewResponseEnvelopeJSON `json:"-"`
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

type RuleNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    ruleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    ruleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Host      param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []RuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                           `json:"success,required"`
	Result  RUMRule                        `json:"result"`
	JSON    ruleUpdateResponseEnvelopeJSON `json:"-"`
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

type RuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeErrors]
type ruleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeMessages]
type ruleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type RuleListResponseEnvelope struct {
	Errors   []RuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                         `json:"success,required"`
	Result  RuleListResponse             `json:"result"`
	JSON    ruleListResponseEnvelopeJSON `json:"-"`
}

// ruleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelope]
type ruleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    ruleListResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelopeErrors]
type ruleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleListResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelopeMessages]
type ruleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []RuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                           `json:"success,required"`
	Result  RuleDeleteResponse             `json:"result"`
	JSON    ruleDeleteResponseEnvelopeJSON `json:"-"`
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

type RuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RuleBulkNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of rule identifiers to delete.
	DeleteRules param.Field[[]string] `json:"delete_rules"`
	// A list of rules to create or update.
	Rules param.Field[[]RuleBulkNewParamsRule] `json:"rules"`
}

func (r RuleBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkNewParamsRule struct {
	// The Web Analytics rule identifier.
	ID        param.Field[string]   `json:"id"`
	Host      param.Field[string]   `json:"host"`
	Inclusive param.Field[bool]     `json:"inclusive"`
	IsPaused  param.Field[bool]     `json:"is_paused"`
	Paths     param.Field[[]string] `json:"paths"`
}

func (r RuleBulkNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkNewResponseEnvelope struct {
	Errors   []RuleBulkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RuleBulkNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                            `json:"success,required"`
	Result  RuleBulkNewResponse             `json:"result"`
	JSON    ruleBulkNewResponseEnvelopeJSON `json:"-"`
}

// ruleBulkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleBulkNewResponseEnvelope]
type ruleBulkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RuleBulkNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    ruleBulkNewResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleBulkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleBulkNewResponseEnvelopeErrors]
type ruleBulkNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RuleBulkNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    ruleBulkNewResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleBulkNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RuleBulkNewResponseEnvelopeMessages]
type ruleBulkNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleBulkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleBulkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
