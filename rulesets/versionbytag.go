// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// VersionByTagService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVersionByTagService] method
// instead.
type VersionByTagService struct {
	Options []option.RequestOption
}

// NewVersionByTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVersionByTagService(opts ...option.RequestOption) (r *VersionByTagService) {
	r = &VersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *VersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query VersionByTagGetParams, opts ...option.RequestOption) (res *RulesetsRulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionByTagGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type VersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result RulesetsRulesetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// versionByTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseEnvelope]
type versionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeErrors]
type versionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeErrorsSource]
type versionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeMessages]
type versionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeMessagesSource]
type versionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionByTagGetResponseEnvelopeSuccess bool

const (
	VersionByTagGetResponseEnvelopeSuccessTrue VersionByTagGetResponseEnvelopeSuccess = true
)
