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

// NamespaceInstanceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceInstanceService] method instead.
type NamespaceInstanceService struct {
	Options []option.RequestOption
	Jobs    *NamespaceInstanceJobService
	Items   *NamespaceInstanceItemService
}

// NewNamespaceInstanceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamespaceInstanceService(opts ...option.RequestOption) (r *NamespaceInstanceService) {
	r = &NamespaceInstanceService{}
	r.Options = opts
	r.Jobs = NewNamespaceInstanceJobService(opts...)
	r.Items = NewNamespaceInstanceItemService(opts...)
	return
}

// Create a new AI Search instance with the given configuration.
func (r *NamespaceInstanceService) New(ctx context.Context, name string, params NamespaceInstanceNewParams, opts ...option.RequestOption) (res *NamespaceInstanceNewResponse, err error) {
	var env NamespaceInstanceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances", params.AccountID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the configuration of an AI Search instance.
func (r *NamespaceInstanceService) Update(ctx context.Context, name string, id string, params NamespaceInstanceUpdateParams, opts ...option.RequestOption) (res *NamespaceInstanceUpdateResponse, err error) {
	var env NamespaceInstanceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List all AI Search instances in the account.
func (r *NamespaceInstanceService) List(ctx context.Context, name string, params NamespaceInstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[NamespaceInstanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances", params.AccountID, name)
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

// List all AI Search instances in the account.
func (r *NamespaceInstanceService) ListAutoPaging(ctx context.Context, name string, params NamespaceInstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceInstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, name, params, opts...))
}

// Permanently delete an AI Search instance and all its indexed data.
func (r *NamespaceInstanceService) Delete(ctx context.Context, name string, id string, body NamespaceInstanceDeleteParams, opts ...option.RequestOption) (res *NamespaceInstanceDeleteResponse, err error) {
	var env NamespaceInstanceDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s", body.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Performs a chat completion request against an AI Search instance, using indexed
// content as context for generating responses.
func (r *NamespaceInstanceService) ChatCompletions(ctx context.Context, name string, id string, params NamespaceInstanceChatCompletionsParams, opts ...option.RequestOption) (res *NamespaceInstanceChatCompletionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/chat/completions", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the configuration and status of an AI Search instance.
func (r *NamespaceInstanceService) Read(ctx context.Context, name string, id string, query NamespaceInstanceReadParams, opts ...option.RequestOption) (res *NamespaceInstanceReadResponse, err error) {
	var env NamespaceInstanceReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s", query.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Executes a semantic search query against an AI Search instance to find relevant
// indexed content.
func (r *NamespaceInstanceService) Search(ctx context.Context, name string, id string, params NamespaceInstanceSearchParams, opts ...option.RequestOption) (res *NamespaceInstanceSearchResponse, err error) {
	var env NamespaceInstanceSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/search", params.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve usage and indexing statistics for an AI Search instance.
func (r *NamespaceInstanceService) Stats(ctx context.Context, name string, id string, query NamespaceInstanceStatsParams, opts ...option.RequestOption) (res *NamespaceInstanceStatsResponse, err error) {
	var env NamespaceInstanceStatsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/ai-search/namespaces/%s/instances/%s/stats", query.AccountID, name, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type NamespaceInstanceNewResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID string    `json:"ai_gateway_id" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  string                                     `json:"ai_search_model" api:"nullable"`
	Cache          bool                                       `json:"cache"`
	CacheThreshold NamespaceInstanceNewResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       NamespaceInstanceNewResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                        `json:"chunk_overlap"`
	ChunkSize      int64                                        `json:"chunk_size"`
	CreatedBy      string                                       `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceNewResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                       `json:"embedding_model" api:"nullable"`
	Enable         bool                                         `json:"enable"`
	EngineVersion  float64                                      `json:"engine_version"`
	FusionMethod   NamespaceInstanceNewResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          NamespaceInstanceNewResponseIndexMethod          `json:"index_method"`
	IndexingOptions      NamespaceInstanceNewResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                        `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                            `json:"max_num_results"`
	Metadata             NamespaceInstanceNewResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                           `json:"modified_by" api:"nullable"`
	Namespace            string                                           `json:"namespace" api:"nullable"`
	Paused               bool                                             `json:"paused"`
	PublicEndpointID     string                                           `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceInstanceNewResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                             `json:"reranking"`
	RerankingModel       string                                           `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceNewResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel   string                                   `json:"rewrite_model" api:"nullable"`
	RewriteQuery   bool                                     `json:"rewrite_query"`
	ScoreThreshold float64                                  `json:"score_threshold"`
	Source         string                                   `json:"source" api:"nullable"`
	SourceParams   NamespaceInstanceNewResponseSourceParams `json:"source_params" api:"nullable"`
	Status         string                                   `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval NamespaceInstanceNewResponseSyncInterval `json:"sync_interval"`
	TokenID      string                                   `json:"token_id" format:"uuid"`
	Type         NamespaceInstanceNewResponseType         `json:"type" api:"nullable"`
	JSON         namespaceInstanceNewResponseJSON         `json:"-"`
}

// namespaceInstanceNewResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceNewResponse]
type namespaceInstanceNewResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	AIGatewayID          apijson.Field
	AISearchModel        apijson.Field
	Cache                apijson.Field
	CacheThreshold       apijson.Field
	CacheTTL             apijson.Field
	ChunkOverlap         apijson.Field
	ChunkSize            apijson.Field
	CreatedBy            apijson.Field
	CustomMetadata       apijson.Field
	EmbeddingModel       apijson.Field
	Enable               apijson.Field
	EngineVersion        apijson.Field
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	IndexMethod          apijson.Field
	IndexingOptions      apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Namespace            apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	Source               apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	SyncInterval         apijson.Field
	TokenID              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseCacheThreshold string

const (
	NamespaceInstanceNewResponseCacheThresholdSuperStrictMatch NamespaceInstanceNewResponseCacheThreshold = "super_strict_match"
	NamespaceInstanceNewResponseCacheThresholdCloseEnough      NamespaceInstanceNewResponseCacheThreshold = "close_enough"
	NamespaceInstanceNewResponseCacheThresholdFlexibleFriend   NamespaceInstanceNewResponseCacheThreshold = "flexible_friend"
	NamespaceInstanceNewResponseCacheThresholdAnythingGoes     NamespaceInstanceNewResponseCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceNewResponseCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseCacheThresholdSuperStrictMatch, NamespaceInstanceNewResponseCacheThresholdCloseEnough, NamespaceInstanceNewResponseCacheThresholdFlexibleFriend, NamespaceInstanceNewResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceNewResponseCacheTTL float64

const (
	NamespaceInstanceNewResponseCacheTTL600    NamespaceInstanceNewResponseCacheTTL = 600
	NamespaceInstanceNewResponseCacheTTL1800   NamespaceInstanceNewResponseCacheTTL = 1800
	NamespaceInstanceNewResponseCacheTTL3600   NamespaceInstanceNewResponseCacheTTL = 3600
	NamespaceInstanceNewResponseCacheTTL7200   NamespaceInstanceNewResponseCacheTTL = 7200
	NamespaceInstanceNewResponseCacheTTL21600  NamespaceInstanceNewResponseCacheTTL = 21600
	NamespaceInstanceNewResponseCacheTTL43200  NamespaceInstanceNewResponseCacheTTL = 43200
	NamespaceInstanceNewResponseCacheTTL86400  NamespaceInstanceNewResponseCacheTTL = 86400
	NamespaceInstanceNewResponseCacheTTL172800 NamespaceInstanceNewResponseCacheTTL = 172800
	NamespaceInstanceNewResponseCacheTTL259200 NamespaceInstanceNewResponseCacheTTL = 259200
	NamespaceInstanceNewResponseCacheTTL518400 NamespaceInstanceNewResponseCacheTTL = 518400
)

func (r NamespaceInstanceNewResponseCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseCacheTTL600, NamespaceInstanceNewResponseCacheTTL1800, NamespaceInstanceNewResponseCacheTTL3600, NamespaceInstanceNewResponseCacheTTL7200, NamespaceInstanceNewResponseCacheTTL21600, NamespaceInstanceNewResponseCacheTTL43200, NamespaceInstanceNewResponseCacheTTL86400, NamespaceInstanceNewResponseCacheTTL172800, NamespaceInstanceNewResponseCacheTTL259200, NamespaceInstanceNewResponseCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseCustomMetadata struct {
	DataType  NamespaceInstanceNewResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                             `json:"field_name" api:"required"`
	JSON      namespaceInstanceNewResponseCustomMetadataJSON     `json:"-"`
}

// namespaceInstanceNewResponseCustomMetadataJSON contains the JSON metadata for
// the struct [NamespaceInstanceNewResponseCustomMetadata]
type namespaceInstanceNewResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseCustomMetadataDataType string

const (
	NamespaceInstanceNewResponseCustomMetadataDataTypeText     NamespaceInstanceNewResponseCustomMetadataDataType = "text"
	NamespaceInstanceNewResponseCustomMetadataDataTypeNumber   NamespaceInstanceNewResponseCustomMetadataDataType = "number"
	NamespaceInstanceNewResponseCustomMetadataDataTypeBoolean  NamespaceInstanceNewResponseCustomMetadataDataType = "boolean"
	NamespaceInstanceNewResponseCustomMetadataDataTypeDatetime NamespaceInstanceNewResponseCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceNewResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseCustomMetadataDataTypeText, NamespaceInstanceNewResponseCustomMetadataDataTypeNumber, NamespaceInstanceNewResponseCustomMetadataDataTypeBoolean, NamespaceInstanceNewResponseCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseFusionMethod string

const (
	NamespaceInstanceNewResponseFusionMethodMax NamespaceInstanceNewResponseFusionMethod = "max"
	NamespaceInstanceNewResponseFusionMethodRrf NamespaceInstanceNewResponseFusionMethod = "rrf"
)

func (r NamespaceInstanceNewResponseFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseFusionMethodMax, NamespaceInstanceNewResponseFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceNewResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                        `json:"vector" api:"required"`
	JSON   namespaceInstanceNewResponseIndexMethodJSON `json:"-"`
}

// namespaceInstanceNewResponseIndexMethodJSON contains the JSON metadata for the
// struct [NamespaceInstanceNewResponseIndexMethod]
type namespaceInstanceNewResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             namespaceInstanceNewResponseIndexingOptionsJSON             `json:"-"`
}

// namespaceInstanceNewResponseIndexingOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceNewResponseIndexingOptions]
type namespaceInstanceNewResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceNewResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseMetadata struct {
	CreatedFromAISearchWizard bool                                     `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                                   `json:"worker_domain"`
	JSON                      namespaceInstanceNewResponseMetadataJSON `json:"-"`
}

// namespaceInstanceNewResponseMetadataJSON contains the JSON metadata for the
// struct [NamespaceInstanceNewResponseMetadata]
type namespaceInstanceNewResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                                `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled bool                                                           `json:"default_domain_enabled"`
	Enabled              bool                                                           `json:"enabled"`
	Mcp                  NamespaceInstanceNewResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            NamespaceInstanceNewResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       NamespaceInstanceNewResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 namespaceInstanceNewResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceNewResponsePublicEndpointParams]
type namespaceInstanceNewResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                        `json:"disabled"`
	JSON     namespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                                    `json:"disabled"`
	JSON     namespaceInstanceNewResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsMcpJSON contains the JSON
// metadata for the struct [NamespaceInstanceNewResponsePublicEndpointParamsMcp]
type namespaceInstanceNewResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                              `json:"period_ms"`
	Requests  int64                                                              `json:"requests"`
	Technique NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceInstanceNewResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct
// [NamespaceInstanceNewResponsePublicEndpointParamsRateLimit]
type namespaceInstanceNewResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceNewResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceNewResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                               `json:"disabled"`
	JSON     namespaceInstanceNewResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsSearchEndpointJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceNewResponsePublicEndpointParamsSearchEndpoint]
type namespaceInstanceNewResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []NamespaceInstanceNewResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             namespaceInstanceNewResponseRetrievalOptionsJSON             `json:"-"`
}

// namespaceInstanceNewResponseRetrievalOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceNewResponseRetrievalOptions]
type namespaceInstanceNewResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      namespaceInstanceNewResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceNewResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceNewResponseRetrievalOptionsBoostBy]
type namespaceInstanceNewResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionExists    NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceNewResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionExists, NamespaceInstanceNewResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceNewResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   []string                                           `json:"include_items"`
	Prefix         string                                             `json:"prefix"`
	R2Jurisdiction string                                             `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceNewResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           namespaceInstanceNewResponseSourceParamsJSON       `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsJSON contains the JSON metadata for the
// struct [NamespaceInstanceNewResponseSourceParams]
type namespaceInstanceNewResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      namespaceInstanceNewResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceNewResponseSourceParamsWebCrawler]
type namespaceInstanceNewResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth float64 `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks bool `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit float64 `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge float64 `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   namespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptions]
type namespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceNewResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                           `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                               `json:"use_browser_rendering"`
	JSON                namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions]
type namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                            `json:"selector" api:"required"`
	JSON     namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceNewResponseSyncInterval float64

const (
	NamespaceInstanceNewResponseSyncInterval900   NamespaceInstanceNewResponseSyncInterval = 900
	NamespaceInstanceNewResponseSyncInterval1800  NamespaceInstanceNewResponseSyncInterval = 1800
	NamespaceInstanceNewResponseSyncInterval3600  NamespaceInstanceNewResponseSyncInterval = 3600
	NamespaceInstanceNewResponseSyncInterval7200  NamespaceInstanceNewResponseSyncInterval = 7200
	NamespaceInstanceNewResponseSyncInterval14400 NamespaceInstanceNewResponseSyncInterval = 14400
	NamespaceInstanceNewResponseSyncInterval21600 NamespaceInstanceNewResponseSyncInterval = 21600
	NamespaceInstanceNewResponseSyncInterval43200 NamespaceInstanceNewResponseSyncInterval = 43200
	NamespaceInstanceNewResponseSyncInterval86400 NamespaceInstanceNewResponseSyncInterval = 86400
)

func (r NamespaceInstanceNewResponseSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseSyncInterval900, NamespaceInstanceNewResponseSyncInterval1800, NamespaceInstanceNewResponseSyncInterval3600, NamespaceInstanceNewResponseSyncInterval7200, NamespaceInstanceNewResponseSyncInterval14400, NamespaceInstanceNewResponseSyncInterval21600, NamespaceInstanceNewResponseSyncInterval43200, NamespaceInstanceNewResponseSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseType string

const (
	NamespaceInstanceNewResponseTypeR2         NamespaceInstanceNewResponseType = "r2"
	NamespaceInstanceNewResponseTypeWebCrawler NamespaceInstanceNewResponseType = "web-crawler"
)

func (r NamespaceInstanceNewResponseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseTypeR2, NamespaceInstanceNewResponseTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID string    `json:"ai_gateway_id" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  string                                        `json:"ai_search_model" api:"nullable"`
	Cache          bool                                          `json:"cache"`
	CacheThreshold NamespaceInstanceUpdateResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       NamespaceInstanceUpdateResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                           `json:"chunk_overlap"`
	ChunkSize      int64                                           `json:"chunk_size"`
	CreatedBy      string                                          `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceUpdateResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                          `json:"embedding_model" api:"nullable"`
	Enable         bool                                            `json:"enable"`
	EngineVersion  float64                                         `json:"engine_version"`
	FusionMethod   NamespaceInstanceUpdateResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          NamespaceInstanceUpdateResponseIndexMethod          `json:"index_method"`
	IndexingOptions      NamespaceInstanceUpdateResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                           `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                               `json:"max_num_results"`
	Metadata             NamespaceInstanceUpdateResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                              `json:"modified_by" api:"nullable"`
	Namespace            string                                              `json:"namespace" api:"nullable"`
	Paused               bool                                                `json:"paused"`
	PublicEndpointID     string                                              `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceInstanceUpdateResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                                `json:"reranking"`
	RerankingModel       string                                              `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceUpdateResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel   string                                      `json:"rewrite_model" api:"nullable"`
	RewriteQuery   bool                                        `json:"rewrite_query"`
	ScoreThreshold float64                                     `json:"score_threshold"`
	Source         string                                      `json:"source" api:"nullable"`
	SourceParams   NamespaceInstanceUpdateResponseSourceParams `json:"source_params" api:"nullable"`
	Status         string                                      `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval NamespaceInstanceUpdateResponseSyncInterval `json:"sync_interval"`
	TokenID      string                                      `json:"token_id" format:"uuid"`
	Type         NamespaceInstanceUpdateResponseType         `json:"type" api:"nullable"`
	JSON         namespaceInstanceUpdateResponseJSON         `json:"-"`
}

// namespaceInstanceUpdateResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceUpdateResponse]
type namespaceInstanceUpdateResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	AIGatewayID          apijson.Field
	AISearchModel        apijson.Field
	Cache                apijson.Field
	CacheThreshold       apijson.Field
	CacheTTL             apijson.Field
	ChunkOverlap         apijson.Field
	ChunkSize            apijson.Field
	CreatedBy            apijson.Field
	CustomMetadata       apijson.Field
	EmbeddingModel       apijson.Field
	Enable               apijson.Field
	EngineVersion        apijson.Field
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	IndexMethod          apijson.Field
	IndexingOptions      apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Namespace            apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	Source               apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	SyncInterval         apijson.Field
	TokenID              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseCacheThreshold string

const (
	NamespaceInstanceUpdateResponseCacheThresholdSuperStrictMatch NamespaceInstanceUpdateResponseCacheThreshold = "super_strict_match"
	NamespaceInstanceUpdateResponseCacheThresholdCloseEnough      NamespaceInstanceUpdateResponseCacheThreshold = "close_enough"
	NamespaceInstanceUpdateResponseCacheThresholdFlexibleFriend   NamespaceInstanceUpdateResponseCacheThreshold = "flexible_friend"
	NamespaceInstanceUpdateResponseCacheThresholdAnythingGoes     NamespaceInstanceUpdateResponseCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceUpdateResponseCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseCacheThresholdSuperStrictMatch, NamespaceInstanceUpdateResponseCacheThresholdCloseEnough, NamespaceInstanceUpdateResponseCacheThresholdFlexibleFriend, NamespaceInstanceUpdateResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceUpdateResponseCacheTTL float64

const (
	NamespaceInstanceUpdateResponseCacheTTL600    NamespaceInstanceUpdateResponseCacheTTL = 600
	NamespaceInstanceUpdateResponseCacheTTL1800   NamespaceInstanceUpdateResponseCacheTTL = 1800
	NamespaceInstanceUpdateResponseCacheTTL3600   NamespaceInstanceUpdateResponseCacheTTL = 3600
	NamespaceInstanceUpdateResponseCacheTTL7200   NamespaceInstanceUpdateResponseCacheTTL = 7200
	NamespaceInstanceUpdateResponseCacheTTL21600  NamespaceInstanceUpdateResponseCacheTTL = 21600
	NamespaceInstanceUpdateResponseCacheTTL43200  NamespaceInstanceUpdateResponseCacheTTL = 43200
	NamespaceInstanceUpdateResponseCacheTTL86400  NamespaceInstanceUpdateResponseCacheTTL = 86400
	NamespaceInstanceUpdateResponseCacheTTL172800 NamespaceInstanceUpdateResponseCacheTTL = 172800
	NamespaceInstanceUpdateResponseCacheTTL259200 NamespaceInstanceUpdateResponseCacheTTL = 259200
	NamespaceInstanceUpdateResponseCacheTTL518400 NamespaceInstanceUpdateResponseCacheTTL = 518400
)

func (r NamespaceInstanceUpdateResponseCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseCacheTTL600, NamespaceInstanceUpdateResponseCacheTTL1800, NamespaceInstanceUpdateResponseCacheTTL3600, NamespaceInstanceUpdateResponseCacheTTL7200, NamespaceInstanceUpdateResponseCacheTTL21600, NamespaceInstanceUpdateResponseCacheTTL43200, NamespaceInstanceUpdateResponseCacheTTL86400, NamespaceInstanceUpdateResponseCacheTTL172800, NamespaceInstanceUpdateResponseCacheTTL259200, NamespaceInstanceUpdateResponseCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseCustomMetadata struct {
	DataType  NamespaceInstanceUpdateResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                                `json:"field_name" api:"required"`
	JSON      namespaceInstanceUpdateResponseCustomMetadataJSON     `json:"-"`
}

// namespaceInstanceUpdateResponseCustomMetadataJSON contains the JSON metadata for
// the struct [NamespaceInstanceUpdateResponseCustomMetadata]
type namespaceInstanceUpdateResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseCustomMetadataDataType string

const (
	NamespaceInstanceUpdateResponseCustomMetadataDataTypeText     NamespaceInstanceUpdateResponseCustomMetadataDataType = "text"
	NamespaceInstanceUpdateResponseCustomMetadataDataTypeNumber   NamespaceInstanceUpdateResponseCustomMetadataDataType = "number"
	NamespaceInstanceUpdateResponseCustomMetadataDataTypeBoolean  NamespaceInstanceUpdateResponseCustomMetadataDataType = "boolean"
	NamespaceInstanceUpdateResponseCustomMetadataDataTypeDatetime NamespaceInstanceUpdateResponseCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceUpdateResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseCustomMetadataDataTypeText, NamespaceInstanceUpdateResponseCustomMetadataDataTypeNumber, NamespaceInstanceUpdateResponseCustomMetadataDataTypeBoolean, NamespaceInstanceUpdateResponseCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseFusionMethod string

const (
	NamespaceInstanceUpdateResponseFusionMethodMax NamespaceInstanceUpdateResponseFusionMethod = "max"
	NamespaceInstanceUpdateResponseFusionMethodRrf NamespaceInstanceUpdateResponseFusionMethod = "rrf"
)

func (r NamespaceInstanceUpdateResponseFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseFusionMethodMax, NamespaceInstanceUpdateResponseFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceUpdateResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                           `json:"vector" api:"required"`
	JSON   namespaceInstanceUpdateResponseIndexMethodJSON `json:"-"`
}

// namespaceInstanceUpdateResponseIndexMethodJSON contains the JSON metadata for
// the struct [NamespaceInstanceUpdateResponseIndexMethod]
type namespaceInstanceUpdateResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             namespaceInstanceUpdateResponseIndexingOptionsJSON             `json:"-"`
}

// namespaceInstanceUpdateResponseIndexingOptionsJSON contains the JSON metadata
// for the struct [NamespaceInstanceUpdateResponseIndexingOptions]
type namespaceInstanceUpdateResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceUpdateResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseMetadata struct {
	CreatedFromAISearchWizard bool                                        `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                                      `json:"worker_domain"`
	JSON                      namespaceInstanceUpdateResponseMetadataJSON `json:"-"`
}

// namespaceInstanceUpdateResponseMetadataJSON contains the JSON metadata for the
// struct [NamespaceInstanceUpdateResponseMetadata]
type namespaceInstanceUpdateResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                                   `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled bool                                                              `json:"default_domain_enabled"`
	Enabled              bool                                                              `json:"enabled"`
	Mcp                  NamespaceInstanceUpdateResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       NamespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 namespaceInstanceUpdateResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponsePublicEndpointParams]
type namespaceInstanceUpdateResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                           `json:"disabled"`
	JSON     namespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                                       `json:"disabled"`
	JSON     namespaceInstanceUpdateResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsMcpJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponsePublicEndpointParamsMcp]
type namespaceInstanceUpdateResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                                 `json:"period_ms"`
	Requests  int64                                                                 `json:"requests"`
	Technique NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceInstanceUpdateResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsRateLimitJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimit]
type namespaceInstanceUpdateResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                                  `json:"disabled"`
	JSON     namespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpointJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpoint]
type namespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []NamespaceInstanceUpdateResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             namespaceInstanceUpdateResponseRetrievalOptionsJSON             `json:"-"`
}

// namespaceInstanceUpdateResponseRetrievalOptionsJSON contains the JSON metadata
// for the struct [NamespaceInstanceUpdateResponseRetrievalOptions]
type namespaceInstanceUpdateResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      namespaceInstanceUpdateResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceUpdateResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponseRetrievalOptionsBoostBy]
type namespaceInstanceUpdateResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionExists    NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionExists, NamespaceInstanceUpdateResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceUpdateResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   []string                                              `json:"include_items"`
	Prefix         string                                                `json:"prefix"`
	R2Jurisdiction string                                                `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceUpdateResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           namespaceInstanceUpdateResponseSourceParamsJSON       `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsJSON contains the JSON metadata for
// the struct [NamespaceInstanceUpdateResponseSourceParams]
type namespaceInstanceUpdateResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponseSourceParamsWebCrawler]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth float64 `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks bool `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit float64 `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge float64 `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   namespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                              `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                  `json:"use_browser_rendering"`
	JSON                namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                               `json:"selector" api:"required"`
	JSON     namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceUpdateResponseSyncInterval float64

const (
	NamespaceInstanceUpdateResponseSyncInterval900   NamespaceInstanceUpdateResponseSyncInterval = 900
	NamespaceInstanceUpdateResponseSyncInterval1800  NamespaceInstanceUpdateResponseSyncInterval = 1800
	NamespaceInstanceUpdateResponseSyncInterval3600  NamespaceInstanceUpdateResponseSyncInterval = 3600
	NamespaceInstanceUpdateResponseSyncInterval7200  NamespaceInstanceUpdateResponseSyncInterval = 7200
	NamespaceInstanceUpdateResponseSyncInterval14400 NamespaceInstanceUpdateResponseSyncInterval = 14400
	NamespaceInstanceUpdateResponseSyncInterval21600 NamespaceInstanceUpdateResponseSyncInterval = 21600
	NamespaceInstanceUpdateResponseSyncInterval43200 NamespaceInstanceUpdateResponseSyncInterval = 43200
	NamespaceInstanceUpdateResponseSyncInterval86400 NamespaceInstanceUpdateResponseSyncInterval = 86400
)

func (r NamespaceInstanceUpdateResponseSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseSyncInterval900, NamespaceInstanceUpdateResponseSyncInterval1800, NamespaceInstanceUpdateResponseSyncInterval3600, NamespaceInstanceUpdateResponseSyncInterval7200, NamespaceInstanceUpdateResponseSyncInterval14400, NamespaceInstanceUpdateResponseSyncInterval21600, NamespaceInstanceUpdateResponseSyncInterval43200, NamespaceInstanceUpdateResponseSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseType string

const (
	NamespaceInstanceUpdateResponseTypeR2         NamespaceInstanceUpdateResponseType = "r2"
	NamespaceInstanceUpdateResponseTypeWebCrawler NamespaceInstanceUpdateResponseType = "web-crawler"
)

func (r NamespaceInstanceUpdateResponseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseTypeR2, NamespaceInstanceUpdateResponseTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceListResponse struct {
	ID                             string                                            `json:"id" api:"required"`
	AIGatewayID                    string                                            `json:"ai_gateway_id" api:"required,nullable"`
	AISearchModel                  string                                            `json:"ai_search_model" api:"required,nullable"`
	Cache                          bool                                              `json:"cache" api:"required"`
	CacheThreshold                 NamespaceInstanceListResponseCacheThreshold       `json:"cache_threshold" api:"required,nullable"`
	CacheTTL                       NamespaceInstanceListResponseCacheTTL             `json:"cache_ttl" api:"required"`
	Chunk                          bool                                              `json:"chunk" api:"required"`
	ChunkOverlap                   float64                                           `json:"chunk_overlap" api:"required,nullable"`
	ChunkSize                      float64                                           `json:"chunk_size" api:"required,nullable"`
	CreatedAt                      time.Time                                         `json:"created_at" api:"required" format:"date-time"`
	CreatedBy                      string                                            `json:"created_by" api:"required,nullable"`
	CustomMetadata                 []NamespaceInstanceListResponseCustomMetadata     `json:"custom_metadata" api:"required,nullable"`
	EmbeddingModel                 string                                            `json:"embedding_model" api:"required,nullable"`
	Enable                         bool                                              `json:"enable" api:"required"`
	EngineVersion                  float64                                           `json:"engine_version" api:"required"`
	FusionMethod                   NamespaceInstanceListResponseFusionMethod         `json:"fusion_method" api:"required"`
	HybridSearchEnabled            bool                                              `json:"hybrid_search_enabled" api:"required"`
	IndexMethod                    NamespaceInstanceListResponseIndexMethod          `json:"index_method" api:"required"`
	IndexingOptions                NamespaceInstanceListResponseIndexingOptions      `json:"indexing_options" api:"required,nullable"`
	LastActivity                   time.Time                                         `json:"last_activity" api:"required,nullable" format:"date-time"`
	MaxNumResults                  float64                                           `json:"max_num_results" api:"required,nullable"`
	Metadata                       NamespaceInstanceListResponseMetadata             `json:"metadata" api:"required,nullable"`
	ModifiedAt                     time.Time                                         `json:"modified_at" api:"required" format:"date-time"`
	ModifiedBy                     string                                            `json:"modified_by" api:"required,nullable"`
	Namespace                      string                                            `json:"namespace" api:"required"`
	Paused                         bool                                              `json:"paused" api:"required"`
	PublicEndpointID               string                                            `json:"public_endpoint_id" api:"required,nullable"`
	PublicEndpointParams           NamespaceInstanceListResponsePublicEndpointParams `json:"public_endpoint_params" api:"required,nullable"`
	Reranking                      bool                                              `json:"reranking" api:"required"`
	RerankingModel                 string                                            `json:"reranking_model" api:"required,nullable"`
	RetrievalOptions               NamespaceInstanceListResponseRetrievalOptions     `json:"retrieval_options" api:"required,nullable"`
	RewriteModel                   string                                            `json:"rewrite_model" api:"required,nullable"`
	RewriteQuery                   bool                                              `json:"rewrite_query" api:"required"`
	ScoreThreshold                 float64                                           `json:"score_threshold" api:"required,nullable"`
	Source                         string                                            `json:"source" api:"required,nullable"`
	SourceParams                   NamespaceInstanceListResponseSourceParams         `json:"source_params" api:"required,nullable"`
	Status                         string                                            `json:"status" api:"required"`
	Summarization                  bool                                              `json:"summarization" api:"required"`
	SummarizationModel             string                                            `json:"summarization_model" api:"required,nullable"`
	SyncInterval                   NamespaceInstanceListResponseSyncInterval         `json:"sync_interval" api:"required"`
	SystemPromptAISearch           string                                            `json:"system_prompt_ai_search" api:"required,nullable"`
	SystemPromptIndexSummarization string                                            `json:"system_prompt_index_summarization" api:"required,nullable"`
	SystemPromptRewriteQuery       string                                            `json:"system_prompt_rewrite_query" api:"required,nullable"`
	TokenID                        string                                            `json:"token_id" api:"required,nullable"`
	Type                           NamespaceInstanceListResponseType                 `json:"type" api:"required,nullable"`
	JSON                           namespaceInstanceListResponseJSON                 `json:"-"`
}

// namespaceInstanceListResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceListResponse]
type namespaceInstanceListResponseJSON struct {
	ID                             apijson.Field
	AIGatewayID                    apijson.Field
	AISearchModel                  apijson.Field
	Cache                          apijson.Field
	CacheThreshold                 apijson.Field
	CacheTTL                       apijson.Field
	Chunk                          apijson.Field
	ChunkOverlap                   apijson.Field
	ChunkSize                      apijson.Field
	CreatedAt                      apijson.Field
	CreatedBy                      apijson.Field
	CustomMetadata                 apijson.Field
	EmbeddingModel                 apijson.Field
	Enable                         apijson.Field
	EngineVersion                  apijson.Field
	FusionMethod                   apijson.Field
	HybridSearchEnabled            apijson.Field
	IndexMethod                    apijson.Field
	IndexingOptions                apijson.Field
	LastActivity                   apijson.Field
	MaxNumResults                  apijson.Field
	Metadata                       apijson.Field
	ModifiedAt                     apijson.Field
	ModifiedBy                     apijson.Field
	Namespace                      apijson.Field
	Paused                         apijson.Field
	PublicEndpointID               apijson.Field
	PublicEndpointParams           apijson.Field
	Reranking                      apijson.Field
	RerankingModel                 apijson.Field
	RetrievalOptions               apijson.Field
	RewriteModel                   apijson.Field
	RewriteQuery                   apijson.Field
	ScoreThreshold                 apijson.Field
	Source                         apijson.Field
	SourceParams                   apijson.Field
	Status                         apijson.Field
	Summarization                  apijson.Field
	SummarizationModel             apijson.Field
	SyncInterval                   apijson.Field
	SystemPromptAISearch           apijson.Field
	SystemPromptIndexSummarization apijson.Field
	SystemPromptRewriteQuery       apijson.Field
	TokenID                        apijson.Field
	Type                           apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *NamespaceInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseCacheThreshold string

const (
	NamespaceInstanceListResponseCacheThresholdSuperStrictMatch NamespaceInstanceListResponseCacheThreshold = "super_strict_match"
	NamespaceInstanceListResponseCacheThresholdCloseEnough      NamespaceInstanceListResponseCacheThreshold = "close_enough"
	NamespaceInstanceListResponseCacheThresholdFlexibleFriend   NamespaceInstanceListResponseCacheThreshold = "flexible_friend"
	NamespaceInstanceListResponseCacheThresholdAnythingGoes     NamespaceInstanceListResponseCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceListResponseCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseCacheThresholdSuperStrictMatch, NamespaceInstanceListResponseCacheThresholdCloseEnough, NamespaceInstanceListResponseCacheThresholdFlexibleFriend, NamespaceInstanceListResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type NamespaceInstanceListResponseCacheTTL float64

const (
	NamespaceInstanceListResponseCacheTTL600    NamespaceInstanceListResponseCacheTTL = 600
	NamespaceInstanceListResponseCacheTTL1800   NamespaceInstanceListResponseCacheTTL = 1800
	NamespaceInstanceListResponseCacheTTL3600   NamespaceInstanceListResponseCacheTTL = 3600
	NamespaceInstanceListResponseCacheTTL7200   NamespaceInstanceListResponseCacheTTL = 7200
	NamespaceInstanceListResponseCacheTTL21600  NamespaceInstanceListResponseCacheTTL = 21600
	NamespaceInstanceListResponseCacheTTL43200  NamespaceInstanceListResponseCacheTTL = 43200
	NamespaceInstanceListResponseCacheTTL86400  NamespaceInstanceListResponseCacheTTL = 86400
	NamespaceInstanceListResponseCacheTTL172800 NamespaceInstanceListResponseCacheTTL = 172800
	NamespaceInstanceListResponseCacheTTL259200 NamespaceInstanceListResponseCacheTTL = 259200
	NamespaceInstanceListResponseCacheTTL518400 NamespaceInstanceListResponseCacheTTL = 518400
)

func (r NamespaceInstanceListResponseCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseCacheTTL600, NamespaceInstanceListResponseCacheTTL1800, NamespaceInstanceListResponseCacheTTL3600, NamespaceInstanceListResponseCacheTTL7200, NamespaceInstanceListResponseCacheTTL21600, NamespaceInstanceListResponseCacheTTL43200, NamespaceInstanceListResponseCacheTTL86400, NamespaceInstanceListResponseCacheTTL172800, NamespaceInstanceListResponseCacheTTL259200, NamespaceInstanceListResponseCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceListResponseCustomMetadata struct {
	DataType  NamespaceInstanceListResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                              `json:"field_name" api:"required"`
	JSON      namespaceInstanceListResponseCustomMetadataJSON     `json:"-"`
}

// namespaceInstanceListResponseCustomMetadataJSON contains the JSON metadata for
// the struct [NamespaceInstanceListResponseCustomMetadata]
type namespaceInstanceListResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseCustomMetadataDataType string

const (
	NamespaceInstanceListResponseCustomMetadataDataTypeText     NamespaceInstanceListResponseCustomMetadataDataType = "text"
	NamespaceInstanceListResponseCustomMetadataDataTypeNumber   NamespaceInstanceListResponseCustomMetadataDataType = "number"
	NamespaceInstanceListResponseCustomMetadataDataTypeBoolean  NamespaceInstanceListResponseCustomMetadataDataType = "boolean"
	NamespaceInstanceListResponseCustomMetadataDataTypeDatetime NamespaceInstanceListResponseCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceListResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseCustomMetadataDataTypeText, NamespaceInstanceListResponseCustomMetadataDataTypeNumber, NamespaceInstanceListResponseCustomMetadataDataTypeBoolean, NamespaceInstanceListResponseCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceListResponseFusionMethod string

const (
	NamespaceInstanceListResponseFusionMethodMax NamespaceInstanceListResponseFusionMethod = "max"
	NamespaceInstanceListResponseFusionMethodRrf NamespaceInstanceListResponseFusionMethod = "rrf"
)

func (r NamespaceInstanceListResponseFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseFusionMethodMax, NamespaceInstanceListResponseFusionMethodRrf:
		return true
	}
	return false
}

type NamespaceInstanceListResponseIndexMethod struct {
	Keyword     bool                                         `json:"keyword" api:"required"`
	Vector      bool                                         `json:"vector" api:"required"`
	ExtraFields map[string]interface{}                       `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponseIndexMethodJSON `json:"-"`
}

// namespaceInstanceListResponseIndexMethodJSON contains the JSON metadata for the
// struct [NamespaceInstanceListResponseIndexMethod]
type namespaceInstanceListResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseIndexingOptions struct {
	KeywordTokenizer NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	ExtraFields      map[string]interface{}                                       `json:"-" api:"extrafields"`
	JSON             namespaceInstanceListResponseIndexingOptionsJSON             `json:"-"`
}

// namespaceInstanceListResponseIndexingOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceListResponseIndexingOptions]
type namespaceInstanceListResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceListResponseIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceListResponseIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceListResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceListResponseMetadata struct {
	CreatedFromAISearchWizard bool                                      `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                                    `json:"worker_domain"`
	ExtraFields               map[string]interface{}                    `json:"-" api:"extrafields"`
	JSON                      namespaceInstanceListResponseMetadataJSON `json:"-"`
}

// namespaceInstanceListResponseMetadataJSON contains the JSON metadata for the
// struct [NamespaceInstanceListResponseMetadata]
type namespaceInstanceListResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                                 `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	CustomDomains           []string                                                                 `json:"custom_domains" api:"nullable"`
	DefaultDomainEnabled    bool                                                                     `json:"default_domain_enabled"`
	Enabled                 bool                                                                     `json:"enabled"`
	Mcp                     NamespaceInstanceListResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceListResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceListResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	ExtraFields             map[string]interface{}                                                   `json:"-" api:"extrafields"`
	JSON                    namespaceInstanceListResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceListResponsePublicEndpointParams]
type namespaceInstanceListResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceInstanceListResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	Disabled    bool                                                                         `json:"disabled"`
	ExtraFields map[string]interface{}                                                       `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponsePublicEndpointParamsMcp struct {
	Description string                                                   `json:"description"`
	Disabled    bool                                                     `json:"disabled"`
	ExtraFields map[string]interface{}                                   `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsMcpJSON contains the JSON
// metadata for the struct [NamespaceInstanceListResponsePublicEndpointParamsMcp]
type namespaceInstanceListResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponsePublicEndpointParamsRateLimit struct {
	PeriodMs    int64                                                               `json:"period_ms"`
	Requests    int64                                                               `json:"requests"`
	Technique   NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	ExtraFields map[string]interface{}                                              `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct
// [NamespaceInstanceListResponsePublicEndpointParamsRateLimit]
type namespaceInstanceListResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceListResponsePublicEndpointParamsSearchEndpoint struct {
	Disabled    bool                                                                `json:"disabled"`
	ExtraFields map[string]interface{}                                              `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsSearchEndpointJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceListResponsePublicEndpointParamsSearchEndpoint]
type namespaceInstanceListResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseRetrievalOptions struct {
	BoostBy          []NamespaceInstanceListResponseRetrievalOptionsBoostBy        `json:"boost_by"`
	KeywordMatchMode NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	ExtraFields      map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON             namespaceInstanceListResponseRetrievalOptionsJSON             `json:"-"`
}

// namespaceInstanceListResponseRetrievalOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceListResponseRetrievalOptions]
type namespaceInstanceListResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseRetrievalOptionsBoostBy struct {
	Field       string                                                        `json:"field" api:"required"`
	DataType    NamespaceInstanceListResponseRetrievalOptionsBoostByDataType  `json:"dataType"`
	Direction   NamespaceInstanceListResponseRetrievalOptionsBoostByDirection `json:"direction"`
	ExtraFields map[string]interface{}                                        `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceListResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceListResponseRetrievalOptionsBoostBy]
type namespaceInstanceListResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	DataType    apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseRetrievalOptionsBoostByDataType string

const (
	NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeNumber   NamespaceInstanceListResponseRetrievalOptionsBoostByDataType = "number"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeDatetime NamespaceInstanceListResponseRetrievalOptionsBoostByDataType = "datetime"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeText     NamespaceInstanceListResponseRetrievalOptionsBoostByDataType = "text"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeBoolean  NamespaceInstanceListResponseRetrievalOptionsBoostByDataType = "boolean"
)

func (r NamespaceInstanceListResponseRetrievalOptionsBoostByDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeNumber, NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeDatetime, NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeText, NamespaceInstanceListResponseRetrievalOptionsBoostByDataTypeBoolean:
		return true
	}
	return false
}

type NamespaceInstanceListResponseRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceListResponseRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceListResponseRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionExists    NamespaceInstanceListResponseRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceListResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceListResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionExists, NamespaceInstanceListResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

type NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceListResponseRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceListResponseRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceListResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSourceParams struct {
	ExcludeItems   []string                                            `json:"exclude_items"`
	IncludeItems   []string                                            `json:"include_items"`
	Prefix         string                                              `json:"prefix"`
	R2Jurisdiction string                                              `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceListResponseSourceParamsWebCrawler `json:"web_crawler"`
	ExtraFields    map[string]interface{}                              `json:"-" api:"extrafields"`
	JSON           namespaceInstanceListResponseSourceParamsJSON       `json:"-"`
}

// namespaceInstanceListResponseSourceParamsJSON contains the JSON metadata for the
// struct [NamespaceInstanceListResponseSourceParams]
type namespaceInstanceListResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawler struct {
	DiscoverOptions NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	ParseType       NamespaceInstanceListResponseSourceParamsWebCrawlerParseType       `json:"parse_type"`
	ExtraFields     map[string]interface{}                                             `json:"-" api:"extrafields"`
	JSON            namespaceInstanceListResponseSourceParamsWebCrawlerJSON            `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceListResponseSourceParamsWebCrawler]
type namespaceInstanceListResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptions struct {
	Depth                float64 `json:"depth"`
	IncludeExternalLinks bool    `json:"include_external_links"`
	IncludeSubdomains    bool    `json:"include_subdomains"`
	// Maximum number of pages to crawl. New values are capped at 100000; instances
	// configured before that cap may report a higher stored value, which the crawler
	// clamps at run time.
	Limit       float64                                                                  `json:"limit"`
	MaxAge      float64                                                                  `json:"max_age"`
	Source      NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	ExtraFields map[string]interface{}                                                   `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptions]
type namespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceListResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions struct {
	ContentSelector     []NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders      map[string]string                                                                `json:"include_headers"`
	IncludeImages       bool                                                                             `json:"include_images"`
	SpecificSitemaps    []string                                                                         `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                             `json:"use_browser_rendering"`
	ExtraFields         map[string]interface{}                                                           `json:"-" api:"extrafields"`
	JSON                namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsJSON              `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions]
type namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	Path        string                                                                             `json:"path" api:"required"`
	Selector    string                                                                             `json:"selector" api:"required"`
	ExtraFields map[string]interface{}                                                             `json:"-" api:"extrafields"`
	JSON        namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceListResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceListResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceListResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSyncInterval float64

const (
	NamespaceInstanceListResponseSyncInterval900   NamespaceInstanceListResponseSyncInterval = 900
	NamespaceInstanceListResponseSyncInterval1800  NamespaceInstanceListResponseSyncInterval = 1800
	NamespaceInstanceListResponseSyncInterval3600  NamespaceInstanceListResponseSyncInterval = 3600
	NamespaceInstanceListResponseSyncInterval7200  NamespaceInstanceListResponseSyncInterval = 7200
	NamespaceInstanceListResponseSyncInterval14400 NamespaceInstanceListResponseSyncInterval = 14400
	NamespaceInstanceListResponseSyncInterval21600 NamespaceInstanceListResponseSyncInterval = 21600
	NamespaceInstanceListResponseSyncInterval43200 NamespaceInstanceListResponseSyncInterval = 43200
	NamespaceInstanceListResponseSyncInterval86400 NamespaceInstanceListResponseSyncInterval = 86400
)

func (r NamespaceInstanceListResponseSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseSyncInterval900, NamespaceInstanceListResponseSyncInterval1800, NamespaceInstanceListResponseSyncInterval3600, NamespaceInstanceListResponseSyncInterval7200, NamespaceInstanceListResponseSyncInterval14400, NamespaceInstanceListResponseSyncInterval21600, NamespaceInstanceListResponseSyncInterval43200, NamespaceInstanceListResponseSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceListResponseType string

const (
	NamespaceInstanceListResponseTypeR2         NamespaceInstanceListResponseType = "r2"
	NamespaceInstanceListResponseTypeWebCrawler NamespaceInstanceListResponseType = "web-crawler"
)

func (r NamespaceInstanceListResponseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseTypeR2, NamespaceInstanceListResponseTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID string    `json:"ai_gateway_id" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  string                                        `json:"ai_search_model" api:"nullable"`
	Cache          bool                                          `json:"cache"`
	CacheThreshold NamespaceInstanceDeleteResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       NamespaceInstanceDeleteResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                           `json:"chunk_overlap"`
	ChunkSize      int64                                           `json:"chunk_size"`
	CreatedBy      string                                          `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceDeleteResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                          `json:"embedding_model" api:"nullable"`
	Enable         bool                                            `json:"enable"`
	EngineVersion  float64                                         `json:"engine_version"`
	FusionMethod   NamespaceInstanceDeleteResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          NamespaceInstanceDeleteResponseIndexMethod          `json:"index_method"`
	IndexingOptions      NamespaceInstanceDeleteResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                           `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                               `json:"max_num_results"`
	Metadata             NamespaceInstanceDeleteResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                              `json:"modified_by" api:"nullable"`
	Namespace            string                                              `json:"namespace" api:"nullable"`
	Paused               bool                                                `json:"paused"`
	PublicEndpointID     string                                              `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceInstanceDeleteResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                                `json:"reranking"`
	RerankingModel       string                                              `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceDeleteResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel   string                                      `json:"rewrite_model" api:"nullable"`
	RewriteQuery   bool                                        `json:"rewrite_query"`
	ScoreThreshold float64                                     `json:"score_threshold"`
	Source         string                                      `json:"source" api:"nullable"`
	SourceParams   NamespaceInstanceDeleteResponseSourceParams `json:"source_params" api:"nullable"`
	Status         string                                      `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval NamespaceInstanceDeleteResponseSyncInterval `json:"sync_interval"`
	TokenID      string                                      `json:"token_id" format:"uuid"`
	Type         NamespaceInstanceDeleteResponseType         `json:"type" api:"nullable"`
	JSON         namespaceInstanceDeleteResponseJSON         `json:"-"`
}

// namespaceInstanceDeleteResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceDeleteResponse]
type namespaceInstanceDeleteResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	AIGatewayID          apijson.Field
	AISearchModel        apijson.Field
	Cache                apijson.Field
	CacheThreshold       apijson.Field
	CacheTTL             apijson.Field
	ChunkOverlap         apijson.Field
	ChunkSize            apijson.Field
	CreatedBy            apijson.Field
	CustomMetadata       apijson.Field
	EmbeddingModel       apijson.Field
	Enable               apijson.Field
	EngineVersion        apijson.Field
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	IndexMethod          apijson.Field
	IndexingOptions      apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Namespace            apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	Source               apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	SyncInterval         apijson.Field
	TokenID              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseCacheThreshold string

const (
	NamespaceInstanceDeleteResponseCacheThresholdSuperStrictMatch NamespaceInstanceDeleteResponseCacheThreshold = "super_strict_match"
	NamespaceInstanceDeleteResponseCacheThresholdCloseEnough      NamespaceInstanceDeleteResponseCacheThreshold = "close_enough"
	NamespaceInstanceDeleteResponseCacheThresholdFlexibleFriend   NamespaceInstanceDeleteResponseCacheThreshold = "flexible_friend"
	NamespaceInstanceDeleteResponseCacheThresholdAnythingGoes     NamespaceInstanceDeleteResponseCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceDeleteResponseCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseCacheThresholdSuperStrictMatch, NamespaceInstanceDeleteResponseCacheThresholdCloseEnough, NamespaceInstanceDeleteResponseCacheThresholdFlexibleFriend, NamespaceInstanceDeleteResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceDeleteResponseCacheTTL float64

const (
	NamespaceInstanceDeleteResponseCacheTTL600    NamespaceInstanceDeleteResponseCacheTTL = 600
	NamespaceInstanceDeleteResponseCacheTTL1800   NamespaceInstanceDeleteResponseCacheTTL = 1800
	NamespaceInstanceDeleteResponseCacheTTL3600   NamespaceInstanceDeleteResponseCacheTTL = 3600
	NamespaceInstanceDeleteResponseCacheTTL7200   NamespaceInstanceDeleteResponseCacheTTL = 7200
	NamespaceInstanceDeleteResponseCacheTTL21600  NamespaceInstanceDeleteResponseCacheTTL = 21600
	NamespaceInstanceDeleteResponseCacheTTL43200  NamespaceInstanceDeleteResponseCacheTTL = 43200
	NamespaceInstanceDeleteResponseCacheTTL86400  NamespaceInstanceDeleteResponseCacheTTL = 86400
	NamespaceInstanceDeleteResponseCacheTTL172800 NamespaceInstanceDeleteResponseCacheTTL = 172800
	NamespaceInstanceDeleteResponseCacheTTL259200 NamespaceInstanceDeleteResponseCacheTTL = 259200
	NamespaceInstanceDeleteResponseCacheTTL518400 NamespaceInstanceDeleteResponseCacheTTL = 518400
)

func (r NamespaceInstanceDeleteResponseCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseCacheTTL600, NamespaceInstanceDeleteResponseCacheTTL1800, NamespaceInstanceDeleteResponseCacheTTL3600, NamespaceInstanceDeleteResponseCacheTTL7200, NamespaceInstanceDeleteResponseCacheTTL21600, NamespaceInstanceDeleteResponseCacheTTL43200, NamespaceInstanceDeleteResponseCacheTTL86400, NamespaceInstanceDeleteResponseCacheTTL172800, NamespaceInstanceDeleteResponseCacheTTL259200, NamespaceInstanceDeleteResponseCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseCustomMetadata struct {
	DataType  NamespaceInstanceDeleteResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                                `json:"field_name" api:"required"`
	JSON      namespaceInstanceDeleteResponseCustomMetadataJSON     `json:"-"`
}

// namespaceInstanceDeleteResponseCustomMetadataJSON contains the JSON metadata for
// the struct [NamespaceInstanceDeleteResponseCustomMetadata]
type namespaceInstanceDeleteResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseCustomMetadataDataType string

const (
	NamespaceInstanceDeleteResponseCustomMetadataDataTypeText     NamespaceInstanceDeleteResponseCustomMetadataDataType = "text"
	NamespaceInstanceDeleteResponseCustomMetadataDataTypeNumber   NamespaceInstanceDeleteResponseCustomMetadataDataType = "number"
	NamespaceInstanceDeleteResponseCustomMetadataDataTypeBoolean  NamespaceInstanceDeleteResponseCustomMetadataDataType = "boolean"
	NamespaceInstanceDeleteResponseCustomMetadataDataTypeDatetime NamespaceInstanceDeleteResponseCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceDeleteResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseCustomMetadataDataTypeText, NamespaceInstanceDeleteResponseCustomMetadataDataTypeNumber, NamespaceInstanceDeleteResponseCustomMetadataDataTypeBoolean, NamespaceInstanceDeleteResponseCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseFusionMethod string

const (
	NamespaceInstanceDeleteResponseFusionMethodMax NamespaceInstanceDeleteResponseFusionMethod = "max"
	NamespaceInstanceDeleteResponseFusionMethodRrf NamespaceInstanceDeleteResponseFusionMethod = "rrf"
)

func (r NamespaceInstanceDeleteResponseFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseFusionMethodMax, NamespaceInstanceDeleteResponseFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceDeleteResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                           `json:"vector" api:"required"`
	JSON   namespaceInstanceDeleteResponseIndexMethodJSON `json:"-"`
}

// namespaceInstanceDeleteResponseIndexMethodJSON contains the JSON metadata for
// the struct [NamespaceInstanceDeleteResponseIndexMethod]
type namespaceInstanceDeleteResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             namespaceInstanceDeleteResponseIndexingOptionsJSON             `json:"-"`
}

// namespaceInstanceDeleteResponseIndexingOptionsJSON contains the JSON metadata
// for the struct [NamespaceInstanceDeleteResponseIndexingOptions]
type namespaceInstanceDeleteResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceDeleteResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseMetadata struct {
	CreatedFromAISearchWizard bool                                        `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                                      `json:"worker_domain"`
	JSON                      namespaceInstanceDeleteResponseMetadataJSON `json:"-"`
}

// namespaceInstanceDeleteResponseMetadataJSON contains the JSON metadata for the
// struct [NamespaceInstanceDeleteResponseMetadata]
type namespaceInstanceDeleteResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                                   `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled bool                                                              `json:"default_domain_enabled"`
	Enabled              bool                                                              `json:"enabled"`
	Mcp                  NamespaceInstanceDeleteResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       NamespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 namespaceInstanceDeleteResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponsePublicEndpointParams]
type namespaceInstanceDeleteResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                           `json:"disabled"`
	JSON     namespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                                       `json:"disabled"`
	JSON     namespaceInstanceDeleteResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsMcpJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponsePublicEndpointParamsMcp]
type namespaceInstanceDeleteResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                                 `json:"period_ms"`
	Requests  int64                                                                 `json:"requests"`
	Technique NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceInstanceDeleteResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsRateLimitJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimit]
type namespaceInstanceDeleteResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                                  `json:"disabled"`
	JSON     namespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpointJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpoint]
type namespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []NamespaceInstanceDeleteResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             namespaceInstanceDeleteResponseRetrievalOptionsJSON             `json:"-"`
}

