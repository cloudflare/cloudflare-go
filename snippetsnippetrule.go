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

// Put Rules
func (r *SnippetSnippetRuleService) Update(ctx context.Context, zoneIdentifier string, body SnippetSnippetRuleUpdateParams, opts ...option.RequestOption) (res *SnippetSnippetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Rules
func (r *SnippetSnippetRuleService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SnippetSnippetRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SnippetSnippetRuleUpdateResponse struct {
	Errors   []SnippetSnippetRuleUpdateResponseError   `json:"errors"`
	Messages []SnippetSnippetRuleUpdateResponseMessage `json:"messages"`
	// List of snippet rules
	Result []SnippetSnippetRuleUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success SnippetSnippetRuleUpdateResponseSuccess `json:"success"`
	JSON    snippetSnippetRuleUpdateResponseJSON    `json:"-"`
}

// snippetSnippetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [SnippetSnippetRuleUpdateResponse]
type snippetSnippetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    snippetSnippetRuleUpdateResponseErrorJSON `json:"-"`
}

// snippetSnippetRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleUpdateResponseError]
type snippetSnippetRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    snippetSnippetRuleUpdateResponseMessageJSON `json:"-"`
}

// snippetSnippetRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleUpdateResponseMessage]
type snippetSnippetRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleUpdateResponseResult struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                                     `json:"snippet_name"`
	JSON        snippetSnippetRuleUpdateResponseResultJSON `json:"-"`
}

// snippetSnippetRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleUpdateResponseResult]
type snippetSnippetRuleUpdateResponseResultJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetSnippetRuleUpdateResponseSuccess bool

const (
	SnippetSnippetRuleUpdateResponseSuccessTrue SnippetSnippetRuleUpdateResponseSuccess = true
)

type SnippetSnippetRuleListResponse struct {
	Errors   []SnippetSnippetRuleListResponseError   `json:"errors"`
	Messages []SnippetSnippetRuleListResponseMessage `json:"messages"`
	// List of snippet rules
	Result []SnippetSnippetRuleListResponseResult `json:"result"`
	// Whether the API call was successful
	Success SnippetSnippetRuleListResponseSuccess `json:"success"`
	JSON    snippetSnippetRuleListResponseJSON    `json:"-"`
}

// snippetSnippetRuleListResponseJSON contains the JSON metadata for the struct
// [SnippetSnippetRuleListResponse]
type snippetSnippetRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    snippetSnippetRuleListResponseErrorJSON `json:"-"`
}

// snippetSnippetRuleListResponseErrorJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleListResponseError]
type snippetSnippetRuleListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    snippetSnippetRuleListResponseMessageJSON `json:"-"`
}

// snippetSnippetRuleListResponseMessageJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleListResponseMessage]
type snippetSnippetRuleListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetSnippetRuleListResponseResult struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                                   `json:"snippet_name"`
	JSON        snippetSnippetRuleListResponseResultJSON `json:"-"`
}

// snippetSnippetRuleListResponseResultJSON contains the JSON metadata for the
// struct [SnippetSnippetRuleListResponseResult]
type snippetSnippetRuleListResponseResultJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetSnippetRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetSnippetRuleListResponseSuccess bool

const (
	SnippetSnippetRuleListResponseSuccessTrue SnippetSnippetRuleListResponseSuccess = true
)

type SnippetSnippetRuleUpdateParams struct {
	// List of snippet rules
	Rules param.Field[[]SnippetSnippetRuleUpdateParamsRule] `json:"rules"`
}

func (r SnippetSnippetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetSnippetRuleUpdateParamsRule struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Snippet identifying name
	SnippetName param.Field[string] `json:"snippet_name"`
}

func (r SnippetSnippetRuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
