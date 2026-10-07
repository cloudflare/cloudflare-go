// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// DispatchNamespaceScriptTagService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceScriptTagService] method instead.
type DispatchNamespaceScriptTagService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptTagService(opts ...option.RequestOption) (r *DispatchNamespaceScriptTagService) {
	r = &DispatchNamespaceScriptTagService{}
	r.Options = opts
	return
}

// Replace tags for a script uploaded to a Workers for Platforms dispatch
// namespace.
func (r *DispatchNamespaceScriptTagService) Update(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptTagUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags", params.AccountID, dispatchNamespace, scriptName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Replace tags for a script uploaded to a Workers for Platforms dispatch
// namespace.
func (r *DispatchNamespaceScriptTagService) UpdateAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, params DispatchNamespaceScriptTagUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, dispatchNamespace, scriptName, params, opts...))
}

// Fetch tags from a script uploaded to a Workers for Platforms dispatch namespace.
func (r *DispatchNamespaceScriptTagService) List(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptTagListParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags", query.AccountID, dispatchNamespace, scriptName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch tags from a script uploaded to a Workers for Platforms dispatch namespace.
func (r *DispatchNamespaceScriptTagService) ListAutoPaging(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptTagListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, dispatchNamespace, scriptName, query, opts...))
}

// Delete a tag from a script uploaded to a Workers for Platforms dispatch
// namespace.
//
// On `api-version` dates on or after `2026-10-01`, `tag` identifies a key and the
// operation returns the complete updated tag map. Deleting a missing key succeeds.
// Earlier versions retain the legacy string-tag behavior and return a null result.
func (r *DispatchNamespaceScriptTagService) Delete(ctx context.Context, dispatchNamespace string, scriptName string, tag string, params DispatchNamespaceScriptTagDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceScriptTagDeleteResponse, err error) {
	if params.APIVersion.Present {
		opts = append(opts, option.WithHeader("api-version", fmt.Sprintf("%v", params.APIVersion)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return nil, err
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return nil, err
	}
	if tag == "" {
		err = errors.New("missing required tag parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/tags/%s", params.AccountID, dispatchNamespace, scriptName, tag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type DispatchNamespaceScriptTagDeleteResponse struct {
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseError],
	// [[]DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultError].
	Errors interface{} `json:"errors" api:"required"`
	// This field can have the runtime type of
	// [[]DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessage],
	// [[]DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessage].
	Messages interface{} `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptTagDeleteResponseSuccess `json:"success" api:"required"`
	// This field can have the runtime type of [map[string]string], [interface{}].
	Result interface{}                                  `json:"result"`
	JSON   dispatchNamespaceScriptTagDeleteResponseJSON `json:"-"`
	union  DispatchNamespaceScriptTagDeleteResponseUnion
}

// dispatchNamespaceScriptTagDeleteResponseJSON contains the JSON metadata for the
// struct [DispatchNamespaceScriptTagDeleteResponse]
type dispatchNamespaceScriptTagDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r dispatchNamespaceScriptTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *DispatchNamespaceScriptTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = DispatchNamespaceScriptTagDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [DispatchNamespaceScriptTagDeleteResponseUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse],
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult].
func (r DispatchNamespaceScriptTagDeleteResponse) AsUnion() DispatchNamespaceScriptTagDeleteResponseUnion {
	return r.union
}

// Union satisfied by
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse] or
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult].
type DispatchNamespaceScriptTagDeleteResponseUnion interface {
	implementsDispatchNamespaceScriptTagDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DispatchNamespaceScriptTagDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult{}),
		},
	)
}

type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse struct {
	Errors   []DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseError   `json:"errors" api:"required"`
	Messages []DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessage `json:"messages" api:"required"`
	// Tags associated with the Worker, as a key/value object. Each complete UTF-8
	// encoded tag (the key plus `=` and the value when the value is non-empty) must
	// not exceed 1024 bytes.
	Result map[string]string `json:"result" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccess `json:"success" api:"required"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseJSON contains the
// JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse]
type dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponse) implementsDispatchNamespaceScriptTagDeleteResponse() {
}

type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseError struct {
	Code             int64                                                                     `json:"code" api:"required"`
	Message          string                                                                    `json:"message" api:"required"`
	DocumentationURL string                                                                    `json:"documentation_url"`
	Source           DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorJSON contains
// the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseError]
type dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSource struct {
	Pointer string                                                                        `json:"pointer"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSourceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSource]
type dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessage struct {
	Code             int64                                                                       `json:"code" api:"required"`
	Message          string                                                                      `json:"message" api:"required"`
	DocumentationURL string                                                                      `json:"documentation_url"`
	Source           DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessageJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessageJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessage]
type dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessageJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSource struct {
	Pointer string                                                                          `json:"pointer"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSourceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSource]
type dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccess bool

const (
	DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccessTrue DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccess = true
)

func (r DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptTagDeleteResponseWorkersTagsKVResponseSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult struct {
	Errors   []DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultError   `json:"errors" api:"required"`
	Messages []DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessage `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccess `json:"success" api:"required"`
	Result  interface{}                                                                 `json:"result" api:"nullable"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult]
type dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultJSON) RawJSON() string {
	return r.raw
}

func (r DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResult) implementsDispatchNamespaceScriptTagDeleteResponse() {
}

type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultError struct {
	Code             int64                                                                            `json:"code" api:"required"`
	Message          string                                                                           `json:"message" api:"required"`
	DocumentationURL string                                                                           `json:"documentation_url"`
	Source           DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSource `json:"source"`
	JSON             dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultError]
type dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSource struct {
	Pointer string                                                                               `json:"pointer"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSourceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSource]
type dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessage struct {
	Code             int64                                                                              `json:"code" api:"required"`
	Message          string                                                                             `json:"message" api:"required"`
	DocumentationURL string                                                                             `json:"documentation_url"`
	Source           DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSource `json:"source"`
	JSON             dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessageJSON    `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessageJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessage]
type dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessageJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSource struct {
	Pointer string                                                                                 `json:"pointer"`
	JSON    dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSourceJSON
// contains the JSON metadata for the struct
// [DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSource]
type dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccess bool

const (
	DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccessTrue DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccess = true
)

func (r DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptTagDeleteResponseWorkersAPIResponseNullResultSuccessTrue:
		return true
	}
	return false
}

// Whether the API call was successful.
type DispatchNamespaceScriptTagDeleteResponseSuccess bool

const (
	DispatchNamespaceScriptTagDeleteResponseSuccessTrue DispatchNamespaceScriptTagDeleteResponseSuccess = true
)

func (r DispatchNamespaceScriptTagDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceScriptTagDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceScriptTagUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Tags associated with the Worker.
	Body []string `json:"body" api:"required"`
}

func (r DispatchNamespaceScriptTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DispatchNamespaceScriptTagListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DispatchNamespaceScriptTagDeleteParams struct {
	// Identifier.
	AccountID  param.Field[string] `path:"account_id" api:"required"`
	APIVersion param.Field[string] `header:"api-version"`
}
