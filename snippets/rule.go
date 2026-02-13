// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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

// Updates all snippet rules belonging to the zone.
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *RuleUpdateResponse, err error) {
	var env RuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all snippet rules belonging to the zone.
func (r *RuleService) List(ctx context.Context, query RuleListParams, opts ...option.RequestOption) (res *RuleListResponse, err error) {
	var env RuleListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all snippet rules belonging to the zone.
func (r *RuleService) Delete(ctx context.Context, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	var env RuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RuleUpdateResponse = interface{}

type RuleListResponse = interface{}

type RuleDeleteResponse = interface{}

type RuleUpdateParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Lists snippet rules.
	Rules param.Field[[]RuleUpdateParamsRule] `json:"rules,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define a snippet rule.
type RuleUpdateParamsRule struct {
	// Define the expression that determines which traffic matches the rule.
	Expression param.Field[string] `json:"expression,required"`
	// Identify the snippet.
	SnippetName param.Field[string] `json:"snippet_name,required"`
	// Provide an informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Indicate whether to execute the rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Return all API responses using this object.
type RuleUpdateResponseEnvelope struct {
	// Lists error messages.
	Errors []RuleUpdateResponseEnvelopeErrors `json:"errors,required"`
	// Contain warning messages.
	Messages []RuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Contain the response result.
	Result RuleUpdateResponse `json:"result,required"`
	// Indicate whether the API call was successful.
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Describes an API message.
type RuleUpdateResponseEnvelopeErrors struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                                `json:"code"`
	JSON ruleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeErrors]
type ruleUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Describes an API message.
type RuleUpdateResponseEnvelopeMessages struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                                  `json:"code"`
	JSON ruleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelopeMessages]
type ruleUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
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
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// Return all API responses using this object.
type RuleListResponseEnvelope struct {
	// Lists error messages.
	Errors []RuleListResponseEnvelopeErrors `json:"errors,required"`
	// Contain warning messages.
	Messages []RuleListResponseEnvelopeMessages `json:"messages,required"`
	// Contain the response result.
	Result RuleListResponse `json:"result,required"`
	// Indicate whether the API call was successful.
	Success RuleListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleListResponseEnvelopeJSON    `json:"-"`
}

// ruleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelope]
type ruleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Describes an API message.
type RuleListResponseEnvelopeErrors struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                              `json:"code"`
	JSON ruleListResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelopeErrors]
type ruleListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Describes an API message.
type RuleListResponseEnvelopeMessages struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                                `json:"code"`
	JSON ruleListResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleListResponseEnvelopeMessages]
type ruleListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type RuleListResponseEnvelopeSuccess bool

const (
	RuleListResponseEnvelopeSuccessTrue RuleListResponseEnvelopeSuccess = true
)

func (r RuleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// Use this field to specify the unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// Return all API responses using this object.
type RuleDeleteResponseEnvelope struct {
	// Lists error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// Contain warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Contain the response result.
	Result RuleDeleteResponse `json:"result,required"`
	// Indicate whether the API call was successful.
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

// Describes an API message.
type RuleDeleteResponseEnvelopeErrors struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                                `json:"code"`
	JSON ruleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Describes an API message.
type RuleDeleteResponseEnvelopeMessages struct {
	// Describes the message text.
	Message string `json:"message,required"`
	// Identify the message code.
	Code int64                                  `json:"code"`
	JSON ruleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
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