// namespaceInstanceDeleteResponseRetrievalOptionsJSON contains the JSON metadata
// for the struct [NamespaceInstanceDeleteResponseRetrievalOptions]
type namespaceInstanceDeleteResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      namespaceInstanceDeleteResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceDeleteResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponseRetrievalOptionsBoostBy]
type namespaceInstanceDeleteResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionExists    NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionExists, NamespaceInstanceDeleteResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceDeleteResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   []string                                              `json:"include_items"`
	Prefix         string                                                `json:"prefix"`
	R2Jurisdiction string                                                `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceDeleteResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           namespaceInstanceDeleteResponseSourceParamsJSON       `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsJSON contains the JSON metadata for
// the struct [NamespaceInstanceDeleteResponseSourceParams]
type namespaceInstanceDeleteResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponseSourceParamsWebCrawler]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth float64 `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks bool `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit float64 `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge float64 `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   namespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                              `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                  `json:"use_browser_rendering"`
	JSON                namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                               `json:"selector" api:"required"`
	JSON     namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceDeleteResponseSyncInterval float64

const (
	NamespaceInstanceDeleteResponseSyncInterval900   NamespaceInstanceDeleteResponseSyncInterval = 900
	NamespaceInstanceDeleteResponseSyncInterval1800  NamespaceInstanceDeleteResponseSyncInterval = 1800
	NamespaceInstanceDeleteResponseSyncInterval3600  NamespaceInstanceDeleteResponseSyncInterval = 3600
	NamespaceInstanceDeleteResponseSyncInterval7200  NamespaceInstanceDeleteResponseSyncInterval = 7200
	NamespaceInstanceDeleteResponseSyncInterval14400 NamespaceInstanceDeleteResponseSyncInterval = 14400
	NamespaceInstanceDeleteResponseSyncInterval21600 NamespaceInstanceDeleteResponseSyncInterval = 21600
	NamespaceInstanceDeleteResponseSyncInterval43200 NamespaceInstanceDeleteResponseSyncInterval = 43200
	NamespaceInstanceDeleteResponseSyncInterval86400 NamespaceInstanceDeleteResponseSyncInterval = 86400
)

func (r NamespaceInstanceDeleteResponseSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseSyncInterval900, NamespaceInstanceDeleteResponseSyncInterval1800, NamespaceInstanceDeleteResponseSyncInterval3600, NamespaceInstanceDeleteResponseSyncInterval7200, NamespaceInstanceDeleteResponseSyncInterval14400, NamespaceInstanceDeleteResponseSyncInterval21600, NamespaceInstanceDeleteResponseSyncInterval43200, NamespaceInstanceDeleteResponseSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseType string

const (
	NamespaceInstanceDeleteResponseTypeR2         NamespaceInstanceDeleteResponseType = "r2"
	NamespaceInstanceDeleteResponseTypeWebCrawler NamespaceInstanceDeleteResponseType = "web-crawler"
)

func (r NamespaceInstanceDeleteResponseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseTypeR2, NamespaceInstanceDeleteResponseTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsResponse struct {
	Choices     []NamespaceInstanceChatCompletionsResponseChoice `json:"choices" api:"required"`
	Chunks      []NamespaceInstanceChatCompletionsResponseChunk  `json:"chunks" api:"required"`
	ID          string                                           `json:"id"`
	Model       string                                           `json:"model"`
	Object      string                                           `json:"object"`
	ExtraFields map[string]interface{}                           `json:"-" api:"extrafields"`
	JSON        namespaceInstanceChatCompletionsResponseJSON     `json:"-"`
}

// namespaceInstanceChatCompletionsResponseJSON contains the JSON metadata for the
// struct [NamespaceInstanceChatCompletionsResponse]
type namespaceInstanceChatCompletionsResponseJSON struct {
	Choices     apijson.Field
	Chunks      apijson.Field
	ID          apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsResponseChoice struct {
	Message NamespaceInstanceChatCompletionsResponseChoicesMessage `json:"message" api:"required"`
	Index   int64                                                  `json:"index"`
	JSON    namespaceInstanceChatCompletionsResponseChoiceJSON     `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChoiceJSON contains the JSON metadata
