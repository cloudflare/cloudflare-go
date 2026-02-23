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

// InstanceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options []option.RequestOption
	Items   *InstanceItemService
	Jobs    *InstanceJobService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r *InstanceService) {
	r = &InstanceService{}
	r.Options = opts
	r.Items = NewInstanceItemService(opts...)
	r.Jobs = NewInstanceJobService(opts...)
	return
}

// Create new instances.
func (r *InstanceService) New(ctx context.Context, params InstanceNewParams, opts ...option.RequestOption) (res *InstanceNewResponse, err error) {
	var env InstanceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update instances.
func (r *InstanceService) Update(ctx context.Context, id string, params InstanceUpdateParams, opts ...option.RequestOption) (res *InstanceUpdateResponse, err error) {
	var env InstanceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List instances.
func (r *InstanceService) List(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InstanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances", params.AccountID)
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
func (r *InstanceService) ListAutoPaging(ctx context.Context, params InstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete instances.
func (r *InstanceService) Delete(ctx context.Context, id string, body InstanceDeleteParams, opts ...option.RequestOption) (res *InstanceDeleteResponse, err error) {
	var env InstanceDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Chat Completions
func (r *InstanceService) ChatCompletions(ctx context.Context, id string, params InstanceChatCompletionsParams, opts ...option.RequestOption) (res *InstanceChatCompletionsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/chat/completions", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Read instances.
func (r *InstanceService) Read(ctx context.Context, id string, query InstanceReadParams, opts ...option.RequestOption) (res *InstanceReadResponse, err error) {
	var env InstanceReadResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Search
func (r *InstanceService) Search(ctx context.Context, id string, params InstanceSearchParams, opts ...option.RequestOption) (res *InstanceSearchResponse, err error) {
	var env InstanceSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/search", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Stats
func (r *InstanceService) Stats(ctx context.Context, id string, query InstanceStatsParams, opts ...option.RequestOption) (res *InstanceStatsResponse, err error) {
	var env InstanceStatsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-search/instances/%s/stats", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceNewResponse struct {
	// Use your AI Search ID.
	ID                   string                                  `json:"id,required"`
	CreatedAt            time.Time                               `json:"created_at,required" format:"date-time"`
	ModifiedAt           time.Time                               `json:"modified_at,required" format:"date-time"`
	Source               string                                  `json:"source,required"`
	Type                 InstanceNewResponseType                 `json:"type,required"`
	VectorizeName        string                                  `json:"vectorize_name,required"`
	AIGatewayID          string                                  `json:"ai_gateway_id,nullable"`
	AISearchModel        InstanceNewResponseAISearchModel        `json:"ai_search_model"`
	Cache                bool                                    `json:"cache"`
	CacheThreshold       InstanceNewResponseCacheThreshold       `json:"cache_threshold"`
	ChunkOverlap         int64                                   `json:"chunk_overlap"`
	ChunkSize            int64                                   `json:"chunk_size"`
	CreatedBy            string                                  `json:"created_by,nullable"`
	CustomMetadata       []InstanceNewResponseCustomMetadata     `json:"custom_metadata"`
	EmbeddingModel       InstanceNewResponseEmbeddingModel       `json:"embedding_model"`
	Enable               bool                                    `json:"enable"`
	FusionMethod         InstanceNewResponseFusionMethod         `json:"fusion_method"`
	HybridSearchEnabled  bool                                    `json:"hybrid_search_enabled"`
	LastActivity         time.Time                               `json:"last_activity,nullable" format:"date-time"`
	MaxNumResults        int64                                   `json:"max_num_results"`
	Metadata             InstanceNewResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                  `json:"modified_by,nullable"`
	Paused               bool                                    `json:"paused"`
	PublicEndpointID     string                                  `json:"public_endpoint_id,nullable"`
	PublicEndpointParams InstanceNewResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                    `json:"reranking"`
	RerankingModel       InstanceNewResponseRerankingModel       `json:"reranking_model"`
	RetrievalOptions     InstanceNewResponseRetrievalOptions     `json:"retrieval_options,nullable"`
	RewriteModel         InstanceNewResponseRewriteModel         `json:"rewrite_model"`
	RewriteQuery         bool                                    `json:"rewrite_query"`
	ScoreThreshold       float64                                 `json:"score_threshold"`
	SourceParams         InstanceNewResponseSourceParams         `json:"source_params,nullable"`
	Status               string                                  `json:"status"`
	TokenID              string                                  `json:"token_id" format:"uuid"`
	JSON                 instanceNewResponseJSON                 `json:"-"`
}

// instanceNewResponseJSON contains the JSON metadata for the struct
// [InstanceNewResponse]
type instanceNewResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	Source               apijson.Field
	Type                 apijson.Field
	VectorizeName        apijson.Field
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
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	TokenID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseType string

const (
	InstanceNewResponseTypeR2         InstanceNewResponseType = "r2"
	InstanceNewResponseTypeWebCrawler InstanceNewResponseType = "web-crawler"
)

func (r InstanceNewResponseType) IsKnown() bool {
	switch r {
	case InstanceNewResponseTypeR2, InstanceNewResponseTypeWebCrawler:
		return true
	}
	return false
}

type InstanceNewResponseAISearchModel string

const (
	InstanceNewResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceNewResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceNewResponseAISearchModelCfZaiOrgGlm4_7Flash                   InstanceNewResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceNewResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceNewResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceNewResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceNewResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceNewResponseAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceNewResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceNewResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceNewResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceNewResponseAISearchModelCfMoonshotaiKimiK2Instruct            InstanceNewResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceNewResponseAISearchModelCfGoogleGemma3_12bIt                  InstanceNewResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceNewResponseAISearchModelAnthropicClaude3_7Sonnet              InstanceNewResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceNewResponseAISearchModelAnthropicClaudeSonnet4                InstanceNewResponseAISearchModel = "anthropic/claude-sonnet-4"
	InstanceNewResponseAISearchModelAnthropicClaudeOpus4                  InstanceNewResponseAISearchModel = "anthropic/claude-opus-4"
	InstanceNewResponseAISearchModelAnthropicClaude3_5Haiku               InstanceNewResponseAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceNewResponseAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceNewResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceNewResponseAISearchModelCerebrasQwen3_235bA22bThinking        InstanceNewResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceNewResponseAISearchModelCerebrasLlama3_3_70b                  InstanceNewResponseAISearchModel = "cerebras/llama-3.3-70b"
	InstanceNewResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceNewResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceNewResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceNewResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceNewResponseAISearchModelCerebrasGptOSs120b                    InstanceNewResponseAISearchModel = "cerebras/gpt-oss-120b"
	InstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Flash          InstanceNewResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Pro            InstanceNewResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceNewResponseAISearchModelGrokGrok4                             InstanceNewResponseAISearchModel = "grok/grok-4"
	InstanceNewResponseAISearchModelGroqLlama3_3_70bVersatile             InstanceNewResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceNewResponseAISearchModelGroqLlama3_1_8bInstant                InstanceNewResponseAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceNewResponseAISearchModelOpenAIGpt5                            InstanceNewResponseAISearchModel = "openai/gpt-5"
	InstanceNewResponseAISearchModelOpenAIGpt5Mini                        InstanceNewResponseAISearchModel = "openai/gpt-5-mini"
	InstanceNewResponseAISearchModelOpenAIGpt5Nano                        InstanceNewResponseAISearchModel = "openai/gpt-5-nano"
	InstanceNewResponseAISearchModelEmpty                                 InstanceNewResponseAISearchModel = ""
)

func (r InstanceNewResponseAISearchModel) IsKnown() bool {
	switch r {
	case InstanceNewResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceNewResponseAISearchModelCfZaiOrgGlm4_7Flash, InstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceNewResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceNewResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceNewResponseAISearchModelCfQwenQwen3_30bA3bFp8, InstanceNewResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceNewResponseAISearchModelCfMoonshotaiKimiK2Instruct, InstanceNewResponseAISearchModelCfGoogleGemma3_12bIt, InstanceNewResponseAISearchModelAnthropicClaude3_7Sonnet, InstanceNewResponseAISearchModelAnthropicClaudeSonnet4, InstanceNewResponseAISearchModelAnthropicClaudeOpus4, InstanceNewResponseAISearchModelAnthropicClaude3_5Haiku, InstanceNewResponseAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceNewResponseAISearchModelCerebrasQwen3_235bA22bThinking, InstanceNewResponseAISearchModelCerebrasLlama3_3_70b, InstanceNewResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceNewResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceNewResponseAISearchModelCerebrasGptOSs120b, InstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Flash, InstanceNewResponseAISearchModelGoogleAIStudioGemini2_5Pro, InstanceNewResponseAISearchModelGrokGrok4, InstanceNewResponseAISearchModelGroqLlama3_3_70bVersatile, InstanceNewResponseAISearchModelGroqLlama3_1_8bInstant, InstanceNewResponseAISearchModelOpenAIGpt5, InstanceNewResponseAISearchModelOpenAIGpt5Mini, InstanceNewResponseAISearchModelOpenAIGpt5Nano, InstanceNewResponseAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceNewResponseCacheThreshold string

const (
	InstanceNewResponseCacheThresholdSuperStrictMatch InstanceNewResponseCacheThreshold = "super_strict_match"
	InstanceNewResponseCacheThresholdCloseEnough      InstanceNewResponseCacheThreshold = "close_enough"
	InstanceNewResponseCacheThresholdFlexibleFriend   InstanceNewResponseCacheThreshold = "flexible_friend"
	InstanceNewResponseCacheThresholdAnythingGoes     InstanceNewResponseCacheThreshold = "anything_goes"
)

func (r InstanceNewResponseCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceNewResponseCacheThresholdSuperStrictMatch, InstanceNewResponseCacheThresholdCloseEnough, InstanceNewResponseCacheThresholdFlexibleFriend, InstanceNewResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceNewResponseCustomMetadata struct {
	DataType  InstanceNewResponseCustomMetadataDataType `json:"data_type,required"`
	FieldName string                                    `json:"field_name,required"`
	JSON      instanceNewResponseCustomMetadataJSON     `json:"-"`
}

// instanceNewResponseCustomMetadataJSON contains the JSON metadata for the struct
// [InstanceNewResponseCustomMetadata]
type instanceNewResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseCustomMetadataDataType string

const (
	InstanceNewResponseCustomMetadataDataTypeText    InstanceNewResponseCustomMetadataDataType = "text"
	InstanceNewResponseCustomMetadataDataTypeNumber  InstanceNewResponseCustomMetadataDataType = "number"
	InstanceNewResponseCustomMetadataDataTypeBoolean InstanceNewResponseCustomMetadataDataType = "boolean"
)

func (r InstanceNewResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceNewResponseCustomMetadataDataTypeText, InstanceNewResponseCustomMetadataDataTypeNumber, InstanceNewResponseCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceNewResponseEmbeddingModel string

const (
	InstanceNewResponseEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceNewResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceNewResponseEmbeddingModelCfBaaiBgeM3                      InstanceNewResponseEmbeddingModel = "@cf/baai/bge-m3"
	InstanceNewResponseEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceNewResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceNewResponseEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceNewResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceNewResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceNewResponseEmbeddingModelOpenAITextEmbedding3Small        InstanceNewResponseEmbeddingModel = "openai/text-embedding-3-small"
	InstanceNewResponseEmbeddingModelOpenAITextEmbedding3Large        InstanceNewResponseEmbeddingModel = "openai/text-embedding-3-large"
	InstanceNewResponseEmbeddingModelEmpty                            InstanceNewResponseEmbeddingModel = ""
)

func (r InstanceNewResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceNewResponseEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceNewResponseEmbeddingModelCfBaaiBgeM3, InstanceNewResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceNewResponseEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceNewResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceNewResponseEmbeddingModelOpenAITextEmbedding3Small, InstanceNewResponseEmbeddingModelOpenAITextEmbedding3Large, InstanceNewResponseEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceNewResponseFusionMethod string

const (
	InstanceNewResponseFusionMethodMax InstanceNewResponseFusionMethod = "max"
	InstanceNewResponseFusionMethodRrf InstanceNewResponseFusionMethod = "rrf"
)

func (r InstanceNewResponseFusionMethod) IsKnown() bool {
	switch r {
	case InstanceNewResponseFusionMethodMax, InstanceNewResponseFusionMethodRrf:
		return true
	}
	return false
}

type InstanceNewResponseMetadata struct {
	CreatedFromAISearchWizard bool                            `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                          `json:"worker_domain"`
	JSON                      instanceNewResponseMetadataJSON `json:"-"`
}

// instanceNewResponseMetadataJSON contains the JSON metadata for the struct
// [InstanceNewResponseMetadata]
type instanceNewResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InstanceNewResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                       `json:"authorized_hosts"`
	ChatCompletionsEndpoint InstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	Enabled                 bool                                                           `json:"enabled"`
	Mcp                     InstanceNewResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceNewResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceNewResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    instanceNewResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceNewResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceNewResponsePublicEndpointParams]
type instanceNewResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstanceNewResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                               `json:"disabled"`
	JSON     instanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// instanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON contains the
// JSON metadata for the struct
// [InstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint]
type instanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                           `json:"disabled"`
	JSON     instanceNewResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// instanceNewResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [InstanceNewResponsePublicEndpointParamsMcp]
type instanceNewResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                     `json:"period_ms"`
	Requests  int64                                                     `json:"requests"`
	Technique InstanceNewResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      instanceNewResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// instanceNewResponsePublicEndpointParamsRateLimitJSON contains the JSON metadata
// for the struct [InstanceNewResponsePublicEndpointParamsRateLimit]
type instanceNewResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponsePublicEndpointParamsRateLimitTechnique string

const (
	InstanceNewResponsePublicEndpointParamsRateLimitTechniqueFixed   InstanceNewResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceNewResponsePublicEndpointParamsRateLimitTechniqueSliding InstanceNewResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceNewResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceNewResponsePublicEndpointParamsRateLimitTechniqueFixed, InstanceNewResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceNewResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                      `json:"disabled"`
	JSON     instanceNewResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// instanceNewResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct [InstanceNewResponsePublicEndpointParamsSearchEndpoint]
type instanceNewResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseRerankingModel string

const (
	InstanceNewResponseRerankingModelCfBaaiBgeRerankerBase InstanceNewResponseRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceNewResponseRerankingModelEmpty                 InstanceNewResponseRerankingModel = ""
)

func (r InstanceNewResponseRerankingModel) IsKnown() bool {
	switch r {
	case InstanceNewResponseRerankingModelCfBaaiBgeRerankerBase, InstanceNewResponseRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceNewResponseRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode InstanceNewResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceNewResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceNewResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceNewResponseRetrievalOptions]
type instanceNewResponseRetrievalOptionsJSON struct {
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceNewResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceNewResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceNewResponseRetrievalOptionsKeywordMatchModeExactMatch InstanceNewResponseRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceNewResponseRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceNewResponseRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceNewResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceNewResponseRetrievalOptionsKeywordMatchModeExactMatch, InstanceNewResponseRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceNewResponseRewriteModel string

const (
	InstanceNewResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceNewResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceNewResponseRewriteModelCfZaiOrgGlm4_7Flash                   InstanceNewResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceNewResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceNewResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceNewResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceNewResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceNewResponseRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceNewResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceNewResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceNewResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceNewResponseRewriteModelCfMoonshotaiKimiK2Instruct            InstanceNewResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceNewResponseRewriteModelCfGoogleGemma3_12bIt                  InstanceNewResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceNewResponseRewriteModelAnthropicClaude3_7Sonnet              InstanceNewResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceNewResponseRewriteModelAnthropicClaudeSonnet4                InstanceNewResponseRewriteModel = "anthropic/claude-sonnet-4"
	InstanceNewResponseRewriteModelAnthropicClaudeOpus4                  InstanceNewResponseRewriteModel = "anthropic/claude-opus-4"
	InstanceNewResponseRewriteModelAnthropicClaude3_5Haiku               InstanceNewResponseRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceNewResponseRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceNewResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceNewResponseRewriteModelCerebrasQwen3_235bA22bThinking        InstanceNewResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceNewResponseRewriteModelCerebrasLlama3_3_70b                  InstanceNewResponseRewriteModel = "cerebras/llama-3.3-70b"
	InstanceNewResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceNewResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceNewResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceNewResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceNewResponseRewriteModelCerebrasGptOSs120b                    InstanceNewResponseRewriteModel = "cerebras/gpt-oss-120b"
	InstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Flash          InstanceNewResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Pro            InstanceNewResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceNewResponseRewriteModelGrokGrok4                             InstanceNewResponseRewriteModel = "grok/grok-4"
	InstanceNewResponseRewriteModelGroqLlama3_3_70bVersatile             InstanceNewResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceNewResponseRewriteModelGroqLlama3_1_8bInstant                InstanceNewResponseRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceNewResponseRewriteModelOpenAIGpt5                            InstanceNewResponseRewriteModel = "openai/gpt-5"
	InstanceNewResponseRewriteModelOpenAIGpt5Mini                        InstanceNewResponseRewriteModel = "openai/gpt-5-mini"
	InstanceNewResponseRewriteModelOpenAIGpt5Nano                        InstanceNewResponseRewriteModel = "openai/gpt-5-nano"
	InstanceNewResponseRewriteModelEmpty                                 InstanceNewResponseRewriteModel = ""
)

func (r InstanceNewResponseRewriteModel) IsKnown() bool {
	switch r {
	case InstanceNewResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceNewResponseRewriteModelCfZaiOrgGlm4_7Flash, InstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceNewResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceNewResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceNewResponseRewriteModelCfQwenQwen3_30bA3bFp8, InstanceNewResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceNewResponseRewriteModelCfMoonshotaiKimiK2Instruct, InstanceNewResponseRewriteModelCfGoogleGemma3_12bIt, InstanceNewResponseRewriteModelAnthropicClaude3_7Sonnet, InstanceNewResponseRewriteModelAnthropicClaudeSonnet4, InstanceNewResponseRewriteModelAnthropicClaudeOpus4, InstanceNewResponseRewriteModelAnthropicClaude3_5Haiku, InstanceNewResponseRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceNewResponseRewriteModelCerebrasQwen3_235bA22bThinking, InstanceNewResponseRewriteModelCerebrasLlama3_3_70b, InstanceNewResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceNewResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceNewResponseRewriteModelCerebrasGptOSs120b, InstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Flash, InstanceNewResponseRewriteModelGoogleAIStudioGemini2_5Pro, InstanceNewResponseRewriteModelGrokGrok4, InstanceNewResponseRewriteModelGroqLlama3_3_70bVersatile, InstanceNewResponseRewriteModelGroqLlama3_1_8bInstant, InstanceNewResponseRewriteModelOpenAIGpt5, InstanceNewResponseRewriteModelOpenAIGpt5Mini, InstanceNewResponseRewriteModelOpenAIGpt5Nano, InstanceNewResponseRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceNewResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                  `json:"include_items"`
	Prefix         string                                    `json:"prefix"`
	R2Jurisdiction string                                    `json:"r2_jurisdiction"`
	WebCrawler     InstanceNewResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           instanceNewResponseSourceParamsJSON       `json:"-"`
}

// instanceNewResponseSourceParamsJSON contains the JSON metadata for the struct
// [InstanceNewResponseSourceParams]
type instanceNewResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseSourceParamsWebCrawler struct {
	ParseOptions InstanceNewResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    InstanceNewResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions InstanceNewResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         instanceNewResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerJSON contains the JSON metadata for the
// struct [InstanceNewResponseSourceParamsWebCrawler]
type instanceNewResponseSourceParamsWebCrawlerJSON struct {
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                  `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                      `json:"use_browser_rendering"`
	JSON                instanceNewResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct [InstanceNewResponseSourceParamsWebCrawlerParseOptions]
type instanceNewResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseSourceParamsWebCrawlerParseType string

const (
	InstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap InstanceNewResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceNewResponseSourceParamsWebCrawlerParseTypeFeedRss InstanceNewResponseSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceNewResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceNewResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceNewResponseSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceNewResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                    `json:"storage_id,required"`
	R2Jurisdiction string                                                    `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                               `json:"storage_type"`
	JSON           instanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// instanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON contains the JSON
// metadata for the struct [InstanceNewResponseSourceParamsWebCrawlerStoreOptions]
type instanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceNewResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponse struct {
	// Use your AI Search ID.
	ID                   string                                     `json:"id,required"`
	CreatedAt            time.Time                                  `json:"created_at,required" format:"date-time"`
	ModifiedAt           time.Time                                  `json:"modified_at,required" format:"date-time"`
	Source               string                                     `json:"source,required"`
	Type                 InstanceUpdateResponseType                 `json:"type,required"`
	VectorizeName        string                                     `json:"vectorize_name,required"`
	AIGatewayID          string                                     `json:"ai_gateway_id,nullable"`
	AISearchModel        InstanceUpdateResponseAISearchModel        `json:"ai_search_model"`
	Cache                bool                                       `json:"cache"`
	CacheThreshold       InstanceUpdateResponseCacheThreshold       `json:"cache_threshold"`
	ChunkOverlap         int64                                      `json:"chunk_overlap"`
	ChunkSize            int64                                      `json:"chunk_size"`
	CreatedBy            string                                     `json:"created_by,nullable"`
	CustomMetadata       []InstanceUpdateResponseCustomMetadata     `json:"custom_metadata"`
	EmbeddingModel       InstanceUpdateResponseEmbeddingModel       `json:"embedding_model"`
	Enable               bool                                       `json:"enable"`
	FusionMethod         InstanceUpdateResponseFusionMethod         `json:"fusion_method"`
	HybridSearchEnabled  bool                                       `json:"hybrid_search_enabled"`
	LastActivity         time.Time                                  `json:"last_activity,nullable" format:"date-time"`
	MaxNumResults        int64                                      `json:"max_num_results"`
	Metadata             InstanceUpdateResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                     `json:"modified_by,nullable"`
	Paused               bool                                       `json:"paused"`
	PublicEndpointID     string                                     `json:"public_endpoint_id,nullable"`
	PublicEndpointParams InstanceUpdateResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                       `json:"reranking"`
	RerankingModel       InstanceUpdateResponseRerankingModel       `json:"reranking_model"`
	RetrievalOptions     InstanceUpdateResponseRetrievalOptions     `json:"retrieval_options,nullable"`
	RewriteModel         InstanceUpdateResponseRewriteModel         `json:"rewrite_model"`
	RewriteQuery         bool                                       `json:"rewrite_query"`
	ScoreThreshold       float64                                    `json:"score_threshold"`
	SourceParams         InstanceUpdateResponseSourceParams         `json:"source_params,nullable"`
	Status               string                                     `json:"status"`
	TokenID              string                                     `json:"token_id" format:"uuid"`
	JSON                 instanceUpdateResponseJSON                 `json:"-"`
}

// instanceUpdateResponseJSON contains the JSON metadata for the struct
// [InstanceUpdateResponse]
type instanceUpdateResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	Source               apijson.Field
	Type                 apijson.Field
	VectorizeName        apijson.Field
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
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	TokenID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseType string

const (
	InstanceUpdateResponseTypeR2         InstanceUpdateResponseType = "r2"
	InstanceUpdateResponseTypeWebCrawler InstanceUpdateResponseType = "web-crawler"
)

func (r InstanceUpdateResponseType) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseTypeR2, InstanceUpdateResponseTypeWebCrawler:
		return true
	}
	return false
}

type InstanceUpdateResponseAISearchModel string

const (
	InstanceUpdateResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceUpdateResponseAISearchModelCfZaiOrgGlm4_7Flash                   InstanceUpdateResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceUpdateResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceUpdateResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceUpdateResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceUpdateResponseAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceUpdateResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceUpdateResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceUpdateResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2Instruct            InstanceUpdateResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceUpdateResponseAISearchModelCfGoogleGemma3_12bIt                  InstanceUpdateResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceUpdateResponseAISearchModelAnthropicClaude3_7Sonnet              InstanceUpdateResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceUpdateResponseAISearchModelAnthropicClaudeSonnet4                InstanceUpdateResponseAISearchModel = "anthropic/claude-sonnet-4"
	InstanceUpdateResponseAISearchModelAnthropicClaudeOpus4                  InstanceUpdateResponseAISearchModel = "anthropic/claude-opus-4"
	InstanceUpdateResponseAISearchModelAnthropicClaude3_5Haiku               InstanceUpdateResponseAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceUpdateResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bThinking        InstanceUpdateResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceUpdateResponseAISearchModelCerebrasLlama3_3_70b                  InstanceUpdateResponseAISearchModel = "cerebras/llama-3.3-70b"
	InstanceUpdateResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceUpdateResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceUpdateResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceUpdateResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceUpdateResponseAISearchModelCerebrasGptOSs120b                    InstanceUpdateResponseAISearchModel = "cerebras/gpt-oss-120b"
	InstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Flash          InstanceUpdateResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Pro            InstanceUpdateResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceUpdateResponseAISearchModelGrokGrok4                             InstanceUpdateResponseAISearchModel = "grok/grok-4"
	InstanceUpdateResponseAISearchModelGroqLlama3_3_70bVersatile             InstanceUpdateResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceUpdateResponseAISearchModelGroqLlama3_1_8bInstant                InstanceUpdateResponseAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceUpdateResponseAISearchModelOpenAIGpt5                            InstanceUpdateResponseAISearchModel = "openai/gpt-5"
	InstanceUpdateResponseAISearchModelOpenAIGpt5Mini                        InstanceUpdateResponseAISearchModel = "openai/gpt-5-mini"
	InstanceUpdateResponseAISearchModelOpenAIGpt5Nano                        InstanceUpdateResponseAISearchModel = "openai/gpt-5-nano"
	InstanceUpdateResponseAISearchModelEmpty                                 InstanceUpdateResponseAISearchModel = ""
)

func (r InstanceUpdateResponseAISearchModel) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceUpdateResponseAISearchModelCfZaiOrgGlm4_7Flash, InstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceUpdateResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceUpdateResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceUpdateResponseAISearchModelCfQwenQwen3_30bA3bFp8, InstanceUpdateResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceUpdateResponseAISearchModelCfMoonshotaiKimiK2Instruct, InstanceUpdateResponseAISearchModelCfGoogleGemma3_12bIt, InstanceUpdateResponseAISearchModelAnthropicClaude3_7Sonnet, InstanceUpdateResponseAISearchModelAnthropicClaudeSonnet4, InstanceUpdateResponseAISearchModelAnthropicClaudeOpus4, InstanceUpdateResponseAISearchModelAnthropicClaude3_5Haiku, InstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceUpdateResponseAISearchModelCerebrasQwen3_235bA22bThinking, InstanceUpdateResponseAISearchModelCerebrasLlama3_3_70b, InstanceUpdateResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceUpdateResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceUpdateResponseAISearchModelCerebrasGptOSs120b, InstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Flash, InstanceUpdateResponseAISearchModelGoogleAIStudioGemini2_5Pro, InstanceUpdateResponseAISearchModelGrokGrok4, InstanceUpdateResponseAISearchModelGroqLlama3_3_70bVersatile, InstanceUpdateResponseAISearchModelGroqLlama3_1_8bInstant, InstanceUpdateResponseAISearchModelOpenAIGpt5, InstanceUpdateResponseAISearchModelOpenAIGpt5Mini, InstanceUpdateResponseAISearchModelOpenAIGpt5Nano, InstanceUpdateResponseAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateResponseCacheThreshold string

const (
	InstanceUpdateResponseCacheThresholdSuperStrictMatch InstanceUpdateResponseCacheThreshold = "super_strict_match"
	InstanceUpdateResponseCacheThresholdCloseEnough      InstanceUpdateResponseCacheThreshold = "close_enough"
	InstanceUpdateResponseCacheThresholdFlexibleFriend   InstanceUpdateResponseCacheThreshold = "flexible_friend"
	InstanceUpdateResponseCacheThresholdAnythingGoes     InstanceUpdateResponseCacheThreshold = "anything_goes"
)

func (r InstanceUpdateResponseCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseCacheThresholdSuperStrictMatch, InstanceUpdateResponseCacheThresholdCloseEnough, InstanceUpdateResponseCacheThresholdFlexibleFriend, InstanceUpdateResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceUpdateResponseCustomMetadata struct {
	DataType  InstanceUpdateResponseCustomMetadataDataType `json:"data_type,required"`
	FieldName string                                       `json:"field_name,required"`
	JSON      instanceUpdateResponseCustomMetadataJSON     `json:"-"`
}

// instanceUpdateResponseCustomMetadataJSON contains the JSON metadata for the
// struct [InstanceUpdateResponseCustomMetadata]
type instanceUpdateResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseCustomMetadataDataType string

const (
	InstanceUpdateResponseCustomMetadataDataTypeText    InstanceUpdateResponseCustomMetadataDataType = "text"
	InstanceUpdateResponseCustomMetadataDataTypeNumber  InstanceUpdateResponseCustomMetadataDataType = "number"
	InstanceUpdateResponseCustomMetadataDataTypeBoolean InstanceUpdateResponseCustomMetadataDataType = "boolean"
)

func (r InstanceUpdateResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseCustomMetadataDataTypeText, InstanceUpdateResponseCustomMetadataDataTypeNumber, InstanceUpdateResponseCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceUpdateResponseEmbeddingModel string

const (
	InstanceUpdateResponseEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceUpdateResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceUpdateResponseEmbeddingModelCfBaaiBgeM3                      InstanceUpdateResponseEmbeddingModel = "@cf/baai/bge-m3"
	InstanceUpdateResponseEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceUpdateResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceUpdateResponseEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceUpdateResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceUpdateResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Small        InstanceUpdateResponseEmbeddingModel = "openai/text-embedding-3-small"
	InstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Large        InstanceUpdateResponseEmbeddingModel = "openai/text-embedding-3-large"
	InstanceUpdateResponseEmbeddingModelEmpty                            InstanceUpdateResponseEmbeddingModel = ""
)

func (r InstanceUpdateResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceUpdateResponseEmbeddingModelCfBaaiBgeM3, InstanceUpdateResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceUpdateResponseEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceUpdateResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Small, InstanceUpdateResponseEmbeddingModelOpenAITextEmbedding3Large, InstanceUpdateResponseEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateResponseFusionMethod string

const (
	InstanceUpdateResponseFusionMethodMax InstanceUpdateResponseFusionMethod = "max"
	InstanceUpdateResponseFusionMethodRrf InstanceUpdateResponseFusionMethod = "rrf"
)

func (r InstanceUpdateResponseFusionMethod) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseFusionMethodMax, InstanceUpdateResponseFusionMethodRrf:
		return true
	}
	return false
}

type InstanceUpdateResponseMetadata struct {
	CreatedFromAISearchWizard bool                               `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                             `json:"worker_domain"`
	JSON                      instanceUpdateResponseMetadataJSON `json:"-"`
}

// instanceUpdateResponseMetadataJSON contains the JSON metadata for the struct
// [InstanceUpdateResponseMetadata]
type instanceUpdateResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InstanceUpdateResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                          `json:"authorized_hosts"`
	ChatCompletionsEndpoint InstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	Enabled                 bool                                                              `json:"enabled"`
	Mcp                     InstanceUpdateResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceUpdateResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceUpdateResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    instanceUpdateResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsJSON contains the JSON metadata for
// the struct [InstanceUpdateResponsePublicEndpointParams]
type instanceUpdateResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstanceUpdateResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                  `json:"disabled"`
	JSON     instanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON contains
// the JSON metadata for the struct
// [InstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint]
type instanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                              `json:"disabled"`
	JSON     instanceUpdateResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [InstanceUpdateResponsePublicEndpointParamsMcp]
type instanceUpdateResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                        `json:"period_ms"`
	Requests  int64                                                        `json:"requests"`
	Technique InstanceUpdateResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      instanceUpdateResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct [InstanceUpdateResponsePublicEndpointParamsRateLimit]
type instanceUpdateResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponsePublicEndpointParamsRateLimitTechnique string

const (
	InstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed   InstanceUpdateResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding InstanceUpdateResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceUpdateResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueFixed, InstanceUpdateResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceUpdateResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                         `json:"disabled"`
	JSON     instanceUpdateResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// instanceUpdateResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct
// [InstanceUpdateResponsePublicEndpointParamsSearchEndpoint]
type instanceUpdateResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseRerankingModel string

const (
	InstanceUpdateResponseRerankingModelCfBaaiBgeRerankerBase InstanceUpdateResponseRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceUpdateResponseRerankingModelEmpty                 InstanceUpdateResponseRerankingModel = ""
)

func (r InstanceUpdateResponseRerankingModel) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseRerankingModelCfBaaiBgeRerankerBase, InstanceUpdateResponseRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateResponseRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode InstanceUpdateResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceUpdateResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceUpdateResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceUpdateResponseRetrievalOptions]
type instanceUpdateResponseRetrievalOptionsJSON struct {
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceUpdateResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceUpdateResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceUpdateResponseRetrievalOptionsKeywordMatchModeExactMatch InstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceUpdateResponseRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceUpdateResponseRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceUpdateResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseRetrievalOptionsKeywordMatchModeExactMatch, InstanceUpdateResponseRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceUpdateResponseRewriteModel string

const (
	InstanceUpdateResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceUpdateResponseRewriteModelCfZaiOrgGlm4_7Flash                   InstanceUpdateResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceUpdateResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceUpdateResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceUpdateResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceUpdateResponseRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceUpdateResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceUpdateResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceUpdateResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2Instruct            InstanceUpdateResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceUpdateResponseRewriteModelCfGoogleGemma3_12bIt                  InstanceUpdateResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceUpdateResponseRewriteModelAnthropicClaude3_7Sonnet              InstanceUpdateResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceUpdateResponseRewriteModelAnthropicClaudeSonnet4                InstanceUpdateResponseRewriteModel = "anthropic/claude-sonnet-4"
	InstanceUpdateResponseRewriteModelAnthropicClaudeOpus4                  InstanceUpdateResponseRewriteModel = "anthropic/claude-opus-4"
	InstanceUpdateResponseRewriteModelAnthropicClaude3_5Haiku               InstanceUpdateResponseRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceUpdateResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bThinking        InstanceUpdateResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceUpdateResponseRewriteModelCerebrasLlama3_3_70b                  InstanceUpdateResponseRewriteModel = "cerebras/llama-3.3-70b"
	InstanceUpdateResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceUpdateResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceUpdateResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceUpdateResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceUpdateResponseRewriteModelCerebrasGptOSs120b                    InstanceUpdateResponseRewriteModel = "cerebras/gpt-oss-120b"
	InstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Flash          InstanceUpdateResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Pro            InstanceUpdateResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceUpdateResponseRewriteModelGrokGrok4                             InstanceUpdateResponseRewriteModel = "grok/grok-4"
	InstanceUpdateResponseRewriteModelGroqLlama3_3_70bVersatile             InstanceUpdateResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceUpdateResponseRewriteModelGroqLlama3_1_8bInstant                InstanceUpdateResponseRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceUpdateResponseRewriteModelOpenAIGpt5                            InstanceUpdateResponseRewriteModel = "openai/gpt-5"
	InstanceUpdateResponseRewriteModelOpenAIGpt5Mini                        InstanceUpdateResponseRewriteModel = "openai/gpt-5-mini"
	InstanceUpdateResponseRewriteModelOpenAIGpt5Nano                        InstanceUpdateResponseRewriteModel = "openai/gpt-5-nano"
	InstanceUpdateResponseRewriteModelEmpty                                 InstanceUpdateResponseRewriteModel = ""
)

func (r InstanceUpdateResponseRewriteModel) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceUpdateResponseRewriteModelCfZaiOrgGlm4_7Flash, InstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceUpdateResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceUpdateResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceUpdateResponseRewriteModelCfQwenQwen3_30bA3bFp8, InstanceUpdateResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceUpdateResponseRewriteModelCfMoonshotaiKimiK2Instruct, InstanceUpdateResponseRewriteModelCfGoogleGemma3_12bIt, InstanceUpdateResponseRewriteModelAnthropicClaude3_7Sonnet, InstanceUpdateResponseRewriteModelAnthropicClaudeSonnet4, InstanceUpdateResponseRewriteModelAnthropicClaudeOpus4, InstanceUpdateResponseRewriteModelAnthropicClaude3_5Haiku, InstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceUpdateResponseRewriteModelCerebrasQwen3_235bA22bThinking, InstanceUpdateResponseRewriteModelCerebrasLlama3_3_70b, InstanceUpdateResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceUpdateResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceUpdateResponseRewriteModelCerebrasGptOSs120b, InstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Flash, InstanceUpdateResponseRewriteModelGoogleAIStudioGemini2_5Pro, InstanceUpdateResponseRewriteModelGrokGrok4, InstanceUpdateResponseRewriteModelGroqLlama3_3_70bVersatile, InstanceUpdateResponseRewriteModelGroqLlama3_1_8bInstant, InstanceUpdateResponseRewriteModelOpenAIGpt5, InstanceUpdateResponseRewriteModelOpenAIGpt5Mini, InstanceUpdateResponseRewriteModelOpenAIGpt5Nano, InstanceUpdateResponseRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                     `json:"include_items"`
	Prefix         string                                       `json:"prefix"`
	R2Jurisdiction string                                       `json:"r2_jurisdiction"`
	WebCrawler     InstanceUpdateResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           instanceUpdateResponseSourceParamsJSON       `json:"-"`
}

// instanceUpdateResponseSourceParamsJSON contains the JSON metadata for the struct
// [InstanceUpdateResponseSourceParams]
type instanceUpdateResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseSourceParamsWebCrawler struct {
	ParseOptions InstanceUpdateResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    InstanceUpdateResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions InstanceUpdateResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         instanceUpdateResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceUpdateResponseSourceParamsWebCrawler]
type instanceUpdateResponseSourceParamsWebCrawlerJSON struct {
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                     `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                         `json:"use_browser_rendering"`
	JSON                instanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct
// [InstanceUpdateResponseSourceParamsWebCrawlerParseOptions]
type instanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateResponseSourceParamsWebCrawlerParseType string

const (
	InstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap InstanceUpdateResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceUpdateResponseSourceParamsWebCrawlerParseTypeFeedRss InstanceUpdateResponseSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceUpdateResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceUpdateResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceUpdateResponseSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceUpdateResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                       `json:"storage_id,required"`
	R2Jurisdiction string                                                       `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                  `json:"storage_type"`
	JSON           instanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// instanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON contains the JSON
// metadata for the struct
// [InstanceUpdateResponseSourceParamsWebCrawlerStoreOptions]
type instanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceUpdateResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponse struct {
	// Use your AI Search ID.
	ID                   string                                   `json:"id,required"`
	CreatedAt            time.Time                                `json:"created_at,required" format:"date-time"`
	ModifiedAt           time.Time                                `json:"modified_at,required" format:"date-time"`
	Source               string                                   `json:"source,required"`
	Type                 InstanceListResponseType                 `json:"type,required"`
	VectorizeName        string                                   `json:"vectorize_name,required"`
	AIGatewayID          string                                   `json:"ai_gateway_id,nullable"`
	AISearchModel        InstanceListResponseAISearchModel        `json:"ai_search_model"`
	Cache                bool                                     `json:"cache"`
	CacheThreshold       InstanceListResponseCacheThreshold       `json:"cache_threshold"`
	ChunkOverlap         int64                                    `json:"chunk_overlap"`
	ChunkSize            int64                                    `json:"chunk_size"`
	CreatedBy            string                                   `json:"created_by,nullable"`
	CustomMetadata       []InstanceListResponseCustomMetadata     `json:"custom_metadata"`
	EmbeddingModel       InstanceListResponseEmbeddingModel       `json:"embedding_model"`
	Enable               bool                                     `json:"enable"`
	FusionMethod         InstanceListResponseFusionMethod         `json:"fusion_method"`
	HybridSearchEnabled  bool                                     `json:"hybrid_search_enabled"`
	LastActivity         time.Time                                `json:"last_activity,nullable" format:"date-time"`
	MaxNumResults        int64                                    `json:"max_num_results"`
	Metadata             InstanceListResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                   `json:"modified_by,nullable"`
	Paused               bool                                     `json:"paused"`
	PublicEndpointID     string                                   `json:"public_endpoint_id,nullable"`
	PublicEndpointParams InstanceListResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                     `json:"reranking"`
	RerankingModel       InstanceListResponseRerankingModel       `json:"reranking_model"`
	RetrievalOptions     InstanceListResponseRetrievalOptions     `json:"retrieval_options,nullable"`
	RewriteModel         InstanceListResponseRewriteModel         `json:"rewrite_model"`
	RewriteQuery         bool                                     `json:"rewrite_query"`
	ScoreThreshold       float64                                  `json:"score_threshold"`
	SourceParams         InstanceListResponseSourceParams         `json:"source_params,nullable"`
	Status               string                                   `json:"status"`
	TokenID              string                                   `json:"token_id" format:"uuid"`
	JSON                 instanceListResponseJSON                 `json:"-"`
}

// instanceListResponseJSON contains the JSON metadata for the struct
// [InstanceListResponse]
type instanceListResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	Source               apijson.Field
	Type                 apijson.Field
	VectorizeName        apijson.Field
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
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	TokenID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseType string

const (
	InstanceListResponseTypeR2         InstanceListResponseType = "r2"
	InstanceListResponseTypeWebCrawler InstanceListResponseType = "web-crawler"
)

func (r InstanceListResponseType) IsKnown() bool {
	switch r {
	case InstanceListResponseTypeR2, InstanceListResponseTypeWebCrawler:
		return true
	}
	return false
}

type InstanceListResponseAISearchModel string

const (
	InstanceListResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceListResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceListResponseAISearchModelCfZaiOrgGlm4_7Flash                   InstanceListResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceListResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceListResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceListResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceListResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceListResponseAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceListResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceListResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceListResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceListResponseAISearchModelCfMoonshotaiKimiK2Instruct            InstanceListResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceListResponseAISearchModelCfGoogleGemma3_12bIt                  InstanceListResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceListResponseAISearchModelAnthropicClaude3_7Sonnet              InstanceListResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceListResponseAISearchModelAnthropicClaudeSonnet4                InstanceListResponseAISearchModel = "anthropic/claude-sonnet-4"
	InstanceListResponseAISearchModelAnthropicClaudeOpus4                  InstanceListResponseAISearchModel = "anthropic/claude-opus-4"
	InstanceListResponseAISearchModelAnthropicClaude3_5Haiku               InstanceListResponseAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceListResponseAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceListResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceListResponseAISearchModelCerebrasQwen3_235bA22bThinking        InstanceListResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceListResponseAISearchModelCerebrasLlama3_3_70b                  InstanceListResponseAISearchModel = "cerebras/llama-3.3-70b"
	InstanceListResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceListResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceListResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceListResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceListResponseAISearchModelCerebrasGptOSs120b                    InstanceListResponseAISearchModel = "cerebras/gpt-oss-120b"
	InstanceListResponseAISearchModelGoogleAIStudioGemini2_5Flash          InstanceListResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceListResponseAISearchModelGoogleAIStudioGemini2_5Pro            InstanceListResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceListResponseAISearchModelGrokGrok4                             InstanceListResponseAISearchModel = "grok/grok-4"
	InstanceListResponseAISearchModelGroqLlama3_3_70bVersatile             InstanceListResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceListResponseAISearchModelGroqLlama3_1_8bInstant                InstanceListResponseAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceListResponseAISearchModelOpenAIGpt5                            InstanceListResponseAISearchModel = "openai/gpt-5"
	InstanceListResponseAISearchModelOpenAIGpt5Mini                        InstanceListResponseAISearchModel = "openai/gpt-5-mini"
	InstanceListResponseAISearchModelOpenAIGpt5Nano                        InstanceListResponseAISearchModel = "openai/gpt-5-nano"
	InstanceListResponseAISearchModelEmpty                                 InstanceListResponseAISearchModel = ""
)

func (r InstanceListResponseAISearchModel) IsKnown() bool {
	switch r {
	case InstanceListResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceListResponseAISearchModelCfZaiOrgGlm4_7Flash, InstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceListResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceListResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceListResponseAISearchModelCfQwenQwen3_30bA3bFp8, InstanceListResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceListResponseAISearchModelCfMoonshotaiKimiK2Instruct, InstanceListResponseAISearchModelCfGoogleGemma3_12bIt, InstanceListResponseAISearchModelAnthropicClaude3_7Sonnet, InstanceListResponseAISearchModelAnthropicClaudeSonnet4, InstanceListResponseAISearchModelAnthropicClaudeOpus4, InstanceListResponseAISearchModelAnthropicClaude3_5Haiku, InstanceListResponseAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceListResponseAISearchModelCerebrasQwen3_235bA22bThinking, InstanceListResponseAISearchModelCerebrasLlama3_3_70b, InstanceListResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceListResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceListResponseAISearchModelCerebrasGptOSs120b, InstanceListResponseAISearchModelGoogleAIStudioGemini2_5Flash, InstanceListResponseAISearchModelGoogleAIStudioGemini2_5Pro, InstanceListResponseAISearchModelGrokGrok4, InstanceListResponseAISearchModelGroqLlama3_3_70bVersatile, InstanceListResponseAISearchModelGroqLlama3_1_8bInstant, InstanceListResponseAISearchModelOpenAIGpt5, InstanceListResponseAISearchModelOpenAIGpt5Mini, InstanceListResponseAISearchModelOpenAIGpt5Nano, InstanceListResponseAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceListResponseCacheThreshold string

const (
	InstanceListResponseCacheThresholdSuperStrictMatch InstanceListResponseCacheThreshold = "super_strict_match"
	InstanceListResponseCacheThresholdCloseEnough      InstanceListResponseCacheThreshold = "close_enough"
	InstanceListResponseCacheThresholdFlexibleFriend   InstanceListResponseCacheThreshold = "flexible_friend"
	InstanceListResponseCacheThresholdAnythingGoes     InstanceListResponseCacheThreshold = "anything_goes"
)

func (r InstanceListResponseCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceListResponseCacheThresholdSuperStrictMatch, InstanceListResponseCacheThresholdCloseEnough, InstanceListResponseCacheThresholdFlexibleFriend, InstanceListResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceListResponseCustomMetadata struct {
	DataType  InstanceListResponseCustomMetadataDataType `json:"data_type,required"`
	FieldName string                                     `json:"field_name,required"`
	JSON      instanceListResponseCustomMetadataJSON     `json:"-"`
}

// instanceListResponseCustomMetadataJSON contains the JSON metadata for the struct
// [InstanceListResponseCustomMetadata]
type instanceListResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseCustomMetadataDataType string

const (
	InstanceListResponseCustomMetadataDataTypeText    InstanceListResponseCustomMetadataDataType = "text"
	InstanceListResponseCustomMetadataDataTypeNumber  InstanceListResponseCustomMetadataDataType = "number"
	InstanceListResponseCustomMetadataDataTypeBoolean InstanceListResponseCustomMetadataDataType = "boolean"
)

func (r InstanceListResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceListResponseCustomMetadataDataTypeText, InstanceListResponseCustomMetadataDataTypeNumber, InstanceListResponseCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceListResponseEmbeddingModel string

const (
	InstanceListResponseEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceListResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceListResponseEmbeddingModelCfBaaiBgeM3                      InstanceListResponseEmbeddingModel = "@cf/baai/bge-m3"
	InstanceListResponseEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceListResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceListResponseEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceListResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceListResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceListResponseEmbeddingModelOpenAITextEmbedding3Small        InstanceListResponseEmbeddingModel = "openai/text-embedding-3-small"
	InstanceListResponseEmbeddingModelOpenAITextEmbedding3Large        InstanceListResponseEmbeddingModel = "openai/text-embedding-3-large"
	InstanceListResponseEmbeddingModelEmpty                            InstanceListResponseEmbeddingModel = ""
)

func (r InstanceListResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceListResponseEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceListResponseEmbeddingModelCfBaaiBgeM3, InstanceListResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceListResponseEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceListResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceListResponseEmbeddingModelOpenAITextEmbedding3Small, InstanceListResponseEmbeddingModelOpenAITextEmbedding3Large, InstanceListResponseEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceListResponseFusionMethod string

const (
	InstanceListResponseFusionMethodMax InstanceListResponseFusionMethod = "max"
	InstanceListResponseFusionMethodRrf InstanceListResponseFusionMethod = "rrf"
)

func (r InstanceListResponseFusionMethod) IsKnown() bool {
	switch r {
	case InstanceListResponseFusionMethodMax, InstanceListResponseFusionMethodRrf:
		return true
	}
	return false
}

type InstanceListResponseMetadata struct {
	CreatedFromAISearchWizard bool                             `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                           `json:"worker_domain"`
	JSON                      instanceListResponseMetadataJSON `json:"-"`
}

// instanceListResponseMetadataJSON contains the JSON metadata for the struct
// [InstanceListResponseMetadata]
type instanceListResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InstanceListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                        `json:"authorized_hosts"`
	ChatCompletionsEndpoint InstanceListResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	Enabled                 bool                                                            `json:"enabled"`
	Mcp                     InstanceListResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceListResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceListResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    instanceListResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceListResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceListResponsePublicEndpointParams]
type instanceListResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstanceListResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                `json:"disabled"`
	JSON     instanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// instanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON contains the
// JSON metadata for the struct
// [InstanceListResponsePublicEndpointParamsChatCompletionsEndpoint]
type instanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                            `json:"disabled"`
	JSON     instanceListResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// instanceListResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [InstanceListResponsePublicEndpointParamsMcp]
type instanceListResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                      `json:"period_ms"`
	Requests  int64                                                      `json:"requests"`
	Technique InstanceListResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      instanceListResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// instanceListResponsePublicEndpointParamsRateLimitJSON contains the JSON metadata
// for the struct [InstanceListResponsePublicEndpointParamsRateLimit]
type instanceListResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponsePublicEndpointParamsRateLimitTechnique string

const (
	InstanceListResponsePublicEndpointParamsRateLimitTechniqueFixed   InstanceListResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceListResponsePublicEndpointParamsRateLimitTechniqueSliding InstanceListResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceListResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceListResponsePublicEndpointParamsRateLimitTechniqueFixed, InstanceListResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceListResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                       `json:"disabled"`
	JSON     instanceListResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// instanceListResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct [InstanceListResponsePublicEndpointParamsSearchEndpoint]
type instanceListResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseRerankingModel string

const (
	InstanceListResponseRerankingModelCfBaaiBgeRerankerBase InstanceListResponseRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceListResponseRerankingModelEmpty                 InstanceListResponseRerankingModel = ""
)

func (r InstanceListResponseRerankingModel) IsKnown() bool {
	switch r {
	case InstanceListResponseRerankingModelCfBaaiBgeRerankerBase, InstanceListResponseRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceListResponseRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode InstanceListResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceListResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceListResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceListResponseRetrievalOptions]
type instanceListResponseRetrievalOptionsJSON struct {
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceListResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceListResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceListResponseRetrievalOptionsKeywordMatchModeExactMatch InstanceListResponseRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceListResponseRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceListResponseRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceListResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceListResponseRetrievalOptionsKeywordMatchModeExactMatch, InstanceListResponseRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceListResponseRewriteModel string

const (
	InstanceListResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceListResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceListResponseRewriteModelCfZaiOrgGlm4_7Flash                   InstanceListResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceListResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceListResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceListResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceListResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceListResponseRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceListResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceListResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceListResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceListResponseRewriteModelCfMoonshotaiKimiK2Instruct            InstanceListResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceListResponseRewriteModelCfGoogleGemma3_12bIt                  InstanceListResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceListResponseRewriteModelAnthropicClaude3_7Sonnet              InstanceListResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceListResponseRewriteModelAnthropicClaudeSonnet4                InstanceListResponseRewriteModel = "anthropic/claude-sonnet-4"
	InstanceListResponseRewriteModelAnthropicClaudeOpus4                  InstanceListResponseRewriteModel = "anthropic/claude-opus-4"
	InstanceListResponseRewriteModelAnthropicClaude3_5Haiku               InstanceListResponseRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceListResponseRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceListResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceListResponseRewriteModelCerebrasQwen3_235bA22bThinking        InstanceListResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceListResponseRewriteModelCerebrasLlama3_3_70b                  InstanceListResponseRewriteModel = "cerebras/llama-3.3-70b"
	InstanceListResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceListResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceListResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceListResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceListResponseRewriteModelCerebrasGptOSs120b                    InstanceListResponseRewriteModel = "cerebras/gpt-oss-120b"
	InstanceListResponseRewriteModelGoogleAIStudioGemini2_5Flash          InstanceListResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceListResponseRewriteModelGoogleAIStudioGemini2_5Pro            InstanceListResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceListResponseRewriteModelGrokGrok4                             InstanceListResponseRewriteModel = "grok/grok-4"
	InstanceListResponseRewriteModelGroqLlama3_3_70bVersatile             InstanceListResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceListResponseRewriteModelGroqLlama3_1_8bInstant                InstanceListResponseRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceListResponseRewriteModelOpenAIGpt5                            InstanceListResponseRewriteModel = "openai/gpt-5"
	InstanceListResponseRewriteModelOpenAIGpt5Mini                        InstanceListResponseRewriteModel = "openai/gpt-5-mini"
	InstanceListResponseRewriteModelOpenAIGpt5Nano                        InstanceListResponseRewriteModel = "openai/gpt-5-nano"
	InstanceListResponseRewriteModelEmpty                                 InstanceListResponseRewriteModel = ""
)

func (r InstanceListResponseRewriteModel) IsKnown() bool {
	switch r {
	case InstanceListResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceListResponseRewriteModelCfZaiOrgGlm4_7Flash, InstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceListResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceListResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceListResponseRewriteModelCfQwenQwen3_30bA3bFp8, InstanceListResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceListResponseRewriteModelCfMoonshotaiKimiK2Instruct, InstanceListResponseRewriteModelCfGoogleGemma3_12bIt, InstanceListResponseRewriteModelAnthropicClaude3_7Sonnet, InstanceListResponseRewriteModelAnthropicClaudeSonnet4, InstanceListResponseRewriteModelAnthropicClaudeOpus4, InstanceListResponseRewriteModelAnthropicClaude3_5Haiku, InstanceListResponseRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceListResponseRewriteModelCerebrasQwen3_235bA22bThinking, InstanceListResponseRewriteModelCerebrasLlama3_3_70b, InstanceListResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceListResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceListResponseRewriteModelCerebrasGptOSs120b, InstanceListResponseRewriteModelGoogleAIStudioGemini2_5Flash, InstanceListResponseRewriteModelGoogleAIStudioGemini2_5Pro, InstanceListResponseRewriteModelGrokGrok4, InstanceListResponseRewriteModelGroqLlama3_3_70bVersatile, InstanceListResponseRewriteModelGroqLlama3_1_8bInstant, InstanceListResponseRewriteModelOpenAIGpt5, InstanceListResponseRewriteModelOpenAIGpt5Mini, InstanceListResponseRewriteModelOpenAIGpt5Nano, InstanceListResponseRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceListResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                   `json:"include_items"`
	Prefix         string                                     `json:"prefix"`
	R2Jurisdiction string                                     `json:"r2_jurisdiction"`
	WebCrawler     InstanceListResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           instanceListResponseSourceParamsJSON       `json:"-"`
}

// instanceListResponseSourceParamsJSON contains the JSON metadata for the struct
// [InstanceListResponseSourceParams]
type instanceListResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceListResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawler struct {
	ParseOptions InstanceListResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    InstanceListResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions InstanceListResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         instanceListResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceListResponseSourceParamsWebCrawler]
type instanceListResponseSourceParamsWebCrawlerJSON struct {
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                   `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                       `json:"use_browser_rendering"`
	JSON                instanceListResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct [InstanceListResponseSourceParamsWebCrawlerParseOptions]
type instanceListResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseSourceParamsWebCrawlerParseType string

const (
	InstanceListResponseSourceParamsWebCrawlerParseTypeSitemap InstanceListResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceListResponseSourceParamsWebCrawlerParseTypeFeedRss InstanceListResponseSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceListResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceListResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceListResponseSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceListResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                     `json:"storage_id,required"`
	R2Jurisdiction string                                                     `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                `json:"storage_type"`
	JSON           instanceListResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// instanceListResponseSourceParamsWebCrawlerStoreOptionsJSON contains the JSON
// metadata for the struct [InstanceListResponseSourceParamsWebCrawlerStoreOptions]
type instanceListResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceListResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponse struct {
	// Use your AI Search ID.
	ID                   string                                     `json:"id,required"`
	CreatedAt            time.Time                                  `json:"created_at,required" format:"date-time"`
	ModifiedAt           time.Time                                  `json:"modified_at,required" format:"date-time"`
	Source               string                                     `json:"source,required"`
	Type                 InstanceDeleteResponseType                 `json:"type,required"`
	VectorizeName        string                                     `json:"vectorize_name,required"`
	AIGatewayID          string                                     `json:"ai_gateway_id,nullable"`
	AISearchModel        InstanceDeleteResponseAISearchModel        `json:"ai_search_model"`
	Cache                bool                                       `json:"cache"`
	CacheThreshold       InstanceDeleteResponseCacheThreshold       `json:"cache_threshold"`
	ChunkOverlap         int64                                      `json:"chunk_overlap"`
	ChunkSize            int64                                      `json:"chunk_size"`
	CreatedBy            string                                     `json:"created_by,nullable"`
	CustomMetadata       []InstanceDeleteResponseCustomMetadata     `json:"custom_metadata"`
	EmbeddingModel       InstanceDeleteResponseEmbeddingModel       `json:"embedding_model"`
	Enable               bool                                       `json:"enable"`
	FusionMethod         InstanceDeleteResponseFusionMethod         `json:"fusion_method"`
	HybridSearchEnabled  bool                                       `json:"hybrid_search_enabled"`
	LastActivity         time.Time                                  `json:"last_activity,nullable" format:"date-time"`
	MaxNumResults        int64                                      `json:"max_num_results"`
	Metadata             InstanceDeleteResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                     `json:"modified_by,nullable"`
	Paused               bool                                       `json:"paused"`
	PublicEndpointID     string                                     `json:"public_endpoint_id,nullable"`
	PublicEndpointParams InstanceDeleteResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                       `json:"reranking"`
	RerankingModel       InstanceDeleteResponseRerankingModel       `json:"reranking_model"`
	RetrievalOptions     InstanceDeleteResponseRetrievalOptions     `json:"retrieval_options,nullable"`
	RewriteModel         InstanceDeleteResponseRewriteModel         `json:"rewrite_model"`
	RewriteQuery         bool                                       `json:"rewrite_query"`
	ScoreThreshold       float64                                    `json:"score_threshold"`
	SourceParams         InstanceDeleteResponseSourceParams         `json:"source_params,nullable"`
	Status               string                                     `json:"status"`
	TokenID              string                                     `json:"token_id" format:"uuid"`
	JSON                 instanceDeleteResponseJSON                 `json:"-"`
}

// instanceDeleteResponseJSON contains the JSON metadata for the struct
// [InstanceDeleteResponse]
type instanceDeleteResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	Source               apijson.Field
	Type                 apijson.Field
	VectorizeName        apijson.Field
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
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	TokenID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseType string

const (
	InstanceDeleteResponseTypeR2         InstanceDeleteResponseType = "r2"
	InstanceDeleteResponseTypeWebCrawler InstanceDeleteResponseType = "web-crawler"
)

func (r InstanceDeleteResponseType) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseTypeR2, InstanceDeleteResponseTypeWebCrawler:
		return true
	}
	return false
}

type InstanceDeleteResponseAISearchModel string

const (
	InstanceDeleteResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceDeleteResponseAISearchModelCfZaiOrgGlm4_7Flash                   InstanceDeleteResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceDeleteResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceDeleteResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceDeleteResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceDeleteResponseAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceDeleteResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceDeleteResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceDeleteResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2Instruct            InstanceDeleteResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceDeleteResponseAISearchModelCfGoogleGemma3_12bIt                  InstanceDeleteResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceDeleteResponseAISearchModelAnthropicClaude3_7Sonnet              InstanceDeleteResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceDeleteResponseAISearchModelAnthropicClaudeSonnet4                InstanceDeleteResponseAISearchModel = "anthropic/claude-sonnet-4"
	InstanceDeleteResponseAISearchModelAnthropicClaudeOpus4                  InstanceDeleteResponseAISearchModel = "anthropic/claude-opus-4"
	InstanceDeleteResponseAISearchModelAnthropicClaude3_5Haiku               InstanceDeleteResponseAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceDeleteResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bThinking        InstanceDeleteResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceDeleteResponseAISearchModelCerebrasLlama3_3_70b                  InstanceDeleteResponseAISearchModel = "cerebras/llama-3.3-70b"
	InstanceDeleteResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceDeleteResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceDeleteResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceDeleteResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceDeleteResponseAISearchModelCerebrasGptOSs120b                    InstanceDeleteResponseAISearchModel = "cerebras/gpt-oss-120b"
	InstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Flash          InstanceDeleteResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Pro            InstanceDeleteResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceDeleteResponseAISearchModelGrokGrok4                             InstanceDeleteResponseAISearchModel = "grok/grok-4"
	InstanceDeleteResponseAISearchModelGroqLlama3_3_70bVersatile             InstanceDeleteResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceDeleteResponseAISearchModelGroqLlama3_1_8bInstant                InstanceDeleteResponseAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceDeleteResponseAISearchModelOpenAIGpt5                            InstanceDeleteResponseAISearchModel = "openai/gpt-5"
	InstanceDeleteResponseAISearchModelOpenAIGpt5Mini                        InstanceDeleteResponseAISearchModel = "openai/gpt-5-mini"
	InstanceDeleteResponseAISearchModelOpenAIGpt5Nano                        InstanceDeleteResponseAISearchModel = "openai/gpt-5-nano"
	InstanceDeleteResponseAISearchModelEmpty                                 InstanceDeleteResponseAISearchModel = ""
)

func (r InstanceDeleteResponseAISearchModel) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceDeleteResponseAISearchModelCfZaiOrgGlm4_7Flash, InstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceDeleteResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceDeleteResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceDeleteResponseAISearchModelCfQwenQwen3_30bA3bFp8, InstanceDeleteResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceDeleteResponseAISearchModelCfMoonshotaiKimiK2Instruct, InstanceDeleteResponseAISearchModelCfGoogleGemma3_12bIt, InstanceDeleteResponseAISearchModelAnthropicClaude3_7Sonnet, InstanceDeleteResponseAISearchModelAnthropicClaudeSonnet4, InstanceDeleteResponseAISearchModelAnthropicClaudeOpus4, InstanceDeleteResponseAISearchModelAnthropicClaude3_5Haiku, InstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceDeleteResponseAISearchModelCerebrasQwen3_235bA22bThinking, InstanceDeleteResponseAISearchModelCerebrasLlama3_3_70b, InstanceDeleteResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceDeleteResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceDeleteResponseAISearchModelCerebrasGptOSs120b, InstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Flash, InstanceDeleteResponseAISearchModelGoogleAIStudioGemini2_5Pro, InstanceDeleteResponseAISearchModelGrokGrok4, InstanceDeleteResponseAISearchModelGroqLlama3_3_70bVersatile, InstanceDeleteResponseAISearchModelGroqLlama3_1_8bInstant, InstanceDeleteResponseAISearchModelOpenAIGpt5, InstanceDeleteResponseAISearchModelOpenAIGpt5Mini, InstanceDeleteResponseAISearchModelOpenAIGpt5Nano, InstanceDeleteResponseAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceDeleteResponseCacheThreshold string

const (
	InstanceDeleteResponseCacheThresholdSuperStrictMatch InstanceDeleteResponseCacheThreshold = "super_strict_match"
	InstanceDeleteResponseCacheThresholdCloseEnough      InstanceDeleteResponseCacheThreshold = "close_enough"
	InstanceDeleteResponseCacheThresholdFlexibleFriend   InstanceDeleteResponseCacheThreshold = "flexible_friend"
	InstanceDeleteResponseCacheThresholdAnythingGoes     InstanceDeleteResponseCacheThreshold = "anything_goes"
)

func (r InstanceDeleteResponseCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseCacheThresholdSuperStrictMatch, InstanceDeleteResponseCacheThresholdCloseEnough, InstanceDeleteResponseCacheThresholdFlexibleFriend, InstanceDeleteResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceDeleteResponseCustomMetadata struct {
	DataType  InstanceDeleteResponseCustomMetadataDataType `json:"data_type,required"`
	FieldName string                                       `json:"field_name,required"`
	JSON      instanceDeleteResponseCustomMetadataJSON     `json:"-"`
}

// instanceDeleteResponseCustomMetadataJSON contains the JSON metadata for the
// struct [InstanceDeleteResponseCustomMetadata]
type instanceDeleteResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseCustomMetadataDataType string

const (
	InstanceDeleteResponseCustomMetadataDataTypeText    InstanceDeleteResponseCustomMetadataDataType = "text"
	InstanceDeleteResponseCustomMetadataDataTypeNumber  InstanceDeleteResponseCustomMetadataDataType = "number"
	InstanceDeleteResponseCustomMetadataDataTypeBoolean InstanceDeleteResponseCustomMetadataDataType = "boolean"
)

func (r InstanceDeleteResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseCustomMetadataDataTypeText, InstanceDeleteResponseCustomMetadataDataTypeNumber, InstanceDeleteResponseCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceDeleteResponseEmbeddingModel string

const (
	InstanceDeleteResponseEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceDeleteResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceDeleteResponseEmbeddingModelCfBaaiBgeM3                      InstanceDeleteResponseEmbeddingModel = "@cf/baai/bge-m3"
	InstanceDeleteResponseEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceDeleteResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceDeleteResponseEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceDeleteResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceDeleteResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Small        InstanceDeleteResponseEmbeddingModel = "openai/text-embedding-3-small"
	InstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Large        InstanceDeleteResponseEmbeddingModel = "openai/text-embedding-3-large"
	InstanceDeleteResponseEmbeddingModelEmpty                            InstanceDeleteResponseEmbeddingModel = ""
)

func (r InstanceDeleteResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceDeleteResponseEmbeddingModelCfBaaiBgeM3, InstanceDeleteResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceDeleteResponseEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceDeleteResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Small, InstanceDeleteResponseEmbeddingModelOpenAITextEmbedding3Large, InstanceDeleteResponseEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceDeleteResponseFusionMethod string

const (
	InstanceDeleteResponseFusionMethodMax InstanceDeleteResponseFusionMethod = "max"
	InstanceDeleteResponseFusionMethodRrf InstanceDeleteResponseFusionMethod = "rrf"
)

func (r InstanceDeleteResponseFusionMethod) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseFusionMethodMax, InstanceDeleteResponseFusionMethodRrf:
		return true
	}
	return false
}

type InstanceDeleteResponseMetadata struct {
	CreatedFromAISearchWizard bool                               `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                             `json:"worker_domain"`
	JSON                      instanceDeleteResponseMetadataJSON `json:"-"`
}

// instanceDeleteResponseMetadataJSON contains the JSON metadata for the struct
// [InstanceDeleteResponseMetadata]
type instanceDeleteResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InstanceDeleteResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                          `json:"authorized_hosts"`
	ChatCompletionsEndpoint InstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	Enabled                 bool                                                              `json:"enabled"`
	Mcp                     InstanceDeleteResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceDeleteResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceDeleteResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    instanceDeleteResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsJSON contains the JSON metadata for
// the struct [InstanceDeleteResponsePublicEndpointParams]
type instanceDeleteResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstanceDeleteResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                  `json:"disabled"`
	JSON     instanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON contains
// the JSON metadata for the struct
// [InstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint]
type instanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                              `json:"disabled"`
	JSON     instanceDeleteResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [InstanceDeleteResponsePublicEndpointParamsMcp]
type instanceDeleteResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                        `json:"period_ms"`
	Requests  int64                                                        `json:"requests"`
	Technique InstanceDeleteResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      instanceDeleteResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsRateLimitJSON contains the JSON
// metadata for the struct [InstanceDeleteResponsePublicEndpointParamsRateLimit]
type instanceDeleteResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponsePublicEndpointParamsRateLimitTechnique string

const (
	InstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueFixed   InstanceDeleteResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueSliding InstanceDeleteResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceDeleteResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueFixed, InstanceDeleteResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceDeleteResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                         `json:"disabled"`
	JSON     instanceDeleteResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// instanceDeleteResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct
// [InstanceDeleteResponsePublicEndpointParamsSearchEndpoint]
type instanceDeleteResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseRerankingModel string

const (
	InstanceDeleteResponseRerankingModelCfBaaiBgeRerankerBase InstanceDeleteResponseRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceDeleteResponseRerankingModelEmpty                 InstanceDeleteResponseRerankingModel = ""
)

func (r InstanceDeleteResponseRerankingModel) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseRerankingModelCfBaaiBgeRerankerBase, InstanceDeleteResponseRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceDeleteResponseRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode InstanceDeleteResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceDeleteResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceDeleteResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceDeleteResponseRetrievalOptions]
type instanceDeleteResponseRetrievalOptionsJSON struct {
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceDeleteResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceDeleteResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceDeleteResponseRetrievalOptionsKeywordMatchModeExactMatch InstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceDeleteResponseRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceDeleteResponseRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceDeleteResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseRetrievalOptionsKeywordMatchModeExactMatch, InstanceDeleteResponseRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceDeleteResponseRewriteModel string

const (
	InstanceDeleteResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceDeleteResponseRewriteModelCfZaiOrgGlm4_7Flash                   InstanceDeleteResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceDeleteResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceDeleteResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceDeleteResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceDeleteResponseRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceDeleteResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceDeleteResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceDeleteResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2Instruct            InstanceDeleteResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceDeleteResponseRewriteModelCfGoogleGemma3_12bIt                  InstanceDeleteResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceDeleteResponseRewriteModelAnthropicClaude3_7Sonnet              InstanceDeleteResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceDeleteResponseRewriteModelAnthropicClaudeSonnet4                InstanceDeleteResponseRewriteModel = "anthropic/claude-sonnet-4"
	InstanceDeleteResponseRewriteModelAnthropicClaudeOpus4                  InstanceDeleteResponseRewriteModel = "anthropic/claude-opus-4"
	InstanceDeleteResponseRewriteModelAnthropicClaude3_5Haiku               InstanceDeleteResponseRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceDeleteResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bThinking        InstanceDeleteResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceDeleteResponseRewriteModelCerebrasLlama3_3_70b                  InstanceDeleteResponseRewriteModel = "cerebras/llama-3.3-70b"
	InstanceDeleteResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceDeleteResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceDeleteResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceDeleteResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceDeleteResponseRewriteModelCerebrasGptOSs120b                    InstanceDeleteResponseRewriteModel = "cerebras/gpt-oss-120b"
	InstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Flash          InstanceDeleteResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Pro            InstanceDeleteResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceDeleteResponseRewriteModelGrokGrok4                             InstanceDeleteResponseRewriteModel = "grok/grok-4"
	InstanceDeleteResponseRewriteModelGroqLlama3_3_70bVersatile             InstanceDeleteResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceDeleteResponseRewriteModelGroqLlama3_1_8bInstant                InstanceDeleteResponseRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceDeleteResponseRewriteModelOpenAIGpt5                            InstanceDeleteResponseRewriteModel = "openai/gpt-5"
	InstanceDeleteResponseRewriteModelOpenAIGpt5Mini                        InstanceDeleteResponseRewriteModel = "openai/gpt-5-mini"
	InstanceDeleteResponseRewriteModelOpenAIGpt5Nano                        InstanceDeleteResponseRewriteModel = "openai/gpt-5-nano"
	InstanceDeleteResponseRewriteModelEmpty                                 InstanceDeleteResponseRewriteModel = ""
)

func (r InstanceDeleteResponseRewriteModel) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceDeleteResponseRewriteModelCfZaiOrgGlm4_7Flash, InstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceDeleteResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceDeleteResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceDeleteResponseRewriteModelCfQwenQwen3_30bA3bFp8, InstanceDeleteResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceDeleteResponseRewriteModelCfMoonshotaiKimiK2Instruct, InstanceDeleteResponseRewriteModelCfGoogleGemma3_12bIt, InstanceDeleteResponseRewriteModelAnthropicClaude3_7Sonnet, InstanceDeleteResponseRewriteModelAnthropicClaudeSonnet4, InstanceDeleteResponseRewriteModelAnthropicClaudeOpus4, InstanceDeleteResponseRewriteModelAnthropicClaude3_5Haiku, InstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceDeleteResponseRewriteModelCerebrasQwen3_235bA22bThinking, InstanceDeleteResponseRewriteModelCerebrasLlama3_3_70b, InstanceDeleteResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceDeleteResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceDeleteResponseRewriteModelCerebrasGptOSs120b, InstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Flash, InstanceDeleteResponseRewriteModelGoogleAIStudioGemini2_5Pro, InstanceDeleteResponseRewriteModelGrokGrok4, InstanceDeleteResponseRewriteModelGroqLlama3_3_70bVersatile, InstanceDeleteResponseRewriteModelGroqLlama3_1_8bInstant, InstanceDeleteResponseRewriteModelOpenAIGpt5, InstanceDeleteResponseRewriteModelOpenAIGpt5Mini, InstanceDeleteResponseRewriteModelOpenAIGpt5Nano, InstanceDeleteResponseRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceDeleteResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                     `json:"include_items"`
	Prefix         string                                       `json:"prefix"`
	R2Jurisdiction string                                       `json:"r2_jurisdiction"`
	WebCrawler     InstanceDeleteResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           instanceDeleteResponseSourceParamsJSON       `json:"-"`
}

// instanceDeleteResponseSourceParamsJSON contains the JSON metadata for the struct
// [InstanceDeleteResponseSourceParams]
type instanceDeleteResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseSourceParamsWebCrawler struct {
	ParseOptions InstanceDeleteResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    InstanceDeleteResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions InstanceDeleteResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         instanceDeleteResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceDeleteResponseSourceParamsWebCrawler]
type instanceDeleteResponseSourceParamsWebCrawlerJSON struct {
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                     `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                         `json:"use_browser_rendering"`
	JSON                instanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct
// [InstanceDeleteResponseSourceParamsWebCrawlerParseOptions]
type instanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceDeleteResponseSourceParamsWebCrawlerParseType string

const (
	InstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap InstanceDeleteResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceDeleteResponseSourceParamsWebCrawlerParseTypeFeedRss InstanceDeleteResponseSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceDeleteResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceDeleteResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceDeleteResponseSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceDeleteResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                       `json:"storage_id,required"`
	R2Jurisdiction string                                                       `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                  `json:"storage_type"`
	JSON           instanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// instanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON contains the JSON
// metadata for the struct
// [InstanceDeleteResponseSourceParamsWebCrawlerStoreOptions]
type instanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceDeleteResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponse struct {
	Choices     []InstanceChatCompletionsResponseChoice `json:"choices,required"`
	Chunks      []InstanceChatCompletionsResponseChunk  `json:"chunks,required"`
	ID          string                                  `json:"id"`
	Model       string                                  `json:"model"`
	Object      string                                  `json:"object"`
	ExtraFields map[string]interface{}                  `json:"-,extras"`
	JSON        instanceChatCompletionsResponseJSON     `json:"-"`
}

// instanceChatCompletionsResponseJSON contains the JSON metadata for the struct
// [InstanceChatCompletionsResponse]
type instanceChatCompletionsResponseJSON struct {
	Choices     apijson.Field
	Chunks      apijson.Field
	ID          apijson.Field
	Model       apijson.Field
	Object      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponseChoice struct {
	Message InstanceChatCompletionsResponseChoicesMessage `json:"message,required"`
	Index   int64                                         `json:"index"`
	JSON    instanceChatCompletionsResponseChoiceJSON     `json:"-"`
}

// instanceChatCompletionsResponseChoiceJSON contains the JSON metadata for the
// struct [InstanceChatCompletionsResponseChoice]
type instanceChatCompletionsResponseChoiceJSON struct {
	Message     apijson.Field
	Index       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChoice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChoiceJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponseChoicesMessage struct {
	Content     string                                            `json:"content,required,nullable"`
	Role        InstanceChatCompletionsResponseChoicesMessageRole `json:"role,required"`
	ExtraFields map[string]interface{}                            `json:"-,extras"`
	JSON        instanceChatCompletionsResponseChoicesMessageJSON `json:"-"`
}

// instanceChatCompletionsResponseChoicesMessageJSON contains the JSON metadata for
// the struct [InstanceChatCompletionsResponseChoicesMessage]
type instanceChatCompletionsResponseChoicesMessageJSON struct {
	Content     apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChoicesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChoicesMessageJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponseChoicesMessageRole string

const (
	InstanceChatCompletionsResponseChoicesMessageRoleSystem    InstanceChatCompletionsResponseChoicesMessageRole = "system"
	InstanceChatCompletionsResponseChoicesMessageRoleDeveloper InstanceChatCompletionsResponseChoicesMessageRole = "developer"
	InstanceChatCompletionsResponseChoicesMessageRoleUser      InstanceChatCompletionsResponseChoicesMessageRole = "user"
	InstanceChatCompletionsResponseChoicesMessageRoleAssistant InstanceChatCompletionsResponseChoicesMessageRole = "assistant"
	InstanceChatCompletionsResponseChoicesMessageRoleTool      InstanceChatCompletionsResponseChoicesMessageRole = "tool"
)

func (r InstanceChatCompletionsResponseChoicesMessageRole) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsResponseChoicesMessageRoleSystem, InstanceChatCompletionsResponseChoicesMessageRoleDeveloper, InstanceChatCompletionsResponseChoicesMessageRoleUser, InstanceChatCompletionsResponseChoicesMessageRoleAssistant, InstanceChatCompletionsResponseChoicesMessageRoleTool:
		return true
	}
	return false
}

type InstanceChatCompletionsResponseChunk struct {
	ID             string                                              `json:"id,required"`
	Score          float64                                             `json:"score,required"`
	Text           string                                              `json:"text,required"`
	Type           string                                              `json:"type,required"`
	Item           InstanceChatCompletionsResponseChunksItem           `json:"item"`
	ScoringDetails InstanceChatCompletionsResponseChunksScoringDetails `json:"scoring_details"`
	JSON           instanceChatCompletionsResponseChunkJSON            `json:"-"`
}

// instanceChatCompletionsResponseChunkJSON contains the JSON metadata for the
// struct [InstanceChatCompletionsResponseChunk]
type instanceChatCompletionsResponseChunkJSON struct {
	ID             apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChunkJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponseChunksItem struct {
	Key       string                                        `json:"key,required"`
	Metadata  map[string]interface{}                        `json:"metadata"`
	Timestamp float64                                       `json:"timestamp"`
	JSON      instanceChatCompletionsResponseChunksItemJSON `json:"-"`
}

// instanceChatCompletionsResponseChunksItemJSON contains the JSON metadata for the
// struct [InstanceChatCompletionsResponseChunksItem]
type instanceChatCompletionsResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsResponseChunksScoringDetails struct {
	KeywordRank    float64                                                 `json:"keyword_rank"`
	KeywordScore   float64                                                 `json:"keyword_score"`
	RerankingScore float64                                                 `json:"reranking_score"`
	VectorRank     float64                                                 `json:"vector_rank"`
	VectorScore    float64                                                 `json:"vector_score"`
	JSON           instanceChatCompletionsResponseChunksScoringDetailsJSON `json:"-"`
}

// instanceChatCompletionsResponseChunksScoringDetailsJSON contains the JSON
// metadata for the struct [InstanceChatCompletionsResponseChunksScoringDetails]
type instanceChatCompletionsResponseChunksScoringDetailsJSON struct {
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceChatCompletionsResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceChatCompletionsResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponse struct {
	// Use your AI Search ID.
	ID                   string                                   `json:"id,required"`
	CreatedAt            time.Time                                `json:"created_at,required" format:"date-time"`
	ModifiedAt           time.Time                                `json:"modified_at,required" format:"date-time"`
	Source               string                                   `json:"source,required"`
	Type                 InstanceReadResponseType                 `json:"type,required"`
	VectorizeName        string                                   `json:"vectorize_name,required"`
	AIGatewayID          string                                   `json:"ai_gateway_id,nullable"`
	AISearchModel        InstanceReadResponseAISearchModel        `json:"ai_search_model"`
	Cache                bool                                     `json:"cache"`
	CacheThreshold       InstanceReadResponseCacheThreshold       `json:"cache_threshold"`
	ChunkOverlap         int64                                    `json:"chunk_overlap"`
	ChunkSize            int64                                    `json:"chunk_size"`
	CreatedBy            string                                   `json:"created_by,nullable"`
	CustomMetadata       []InstanceReadResponseCustomMetadata     `json:"custom_metadata"`
	EmbeddingModel       InstanceReadResponseEmbeddingModel       `json:"embedding_model"`
	Enable               bool                                     `json:"enable"`
	FusionMethod         InstanceReadResponseFusionMethod         `json:"fusion_method"`
	HybridSearchEnabled  bool                                     `json:"hybrid_search_enabled"`
	LastActivity         time.Time                                `json:"last_activity,nullable" format:"date-time"`
	MaxNumResults        int64                                    `json:"max_num_results"`
	Metadata             InstanceReadResponseMetadata             `json:"metadata"`
	ModifiedBy           string                                   `json:"modified_by,nullable"`
	Paused               bool                                     `json:"paused"`
	PublicEndpointID     string                                   `json:"public_endpoint_id,nullable"`
	PublicEndpointParams InstanceReadResponsePublicEndpointParams `json:"public_endpoint_params"`
	Reranking            bool                                     `json:"reranking"`
	RerankingModel       InstanceReadResponseRerankingModel       `json:"reranking_model"`
	RetrievalOptions     InstanceReadResponseRetrievalOptions     `json:"retrieval_options,nullable"`
	RewriteModel         InstanceReadResponseRewriteModel         `json:"rewrite_model"`
	RewriteQuery         bool                                     `json:"rewrite_query"`
	ScoreThreshold       float64                                  `json:"score_threshold"`
	SourceParams         InstanceReadResponseSourceParams         `json:"source_params,nullable"`
	Status               string                                   `json:"status"`
	TokenID              string                                   `json:"token_id" format:"uuid"`
	JSON                 instanceReadResponseJSON                 `json:"-"`
}

// instanceReadResponseJSON contains the JSON metadata for the struct
// [InstanceReadResponse]
type instanceReadResponseJSON struct {
	ID                   apijson.Field
	CreatedAt            apijson.Field
	ModifiedAt           apijson.Field
	Source               apijson.Field
	Type                 apijson.Field
	VectorizeName        apijson.Field
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
	FusionMethod         apijson.Field
	HybridSearchEnabled  apijson.Field
	LastActivity         apijson.Field
	MaxNumResults        apijson.Field
	Metadata             apijson.Field
	ModifiedBy           apijson.Field
	Paused               apijson.Field
	PublicEndpointID     apijson.Field
	PublicEndpointParams apijson.Field
	Reranking            apijson.Field
	RerankingModel       apijson.Field
	RetrievalOptions     apijson.Field
	RewriteModel         apijson.Field
	RewriteQuery         apijson.Field
	ScoreThreshold       apijson.Field
	SourceParams         apijson.Field
	Status               apijson.Field
	TokenID              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *InstanceReadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseType string

const (
	InstanceReadResponseTypeR2         InstanceReadResponseType = "r2"
	InstanceReadResponseTypeWebCrawler InstanceReadResponseType = "web-crawler"
)

func (r InstanceReadResponseType) IsKnown() bool {
	switch r {
	case InstanceReadResponseTypeR2, InstanceReadResponseTypeWebCrawler:
		return true
	}
	return false
}

type InstanceReadResponseAISearchModel string

const (
	InstanceReadResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceReadResponseAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceReadResponseAISearchModelCfZaiOrgGlm4_7Flash                   InstanceReadResponseAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceReadResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceReadResponseAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceReadResponseAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceReadResponseAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceReadResponseAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceReadResponseAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceReadResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceReadResponseAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceReadResponseAISearchModelCfMoonshotaiKimiK2Instruct            InstanceReadResponseAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceReadResponseAISearchModelCfGoogleGemma3_12bIt                  InstanceReadResponseAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceReadResponseAISearchModelAnthropicClaude3_7Sonnet              InstanceReadResponseAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceReadResponseAISearchModelAnthropicClaudeSonnet4                InstanceReadResponseAISearchModel = "anthropic/claude-sonnet-4"
	InstanceReadResponseAISearchModelAnthropicClaudeOpus4                  InstanceReadResponseAISearchModel = "anthropic/claude-opus-4"
	InstanceReadResponseAISearchModelAnthropicClaude3_5Haiku               InstanceReadResponseAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceReadResponseAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceReadResponseAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceReadResponseAISearchModelCerebrasQwen3_235bA22bThinking        InstanceReadResponseAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceReadResponseAISearchModelCerebrasLlama3_3_70b                  InstanceReadResponseAISearchModel = "cerebras/llama-3.3-70b"
	InstanceReadResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceReadResponseAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceReadResponseAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceReadResponseAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceReadResponseAISearchModelCerebrasGptOSs120b                    InstanceReadResponseAISearchModel = "cerebras/gpt-oss-120b"
	InstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Flash          InstanceReadResponseAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Pro            InstanceReadResponseAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceReadResponseAISearchModelGrokGrok4                             InstanceReadResponseAISearchModel = "grok/grok-4"
	InstanceReadResponseAISearchModelGroqLlama3_3_70bVersatile             InstanceReadResponseAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceReadResponseAISearchModelGroqLlama3_1_8bInstant                InstanceReadResponseAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceReadResponseAISearchModelOpenAIGpt5                            InstanceReadResponseAISearchModel = "openai/gpt-5"
	InstanceReadResponseAISearchModelOpenAIGpt5Mini                        InstanceReadResponseAISearchModel = "openai/gpt-5-mini"
	InstanceReadResponseAISearchModelOpenAIGpt5Nano                        InstanceReadResponseAISearchModel = "openai/gpt-5-nano"
	InstanceReadResponseAISearchModelEmpty                                 InstanceReadResponseAISearchModel = ""
)

func (r InstanceReadResponseAISearchModel) IsKnown() bool {
	switch r {
	case InstanceReadResponseAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceReadResponseAISearchModelCfZaiOrgGlm4_7Flash, InstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceReadResponseAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceReadResponseAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceReadResponseAISearchModelCfQwenQwen3_30bA3bFp8, InstanceReadResponseAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceReadResponseAISearchModelCfMoonshotaiKimiK2Instruct, InstanceReadResponseAISearchModelCfGoogleGemma3_12bIt, InstanceReadResponseAISearchModelAnthropicClaude3_7Sonnet, InstanceReadResponseAISearchModelAnthropicClaudeSonnet4, InstanceReadResponseAISearchModelAnthropicClaudeOpus4, InstanceReadResponseAISearchModelAnthropicClaude3_5Haiku, InstanceReadResponseAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceReadResponseAISearchModelCerebrasQwen3_235bA22bThinking, InstanceReadResponseAISearchModelCerebrasLlama3_3_70b, InstanceReadResponseAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceReadResponseAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceReadResponseAISearchModelCerebrasGptOSs120b, InstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Flash, InstanceReadResponseAISearchModelGoogleAIStudioGemini2_5Pro, InstanceReadResponseAISearchModelGrokGrok4, InstanceReadResponseAISearchModelGroqLlama3_3_70bVersatile, InstanceReadResponseAISearchModelGroqLlama3_1_8bInstant, InstanceReadResponseAISearchModelOpenAIGpt5, InstanceReadResponseAISearchModelOpenAIGpt5Mini, InstanceReadResponseAISearchModelOpenAIGpt5Nano, InstanceReadResponseAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceReadResponseCacheThreshold string

const (
	InstanceReadResponseCacheThresholdSuperStrictMatch InstanceReadResponseCacheThreshold = "super_strict_match"
	InstanceReadResponseCacheThresholdCloseEnough      InstanceReadResponseCacheThreshold = "close_enough"
	InstanceReadResponseCacheThresholdFlexibleFriend   InstanceReadResponseCacheThreshold = "flexible_friend"
	InstanceReadResponseCacheThresholdAnythingGoes     InstanceReadResponseCacheThreshold = "anything_goes"
)

func (r InstanceReadResponseCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceReadResponseCacheThresholdSuperStrictMatch, InstanceReadResponseCacheThresholdCloseEnough, InstanceReadResponseCacheThresholdFlexibleFriend, InstanceReadResponseCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceReadResponseCustomMetadata struct {
	DataType  InstanceReadResponseCustomMetadataDataType `json:"data_type,required"`
	FieldName string                                     `json:"field_name,required"`
	JSON      instanceReadResponseCustomMetadataJSON     `json:"-"`
}

// instanceReadResponseCustomMetadataJSON contains the JSON metadata for the struct
// [InstanceReadResponseCustomMetadata]
type instanceReadResponseCustomMetadataJSON struct {
	DataType    apijson.Field
	FieldName   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseCustomMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseCustomMetadataDataType string

const (
	InstanceReadResponseCustomMetadataDataTypeText    InstanceReadResponseCustomMetadataDataType = "text"
	InstanceReadResponseCustomMetadataDataTypeNumber  InstanceReadResponseCustomMetadataDataType = "number"
	InstanceReadResponseCustomMetadataDataTypeBoolean InstanceReadResponseCustomMetadataDataType = "boolean"
)

func (r InstanceReadResponseCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceReadResponseCustomMetadataDataTypeText, InstanceReadResponseCustomMetadataDataTypeNumber, InstanceReadResponseCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceReadResponseEmbeddingModel string

const (
	InstanceReadResponseEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceReadResponseEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceReadResponseEmbeddingModelCfBaaiBgeM3                      InstanceReadResponseEmbeddingModel = "@cf/baai/bge-m3"
	InstanceReadResponseEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceReadResponseEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceReadResponseEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceReadResponseEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceReadResponseEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceReadResponseEmbeddingModelOpenAITextEmbedding3Small        InstanceReadResponseEmbeddingModel = "openai/text-embedding-3-small"
	InstanceReadResponseEmbeddingModelOpenAITextEmbedding3Large        InstanceReadResponseEmbeddingModel = "openai/text-embedding-3-large"
	InstanceReadResponseEmbeddingModelEmpty                            InstanceReadResponseEmbeddingModel = ""
)

func (r InstanceReadResponseEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceReadResponseEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceReadResponseEmbeddingModelCfBaaiBgeM3, InstanceReadResponseEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceReadResponseEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceReadResponseEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceReadResponseEmbeddingModelOpenAITextEmbedding3Small, InstanceReadResponseEmbeddingModelOpenAITextEmbedding3Large, InstanceReadResponseEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceReadResponseFusionMethod string

const (
	InstanceReadResponseFusionMethodMax InstanceReadResponseFusionMethod = "max"
	InstanceReadResponseFusionMethodRrf InstanceReadResponseFusionMethod = "rrf"
)

func (r InstanceReadResponseFusionMethod) IsKnown() bool {
	switch r {
	case InstanceReadResponseFusionMethodMax, InstanceReadResponseFusionMethodRrf:
		return true
	}
	return false
}

type InstanceReadResponseMetadata struct {
	CreatedFromAISearchWizard bool                             `json:"created_from_aisearch_wizard"`
	WorkerDomain              string                           `json:"worker_domain"`
	JSON                      instanceReadResponseMetadataJSON `json:"-"`
}

// instanceReadResponseMetadataJSON contains the JSON metadata for the struct
// [InstanceReadResponseMetadata]
type instanceReadResponseMetadataJSON struct {
	CreatedFromAISearchWizard apijson.Field
	WorkerDomain              apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *InstanceReadResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponsePublicEndpointParams struct {
	AuthorizedHosts         []string                                                        `json:"authorized_hosts"`
	ChatCompletionsEndpoint InstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint `json:"chat_completions_endpoint"`
	Enabled                 bool                                                            `json:"enabled"`
	Mcp                     InstanceReadResponsePublicEndpointParamsMcp                     `json:"mcp"`
	RateLimit               InstanceReadResponsePublicEndpointParamsRateLimit               `json:"rate_limit"`
	SearchEndpoint          InstanceReadResponsePublicEndpointParamsSearchEndpoint          `json:"search_endpoint"`
	JSON                    instanceReadResponsePublicEndpointParamsJSON                    `json:"-"`
}

// instanceReadResponsePublicEndpointParamsJSON contains the JSON metadata for the
// struct [InstanceReadResponsePublicEndpointParams]
type instanceReadResponsePublicEndpointParamsJSON struct {
	AuthorizedHosts         apijson.Field
	ChatCompletionsEndpoint apijson.Field
	Enabled                 apijson.Field
	Mcp                     apijson.Field
	RateLimit               apijson.Field
	SearchEndpoint          apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *InstanceReadResponsePublicEndpointParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponsePublicEndpointParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled bool                                                                `json:"disabled"`
	JSON     instanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON `json:"-"`
}

// instanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON contains the
// JSON metadata for the struct
// [InstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint]
type instanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponsePublicEndpointParamsChatCompletionsEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponsePublicEndpointParamsChatCompletionsEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponsePublicEndpointParamsMcp struct {
	Description string `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled bool                                            `json:"disabled"`
	JSON     instanceReadResponsePublicEndpointParamsMcpJSON `json:"-"`
}

// instanceReadResponsePublicEndpointParamsMcpJSON contains the JSON metadata for
// the struct [InstanceReadResponsePublicEndpointParamsMcp]
type instanceReadResponsePublicEndpointParamsMcpJSON struct {
	Description apijson.Field
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponsePublicEndpointParamsMcp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponsePublicEndpointParamsMcpJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponsePublicEndpointParamsRateLimit struct {
	PeriodMs  int64                                                      `json:"period_ms"`
	Requests  int64                                                      `json:"requests"`
	Technique InstanceReadResponsePublicEndpointParamsRateLimitTechnique `json:"technique"`
	JSON      instanceReadResponsePublicEndpointParamsRateLimitJSON      `json:"-"`
}

// instanceReadResponsePublicEndpointParamsRateLimitJSON contains the JSON metadata
// for the struct [InstanceReadResponsePublicEndpointParamsRateLimit]
type instanceReadResponsePublicEndpointParamsRateLimitJSON struct {
	PeriodMs    apijson.Field
	Requests    apijson.Field
	Technique   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponsePublicEndpointParamsRateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponsePublicEndpointParamsRateLimitJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponsePublicEndpointParamsRateLimitTechnique string

const (
	InstanceReadResponsePublicEndpointParamsRateLimitTechniqueFixed   InstanceReadResponsePublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceReadResponsePublicEndpointParamsRateLimitTechniqueSliding InstanceReadResponsePublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceReadResponsePublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceReadResponsePublicEndpointParamsRateLimitTechniqueFixed, InstanceReadResponsePublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceReadResponsePublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled bool                                                       `json:"disabled"`
	JSON     instanceReadResponsePublicEndpointParamsSearchEndpointJSON `json:"-"`
}

// instanceReadResponsePublicEndpointParamsSearchEndpointJSON contains the JSON
// metadata for the struct [InstanceReadResponsePublicEndpointParamsSearchEndpoint]
type instanceReadResponsePublicEndpointParamsSearchEndpointJSON struct {
	Disabled    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponsePublicEndpointParamsSearchEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponsePublicEndpointParamsSearchEndpointJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseRerankingModel string

const (
	InstanceReadResponseRerankingModelCfBaaiBgeRerankerBase InstanceReadResponseRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceReadResponseRerankingModelEmpty                 InstanceReadResponseRerankingModel = ""
)

func (r InstanceReadResponseRerankingModel) IsKnown() bool {
	switch r {
	case InstanceReadResponseRerankingModelCfBaaiBgeRerankerBase, InstanceReadResponseRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceReadResponseRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode InstanceReadResponseRetrievalOptionsKeywordMatchMode `json:"keyword_match_mode"`
	JSON             instanceReadResponseRetrievalOptionsJSON             `json:"-"`
}

// instanceReadResponseRetrievalOptionsJSON contains the JSON metadata for the
// struct [InstanceReadResponseRetrievalOptions]
type instanceReadResponseRetrievalOptionsJSON struct {
	KeywordMatchMode apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InstanceReadResponseRetrievalOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseRetrievalOptionsJSON) RawJSON() string {
	return r.raw
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceReadResponseRetrievalOptionsKeywordMatchMode string

const (
	InstanceReadResponseRetrievalOptionsKeywordMatchModeExactMatch InstanceReadResponseRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceReadResponseRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceReadResponseRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceReadResponseRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceReadResponseRetrievalOptionsKeywordMatchModeExactMatch, InstanceReadResponseRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceReadResponseRewriteModel string

const (
	InstanceReadResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceReadResponseRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceReadResponseRewriteModelCfZaiOrgGlm4_7Flash                   InstanceReadResponseRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceReadResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceReadResponseRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceReadResponseRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceReadResponseRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceReadResponseRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceReadResponseRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceReadResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceReadResponseRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceReadResponseRewriteModelCfMoonshotaiKimiK2Instruct            InstanceReadResponseRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceReadResponseRewriteModelCfGoogleGemma3_12bIt                  InstanceReadResponseRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceReadResponseRewriteModelAnthropicClaude3_7Sonnet              InstanceReadResponseRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceReadResponseRewriteModelAnthropicClaudeSonnet4                InstanceReadResponseRewriteModel = "anthropic/claude-sonnet-4"
	InstanceReadResponseRewriteModelAnthropicClaudeOpus4                  InstanceReadResponseRewriteModel = "anthropic/claude-opus-4"
	InstanceReadResponseRewriteModelAnthropicClaude3_5Haiku               InstanceReadResponseRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceReadResponseRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceReadResponseRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceReadResponseRewriteModelCerebrasQwen3_235bA22bThinking        InstanceReadResponseRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceReadResponseRewriteModelCerebrasLlama3_3_70b                  InstanceReadResponseRewriteModel = "cerebras/llama-3.3-70b"
	InstanceReadResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceReadResponseRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceReadResponseRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceReadResponseRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceReadResponseRewriteModelCerebrasGptOSs120b                    InstanceReadResponseRewriteModel = "cerebras/gpt-oss-120b"
	InstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Flash          InstanceReadResponseRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Pro            InstanceReadResponseRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceReadResponseRewriteModelGrokGrok4                             InstanceReadResponseRewriteModel = "grok/grok-4"
	InstanceReadResponseRewriteModelGroqLlama3_3_70bVersatile             InstanceReadResponseRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceReadResponseRewriteModelGroqLlama3_1_8bInstant                InstanceReadResponseRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceReadResponseRewriteModelOpenAIGpt5                            InstanceReadResponseRewriteModel = "openai/gpt-5"
	InstanceReadResponseRewriteModelOpenAIGpt5Mini                        InstanceReadResponseRewriteModel = "openai/gpt-5-mini"
	InstanceReadResponseRewriteModelOpenAIGpt5Nano                        InstanceReadResponseRewriteModel = "openai/gpt-5-nano"
	InstanceReadResponseRewriteModelEmpty                                 InstanceReadResponseRewriteModel = ""
)

func (r InstanceReadResponseRewriteModel) IsKnown() bool {
	switch r {
	case InstanceReadResponseRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceReadResponseRewriteModelCfZaiOrgGlm4_7Flash, InstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceReadResponseRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceReadResponseRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceReadResponseRewriteModelCfQwenQwen3_30bA3bFp8, InstanceReadResponseRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceReadResponseRewriteModelCfMoonshotaiKimiK2Instruct, InstanceReadResponseRewriteModelCfGoogleGemma3_12bIt, InstanceReadResponseRewriteModelAnthropicClaude3_7Sonnet, InstanceReadResponseRewriteModelAnthropicClaudeSonnet4, InstanceReadResponseRewriteModelAnthropicClaudeOpus4, InstanceReadResponseRewriteModelAnthropicClaude3_5Haiku, InstanceReadResponseRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceReadResponseRewriteModelCerebrasQwen3_235bA22bThinking, InstanceReadResponseRewriteModelCerebrasLlama3_3_70b, InstanceReadResponseRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceReadResponseRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceReadResponseRewriteModelCerebrasGptOSs120b, InstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Flash, InstanceReadResponseRewriteModelGoogleAIStudioGemini2_5Pro, InstanceReadResponseRewriteModelGrokGrok4, InstanceReadResponseRewriteModelGroqLlama3_3_70bVersatile, InstanceReadResponseRewriteModelGroqLlama3_1_8bInstant, InstanceReadResponseRewriteModelOpenAIGpt5, InstanceReadResponseRewriteModelOpenAIGpt5Mini, InstanceReadResponseRewriteModelOpenAIGpt5Nano, InstanceReadResponseRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceReadResponseSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems []string `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   []string                                   `json:"include_items"`
	Prefix         string                                     `json:"prefix"`
	R2Jurisdiction string                                     `json:"r2_jurisdiction"`
	WebCrawler     InstanceReadResponseSourceParamsWebCrawler `json:"web_crawler"`
	JSON           instanceReadResponseSourceParamsJSON       `json:"-"`
}

// instanceReadResponseSourceParamsJSON contains the JSON metadata for the struct
// [InstanceReadResponseSourceParams]
type instanceReadResponseSourceParamsJSON struct {
	ExcludeItems   apijson.Field
	IncludeItems   apijson.Field
	Prefix         apijson.Field
	R2Jurisdiction apijson.Field
	WebCrawler     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseSourceParamsWebCrawler struct {
	ParseOptions InstanceReadResponseSourceParamsWebCrawlerParseOptions `json:"parse_options"`
	ParseType    InstanceReadResponseSourceParamsWebCrawlerParseType    `json:"parse_type"`
	StoreOptions InstanceReadResponseSourceParamsWebCrawlerStoreOptions `json:"store_options"`
	JSON         instanceReadResponseSourceParamsWebCrawlerJSON         `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerJSON contains the JSON metadata for
// the struct [InstanceReadResponseSourceParamsWebCrawler]
type instanceReadResponseSourceParamsWebCrawlerJSON struct {
	ParseOptions apijson.Field
	ParseType    apijson.Field
	StoreOptions apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawler) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders map[string]string `json:"include_headers"`
	IncludeImages  bool              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    []string                                                   `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering bool                                                       `json:"use_browser_rendering"`
	JSON                instanceReadResponseSourceParamsWebCrawlerParseOptionsJSON `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerParseOptionsJSON contains the JSON
// metadata for the struct [InstanceReadResponseSourceParamsWebCrawlerParseOptions]
type instanceReadResponseSourceParamsWebCrawlerParseOptionsJSON struct {
	IncludeHeaders      apijson.Field
	IncludeImages       apijson.Field
	SpecificSitemaps    apijson.Field
	UseBrowserRendering apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawlerParseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerParseOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceReadResponseSourceParamsWebCrawlerParseType string

const (
	InstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap InstanceReadResponseSourceParamsWebCrawlerParseType = "sitemap"
	InstanceReadResponseSourceParamsWebCrawlerParseTypeFeedRss InstanceReadResponseSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceReadResponseSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceReadResponseSourceParamsWebCrawlerParseTypeSitemap, InstanceReadResponseSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceReadResponseSourceParamsWebCrawlerStoreOptions struct {
	StorageID      string                                                     `json:"storage_id,required"`
	R2Jurisdiction string                                                     `json:"r2_jurisdiction"`
	StorageType    r2.Provider                                                `json:"storage_type"`
	JSON           instanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON `json:"-"`
}

// instanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON contains the JSON
// metadata for the struct [InstanceReadResponseSourceParamsWebCrawlerStoreOptions]
type instanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON struct {
	StorageID      apijson.Field
	R2Jurisdiction apijson.Field
	StorageType    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceReadResponseSourceParamsWebCrawlerStoreOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseSourceParamsWebCrawlerStoreOptionsJSON) RawJSON() string {
	return r.raw
}

type InstanceSearchResponse struct {
	Chunks      []InstanceSearchResponseChunk `json:"chunks,required"`
	SearchQuery string                        `json:"search_query,required"`
	JSON        instanceSearchResponseJSON    `json:"-"`
}

// instanceSearchResponseJSON contains the JSON metadata for the struct
// [InstanceSearchResponse]
type instanceSearchResponseJSON struct {
	Chunks      apijson.Field
	SearchQuery apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceSearchResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceSearchResponseChunk struct {
	ID             string                                     `json:"id,required"`
	Score          float64                                    `json:"score,required"`
	Text           string                                     `json:"text,required"`
	Type           string                                     `json:"type,required"`
	Item           InstanceSearchResponseChunksItem           `json:"item"`
	ScoringDetails InstanceSearchResponseChunksScoringDetails `json:"scoring_details"`
	JSON           instanceSearchResponseChunkJSON            `json:"-"`
}

// instanceSearchResponseChunkJSON contains the JSON metadata for the struct
// [InstanceSearchResponseChunk]
type instanceSearchResponseChunkJSON struct {
	ID             apijson.Field
	Score          apijson.Field
	Text           apijson.Field
	Type           apijson.Field
	Item           apijson.Field
	ScoringDetails apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceSearchResponseChunk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceSearchResponseChunkJSON) RawJSON() string {
	return r.raw
}

type InstanceSearchResponseChunksItem struct {
	Key       string                               `json:"key,required"`
	Metadata  map[string]interface{}               `json:"metadata"`
	Timestamp float64                              `json:"timestamp"`
	JSON      instanceSearchResponseChunksItemJSON `json:"-"`
}

// instanceSearchResponseChunksItemJSON contains the JSON metadata for the struct
// [InstanceSearchResponseChunksItem]
type instanceSearchResponseChunksItemJSON struct {
	Key         apijson.Field
	Metadata    apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceSearchResponseChunksItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceSearchResponseChunksItemJSON) RawJSON() string {
	return r.raw
}

type InstanceSearchResponseChunksScoringDetails struct {
	KeywordRank    float64                                        `json:"keyword_rank"`
	KeywordScore   float64                                        `json:"keyword_score"`
	RerankingScore float64                                        `json:"reranking_score"`
	VectorRank     float64                                        `json:"vector_rank"`
	VectorScore    float64                                        `json:"vector_score"`
	JSON           instanceSearchResponseChunksScoringDetailsJSON `json:"-"`
}

// instanceSearchResponseChunksScoringDetailsJSON contains the JSON metadata for
// the struct [InstanceSearchResponseChunksScoringDetails]
type instanceSearchResponseChunksScoringDetailsJSON struct {
	KeywordRank    apijson.Field
	KeywordScore   apijson.Field
	RerankingScore apijson.Field
	VectorRank     apijson.Field
	VectorScore    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InstanceSearchResponseChunksScoringDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceSearchResponseChunksScoringDetailsJSON) RawJSON() string {
	return r.raw
}

type InstanceStatsResponse struct {
	Completed         int64                     `json:"completed"`
	Error             int64                     `json:"error"`
	FileEmbedErrors   map[string]interface{}    `json:"file_embed_errors"`
	IndexSourceErrors map[string]interface{}    `json:"index_source_errors"`
	LastActivity      time.Time                 `json:"last_activity" format:"date-time"`
	Queued            int64                     `json:"queued"`
	Running           int64                     `json:"running"`
	Skipped           int64                     `json:"skipped"`
	JSON              instanceStatsResponseJSON `json:"-"`
}

// instanceStatsResponseJSON contains the JSON metadata for the struct
// [InstanceStatsResponse]
type instanceStatsResponseJSON struct {
	Completed         apijson.Field
	Error             apijson.Field
	FileEmbedErrors   apijson.Field
	IndexSourceErrors apijson.Field
	LastActivity      apijson.Field
	Queued            apijson.Field
	Running           apijson.Field
	Skipped           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *InstanceStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Use your AI Search ID.
	ID                   param.Field[string]                                `json:"id,required"`
	Source               param.Field[string]                                `json:"source,required"`
	Type                 param.Field[InstanceNewParamsType]                 `json:"type,required"`
	AIGatewayID          param.Field[string]                                `json:"ai_gateway_id"`
	AISearchModel        param.Field[InstanceNewParamsAISearchModel]        `json:"ai_search_model"`
	Chunk                param.Field[bool]                                  `json:"chunk"`
	ChunkOverlap         param.Field[int64]                                 `json:"chunk_overlap"`
	ChunkSize            param.Field[int64]                                 `json:"chunk_size"`
	CustomMetadata       param.Field[[]InstanceNewParamsCustomMetadata]     `json:"custom_metadata"`
	EmbeddingModel       param.Field[InstanceNewParamsEmbeddingModel]       `json:"embedding_model"`
	FusionMethod         param.Field[InstanceNewParamsFusionMethod]         `json:"fusion_method"`
	HybridSearchEnabled  param.Field[bool]                                  `json:"hybrid_search_enabled"`
	MaxNumResults        param.Field[int64]                                 `json:"max_num_results"`
	Metadata             param.Field[InstanceNewParamsMetadata]             `json:"metadata"`
	PublicEndpointParams param.Field[InstanceNewParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking            param.Field[bool]                                  `json:"reranking"`
	RerankingModel       param.Field[InstanceNewParamsRerankingModel]       `json:"reranking_model"`
	RetrievalOptions     param.Field[InstanceNewParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel         param.Field[InstanceNewParamsRewriteModel]         `json:"rewrite_model"`
	RewriteQuery         param.Field[bool]                                  `json:"rewrite_query"`
	ScoreThreshold       param.Field[float64]                               `json:"score_threshold"`
	SourceParams         param.Field[InstanceNewParamsSourceParams]         `json:"source_params"`
	TokenID              param.Field[string]                                `json:"token_id" format:"uuid"`
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsType string

const (
	InstanceNewParamsTypeR2         InstanceNewParamsType = "r2"
	InstanceNewParamsTypeWebCrawler InstanceNewParamsType = "web-crawler"
)

func (r InstanceNewParamsType) IsKnown() bool {
	switch r {
	case InstanceNewParamsTypeR2, InstanceNewParamsTypeWebCrawler:
		return true
	}
	return false
}

type InstanceNewParamsAISearchModel string

const (
	InstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceNewParamsAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceNewParamsAISearchModelCfZaiOrgGlm4_7Flash                   InstanceNewParamsAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceNewParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceNewParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceNewParamsAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceNewParamsAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceNewParamsAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceNewParamsAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceNewParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceNewParamsAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceNewParamsAISearchModelCfMoonshotaiKimiK2Instruct            InstanceNewParamsAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceNewParamsAISearchModelCfGoogleGemma3_12bIt                  InstanceNewParamsAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceNewParamsAISearchModelAnthropicClaude3_7Sonnet              InstanceNewParamsAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceNewParamsAISearchModelAnthropicClaudeSonnet4                InstanceNewParamsAISearchModel = "anthropic/claude-sonnet-4"
	InstanceNewParamsAISearchModelAnthropicClaudeOpus4                  InstanceNewParamsAISearchModel = "anthropic/claude-opus-4"
	InstanceNewParamsAISearchModelAnthropicClaude3_5Haiku               InstanceNewParamsAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceNewParamsAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceNewParamsAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceNewParamsAISearchModelCerebrasQwen3_235bA22bThinking        InstanceNewParamsAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceNewParamsAISearchModelCerebrasLlama3_3_70b                  InstanceNewParamsAISearchModel = "cerebras/llama-3.3-70b"
	InstanceNewParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceNewParamsAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceNewParamsAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceNewParamsAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceNewParamsAISearchModelCerebrasGptOSs120b                    InstanceNewParamsAISearchModel = "cerebras/gpt-oss-120b"
	InstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Flash          InstanceNewParamsAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Pro            InstanceNewParamsAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceNewParamsAISearchModelGrokGrok4                             InstanceNewParamsAISearchModel = "grok/grok-4"
	InstanceNewParamsAISearchModelGroqLlama3_3_70bVersatile             InstanceNewParamsAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceNewParamsAISearchModelGroqLlama3_1_8bInstant                InstanceNewParamsAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceNewParamsAISearchModelOpenAIGpt5                            InstanceNewParamsAISearchModel = "openai/gpt-5"
	InstanceNewParamsAISearchModelOpenAIGpt5Mini                        InstanceNewParamsAISearchModel = "openai/gpt-5-mini"
	InstanceNewParamsAISearchModelOpenAIGpt5Nano                        InstanceNewParamsAISearchModel = "openai/gpt-5-nano"
	InstanceNewParamsAISearchModelEmpty                                 InstanceNewParamsAISearchModel = ""
)

func (r InstanceNewParamsAISearchModel) IsKnown() bool {
	switch r {
	case InstanceNewParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceNewParamsAISearchModelCfZaiOrgGlm4_7Flash, InstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceNewParamsAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceNewParamsAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceNewParamsAISearchModelCfQwenQwen3_30bA3bFp8, InstanceNewParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceNewParamsAISearchModelCfMoonshotaiKimiK2Instruct, InstanceNewParamsAISearchModelCfGoogleGemma3_12bIt, InstanceNewParamsAISearchModelAnthropicClaude3_7Sonnet, InstanceNewParamsAISearchModelAnthropicClaudeSonnet4, InstanceNewParamsAISearchModelAnthropicClaudeOpus4, InstanceNewParamsAISearchModelAnthropicClaude3_5Haiku, InstanceNewParamsAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceNewParamsAISearchModelCerebrasQwen3_235bA22bThinking, InstanceNewParamsAISearchModelCerebrasLlama3_3_70b, InstanceNewParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceNewParamsAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceNewParamsAISearchModelCerebrasGptOSs120b, InstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Flash, InstanceNewParamsAISearchModelGoogleAIStudioGemini2_5Pro, InstanceNewParamsAISearchModelGrokGrok4, InstanceNewParamsAISearchModelGroqLlama3_3_70bVersatile, InstanceNewParamsAISearchModelGroqLlama3_1_8bInstant, InstanceNewParamsAISearchModelOpenAIGpt5, InstanceNewParamsAISearchModelOpenAIGpt5Mini, InstanceNewParamsAISearchModelOpenAIGpt5Nano, InstanceNewParamsAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceNewParamsCustomMetadata struct {
	DataType  param.Field[InstanceNewParamsCustomMetadataDataType] `json:"data_type,required"`
	FieldName param.Field[string]                                  `json:"field_name,required"`
}

func (r InstanceNewParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsCustomMetadataDataType string

const (
	InstanceNewParamsCustomMetadataDataTypeText    InstanceNewParamsCustomMetadataDataType = "text"
	InstanceNewParamsCustomMetadataDataTypeNumber  InstanceNewParamsCustomMetadataDataType = "number"
	InstanceNewParamsCustomMetadataDataTypeBoolean InstanceNewParamsCustomMetadataDataType = "boolean"
)

func (r InstanceNewParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceNewParamsCustomMetadataDataTypeText, InstanceNewParamsCustomMetadataDataTypeNumber, InstanceNewParamsCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceNewParamsEmbeddingModel string

const (
	InstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceNewParamsEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceNewParamsEmbeddingModelCfBaaiBgeM3                      InstanceNewParamsEmbeddingModel = "@cf/baai/bge-m3"
	InstanceNewParamsEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceNewParamsEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceNewParamsEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceNewParamsEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceNewParamsEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceNewParamsEmbeddingModelOpenAITextEmbedding3Small        InstanceNewParamsEmbeddingModel = "openai/text-embedding-3-small"
	InstanceNewParamsEmbeddingModelOpenAITextEmbedding3Large        InstanceNewParamsEmbeddingModel = "openai/text-embedding-3-large"
	InstanceNewParamsEmbeddingModelEmpty                            InstanceNewParamsEmbeddingModel = ""
)

func (r InstanceNewParamsEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceNewParamsEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceNewParamsEmbeddingModelCfBaaiBgeM3, InstanceNewParamsEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceNewParamsEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceNewParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceNewParamsEmbeddingModelOpenAITextEmbedding3Small, InstanceNewParamsEmbeddingModelOpenAITextEmbedding3Large, InstanceNewParamsEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceNewParamsFusionMethod string

const (
	InstanceNewParamsFusionMethodMax InstanceNewParamsFusionMethod = "max"
	InstanceNewParamsFusionMethodRrf InstanceNewParamsFusionMethod = "rrf"
)

func (r InstanceNewParamsFusionMethod) IsKnown() bool {
	switch r {
	case InstanceNewParamsFusionMethodMax, InstanceNewParamsFusionMethodRrf:
		return true
	}
	return false
}

type InstanceNewParamsMetadata struct {
	CreatedFromAISearchWizard param.Field[bool]   `json:"created_from_aisearch_wizard"`
	WorkerDomain              param.Field[string] `json:"worker_domain"`
}

func (r InstanceNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                     `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[InstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
	Enabled                 param.Field[bool]                                                         `json:"enabled"`
	Mcp                     param.Field[InstanceNewParamsPublicEndpointParamsMcp]                     `json:"mcp"`
	RateLimit               param.Field[InstanceNewParamsPublicEndpointParamsRateLimit]               `json:"rate_limit"`
	SearchEndpoint          param.Field[InstanceNewParamsPublicEndpointParamsSearchEndpoint]          `json:"search_endpoint"`
}

func (r InstanceNewParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceNewParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceNewParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                   `json:"period_ms"`
	Requests  param.Field[int64]                                                   `json:"requests"`
	Technique param.Field[InstanceNewParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r InstanceNewParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsPublicEndpointParamsRateLimitTechnique string

const (
	InstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed   InstanceNewParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceNewParamsPublicEndpointParamsRateLimitTechniqueSliding InstanceNewParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceNewParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceNewParamsPublicEndpointParamsRateLimitTechniqueFixed, InstanceNewParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceNewParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceNewParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsRerankingModel string

const (
	InstanceNewParamsRerankingModelCfBaaiBgeRerankerBase InstanceNewParamsRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceNewParamsRerankingModelEmpty                 InstanceNewParamsRerankingModel = ""
)

func (r InstanceNewParamsRerankingModel) IsKnown() bool {
	switch r {
	case InstanceNewParamsRerankingModelCfBaaiBgeRerankerBase, InstanceNewParamsRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceNewParamsRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode param.Field[InstanceNewParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r InstanceNewParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceNewParamsRetrievalOptionsKeywordMatchMode string

const (
	InstanceNewParamsRetrievalOptionsKeywordMatchModeExactMatch InstanceNewParamsRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceNewParamsRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceNewParamsRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceNewParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceNewParamsRetrievalOptionsKeywordMatchModeExactMatch, InstanceNewParamsRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceNewParamsRewriteModel string

const (
	InstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceNewParamsRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceNewParamsRewriteModelCfZaiOrgGlm4_7Flash                   InstanceNewParamsRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceNewParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceNewParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceNewParamsRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceNewParamsRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceNewParamsRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceNewParamsRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceNewParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceNewParamsRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceNewParamsRewriteModelCfMoonshotaiKimiK2Instruct            InstanceNewParamsRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceNewParamsRewriteModelCfGoogleGemma3_12bIt                  InstanceNewParamsRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceNewParamsRewriteModelAnthropicClaude3_7Sonnet              InstanceNewParamsRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceNewParamsRewriteModelAnthropicClaudeSonnet4                InstanceNewParamsRewriteModel = "anthropic/claude-sonnet-4"
	InstanceNewParamsRewriteModelAnthropicClaudeOpus4                  InstanceNewParamsRewriteModel = "anthropic/claude-opus-4"
	InstanceNewParamsRewriteModelAnthropicClaude3_5Haiku               InstanceNewParamsRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceNewParamsRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceNewParamsRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceNewParamsRewriteModelCerebrasQwen3_235bA22bThinking        InstanceNewParamsRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceNewParamsRewriteModelCerebrasLlama3_3_70b                  InstanceNewParamsRewriteModel = "cerebras/llama-3.3-70b"
	InstanceNewParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceNewParamsRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceNewParamsRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceNewParamsRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceNewParamsRewriteModelCerebrasGptOSs120b                    InstanceNewParamsRewriteModel = "cerebras/gpt-oss-120b"
	InstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Flash          InstanceNewParamsRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Pro            InstanceNewParamsRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceNewParamsRewriteModelGrokGrok4                             InstanceNewParamsRewriteModel = "grok/grok-4"
	InstanceNewParamsRewriteModelGroqLlama3_3_70bVersatile             InstanceNewParamsRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceNewParamsRewriteModelGroqLlama3_1_8bInstant                InstanceNewParamsRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceNewParamsRewriteModelOpenAIGpt5                            InstanceNewParamsRewriteModel = "openai/gpt-5"
	InstanceNewParamsRewriteModelOpenAIGpt5Mini                        InstanceNewParamsRewriteModel = "openai/gpt-5-mini"
	InstanceNewParamsRewriteModelOpenAIGpt5Nano                        InstanceNewParamsRewriteModel = "openai/gpt-5-nano"
	InstanceNewParamsRewriteModelEmpty                                 InstanceNewParamsRewriteModel = ""
)

func (r InstanceNewParamsRewriteModel) IsKnown() bool {
	switch r {
	case InstanceNewParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceNewParamsRewriteModelCfZaiOrgGlm4_7Flash, InstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceNewParamsRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceNewParamsRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceNewParamsRewriteModelCfQwenQwen3_30bA3bFp8, InstanceNewParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceNewParamsRewriteModelCfMoonshotaiKimiK2Instruct, InstanceNewParamsRewriteModelCfGoogleGemma3_12bIt, InstanceNewParamsRewriteModelAnthropicClaude3_7Sonnet, InstanceNewParamsRewriteModelAnthropicClaudeSonnet4, InstanceNewParamsRewriteModelAnthropicClaudeOpus4, InstanceNewParamsRewriteModelAnthropicClaude3_5Haiku, InstanceNewParamsRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceNewParamsRewriteModelCerebrasQwen3_235bA22bThinking, InstanceNewParamsRewriteModelCerebrasLlama3_3_70b, InstanceNewParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceNewParamsRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceNewParamsRewriteModelCerebrasGptOSs120b, InstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Flash, InstanceNewParamsRewriteModelGoogleAIStudioGemini2_5Pro, InstanceNewParamsRewriteModelGrokGrok4, InstanceNewParamsRewriteModelGroqLlama3_3_70bVersatile, InstanceNewParamsRewriteModelGroqLlama3_1_8bInstant, InstanceNewParamsRewriteModelOpenAIGpt5, InstanceNewParamsRewriteModelOpenAIGpt5Mini, InstanceNewParamsRewriteModelOpenAIGpt5Nano, InstanceNewParamsRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceNewParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   param.Field[[]string]                                `json:"include_items"`
	Prefix         param.Field[string]                                  `json:"prefix"`
	R2Jurisdiction param.Field[string]                                  `json:"r2_jurisdiction"`
	WebCrawler     param.Field[InstanceNewParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r InstanceNewParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsSourceParamsWebCrawler struct {
	ParseOptions param.Field[InstanceNewParamsSourceParamsWebCrawlerParseOptions] `json:"parse_options"`
	ParseType    param.Field[InstanceNewParamsSourceParamsWebCrawlerParseType]    `json:"parse_type"`
	StoreOptions param.Field[InstanceNewParamsSourceParamsWebCrawlerStoreOptions] `json:"store_options"`
}

func (r InstanceNewParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders param.Field[map[string]string] `json:"include_headers"`
	IncludeImages  param.Field[bool]              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    param.Field[[]string] `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering param.Field[bool]     `json:"use_browser_rendering"`
}

func (r InstanceNewParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewParamsSourceParamsWebCrawlerParseType string

const (
	InstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap InstanceNewParamsSourceParamsWebCrawlerParseType = "sitemap"
	InstanceNewParamsSourceParamsWebCrawlerParseTypeFeedRss InstanceNewParamsSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceNewParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceNewParamsSourceParamsWebCrawlerParseTypeSitemap, InstanceNewParamsSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceNewParamsSourceParamsWebCrawlerStoreOptions struct {
	StorageID      param.Field[string]      `json:"storage_id,required"`
	R2Jurisdiction param.Field[string]      `json:"r2_jurisdiction"`
	StorageType    param.Field[r2.Provider] `json:"storage_type"`
}

func (r InstanceNewParamsSourceParamsWebCrawlerStoreOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewResponseEnvelope struct {
	Result  InstanceNewResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    instanceNewResponseEnvelopeJSON `json:"-"`
}

// instanceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceNewResponseEnvelope]
type instanceNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceUpdateParams struct {
	AccountID                      param.Field[string]                                   `path:"account_id,required"`
	AIGatewayID                    param.Field[string]                                   `json:"ai_gateway_id"`
	AISearchModel                  param.Field[InstanceUpdateParamsAISearchModel]        `json:"ai_search_model"`
	Cache                          param.Field[bool]                                     `json:"cache"`
	CacheThreshold                 param.Field[InstanceUpdateParamsCacheThreshold]       `json:"cache_threshold"`
	Chunk                          param.Field[bool]                                     `json:"chunk"`
	ChunkOverlap                   param.Field[int64]                                    `json:"chunk_overlap"`
	ChunkSize                      param.Field[int64]                                    `json:"chunk_size"`
	CustomMetadata                 param.Field[[]InstanceUpdateParamsCustomMetadata]     `json:"custom_metadata"`
	EmbeddingModel                 param.Field[InstanceUpdateParamsEmbeddingModel]       `json:"embedding_model"`
	FusionMethod                   param.Field[InstanceUpdateParamsFusionMethod]         `json:"fusion_method"`
	HybridSearchEnabled            param.Field[bool]                                     `json:"hybrid_search_enabled"`
	MaxNumResults                  param.Field[int64]                                    `json:"max_num_results"`
	Metadata                       param.Field[InstanceUpdateParamsMetadata]             `json:"metadata"`
	Paused                         param.Field[bool]                                     `json:"paused"`
	PublicEndpointParams           param.Field[InstanceUpdateParamsPublicEndpointParams] `json:"public_endpoint_params"`
	Reranking                      param.Field[bool]                                     `json:"reranking"`
	RerankingModel                 param.Field[InstanceUpdateParamsRerankingModel]       `json:"reranking_model"`
	RetrievalOptions               param.Field[InstanceUpdateParamsRetrievalOptions]     `json:"retrieval_options"`
	RewriteModel                   param.Field[InstanceUpdateParamsRewriteModel]         `json:"rewrite_model"`
	RewriteQuery                   param.Field[bool]                                     `json:"rewrite_query"`
	ScoreThreshold                 param.Field[float64]                                  `json:"score_threshold"`
	SourceParams                   param.Field[InstanceUpdateParamsSourceParams]         `json:"source_params"`
	Summarization                  param.Field[bool]                                     `json:"summarization"`
	SummarizationModel             param.Field[InstanceUpdateParamsSummarizationModel]   `json:"summarization_model"`
	SystemPromptAISearch           param.Field[string]                                   `json:"system_prompt_ai_search"`
	SystemPromptIndexSummarization param.Field[string]                                   `json:"system_prompt_index_summarization"`
	SystemPromptRewriteQuery       param.Field[string]                                   `json:"system_prompt_rewrite_query"`
	TokenID                        param.Field[string]                                   `json:"token_id" format:"uuid"`
}

func (r InstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsAISearchModel string

const (
	InstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceUpdateParamsAISearchModelCfZaiOrgGlm4_7Flash                   InstanceUpdateParamsAISearchModel = "@cf/zai-org/glm-4.7-flash"
	InstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFast         InstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFp8          InstanceUpdateParamsAISearchModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceUpdateParamsAISearchModelCfMetaLlama4Scout17b16eInstruct       InstanceUpdateParamsAISearchModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsAISearchModelCfQwenQwen3_30bA3bFp8                 InstanceUpdateParamsAISearchModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceUpdateParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceUpdateParamsAISearchModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2Instruct            InstanceUpdateParamsAISearchModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceUpdateParamsAISearchModelCfGoogleGemma3_12bIt                  InstanceUpdateParamsAISearchModel = "@cf/google/gemma-3-12b-it"
	InstanceUpdateParamsAISearchModelAnthropicClaude3_7Sonnet              InstanceUpdateParamsAISearchModel = "anthropic/claude-3-7-sonnet"
	InstanceUpdateParamsAISearchModelAnthropicClaudeSonnet4                InstanceUpdateParamsAISearchModel = "anthropic/claude-sonnet-4"
	InstanceUpdateParamsAISearchModelAnthropicClaudeOpus4                  InstanceUpdateParamsAISearchModel = "anthropic/claude-opus-4"
	InstanceUpdateParamsAISearchModelAnthropicClaude3_5Haiku               InstanceUpdateParamsAISearchModel = "anthropic/claude-3-5-haiku"
	InstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bInstruct        InstanceUpdateParamsAISearchModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bThinking        InstanceUpdateParamsAISearchModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceUpdateParamsAISearchModelCerebrasLlama3_3_70b                  InstanceUpdateParamsAISearchModel = "cerebras/llama-3.3-70b"
	InstanceUpdateParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct InstanceUpdateParamsAISearchModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceUpdateParamsAISearchModelCerebrasLlama4Scout17b16eInstruct     InstanceUpdateParamsAISearchModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsAISearchModelCerebrasGptOSs120b                    InstanceUpdateParamsAISearchModel = "cerebras/gpt-oss-120b"
	InstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Flash          InstanceUpdateParamsAISearchModel = "google-ai-studio/gemini-2.5-flash"
	InstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Pro            InstanceUpdateParamsAISearchModel = "google-ai-studio/gemini-2.5-pro"
	InstanceUpdateParamsAISearchModelGrokGrok4                             InstanceUpdateParamsAISearchModel = "grok/grok-4"
	InstanceUpdateParamsAISearchModelGroqLlama3_3_70bVersatile             InstanceUpdateParamsAISearchModel = "groq/llama-3.3-70b-versatile"
	InstanceUpdateParamsAISearchModelGroqLlama3_1_8bInstant                InstanceUpdateParamsAISearchModel = "groq/llama-3.1-8b-instant"
	InstanceUpdateParamsAISearchModelOpenAIGpt5                            InstanceUpdateParamsAISearchModel = "openai/gpt-5"
	InstanceUpdateParamsAISearchModelOpenAIGpt5Mini                        InstanceUpdateParamsAISearchModel = "openai/gpt-5-mini"
	InstanceUpdateParamsAISearchModelOpenAIGpt5Nano                        InstanceUpdateParamsAISearchModel = "openai/gpt-5-nano"
	InstanceUpdateParamsAISearchModelEmpty                                 InstanceUpdateParamsAISearchModel = ""
)

func (r InstanceUpdateParamsAISearchModel) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsAISearchModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceUpdateParamsAISearchModelCfZaiOrgGlm4_7Flash, InstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFast, InstanceUpdateParamsAISearchModelCfMetaLlama3_1_8bInstructFp8, InstanceUpdateParamsAISearchModelCfMetaLlama4Scout17b16eInstruct, InstanceUpdateParamsAISearchModelCfQwenQwen3_30bA3bFp8, InstanceUpdateParamsAISearchModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceUpdateParamsAISearchModelCfMoonshotaiKimiK2Instruct, InstanceUpdateParamsAISearchModelCfGoogleGemma3_12bIt, InstanceUpdateParamsAISearchModelAnthropicClaude3_7Sonnet, InstanceUpdateParamsAISearchModelAnthropicClaudeSonnet4, InstanceUpdateParamsAISearchModelAnthropicClaudeOpus4, InstanceUpdateParamsAISearchModelAnthropicClaude3_5Haiku, InstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bInstruct, InstanceUpdateParamsAISearchModelCerebrasQwen3_235bA22bThinking, InstanceUpdateParamsAISearchModelCerebrasLlama3_3_70b, InstanceUpdateParamsAISearchModelCerebrasLlama4Maverick17b128eInstruct, InstanceUpdateParamsAISearchModelCerebrasLlama4Scout17b16eInstruct, InstanceUpdateParamsAISearchModelCerebrasGptOSs120b, InstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Flash, InstanceUpdateParamsAISearchModelGoogleAIStudioGemini2_5Pro, InstanceUpdateParamsAISearchModelGrokGrok4, InstanceUpdateParamsAISearchModelGroqLlama3_3_70bVersatile, InstanceUpdateParamsAISearchModelGroqLlama3_1_8bInstant, InstanceUpdateParamsAISearchModelOpenAIGpt5, InstanceUpdateParamsAISearchModelOpenAIGpt5Mini, InstanceUpdateParamsAISearchModelOpenAIGpt5Nano, InstanceUpdateParamsAISearchModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateParamsCacheThreshold string

const (
	InstanceUpdateParamsCacheThresholdSuperStrictMatch InstanceUpdateParamsCacheThreshold = "super_strict_match"
	InstanceUpdateParamsCacheThresholdCloseEnough      InstanceUpdateParamsCacheThreshold = "close_enough"
	InstanceUpdateParamsCacheThresholdFlexibleFriend   InstanceUpdateParamsCacheThreshold = "flexible_friend"
	InstanceUpdateParamsCacheThresholdAnythingGoes     InstanceUpdateParamsCacheThreshold = "anything_goes"
)

func (r InstanceUpdateParamsCacheThreshold) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsCacheThresholdSuperStrictMatch, InstanceUpdateParamsCacheThresholdCloseEnough, InstanceUpdateParamsCacheThresholdFlexibleFriend, InstanceUpdateParamsCacheThresholdAnythingGoes:
		return true
	}
	return false
}

type InstanceUpdateParamsCustomMetadata struct {
	DataType  param.Field[InstanceUpdateParamsCustomMetadataDataType] `json:"data_type,required"`
	FieldName param.Field[string]                                     `json:"field_name,required"`
}

func (r InstanceUpdateParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsCustomMetadataDataType string

const (
	InstanceUpdateParamsCustomMetadataDataTypeText    InstanceUpdateParamsCustomMetadataDataType = "text"
	InstanceUpdateParamsCustomMetadataDataTypeNumber  InstanceUpdateParamsCustomMetadataDataType = "number"
	InstanceUpdateParamsCustomMetadataDataTypeBoolean InstanceUpdateParamsCustomMetadataDataType = "boolean"
)

func (r InstanceUpdateParamsCustomMetadataDataType) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsCustomMetadataDataTypeText, InstanceUpdateParamsCustomMetadataDataTypeNumber, InstanceUpdateParamsCustomMetadataDataTypeBoolean:
		return true
	}
	return false
}

type InstanceUpdateParamsEmbeddingModel string

const (
	InstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b         InstanceUpdateParamsEmbeddingModel = "@cf/qwen/qwen3-embedding-0.6b"
	InstanceUpdateParamsEmbeddingModelCfBaaiBgeM3                      InstanceUpdateParamsEmbeddingModel = "@cf/baai/bge-m3"
	InstanceUpdateParamsEmbeddingModelCfBaaiBgeLargeEnV1_5             InstanceUpdateParamsEmbeddingModel = "@cf/baai/bge-large-en-v1.5"
	InstanceUpdateParamsEmbeddingModelCfGoogleEmbeddinggemma300m       InstanceUpdateParamsEmbeddingModel = "@cf/google/embeddinggemma-300m"
	InstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001 InstanceUpdateParamsEmbeddingModel = "google-ai-studio/gemini-embedding-001"
	InstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Small        InstanceUpdateParamsEmbeddingModel = "openai/text-embedding-3-small"
	InstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Large        InstanceUpdateParamsEmbeddingModel = "openai/text-embedding-3-large"
	InstanceUpdateParamsEmbeddingModelEmpty                            InstanceUpdateParamsEmbeddingModel = ""
)

func (r InstanceUpdateParamsEmbeddingModel) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsEmbeddingModelCfQwenQwen3Embedding0_6b, InstanceUpdateParamsEmbeddingModelCfBaaiBgeM3, InstanceUpdateParamsEmbeddingModelCfBaaiBgeLargeEnV1_5, InstanceUpdateParamsEmbeddingModelCfGoogleEmbeddinggemma300m, InstanceUpdateParamsEmbeddingModelGoogleAIStudioGeminiEmbedding001, InstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Small, InstanceUpdateParamsEmbeddingModelOpenAITextEmbedding3Large, InstanceUpdateParamsEmbeddingModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateParamsFusionMethod string

const (
	InstanceUpdateParamsFusionMethodMax InstanceUpdateParamsFusionMethod = "max"
	InstanceUpdateParamsFusionMethodRrf InstanceUpdateParamsFusionMethod = "rrf"
)

func (r InstanceUpdateParamsFusionMethod) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsFusionMethodMax, InstanceUpdateParamsFusionMethodRrf:
		return true
	}
	return false
}

type InstanceUpdateParamsMetadata struct {
	CreatedFromAISearchWizard param.Field[bool]   `json:"created_from_aisearch_wizard"`
	WorkerDomain              param.Field[string] `json:"worker_domain"`
}

func (r InstanceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsPublicEndpointParams struct {
	AuthorizedHosts         param.Field[[]string]                                                        `json:"authorized_hosts"`
	ChatCompletionsEndpoint param.Field[InstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint] `json:"chat_completions_endpoint"`
	Enabled                 param.Field[bool]                                                            `json:"enabled"`
	Mcp                     param.Field[InstanceUpdateParamsPublicEndpointParamsMcp]                     `json:"mcp"`
	RateLimit               param.Field[InstanceUpdateParamsPublicEndpointParamsRateLimit]               `json:"rate_limit"`
	SearchEndpoint          param.Field[InstanceUpdateParamsPublicEndpointParamsSearchEndpoint]          `json:"search_endpoint"`
}

func (r InstanceUpdateParamsPublicEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint struct {
	// Disable chat completions endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceUpdateParamsPublicEndpointParamsChatCompletionsEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsPublicEndpointParamsMcp struct {
	Description param.Field[string] `json:"description"`
	// Disable MCP endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceUpdateParamsPublicEndpointParamsMcp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsPublicEndpointParamsRateLimit struct {
	PeriodMs  param.Field[int64]                                                      `json:"period_ms"`
	Requests  param.Field[int64]                                                      `json:"requests"`
	Technique param.Field[InstanceUpdateParamsPublicEndpointParamsRateLimitTechnique] `json:"technique"`
}

func (r InstanceUpdateParamsPublicEndpointParamsRateLimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsPublicEndpointParamsRateLimitTechnique string

const (
	InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed   InstanceUpdateParamsPublicEndpointParamsRateLimitTechnique = "fixed"
	InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding InstanceUpdateParamsPublicEndpointParamsRateLimitTechnique = "sliding"
)

func (r InstanceUpdateParamsPublicEndpointParamsRateLimitTechnique) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueFixed, InstanceUpdateParamsPublicEndpointParamsRateLimitTechniqueSliding:
		return true
	}
	return false
}

type InstanceUpdateParamsPublicEndpointParamsSearchEndpoint struct {
	// Disable search endpoint for this public endpoint
	Disabled param.Field[bool] `json:"disabled"`
}

func (r InstanceUpdateParamsPublicEndpointParamsSearchEndpoint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsRerankingModel string

const (
	InstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase InstanceUpdateParamsRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceUpdateParamsRerankingModelEmpty                 InstanceUpdateParamsRerankingModel = ""
)

func (r InstanceUpdateParamsRerankingModel) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsRerankingModelCfBaaiBgeRerankerBase, InstanceUpdateParamsRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateParamsRetrievalOptions struct {
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode param.Field[InstanceUpdateParamsRetrievalOptionsKeywordMatchMode] `json:"keyword_match_mode"`
}

func (r InstanceUpdateParamsRetrievalOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceUpdateParamsRetrievalOptionsKeywordMatchMode string

const (
	InstanceUpdateParamsRetrievalOptionsKeywordMatchModeExactMatch InstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "exact_match"
	InstanceUpdateParamsRetrievalOptionsKeywordMatchModeFuzzyMatch InstanceUpdateParamsRetrievalOptionsKeywordMatchMode = "fuzzy_match"
)

func (r InstanceUpdateParamsRetrievalOptionsKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsRetrievalOptionsKeywordMatchModeExactMatch, InstanceUpdateParamsRetrievalOptionsKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceUpdateParamsRewriteModel string

const (
	InstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceUpdateParamsRewriteModelCfZaiOrgGlm4_7Flash                   InstanceUpdateParamsRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceUpdateParamsRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceUpdateParamsRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceUpdateParamsRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceUpdateParamsRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceUpdateParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceUpdateParamsRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2Instruct            InstanceUpdateParamsRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceUpdateParamsRewriteModelCfGoogleGemma3_12bIt                  InstanceUpdateParamsRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceUpdateParamsRewriteModelAnthropicClaude3_7Sonnet              InstanceUpdateParamsRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceUpdateParamsRewriteModelAnthropicClaudeSonnet4                InstanceUpdateParamsRewriteModel = "anthropic/claude-sonnet-4"
	InstanceUpdateParamsRewriteModelAnthropicClaudeOpus4                  InstanceUpdateParamsRewriteModel = "anthropic/claude-opus-4"
	InstanceUpdateParamsRewriteModelAnthropicClaude3_5Haiku               InstanceUpdateParamsRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceUpdateParamsRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bThinking        InstanceUpdateParamsRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceUpdateParamsRewriteModelCerebrasLlama3_3_70b                  InstanceUpdateParamsRewriteModel = "cerebras/llama-3.3-70b"
	InstanceUpdateParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceUpdateParamsRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceUpdateParamsRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceUpdateParamsRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsRewriteModelCerebrasGptOSs120b                    InstanceUpdateParamsRewriteModel = "cerebras/gpt-oss-120b"
	InstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Flash          InstanceUpdateParamsRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Pro            InstanceUpdateParamsRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceUpdateParamsRewriteModelGrokGrok4                             InstanceUpdateParamsRewriteModel = "grok/grok-4"
	InstanceUpdateParamsRewriteModelGroqLlama3_3_70bVersatile             InstanceUpdateParamsRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceUpdateParamsRewriteModelGroqLlama3_1_8bInstant                InstanceUpdateParamsRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceUpdateParamsRewriteModelOpenAIGpt5                            InstanceUpdateParamsRewriteModel = "openai/gpt-5"
	InstanceUpdateParamsRewriteModelOpenAIGpt5Mini                        InstanceUpdateParamsRewriteModel = "openai/gpt-5-mini"
	InstanceUpdateParamsRewriteModelOpenAIGpt5Nano                        InstanceUpdateParamsRewriteModel = "openai/gpt-5-nano"
	InstanceUpdateParamsRewriteModelEmpty                                 InstanceUpdateParamsRewriteModel = ""
)

func (r InstanceUpdateParamsRewriteModel) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceUpdateParamsRewriteModelCfZaiOrgGlm4_7Flash, InstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceUpdateParamsRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceUpdateParamsRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceUpdateParamsRewriteModelCfQwenQwen3_30bA3bFp8, InstanceUpdateParamsRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceUpdateParamsRewriteModelCfMoonshotaiKimiK2Instruct, InstanceUpdateParamsRewriteModelCfGoogleGemma3_12bIt, InstanceUpdateParamsRewriteModelAnthropicClaude3_7Sonnet, InstanceUpdateParamsRewriteModelAnthropicClaudeSonnet4, InstanceUpdateParamsRewriteModelAnthropicClaudeOpus4, InstanceUpdateParamsRewriteModelAnthropicClaude3_5Haiku, InstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceUpdateParamsRewriteModelCerebrasQwen3_235bA22bThinking, InstanceUpdateParamsRewriteModelCerebrasLlama3_3_70b, InstanceUpdateParamsRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceUpdateParamsRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceUpdateParamsRewriteModelCerebrasGptOSs120b, InstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Flash, InstanceUpdateParamsRewriteModelGoogleAIStudioGemini2_5Pro, InstanceUpdateParamsRewriteModelGrokGrok4, InstanceUpdateParamsRewriteModelGroqLlama3_3_70bVersatile, InstanceUpdateParamsRewriteModelGroqLlama3_1_8bInstant, InstanceUpdateParamsRewriteModelOpenAIGpt5, InstanceUpdateParamsRewriteModelOpenAIGpt5Mini, InstanceUpdateParamsRewriteModelOpenAIGpt5Nano, InstanceUpdateParamsRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateParamsSourceParams struct {
	// List of path patterns to exclude. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /admin/** matches
	// /admin/users and /admin/settings/advanced)
	ExcludeItems param.Field[[]string] `json:"exclude_items"`
	// List of path patterns to include. Uses micromatch glob syntax: \* matches within
	// a path segment, ** matches across path segments (e.g., /blog/** matches
	// /blog/post and /blog/2024/post)
	IncludeItems   param.Field[[]string]                                   `json:"include_items"`
	Prefix         param.Field[string]                                     `json:"prefix"`
	R2Jurisdiction param.Field[string]                                     `json:"r2_jurisdiction"`
	WebCrawler     param.Field[InstanceUpdateParamsSourceParamsWebCrawler] `json:"web_crawler"`
}

func (r InstanceUpdateParamsSourceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSourceParamsWebCrawler struct {
	ParseOptions param.Field[InstanceUpdateParamsSourceParamsWebCrawlerParseOptions] `json:"parse_options"`
	ParseType    param.Field[InstanceUpdateParamsSourceParamsWebCrawlerParseType]    `json:"parse_type"`
	StoreOptions param.Field[InstanceUpdateParamsSourceParamsWebCrawlerStoreOptions] `json:"store_options"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawler) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSourceParamsWebCrawlerParseOptions struct {
	IncludeHeaders param.Field[map[string]string] `json:"include_headers"`
	IncludeImages  param.Field[bool]              `json:"include_images"`
	// List of specific sitemap URLs to use for crawling. Only valid when parse_type is
	// 'sitemap'.
	SpecificSitemaps    param.Field[[]string] `json:"specific_sitemaps" format:"uri"`
	UseBrowserRendering param.Field[bool]     `json:"use_browser_rendering"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawlerParseOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSourceParamsWebCrawlerParseType string

const (
	InstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap InstanceUpdateParamsSourceParamsWebCrawlerParseType = "sitemap"
	InstanceUpdateParamsSourceParamsWebCrawlerParseTypeFeedRss InstanceUpdateParamsSourceParamsWebCrawlerParseType = "feed-rss"
)

func (r InstanceUpdateParamsSourceParamsWebCrawlerParseType) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsSourceParamsWebCrawlerParseTypeSitemap, InstanceUpdateParamsSourceParamsWebCrawlerParseTypeFeedRss:
		return true
	}
	return false
}

type InstanceUpdateParamsSourceParamsWebCrawlerStoreOptions struct {
	StorageID      param.Field[string]      `json:"storage_id,required"`
	R2Jurisdiction param.Field[string]      `json:"r2_jurisdiction"`
	StorageType    param.Field[r2.Provider] `json:"storage_type"`
}

func (r InstanceUpdateParamsSourceParamsWebCrawlerStoreOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceUpdateParamsSummarizationModel string

const (
	InstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceUpdateParamsSummarizationModelCfZaiOrgGlm4_7Flash                   InstanceUpdateParamsSummarizationModel = "@cf/zai-org/glm-4.7-flash"
	InstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFast         InstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFp8          InstanceUpdateParamsSummarizationModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceUpdateParamsSummarizationModelCfMetaLlama4Scout17b16eInstruct       InstanceUpdateParamsSummarizationModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsSummarizationModelCfQwenQwen3_30bA3bFp8                 InstanceUpdateParamsSummarizationModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceUpdateParamsSummarizationModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceUpdateParamsSummarizationModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2Instruct            InstanceUpdateParamsSummarizationModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceUpdateParamsSummarizationModelCfGoogleGemma3_12bIt                  InstanceUpdateParamsSummarizationModel = "@cf/google/gemma-3-12b-it"
	InstanceUpdateParamsSummarizationModelAnthropicClaude3_7Sonnet              InstanceUpdateParamsSummarizationModel = "anthropic/claude-3-7-sonnet"
	InstanceUpdateParamsSummarizationModelAnthropicClaudeSonnet4                InstanceUpdateParamsSummarizationModel = "anthropic/claude-sonnet-4"
	InstanceUpdateParamsSummarizationModelAnthropicClaudeOpus4                  InstanceUpdateParamsSummarizationModel = "anthropic/claude-opus-4"
	InstanceUpdateParamsSummarizationModelAnthropicClaude3_5Haiku               InstanceUpdateParamsSummarizationModel = "anthropic/claude-3-5-haiku"
	InstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bInstruct        InstanceUpdateParamsSummarizationModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bThinking        InstanceUpdateParamsSummarizationModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceUpdateParamsSummarizationModelCerebrasLlama3_3_70b                  InstanceUpdateParamsSummarizationModel = "cerebras/llama-3.3-70b"
	InstanceUpdateParamsSummarizationModelCerebrasLlama4Maverick17b128eInstruct InstanceUpdateParamsSummarizationModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceUpdateParamsSummarizationModelCerebrasLlama4Scout17b16eInstruct     InstanceUpdateParamsSummarizationModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceUpdateParamsSummarizationModelCerebrasGptOSs120b                    InstanceUpdateParamsSummarizationModel = "cerebras/gpt-oss-120b"
	InstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Flash          InstanceUpdateParamsSummarizationModel = "google-ai-studio/gemini-2.5-flash"
	InstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Pro            InstanceUpdateParamsSummarizationModel = "google-ai-studio/gemini-2.5-pro"
	InstanceUpdateParamsSummarizationModelGrokGrok4                             InstanceUpdateParamsSummarizationModel = "grok/grok-4"
	InstanceUpdateParamsSummarizationModelGroqLlama3_3_70bVersatile             InstanceUpdateParamsSummarizationModel = "groq/llama-3.3-70b-versatile"
	InstanceUpdateParamsSummarizationModelGroqLlama3_1_8bInstant                InstanceUpdateParamsSummarizationModel = "groq/llama-3.1-8b-instant"
	InstanceUpdateParamsSummarizationModelOpenAIGpt5                            InstanceUpdateParamsSummarizationModel = "openai/gpt-5"
	InstanceUpdateParamsSummarizationModelOpenAIGpt5Mini                        InstanceUpdateParamsSummarizationModel = "openai/gpt-5-mini"
	InstanceUpdateParamsSummarizationModelOpenAIGpt5Nano                        InstanceUpdateParamsSummarizationModel = "openai/gpt-5-nano"
	InstanceUpdateParamsSummarizationModelEmpty                                 InstanceUpdateParamsSummarizationModel = ""
)

func (r InstanceUpdateParamsSummarizationModel) IsKnown() bool {
	switch r {
	case InstanceUpdateParamsSummarizationModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceUpdateParamsSummarizationModelCfZaiOrgGlm4_7Flash, InstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFast, InstanceUpdateParamsSummarizationModelCfMetaLlama3_1_8bInstructFp8, InstanceUpdateParamsSummarizationModelCfMetaLlama4Scout17b16eInstruct, InstanceUpdateParamsSummarizationModelCfQwenQwen3_30bA3bFp8, InstanceUpdateParamsSummarizationModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceUpdateParamsSummarizationModelCfMoonshotaiKimiK2Instruct, InstanceUpdateParamsSummarizationModelCfGoogleGemma3_12bIt, InstanceUpdateParamsSummarizationModelAnthropicClaude3_7Sonnet, InstanceUpdateParamsSummarizationModelAnthropicClaudeSonnet4, InstanceUpdateParamsSummarizationModelAnthropicClaudeOpus4, InstanceUpdateParamsSummarizationModelAnthropicClaude3_5Haiku, InstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bInstruct, InstanceUpdateParamsSummarizationModelCerebrasQwen3_235bA22bThinking, InstanceUpdateParamsSummarizationModelCerebrasLlama3_3_70b, InstanceUpdateParamsSummarizationModelCerebrasLlama4Maverick17b128eInstruct, InstanceUpdateParamsSummarizationModelCerebrasLlama4Scout17b16eInstruct, InstanceUpdateParamsSummarizationModelCerebrasGptOSs120b, InstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Flash, InstanceUpdateParamsSummarizationModelGoogleAIStudioGemini2_5Pro, InstanceUpdateParamsSummarizationModelGrokGrok4, InstanceUpdateParamsSummarizationModelGroqLlama3_3_70bVersatile, InstanceUpdateParamsSummarizationModelGroqLlama3_1_8bInstant, InstanceUpdateParamsSummarizationModelOpenAIGpt5, InstanceUpdateParamsSummarizationModelOpenAIGpt5Mini, InstanceUpdateParamsSummarizationModelOpenAIGpt5Nano, InstanceUpdateParamsSummarizationModelEmpty:
		return true
	}
	return false
}

type InstanceUpdateResponseEnvelope struct {
	Result  InstanceUpdateResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    instanceUpdateResponseEnvelopeJSON `json:"-"`
}

// instanceUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceUpdateResponseEnvelope]
type instanceUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceDeleteResponseEnvelope struct {
	Result  InstanceDeleteResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    instanceDeleteResponseEnvelopeJSON `json:"-"`
}

// instanceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceDeleteResponseEnvelope]
type instanceDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceChatCompletionsParams struct {
	AccountID       param.Field[string]                                       `path:"account_id,required"`
	Messages        param.Field[[]InstanceChatCompletionsParamsMessage]       `json:"messages,required"`
	AISearchOptions param.Field[InstanceChatCompletionsParamsAISearchOptions] `json:"ai_search_options"`
	Model           param.Field[InstanceChatCompletionsParamsModel]           `json:"model"`
	Stream          param.Field[bool]                                         `json:"stream"`
}

func (r InstanceChatCompletionsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsMessage struct {
	Content     param.Field[string]                                    `json:"content,required"`
	Role        param.Field[InstanceChatCompletionsParamsMessagesRole] `json:"role,required"`
	ExtraFields map[string]interface{}                                 `json:"-,extras"`
}

func (r InstanceChatCompletionsParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsMessagesRole string

const (
	InstanceChatCompletionsParamsMessagesRoleSystem    InstanceChatCompletionsParamsMessagesRole = "system"
	InstanceChatCompletionsParamsMessagesRoleDeveloper InstanceChatCompletionsParamsMessagesRole = "developer"
	InstanceChatCompletionsParamsMessagesRoleUser      InstanceChatCompletionsParamsMessagesRole = "user"
	InstanceChatCompletionsParamsMessagesRoleAssistant InstanceChatCompletionsParamsMessagesRole = "assistant"
	InstanceChatCompletionsParamsMessagesRoleTool      InstanceChatCompletionsParamsMessagesRole = "tool"
)

func (r InstanceChatCompletionsParamsMessagesRole) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsMessagesRoleSystem, InstanceChatCompletionsParamsMessagesRoleDeveloper, InstanceChatCompletionsParamsMessagesRoleUser, InstanceChatCompletionsParamsMessagesRoleAssistant, InstanceChatCompletionsParamsMessagesRoleTool:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsAISearchOptions struct {
	QueryRewrite param.Field[InstanceChatCompletionsParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[InstanceChatCompletionsParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r InstanceChatCompletionsParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]                                                          `json:"enabled"`
	Model         param.Field[InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                                        `json:"rewrite_prompt"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel string

const (
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4                             InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty                                 InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel = ""
)

func (r InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGrokGrok4, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, InstanceChatCompletionsParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                                       `json:"enabled"`
	MatchThreshold param.Field[float64]                                                    `json:"match_threshold"`
	Model          param.Field[InstanceChatCompletionsParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsRerankingModel string

const (
	InstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase InstanceChatCompletionsParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceChatCompletionsParamsAISearchOptionsRerankingModelEmpty                 InstanceChatCompletionsParamsAISearchOptionsRerankingModel = ""
)

func (r InstanceChatCompletionsParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, InstanceChatCompletionsParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsAISearchOptionsRetrieval struct {
	ContextExpansion param.Field[int64]                                                             `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                            `json:"filters"`
	FusionMethod     param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                               `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                                 `json:"max_num_results"`
	RetrievalType    param.Field[InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                                  `json:"return_on_failure"`
}

func (r InstanceChatCompletionsParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod string

const (
	InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "max"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodMax, InstanceChatCompletionsParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeExactMatch InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "exact_match"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeFuzzyMatch InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode = "fuzzy_match"
)

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeExactMatch, InstanceChatCompletionsParamsAISearchOptionsRetrievalKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType string

const (
	InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector  InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "vector"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid  InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeVector, InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeKeyword, InstanceChatCompletionsParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type InstanceChatCompletionsParamsModel string

const (
	InstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceChatCompletionsParamsModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash                   InstanceChatCompletionsParamsModel = "@cf/zai-org/glm-4.7-flash"
	InstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast         InstanceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8          InstanceChatCompletionsParamsModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct       InstanceChatCompletionsParamsModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8                 InstanceChatCompletionsParamsModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceChatCompletionsParamsModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct            InstanceChatCompletionsParamsModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceChatCompletionsParamsModelCfGoogleGemma3_12bIt                  InstanceChatCompletionsParamsModel = "@cf/google/gemma-3-12b-it"
	InstanceChatCompletionsParamsModelAnthropicClaude3_7Sonnet              InstanceChatCompletionsParamsModel = "anthropic/claude-3-7-sonnet"
	InstanceChatCompletionsParamsModelAnthropicClaudeSonnet4                InstanceChatCompletionsParamsModel = "anthropic/claude-sonnet-4"
	InstanceChatCompletionsParamsModelAnthropicClaudeOpus4                  InstanceChatCompletionsParamsModel = "anthropic/claude-opus-4"
	InstanceChatCompletionsParamsModelAnthropicClaude3_5Haiku               InstanceChatCompletionsParamsModel = "anthropic/claude-3-5-haiku"
	InstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct        InstanceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking        InstanceChatCompletionsParamsModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceChatCompletionsParamsModelCerebrasLlama3_3_70b                  InstanceChatCompletionsParamsModel = "cerebras/llama-3.3-70b"
	InstanceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct InstanceChatCompletionsParamsModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct     InstanceChatCompletionsParamsModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceChatCompletionsParamsModelCerebrasGptOSs120b                    InstanceChatCompletionsParamsModel = "cerebras/gpt-oss-120b"
	InstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash          InstanceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-flash"
	InstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro            InstanceChatCompletionsParamsModel = "google-ai-studio/gemini-2.5-pro"
	InstanceChatCompletionsParamsModelGrokGrok4                             InstanceChatCompletionsParamsModel = "grok/grok-4"
	InstanceChatCompletionsParamsModelGroqLlama3_3_70bVersatile             InstanceChatCompletionsParamsModel = "groq/llama-3.3-70b-versatile"
	InstanceChatCompletionsParamsModelGroqLlama3_1_8bInstant                InstanceChatCompletionsParamsModel = "groq/llama-3.1-8b-instant"
	InstanceChatCompletionsParamsModelOpenAIGpt5                            InstanceChatCompletionsParamsModel = "openai/gpt-5"
	InstanceChatCompletionsParamsModelOpenAIGpt5Mini                        InstanceChatCompletionsParamsModel = "openai/gpt-5-mini"
	InstanceChatCompletionsParamsModelOpenAIGpt5Nano                        InstanceChatCompletionsParamsModel = "openai/gpt-5-nano"
	InstanceChatCompletionsParamsModelEmpty                                 InstanceChatCompletionsParamsModel = ""
)

func (r InstanceChatCompletionsParamsModel) IsKnown() bool {
	switch r {
	case InstanceChatCompletionsParamsModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceChatCompletionsParamsModelCfZaiOrgGlm4_7Flash, InstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFast, InstanceChatCompletionsParamsModelCfMetaLlama3_1_8bInstructFp8, InstanceChatCompletionsParamsModelCfMetaLlama4Scout17b16eInstruct, InstanceChatCompletionsParamsModelCfQwenQwen3_30bA3bFp8, InstanceChatCompletionsParamsModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceChatCompletionsParamsModelCfMoonshotaiKimiK2Instruct, InstanceChatCompletionsParamsModelCfGoogleGemma3_12bIt, InstanceChatCompletionsParamsModelAnthropicClaude3_7Sonnet, InstanceChatCompletionsParamsModelAnthropicClaudeSonnet4, InstanceChatCompletionsParamsModelAnthropicClaudeOpus4, InstanceChatCompletionsParamsModelAnthropicClaude3_5Haiku, InstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bInstruct, InstanceChatCompletionsParamsModelCerebrasQwen3_235bA22bThinking, InstanceChatCompletionsParamsModelCerebrasLlama3_3_70b, InstanceChatCompletionsParamsModelCerebrasLlama4Maverick17b128eInstruct, InstanceChatCompletionsParamsModelCerebrasLlama4Scout17b16eInstruct, InstanceChatCompletionsParamsModelCerebrasGptOSs120b, InstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Flash, InstanceChatCompletionsParamsModelGoogleAIStudioGemini2_5Pro, InstanceChatCompletionsParamsModelGrokGrok4, InstanceChatCompletionsParamsModelGroqLlama3_3_70bVersatile, InstanceChatCompletionsParamsModelGroqLlama3_1_8bInstant, InstanceChatCompletionsParamsModelOpenAIGpt5, InstanceChatCompletionsParamsModelOpenAIGpt5Mini, InstanceChatCompletionsParamsModelOpenAIGpt5Nano, InstanceChatCompletionsParamsModelEmpty:
		return true
	}
	return false
}

type InstanceReadParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceReadResponseEnvelope struct {
	Result  InstanceReadResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    instanceReadResponseEnvelopeJSON `json:"-"`
}

// instanceReadResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceReadResponseEnvelope]
type instanceReadResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceReadResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceReadResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceSearchParams struct {
	AccountID       param.Field[string]                              `path:"account_id,required"`
	Messages        param.Field[[]InstanceSearchParamsMessage]       `json:"messages,required"`
	AISearchOptions param.Field[InstanceSearchParamsAISearchOptions] `json:"ai_search_options"`
}

func (r InstanceSearchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsMessage struct {
	Content     param.Field[string]                           `json:"content,required"`
	Role        param.Field[InstanceSearchParamsMessagesRole] `json:"role,required"`
	ExtraFields map[string]interface{}                        `json:"-,extras"`
}

func (r InstanceSearchParamsMessage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsMessagesRole string

const (
	InstanceSearchParamsMessagesRoleSystem    InstanceSearchParamsMessagesRole = "system"
	InstanceSearchParamsMessagesRoleDeveloper InstanceSearchParamsMessagesRole = "developer"
	InstanceSearchParamsMessagesRoleUser      InstanceSearchParamsMessagesRole = "user"
	InstanceSearchParamsMessagesRoleAssistant InstanceSearchParamsMessagesRole = "assistant"
	InstanceSearchParamsMessagesRoleTool      InstanceSearchParamsMessagesRole = "tool"
)

func (r InstanceSearchParamsMessagesRole) IsKnown() bool {
	switch r {
	case InstanceSearchParamsMessagesRoleSystem, InstanceSearchParamsMessagesRoleDeveloper, InstanceSearchParamsMessagesRoleUser, InstanceSearchParamsMessagesRoleAssistant, InstanceSearchParamsMessagesRoleTool:
		return true
	}
	return false
}

type InstanceSearchParamsAISearchOptions struct {
	QueryRewrite param.Field[InstanceSearchParamsAISearchOptionsQueryRewrite] `json:"query_rewrite"`
	Reranking    param.Field[InstanceSearchParamsAISearchOptionsReranking]    `json:"reranking"`
	Retrieval    param.Field[InstanceSearchParamsAISearchOptionsRetrieval]    `json:"retrieval"`
}

func (r InstanceSearchParamsAISearchOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsQueryRewrite struct {
	Enabled       param.Field[bool]                                                 `json:"enabled"`
	Model         param.Field[InstanceSearchParamsAISearchOptionsQueryRewriteModel] `json:"model"`
	RewritePrompt param.Field[string]                                               `json:"rewrite_prompt"`
}

func (r InstanceSearchParamsAISearchOptionsQueryRewrite) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsQueryRewriteModel string

const (
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast     InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.3-70b-instruct-fp8-fast"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash                   InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/zai-org/glm-4.7-flash"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast         InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fast"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8          InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-3.1-8b-instruct-fp8"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct       InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/meta/llama-4-scout-17b-16e-instruct"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8                 InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/qwen/qwen3-30b-a3b-fp8"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b  InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/deepseek-ai/deepseek-r1-distill-qwen-32b"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct            InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/moonshotai/kimi-k2-instruct"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt                  InstanceSearchParamsAISearchOptionsQueryRewriteModel = "@cf/google/gemma-3-12b-it"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet              InstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-7-sonnet"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4                InstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-sonnet-4"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4                  InstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-opus-4"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku               InstanceSearchParamsAISearchOptionsQueryRewriteModel = "anthropic/claude-3-5-haiku"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct        InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-instruct"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking        InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/qwen-3-235b-a22b-thinking"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b                  InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-3.3-70b"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-maverick-17b-128e-instruct"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct     InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/llama-4-scout-17b-16e-instruct"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b                    InstanceSearchParamsAISearchOptionsQueryRewriteModel = "cerebras/gpt-oss-120b"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash          InstanceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-flash"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro            InstanceSearchParamsAISearchOptionsQueryRewriteModel = "google-ai-studio/gemini-2.5-pro"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4                             InstanceSearchParamsAISearchOptionsQueryRewriteModel = "grok/grok-4"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile             InstanceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.3-70b-versatile"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant                InstanceSearchParamsAISearchOptionsQueryRewriteModel = "groq/llama-3.1-8b-instant"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5                            InstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini                        InstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-mini"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano                        InstanceSearchParamsAISearchOptionsQueryRewriteModel = "openai/gpt-5-nano"
	InstanceSearchParamsAISearchOptionsQueryRewriteModelEmpty                                 InstanceSearchParamsAISearchOptionsQueryRewriteModel = ""
)

func (r InstanceSearchParamsAISearchOptionsQueryRewriteModel) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_3_70bInstructFp8Fast, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfZaiOrgGlm4_7Flash, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFast, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama3_1_8bInstructFp8, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMetaLlama4Scout17b16eInstruct, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfQwenQwen3_30bA3bFp8, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfDeepseekAIDeepseekR1DistillQwen32b, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfMoonshotaiKimiK2Instruct, InstanceSearchParamsAISearchOptionsQueryRewriteModelCfGoogleGemma3_12bIt, InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_7Sonnet, InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeSonnet4, InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaudeOpus4, InstanceSearchParamsAISearchOptionsQueryRewriteModelAnthropicClaude3_5Haiku, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bInstruct, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasQwen3_235bA22bThinking, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama3_3_70b, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Maverick17b128eInstruct, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasLlama4Scout17b16eInstruct, InstanceSearchParamsAISearchOptionsQueryRewriteModelCerebrasGptOSs120b, InstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Flash, InstanceSearchParamsAISearchOptionsQueryRewriteModelGoogleAIStudioGemini2_5Pro, InstanceSearchParamsAISearchOptionsQueryRewriteModelGrokGrok4, InstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_3_70bVersatile, InstanceSearchParamsAISearchOptionsQueryRewriteModelGroqLlama3_1_8bInstant, InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5, InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Mini, InstanceSearchParamsAISearchOptionsQueryRewriteModelOpenAIGpt5Nano, InstanceSearchParamsAISearchOptionsQueryRewriteModelEmpty:
		return true
	}
	return false
}

type InstanceSearchParamsAISearchOptionsReranking struct {
	Enabled        param.Field[bool]                                              `json:"enabled"`
	MatchThreshold param.Field[float64]                                           `json:"match_threshold"`
	Model          param.Field[InstanceSearchParamsAISearchOptionsRerankingModel] `json:"model"`
}

func (r InstanceSearchParamsAISearchOptionsReranking) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsRerankingModel string

const (
	InstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase InstanceSearchParamsAISearchOptionsRerankingModel = "@cf/baai/bge-reranker-base"
	InstanceSearchParamsAISearchOptionsRerankingModelEmpty                 InstanceSearchParamsAISearchOptionsRerankingModel = ""
)

func (r InstanceSearchParamsAISearchOptionsRerankingModel) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRerankingModelCfBaaiBgeRerankerBase, InstanceSearchParamsAISearchOptionsRerankingModelEmpty:
		return true
	}
	return false
}

type InstanceSearchParamsAISearchOptionsRetrieval struct {
	ContextExpansion param.Field[int64]                                                    `json:"context_expansion"`
	Filters          param.Field[map[string]interface{}]                                   `json:"filters"`
	FusionMethod     param.Field[InstanceSearchParamsAISearchOptionsRetrievalFusionMethod] `json:"fusion_method"`
	// Controls how keyword search terms are matched. exact_match requires all terms to
	// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
	// exact_match.
	KeywordMatchMode param.Field[InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode] `json:"keyword_match_mode"`
	MatchThreshold   param.Field[float64]                                                      `json:"match_threshold"`
	MaxNumResults    param.Field[int64]                                                        `json:"max_num_results"`
	RetrievalType    param.Field[InstanceSearchParamsAISearchOptionsRetrievalRetrievalType]    `json:"retrieval_type"`
	ReturnOnFailure  param.Field[bool]                                                         `json:"return_on_failure"`
}

func (r InstanceSearchParamsAISearchOptionsRetrieval) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceSearchParamsAISearchOptionsRetrievalFusionMethod string

const (
	InstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax InstanceSearchParamsAISearchOptionsRetrievalFusionMethod = "max"
	InstanceSearchParamsAISearchOptionsRetrievalFusionMethodRrf InstanceSearchParamsAISearchOptionsRetrievalFusionMethod = "rrf"
)

func (r InstanceSearchParamsAISearchOptionsRetrievalFusionMethod) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRetrievalFusionMethodMax, InstanceSearchParamsAISearchOptionsRetrievalFusionMethodRrf:
		return true
	}
	return false
}

// Controls how keyword search terms are matched. exact_match requires all terms to
// appear (AND); fuzzy_match returns results containing any term (OR). Defaults to
// exact_match.
type InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode string

const (
	InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeExactMatch InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "exact_match"
	InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeFuzzyMatch InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode = "fuzzy_match"
)

func (r InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchMode) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeExactMatch, InstanceSearchParamsAISearchOptionsRetrievalKeywordMatchModeFuzzyMatch:
		return true
	}
	return false
}

type InstanceSearchParamsAISearchOptionsRetrievalRetrievalType string

const (
	InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector  InstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "vector"
	InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword InstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "keyword"
	InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid  InstanceSearchParamsAISearchOptionsRetrievalRetrievalType = "hybrid"
)

func (r InstanceSearchParamsAISearchOptionsRetrievalRetrievalType) IsKnown() bool {
	switch r {
	case InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeVector, InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeKeyword, InstanceSearchParamsAISearchOptionsRetrievalRetrievalTypeHybrid:
		return true
	}
	return false
}

type InstanceSearchResponseEnvelope struct {
	Result  InstanceSearchResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    instanceSearchResponseEnvelopeJSON `json:"-"`
}

// instanceSearchResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceSearchResponseEnvelope]
type instanceSearchResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceStatsParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceStatsResponseEnvelope struct {
	Result  InstanceStatsResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    instanceStatsResponseEnvelopeJSON `json:"-"`
}

// instanceStatsResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceStatsResponseEnvelope]
type instanceStatsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceStatsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceStatsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
