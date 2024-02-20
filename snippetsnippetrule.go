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

// SnippetSnippetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSnippetSnippetRuleService] method
// instead.
type SnippetSnippetRuleService struct {
	Options []option.RequestOption
}

// NewSnippetSnippetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSnippetSnippetRuleService(opts ...option.RequestOption) (r *SnippetSnippetRuleService) {
	r = &SnippetSnippetRuleService{}
	r.Options = opts
	return
}

// Rules
func (r *SnippetSnippetRuleService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]SnippetSnippetRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetSnippetRuleListResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Put Rules
func (r *SnippetSnippetRuleService) Replace(ctx context.Context, zoneIdentifier string, body SnippetSnippetRuleReplaceParams, opts ...option.RequestOption) (res *[]SnippetSnippetRuleReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetSnippetRuleReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SnippetSnippetRuleListResponse struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                             `json:"snippet_name"`
	JSON        snippetSnippetRuleListResponseJSON `json:"-"`
}

// snippetSnippetRuleListResponseJSON contains the JSON metadata for the struct
// [SnippetSnippetRuleListResponse]
type snippetSnippetRuleListResponseJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleReplaceResponse struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                                `json:"snippet_name"`
	JSON        snippetSnippetRuleReplaceResponseJSON `json:"-"`
}

// snippetSnippetRuleReplaceResponseJSON contains the JSON metadata for the struct
// [SnippetSnippetRuleReplaceResponse]
type snippetSnippetRuleReplaceResponseJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseEnvelope struct {
	Errors   []SnippetSnippetRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetSnippetRuleListResponseEnvelopeMessages `json:"messages,required"`
	// List of snippet rules
	Result []SnippetSnippetRuleListResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetSnippetRuleListResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetSnippetRuleListResponseEnvelopeJSON    `json:"-"`
}

// snippetSnippetRuleListResponseEnvelopeJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleListResponseEnvelope]
type snippetSnippetRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    snippetSnippetRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetSnippetRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SnippetSnippetRuleListResponseEnvelopeErrors]
type snippetSnippetRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    snippetSnippetRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetSnippetRuleListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SnippetSnippetRuleListResponseEnvelopeMessages]
type snippetSnippetRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetSnippetRuleListResponseEnvelopeSuccess bool

const (
	SnippetSnippetRuleListResponseEnvelopeSuccessTrue SnippetSnippetRuleListResponseEnvelopeSuccess = true
)

type SnippetSnippetRuleReplaceParams struct {
	// List of snippet rules
	Rules param.Field[[]SnippetSnippetRuleReplaceParamsRule] `json:"rules"`
}

func (r SnippetSnippetRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetSnippetRuleReplaceParamsRule struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Snippet identifying name
	SnippetName param.Field[string] `json:"snippet_name"`
}

func (r SnippetSnippetRuleReplaceParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetSnippetRuleReplaceResponseEnvelope struct {
	Errors   []SnippetSnippetRuleReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetSnippetRuleReplaceResponseEnvelopeMessages `json:"messages,required"`
	// List of snippet rules
	Result []SnippetSnippetRuleReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetSnippetRuleReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetSnippetRuleReplaceResponseEnvelopeJSON    `json:"-"`
}

// snippetSnippetRuleReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleReplaceResponseEnvelope]
type snippetSnippetRuleReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleReplaceResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    snippetSnippetRuleReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetSnippetRuleReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SnippetSnippetRuleReplaceResponseEnvelopeErrors]
type snippetSnippetRuleReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleReplaceResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    snippetSnippetRuleReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetSnippetRuleReplaceResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SnippetSnippetRuleReplaceResponseEnvelopeMessages]
type snippetSnippetRuleReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetSnippetRuleReplaceResponseEnvelopeSuccess bool

const (
	SnippetSnippetRuleReplaceResponseEnvelopeSuccessTrue SnippetSnippetRuleReplaceResponseEnvelopeSuccess = true
)
