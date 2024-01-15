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

// ZoneSnippetsRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSnippetsRuleService] method
// instead.
type ZoneSnippetsRuleService struct {
	Options []option.RequestOption
}

// NewZoneSnippetsRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSnippetsRuleService(opts ...option.RequestOption) (r *ZoneSnippetsRuleService) {
	r = &ZoneSnippetsRuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *ZoneSnippetsRuleService) Update(ctx context.Context, zoneIdentifier string, body ZoneSnippetsRuleUpdateParams, opts ...option.RequestOption) (res *ZoneSnippetsRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Rules
func (r *ZoneSnippetsRuleService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSnippetsRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/snippet_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSnippetsRuleUpdateResponse struct {
	Errors   []ZoneSnippetsRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneSnippetsRuleUpdateResponseMessage `json:"messages"`
	// List of snippet rules
	Result []ZoneSnippetsRuleUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneSnippetsRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneSnippetsRuleUpdateResponseJSON    `json:"-"`
}

// zoneSnippetsRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetsRuleUpdateResponse]
type zoneSnippetsRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSnippetsRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneSnippetsRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneSnippetsRuleUpdateResponseError]
type zoneSnippetsRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneSnippetsRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneSnippetsRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSnippetsRuleUpdateResponseMessage]
type zoneSnippetsRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleUpdateResponseResult struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                                   `json:"snippet_name"`
	JSON        zoneSnippetsRuleUpdateResponseResultJSON `json:"-"`
}

// zoneSnippetsRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneSnippetsRuleUpdateResponseResult]
type zoneSnippetsRuleUpdateResponseResultJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSnippetsRuleUpdateResponseSuccess bool

const (
	ZoneSnippetsRuleUpdateResponseSuccessTrue ZoneSnippetsRuleUpdateResponseSuccess = true
)

type ZoneSnippetsRuleListResponse struct {
	Errors   []ZoneSnippetsRuleListResponseError   `json:"errors"`
	Messages []ZoneSnippetsRuleListResponseMessage `json:"messages"`
	// List of snippet rules
	Result []ZoneSnippetsRuleListResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneSnippetsRuleListResponseSuccess `json:"success"`
	JSON    zoneSnippetsRuleListResponseJSON    `json:"-"`
}

// zoneSnippetsRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetsRuleListResponse]
type zoneSnippetsRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleListResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zoneSnippetsRuleListResponseErrorJSON `json:"-"`
}

// zoneSnippetsRuleListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetsRuleListResponseError]
type zoneSnippetsRuleListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleListResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneSnippetsRuleListResponseMessageJSON `json:"-"`
}

// zoneSnippetsRuleListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneSnippetsRuleListResponseMessage]
type zoneSnippetsRuleListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetsRuleListResponseResult struct {
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Snippet identifying name
	SnippetName string                                 `json:"snippet_name"`
	JSON        zoneSnippetsRuleListResponseResultJSON `json:"-"`
}

// zoneSnippetsRuleListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetsRuleListResponseResult]
type zoneSnippetsRuleListResponseResultJSON struct {
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetsRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSnippetsRuleListResponseSuccess bool

const (
	ZoneSnippetsRuleListResponseSuccessTrue ZoneSnippetsRuleListResponseSuccess = true
)

type ZoneSnippetsRuleUpdateParams struct {
	// List of snippet rules
	Rules param.Field[[]ZoneSnippetsRuleUpdateParamsRule] `json:"rules"`
}

func (r ZoneSnippetsRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSnippetsRuleUpdateParamsRule struct {
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Snippet identifying name
	SnippetName param.Field[string] `json:"snippet_name"`
}

func (r ZoneSnippetsRuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