// for the struct [NamespaceInstanceChatCompletionsResponseChoice]
type namespaceInstanceChatCompletionsResponseChoiceJSON struct {
	Message     apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChoiceJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsResponseChoicesMessage struct {
	Content     NamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion `json:"content" api:"required"`
	Role        NamespaceInstanceChatCompletionsResponseChoicesMessageRole         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                             `json:"-" api:"extrafields"`
	JSON        namespaceInstanceChatCompletionsResponseChoicesMessageJSON         `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChoicesMessageJSON contains the JSON
// metadata for the struct [NamespaceInstanceChatCompletionsResponseChoicesMessage]
type namespaceInstanceChatCompletionsResponseChoicesMessageJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChoicesMessageJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString],
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArray] or
// [shared.UnionString].
type NamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion interface {
	ImplementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceInstanceChatCompletionsResponseChoicesMessageContentArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArray []NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem

func (r NamespaceInstanceChatCompletionsResponseChoicesMessageContentArray) ImplementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentUnion() {
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem struct {
	Type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType `json:"type" api:"required"`
	// This field can have the runtime type of
	// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectFile].
	File interface{} `json:"file"`
	// This field can have the runtime type of
	// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectImageURL].
	ImageURL interface{}                                                                `json:"image_url"`
	Text     string                                                                     `json:"text"`
	JSON     namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItemJSON `json:"-"`
	union    NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem
}

// namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItemJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem]
type namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItemJSON struct {
	Type        apijson.Field
	File        apijson.Field
	ImageURL    apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItemJSON) RawJSON() string {
	return r.raw
}

func (r *NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem) UnmarshalJSON(data []byte) (err error) {
	*r = NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject].
func (r NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem) AsUnion() NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem {
	return r.union
}

// Union satisfied by
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject],
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject] or
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject].
type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem interface {
	implementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject{}),
		},
	)
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject struct {
	Text string                                                                       `json:"text" api:"required"`
	Type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectType `json:"type" api:"required"`
	JSON namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject]
type namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectJSON) RawJSON() string {
	return r.raw
}

func (r NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObject) implementsNamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayItem() {
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectType string

const (
	NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectType = "text"
)

func (r NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType string

const (
	NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeText     NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType = "text"
	NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType = "image_url"
	NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeFile     NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType = "file"
)

func (r NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeText, NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeImageURL, NamespaceInstanceChatCompletionsResponseChoicesMessageContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsResponseChoicesMessageRole string

const (
	NamespaceInstanceChatCompletionsResponseChoicesMessageRoleSystem    NamespaceInstanceChatCompletionsResponseChoicesMessageRole = "system"
	NamespaceInstanceChatCompletionsResponseChoicesMessageRoleDeveloper NamespaceInstanceChatCompletionsResponseChoicesMessageRole = "developer"
	NamespaceInstanceChatCompletionsResponseChoicesMessageRoleUser      NamespaceInstanceChatCompletionsResponseChoicesMessageRole = "user"
	NamespaceInstanceChatCompletionsResponseChoicesMessageRoleAssistant NamespaceInstanceChatCompletionsResponseChoicesMessageRole = "assistant"
	NamespaceInstanceChatCompletionsResponseChoicesMessageRoleTool      NamespaceInstanceChatCompletionsResponseChoicesMessageRole = "tool"
)

func (r NamespaceInstanceChatCompletionsResponseChoicesMessageRole) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsResponseChoicesMessageRoleSystem, NamespaceInstanceChatCompletionsResponseChoicesMessageRoleDeveloper, NamespaceInstanceChatCompletionsResponseChoicesMessageRoleUser, NamespaceInstanceChatCompletionsResponseChoicesMessageRoleAssistant, NamespaceInstanceChatCompletionsResponseChoicesMessageRoleTool:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsResponseChunk struct {
	ID             string                                                       `json:"id" api:"required"`
	Score          float64                                                      `json:"score" api:"required"`
	Text           string                                                       `json:"text" api:"required"`
	Type           string                                                       `json:"type" api:"required"`
	Item           NamespaceInstanceChatCompletionsResponseChunksItem           `json:"item"`
	ScoringDetails NamespaceInstanceChatCompletionsResponseChunksScoringDetails `json:"scoring_details"`
	JSON           namespaceInstanceChatCompletionsResponseChunkJSON            `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChunkJSON contains the JSON metadata for
// the struct [NamespaceInstanceChatCompletionsResponseChunk]
type namespaceInstanceChatCompletionsResponseChunkJSON struct {
	ID             apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChunkJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsResponseChunksItem struct {
	Key       string                                                 `json:"key" api:"required"`
	Metadata  map[string]interface{}                                 `json:"metadata"`
	Timestamp float64                                                `json:"timestamp"`
	JSON      namespaceInstanceChatCompletionsResponseChunksItemJSON `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChunksItemJSON contains the JSON
// metadata for the struct [NamespaceInstanceChatCompletionsResponseChunksItem]
type namespaceInstanceChatCompletionsResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsResponseChunksScoringDetails struct {
	FusionMethod   NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                                  `json:"keyword_rank"`
	KeywordScore   float64                                                                  `json:"keyword_score"`
	RerankingScore float64                                                                  `json:"reranking_score"`
	VectorRank     float64                                                                  `json:"vector_rank"`
	VectorScore    float64                                                                  `json:"vector_score"`
	JSON           namespaceInstanceChatCompletionsResponseChunksScoringDetailsJSON         `json:"-"`
}

// namespaceInstanceChatCompletionsResponseChunksScoringDetailsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceChatCompletionsResponseChunksScoringDetails]
type namespaceInstanceChatCompletionsResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceChatCompletionsResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceChatCompletionsResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethod string

const (
	NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethod = "rrf"
	NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethodMax NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethod = "max"
)

func (r NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethodRrf, NamespaceInstanceChatCompletionsResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type NamespaceInstanceReadResponse struct {
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt  time.Time `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID string    `json:"ai_gateway_id" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  string                                      `json:"ai_search_model" api:"nullable"`
	Cache          bool                                        `json:"cache"`
	CacheThreshold NamespaceInstanceReadResponseCacheThreshold `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       NamespaceInstanceReadResponseCacheTTL         `json:"cache_ttl"`
	ChunkOverlap   int64                                         `json:"chunk_overlap"`
	ChunkSize      int64                                         `json:"chunk_size"`
	CreatedBy      string                                        `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceReadResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel string                                        `json:"embedding_model" api:"nullable"`
	Enable         bool                                          `json:"enable"`
	EngineVersion  float64                                       `json:"engine_version"`
	FusionMethod   NamespaceInstanceReadResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          NamespaceInstanceReadResponseIndexMethod          `json:"index_method"`
	IndexingOptions      NamespaceInstanceReadResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                         `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                             `json:"max_num_results"`
	Metadata             NamespaceInstanceReadResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                            `json:"modified_by" api:"nullable"`
	Namespace            string                                            `json:"namespace" api:"nullable"`
	Paused               bool                                              `json:"paused"`
	PublicEndpointID     string                                            `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceInstanceReadResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                              `json:"reranking"`
	RerankingModel       string                                            `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceReadResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel   string                                    `json:"rewrite_model" api:"nullable"`
	RewriteQuery   bool                                      `json:"rewrite_query"`
	ScoreThreshold float64                                   `json:"score_threshold"`
	Source         string                                    `json:"source" api:"nullable"`
	SourceParams   NamespaceInstanceReadResponseSourceParams `json:"source_params" api:"nullable"`
	Status         string                                    `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval NamespaceInstanceReadResponseSyncInterval `json:"sync_interval"`
	TokenID      string                                    `json:"token_id" format:"uuid"`
	Type         NamespaceInstanceReadResponseType         `json:"type" api:"nullable"`
	JSON         namespaceInstanceReadResponseJSON         `json:"-"`
}

// namespaceInstanceReadResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceReadResponse]
type namespaceInstanceReadResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	AIGatewayID          apijson.Field
	AISearchModel        apijson.Field
	Cache                apijson.Field
	CacheThreshold       apijson.Field
	CacheTTL             apijson.Field
	ChunkOverlap         apijson.Field
	ChunkSize            apijson.Field
	CreatedBy            apijson.Field
	CustomMetadata       apijson.Field
	EmbeddingModel       apijson.Field
	Enable               apijson.Field
	EngineVersion        apijson.Field
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	IndexMethod          apijson.Field
	IndexingOptions      apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Namespace            apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	Source               apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	SyncInterval         apijson.Field
	TokenID              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseCacheThreshold string

const (
	NamespaceInstanceReadResponseCacheThresholdSuperStrictMatch NamespaceInstanceReadResponseCacheThreshold = "super_strict_match"
	NamespaceInstanceReadResponseCacheThresholdCloseEnough      NamespaceInstanceReadResponseCacheThreshold = "close_enough"
	NamespaceInstanceReadResponseCacheThresholdFlexibleFriend   NamespaceInstanceReadResponseCacheThreshold = "flexible_friend"
	NamespaceInstanceReadResponseCacheThresholdAnythingGoes     NamespaceInstanceReadResponseCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceReadResponseCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseCacheThresholdSuperStrictMatch, NamespaceInstanceReadResponseCacheThresholdCloseEnough, NamespaceInstanceReadResponseCacheThresholdFlexibleFriend, NamespaceInstanceReadResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceReadResponseCacheTTL float64

const (
	NamespaceInstanceReadResponseCacheTTL600    NamespaceInstanceReadResponseCacheTTL = 600
	NamespaceInstanceReadResponseCacheTTL1800   NamespaceInstanceReadResponseCacheTTL = 1800
	NamespaceInstanceReadResponseCacheTTL3600   NamespaceInstanceReadResponseCacheTTL = 3600
	NamespaceInstanceReadResponseCacheTTL7200   NamespaceInstanceReadResponseCacheTTL = 7200
	NamespaceInstanceReadResponseCacheTTL21600  NamespaceInstanceReadResponseCacheTTL = 21600
	NamespaceInstanceReadResponseCacheTTL43200  NamespaceInstanceReadResponseCacheTTL = 43200
	NamespaceInstanceReadResponseCacheTTL86400  NamespaceInstanceReadResponseCacheTTL = 86400
	NamespaceInstanceReadResponseCacheTTL172800 NamespaceInstanceReadResponseCacheTTL = 172800
	NamespaceInstanceReadResponseCacheTTL259200 NamespaceInstanceReadResponseCacheTTL = 259200
	NamespaceInstanceReadResponseCacheTTL518400 NamespaceInstanceReadResponseCacheTTL = 518400
)

func (r NamespaceInstanceReadResponseCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseCacheTTL600, NamespaceInstanceReadResponseCacheTTL1800, NamespaceInstanceReadResponseCacheTTL3600, NamespaceInstanceReadResponseCacheTTL7200, NamespaceInstanceReadResponseCacheTTL21600, NamespaceInstanceReadResponseCacheTTL43200, NamespaceInstanceReadResponseCacheTTL86400, NamespaceInstanceReadResponseCacheTTL172800, NamespaceInstanceReadResponseCacheTTL259200, NamespaceInstanceReadResponseCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseCustomMetadata struct {
	DataType  NamespaceInstanceReadResponseCustomMetadataDataType `json:"data_type" api:"required"`
	FieldName string                                              `json:"field_name" api:"required"`
	JSON      namespaceInstanceReadResponseCustomMetadataJSON     `json:"-"`
}

// namespaceInstanceReadResponseCustomMetadataJSON contains the JSON metadata for
// the struct [NamespaceInstanceReadResponseCustomMetadata]
type namespaceInstanceReadResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseCustomMetadataDataType string

const (
	NamespaceInstanceReadResponseCustomMetadataDataTypeText     NamespaceInstanceReadResponseCustomMetadataDataType = "text"
	NamespaceInstanceReadResponseCustomMetadataDataTypeNumber   NamespaceInstanceReadResponseCustomMetadataDataType = "number"
	NamespaceInstanceReadResponseCustomMetadataDataTypeBoolean  NamespaceInstanceReadResponseCustomMetadataDataType = "boolean"
	NamespaceInstanceReadResponseCustomMetadataDataTypeDatetime NamespaceInstanceReadResponseCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceReadResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseCustomMetadataDataTypeText, NamespaceInstanceReadResponseCustomMetadataDataTypeNumber, NamespaceInstanceReadResponseCustomMetadataDataTypeBoolean, NamespaceInstanceReadResponseCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseFusionMethod string

const (
	NamespaceInstanceReadResponseFusionMethodMax NamespaceInstanceReadResponseFusionMethod = "max"
	NamespaceInstanceReadResponseFusionMethodRrf NamespaceInstanceReadResponseFusionMethod = "rrf"
)

func (r NamespaceInstanceReadResponseFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseFusionMethodMax, NamespaceInstanceReadResponseFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceReadResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                         `json:"vector" api:"required"`
	JSON   namespaceInstanceReadResponseIndexMethodJSON `json:"-"`
}

// namespaceInstanceReadResponseIndexMethodJSON contains the JSON metadata for the
// struct [NamespaceInstanceReadResponseIndexMethod]
type namespaceInstanceReadResponseIndexMethodJSON struct {
	Keyword     apijson.Field
	Vector      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseIndexMethod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseIndexMethodJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
	JSON             namespaceInstanceReadResponseIndexingOptionsJSON             `json:"-"`
}

// namespaceInstanceReadResponseIndexingOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceReadResponseIndexingOptions]
type namespaceInstanceReadResponseIndexingOptionsJSON struct {
	KeywordTokenizer apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseIndexingOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseIndexingOptionsJSON) RawJSON() string {
	return r.raw
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceReadResponseIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseMetadata struct {
	CreatedFromAISearchWizard bool                                      `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                                    `json:"worker_domain"`
	JSON                      namespaceInstanceReadResponseMetadataJSON `json:"-"`
}

// namespaceInstanceReadResponseMetadataJSON contains the JSON metadata for the
// struct [NamespaceInstanceReadResponseMetadata]
type namespaceInstanceReadResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                                 `json:"authorized_hosts"`
	ChatCompletionsEndpoint NamespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled bool                                                            `json:"default_domain_enabled"`
	Enabled              bool                                                            `json:"enabled"`
	Mcp                  NamespaceInstanceReadResponsePublicEndpointParamsMcp            `json:"mcp"`
	RateLimit            NamespaceInstanceReadResponsePublicEndpointParamsRateLimit      `json:"rate_limit"`
	SearchEndpoint       NamespaceInstanceReadResponsePublicEndpointParamsSearchEndpoint `json:"search_endpoint"`
	JSON                 namespaceInstanceReadResponsePublicEndpointParamsJSON           `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceReadResponsePublicEndpointParams]
type namespaceInstanceReadResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	CustomDomains           apijson.Field
	DefaultDomainEnabled    apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                         `json:"disabled"`
	JSON     namespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint]
type namespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                                     `json:"disabled"`
	JSON     namespaceInstanceReadResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsMcpJSON contains the JSON
// metadata for the struct [NamespaceInstanceReadResponsePublicEndpointParamsMcp]
type namespaceInstanceReadResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                               `json:"period_ms"`
	Requests  int64                                                               `json:"requests"`
	Technique NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceInstanceReadResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct
// [NamespaceInstanceReadResponsePublicEndpointParamsRateLimit]
type namespaceInstanceReadResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceReadResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceReadResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                                `json:"disabled"`
	JSON     namespaceInstanceReadResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsSearchEndpointJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceReadResponsePublicEndpointParamsSearchEndpoint]
type namespaceInstanceReadResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy []NamespaceInstanceReadResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             namespaceInstanceReadResponseRetrievalOptionsJSON             `json:"-"`
}

// namespaceInstanceReadResponseRetrievalOptionsJSON contains the JSON metadata for
// the struct [NamespaceInstanceReadResponseRetrievalOptions]
type namespaceInstanceReadResponseRetrievalOptionsJSON struct {
	BoostBy          apijson.Field
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      namespaceInstanceReadResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceReadResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceReadResponseRetrievalOptionsBoostBy]
type namespaceInstanceReadResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
	Direction   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseRetrievalOptionsBoostBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseRetrievalOptionsBoostByJSON) RawJSON() string {
	return r.raw
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionExists    NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceReadResponseRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionExists, NamespaceInstanceReadResponseRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceReadResponseRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   []string                                            `json:"include_items"`
	Prefix         string                                              `json:"prefix"`
	R2Jurisdiction string                                              `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceReadResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           namespaceInstanceReadResponseSourceParamsJSON       `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsJSON contains the JSON metadata for the
// struct [NamespaceInstanceReadResponseSourceParams]
type namespaceInstanceReadResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptions `json:"discover_options"`
	ParseOptions    NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType `json:"parse_type"`
	JSON      namespaceInstanceReadResponseSourceParamsWebCrawlerJSON      `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceReadResponseSourceParamsWebCrawler]
type namespaceInstanceReadResponseSourceParamsWebCrawlerJSON struct {
	DiscoverOptions apijson.Field
	ParseOptions    apijson.Field
	ParseType       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth float64 `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks bool `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit float64 `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge float64 `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource `json:"source"`
	JSON   namespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON   `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptions]
type namespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	Limit                apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsJSON) RawJSON() string {
	return r.raw
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceReadResponseSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector []NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                            `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                `json:"use_browser_rendering"`
	JSON                namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions]
type namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	ContentSelector     apijson.Field
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector string                                                                             `json:"selector" api:"required"`
	JSON     namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON
// contains the JSON metadata for the struct
// [NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector]
type namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON struct {
	Path        apijson.Field
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON) RawJSON() string {
	return r.raw
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceReadResponseSyncInterval float64

const (
	NamespaceInstanceReadResponseSyncInterval900   NamespaceInstanceReadResponseSyncInterval = 900
	NamespaceInstanceReadResponseSyncInterval1800  NamespaceInstanceReadResponseSyncInterval = 1800
	NamespaceInstanceReadResponseSyncInterval3600  NamespaceInstanceReadResponseSyncInterval = 3600
	NamespaceInstanceReadResponseSyncInterval7200  NamespaceInstanceReadResponseSyncInterval = 7200
	NamespaceInstanceReadResponseSyncInterval14400 NamespaceInstanceReadResponseSyncInterval = 14400
	NamespaceInstanceReadResponseSyncInterval21600 NamespaceInstanceReadResponseSyncInterval = 21600
	NamespaceInstanceReadResponseSyncInterval43200 NamespaceInstanceReadResponseSyncInterval = 43200
	NamespaceInstanceReadResponseSyncInterval86400 NamespaceInstanceReadResponseSyncInterval = 86400
)

func (r NamespaceInstanceReadResponseSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseSyncInterval900, NamespaceInstanceReadResponseSyncInterval1800, NamespaceInstanceReadResponseSyncInterval3600, NamespaceInstanceReadResponseSyncInterval7200, NamespaceInstanceReadResponseSyncInterval14400, NamespaceInstanceReadResponseSyncInterval21600, NamespaceInstanceReadResponseSyncInterval43200, NamespaceInstanceReadResponseSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseType string

const (
	NamespaceInstanceReadResponseTypeR2         NamespaceInstanceReadResponseType = "r2"
	NamespaceInstanceReadResponseTypeWebCrawler NamespaceInstanceReadResponseType = "web-crawler"
)

func (r NamespaceInstanceReadResponseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseTypeR2, NamespaceInstanceReadResponseTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceSearchResponse struct {
	Chunks      []NamespaceInstanceSearchResponseChunk   `json:"chunks" api:"required"`
	QueryKind   NamespaceInstanceSearchResponseQueryKind `json:"query_kind" api:"required"`
	SearchQuery string                                   `json:"search_query"`
	JSON        namespaceInstanceSearchResponseJSON      `json:"-"`
}

// namespaceInstanceSearchResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceSearchResponse]
type namespaceInstanceSearchResponseJSON struct {
	Chunks      apijson.Field
	QueryKind   apijson.Field
	SearchQuery apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceSearchResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceSearchResponseChunk struct {
	ID             string                                              `json:"id" api:"required"`
	Score          float64                                             `json:"score" api:"required"`
	Text           string                                              `json:"text" api:"required"`
	Type           string                                              `json:"type" api:"required"`
	Item           NamespaceInstanceSearchResponseChunksItem           `json:"item"`
	ScoringDetails NamespaceInstanceSearchResponseChunksScoringDetails `json:"scoring_details"`
	JSON           namespaceInstanceSearchResponseChunkJSON            `json:"-"`
}

// namespaceInstanceSearchResponseChunkJSON contains the JSON metadata for the
// struct [NamespaceInstanceSearchResponseChunk]
type namespaceInstanceSearchResponseChunkJSON struct {
	ID             apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceSearchResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceSearchResponseChunkJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceSearchResponseChunksItem struct {
	Key       string                                        `json:"key" api:"required"`
	Metadata  map[string]interface{}                        `json:"metadata"`
	Timestamp float64                                       `json:"timestamp"`
	JSON      namespaceInstanceSearchResponseChunksItemJSON `json:"-"`
}

// namespaceInstanceSearchResponseChunksItemJSON contains the JSON metadata for the
// struct [NamespaceInstanceSearchResponseChunksItem]
type namespaceInstanceSearchResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceSearchResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceSearchResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceSearchResponseChunksScoringDetails struct {
	FusionMethod   NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethod `json:"fusion_method"`
	KeywordRank    float64                                                         `json:"keyword_rank"`
	KeywordScore   float64                                                         `json:"keyword_score"`
	RerankingScore float64                                                         `json:"reranking_score"`
	VectorRank     float64                                                         `json:"vector_rank"`
	VectorScore    float64                                                         `json:"vector_score"`
	JSON           namespaceInstanceSearchResponseChunksScoringDetailsJSON         `json:"-"`
}

// namespaceInstanceSearchResponseChunksScoringDetailsJSON contains the JSON
// metadata for the struct [NamespaceInstanceSearchResponseChunksScoringDetails]
type namespaceInstanceSearchResponseChunksScoringDetailsJSON struct {
	FusionMethod   apijson.Field
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceSearchResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceSearchResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethod string

const (
	NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethodRrf NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethod = "rrf"
	NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethodMax NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethod = "max"
)

func (r NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethodRrf, NamespaceInstanceSearchResponseChunksScoringDetailsFusionMethodMax:
		return true
	}
	return false
}

type NamespaceInstanceSearchResponseQueryKind string

const (
	NamespaceInstanceSearchResponseQueryKindText       NamespaceInstanceSearchResponseQueryKind = "text"
	NamespaceInstanceSearchResponseQueryKindImage      NamespaceInstanceSearchResponseQueryKind = "image"
	NamespaceInstanceSearchResponseQueryKindMultimodal NamespaceInstanceSearchResponseQueryKind = "multimodal"
)

func (r NamespaceInstanceSearchResponseQueryKind) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchResponseQueryKindText, NamespaceInstanceSearchResponseQueryKindImage, NamespaceInstanceSearchResponseQueryKindMultimodal:
		return true
	}
	return false
}

type NamespaceInstanceStatsResponse struct {
	Completed int64 `json:"completed"`
	// True when status counts are unavailable (e.g. legacy stats query exceeded D1
	// statement-size limit). Counts are omitted in this case.
	Degraded bool `json:"degraded"`
	// Engine-specific metadata. Present only for managed (v3) instances.
	Engine            NamespaceInstanceStatsResponseEngine `json:"engine"`
	Error             int64                                `json:"error"`
	FileEmbedErrors   map[string]interface{}               `json:"file_embed_errors"`
	IndexSourceErrors map[string]interface{}               `json:"index_source_errors"`
	LastActivity      time.Time                            `json:"last_activity" format:"date-time"`
	Outdated          int64                                `json:"outdated"`
	Queued            int64                                `json:"queued"`
	Running           int64                                `json:"running"`
	Skipped           int64                                `json:"skipped"`
	JSON              namespaceInstanceStatsResponseJSON   `json:"-"`
}

// namespaceInstanceStatsResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceStatsResponse]
type namespaceInstanceStatsResponseJSON struct {
	Completed         apijson.Field
	Degraded          apijson.Field
	Engine            apijson.Field
	Error             apijson.Field
	FileEmbedErrors   apijson.Field
	IndexSourceErrors apijson.Field
	LastActivity      apijson.Field
	Outdated          apijson.Field
	Queued            apijson.Field
	Running           apijson.Field
	Skipped           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *NamespaceInstanceStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceStatsResponseJSON) RawJSON() string {
	return r.raw
}

// Engine-specific metadata. Present only for managed (v3) instances.
type NamespaceInstanceStatsResponseEngine struct {
	// R2 bucket storage usage in bytes.
	R2 NamespaceInstanceStatsResponseEngineR2 `json:"r2"`
	// Vectorize index metadata (dimensions, vector count).
	Vectorize NamespaceInstanceStatsResponseEngineVectorize `json:"vectorize"`
	JSON      namespaceInstanceStatsResponseEngineJSON      `json:"-"`
}

// namespaceInstanceStatsResponseEngineJSON contains the JSON metadata for the
// struct [NamespaceInstanceStatsResponseEngine]
type namespaceInstanceStatsResponseEngineJSON struct {
	R2          apijson.Field
	Vectorize   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceStatsResponseEngine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceStatsResponseEngineJSON) RawJSON() string {
	return r.raw
}

// R2 bucket storage usage in bytes.
type NamespaceInstanceStatsResponseEngineR2 struct {
	MetadataSizeBytes int64                                      `json:"metadataSizeBytes" api:"required"`
	ObjectCount       int64                                      `json:"objectCount" api:"required"`
	PayloadSizeBytes  int64                                      `json:"payloadSizeBytes" api:"required"`
	JSON              namespaceInstanceStatsResponseEngineR2JSON `json:"-"`
}

// namespaceInstanceStatsResponseEngineR2JSON contains the JSON metadata for the
// struct [NamespaceInstanceStatsResponseEngineR2]
type namespaceInstanceStatsResponseEngineR2JSON struct {
	MetadataSizeBytes apijson.Field
	ObjectCount       apijson.Field
	PayloadSizeBytes  apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *NamespaceInstanceStatsResponseEngineR2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceStatsResponseEngineR2JSON) RawJSON() string {
	return r.raw
}

// Vectorize index metadata (dimensions, vector count).
type NamespaceInstanceStatsResponseEngineVectorize struct {
	Dimensions   int64                                             `json:"dimensions" api:"required"`
	VectorsCount int64                                             `json:"vectorsCount" api:"required"`
	JSON         namespaceInstanceStatsResponseEngineVectorizeJSON `json:"-"`
}

// namespaceInstanceStatsResponseEngineVectorizeJSON contains the JSON metadata for
// the struct [NamespaceInstanceStatsResponseEngineVectorize]
type namespaceInstanceStatsResponseEngineVectorizeJSON struct {
	Dimensions   apijson.Field
	VectorsCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceStatsResponseEngineVectorize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceStatsResponseEngineVectorizeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID          param.Field[string] `json:"id" api:"required"`
	AIGatewayID param.Field[string] `json:"ai_gateway_id"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  param.Field[string]                                   `json:"ai_search_model"`
	Cache          param.Field[bool]                                     `json:"cache"`
	CacheThreshold param.Field[NamespaceInstanceNewParamsCacheThreshold] `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       param.Field[NamespaceInstanceNewParamsCacheTTL]         `json:"cache_ttl"`
	Chunk          param.Field[bool]                                       `json:"chunk"`
	ChunkOverlap   param.Field[int64]                                      `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                                      `json:"chunk_size"`
	CustomMetadata param.Field[[]NamespaceInstanceNewParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[string]                                     `json:"embedding_model"`
	FusionMethod   param.Field[NamespaceInstanceNewParamsFusionMethod]     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	HybridSearchEnabled param.Field[bool] `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          param.Field[NamespaceInstanceNewParamsIndexMethod]          `json:"index_method"`
	IndexingOptions      param.Field[NamespaceInstanceNewParamsIndexingOptions]      `json:"indexing_options"`
	MaxNumResults        param.Field[int64]                                          `json:"max_num_results"`
	Metadata             param.Field[NamespaceInstanceNewParamsMetadata]             `json:"metadata"`
	PublicEndpointParams param.Field[NamespaceInstanceNewParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking            param.Field[bool]                                           `json:"reranking"`
	RerankingModel       param.Field[string]                                         `json:"reranking_model"`
	RetrievalOptions     param.Field[NamespaceInstanceNewParamsRetrievalOptions]     `json:"retrieval_options"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel   param.Field[string]                                 `json:"rewrite_model"`
	RewriteQuery   param.Field[bool]                                   `json:"rewrite_query"`
	ScoreThreshold param.Field[float64]                                `json:"score_threshold"`
	Source         param.Field[string]                                 `json:"source"`
	SourceParams   param.Field[NamespaceInstanceNewParamsSourceParams] `json:"source_params"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval param.Field[NamespaceInstanceNewParamsSyncInterval] `json:"sync_interval"`
	TokenID      param.Field[string]                                 `json:"token_id" format:"uuid"`
	Type         param.Field[NamespaceInstanceNewParamsType]         `json:"type"`
}

func (r NamespaceInstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsCacheThreshold string

const (
	NamespaceInstanceNewParamsCacheThresholdSuperStrictMatch NamespaceInstanceNewParamsCacheThreshold = "super_strict_match"
	NamespaceInstanceNewParamsCacheThresholdCloseEnough      NamespaceInstanceNewParamsCacheThreshold = "close_enough"
	NamespaceInstanceNewParamsCacheThresholdFlexibleFriend   NamespaceInstanceNewParamsCacheThreshold = "flexible_friend"
	NamespaceInstanceNewParamsCacheThresholdAnythingGoes     NamespaceInstanceNewParamsCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceNewParamsCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsCacheThresholdSuperStrictMatch, NamespaceInstanceNewParamsCacheThresholdCloseEnough, NamespaceInstanceNewParamsCacheThresholdFlexibleFriend, NamespaceInstanceNewParamsCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceNewParamsCacheTTL float64

const (
	NamespaceInstanceNewParamsCacheTTL600    NamespaceInstanceNewParamsCacheTTL = 600
	NamespaceInstanceNewParamsCacheTTL1800   NamespaceInstanceNewParamsCacheTTL = 1800
	NamespaceInstanceNewParamsCacheTTL3600   NamespaceInstanceNewParamsCacheTTL = 3600
	NamespaceInstanceNewParamsCacheTTL7200   NamespaceInstanceNewParamsCacheTTL = 7200
	NamespaceInstanceNewParamsCacheTTL21600  NamespaceInstanceNewParamsCacheTTL = 21600
	NamespaceInstanceNewParamsCacheTTL43200  NamespaceInstanceNewParamsCacheTTL = 43200
	NamespaceInstanceNewParamsCacheTTL86400  NamespaceInstanceNewParamsCacheTTL = 86400
	NamespaceInstanceNewParamsCacheTTL172800 NamespaceInstanceNewParamsCacheTTL = 172800
	NamespaceInstanceNewParamsCacheTTL259200 NamespaceInstanceNewParamsCacheTTL = 259200
	NamespaceInstanceNewParamsCacheTTL518400 NamespaceInstanceNewParamsCacheTTL = 518400
)

func (r NamespaceInstanceNewParamsCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsCacheTTL600, NamespaceInstanceNewParamsCacheTTL1800, NamespaceInstanceNewParamsCacheTTL3600, NamespaceInstanceNewParamsCacheTTL7200, NamespaceInstanceNewParamsCacheTTL21600, NamespaceInstanceNewParamsCacheTTL43200, NamespaceInstanceNewParamsCacheTTL86400, NamespaceInstanceNewParamsCacheTTL172800, NamespaceInstanceNewParamsCacheTTL259200, NamespaceInstanceNewParamsCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsCustomMetadata struct {
	DataType  param.Field[NamespaceInstanceNewParamsCustomMetadataDataType] `json:"data_type" api:"required"`
	FieldName param.Field[string]                                           `json:"field_name" api:"required"`
}

func (r NamespaceInstanceNewParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsCustomMetadataDataType string

const (
	NamespaceInstanceNewParamsCustomMetadataDataTypeText     NamespaceInstanceNewParamsCustomMetadataDataType = "text"
	NamespaceInstanceNewParamsCustomMetadataDataTypeNumber   NamespaceInstanceNewParamsCustomMetadataDataType = "number"
	NamespaceInstanceNewParamsCustomMetadataDataTypeBoolean  NamespaceInstanceNewParamsCustomMetadataDataType = "boolean"
	NamespaceInstanceNewParamsCustomMetadataDataTypeDatetime NamespaceInstanceNewParamsCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceNewParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsCustomMetadataDataTypeText, NamespaceInstanceNewParamsCustomMetadataDataTypeNumber, NamespaceInstanceNewParamsCustomMetadataDataTypeBoolean, NamespaceInstanceNewParamsCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsFusionMethod string

const (
	NamespaceInstanceNewParamsFusionMethodMax NamespaceInstanceNewParamsFusionMethod = "max"
	NamespaceInstanceNewParamsFusionMethodRrf NamespaceInstanceNewParamsFusionMethod = "rrf"
)

func (r NamespaceInstanceNewParamsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsFusionMethodMax, NamespaceInstanceNewParamsFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceNewParamsIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword param.Field[bool] `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector param.Field[bool] `json:"vector" api:"required"`
}

func (r NamespaceInstanceNewParamsIndexMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer param.Field[NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizer] `json:"keyword_tokenizer"`
}

func (r NamespaceInstanceNewParamsIndexingOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceNewParamsIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsMetadata struct {
	CreatedFromAISearchWizard param.Field[bool]   `json:"created_from_aisearch_wizard"`
	WorkerDomain              param.Field[string] `json:"worker_domain"`
}

func (r NamespaceInstanceNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                              `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[NamespaceInstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled param.Field[bool]                                                         `json:"default_domain_enabled"`
	Enabled              param.Field[bool]                                                         `json:"enabled"`
	Mcp                  param.Field[NamespaceInstanceNewParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit            param.Field[NamespaceInstanceNewParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint       param.Field[NamespaceInstanceNewParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
}

func (r NamespaceInstanceNewParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceNewParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                            `json:"period_ms"`
	Requests  param.Field[int64]                                                            `json:"requests"`
	Technique param.Field[NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r NamespaceInstanceNewParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceNewParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceNewParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy param.Field[[]NamespaceInstanceNewParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r NamespaceInstanceNewParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceNewParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionExists    NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionExists, NamespaceInstanceNewParamsRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   param.Field[[]string]                                         `json:"include_items"`
	Prefix         param.Field[string]                                           `json:"prefix"`
	R2Jurisdiction param.Field[string]                                           `json:"r2_jurisdiction"`
	WebCrawler     param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r NamespaceInstanceNewParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptions] `json:"discover_options"`
	ParseOptions    param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions]    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType] `json:"parse_type"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth param.Field[float64] `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks param.Field[bool] `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit param.Field[float64] `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge param.Field[float64] `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource] `json:"source"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceNewParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector param.Field[[]NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders param.Field[map[string]string] `json:"include_headers"`
	IncludeImages  param.Field[bool]              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    param.Field[[]string] `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering param.Field[bool]     `json:"use_browser_rendering"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path param.Field[string] `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceNewParamsSyncInterval float64

const (
	NamespaceInstanceNewParamsSyncInterval900   NamespaceInstanceNewParamsSyncInterval = 900
	NamespaceInstanceNewParamsSyncInterval1800  NamespaceInstanceNewParamsSyncInterval = 1800
	NamespaceInstanceNewParamsSyncInterval3600  NamespaceInstanceNewParamsSyncInterval = 3600
	NamespaceInstanceNewParamsSyncInterval7200  NamespaceInstanceNewParamsSyncInterval = 7200
	NamespaceInstanceNewParamsSyncInterval14400 NamespaceInstanceNewParamsSyncInterval = 14400
	NamespaceInstanceNewParamsSyncInterval21600 NamespaceInstanceNewParamsSyncInterval = 21600
	NamespaceInstanceNewParamsSyncInterval43200 NamespaceInstanceNewParamsSyncInterval = 43200
	NamespaceInstanceNewParamsSyncInterval86400 NamespaceInstanceNewParamsSyncInterval = 86400
)

func (r NamespaceInstanceNewParamsSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsSyncInterval900, NamespaceInstanceNewParamsSyncInterval1800, NamespaceInstanceNewParamsSyncInterval3600, NamespaceInstanceNewParamsSyncInterval7200, NamespaceInstanceNewParamsSyncInterval14400, NamespaceInstanceNewParamsSyncInterval21600, NamespaceInstanceNewParamsSyncInterval43200, NamespaceInstanceNewParamsSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsType string

const (
	NamespaceInstanceNewParamsTypeR2         NamespaceInstanceNewParamsType = "r2"
	NamespaceInstanceNewParamsTypeWebCrawler NamespaceInstanceNewParamsType = "web-crawler"
)

func (r NamespaceInstanceNewParamsType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsTypeR2, NamespaceInstanceNewParamsTypeWebCrawler:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseEnvelope struct {
	Result  NamespaceInstanceNewResponse             `json:"result" api:"required"`
	Success bool                                     `json:"success" api:"required"`
	JSON    namespaceInstanceNewResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceNewResponseEnvelope]
type namespaceInstanceNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	AIGatewayID param.Field[string] `json:"ai_gateway_id"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	AISearchModel  param.Field[string]                                      `json:"ai_search_model"`
	Cache          param.Field[bool]                                        `json:"cache"`
	CacheThreshold param.Field[NamespaceInstanceUpdateParamsCacheThreshold] `json:"cache_threshold"`
	// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
	// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
	// (72h), 518400 (6d).
	CacheTTL       param.Field[NamespaceInstanceUpdateParamsCacheTTL]         `json:"cache_ttl"`
	Chunk          param.Field[bool]                                          `json:"chunk"`
	ChunkOverlap   param.Field[int64]                                         `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                                         `json:"chunk_size"`
	CustomMetadata param.Field[[]NamespaceInstanceUpdateParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[string]                                        `json:"embedding_model"`
	FusionMethod   param.Field[NamespaceInstanceUpdateParamsFusionMethod]     `json:"fusion_method"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          param.Field[NamespaceInstanceUpdateParamsIndexMethod]          `json:"index_method"`
	IndexingOptions      param.Field[NamespaceInstanceUpdateParamsIndexingOptions]      `json:"indexing_options"`
	MaxNumResults        param.Field[int64]                                             `json:"max_num_results"`
	Metadata             param.Field[NamespaceInstanceUpdateParamsMetadata]             `json:"metadata"`
	Paused               param.Field[bool]                                              `json:"paused"`
	PublicEndpointParams param.Field[NamespaceInstanceUpdateParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking            param.Field[bool]                                              `json:"reranking"`
	RerankingModel       param.Field[string]                                            `json:"reranking_model"`
	RetrievalOptions     param.Field[NamespaceInstanceUpdateParamsRetrievalOptions]     `json:"retrieval_options"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	RewriteModel       param.Field[string]                                    `json:"rewrite_model"`
	RewriteQuery       param.Field[bool]                                      `json:"rewrite_query"`
	ScoreThreshold     param.Field[float64]                                   `json:"score_threshold"`
	Source             param.Field[string]                                    `json:"source"`
	SourceParams       param.Field[NamespaceInstanceUpdateParamsSourceParams] `json:"source_params"`
	Summarization      param.Field[bool]                                      `json:"summarization"`
	SummarizationModel param.Field[string]                                    `json:"summarization_model"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval                   param.Field[NamespaceInstanceUpdateParamsSyncInterval] `json:"sync_interval"`
	SystemPromptAISearch           param.Field[string]                                    `json:"system_prompt_ai_search"`
	SystemPromptIndexSummarization param.Field[string]                                    `json:"system_prompt_index_summarization"`
	SystemPromptRewriteQuery       param.Field[string]                                    `json:"system_prompt_rewrite_query"`
	TokenID                        param.Field[string]                                    `json:"token_id" format:"uuid"`
}

func (r NamespaceInstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsCacheThreshold string

const (
	NamespaceInstanceUpdateParamsCacheThresholdSuperStrictMatch NamespaceInstanceUpdateParamsCacheThreshold = "super_strict_match"
	NamespaceInstanceUpdateParamsCacheThresholdCloseEnough      NamespaceInstanceUpdateParamsCacheThreshold = "close_enough"
	NamespaceInstanceUpdateParamsCacheThresholdFlexibleFriend   NamespaceInstanceUpdateParamsCacheThreshold = "flexible_friend"
	NamespaceInstanceUpdateParamsCacheThresholdAnythingGoes     NamespaceInstanceUpdateParamsCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceUpdateParamsCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsCacheThresholdSuperStrictMatch, NamespaceInstanceUpdateParamsCacheThresholdCloseEnough, NamespaceInstanceUpdateParamsCacheThresholdFlexibleFriend, NamespaceInstanceUpdateParamsCacheThresholdAnythingGoes:
		return true
	}
	return false
}

// Cache entry TTL in seconds. Allowed values: 600 (10min), 1800 (30min), 3600
// (1h), 7200 (2h), 21600 (6h), 43200 (12h), 86400 (24h), 172800 (48h), 259200
// (72h), 518400 (6d).
type NamespaceInstanceUpdateParamsCacheTTL float64

const (
	NamespaceInstanceUpdateParamsCacheTTL600    NamespaceInstanceUpdateParamsCacheTTL = 600
	NamespaceInstanceUpdateParamsCacheTTL1800   NamespaceInstanceUpdateParamsCacheTTL = 1800
	NamespaceInstanceUpdateParamsCacheTTL3600   NamespaceInstanceUpdateParamsCacheTTL = 3600
	NamespaceInstanceUpdateParamsCacheTTL7200   NamespaceInstanceUpdateParamsCacheTTL = 7200
	NamespaceInstanceUpdateParamsCacheTTL21600  NamespaceInstanceUpdateParamsCacheTTL = 21600
	NamespaceInstanceUpdateParamsCacheTTL43200  NamespaceInstanceUpdateParamsCacheTTL = 43200
	NamespaceInstanceUpdateParamsCacheTTL86400  NamespaceInstanceUpdateParamsCacheTTL = 86400
	NamespaceInstanceUpdateParamsCacheTTL172800 NamespaceInstanceUpdateParamsCacheTTL = 172800
	NamespaceInstanceUpdateParamsCacheTTL259200 NamespaceInstanceUpdateParamsCacheTTL = 259200
	NamespaceInstanceUpdateParamsCacheTTL518400 NamespaceInstanceUpdateParamsCacheTTL = 518400
)

func (r NamespaceInstanceUpdateParamsCacheTTL) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsCacheTTL600, NamespaceInstanceUpdateParamsCacheTTL1800, NamespaceInstanceUpdateParamsCacheTTL3600, NamespaceInstanceUpdateParamsCacheTTL7200, NamespaceInstanceUpdateParamsCacheTTL21600, NamespaceInstanceUpdateParamsCacheTTL43200, NamespaceInstanceUpdateParamsCacheTTL86400, NamespaceInstanceUpdateParamsCacheTTL172800, NamespaceInstanceUpdateParamsCacheTTL259200, NamespaceInstanceUpdateParamsCacheTTL518400:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsCustomMetadata struct {
	DataType  param.Field[NamespaceInstanceUpdateParamsCustomMetadataDataType] `json:"data_type" api:"required"`
	FieldName param.Field[string]                                              `json:"field_name" api:"required"`
}

func (r NamespaceInstanceUpdateParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsCustomMetadataDataType string

const (
	NamespaceInstanceUpdateParamsCustomMetadataDataTypeText     NamespaceInstanceUpdateParamsCustomMetadataDataType = "text"
	NamespaceInstanceUpdateParamsCustomMetadataDataTypeNumber   NamespaceInstanceUpdateParamsCustomMetadataDataType = "number"
	NamespaceInstanceUpdateParamsCustomMetadataDataTypeBoolean  NamespaceInstanceUpdateParamsCustomMetadataDataType = "boolean"
	NamespaceInstanceUpdateParamsCustomMetadataDataTypeDatetime NamespaceInstanceUpdateParamsCustomMetadataDataType = "datetime"
)

func (r NamespaceInstanceUpdateParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsCustomMetadataDataTypeText, NamespaceInstanceUpdateParamsCustomMetadataDataTypeNumber, NamespaceInstanceUpdateParamsCustomMetadataDataTypeBoolean, NamespaceInstanceUpdateParamsCustomMetadataDataTypeDatetime:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsFusionMethod string

const (
	NamespaceInstanceUpdateParamsFusionMethodMax NamespaceInstanceUpdateParamsFusionMethod = "max"
	NamespaceInstanceUpdateParamsFusionMethodRrf NamespaceInstanceUpdateParamsFusionMethod = "rrf"
)

func (r NamespaceInstanceUpdateParamsFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsFusionMethodMax, NamespaceInstanceUpdateParamsFusionMethodRrf:
		return true
	}
	return false
}

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceUpdateParamsIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword param.Field[bool] `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector param.Field[bool] `json:"vector" api:"required"`
}

func (r NamespaceInstanceUpdateParamsIndexMethod) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsIndexingOptions struct {
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer param.Field[NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizer] `json:"keyword_tokenizer"`
}

func (r NamespaceInstanceUpdateParamsIndexingOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
type NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizer string

const (
	NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter  NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizer = "porter"
	NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizerTrigram NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizer = "trigram"
)

func (r NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizer) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizerPorter, NamespaceInstanceUpdateParamsIndexingOptionsKeywordTokenizerTrigram:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsMetadata struct {
	CreatedFromAISearchWizard param.Field[bool]   `json:"created_from_aisearch_wizard"`
	WorkerDomain              param.Field[string] `json:"worker_domain"`
}

func (r NamespaceInstanceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                                 `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
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
	DefaultDomainEnabled param.Field[bool]                                                            `json:"default_domain_enabled"`
	Enabled              param.Field[bool]                                                            `json:"enabled"`
	Mcp                  param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsMcp]            `json:"mcp"`
	RateLimit            param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimit]      `json:"rate_limit"`
	SearchEndpoint       param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsSearchEndpoint] `json:"search_endpoint"`
}

func (r NamespaceInstanceUpdateParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceUpdateParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                               `json:"period_ms"`
	Requests  param.Field[int64]                                                               `json:"requests"`
	Technique param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechnique string

const (
	NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed   NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed, NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r NamespaceInstanceUpdateParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for
	// numeric/datetime fields and 'exists' for text/boolean fields. Fields must match
	// 'timestamp' or a defined custom_metadata field.
	BoostBy param.Field[[]NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted on an
	// update, the existing stored value is preserved; when never set, search falls
	// back to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r NamespaceInstanceUpdateParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection string

const (
	NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionAsc       NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection = "asc"
	NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc      NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection = "desc"
	NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionExists    NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection = "exists"
	NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionNotExists NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection = "not_exists"
)

func (r NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionAsc, NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionDesc, NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionExists, NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirectionNotExists:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted on an
// update, the existing stored value is preserved; when never set, search falls
// back to 'and'.
type NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode string

const (
	NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "and"
	NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchModeOr  NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "or"
)

func (r NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchModeAnd, NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced). Most accounts are limited to 10
	// rules; contact support to raise it.
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post). Most accounts are limited to 10 rules; contact
	// support to raise it.
	IncludeItems   param.Field[[]string]                                            `json:"include_items"`
	Prefix         param.Field[string]                                              `json:"prefix"`
	R2Jurisdiction param.Field[string]                                              `json:"r2_jurisdiction"`
	WebCrawler     param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r NamespaceInstanceUpdateParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawler struct {
	// Options for parse_type 'discover', where Browser Run discovers URLs by link
	// following and sitemaps. Ignored for 'sitemap'.
	DiscoverOptions param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions] `json:"discover_options"`
	ParseOptions    param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions]    `json:"parse_options"`
	// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
	// recursively and requires the source to be a Verified zone on this account.
	ParseType param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType] `json:"parse_type"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Options for parse_type 'discover', where Browser Run discovers URLs by link
// following and sitemaps. Ignored for 'sitemap'.
type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions struct {
	// Maximum link-follow depth from the seed URL.
	Depth param.Field[float64] `json:"depth"`
	// Follow links that point outside the source domain. Must stay `false` — discover
	// crawls are restricted to the zone you own.
	IncludeExternalLinks param.Field[bool] `json:"include_external_links"`
	// Follow links to subdomains of the source host.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Maximum number of pages to crawl (1-100000).
	Limit param.Field[float64] `json:"limit"`
	// Maximum content age in seconds to accept (0–604800).
	MaxAge param.Field[float64] `json:"max_age"`
	// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
	// follows page links only, 'all' does both.
	Source param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource] `json:"source"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Where the crawler looks for URLs: 'sitemaps' reads sitemap XML only, 'links'
// follows page links only, 'all' does both.
type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource string

const (
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll      NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "all"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "sitemaps"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks    NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource = "links"
)

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceAll, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceSitemaps, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerDiscoverOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed. Omit
	// the field to disable content selection — empty arrays are rejected.
	ContentSelector param.Field[[]NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
	// Up to 5 custom HTTP headers sent with each crawl request. Names must be RFC-7230
	// token characters (no spaces, colons, or control characters); values must be
	// HTAB + printable ASCII (no CR/LF).
	IncludeHeaders param.Field[map[string]string] `json:"include_headers"`
	IncludeImages  param.Field[bool]              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    param.Field[[]string] `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering param.Field[bool]     `json:"use_browser_rendering"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector struct {
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path param.Field[string] `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Must not
	// contain disallowed characters (;, `, $, {, }, \). Must target a single element;
	// if multiple elements match, the selector is ignored and the full page is used.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How URLs are discovered. 'sitemap' reads XML sitemaps; 'discover' follows links
// recursively and requires the source to be a Verified zone on this account.
type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap  NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeDiscover NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType = "discover"
)

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeDiscover:
		return true
	}
	return false
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
type NamespaceInstanceUpdateParamsSyncInterval float64

const (
	NamespaceInstanceUpdateParamsSyncInterval900   NamespaceInstanceUpdateParamsSyncInterval = 900
	NamespaceInstanceUpdateParamsSyncInterval1800  NamespaceInstanceUpdateParamsSyncInterval = 1800
	NamespaceInstanceUpdateParamsSyncInterval3600  NamespaceInstanceUpdateParamsSyncInterval = 3600
	NamespaceInstanceUpdateParamsSyncInterval7200  NamespaceInstanceUpdateParamsSyncInterval = 7200
	NamespaceInstanceUpdateParamsSyncInterval14400 NamespaceInstanceUpdateParamsSyncInterval = 14400
	NamespaceInstanceUpdateParamsSyncInterval21600 NamespaceInstanceUpdateParamsSyncInterval = 21600
	NamespaceInstanceUpdateParamsSyncInterval43200 NamespaceInstanceUpdateParamsSyncInterval = 43200
	NamespaceInstanceUpdateParamsSyncInterval86400 NamespaceInstanceUpdateParamsSyncInterval = 86400
)

func (r NamespaceInstanceUpdateParamsSyncInterval) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSyncInterval900, NamespaceInstanceUpdateParamsSyncInterval1800, NamespaceInstanceUpdateParamsSyncInterval3600, NamespaceInstanceUpdateParamsSyncInterval7200, NamespaceInstanceUpdateParamsSyncInterval14400, NamespaceInstanceUpdateParamsSyncInterval21600, NamespaceInstanceUpdateParamsSyncInterval43200, NamespaceInstanceUpdateParamsSyncInterval86400:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseEnvelope struct {
	Result  NamespaceInstanceUpdateResponse             `json:"result" api:"required"`
	Success bool                                        `json:"success" api:"required"`
	JSON    namespaceInstanceUpdateResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceUpdateResponseEnvelope]
type namespaceInstanceUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by namespace.
	Namespace param.Field[string] `query:"namespace"`
	// Field to order results by.
	OrderBy param.Field[NamespaceInstanceListParamsOrderBy] `query:"order_by"`
	// Order direction.
	OrderByDirection param.Field[NamespaceInstanceListParamsOrderByDirection] `query:"order_by_direction"`
	// Page number (1-indexed).
	Page param.Field[int64] `query:"page"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter instances whose id contains this string (case-insensitive).
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [NamespaceInstanceListParams]'s query parameters as
// `url.Values`.
func (r NamespaceInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order results by.
type NamespaceInstanceListParamsOrderBy string

const (
	NamespaceInstanceListParamsOrderByCreatedAt NamespaceInstanceListParamsOrderBy = "created_at"
)

func (r NamespaceInstanceListParamsOrderBy) IsKnown() bool {
	switch r {
	case NamespaceInstanceListParamsOrderByCreatedAt:
		return true
	}
	return false
}

// Order direction.
type NamespaceInstanceListParamsOrderByDirection string

const (
	NamespaceInstanceListParamsOrderByDirectionAsc  NamespaceInstanceListParamsOrderByDirection = "asc"
	NamespaceInstanceListParamsOrderByDirectionDesc NamespaceInstanceListParamsOrderByDirection = "desc"
)

func (r NamespaceInstanceListParamsOrderByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceListParamsOrderByDirectionAsc, NamespaceInstanceListParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type NamespaceInstanceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceDeleteResponseEnvelope struct {
	Result  NamespaceInstanceDeleteResponse             `json:"result" api:"required"`
	Success bool                                        `json:"success" api:"required"`
	JSON    namespaceInstanceDeleteResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceDeleteResponseEnvelope]
type namespaceInstanceDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceChatCompletionsParams struct {
	AccountID       param.Field[string]                                                `path:"account_id" api:"required"`
	Messages        param.Field[[]NamespaceInstanceChatCompletionsParamsMessage]       `json:"messages" api:"required"`
	AISearchOptions param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptions] `json:"ai_search_options"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	Model  param.Field[string] `json:"model"`
	Stream param.Field[bool]   `json:"stream"`
}

func (r NamespaceInstanceChatCompletionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsMessage struct {
	Content     param.Field[NamespaceInstanceChatCompletionsParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[NamespaceInstanceChatCompletionsParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                                  `json:"-,extras"`
}

func (r NamespaceInstanceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.NamespaceInstanceChatCompletionsParamsMessagesContentArray],
// [shared.UnionString].
type NamespaceInstanceChatCompletionsParamsMessagesContentUnion interface {
	ImplementsNamespaceInstanceChatCompletionsParamsMessagesContentUnion()
}

type NamespaceInstanceChatCompletionsParamsMessagesContentArray []NamespaceInstanceChatCompletionsParamsMessagesContentArrayItemUnion

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArray) ImplementsNamespaceInstanceChatCompletionsParamsMessagesContentUnion() {
}

type NamespaceInstanceChatCompletionsParamsMessagesContentArrayItem struct {
	Type     param.Field[NamespaceInstanceChatCompletionsParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                                    `json:"file"`
	ImageURL param.Field[interface{}]                                                    `json:"image_url"`
	Text     param.Field[string]                                                         `json:"text"`
}

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayItem) implementsNamespaceInstanceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [ai_search.NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject],
// [ai_search.NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject],
// [NamespaceInstanceChatCompletionsParamsMessagesContentArrayItem].
type NamespaceInstanceChatCompletionsParamsMessagesContentArrayItemUnion interface {
	implementsNamespaceInstanceChatCompletionsParamsMessagesContentArrayItemUnion()
}

type NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                                               `json:"text" api:"required"`
	Type param.Field[NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayObject) implementsNamespaceInstanceChatCompletionsParamsMessagesContentArrayItemUnion() {
}

type NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectType string

const (
	NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectTypeText NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectType = "text"
)

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsMessagesContentArrayType string

const (
	NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeText     NamespaceInstanceChatCompletionsParamsMessagesContentArrayType = "text"
	NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeImageURL NamespaceInstanceChatCompletionsParamsMessagesContentArrayType = "image_url"
	NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeFile     NamespaceInstanceChatCompletionsParamsMessagesContentArrayType = "file"
)

func (r NamespaceInstanceChatCompletionsParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeText, NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeImageURL, NamespaceInstanceChatCompletionsParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsMessagesRole string

const (
	NamespaceInstanceChatCompletionsParamsMessagesRoleSystem    NamespaceInstanceChatCompletionsParamsMessagesRole = "system"
	NamespaceInstanceChatCompletionsParamsMessagesRoleDeveloper NamespaceInstanceChatCompletionsParamsMessagesRole = "developer"
	NamespaceInstanceChatCompletionsParamsMessagesRoleUser      NamespaceInstanceChatCompletionsParamsMessagesRole = "user"
	NamespaceInstanceChatCompletionsParamsMessagesRoleAssistant NamespaceInstanceChatCompletionsParamsMessagesRole = "assistant"
	NamespaceInstanceChatCompletionsParamsMessagesRoleTool      NamespaceInstanceChatCompletionsParamsMessagesRole = "tool"
)

func (r NamespaceInstanceChatCompletionsParamsMessagesRole) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsMessagesRoleSystem, NamespaceInstanceChatCompletionsParamsMessagesRoleDeveloper, NamespaceInstanceChatCompletionsParamsMessagesRoleUser, NamespaceInstanceChatCompletionsParamsMessagesRoleAssistant, NamespaceInstanceChatCompletionsParamsMessagesRoleTool:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsAISearchOptions struct {
	Cache        param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                                     `json:"enabled"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough      NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes     NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdCloseEnough, NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, NamespaceInstanceChatCompletionsParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewrite struct {
	Enabled param.Field[bool] `json:"enabled"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                                      `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                                     `json:"filters"`
	FusionMethod     param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                                        `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                                          `json:"max_num_results"`
	RetrievalType    param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                                           `json:"return_on_failure"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc       NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc      NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists    NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionAsc, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionDesc, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionExists, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "max"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr  NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeAnd, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector  NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "vector"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid  NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword, NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type NamespaceInstanceReadParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceReadResponseEnvelope struct {
	Result  NamespaceInstanceReadResponse             `json:"result" api:"required"`
	Success bool                                      `json:"success" api:"required"`
	JSON    namespaceInstanceReadResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceReadResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceReadResponseEnvelope]
type namespaceInstanceReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceSearchParams struct {
	AccountID       param.Field[string]                                       `path:"account_id" api:"required"`
	AISearchOptions param.Field[NamespaceInstanceSearchParamsAISearchOptions] `json:"ai_search_options"`
	// OpenAI-compatible message array. For multimodal queries, set the last user
	// message's `content` to an array of typed parts:
	// `[{type:'text', text:'…'}, {type:'image_url', image_url:{url:'…'}}]`. Image
	// inputs require the RAG's embedding_model to declare 'image' in
	// supported_modalities.
	Messages param.Field[[]NamespaceInstanceSearchParamsMessage] `json:"messages"`
	// A simple text query string. Alternative to 'messages' — provide either this or
	// 'messages', not both.
	Query param.Field[string] `json:"query"`
}

func (r NamespaceInstanceSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptions struct {
	Cache        param.Field[NamespaceInstanceSearchParamsAISearchOptionsCache]        `json:"cache"`
	QueryRewrite param.Field[NamespaceInstanceSearchParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[NamespaceInstanceSearchParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r NamespaceInstanceSearchParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsCache struct {
	CacheThreshold param.Field[NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold] `json:"cache_threshold"`
	Enabled        param.Field[bool]                                                            `json:"enabled"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold string

const (
	NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold = "super_strict_match"
	NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough      NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold = "close_enough"
	NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend   NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold = "flexible_friend"
	NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes     NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold = "anything_goes"
)

func (r NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThreshold) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdSuperStrictMatch, NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdCloseEnough, NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdFlexibleFriend, NamespaceInstanceSearchParamsAISearchOptionsCacheCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsAISearchOptionsQueryRewrite struct {
	Enabled param.Field[bool] `json:"enabled"`
	// A Workers AI model ID or an AI Gateway model ID compatible with the OpenAI Chat
	// Completions API. An empty string uses the configured or default model.
	Model         param.Field[string] `json:"model"`
	RewritePrompt param.Field[string] `json:"rewrite_prompt"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]    `json:"enabled"`
	MatchThreshold param.Field[float64] `json:"match_threshold"`
	Model          param.Field[string]  `json:"model"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsRetrieval struct {
	// Metadata fields to boost search results by. Overrides the instance-level
	// boost_by config. Direction defaults to 'asc' for numeric/datetime fields,
	// 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy          param.Field[[]NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostBy]    `json:"boost_by"`
	ContextExpansion param.Field[int64]                                                             `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                            `json:"filters"`
	FusionMethod     param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. When omitted, falls back
	// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                               `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                                 `json:"max_num_results"`
	RetrievalType    param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                                  `json:"return_on_failure"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support all four directions
	// (asc, desc, exists, not_exists); text/boolean fields only support
	// exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional — defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
type NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection string

const (
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc       NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "asc"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc      NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "desc"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists    NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "exists"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection = "not_exists"
)

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionAsc, NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionDesc, NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionExists, NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirectionNotExists:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethod string

const (
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethod = "max"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethodRrf NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax, NamespaceInstanceSearchParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. When omitted, falls back
// to the instance-level retrieval_options.keyword_match_mode, then to 'and'.
type NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "and"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr  NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "or"
)

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeAnd, NamespaceInstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeOr:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType string

const (
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector  NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "vector"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid  NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector, NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword, NamespaceInstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsMessage struct {
	Content     param.Field[NamespaceInstanceSearchParamsMessagesContentUnion] `json:"content" api:"required"`
	Role        param.Field[NamespaceInstanceSearchParamsMessagesRole]         `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                         `json:"-,extras"`
}

func (r NamespaceInstanceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString],
// [ai_search.NamespaceInstanceSearchParamsMessagesContentArray],
// [shared.UnionString].
type NamespaceInstanceSearchParamsMessagesContentUnion interface {
	ImplementsNamespaceInstanceSearchParamsMessagesContentUnion()
}

type NamespaceInstanceSearchParamsMessagesContentArray []NamespaceInstanceSearchParamsMessagesContentArrayItemUnion

func (r NamespaceInstanceSearchParamsMessagesContentArray) ImplementsNamespaceInstanceSearchParamsMessagesContentUnion() {
}

type NamespaceInstanceSearchParamsMessagesContentArrayItem struct {
	Type     param.Field[NamespaceInstanceSearchParamsMessagesContentArrayType] `json:"type" api:"required"`
	File     param.Field[interface{}]                                           `json:"file"`
	ImageURL param.Field[interface{}]                                           `json:"image_url"`
	Text     param.Field[string]                                                `json:"text"`
}

func (r NamespaceInstanceSearchParamsMessagesContentArrayItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceInstanceSearchParamsMessagesContentArrayItem) implementsNamespaceInstanceSearchParamsMessagesContentArrayItemUnion() {
}

// Satisfied by
// [ai_search.NamespaceInstanceSearchParamsMessagesContentArrayObject],
// [ai_search.NamespaceInstanceSearchParamsMessagesContentArrayObject],
// [ai_search.NamespaceInstanceSearchParamsMessagesContentArrayObject],
// [NamespaceInstanceSearchParamsMessagesContentArrayItem].
type NamespaceInstanceSearchParamsMessagesContentArrayItemUnion interface {
	implementsNamespaceInstanceSearchParamsMessagesContentArrayItemUnion()
}

type NamespaceInstanceSearchParamsMessagesContentArrayObject struct {
	Text param.Field[string]                                                      `json:"text" api:"required"`
	Type param.Field[NamespaceInstanceSearchParamsMessagesContentArrayObjectType] `json:"type" api:"required"`
}

func (r NamespaceInstanceSearchParamsMessagesContentArrayObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r NamespaceInstanceSearchParamsMessagesContentArrayObject) implementsNamespaceInstanceSearchParamsMessagesContentArrayItemUnion() {
}

type NamespaceInstanceSearchParamsMessagesContentArrayObjectType string

const (
	NamespaceInstanceSearchParamsMessagesContentArrayObjectTypeText NamespaceInstanceSearchParamsMessagesContentArrayObjectType = "text"
)

func (r NamespaceInstanceSearchParamsMessagesContentArrayObjectType) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsMessagesContentArrayObjectTypeText:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsMessagesContentArrayType string

const (
	NamespaceInstanceSearchParamsMessagesContentArrayTypeText     NamespaceInstanceSearchParamsMessagesContentArrayType = "text"
	NamespaceInstanceSearchParamsMessagesContentArrayTypeImageURL NamespaceInstanceSearchParamsMessagesContentArrayType = "image_url"
	NamespaceInstanceSearchParamsMessagesContentArrayTypeFile     NamespaceInstanceSearchParamsMessagesContentArrayType = "file"
)

func (r NamespaceInstanceSearchParamsMessagesContentArrayType) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsMessagesContentArrayTypeText, NamespaceInstanceSearchParamsMessagesContentArrayTypeImageURL, NamespaceInstanceSearchParamsMessagesContentArrayTypeFile:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsMessagesRole string

const (
	NamespaceInstanceSearchParamsMessagesRoleSystem    NamespaceInstanceSearchParamsMessagesRole = "system"
	NamespaceInstanceSearchParamsMessagesRoleDeveloper NamespaceInstanceSearchParamsMessagesRole = "developer"
	NamespaceInstanceSearchParamsMessagesRoleUser      NamespaceInstanceSearchParamsMessagesRole = "user"
	NamespaceInstanceSearchParamsMessagesRoleAssistant NamespaceInstanceSearchParamsMessagesRole = "assistant"
	NamespaceInstanceSearchParamsMessagesRoleTool      NamespaceInstanceSearchParamsMessagesRole = "tool"
)

func (r NamespaceInstanceSearchParamsMessagesRole) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsMessagesRoleSystem, NamespaceInstanceSearchParamsMessagesRoleDeveloper, NamespaceInstanceSearchParamsMessagesRoleUser, NamespaceInstanceSearchParamsMessagesRoleAssistant, NamespaceInstanceSearchParamsMessagesRoleTool:
		return true
	}
	return false
}

type NamespaceInstanceSearchResponseEnvelope struct {
	Result  NamespaceInstanceSearchResponse             `json:"result" api:"required"`
	Success bool                                        `json:"success" api:"required"`
	JSON    namespaceInstanceSearchResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceSearchResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceSearchResponseEnvelope]
type namespaceInstanceSearchResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceStatsParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type NamespaceInstanceStatsResponseEnvelope struct {
	Result  NamespaceInstanceStatsResponse             `json:"result" api:"required"`
	Success bool                                       `json:"success" api:"required"`
	JSON    namespaceInstanceStatsResponseEnvelopeJSON `json:"-"`
}

// namespaceInstanceStatsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceInstanceStatsResponseEnvelope]
type namespaceInstanceStatsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceInstanceStatsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceStatsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
