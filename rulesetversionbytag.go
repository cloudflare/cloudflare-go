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

// RulesetVersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRulesetVersionByTagService]
// method instead.
type RulesetVersionByTagService struct {
	Options []option.RequestOption
}

// NewRulesetVersionByTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRulesetVersionByTagService(opts ...option.RequestOption) (r *RulesetVersionByTagService) {
	r = &RulesetVersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *RulesetVersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query RulesetVersionByTagGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RulesetVersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RulesetVersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type RulesetVersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetVersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetVersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RulesetVersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RulesetVersionByTagGetResponseEnvelope]
type rulesetVersionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   rulesetVersionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RulesetVersionByTagGetResponseEnvelopeErrors]
type rulesetVersionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseEnvelopeErrorsSource]
type rulesetVersionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A message.
type RulesetVersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RulesetVersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   rulesetVersionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [RulesetVersionByTagGetResponseEnvelopeMessages]
type rulesetVersionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The source of this message.
type RulesetVersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RulesetVersionByTagGetResponseEnvelopeMessagesSource]
type rulesetVersionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetVersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type RulesetVersionByTagGetResponseEnvelopeSuccess bool

const (
	RulesetVersionByTagGetResponseEnvelopeSuccessTrue RulesetVersionByTagGetResponseEnvelopeSuccess = true
)
