// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_search

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/r2"
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

// Create a new instance.
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

// Update instance.
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

// List instances.
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

// List instances.
func (r *NamespaceInstanceService) ListAutoPaging(ctx context.Context, name string, params NamespaceInstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[NamespaceInstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, name, params, opts...))
}

// Delete instance.
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

// Read instance.
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

// Retrieves usage statistics for AI Search instances.
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
	ID             string                                       `json:"id" api:"required"`
	CreatedAt      time.Time                                    `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                                    `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                                       `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  NamespaceInstanceNewResponseAISearchModel    `json:"ai_search_model" api:"nullable"`
	Cache          bool                                         `json:"cache"`
	CacheThreshold NamespaceInstanceNewResponseCacheThreshold   `json:"cache_threshold"`
	ChunkOverlap   int64                                        `json:"chunk_overlap"`
	ChunkSize      int64                                        `json:"chunk_size"`
	CreatedBy      string                                       `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceNewResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel NamespaceInstanceNewResponseEmbeddingModel   `json:"embedding_model" api:"nullable"`
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
	RerankingModel       NamespaceInstanceNewResponseRerankingModel       `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceNewResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         NamespaceInstanceNewResponseRewriteModel         `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                             `json:"rewrite_query"`
	ScoreThreshold       float64                                          `json:"score_threshold"`
	Source               string                                           `json:"source" api:"nullable"`
	SourceParams         NamespaceInstanceNewResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                           `json:"status"`
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

type NamespaceInstanceNewResponseAISearchModel string

const (
	NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceNewResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceNewResponseAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceNewResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceNewResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceNewResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceNewResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceNewResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewResponseAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceNewResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceNewResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceNewResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceNewResponseAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceNewResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceNewResponseAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceNewResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceNewResponseAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceNewResponseAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceNewResponseAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceNewResponseAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceNewResponseAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceNewResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceNewResponseAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceNewResponseAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceNewResponseAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceNewResponseAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceNewResponseAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceNewResponseAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceNewResponseAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceNewResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceNewResponseAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceNewResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceNewResponseAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceNewResponseAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceNewResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceNewResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceNewResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceNewResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewResponseAISearchModelCerebrasGptOSs120b                    NamespaceInstanceNewResponseAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceNewResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceNewResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceNewResponseAISearchModelGrokGrok4                             NamespaceInstanceNewResponseAISearchModel = "grok/grok-4"
	NamespaceInstanceNewResponseAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceNewResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceNewResponseAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceNewResponseAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceNewResponseAISearchModelOpenAIGpt5                            NamespaceInstanceNewResponseAISearchModel = "openai/gpt-5"
	NamespaceInstanceNewResponseAISearchModelOpenAIGpt5Mini                        NamespaceInstanceNewResponseAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceNewResponseAISearchModelOpenAIGpt5Nano                        NamespaceInstanceNewResponseAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceNewResponseAISearchModelEmpty                                 NamespaceInstanceNewResponseAISearchModel = ""
)

func (r NamespaceInstanceNewResponseAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceNewResponseAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceNewResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceNewResponseAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceNewResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceNewResponseAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceNewResponseAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceNewResponseAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceNewResponseAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceNewResponseAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceNewResponseAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceNewResponseAISearchModelAnthropicClaudeOpus4, NamespaceInstanceNewResponseAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceNewResponseAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceNewResponseAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceNewResponseAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceNewResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceNewResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceNewResponseAISearchModelCerebrasGptOSs120b, NamespaceInstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceNewResponseAISearchModelGrokGrok4, NamespaceInstanceNewResponseAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceNewResponseAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceNewResponseAISearchModelOpenAIGpt5, NamespaceInstanceNewResponseAISearchModelOpenAIGpt5Mini, NamespaceInstanceNewResponseAISearchModelOpenAIGpt5Nano, NamespaceInstanceNewResponseAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceNewResponseEmbeddingModel string

const (
	NamespaceInstanceNewResponseEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceNewResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceNewResponseEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceNewResponseEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceNewResponseEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceNewResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceNewResponseEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceNewResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceNewResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceNewResponseEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceNewResponseEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceNewResponseEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceNewResponseEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceNewResponseEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceNewResponseEmbeddingModelEmpty                                 NamespaceInstanceNewResponseEmbeddingModel = ""
)

func (r NamespaceInstanceNewResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceNewResponseEmbeddingModelCfBaaiBgeM3, NamespaceInstanceNewResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceNewResponseEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceNewResponseEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceNewResponseEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceNewResponseEmbeddingModelEmpty:
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
	Enabled                 bool                                                                    `json:"enabled"`
	Mcp                     NamespaceInstanceNewResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceNewResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceNewResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    namespaceInstanceNewResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceNewResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceNewResponsePublicEndpointParams]
type namespaceInstanceNewResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
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

type NamespaceInstanceNewResponseRerankingModel string

const (
	NamespaceInstanceNewResponseRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceNewResponseRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceNewResponseRerankingModelEmpty                 NamespaceInstanceNewResponseRerankingModel = ""
)

func (r NamespaceInstanceNewResponseRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceNewResponseRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy []NamespaceInstanceNewResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceNewResponseRewriteModel string

const (
	NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceNewResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceNewResponseRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceNewResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceNewResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceNewResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceNewResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceNewResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewResponseRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceNewResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceNewResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceNewResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceNewResponseRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceNewResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceNewResponseRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceNewResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceNewResponseRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceNewResponseRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceNewResponseRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceNewResponseRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceNewResponseRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceNewResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceNewResponseRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceNewResponseRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceNewResponseRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceNewResponseRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceNewResponseRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceNewResponseRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceNewResponseRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceNewResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceNewResponseRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceNewResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceNewResponseRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceNewResponseRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceNewResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceNewResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceNewResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceNewResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewResponseRewriteModelCerebrasGptOSs120b                    NamespaceInstanceNewResponseRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceNewResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceNewResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceNewResponseRewriteModelGrokGrok4                             NamespaceInstanceNewResponseRewriteModel = "grok/grok-4"
	NamespaceInstanceNewResponseRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceNewResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceNewResponseRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceNewResponseRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceNewResponseRewriteModelOpenAIGpt5                            NamespaceInstanceNewResponseRewriteModel = "openai/gpt-5"
	NamespaceInstanceNewResponseRewriteModelOpenAIGpt5Mini                        NamespaceInstanceNewResponseRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceNewResponseRewriteModelOpenAIGpt5Nano                        NamespaceInstanceNewResponseRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceNewResponseRewriteModelEmpty                                 NamespaceInstanceNewResponseRewriteModel = ""
)

func (r NamespaceInstanceNewResponseRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceNewResponseRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceNewResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceNewResponseRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceNewResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceNewResponseRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceNewResponseRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceNewResponseRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceNewResponseRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceNewResponseRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceNewResponseRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceNewResponseRewriteModelAnthropicClaudeOpus4, NamespaceInstanceNewResponseRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceNewResponseRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceNewResponseRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceNewResponseRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceNewResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceNewResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceNewResponseRewriteModelCerebrasGptOSs120b, NamespaceInstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceNewResponseRewriteModelGrokGrok4, NamespaceInstanceNewResponseRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceNewResponseRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceNewResponseRewriteModelOpenAIGpt5, NamespaceInstanceNewResponseRewriteModelOpenAIGpt5Mini, NamespaceInstanceNewResponseRewriteModelOpenAIGpt5Nano, NamespaceInstanceNewResponseRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
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
	CrawlOptions NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptions `json:"crawl_options"`
	ParseOptions NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions NamespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         namespaceInstanceNewResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceNewResponseSourceParamsWebCrawler]
type namespaceInstanceNewResponseSourceParamsWebCrawlerJSON struct {
	CrawlOptions apijson.Field
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptions struct {
	Depth                float64                                                              `json:"depth"`
	IncludeExternalLinks bool                                                                 `json:"include_external_links"`
	IncludeSubdomains    bool                                                                 `json:"include_subdomains"`
	MaxAge               float64                                                              `json:"max_age"`
	Source               NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource `json:"source"`
	JSON                 namespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsJSON   `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptions]
type namespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceNewResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector []NamespaceInstanceNewResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders  map[string]string                                                               `json:"include_headers"`
	IncludeImages   bool                                                                            `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
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

type NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceNewResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceNewResponseSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                             `json:"storage_id" api:"required"`
	R2Jurisdiction string                                                             `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                        `json:"storage_type"`
	JSON           namespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// namespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptions]
type namespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
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
	ID             string                                          `json:"id" api:"required"`
	CreatedAt      time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                                       `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                                          `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  NamespaceInstanceUpdateResponseAISearchModel    `json:"ai_search_model" api:"nullable"`
	Cache          bool                                            `json:"cache"`
	CacheThreshold NamespaceInstanceUpdateResponseCacheThreshold   `json:"cache_threshold"`
	ChunkOverlap   int64                                           `json:"chunk_overlap"`
	ChunkSize      int64                                           `json:"chunk_size"`
	CreatedBy      string                                          `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceUpdateResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel NamespaceInstanceUpdateResponseEmbeddingModel   `json:"embedding_model" api:"nullable"`
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
	RerankingModel       NamespaceInstanceUpdateResponseRerankingModel       `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceUpdateResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         NamespaceInstanceUpdateResponseRewriteModel         `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                                `json:"rewrite_query"`
	ScoreThreshold       float64                                             `json:"score_threshold"`
	Source               string                                              `json:"source" api:"nullable"`
	SourceParams         NamespaceInstanceUpdateResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                              `json:"status"`
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

type NamespaceInstanceUpdateResponseAISearchModel string

const (
	NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceUpdateResponseAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceUpdateResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceUpdateResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateResponseAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceUpdateResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceUpdateResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceUpdateResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceUpdateResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceUpdateResponseAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceUpdateResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceUpdateResponseAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceUpdateResponseAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceUpdateResponseAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceUpdateResponseAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceUpdateResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceUpdateResponseAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceUpdateResponseAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceUpdateResponseAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceUpdateResponseAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceUpdateResponseAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceUpdateResponseAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceUpdateResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceUpdateResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceUpdateResponseAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceUpdateResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceUpdateResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateResponseAISearchModelCerebrasGptOSs120b                    NamespaceInstanceUpdateResponseAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceUpdateResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceUpdateResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceUpdateResponseAISearchModelGrokGrok4                             NamespaceInstanceUpdateResponseAISearchModel = "grok/grok-4"
	NamespaceInstanceUpdateResponseAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceUpdateResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceUpdateResponseAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceUpdateResponseAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5                            NamespaceInstanceUpdateResponseAISearchModel = "openai/gpt-5"
	NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5Mini                        NamespaceInstanceUpdateResponseAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5Nano                        NamespaceInstanceUpdateResponseAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceUpdateResponseAISearchModelEmpty                                 NamespaceInstanceUpdateResponseAISearchModel = ""
)

func (r NamespaceInstanceUpdateResponseAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceUpdateResponseAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceUpdateResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceUpdateResponseAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceUpdateResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceUpdateResponseAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceUpdateResponseAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceUpdateResponseAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceUpdateResponseAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceUpdateResponseAISearchModelAnthropicClaudeOpus4, NamespaceInstanceUpdateResponseAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceUpdateResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceUpdateResponseAISearchModelCerebrasGptOSs120b, NamespaceInstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceUpdateResponseAISearchModelGrokGrok4, NamespaceInstanceUpdateResponseAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceUpdateResponseAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5, NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5Mini, NamespaceInstanceUpdateResponseAISearchModelOpenAIGpt5Nano, NamespaceInstanceUpdateResponseAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceUpdateResponseEmbeddingModel string

const (
	NamespaceInstanceUpdateResponseEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceUpdateResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceUpdateResponseEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceUpdateResponseEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceUpdateResponseEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceUpdateResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceUpdateResponseEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceUpdateResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceUpdateResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceUpdateResponseEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceUpdateResponseEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceUpdateResponseEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceUpdateResponseEmbeddingModelEmpty                                 NamespaceInstanceUpdateResponseEmbeddingModel = ""
)

func (r NamespaceInstanceUpdateResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceUpdateResponseEmbeddingModelCfBaaiBgeM3, NamespaceInstanceUpdateResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceUpdateResponseEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceUpdateResponseEmbeddingModelEmpty:
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
	Enabled                 bool                                                                       `json:"enabled"`
	Mcp                     NamespaceInstanceUpdateResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceUpdateResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceUpdateResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    namespaceInstanceUpdateResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceUpdateResponsePublicEndpointParamsJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponsePublicEndpointParams]
type namespaceInstanceUpdateResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
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

type NamespaceInstanceUpdateResponseRerankingModel string

const (
	NamespaceInstanceUpdateResponseRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceUpdateResponseRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceUpdateResponseRerankingModelEmpty                 NamespaceInstanceUpdateResponseRerankingModel = ""
)

func (r NamespaceInstanceUpdateResponseRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceUpdateResponseRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy []NamespaceInstanceUpdateResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceUpdateResponseRewriteModel string

const (
	NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceUpdateResponseRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceUpdateResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceUpdateResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateResponseRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceUpdateResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceUpdateResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceUpdateResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceUpdateResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceUpdateResponseRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceUpdateResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceUpdateResponseRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceUpdateResponseRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceUpdateResponseRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceUpdateResponseRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceUpdateResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceUpdateResponseRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceUpdateResponseRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceUpdateResponseRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceUpdateResponseRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceUpdateResponseRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceUpdateResponseRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceUpdateResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceUpdateResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceUpdateResponseRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceUpdateResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceUpdateResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateResponseRewriteModelCerebrasGptOSs120b                    NamespaceInstanceUpdateResponseRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceUpdateResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceUpdateResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceUpdateResponseRewriteModelGrokGrok4                             NamespaceInstanceUpdateResponseRewriteModel = "grok/grok-4"
	NamespaceInstanceUpdateResponseRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceUpdateResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceUpdateResponseRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceUpdateResponseRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5                            NamespaceInstanceUpdateResponseRewriteModel = "openai/gpt-5"
	NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5Mini                        NamespaceInstanceUpdateResponseRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5Nano                        NamespaceInstanceUpdateResponseRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceUpdateResponseRewriteModelEmpty                                 NamespaceInstanceUpdateResponseRewriteModel = ""
)

func (r NamespaceInstanceUpdateResponseRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceUpdateResponseRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceUpdateResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceUpdateResponseRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceUpdateResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceUpdateResponseRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceUpdateResponseRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceUpdateResponseRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceUpdateResponseRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceUpdateResponseRewriteModelAnthropicClaudeOpus4, NamespaceInstanceUpdateResponseRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceUpdateResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceUpdateResponseRewriteModelCerebrasGptOSs120b, NamespaceInstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceUpdateResponseRewriteModelGrokGrok4, NamespaceInstanceUpdateResponseRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceUpdateResponseRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5, NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5Mini, NamespaceInstanceUpdateResponseRewriteModelOpenAIGpt5Nano, NamespaceInstanceUpdateResponseRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
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
	CrawlOptions NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptions `json:"crawl_options"`
	ParseOptions NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions NamespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceUpdateResponseSourceParamsWebCrawler]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON struct {
	CrawlOptions apijson.Field
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptions struct {
	Depth                float64                                                                 `json:"depth"`
	IncludeExternalLinks bool                                                                    `json:"include_external_links"`
	IncludeSubdomains    bool                                                                    `json:"include_subdomains"`
	MaxAge               float64                                                                 `json:"max_age"`
	Source               NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource `json:"source"`
	JSON                 namespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsJSON   `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptions]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector []NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders  map[string]string                                                                  `json:"include_headers"`
	IncludeImages   bool                                                                               `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
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

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceUpdateResponseSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                                `json:"storage_id" api:"required"`
	R2Jurisdiction string                                                                `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                           `json:"storage_type"`
	JSON           namespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// namespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptions]
type namespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
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
	// AI Search instance ID. Lowercase alphanumeric, hyphens, and underscores.
	ID             string                                        `json:"id" api:"required"`
	CreatedAt      time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                                     `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                                        `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  NamespaceInstanceListResponseAISearchModel    `json:"ai_search_model" api:"nullable"`
	Cache          bool                                          `json:"cache"`
	CacheThreshold NamespaceInstanceListResponseCacheThreshold   `json:"cache_threshold"`
	ChunkOverlap   int64                                         `json:"chunk_overlap"`
	ChunkSize      int64                                         `json:"chunk_size"`
	CreatedBy      string                                        `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceListResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel NamespaceInstanceListResponseEmbeddingModel   `json:"embedding_model" api:"nullable"`
	Enable         bool                                          `json:"enable"`
	EngineVersion  float64                                       `json:"engine_version"`
	FusionMethod   NamespaceInstanceListResponseFusionMethod     `json:"fusion_method"`
	// Deprecated — use index_method instead.
	//
	// Deprecated: deprecated
	HybridSearchEnabled bool `json:"hybrid_search_enabled"`
	// Controls which storage backends are used during indexing. Defaults to
	// vector-only.
	IndexMethod          NamespaceInstanceListResponseIndexMethod          `json:"index_method"`
	IndexingOptions      NamespaceInstanceListResponseIndexingOptions      `json:"indexing_options" api:"nullable"`
	LastActivity         time.Time                                         `json:"last_activity" api:"nullable" format:"date-time"`
	MaxNumResults        int64                                             `json:"max_num_results"`
	Metadata             NamespaceInstanceListResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                            `json:"modified_by" api:"nullable"`
	Namespace            string                                            `json:"namespace" api:"nullable"`
	Paused               bool                                              `json:"paused"`
	PublicEndpointID     string                                            `json:"public_endpoint_id" api:"nullable"`
	PublicEndpointParams NamespaceInstanceListResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                              `json:"reranking"`
	RerankingModel       NamespaceInstanceListResponseRerankingModel       `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceListResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         NamespaceInstanceListResponseRewriteModel         `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                              `json:"rewrite_query"`
	ScoreThreshold       float64                                           `json:"score_threshold"`
	Source               string                                            `json:"source" api:"nullable"`
	SourceParams         NamespaceInstanceListResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                            `json:"status"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval NamespaceInstanceListResponseSyncInterval `json:"sync_interval"`
	TokenID      string                                    `json:"token_id" format:"uuid"`
	Type         NamespaceInstanceListResponseType         `json:"type" api:"nullable"`
	JSON         namespaceInstanceListResponseJSON         `json:"-"`
}

// namespaceInstanceListResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceListResponse]
type namespaceInstanceListResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	AIGatewayID          apijson.Field
	AISearchModel        apijson.Field
	Cache                apijson.Field
	CacheThreshold       apijson.Field
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

func (r *NamespaceInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseAISearchModel string

const (
	NamespaceInstanceListResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceListResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceListResponseAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceListResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceListResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceListResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceListResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceListResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceListResponseAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceListResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceListResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceListResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceListResponseAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceListResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceListResponseAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceListResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceListResponseAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceListResponseAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceListResponseAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceListResponseAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceListResponseAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceListResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceListResponseAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceListResponseAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceListResponseAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceListResponseAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceListResponseAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceListResponseAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceListResponseAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceListResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceListResponseAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceListResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceListResponseAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceListResponseAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceListResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceListResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceListResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceListResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceListResponseAISearchModelCerebrasGptOSs120b                    NamespaceInstanceListResponseAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceListResponseAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceListResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceListResponseAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceListResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceListResponseAISearchModelGrokGrok4                             NamespaceInstanceListResponseAISearchModel = "grok/grok-4"
	NamespaceInstanceListResponseAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceListResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceListResponseAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceListResponseAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceListResponseAISearchModelOpenAIGpt5                            NamespaceInstanceListResponseAISearchModel = "openai/gpt-5"
	NamespaceInstanceListResponseAISearchModelOpenAIGpt5Mini                        NamespaceInstanceListResponseAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceListResponseAISearchModelOpenAIGpt5Nano                        NamespaceInstanceListResponseAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceListResponseAISearchModelEmpty                                 NamespaceInstanceListResponseAISearchModel = ""
)

func (r NamespaceInstanceListResponseAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceListResponseAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceListResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceListResponseAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceListResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceListResponseAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceListResponseAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceListResponseAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceListResponseAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceListResponseAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceListResponseAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceListResponseAISearchModelAnthropicClaudeOpus4, NamespaceInstanceListResponseAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceListResponseAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceListResponseAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceListResponseAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceListResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceListResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceListResponseAISearchModelCerebrasGptOSs120b, NamespaceInstanceListResponseAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceListResponseAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceListResponseAISearchModelGrokGrok4, NamespaceInstanceListResponseAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceListResponseAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceListResponseAISearchModelOpenAIGpt5, NamespaceInstanceListResponseAISearchModelOpenAIGpt5Mini, NamespaceInstanceListResponseAISearchModelOpenAIGpt5Nano, NamespaceInstanceListResponseAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceListResponseEmbeddingModel string

const (
	NamespaceInstanceListResponseEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceListResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceListResponseEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceListResponseEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceListResponseEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceListResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceListResponseEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceListResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceListResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceListResponseEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceListResponseEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceListResponseEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceListResponseEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceListResponseEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceListResponseEmbeddingModelEmpty                                 NamespaceInstanceListResponseEmbeddingModel = ""
)

func (r NamespaceInstanceListResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceListResponseEmbeddingModelCfBaaiBgeM3, NamespaceInstanceListResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceListResponseEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceListResponseEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceListResponseEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceListResponseEmbeddingModelEmpty:
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

// Controls which storage backends are used during indexing. Defaults to
// vector-only.
type NamespaceInstanceListResponseIndexMethod struct {
	// Enable keyword (BM25) storage backend.
	Keyword bool `json:"keyword" api:"required"`
	// Enable vector (embedding) storage backend.
	Vector bool                                         `json:"vector" api:"required"`
	JSON   namespaceInstanceListResponseIndexMethodJSON `json:"-"`
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
	// Tokenizer used for keyword search indexing. porter provides word-level
	// tokenization with Porter stemming (good for natural language queries). trigram
	// enables character-level substring matching (good for partial matches, code,
	// identifiers). Changing this triggers a full re-index. Defaults to porter.
	KeywordTokenizer NamespaceInstanceListResponseIndexingOptionsKeywordTokenizer `json:"keyword_tokenizer"`
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

// Tokenizer used for keyword search indexing. porter provides word-level
// tokenization with Porter stemming (good for natural language queries). trigram
// enables character-level substring matching (good for partial matches, code,
// identifiers). Changing this triggers a full re-index. Defaults to porter.
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
	Enabled                 bool                                                                     `json:"enabled"`
	Mcp                     NamespaceInstanceListResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceListResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceListResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    namespaceInstanceListResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceListResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceListResponsePublicEndpointParams]
type namespaceInstanceListResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
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
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                         `json:"disabled"`
	JSON     namespaceInstanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
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
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                                     `json:"disabled"`
	JSON     namespaceInstanceListResponsePublicEndpointParamsMcpJSON `json:"-"`
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
	PeriodMs  int64                                                               `json:"period_ms"`
	Requests  int64                                                               `json:"requests"`
	Technique NamespaceInstanceListResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      namespaceInstanceListResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
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
	// Disable search endpoint for this public endpoint
	Disabled bool                                                                `json:"disabled"`
	JSON     namespaceInstanceListResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
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

type NamespaceInstanceListResponseRerankingModel string

const (
	NamespaceInstanceListResponseRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceListResponseRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceListResponseRerankingModelEmpty                 NamespaceInstanceListResponseRerankingModel = ""
)

func (r NamespaceInstanceListResponseRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceListResponseRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceListResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy []NamespaceInstanceListResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
	KeywordMatchMode NamespaceInstanceListResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
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
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction NamespaceInstanceListResponseRetrievalOptionsBoostByDirection `json:"direction"`
	JSON      namespaceInstanceListResponseRetrievalOptionsBoostByJSON      `json:"-"`
}

// namespaceInstanceListResponseRetrievalOptionsBoostByJSON contains the JSON
// metadata for the struct [NamespaceInstanceListResponseRetrievalOptionsBoostBy]
type namespaceInstanceListResponseRetrievalOptionsBoostByJSON struct {
	Field       apijson.Field
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

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
// for numeric/datetime fields, 'exists' for text/boolean fields.
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

// Controls which documents are candidates for BM25 scoring. 'and' restricts
// candidates to documents containing all query terms; 'or' includes any document
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceListResponseRewriteModel string

const (
	NamespaceInstanceListResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceListResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceListResponseRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceListResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceListResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceListResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceListResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceListResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceListResponseRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceListResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceListResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceListResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceListResponseRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceListResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceListResponseRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceListResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceListResponseRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceListResponseRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceListResponseRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceListResponseRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceListResponseRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceListResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceListResponseRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceListResponseRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceListResponseRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceListResponseRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceListResponseRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceListResponseRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceListResponseRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceListResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceListResponseRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceListResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceListResponseRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceListResponseRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceListResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceListResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceListResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceListResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceListResponseRewriteModelCerebrasGptOSs120b                    NamespaceInstanceListResponseRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceListResponseRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceListResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceListResponseRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceListResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceListResponseRewriteModelGrokGrok4                             NamespaceInstanceListResponseRewriteModel = "grok/grok-4"
	NamespaceInstanceListResponseRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceListResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceListResponseRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceListResponseRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceListResponseRewriteModelOpenAIGpt5                            NamespaceInstanceListResponseRewriteModel = "openai/gpt-5"
	NamespaceInstanceListResponseRewriteModelOpenAIGpt5Mini                        NamespaceInstanceListResponseRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceListResponseRewriteModelOpenAIGpt5Nano                        NamespaceInstanceListResponseRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceListResponseRewriteModelEmpty                                 NamespaceInstanceListResponseRewriteModel = ""
)

func (r NamespaceInstanceListResponseRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceListResponseRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceListResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceListResponseRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceListResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceListResponseRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceListResponseRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceListResponseRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceListResponseRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceListResponseRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceListResponseRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceListResponseRewriteModelAnthropicClaudeOpus4, NamespaceInstanceListResponseRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceListResponseRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceListResponseRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceListResponseRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceListResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceListResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceListResponseRewriteModelCerebrasGptOSs120b, NamespaceInstanceListResponseRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceListResponseRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceListResponseRewriteModelGrokGrok4, NamespaceInstanceListResponseRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceListResponseRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceListResponseRewriteModelOpenAIGpt5, NamespaceInstanceListResponseRewriteModelOpenAIGpt5Mini, NamespaceInstanceListResponseRewriteModelOpenAIGpt5Nano, NamespaceInstanceListResponseRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                            `json:"include_items"`
	Prefix         string                                              `json:"prefix"`
	R2Jurisdiction string                                              `json:"r2_jurisdiction"`
	WebCrawler     NamespaceInstanceListResponseSourceParamsWebCrawler `json:"web_crawler"`
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
	CrawlOptions NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptions `json:"crawl_options"`
	ParseOptions NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    NamespaceInstanceListResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions NamespaceInstanceListResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         namespaceInstanceListResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceListResponseSourceParamsWebCrawler]
type namespaceInstanceListResponseSourceParamsWebCrawlerJSON struct {
	CrawlOptions apijson.Field
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptions struct {
	Depth                float64                                                               `json:"depth"`
	IncludeExternalLinks bool                                                                  `json:"include_external_links"`
	IncludeSubdomains    bool                                                                  `json:"include_subdomains"`
	MaxAge               float64                                                               `json:"max_age"`
	Source               NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource `json:"source"`
	JSON                 namespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsJSON   `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptions]
type namespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceListResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector []NamespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders  map[string]string                                                                `json:"include_headers"`
	IncludeImages   bool                                                                             `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                            `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                                `json:"use_browser_rendering"`
	JSON                namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
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
	// Glob pattern to match against the page URL path. Uses standard glob syntax: \*
	// matches within a segment, \*\* crosses directories.
	Path string `json:"path" api:"required"`
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
	Selector string                                                                             `json:"selector" api:"required"`
	JSON     namespaceInstanceListResponseSourceParamsWebCrawlerParseOptionsContentSelectorJSON `json:"-"`
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
	NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceListResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceListResponseSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceListResponseSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceListResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceListResponseSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceListResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                              `json:"storage_id" api:"required"`
	R2Jurisdiction string                                                              `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                         `json:"storage_type"`
	JSON           namespaceInstanceListResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// namespaceInstanceListResponseSourceParamsWebCrawlerStoreOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceListResponseSourceParamsWebCrawlerStoreOptions]
type namespaceInstanceListResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceListResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceListResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
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
	ID             string                                          `json:"id" api:"required"`
	CreatedAt      time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                                       `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                                          `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  NamespaceInstanceDeleteResponseAISearchModel    `json:"ai_search_model" api:"nullable"`
	Cache          bool                                            `json:"cache"`
	CacheThreshold NamespaceInstanceDeleteResponseCacheThreshold   `json:"cache_threshold"`
	ChunkOverlap   int64                                           `json:"chunk_overlap"`
	ChunkSize      int64                                           `json:"chunk_size"`
	CreatedBy      string                                          `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceDeleteResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel NamespaceInstanceDeleteResponseEmbeddingModel   `json:"embedding_model" api:"nullable"`
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
	RerankingModel       NamespaceInstanceDeleteResponseRerankingModel       `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceDeleteResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         NamespaceInstanceDeleteResponseRewriteModel         `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                                `json:"rewrite_query"`
	ScoreThreshold       float64                                             `json:"score_threshold"`
	Source               string                                              `json:"source" api:"nullable"`
	SourceParams         NamespaceInstanceDeleteResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                              `json:"status"`
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

type NamespaceInstanceDeleteResponseAISearchModel string

const (
	NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceDeleteResponseAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceDeleteResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceDeleteResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceDeleteResponseAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceDeleteResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceDeleteResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceDeleteResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceDeleteResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceDeleteResponseAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceDeleteResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceDeleteResponseAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceDeleteResponseAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceDeleteResponseAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceDeleteResponseAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceDeleteResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceDeleteResponseAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceDeleteResponseAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceDeleteResponseAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceDeleteResponseAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceDeleteResponseAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceDeleteResponseAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceDeleteResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceDeleteResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceDeleteResponseAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceDeleteResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceDeleteResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceDeleteResponseAISearchModelCerebrasGptOSs120b                    NamespaceInstanceDeleteResponseAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceDeleteResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceDeleteResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceDeleteResponseAISearchModelGrokGrok4                             NamespaceInstanceDeleteResponseAISearchModel = "grok/grok-4"
	NamespaceInstanceDeleteResponseAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceDeleteResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceDeleteResponseAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceDeleteResponseAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5                            NamespaceInstanceDeleteResponseAISearchModel = "openai/gpt-5"
	NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5Mini                        NamespaceInstanceDeleteResponseAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5Nano                        NamespaceInstanceDeleteResponseAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceDeleteResponseAISearchModelEmpty                                 NamespaceInstanceDeleteResponseAISearchModel = ""
)

func (r NamespaceInstanceDeleteResponseAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceDeleteResponseAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceDeleteResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceDeleteResponseAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceDeleteResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceDeleteResponseAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceDeleteResponseAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceDeleteResponseAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceDeleteResponseAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceDeleteResponseAISearchModelAnthropicClaudeOpus4, NamespaceInstanceDeleteResponseAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceDeleteResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceDeleteResponseAISearchModelCerebrasGptOSs120b, NamespaceInstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceDeleteResponseAISearchModelGrokGrok4, NamespaceInstanceDeleteResponseAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceDeleteResponseAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5, NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5Mini, NamespaceInstanceDeleteResponseAISearchModelOpenAIGpt5Nano, NamespaceInstanceDeleteResponseAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceDeleteResponseEmbeddingModel string

const (
	NamespaceInstanceDeleteResponseEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceDeleteResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceDeleteResponseEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceDeleteResponseEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceDeleteResponseEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceDeleteResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceDeleteResponseEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceDeleteResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceDeleteResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceDeleteResponseEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceDeleteResponseEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceDeleteResponseEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceDeleteResponseEmbeddingModelEmpty                                 NamespaceInstanceDeleteResponseEmbeddingModel = ""
)

func (r NamespaceInstanceDeleteResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceDeleteResponseEmbeddingModelCfBaaiBgeM3, NamespaceInstanceDeleteResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceDeleteResponseEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceDeleteResponseEmbeddingModelEmpty:
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
	Enabled                 bool                                                                       `json:"enabled"`
	Mcp                     NamespaceInstanceDeleteResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceDeleteResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceDeleteResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    namespaceInstanceDeleteResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceDeleteResponsePublicEndpointParamsJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponsePublicEndpointParams]
type namespaceInstanceDeleteResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
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

type NamespaceInstanceDeleteResponseRerankingModel string

const (
	NamespaceInstanceDeleteResponseRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceDeleteResponseRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceDeleteResponseRerankingModelEmpty                 NamespaceInstanceDeleteResponseRerankingModel = ""
)

func (r NamespaceInstanceDeleteResponseRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceDeleteResponseRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy []NamespaceInstanceDeleteResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceDeleteResponseRewriteModel string

const (
	NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceDeleteResponseRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceDeleteResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceDeleteResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceDeleteResponseRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceDeleteResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceDeleteResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceDeleteResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceDeleteResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceDeleteResponseRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceDeleteResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceDeleteResponseRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceDeleteResponseRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceDeleteResponseRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceDeleteResponseRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceDeleteResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceDeleteResponseRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceDeleteResponseRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceDeleteResponseRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceDeleteResponseRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceDeleteResponseRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceDeleteResponseRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceDeleteResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceDeleteResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceDeleteResponseRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceDeleteResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceDeleteResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceDeleteResponseRewriteModelCerebrasGptOSs120b                    NamespaceInstanceDeleteResponseRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceDeleteResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceDeleteResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceDeleteResponseRewriteModelGrokGrok4                             NamespaceInstanceDeleteResponseRewriteModel = "grok/grok-4"
	NamespaceInstanceDeleteResponseRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceDeleteResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceDeleteResponseRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceDeleteResponseRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5                            NamespaceInstanceDeleteResponseRewriteModel = "openai/gpt-5"
	NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5Mini                        NamespaceInstanceDeleteResponseRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5Nano                        NamespaceInstanceDeleteResponseRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceDeleteResponseRewriteModelEmpty                                 NamespaceInstanceDeleteResponseRewriteModel = ""
)

func (r NamespaceInstanceDeleteResponseRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceDeleteResponseRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceDeleteResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceDeleteResponseRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceDeleteResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceDeleteResponseRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceDeleteResponseRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceDeleteResponseRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceDeleteResponseRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceDeleteResponseRewriteModelAnthropicClaudeOpus4, NamespaceInstanceDeleteResponseRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceDeleteResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceDeleteResponseRewriteModelCerebrasGptOSs120b, NamespaceInstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceDeleteResponseRewriteModelGrokGrok4, NamespaceInstanceDeleteResponseRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceDeleteResponseRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5, NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5Mini, NamespaceInstanceDeleteResponseRewriteModelOpenAIGpt5Nano, NamespaceInstanceDeleteResponseRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
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
	CrawlOptions NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptions `json:"crawl_options"`
	ParseOptions NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions NamespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceDeleteResponseSourceParamsWebCrawler]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON struct {
	CrawlOptions apijson.Field
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptions struct {
	Depth                float64                                                                 `json:"depth"`
	IncludeExternalLinks bool                                                                    `json:"include_external_links"`
	IncludeSubdomains    bool                                                                    `json:"include_subdomains"`
	MaxAge               float64                                                                 `json:"max_age"`
	Source               NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource `json:"source"`
	JSON                 namespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsJSON   `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptions]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector []NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders  map[string]string                                                                  `json:"include_headers"`
	IncludeImages   bool                                                                               `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
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

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceDeleteResponseSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                                `json:"storage_id" api:"required"`
	R2Jurisdiction string                                                                `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                           `json:"storage_type"`
	JSON           namespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// namespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON contains
// the JSON metadata for the struct
// [NamespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptions]
type namespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
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
	Content     string                                                     `json:"content" api:"required,nullable"`
	Role        NamespaceInstanceChatCompletionsResponseChoicesMessageRole `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                     `json:"-" api:"extrafields"`
	JSON        namespaceInstanceChatCompletionsResponseChoicesMessageJSON `json:"-"`
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
	ID             string                                        `json:"id" api:"required"`
	CreatedAt      time.Time                                     `json:"created_at" api:"required" format:"date-time"`
	ModifiedAt     time.Time                                     `json:"modified_at" api:"required" format:"date-time"`
	AIGatewayID    string                                        `json:"ai_gateway_id" api:"nullable"`
	AISearchModel  NamespaceInstanceReadResponseAISearchModel    `json:"ai_search_model" api:"nullable"`
	Cache          bool                                          `json:"cache"`
	CacheThreshold NamespaceInstanceReadResponseCacheThreshold   `json:"cache_threshold"`
	ChunkOverlap   int64                                         `json:"chunk_overlap"`
	ChunkSize      int64                                         `json:"chunk_size"`
	CreatedBy      string                                        `json:"created_by" api:"nullable"`
	CustomMetadata []NamespaceInstanceReadResponseCustomMetadata `json:"custom_metadata"`
	EmbeddingModel NamespaceInstanceReadResponseEmbeddingModel   `json:"embedding_model" api:"nullable"`
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
	RerankingModel       NamespaceInstanceReadResponseRerankingModel       `json:"reranking_model" api:"nullable"`
	RetrievalOptions     NamespaceInstanceReadResponseRetrievalOptions     `json:"retrieval_options" api:"nullable"`
	RewriteModel         NamespaceInstanceReadResponseRewriteModel         `json:"rewrite_model" api:"nullable"`
	RewriteQuery         bool                                              `json:"rewrite_query"`
	ScoreThreshold       float64                                           `json:"score_threshold"`
	Source               string                                            `json:"source" api:"nullable"`
	SourceParams         NamespaceInstanceReadResponseSourceParams         `json:"source_params" api:"nullable"`
	Status               string                                            `json:"status"`
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

type NamespaceInstanceReadResponseAISearchModel string

const (
	NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceReadResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceReadResponseAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceReadResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceReadResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceReadResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceReadResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceReadResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceReadResponseAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceReadResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceReadResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceReadResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceReadResponseAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceReadResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceReadResponseAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceReadResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceReadResponseAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceReadResponseAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceReadResponseAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceReadResponseAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceReadResponseAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceReadResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceReadResponseAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceReadResponseAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceReadResponseAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceReadResponseAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceReadResponseAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceReadResponseAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceReadResponseAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceReadResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceReadResponseAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceReadResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceReadResponseAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceReadResponseAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceReadResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceReadResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceReadResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceReadResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceReadResponseAISearchModelCerebrasGptOSs120b                    NamespaceInstanceReadResponseAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceReadResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceReadResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceReadResponseAISearchModelGrokGrok4                             NamespaceInstanceReadResponseAISearchModel = "grok/grok-4"
	NamespaceInstanceReadResponseAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceReadResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceReadResponseAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceReadResponseAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceReadResponseAISearchModelOpenAIGpt5                            NamespaceInstanceReadResponseAISearchModel = "openai/gpt-5"
	NamespaceInstanceReadResponseAISearchModelOpenAIGpt5Mini                        NamespaceInstanceReadResponseAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceReadResponseAISearchModelOpenAIGpt5Nano                        NamespaceInstanceReadResponseAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceReadResponseAISearchModelEmpty                                 NamespaceInstanceReadResponseAISearchModel = ""
)

func (r NamespaceInstanceReadResponseAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceReadResponseAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceReadResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceReadResponseAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceReadResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceReadResponseAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceReadResponseAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceReadResponseAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceReadResponseAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceReadResponseAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceReadResponseAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceReadResponseAISearchModelAnthropicClaudeOpus4, NamespaceInstanceReadResponseAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceReadResponseAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceReadResponseAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceReadResponseAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceReadResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceReadResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceReadResponseAISearchModelCerebrasGptOSs120b, NamespaceInstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceReadResponseAISearchModelGrokGrok4, NamespaceInstanceReadResponseAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceReadResponseAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceReadResponseAISearchModelOpenAIGpt5, NamespaceInstanceReadResponseAISearchModelOpenAIGpt5Mini, NamespaceInstanceReadResponseAISearchModelOpenAIGpt5Nano, NamespaceInstanceReadResponseAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceReadResponseEmbeddingModel string

const (
	NamespaceInstanceReadResponseEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceReadResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceReadResponseEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceReadResponseEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceReadResponseEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceReadResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceReadResponseEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceReadResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceReadResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceReadResponseEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceReadResponseEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceReadResponseEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceReadResponseEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceReadResponseEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceReadResponseEmbeddingModelEmpty                                 NamespaceInstanceReadResponseEmbeddingModel = ""
)

func (r NamespaceInstanceReadResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceReadResponseEmbeddingModelCfBaaiBgeM3, NamespaceInstanceReadResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceReadResponseEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceReadResponseEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceReadResponseEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceReadResponseEmbeddingModelEmpty:
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
	Enabled                 bool                                                                     `json:"enabled"`
	Mcp                     NamespaceInstanceReadResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               NamespaceInstanceReadResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          NamespaceInstanceReadResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    namespaceInstanceReadResponsePublicEndpointParamsJSON                    `json:"-"`
}

// namespaceInstanceReadResponsePublicEndpointParamsJSON contains the JSON metadata
// for the struct [NamespaceInstanceReadResponsePublicEndpointParams]
type namespaceInstanceReadResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
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

type NamespaceInstanceReadResponseRerankingModel string

const (
	NamespaceInstanceReadResponseRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceReadResponseRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceReadResponseRerankingModelEmpty                 NamespaceInstanceReadResponseRerankingModel = ""
)

func (r NamespaceInstanceReadResponseRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceReadResponseRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy []NamespaceInstanceReadResponseRetrievalOptionsBoostBy `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field string `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceReadResponseRewriteModel string

const (
	NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceReadResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceReadResponseRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceReadResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceReadResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceReadResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceReadResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceReadResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceReadResponseRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceReadResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceReadResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceReadResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceReadResponseRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceReadResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceReadResponseRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceReadResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceReadResponseRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceReadResponseRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceReadResponseRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceReadResponseRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceReadResponseRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceReadResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceReadResponseRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceReadResponseRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceReadResponseRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceReadResponseRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceReadResponseRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceReadResponseRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceReadResponseRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceReadResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceReadResponseRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceReadResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceReadResponseRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceReadResponseRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceReadResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceReadResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceReadResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceReadResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceReadResponseRewriteModelCerebrasGptOSs120b                    NamespaceInstanceReadResponseRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceReadResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceReadResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceReadResponseRewriteModelGrokGrok4                             NamespaceInstanceReadResponseRewriteModel = "grok/grok-4"
	NamespaceInstanceReadResponseRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceReadResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceReadResponseRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceReadResponseRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceReadResponseRewriteModelOpenAIGpt5                            NamespaceInstanceReadResponseRewriteModel = "openai/gpt-5"
	NamespaceInstanceReadResponseRewriteModelOpenAIGpt5Mini                        NamespaceInstanceReadResponseRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceReadResponseRewriteModelOpenAIGpt5Nano                        NamespaceInstanceReadResponseRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceReadResponseRewriteModelEmpty                                 NamespaceInstanceReadResponseRewriteModel = ""
)

func (r NamespaceInstanceReadResponseRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceReadResponseRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceReadResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceReadResponseRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceReadResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceReadResponseRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceReadResponseRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceReadResponseRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceReadResponseRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceReadResponseRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceReadResponseRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceReadResponseRewriteModelAnthropicClaudeOpus4, NamespaceInstanceReadResponseRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceReadResponseRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceReadResponseRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceReadResponseRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceReadResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceReadResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceReadResponseRewriteModelCerebrasGptOSs120b, NamespaceInstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceReadResponseRewriteModelGrokGrok4, NamespaceInstanceReadResponseRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceReadResponseRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceReadResponseRewriteModelOpenAIGpt5, NamespaceInstanceReadResponseRewriteModelOpenAIGpt5Mini, NamespaceInstanceReadResponseRewriteModelOpenAIGpt5Nano, NamespaceInstanceReadResponseRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
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
	CrawlOptions NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptions `json:"crawl_options"`
	ParseOptions NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions NamespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         namespaceInstanceReadResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerJSON contains the JSON
// metadata for the struct [NamespaceInstanceReadResponseSourceParamsWebCrawler]
type namespaceInstanceReadResponseSourceParamsWebCrawlerJSON struct {
	CrawlOptions apijson.Field
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptions struct {
	Depth                float64                                                               `json:"depth"`
	IncludeExternalLinks bool                                                                  `json:"include_external_links"`
	IncludeSubdomains    bool                                                                  `json:"include_subdomains"`
	MaxAge               float64                                                               `json:"max_age"`
	Source               NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource `json:"source"`
	JSON                 namespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsJSON   `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptions]
type namespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsJSON struct {
	Depth                apijson.Field
	IncludeExternalLinks apijson.Field
	IncludeSubdomains    apijson.Field
	MaxAge               apijson.Field
	Source               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsJSON) RawJSON() string {
	return r.raw
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceReadResponseSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector []NamespaceInstanceReadResponseSourceParamsWebCrawlerParseOptionsContentSelector `json:"content_selector"`
	IncludeHeaders  map[string]string                                                                `json:"include_headers"`
	IncludeImages   bool                                                                             `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
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

type NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceReadResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceReadResponseSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                              `json:"storage_id" api:"required"`
	R2Jurisdiction string                                                              `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                         `json:"storage_type"`
	JSON           namespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// namespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON contains the
// JSON metadata for the struct
// [NamespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptions]
type namespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NamespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceInstanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
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
	Chunks      []NamespaceInstanceSearchResponseChunk `json:"chunks" api:"required"`
	SearchQuery string                                 `json:"search_query" api:"required"`
	JSON        namespaceInstanceSearchResponseJSON    `json:"-"`
}

// namespaceInstanceSearchResponseJSON contains the JSON metadata for the struct
// [NamespaceInstanceSearchResponse]
type namespaceInstanceSearchResponseJSON struct {
	Chunks      apijson.Field
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

type NamespaceInstanceStatsResponse struct {
	Completed int64 `json:"completed"`
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
	ID             param.Field[string]                                     `json:"id" api:"required"`
	AIGatewayID    param.Field[string]                                     `json:"ai_gateway_id"`
	AISearchModel  param.Field[NamespaceInstanceNewParamsAISearchModel]    `json:"ai_search_model"`
	Cache          param.Field[bool]                                       `json:"cache"`
	CacheThreshold param.Field[NamespaceInstanceNewParamsCacheThreshold]   `json:"cache_threshold"`
	Chunk          param.Field[bool]                                       `json:"chunk"`
	ChunkOverlap   param.Field[int64]                                      `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                                      `json:"chunk_size"`
	CustomMetadata param.Field[[]NamespaceInstanceNewParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[NamespaceInstanceNewParamsEmbeddingModel]   `json:"embedding_model"`
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
	RerankingModel       param.Field[NamespaceInstanceNewParamsRerankingModel]       `json:"reranking_model"`
	RetrievalOptions     param.Field[NamespaceInstanceNewParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel         param.Field[NamespaceInstanceNewParamsRewriteModel]         `json:"rewrite_model"`
	RewriteQuery         param.Field[bool]                                           `json:"rewrite_query"`
	ScoreThreshold       param.Field[float64]                                        `json:"score_threshold"`
	Source               param.Field[string]                                         `json:"source"`
	SourceParams         param.Field[NamespaceInstanceNewParamsSourceParams]         `json:"source_params"`
	// Interval between automatic syncs, in seconds. Allowed values: 900 (15min), 1800
	// (30min), 3600 (1h), 7200 (2h), 14400 (4h), 21600 (6h), 43200 (12h), 86400 (24h).
	SyncInterval param.Field[NamespaceInstanceNewParamsSyncInterval] `json:"sync_interval"`
	TokenID      param.Field[string]                                 `json:"token_id" format:"uuid"`
	Type         param.Field[NamespaceInstanceNewParamsType]         `json:"type"`
}

func (r NamespaceInstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsAISearchModel string

const (
	NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceNewParamsAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceNewParamsAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceNewParamsAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceNewParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceNewParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceNewParamsAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceNewParamsAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewParamsAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceNewParamsAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceNewParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceNewParamsAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceNewParamsAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceNewParamsAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceNewParamsAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceNewParamsAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceNewParamsAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceNewParamsAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceNewParamsAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceNewParamsAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceNewParamsAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceNewParamsAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceNewParamsAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceNewParamsAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceNewParamsAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceNewParamsAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceNewParamsAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceNewParamsAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceNewParamsAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceNewParamsAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceNewParamsAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceNewParamsAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceNewParamsAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceNewParamsAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceNewParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceNewParamsAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceNewParamsAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceNewParamsAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewParamsAISearchModelCerebrasGptOSs120b                    NamespaceInstanceNewParamsAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceNewParamsAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceNewParamsAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceNewParamsAISearchModelGrokGrok4                             NamespaceInstanceNewParamsAISearchModel = "grok/grok-4"
	NamespaceInstanceNewParamsAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceNewParamsAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceNewParamsAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceNewParamsAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceNewParamsAISearchModelOpenAIGpt5                            NamespaceInstanceNewParamsAISearchModel = "openai/gpt-5"
	NamespaceInstanceNewParamsAISearchModelOpenAIGpt5Mini                        NamespaceInstanceNewParamsAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceNewParamsAISearchModelOpenAIGpt5Nano                        NamespaceInstanceNewParamsAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceNewParamsAISearchModelEmpty                                 NamespaceInstanceNewParamsAISearchModel = ""
)

func (r NamespaceInstanceNewParamsAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceNewParamsAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceNewParamsAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceNewParamsAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceNewParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceNewParamsAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceNewParamsAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceNewParamsAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceNewParamsAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceNewParamsAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceNewParamsAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceNewParamsAISearchModelAnthropicClaudeOpus4, NamespaceInstanceNewParamsAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceNewParamsAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceNewParamsAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceNewParamsAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceNewParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceNewParamsAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceNewParamsAISearchModelCerebrasGptOSs120b, NamespaceInstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceNewParamsAISearchModelGrokGrok4, NamespaceInstanceNewParamsAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceNewParamsAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceNewParamsAISearchModelOpenAIGpt5, NamespaceInstanceNewParamsAISearchModelOpenAIGpt5Mini, NamespaceInstanceNewParamsAISearchModelOpenAIGpt5Nano, NamespaceInstanceNewParamsAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceNewParamsEmbeddingModel string

const (
	NamespaceInstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceNewParamsEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceNewParamsEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceNewParamsEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceNewParamsEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceNewParamsEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceNewParamsEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceNewParamsEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceNewParamsEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceNewParamsEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceNewParamsEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceNewParamsEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceNewParamsEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceNewParamsEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceNewParamsEmbeddingModelEmpty                                 NamespaceInstanceNewParamsEmbeddingModel = ""
)

func (r NamespaceInstanceNewParamsEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceNewParamsEmbeddingModelCfBaaiBgeM3, NamespaceInstanceNewParamsEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceNewParamsEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceNewParamsEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceNewParamsEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceNewParamsEmbeddingModelEmpty:
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
	Enabled                 param.Field[bool]                                                                  `json:"enabled"`
	Mcp                     param.Field[NamespaceInstanceNewParamsPublicEndpointParamsMcp]                     `json:"mcp"`
	RateLimit               param.Field[NamespaceInstanceNewParamsPublicEndpointParamsRateLimit]               `json:"rate_limit"`
	SearchEndpoint          param.Field[NamespaceInstanceNewParamsPublicEndpointParamsSearchEndpoint]          `json:"search_endpoint"`
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

type NamespaceInstanceNewParamsRerankingModel string

const (
	NamespaceInstanceNewParamsRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceNewParamsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceNewParamsRerankingModelEmpty                 NamespaceInstanceNewParamsRerankingModel = ""
)

func (r NamespaceInstanceNewParamsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceNewParamsRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy param.Field[[]NamespaceInstanceNewParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceNewParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r NamespaceInstanceNewParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceNewParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceNewParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceNewParamsRewriteModel string

const (
	NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceNewParamsRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceNewParamsRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceNewParamsRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceNewParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceNewParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceNewParamsRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceNewParamsRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewParamsRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceNewParamsRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceNewParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceNewParamsRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceNewParamsRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceNewParamsRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceNewParamsRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceNewParamsRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceNewParamsRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceNewParamsRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceNewParamsRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceNewParamsRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceNewParamsRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceNewParamsRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceNewParamsRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceNewParamsRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceNewParamsRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceNewParamsRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceNewParamsRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceNewParamsRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceNewParamsRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceNewParamsRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceNewParamsRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceNewParamsRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceNewParamsRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceNewParamsRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceNewParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceNewParamsRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceNewParamsRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceNewParamsRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceNewParamsRewriteModelCerebrasGptOSs120b                    NamespaceInstanceNewParamsRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceNewParamsRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceNewParamsRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceNewParamsRewriteModelGrokGrok4                             NamespaceInstanceNewParamsRewriteModel = "grok/grok-4"
	NamespaceInstanceNewParamsRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceNewParamsRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceNewParamsRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceNewParamsRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceNewParamsRewriteModelOpenAIGpt5                            NamespaceInstanceNewParamsRewriteModel = "openai/gpt-5"
	NamespaceInstanceNewParamsRewriteModelOpenAIGpt5Mini                        NamespaceInstanceNewParamsRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceNewParamsRewriteModelOpenAIGpt5Nano                        NamespaceInstanceNewParamsRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceNewParamsRewriteModelEmpty                                 NamespaceInstanceNewParamsRewriteModel = ""
)

func (r NamespaceInstanceNewParamsRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceNewParamsRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceNewParamsRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceNewParamsRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceNewParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceNewParamsRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceNewParamsRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceNewParamsRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceNewParamsRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceNewParamsRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceNewParamsRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceNewParamsRewriteModelAnthropicClaudeOpus4, NamespaceInstanceNewParamsRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceNewParamsRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceNewParamsRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceNewParamsRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceNewParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceNewParamsRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceNewParamsRewriteModelCerebrasGptOSs120b, NamespaceInstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceNewParamsRewriteModelGrokGrok4, NamespaceInstanceNewParamsRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceNewParamsRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceNewParamsRewriteModelOpenAIGpt5, NamespaceInstanceNewParamsRewriteModelOpenAIGpt5Mini, NamespaceInstanceNewParamsRewriteModelOpenAIGpt5Nano, NamespaceInstanceNewParamsRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   param.Field[[]string]                                         `json:"include_items"`
	Prefix         param.Field[string]                                           `json:"prefix"`
	R2Jurisdiction param.Field[string]                                           `json:"r2_jurisdiction"`
	WebCrawler     param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r NamespaceInstanceNewParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawler struct {
	CrawlOptions param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptions] `json:"crawl_options"`
	ParseOptions param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions] `json:"parse_options"`
	ParseType    param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType]    `json:"parse_type"`
	StoreOptions param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerStoreOptions] `json:"store_options"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptions struct {
	Depth                param.Field[float64]                                                            `json:"depth"`
	IncludeExternalLinks param.Field[bool]                                                               `json:"include_external_links"`
	IncludeSubdomains    param.Field[bool]                                                               `json:"include_subdomains"`
	MaxAge               param.Field[float64]                                                            `json:"max_age"`
	Source               param.Field[NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource] `json:"source"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceNewParamsSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector param.Field[[]NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
	IncludeHeaders  param.Field[map[string]string]                                                             `json:"include_headers"`
	IncludeImages   param.Field[bool]                                                                          `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceNewParamsSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceNewParamsSourceParamsWebCrawlerStoreOptions struct {
	StorageID      param.Field[string]      `json:"storage_id" api:"required"`
	R2Jurisdiction param.Field[string]      `json:"r2_jurisdiction"`
	StorageType    param.Field[r2.Provider] `json:"storage_type"`
}

func (r NamespaceInstanceNewParamsSourceParamsWebCrawlerStoreOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	AccountID      param.Field[string]                                        `path:"account_id" api:"required"`
	AIGatewayID    param.Field[string]                                        `json:"ai_gateway_id"`
	AISearchModel  param.Field[NamespaceInstanceUpdateParamsAISearchModel]    `json:"ai_search_model"`
	Cache          param.Field[bool]                                          `json:"cache"`
	CacheThreshold param.Field[NamespaceInstanceUpdateParamsCacheThreshold]   `json:"cache_threshold"`
	Chunk          param.Field[bool]                                          `json:"chunk"`
	ChunkOverlap   param.Field[int64]                                         `json:"chunk_overlap"`
	ChunkSize      param.Field[int64]                                         `json:"chunk_size"`
	CustomMetadata param.Field[[]NamespaceInstanceUpdateParamsCustomMetadata] `json:"custom_metadata"`
	EmbeddingModel param.Field[NamespaceInstanceUpdateParamsEmbeddingModel]   `json:"embedding_model"`
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
	RerankingModel       param.Field[NamespaceInstanceUpdateParamsRerankingModel]       `json:"reranking_model"`
	RetrievalOptions     param.Field[NamespaceInstanceUpdateParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel         param.Field[NamespaceInstanceUpdateParamsRewriteModel]         `json:"rewrite_model"`
	RewriteQuery         param.Field[bool]                                              `json:"rewrite_query"`
	ScoreThreshold       param.Field[float64]                                           `json:"score_threshold"`
	SourceParams         param.Field[NamespaceInstanceUpdateParamsSourceParams]         `json:"source_params"`
	Summarization        param.Field[bool]                                              `json:"summarization"`
	SummarizationModel   param.Field[NamespaceInstanceUpdateParamsSummarizationModel]   `json:"summarization_model"`
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

type NamespaceInstanceUpdateParamsAISearchModel string

const (
	NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceUpdateParamsAISearchModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceUpdateParamsAISearchModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceUpdateParamsAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsAISearchModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceUpdateParamsAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceUpdateParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceUpdateParamsAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceUpdateParamsAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceUpdateParamsAISearchModelCfGoogleGemma3_12bIt                  NamespaceInstanceUpdateParamsAISearchModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceUpdateParamsAISearchModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceUpdateParamsAISearchModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2_5                  NamespaceInstanceUpdateParamsAISearchModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceUpdateParamsAISearchModelAnthropicClaude3_7Sonnet              NamespaceInstanceUpdateParamsAISearchModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceUpdateParamsAISearchModelAnthropicClaudeSonnet4                NamespaceInstanceUpdateParamsAISearchModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceUpdateParamsAISearchModelAnthropicClaudeOpus4                  NamespaceInstanceUpdateParamsAISearchModel = "anthropic/claude-opus-4"
	NamespaceInstanceUpdateParamsAISearchModelAnthropicClaude3_5Haiku               NamespaceInstanceUpdateParamsAISearchModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceUpdateParamsAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceUpdateParamsAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama3_3_70b                  NamespaceInstanceUpdateParamsAISearchModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceUpdateParamsAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceUpdateParamsAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsAISearchModelCerebrasGptOSs120b                    NamespaceInstanceUpdateParamsAISearchModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceUpdateParamsAISearchModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceUpdateParamsAISearchModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceUpdateParamsAISearchModelGrokGrok4                             NamespaceInstanceUpdateParamsAISearchModel = "grok/grok-4"
	NamespaceInstanceUpdateParamsAISearchModelGroqLlama3_3_70bVersatile             NamespaceInstanceUpdateParamsAISearchModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceUpdateParamsAISearchModelGroqLlama3_1_8bInstant                NamespaceInstanceUpdateParamsAISearchModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5                            NamespaceInstanceUpdateParamsAISearchModel = "openai/gpt-5"
	NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5Mini                        NamespaceInstanceUpdateParamsAISearchModel = "openai/gpt-5-mini"
	NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5Nano                        NamespaceInstanceUpdateParamsAISearchModel = "openai/gpt-5-nano"
	NamespaceInstanceUpdateParamsAISearchModelEmpty                                 NamespaceInstanceUpdateParamsAISearchModel = ""
)

func (r NamespaceInstanceUpdateParamsAISearchModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceUpdateParamsAISearchModelCfZaiOrgGlm4_7Flash, NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceUpdateParamsAISearchModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsAISearchModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceUpdateParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceUpdateParamsAISearchModelCfGoogleGemma3_12bIt, NamespaceInstanceUpdateParamsAISearchModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2_5, NamespaceInstanceUpdateParamsAISearchModelAnthropicClaude3_7Sonnet, NamespaceInstanceUpdateParamsAISearchModelAnthropicClaudeSonnet4, NamespaceInstanceUpdateParamsAISearchModelAnthropicClaudeOpus4, NamespaceInstanceUpdateParamsAISearchModelAnthropicClaude3_5Haiku, NamespaceInstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama3_3_70b, NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceUpdateParamsAISearchModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsAISearchModelCerebrasGptOSs120b, NamespaceInstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceUpdateParamsAISearchModelGrokGrok4, NamespaceInstanceUpdateParamsAISearchModelGroqLlama3_3_70bVersatile, NamespaceInstanceUpdateParamsAISearchModelGroqLlama3_1_8bInstant, NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5, NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5Mini, NamespaceInstanceUpdateParamsAISearchModelOpenAIGpt5Nano, NamespaceInstanceUpdateParamsAISearchModelEmpty:
		return true
	}
	return false
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

type NamespaceInstanceUpdateParamsEmbeddingModel string

const (
	NamespaceInstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b              NamespaceInstanceUpdateParamsEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	NamespaceInstanceUpdateParamsEmbeddingModelCfBaaiBgeM3                           NamespaceInstanceUpdateParamsEmbeddingModel = "@cf/baai/bge-m3"
	NamespaceInstanceUpdateParamsEmbeddingModelCfBaaiBgeLargeEnV1_5                  NamespaceInstanceUpdateParamsEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	NamespaceInstanceUpdateParamsEmbeddingModelCfGoogleEmbeddinggemma300m            NamespaceInstanceUpdateParamsEmbeddingModel = "@cf/google/embeddinggemma-300m"
	NamespaceInstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001      NamespaceInstanceUpdateParamsEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	NamespaceInstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview NamespaceInstanceUpdateParamsEmbeddingModel = "google-ai-studio/gemini-embedding-2-preview"
	NamespaceInstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Small             NamespaceInstanceUpdateParamsEmbeddingModel = "openai/text-embedding-3-small"
	NamespaceInstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Large             NamespaceInstanceUpdateParamsEmbeddingModel = "openai/text-embedding-3-large"
	NamespaceInstanceUpdateParamsEmbeddingModelEmpty                                 NamespaceInstanceUpdateParamsEmbeddingModel = ""
)

func (r NamespaceInstanceUpdateParamsEmbeddingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b, NamespaceInstanceUpdateParamsEmbeddingModelCfBaaiBgeM3, NamespaceInstanceUpdateParamsEmbeddingModelCfBaaiBgeLargeEnV1_5, NamespaceInstanceUpdateParamsEmbeddingModelCfGoogleEmbeddinggemma300m, NamespaceInstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001, NamespaceInstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding2Preview, NamespaceInstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Small, NamespaceInstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Large, NamespaceInstanceUpdateParamsEmbeddingModelEmpty:
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
	Enabled                 param.Field[bool]                                                                     `json:"enabled"`
	Mcp                     param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsMcp]                     `json:"mcp"`
	RateLimit               param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsRateLimit]               `json:"rate_limit"`
	SearchEndpoint          param.Field[NamespaceInstanceUpdateParamsPublicEndpointParamsSearchEndpoint]          `json:"search_endpoint"`
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

type NamespaceInstanceUpdateParamsRerankingModel string

const (
	NamespaceInstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceUpdateParamsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceUpdateParamsRerankingModelEmpty                 NamespaceInstanceUpdateParamsRerankingModel = ""
)

func (r NamespaceInstanceUpdateParamsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceUpdateParamsRerankingModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsRetrievalOptions struct {
	// Metadata fields to boost search results by. Each entry specifies a metadata
	// field and an optional direction. Direction defaults to 'asc' for numeric fields
	// and 'exists' for text/boolean fields. Fields must match 'timestamp' or a defined
	// custom_metadata field.
	BoostBy param.Field[[]NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy] `json:"boost_by"`
	// Controls which documents are candidates for BM25 scoring. 'and' restricts
	// candidates to documents containing all query terms; 'or' includes any document
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
	KeywordMatchMode param.Field[NamespaceInstanceUpdateParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r NamespaceInstanceUpdateParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy struct {
	// Metadata field name to boost by. Use 'timestamp' for document freshness, or any
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceUpdateParamsRetrievalOptionsBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceUpdateParamsRetrievalOptionsBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceUpdateParamsRewriteModel string

const (
	NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceUpdateParamsRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceUpdateParamsRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceUpdateParamsRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceUpdateParamsRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceUpdateParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceUpdateParamsRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceUpdateParamsRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceUpdateParamsRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceUpdateParamsRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceUpdateParamsRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceUpdateParamsRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceUpdateParamsRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceUpdateParamsRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceUpdateParamsRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceUpdateParamsRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceUpdateParamsRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceUpdateParamsRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceUpdateParamsRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceUpdateParamsRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceUpdateParamsRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceUpdateParamsRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceUpdateParamsRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceUpdateParamsRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceUpdateParamsRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceUpdateParamsRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsRewriteModelCerebrasGptOSs120b                    NamespaceInstanceUpdateParamsRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceUpdateParamsRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceUpdateParamsRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceUpdateParamsRewriteModelGrokGrok4                             NamespaceInstanceUpdateParamsRewriteModel = "grok/grok-4"
	NamespaceInstanceUpdateParamsRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceUpdateParamsRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceUpdateParamsRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceUpdateParamsRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5                            NamespaceInstanceUpdateParamsRewriteModel = "openai/gpt-5"
	NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5Mini                        NamespaceInstanceUpdateParamsRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5Nano                        NamespaceInstanceUpdateParamsRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceUpdateParamsRewriteModelEmpty                                 NamespaceInstanceUpdateParamsRewriteModel = ""
)

func (r NamespaceInstanceUpdateParamsRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceUpdateParamsRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceUpdateParamsRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceUpdateParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceUpdateParamsRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceUpdateParamsRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceUpdateParamsRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceUpdateParamsRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceUpdateParamsRewriteModelAnthropicClaudeOpus4, NamespaceInstanceUpdateParamsRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceUpdateParamsRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsRewriteModelCerebrasGptOSs120b, NamespaceInstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceUpdateParamsRewriteModelGrokGrok4, NamespaceInstanceUpdateParamsRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceUpdateParamsRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5, NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5Mini, NamespaceInstanceUpdateParamsRewriteModelOpenAIGpt5Nano, NamespaceInstanceUpdateParamsRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   param.Field[[]string]                                            `json:"include_items"`
	Prefix         param.Field[string]                                              `json:"prefix"`
	R2Jurisdiction param.Field[string]                                              `json:"r2_jurisdiction"`
	WebCrawler     param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r NamespaceInstanceUpdateParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawler struct {
	CrawlOptions param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptions] `json:"crawl_options"`
	ParseOptions param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions] `json:"parse_options"`
	ParseType    param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType]    `json:"parse_type"`
	StoreOptions param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerStoreOptions] `json:"store_options"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptions struct {
	Depth                param.Field[float64]                                                               `json:"depth"`
	IncludeExternalLinks param.Field[bool]                                                                  `json:"include_external_links"`
	IncludeSubdomains    param.Field[bool]                                                                  `json:"include_subdomains"`
	MaxAge               param.Field[float64]                                                               `json:"max_age"`
	Source               param.Field[NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource] `json:"source"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource string

const (
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceAll      NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource = "all"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceSitemaps NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource = "sitemaps"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceLinks    NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource = "links"
)

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSource) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceAll, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceSitemaps, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerCrawlOptionsSourceLinks:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptions struct {
	// List of path-to-selector mappings for extracting specific content from crawled
	// pages. Each entry pairs a URL glob pattern with a CSS selector. The first
	// matching path wins. Only the matched HTML fragment is stored and indexed.
	ContentSelector param.Field[[]NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector] `json:"content_selector"`
	IncludeHeaders  param.Field[map[string]string]                                                                `json:"include_headers"`
	IncludeImages   param.Field[bool]                                                                             `json:"include_images"`
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
	// CSS selector to extract content from pages matching the path pattern. Supports
	// standard CSS selectors including class, ID, element, and attribute selectors.
	Selector param.Field[string] `json:"selector" api:"required"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseOptionsContentSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType string

const (
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType = "sitemap"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeFeedRss NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType = "feed-rss"
	NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeCrawl   NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType = "crawl"
)

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeFeedRss, NamespaceInstanceUpdateParamsSourceParamsWebCrawlerParseTypeCrawl:
		return true
	}
	return false
}

type NamespaceInstanceUpdateParamsSourceParamsWebCrawlerStoreOptions struct {
	StorageID      param.Field[string]      `json:"storage_id" api:"required"`
	R2Jurisdiction param.Field[string]      `json:"r2_jurisdiction"`
	StorageType    param.Field[r2.Provider] `json:"storage_type"`
}

func (r NamespaceInstanceUpdateParamsSourceParamsWebCrawlerStoreOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceUpdateParamsSummarizationModel string

const (
	NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceUpdateParamsSummarizationModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceUpdateParamsSummarizationModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceUpdateParamsSummarizationModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsSummarizationModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceUpdateParamsSummarizationModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceUpdateParamsSummarizationModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceUpdateParamsSummarizationModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceUpdateParamsSummarizationModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceUpdateParamsSummarizationModelCfGoogleGemma3_12bIt                  NamespaceInstanceUpdateParamsSummarizationModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceUpdateParamsSummarizationModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceUpdateParamsSummarizationModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2_5                  NamespaceInstanceUpdateParamsSummarizationModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaude3_7Sonnet              NamespaceInstanceUpdateParamsSummarizationModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaudeSonnet4                NamespaceInstanceUpdateParamsSummarizationModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaudeOpus4                  NamespaceInstanceUpdateParamsSummarizationModel = "anthropic/claude-opus-4"
	NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaude3_5Haiku               NamespaceInstanceUpdateParamsSummarizationModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama3_3_70b                  NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceUpdateParamsSummarizationModelCerebrasGptOSs120b                    NamespaceInstanceUpdateParamsSummarizationModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceUpdateParamsSummarizationModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceUpdateParamsSummarizationModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceUpdateParamsSummarizationModelGrokGrok4                             NamespaceInstanceUpdateParamsSummarizationModel = "grok/grok-4"
	NamespaceInstanceUpdateParamsSummarizationModelGroqLlama3_3_70bVersatile             NamespaceInstanceUpdateParamsSummarizationModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceUpdateParamsSummarizationModelGroqLlama3_1_8bInstant                NamespaceInstanceUpdateParamsSummarizationModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5                            NamespaceInstanceUpdateParamsSummarizationModel = "openai/gpt-5"
	NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5Mini                        NamespaceInstanceUpdateParamsSummarizationModel = "openai/gpt-5-mini"
	NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5Nano                        NamespaceInstanceUpdateParamsSummarizationModel = "openai/gpt-5-nano"
	NamespaceInstanceUpdateParamsSummarizationModelEmpty                                 NamespaceInstanceUpdateParamsSummarizationModel = ""
)

func (r NamespaceInstanceUpdateParamsSummarizationModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceUpdateParamsSummarizationModelCfZaiOrgGlm4_7Flash, NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceUpdateParamsSummarizationModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsSummarizationModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceUpdateParamsSummarizationModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceUpdateParamsSummarizationModelCfGoogleGemma3_12bIt, NamespaceInstanceUpdateParamsSummarizationModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2_5, NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaude3_7Sonnet, NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaudeSonnet4, NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaudeOpus4, NamespaceInstanceUpdateParamsSummarizationModelAnthropicClaude3_5Haiku, NamespaceInstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama3_3_70b, NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceUpdateParamsSummarizationModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceUpdateParamsSummarizationModelCerebrasGptOSs120b, NamespaceInstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceUpdateParamsSummarizationModelGrokGrok4, NamespaceInstanceUpdateParamsSummarizationModelGroqLlama3_3_70bVersatile, NamespaceInstanceUpdateParamsSummarizationModelGroqLlama3_1_8bInstant, NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5, NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5Mini, NamespaceInstanceUpdateParamsSummarizationModelOpenAIGpt5Nano, NamespaceInstanceUpdateParamsSummarizationModelEmpty:
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
	Model           param.Field[NamespaceInstanceChatCompletionsParamsModel]           `json:"model"`
	Stream          param.Field[bool]                                                  `json:"stream"`
}

func (r NamespaceInstanceChatCompletionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsMessage struct {
	Content     param.Field[string]                                             `json:"content" api:"required"`
	Role        param.Field[NamespaceInstanceChatCompletionsParamsMessagesRole] `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                          `json:"-,extras"`
}

func (r NamespaceInstanceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
	Enabled       param.Field[bool]                                                                   `json:"enabled"`
	Model         param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                                                 `json:"rewrite_prompt"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4                             NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty                                 NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = ""
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, NamespaceInstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                                                `json:"enabled"`
	MatchThreshold param.Field[float64]                                                             `json:"match_threshold"`
	Model          param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModel string

const (
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModelEmpty                 NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModel = ""
)

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceChatCompletionsParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
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
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceChatCompletionsParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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

type NamespaceInstanceChatCompletionsParamsModel string

const (
	NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceChatCompletionsParamsModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceChatCompletionsParamsModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceChatCompletionsParamsModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceChatCompletionsParamsModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceChatCompletionsParamsModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceChatCompletionsParamsModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceChatCompletionsParamsModelCfGoogleGemma3_12bIt                  NamespaceInstanceChatCompletionsParamsModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceChatCompletionsParamsModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceChatCompletionsParamsModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceChatCompletionsParamsModelCfMoonshotaiKimiK2_5                  NamespaceInstanceChatCompletionsParamsModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceChatCompletionsParamsModelAnthropicClaude3_7Sonnet              NamespaceInstanceChatCompletionsParamsModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceChatCompletionsParamsModelAnthropicClaudeSonnet4                NamespaceInstanceChatCompletionsParamsModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceChatCompletionsParamsModelAnthropicClaudeOpus4                  NamespaceInstanceChatCompletionsParamsModel = "anthropic/claude-opus-4"
	NamespaceInstanceChatCompletionsParamsModelAnthropicClaude3_5Haiku               NamespaceInstanceChatCompletionsParamsModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceChatCompletionsParamsModelCerebrasLlama3_3_70b                  NamespaceInstanceChatCompletionsParamsModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceChatCompletionsParamsModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceChatCompletionsParamsModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceChatCompletionsParamsModelCerebrasGptOSs120b                    NamespaceInstanceChatCompletionsParamsModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceChatCompletionsParamsModelGrokGrok4                             NamespaceInstanceChatCompletionsParamsModel = "grok/grok-4"
	NamespaceInstanceChatCompletionsParamsModelGroqLlama3_3_70bVersatile             NamespaceInstanceChatCompletionsParamsModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceChatCompletionsParamsModelGroqLlama3_1_8bInstant                NamespaceInstanceChatCompletionsParamsModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5                            NamespaceInstanceChatCompletionsParamsModel = "openai/gpt-5"
	NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5Mini                        NamespaceInstanceChatCompletionsParamsModel = "openai/gpt-5-mini"
	NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5Nano                        NamespaceInstanceChatCompletionsParamsModel = "openai/gpt-5-nano"
	NamespaceInstanceChatCompletionsParamsModelEmpty                                 NamespaceInstanceChatCompletionsParamsModel = ""
)

func (r NamespaceInstanceChatCompletionsParamsModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash, NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceChatCompletionsParamsModelCfGoogleGemma3_12bIt, NamespaceInstanceChatCompletionsParamsModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceChatCompletionsParamsModelCfMoonshotaiKimiK2_5, NamespaceInstanceChatCompletionsParamsModelAnthropicClaude3_7Sonnet, NamespaceInstanceChatCompletionsParamsModelAnthropicClaudeSonnet4, NamespaceInstanceChatCompletionsParamsModelAnthropicClaudeOpus4, NamespaceInstanceChatCompletionsParamsModelAnthropicClaude3_5Haiku, NamespaceInstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceChatCompletionsParamsModelCerebrasLlama3_3_70b, NamespaceInstanceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceChatCompletionsParamsModelCerebrasGptOSs120b, NamespaceInstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceChatCompletionsParamsModelGrokGrok4, NamespaceInstanceChatCompletionsParamsModelGroqLlama3_3_70bVersatile, NamespaceInstanceChatCompletionsParamsModelGroqLlama3_1_8bInstant, NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5, NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5Mini, NamespaceInstanceChatCompletionsParamsModelOpenAIGpt5Nano, NamespaceInstanceChatCompletionsParamsModelEmpty:
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
	Messages        param.Field[[]NamespaceInstanceSearchParamsMessage]       `json:"messages"`
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
	Enabled       param.Field[bool]                                                          `json:"enabled"`
	Model         param.Field[NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                                        `json:"rewrite_prompt"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel string

const (
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt               NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-4-26b-a4b-it"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5                  NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2.5"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4                             NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelEmpty                                 NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel = ""
)

func (r NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma4_26bA4bIt, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2_5, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, NamespaceInstanceSearchParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type NamespaceInstanceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                                       `json:"enabled"`
	MatchThreshold param.Field[float64]                                                    `json:"match_threshold"`
	Model          param.Field[NamespaceInstanceSearchParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NamespaceInstanceSearchParamsAISearchOptionsRerankingModel string

const (
	NamespaceInstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase NamespaceInstanceSearchParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	NamespaceInstanceSearchParamsAISearchOptionsRerankingModelEmpty                 NamespaceInstanceSearchParamsAISearchOptionsRerankingModel = ""
)

func (r NamespaceInstanceSearchParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case NamespaceInstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, NamespaceInstanceSearchParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
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
	// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	// custom_metadata field. Numeric and datetime fields support asc/desc directions;
	// text/boolean fields support exists/not_exists.
	Field param.Field[string] `json:"field" api:"required"`
	// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
	// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
	// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
	// for numeric/datetime fields, 'exists' for text/boolean fields.
	Direction param.Field[NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostByDirection] `json:"direction"`
}

func (r NamespaceInstanceSearchParamsAISearchOptionsRetrievalBoostBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Boost direction. 'desc' = higher values rank higher (e.g. newer timestamps).
// 'asc' = lower values rank higher. 'exists' = boost chunks that have the field.
// 'not_exists' = boost chunks that lack the field. Optional ��� defaults to 'asc'
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
// containing at least one term, ranked by BM25 relevance. Defaults to 'and'.
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
	Content     param.Field[string]                                    `json:"content" api:"required"`
	Role        param.Field[NamespaceInstanceSearchParamsMessagesRole] `json:"role" api:"required"`
	ExtraFields map[string]interface{}                                 `json:"-,extras"`
}

func (r NamespaceInstanceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
