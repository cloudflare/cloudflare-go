// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// NamespaceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceService] method instead.
type NamespaceService struct {
	Options   []option.RequestOption
	Instances *NamespaceInstanceService
}

// NewNamespaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNamespaceService(opts ...option.RequestOption) (r *NamespaceService) {
	r = &NamespaceService{}
	r.Options = opts
	r.Instances = NewNamespaceInstanceService(opts...)
	return
}

// Create a namespace for organizing AI Search instances.
func (r *NamespaceService) New(ctx context.Context, params NamespaceNewParams, opts ...option.RequestOption) (res *NamespaceNewResponse, err error) {
	var env NamespaceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the description and/or the public endpoint configuration of an existing
// namespace. The default namespace's description cannot be modified, but its
// public endpoint can.
func (r *NamespaceService) Update(ctx context.Context, name string, params NamespaceUpdateParams, opts ...option.RequestOption) (res *NamespaceUpdateResponse, err error) {
	var env NamespaceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s", params.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List namespaces in the account, including their descriptions and creation times.
func (r *NamespaceService) List(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[NamespaceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List namespaces in the account, including their descriptions and creation times.
func (r *NamespaceService) ListAutoPaging(ctx context.Context, params NamespaceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Permanently delete a namespace. The namespace must be empty (no instances), and
// the default namespace cannot be deleted.
func (r *NamespaceService) Delete(ctx context.Context, name string, body NamespaceDeleteParams, opts ...option.RequestOption) (res *NamespaceDeleteResponse, err error) {
	var env NamespaceDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s", body.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Performs a chat completion request against multiple AI Search instances in
// parallel, merging retrieved content as context for generating a response.
func (r *NamespaceService) ChatCompletions(ctx context.Context, name string, params NamespaceChatCompletionsParams, opts ...option.RequestOption) (res *NamespaceChatCompletionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/chat/completions", params.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a namespace and its description.
func (r *NamespaceService) Read(ctx context.Context, name string, query NamespaceReadParams, opts ...option.RequestOption) (res *NamespaceReadResponse, err error) {
	var env NamespaceReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s", query.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Performs a semantic search query against multiple AI Search instances in
// parallel, merging the retrieved results into a single ranked response.
func (r *NamespaceService) Search(ctx context.Context, name string, params NamespaceSearchParams, opts ...option.RequestOption) (res *NamespaceSearchResponse, err error) {
	var env NamespaceSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/search", params.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NamespaceNewResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          string                                   `json:"description" api:"nullable"`
	PublicEndpointID     string                                   `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceNewResponsePublicEndpointParams `json:"public_endpoint_params" api:"nullable"`
	JSON                 namespaceNewResponseJSON                 `json:"-"`
}

// namespaceNewResponseJSON contains the JSON metadata for the struct
// [NamespaceNewResponse]
type namespaceNewResponseJSON struct {
	CreatedAt            apijson.Field
	Name                 apijson.Field
	Description          apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                        `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceNewResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains []string `json:"custom_domains" api:"nullable"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled bool `json:"default_domain_enabled"`
	Enabled              bool `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed []string                                               `json:"instances_allowed"`
	Mcp              NamespaceNewResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit        NamespaceNewResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint   NamespaceNewResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON             namespaceNewResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceNewResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [NamespaceNewResponsePublicEndpointParams]
type namespaceNewResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	InstancesAllowed        apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceNewResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                `json:"disabled"`
	JSON     namespaceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON contains the
// JSON metadata for the struct
// [NamespaceNewResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                            `json:"disabled"`
	JSON     namespaceNewResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceNewResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [NamespaceNewResponsePublicEndpointParamsMcp]
type namespaceNewResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                      `json:"period_ms"`
	Requests  int64                                                      `json:"requests"`
	Technique NamespaceNewResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceNewResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceNewResponsePublicEndpointParamsRateLimitJSON contains the JSON metadata
// for the struct [NamespaceNewResponsePublicEndpointParamsRateLimit]
type namespaceNewResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceNewResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceNewResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceNewResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceNewResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceNewResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceNewResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceNewResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceNewResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                       `json:"disabled"`
	JSON     namespaceNewResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceNewResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct [NamespaceNewResponsePublicEndpointParamsSearchEndpoint]
type namespaceNewResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          string                                      `json:"description" api:"nullable"`
	PublicEndpointID     string                                      `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceUpdateResponsePublicEndpointParams `json:"public_endpoint_params" api:"nullable"`
	JSON                 namespaceUpdateResponseJSON                 `json:"-"`
}

// namespaceUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceUpdateResponse]
type namespaceUpdateResponseJSON struct {
	CreatedAt            apijson.Field
	Name                 apijson.Field
	Description          apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                           `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains []string `json:"custom_domains" api:"nullable"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled bool `json:"default_domain_enabled"`
	Enabled              bool `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed []string                                                  `json:"instances_allowed"`
	Mcp              NamespaceUpdateResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit        NamespaceUpdateResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint   NamespaceUpdateResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON             namespaceUpdateResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceUpdateResponsePublicEndpointParamsJSON contains the JSON metadata for
// the struct [NamespaceUpdateResponsePublicEndpointParams]
type namespaceUpdateResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	InstancesAllowed        apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceUpdateResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                   `json:"disabled"`
	JSON     namespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON contains
// the JSON metadata for the struct
// [NamespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                               `json:"disabled"`
	JSON     namespaceUpdateResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceUpdateResponsePublicEndpointParamsMcpJSON contains the JSON metadata
// for the struct [NamespaceUpdateResponsePublicEndpointParamsMcp]
type namespaceUpdateResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                         `json:"period_ms"`
	Requests  int64                                                         `json:"requests"`
	Technique NamespaceUpdateResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceUpdateResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceUpdateResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct [NamespaceUpdateResponsePublicEndpointParamsRateLimit]
type namespaceUpdateResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceUpdateResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceUpdateResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceUpdateResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceUpdateResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                          `json:"disabled"`
	JSON     namespaceUpdateResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceUpdateResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct
// [NamespaceUpdateResponsePublicEndpointParamsSearchEndpoint]
type namespaceUpdateResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          string                                    `json:"description" api:"nullable"`
	PublicEndpointID     string                                    `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceListResponsePublicEndpointParams `json:"public_endpoint_params" api:"nullable"`
	JSON                 namespaceListResponseJSON                 `json:"-"`
}

// namespaceListResponseJSON contains the JSON metadata for the struct
// [NamespaceListResponse]
type namespaceListResponseJSON struct {
	CreatedAt            apijson.Field
	Name                 apijson.Field
	Description          apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                         `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceListResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains []string `json:"custom_domains" api:"nullable"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled bool `json:"default_domain_enabled"`
	Enabled              bool `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed []string                                                `json:"instances_allowed"`
	Mcp              NamespaceListResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit        NamespaceListResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint   NamespaceListResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON             namespaceListResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceListResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [NamespaceListResponsePublicEndpointParams]
type namespaceListResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	InstancesAllowed        apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceListResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                 `json:"disabled"`
	JSON     namespaceListResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceListResponsePublicEndpointParamsChatCompletionsEndpointJSON contains
// the JSON metadata for the struct
// [NamespaceListResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceListResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                             `json:"disabled"`
	JSON     namespaceListResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceListResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [NamespaceListResponsePublicEndpointParamsMcp]
type namespaceListResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                       `json:"period_ms"`
	Requests  int64                                                       `json:"requests"`
	Technique NamespaceListResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceListResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceListResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct [NamespaceListResponsePublicEndpointParamsRateLimit]
type namespaceListResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceListResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceListResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceListResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceListResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceListResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceListResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceListResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceListResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceListResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                        `json:"disabled"`
	JSON     namespaceListResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceListResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct
// [NamespaceListResponsePublicEndpointParamsSearchEndpoint]
type namespaceListResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceListResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceListResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceDeleteResponse = interface{}

type NamespaceChatCompletionsResponse struct {
	Choices     []NamespaceChatCompletionsResponseChoice `json:"choices" api:"required"`
	Chunks      []NamespaceChatCompletionsResponseChunk  `json:"chunks" api:"required"`
	ID          string                                   `json:"id"`
	Errors      []NamespaceChatCompletionsResponseError  `json:"errors"`
	Model       string                                   `json:"model"`
	Object      string                                   `json:"object"`
	ExtraFields map[string]interface{}                   `json:"-" api:"extrafields"`
	JSON        namespaceChatCompletionsResponseJSON     `json:"-"`
}

// namespaceChatCompletionsResponseJSON contains the JSON metadata for the struct
// [NamespaceChatCompletionsResponse]
type namespaceChatCompletionsResponseJSON struct {
	Choices     apijson.Field
	Chunks      apijson.Field
	ID          apijson.Field
	Errors      apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceChatCompletionsResponseChoice struct {
	Message NamespaceChatCompletionsResponseChoicesMessage `json:"message" api:"required"`
	Index   int64                                          `json:"index"`
	JSON    namespaceChatCompletionsResponseChoiceJSON     `json:"-"`
}

// namespaceChatCompletionsResponseChoiceJSON contains the JSON metadata for the
// struct [NamespaceChatCompletionsResponseChoice]
type namespaceChatCompletionsResponseChoiceJSON struct {
	Message     apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChoiceJSON) RawJSON() string {
	return r.raw
}

type NamespaceChatCompletionsResponseChoicesMessage struct {
	Content     NamespaceChatCompletionsResponseChoicesMessageContentUnion `json:"content" api:"required"`
	Role        NamespaceChatCompletionsResponseChoicesMessageRole         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                     `json:"-" api:"extrafields"`
	JSON        namespaceChatCompletionsResponseChoicesMessageJSON         `json:"-"`
}

// namespaceChatCompletionsResponseChoicesMessageJSON contains the JSON metadata
// for the struct [NamespaceChatCompletionsResponseChoicesMessage]
type namespaceChatCompletionsResponseChoicesMessageJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChoicesMessageJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString],
// [NamespaceChatCompletionsResponseChoicesMessageContentArray] or
// [shared.UnionString].
type NamespaceChatCompletionsResponseChoicesMessageContentUnion interface {
	ImplementsNamespaceChatCompletionsResponseChoicesMessageContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceChatCompletionsResponseChoicesMessageContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceChatCompletionsResponseChoicesMessageContentArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NamespaceChatCompletionsResponseChoicesMessageContentArray []NamespaceChatCompletionsResponseChoicesMessageContentArrayItem

func (r NamespaceChatCompletionsResponseChoicesMessageContentArray) ImplementsNamespaceChatCompletionsResponseChoicesMessageContentUnion() {
}

type NamespaceChatCompletionsResponseChoicesMessageContentArrayItem struct {
	Type NamespaceChatCompletionsResponseChoicesMessageContentArrayType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectFile].
	File interface{} `json:"file"`
	// This field can have the runtime type of
	// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectImageURL].
	ImageURL interface{}                                                        `json:"image_url"`
	Text     string                                                             `json:"text"`
	JSON     namespaceChatCompletionsResponseChoicesMessageContentArrayItemJSON `json:"-"`
	union    NamespaceChatCompletionsResponseChoicesMessageContentArrayUnionItem
}

// namespaceChatCompletionsResponseChoicesMessageContentArrayItemJSON contains the
// JSON metadata for the struct
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayItem]
type namespaceChatCompletionsResponseChoicesMessageContentArrayItemJSON struct {
	Type        apijson.Field
	File        apijson.Field
	ImageURL    apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r namespaceChatCompletionsResponseChoicesMessageContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *NamespaceChatCompletionsResponseChoicesMessageContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = NamespaceChatCompletionsResponseChoicesMessageContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayUnionItem] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject].
func (r NamespaceChatCompletionsResponseChoicesMessageContentArrayItem) AsUnion() NamespaceChatCompletionsResponseChoicesMessageContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject] or
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject].
type NamespaceChatCompletionsResponseChoicesMessageContentArrayUnionItem interface {
	implementsNamespaceChatCompletionsResponseChoicesMessageContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceChatCompletionsResponseChoicesMessageContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
	)
}

type NamespaceChatCompletionsResponseChoicesMessageContentArrayObject struct {
	Text string                                                               `json:"text" api:"required"`
	Type NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectType `json:"type" api:"required"`
	JSON namespaceChatCompletionsResponseChoicesMessageContentArrayObjectJSON `json:"-"`
}

// namespaceChatCompletionsResponseChoicesMessageContentArrayObjectJSON contains
// the JSON metadata for the struct
// [NamespaceChatCompletionsResponseChoicesMessageContentArrayObject]
type namespaceChatCompletionsResponseChoicesMessageContentArrayObjectJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChoicesMessageContentArrayObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChoicesMessageContentArrayObjectJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceChatCompletionsResponseChoicesMessageContentArrayObject) implementsNamespaceChatCompletionsResponseChoicesMessageContentArrayItem() {
}

type NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectType string

const (
	NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectType = "text"
)

func (r NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceChatCompletionsResponseChoicesMessageContentArrayType string

const (
	NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeText     NamespaceChatCompletionsResponseChoicesMessageContentArrayType = "text"
	NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL NamespaceChatCompletionsResponseChoicesMessageContentArrayType = "image_url"
	NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeFile     NamespaceChatCompletionsResponseChoicesMessageContentArrayType = "file"
)

func (r NamespaceChatCompletionsResponseChoicesMessageContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeText, NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL, NamespaceChatCompletionsResponseChoicesMessageContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceChatCompletionsResponseChoicesMessageRole string

const (
	NamespaceChatCompletionsResponseChoicesMessageRoleSystem    NamespaceChatCompletionsResponseChoicesMessageRole = "system"
	NamespaceChatCompletionsResponseChoicesMessageRoleDeveloper NamespaceChatCompletionsResponseChoicesMessageRole = "developer"
	NamespaceChatCompletionsResponseChoicesMessageRoleUser      NamespaceChatCompletionsResponseChoicesMessageRole = "user"
	NamespaceChatCompletionsResponseChoicesMessageRoleAssistant NamespaceChatCompletionsResponseChoicesMessageRole = "assistant"
	NamespaceChatCompletionsResponseChoicesMessageRoleTool      NamespaceChatCompletionsResponseChoicesMessageRole = "tool"
)

func (r NamespaceChatCompletionsResponseChoicesMessageRole) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsResponseChoicesMessageRoleSystem, NamespaceChatCompletionsResponseChoicesMessageRoleDeveloper, NamespaceChatCompletionsResponseChoicesMessageRoleUser, NamespaceChatCompletionsResponseChoicesMessageRoleAssistant, NamespaceChatCompletionsResponseChoicesMessageRoleTool:
		return true
	}
	return false
}

type NamespaceChatCompletionsResponseChunk struct {
	ID             string                                               `json:"id" api:"required"`
	InstanceID     string                                               `json:"instance_id" api:"required"`
	Score          float64                                              `json:"score" api:"required"`
	Text           string                                               `json:"text" api:"required"`
	Type           string                                               `json:"type" api:"required"`
	Item           NamespaceChatCompletionsResponseChunksItem           `json:"item"`
	ScoringDetails NamespaceChatCompletionsResponseChunksScoringDetails `json:"scoring_details"`
	JSON           namespaceChatCompletionsResponseChunkJSON            `json:"-"`
}

// namespaceChatCompletionsResponseChunkJSON contains the JSON metadata for the
// struct [NamespaceChatCompletionsResponseChunk]
type namespaceChatCompletionsResponseChunkJSON struct {
	ID             apijson.Field
	InstanceID     apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChunkJSON) RawJSON() string {
	return r.raw
}

type NamespaceChatCompletionsResponseChunksItem struct {
	Key       string                                         `json:"key" api:"required"`
	Metadata  map[string]interface{}                         `json:"metadata"`
	Timestamp float64                                        `json:"timestamp"`
	JSON      namespaceChatCompletionsResponseChunksItemJSON `json:"-"`
}

// namespaceChatCompletionsResponseChunksItemJSON contains the JSON metadata for
// the struct [NamespaceChatCompletionsResponseChunksItem]
type namespaceChatCompletionsResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type NamespaceChatCompletionsResponseChunksScoringDetails struct {
	FusionMethod   NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                          `json:"keyword_rank"`
	KeywordScore   float64                                                          `json:"keyword_score"`
	RerankingScore float64                                                          `json:"reranking_score"`
	VectorRank     float64                                                          `json:"vector_rank"`
	VectorScore    float64                                                          `json:"vector_score"`
	JSON           namespaceChatCompletionsResponseChunksScoringDetailsJSON         `json:"-"`
}

// namespaceChatCompletionsResponseChunksScoringDetailsJSON contains the JSON
// metadata for the struct [NamespaceChatCompletionsResponseChunksScoringDetails]
type namespaceChatCompletionsResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethod string

const (
	NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethod = "rrf"
	NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethodMax NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethod = "max"
)

func (r NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf, NamespaceChatCompletionsResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type NamespaceChatCompletionsResponseError struct {
	InstanceID string                                    `json:"instance_id" api:"required"`
	Message    string                                    `json:"message" api:"required"`
	JSON       namespaceChatCompletionsResponseErrorJSON `json:"-"`
}

// namespaceChatCompletionsResponseErrorJSON contains the JSON metadata for the
// struct [NamespaceChatCompletionsResponseError]
type namespaceChatCompletionsResponseErrorJSON struct {
	InstanceID  apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceChatCompletionsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceChatCompletionsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          string                                    `json:"description" api:"nullable"`
	PublicEndpointID     string                                    `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceReadResponsePublicEndpointParams `json:"public_endpoint_params" api:"nullable"`
	JSON                 namespaceReadResponseJSON                 `json:"-"`
}

// namespaceReadResponseJSON contains the JSON metadata for the struct
// [NamespaceReadResponse]
type namespaceReadResponseJSON struct {
	CreatedAt            apijson.Field
	Name                 apijson.Field
	Description          apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                         `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceReadResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains []string `json:"custom_domains" api:"nullable"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled bool `json:"default_domain_enabled"`
	Enabled              bool `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed []string                                                `json:"instances_allowed"`
	Mcp              NamespaceReadResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit        NamespaceReadResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint   NamespaceReadResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON             namespaceReadResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceReadResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [NamespaceReadResponsePublicEndpointParams]
type namespaceReadResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	InstancesAllowed        apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceReadResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                 `json:"disabled"`
	JSON     namespaceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON contains
// the JSON metadata for the struct
// [NamespaceReadResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                             `json:"disabled"`
	JSON     namespaceReadResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceReadResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [NamespaceReadResponsePublicEndpointParamsMcp]
type namespaceReadResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                       `json:"period_ms"`
	Requests  int64                                                       `json:"requests"`
	Technique NamespaceReadResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceReadResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceReadResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct [NamespaceReadResponsePublicEndpointParamsRateLimit]
type namespaceReadResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceReadResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceReadResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceReadResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceReadResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceReadResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceReadResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceReadResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceReadResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                        `json:"disabled"`
	JSON     namespaceReadResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceReadResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct
// [NamespaceReadResponsePublicEndpointParamsSearchEndpoint]
type namespaceReadResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponse struct {
	Chunks      []NamespaceSearchResponseChunk   `json:"chunks" api:"required"`
	QueryKind   NamespaceSearchResponseQueryKind `json:"query_kind" api:"required"`
	Errors      []NamespaceSearchResponseError   `json:"errors"`
	SearchQuery string                           `json:"search_query"`
	JSON        namespaceSearchResponseJSON      `json:"-"`
}

// namespaceSearchResponseJSON contains the JSON metadata for the struct
// [NamespaceSearchResponse]
type namespaceSearchResponseJSON struct {
	Chunks      apijson.Field
	QueryKind   apijson.Field
	Errors      apijson.Field
	SearchQuery apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponseChunk struct {
	ID             string                                      `json:"id" api:"required"`
	InstanceID     string                                      `json:"instance_id" api:"required"`
	Score          float64                                     `json:"score" api:"required"`
	Text           string                                      `json:"text" api:"required"`
	Type           string                                      `json:"type" api:"required"`
	Item           NamespaceSearchResponseChunksItem           `json:"item"`
	ScoringDetails NamespaceSearchResponseChunksScoringDetails `json:"scoring_details"`
	JSON           namespaceSearchResponseChunkJSON            `json:"-"`
}

// namespaceSearchResponseChunkJSON contains the JSON metadata for the struct
// [NamespaceSearchResponseChunk]
type namespaceSearchResponseChunkJSON struct {
	ID             apijson.Field
	InstanceID     apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceSearchResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseChunkJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponseChunksItem struct {
	Key       string                                `json:"key" api:"required"`
	Metadata  map[string]interface{}                `json:"metadata"`
	Timestamp float64                               `json:"timestamp"`
	JSON      namespaceSearchResponseChunksItemJSON `json:"-"`
}

// namespaceSearchResponseChunksItemJSON contains the JSON metadata for the struct
// [NamespaceSearchResponseChunksItem]
type namespaceSearchResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceSearchResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponseChunksScoringDetails struct {
	FusionMethod   NamespaceSearchResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                 `json:"keyword_rank"`
	KeywordScore   float64                                                 `json:"keyword_score"`
	RerankingScore float64                                                 `json:"reranking_score"`
	VectorRank     float64                                                 `json:"vector_rank"`
	VectorScore    float64                                                 `json:"vector_score"`
	JSON           namespaceSearchResponseChunksScoringDetailsJSON         `json:"-"`
}

// namespaceSearchResponseChunksScoringDetailsJSON contains the JSON metadata for
// the struct [NamespaceSearchResponseChunksScoringDetails]
type namespaceSearchResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceSearchResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type NamespaceSearchResponseChunksScoringDetailsFusionMethod string

const (
	NamespaceSearchResponseChunksScoringDetailsFusionMethodRrf NamespaceSearchResponseChunksScoringDetailsFusionMethod = "rrf"
	NamespaceSearchResponseChunksScoringDetailsFusionMethodMax NamespaceSearchResponseChunksScoringDetailsFusionMethod = "max"
)

func (r NamespaceSearchResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceSearchResponseChunksScoringDetailsFusionMethodRrf, NamespaceSearchResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type NamespaceSearchResponseQueryKind string

const (
	NamespaceSearchResponseQueryKindText       NamespaceSearchResponseQueryKind = "text"
	NamespaceSearchResponseQueryKindImage      NamespaceSearchResponseQueryKind = "image"
	NamespaceSearchResponseQueryKindMultimodal NamespaceSearchResponseQueryKind = "multimodal"
)

func (r NamespaceSearchResponseQueryKind) IsKnown() bool {
	switch r {
	case NamespaceSearchResponseQueryKindText, NamespaceSearchResponseQueryKindImage, NamespaceSearchResponseQueryKindMultimodal:
		return true
	}
	return false
}

type NamespaceSearchResponseError struct {
	InstanceID string                           `json:"instance_id" api:"required"`
	Message    string                           `json:"message" api:"required"`
	JSON       namespaceSearchResponseErrorJSON `json:"-"`
}

// namespaceSearchResponseErrorJSON contains the JSON metadata for the struct
// [NamespaceSearchResponseError]
type namespaceSearchResponseErrorJSON struct {
	InstanceID  apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceSearchResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseErrorJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Name      param.Field[string] `json:"name" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          param.Field[string]                                 `json:"description"`
	PublicEndpointParams param.Field[NamespaceNewParamsPublicEndpointParams] `json:"public_endpoint_params"`
}

func (r NamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                      `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[NamespaceNewParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains param.Field[[]string] `json:"custom_domains"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled param.Field[bool] `json:"default_domain_enabled"`
	Enabled              param.Field[bool] `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed param.Field[[]string]                                             `json:"instances_allowed"`
	Mcp              param.Field[NamespaceNewParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit        param.Field[NamespaceNewParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint   param.Field[NamespaceNewParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
}

func (r NamespaceNewParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceNewParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceNewParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                    `json:"period_ms"`
	Requests  param.Field[int64]                                                    `json:"requests"`
	Technique param.Field[NamespaceNewParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r NamespaceNewParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewParamsPublicEndpointParamsRateLimitTechnique string

const (
	NamespaceNewParamsPublicEndpointParamsRateLimitTechniqueFixed   NamespaceNewParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceNewParamsPublicEndpointParamsRateLimitTechniqueSliding NamespaceNewParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceNewParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceNewParamsPublicEndpointParamsRateLimitTechniqueFixed, NamespaceNewParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceNewParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceNewParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceNewResponseEnvelope struct {
	Result  NamespaceNewResponse                `json:"result" api:"required"`
	Success NamespaceNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    namespaceNewResponseEnvelopeJSON    `json:"-"`
}

// namespaceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceNewResponseEnvelope]
type namespaceNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceNewResponseEnvelopeSuccess bool

const (
	NamespaceNewResponseEnvelopeSuccessTrue NamespaceNewResponseEnvelopeSuccess = true
)

func (r NamespaceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceUpdateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Optional description for the namespace. Max 256 characters.
	Description          param.Field[string]                                    `json:"description"`
	PublicEndpointParams param.Field[NamespaceUpdateParamsPublicEndpointParams] `json:"public_endpoint_params"`
}

func (r NamespaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                         `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[NamespaceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
	// Custom domain hostnames that alias this public endpoint. GET and create
	// responses return the current set; on update (PUT) this field is only echoed back
	// when supplied in the request body, otherwise it is null (omit it to leave
	// domains unchanged).
	CustomDomains param.Field[[]string] `json:"custom_domains"`
	// When false, the instance is reachable only via a registered custom domain and
	// the default <public_endpoint_id>.search.ai.cloudflare.com host returns 404.
	// Requires at least one custom domain. Defaults to true. public_endpoint_params is
	// replaced wholesale on update, so resend default_domain_enabled on every update
	// to keep the default host off — omitting it resets to true.
	DefaultDomainEnabled param.Field[bool] `json:"default_domain_enabled"`
	Enabled              param.Field[bool] `json:"enabled"`
	// Instance IDs exposed through the namespace public endpoint. Empty means nothing
	// is searchable. Every ID must be an existing instance in this namespace, and the
	// list cannot exceed the account's multi-instance search limit.
	InstancesAllowed param.Field[[]string]                                                `json:"instances_allowed"`
	Mcp              param.Field[NamespaceUpdateParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit        param.Field[NamespaceUpdateParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint   param.Field[NamespaceUpdateParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
}

func (r NamespaceUpdateParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceUpdateParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                       `json:"period_ms"`
	Requests  param.Field[int64]                                                       `json:"requests"`
	Technique param.Field[NamespaceUpdateParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r NamespaceUpdateParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateParamsPublicEndpointParamsRateLimitTechnique string

const (
	NamespaceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed   NamespaceUpdateParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding NamespaceUpdateParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceUpdateParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed, NamespaceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceUpdateParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceUpdateParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceUpdateResponseEnvelope struct {
	Result  NamespaceUpdateResponse                `json:"result" api:"required"`
	Success NamespaceUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    namespaceUpdateResponseEnvelopeJSON    `json:"-"`
}

// namespaceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceUpdateResponseEnvelope]
type namespaceUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceUpdateResponseEnvelopeSuccess bool

const (
	NamespaceUpdateResponseEnvelopeSuccessTrue NamespaceUpdateResponseEnvelopeSuccess = true
)

func (r NamespaceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number (1-indexed).
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter namespaces whose name or description contains this string
	// (case-insensitive).
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [NamespaceListParams]'s query parameters as `url.Values`.
func (r NamespaceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type NamespaceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceDeleteResponseEnvelope struct {
	Result  NamespaceDeleteResponse                `json:"result" api:"required"`
	Success NamespaceDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    namespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// namespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceDeleteResponseEnvelope]
type namespaceDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceDeleteResponseEnvelopeSuccess bool

const (
	NamespaceDeleteResponseEnvelopeSuccessTrue NamespaceDeleteResponseEnvelopeSuccess = true
)

func (r NamespaceDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceChatCompletionsParams struct {
	AccountID       param.Field[string]                                        `path:"account_id" api:"required"`
	AISearchOptions param.Field[NamespaceChatCompletionsParamsAISearchOptions] `json:"ai_search_options" api:"required"`
	Messages        param.Field[[]NamespaceChatCompletionsParamsMessage]       `json:"messages" api:"required"`
	Model           param.Field[string]                                        `json:"model"`
	Stream          param.Field[bool]                                          `json:"stream"`
}

func (r NamespaceChatCompletionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptions struct {
	InstanceIDs  param.Field[[]string]                                                  `json:"instance_ids" api:"required"`
	Cache        param.Field[NamespaceChatCompletionsParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[NamespaceChatCompletionsParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[NamespaceChatCompletionsParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r NamespaceChatCompletionsParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                             `json:"enabled"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold string

const (
	NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough      NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes     NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough, NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, NamespaceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                              `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                             `json:"filters"`
	FusionMethod     param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                                `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                                  `json:"max_num_results"`
	RetrievalType    param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                                   `json:"return_on_failure"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection string

const (
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc       NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc      NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists    NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc, NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc, NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists, NamespaceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod string

const (
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "max"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax, NamespaceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr  NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd, NamespaceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType string

const (
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector  NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "vector"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid  NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector, NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword, NamespaceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsMessage struct {
	Content     param.Field[NamespaceChatCompletionsParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[NamespaceChatCompletionsParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                          `json:"-,extras"`
}

func (r NamespaceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.NamespaceChatCompletionsParamsMessagesContentArray],
// [shared.UnionString].
type NamespaceChatCompletionsParamsMessagesContentUnion interface {
	ImplementsNamespaceChatCompletionsParamsMessagesContentUnion()
}

type NamespaceChatCompletionsParamsMessagesContentArray []NamespaceChatCompletionsParamsMessagesContentArrayItemUnion

func (r NamespaceChatCompletionsParamsMessagesContentArray) ImplementsNamespaceChatCompletionsParamsMessagesContentUnion() {
}

type NamespaceChatCompletionsParamsMessagesContentArrayItem struct {
	Type     param.Field[NamespaceChatCompletionsParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                            `json:"file"`
	ImageURL param.Field[interface{}]                                            `json:"image_url"`
	Text     param.Field[string]                                                 `json:"text"`
}

func (r NamespaceChatCompletionsParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceChatCompletionsParamsMessagesContentArrayItem) implementsNamespaceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [ai_search.NamespaceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.NamespaceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.NamespaceChatCompletionsParamsMessagesContentArrayObject],
// [NamespaceChatCompletionsParamsMessagesContentArrayItem].
type NamespaceChatCompletionsParamsMessagesContentArrayItemUnion interface {
	implementsNamespaceChatCompletionsParamsMessagesContentArrayItemUnion()
}

type NamespaceChatCompletionsParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                                       `json:"text" api:"required"`
	Type param.Field[NamespaceChatCompletionsParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r NamespaceChatCompletionsParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceChatCompletionsParamsMessagesContentArrayObject) implementsNamespaceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

type NamespaceChatCompletionsParamsMessagesContentArrayObjectType string

const (
	NamespaceChatCompletionsParamsMessagesContentArrayObjectTypeText NamespaceChatCompletionsParamsMessagesContentArrayObjectType = "text"
)

func (r NamespaceChatCompletionsParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsMessagesContentArrayType string

const (
	NamespaceChatCompletionsParamsMessagesContentArrayTypeText     NamespaceChatCompletionsParamsMessagesContentArrayType = "text"
	NamespaceChatCompletionsParamsMessagesContentArrayTypeImageURL NamespaceChatCompletionsParamsMessagesContentArrayType = "image_url"
	NamespaceChatCompletionsParamsMessagesContentArrayTypeFile     NamespaceChatCompletionsParamsMessagesContentArrayType = "file"
)

func (r NamespaceChatCompletionsParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsMessagesContentArrayTypeText, NamespaceChatCompletionsParamsMessagesContentArrayTypeImageURL, NamespaceChatCompletionsParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceChatCompletionsParamsMessagesRole string

const (
	NamespaceChatCompletionsParamsMessagesRoleSystem    NamespaceChatCompletionsParamsMessagesRole = "system"
	NamespaceChatCompletionsParamsMessagesRoleDeveloper NamespaceChatCompletionsParamsMessagesRole = "developer"
	NamespaceChatCompletionsParamsMessagesRoleUser      NamespaceChatCompletionsParamsMessagesRole = "user"
	NamespaceChatCompletionsParamsMessagesRoleAssistant NamespaceChatCompletionsParamsMessagesRole = "assistant"
	NamespaceChatCompletionsParamsMessagesRoleTool      NamespaceChatCompletionsParamsMessagesRole = "tool"
)

func (r NamespaceChatCompletionsParamsMessagesRole) IsKnown() bool {
	switch r {
	case NamespaceChatCompletionsParamsMessagesRoleSystem, NamespaceChatCompletionsParamsMessagesRoleDeveloper, NamespaceChatCompletionsParamsMessagesRoleUser, NamespaceChatCompletionsParamsMessagesRoleAssistant, NamespaceChatCompletionsParamsMessagesRoleTool:
		return true
	}
	return false
}

type NamespaceReadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceReadResponseEnvelope struct {
	Result  NamespaceReadResponse                `json:"result" api:"required"`
	Success NamespaceReadResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    namespaceReadResponseEnvelopeJSON    `json:"-"`
}

// namespaceReadResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceReadResponseEnvelope]
type namespaceReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceReadResponseEnvelopeSuccess bool

const (
	NamespaceReadResponseEnvelopeSuccessTrue NamespaceReadResponseEnvelopeSuccess = true
)

func (r NamespaceReadResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceReadResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceSearchParams struct {
	AccountID       param.Field[string]                               `path:"account_id" api:"required"`
	AISearchOptions param.Field[NamespaceSearchParamsAISearchOptions] `json:"ai_search_options" api:"required"`
	// OpenAI-compatible message array. For multimodal queries, set the last user
	// message's `content` to an array of typed parts:
	// `[{type:'text', text:'…'}, {type:'image_url', image_url:{url:'…'}}]`. Image
	// inputs require the RAG's embedding_model to declare 'image' in
	// supported_modalities.
	Messages param.Field[[]NamespaceSearchParamsMessage] `json:"messages"`
	// A simple text query string. Alternative to 'messages' — provide either this or
	// 'messages', not both.
	Query param.Field[string] `json:"query"`
}

func (r NamespaceSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptions struct {
	InstanceIDs  param.Field[[]string]                                         `json:"instance_ids" api:"required"`
	Cache        param.Field[NamespaceSearchParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[NamespaceSearchParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[NamespaceSearchParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[NamespaceSearchParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r NamespaceSearchParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[NamespaceSearchParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                    `json:"enabled"`
}

func (r NamespaceSearchParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsCacheCacheThreshold string

const (
	NamespaceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch NamespaceSearchParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	NamespaceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough      NamespaceSearchParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	NamespaceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   NamespaceSearchParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	NamespaceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes     NamespaceSearchParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r NamespaceSearchParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, NamespaceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough, NamespaceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, NamespaceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type NamespaceSearchParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]   `json:"enabled"`
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r NamespaceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r NamespaceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]NamespaceSearchParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                     `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                    `json:"filters"`
	FusionMethod     param.Field[NamespaceSearchParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                       `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                         `json:"max_num_results"`
	RetrievalType    param.Field[NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                          `json:"return_on_failure"`
}

func (r NamespaceSearchParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceSearchParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceSearchParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection string

const (
	NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc       NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc      NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists    NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc, NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc, NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists, NamespaceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
}

type NamespaceSearchParamsAISearchOptionsRetrievalFusionMethod string

const (
	NamespaceSearchParamsAISearchOptionsRetrievalFusionMethodMax NamespaceSearchParamsAISearchOptionsRetrievalFusionMethod = "max"
	NamespaceSearchParamsAISearchOptionsRetrievalFusionMethodRrf NamespaceSearchParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r NamespaceSearchParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsRetrievalFusionMethodMax, NamespaceSearchParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr  NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd, NamespaceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType string

const (
	NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector  NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType = "vector"
	NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid  NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r NamespaceSearchParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector, NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword, NamespaceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type NamespaceSearchParamsMessage struct {
	Content     param.Field[NamespaceSearchParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[NamespaceSearchParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                 `json:"-,extras"`
}

func (r NamespaceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.NamespaceSearchParamsMessagesContentArray], [shared.UnionString].
type NamespaceSearchParamsMessagesContentUnion interface {
	ImplementsNamespaceSearchParamsMessagesContentUnion()
}

type NamespaceSearchParamsMessagesContentArray []NamespaceSearchParamsMessagesContentArrayItemUnion

func (r NamespaceSearchParamsMessagesContentArray) ImplementsNamespaceSearchParamsMessagesContentUnion() {
}

type NamespaceSearchParamsMessagesContentArrayItem struct {
	Type     param.Field[NamespaceSearchParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                   `json:"file"`
	ImageURL param.Field[interface{}]                                   `json:"image_url"`
	Text     param.Field[string]                                        `json:"text"`
}

func (r NamespaceSearchParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceSearchParamsMessagesContentArrayItem) implementsNamespaceSearchParamsMessagesContentArrayItemUnion() {
}

// Satisfied by [ai_search.NamespaceSearchParamsMessagesContentArrayObject],
// [ai_search.NamespaceSearchParamsMessagesContentArrayObject],
// [ai_search.NamespaceSearchParamsMessagesContentArrayObject],
// [NamespaceSearchParamsMessagesContentArrayItem].
type NamespaceSearchParamsMessagesContentArrayItemUnion interface {
	implementsNamespaceSearchParamsMessagesContentArrayItemUnion()
}

type NamespaceSearchParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                              `json:"text" api:"required"`
	Type param.Field[NamespaceSearchParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r NamespaceSearchParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceSearchParamsMessagesContentArrayObject) implementsNamespaceSearchParamsMessagesContentArrayItemUnion() {
}

type NamespaceSearchParamsMessagesContentArrayObjectType string

const (
	NamespaceSearchParamsMessagesContentArrayObjectTypeText NamespaceSearchParamsMessagesContentArrayObjectType = "text"
)

func (r NamespaceSearchParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceSearchParamsMessagesContentArrayType string

const (
	NamespaceSearchParamsMessagesContentArrayTypeText     NamespaceSearchParamsMessagesContentArrayType = "text"
	NamespaceSearchParamsMessagesContentArrayTypeImageURL NamespaceSearchParamsMessagesContentArrayType = "image_url"
	NamespaceSearchParamsMessagesContentArrayTypeFile     NamespaceSearchParamsMessagesContentArrayType = "file"
)

func (r NamespaceSearchParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsMessagesContentArrayTypeText, NamespaceSearchParamsMessagesContentArrayTypeImageURL, NamespaceSearchParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceSearchParamsMessagesRole string

const (
	NamespaceSearchParamsMessagesRoleSystem    NamespaceSearchParamsMessagesRole = "system"
	NamespaceSearchParamsMessagesRoleDeveloper NamespaceSearchParamsMessagesRole = "developer"
	NamespaceSearchParamsMessagesRoleUser      NamespaceSearchParamsMessagesRole = "user"
	NamespaceSearchParamsMessagesRoleAssistant NamespaceSearchParamsMessagesRole = "assistant"
	NamespaceSearchParamsMessagesRoleTool      NamespaceSearchParamsMessagesRole = "tool"
)

func (r NamespaceSearchParamsMessagesRole) IsKnown() bool {
	switch r {
	case NamespaceSearchParamsMessagesRoleSystem, NamespaceSearchParamsMessagesRoleDeveloper, NamespaceSearchParamsMessagesRoleUser, NamespaceSearchParamsMessagesRoleAssistant, NamespaceSearchParamsMessagesRoleTool:
		return true
	}
	return false
}

type NamespaceSearchResponseEnvelope struct {
	Result  NamespaceSearchResponse             `json:"result" api:"required"`
	Success bool                                `json:"success" api:"required"`
	JSON    namespaceSearchResponseEnvelopeJSON `json:"-"`
}

// namespaceSearchResponseEnvelopeJSON contains the JSON metadata for the struct
// [NamespaceSearchResponseEnvelope]
type namespaceSearchResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
