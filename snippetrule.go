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

// SnippetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSnippetRuleService] method
// instead.
type SnippetRuleService struct {
	Options []option.RequestOption
}

// NewSnippetRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnippetRuleService(opts ...option.RequestOption) (r *SnippetRuleService) {
	r = &SnippetRuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *SnippetRuleService) Update(ctx context.Context, zoneIdentifier string, body SnippetRuleUpdateParams, opts ...option.RequestOption) (res *[]SnippetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Rules
func (r *SnippetRuleService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]SnippetRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetRuleListResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SnippetRuleUpdateResponse struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                        `json:"snippet_name"`
	JSON        snippetRuleUpdateResponseJSON `json:"-"`
}

// snippetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [SnippetRuleUpdateResponse]
type snippetRuleUpdateResponseJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleListResponse struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                      `json:"snippet_name"`
	JSON        snippetRuleListResponseJSON `json:"-"`
}

// snippetRuleListResponseJSON contains the JSON metadata for the struct
// [SnippetRuleListResponse]
type snippetRuleListResponseJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleUpdateParams struct {
	// List of snippet rules
	Rules param.Field[[]SnippetRuleUpdateParamsRule] `json:"rules"`
}

func (r SnippetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetRuleUpdateParamsRule struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Snippet identifying name
	SnippetName param.Field[string] `json:"snippet_name"`
}

func (r SnippetRuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetRuleUpdateResponseEnvelope struct {
	Errors   []SnippetRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// List of snippet rules
	Result []SnippetRuleUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// snippetRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetRuleUpdateResponseEnvelope]
type snippetRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    snippetRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetRuleUpdateResponseEnvelopeErrors]
type snippetRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    snippetRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetRuleUpdateResponseEnvelopeMessages]
type snippetRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetRuleUpdateResponseEnvelopeSuccess bool

const (
	SnippetRuleUpdateResponseEnvelopeSuccessTrue SnippetRuleUpdateResponseEnvelopeSuccess = true
)

type SnippetRuleListResponseEnvelope struct {
	Errors   []SnippetRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetRuleListResponseEnvelopeMessages `json:"messages,required"`
	// List of snippet rules
	Result []SnippetRuleListResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetRuleListResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetRuleListResponseEnvelopeJSON    `json:"-"`
}

// snippetRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetRuleListResponseEnvelope]
type snippetRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    snippetRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetRuleListResponseEnvelopeErrors]
type snippetRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SnippetRuleListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    snippetRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetRuleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetRuleListResponseEnvelopeMessages]
type snippetRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetRuleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SnippetRuleListResponseEnvelopeSuccess bool

const (
	SnippetRuleListResponseEnvelopeSuccessTrue SnippetRuleListResponseEnvelopeSuccess = true
)
